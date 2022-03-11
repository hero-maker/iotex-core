// Copyright (c) 2022 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package action

import (
	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/iotex-core/config"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
	"github.com/pkg/errors"
)

// Deserializer de-serializes an action
type Deserializer struct {
	withChainID bool
}

// ActionToSealedEnvelope converts protobuf to SealedEnvelope
func (ad *Deserializer) ActionToSealedEnvelope(pbAct *iotextypes.Action) (SealedEnvelope, error) {
	var (
		selp SealedEnvelope
		err  error
	)
	if ad.withChainID {
		err = selp.LoadProtoWithChainID(pbAct)
	} else {
		err = selp.LoadProto(pbAct)
	}
	return selp, err
}

// WithChainID sets whether or not to use chainID
func (ad *Deserializer) WithChainID(with bool) *Deserializer {
	ad.withChainID = with
	return ad
}

// LoadProtoWithChainID use chainID while loading protobuf
func (sealed *SealedEnvelope) LoadProtoWithChainID(pbAct *iotextypes.Action) error {
	if pbAct == nil {
		return ErrNilProto
	}
	if sealed == nil {
		return ErrNilAction
	}
	sigSize := len(pbAct.GetSignature())
	if sigSize != 65 {
		return errors.Errorf("invalid signature length = %d, expecting 65", sigSize)
	}

	var (
		core = pbAct.GetCore()
		elp  = &envelope{}
	)
	if err := elp.LoadProto(core); err != nil {
		return err
	}
	elp.chainID = core.GetChainID()

	// populate pubkey and signature
	srcPub, err := crypto.BytesToPublicKey(pbAct.GetSenderPubKey())
	if err != nil {
		return err
	}
	encoding := pbAct.GetEncoding()
	switch encoding {
	case iotextypes.Encoding_ETHEREUM_RLP:
		// verify action type can support RLP-encoding
		act, ok := elp.Action().(EthCompatibleAction)
		if !ok {
			return ErrInvalidAct
		}
		tx, err := act.ToEthTx()
		if err != nil {
			return err
		}
		if _, err = rlpSignedHash(tx, config.EVMNetworkID(), pbAct.GetSignature()); err != nil {
			return err
		}
		sealed.evmNetworkID = config.EVMNetworkID()
	case iotextypes.Encoding_IOTEX_PROTOBUF:
		break
	default:
		return errors.Errorf("unknown encoding type %v", encoding)
	}

	// clear 'sealed' and populate new value
	sealed.Envelope = elp
	sealed.srcPubkey = srcPub
	sealed.signature = make([]byte, sigSize)
	copy(sealed.signature, pbAct.GetSignature())
	sealed.encoding = encoding
	elp.Action().SetEnvelopeContext(*sealed)
	return nil
}
