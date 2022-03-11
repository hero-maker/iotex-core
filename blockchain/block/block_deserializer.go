// Copyright (c) 2022 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package block

import (
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	"github.com/iotexproject/iotex-core/action"
)

// Deserializer de-serializes a block
type Deserializer struct {
}

// FromBlockProto converts protobuf to block
func (bd *Deserializer) FromBlockProto(pbBlock *iotextypes.Block) (*Block, error) {
	b := Block{}
	if err := b.Header.LoadFromBlockHeaderProto(pbBlock.GetHeader()); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize block header")
	}
	if err := b.Body.LoadProtoWithChainID(pbBlock.GetBody()); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize block body")
	}
	if err := b.ConvertFromBlockFooterPb(pbBlock.GetFooter()); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize block footer")
	}
	return &b, nil
}

// DeserializeBlock de-serializes a block
func (bd *Deserializer) DeserializeBlock(buf []byte) (*Block, error) {
	pbBlock := iotextypes.Block{}
	if err := proto.Unmarshal(buf, &pbBlock); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal block")
	}
	b, err := bd.FromBlockProto(&pbBlock)
	if err != nil {
		return nil, err
	}
	b.Receipts = nil
	if err = b.VerifyTxRoot(); err != nil {
		return nil, err
	}
	return b, nil
}

// FromBodyProto converts protobuf to body
func (bd *Deserializer) FromBodyProto(pbBody *iotextypes.BlockBody) (*Body, error) {
	b := Body{}
	if err := b.LoadProtoWithChainID(pbBody); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize block body")
	}
	return &b, nil
}

// DeserializeBody de-serializes a block body
func (bd *Deserializer) DeserializeBody(buf []byte) (*Body, error) {
	pb := iotextypes.BlockBody{}
	if err := proto.Unmarshal(buf, &pb); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal block body")
	}
	return bd.FromBodyProto(&pb)
}

// LoadProtoWithChainID loads body from proto
func (b *Body) LoadProtoWithChainID(pbBlock *iotextypes.BlockBody) error {
	b.Actions = []action.SealedEnvelope{}
	for _, actPb := range pbBlock.Actions {
		act := action.SealedEnvelope{}
		if err := act.LoadProtoWithChainID(actPb); err != nil {
			return err
		}
		b.Actions = append(b.Actions, act)
	}
	return nil
}
