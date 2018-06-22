// Code generated by idl2go from JSON generated by Barrister v0.1.6
package explorer

import (
	"fmt"
	"github.com/coopernurse/barrister-go"
	"reflect"
)

const BarristerVersion string = "0.1.6"
const BarristerChecksum string = "825d00b6df957edba7b31dbf47f9800f"
const BarristerDateGenerated int64 = 1529624748052000000

type CoinStatistic struct {
	Height    int64 `json:"height"`
	Supply    int64 `json:"supply"`
	Transfers int64 `json:"transfers"`
	Votes     int64 `json:"votes"`
	Aps       int64 `json:"aps"`
}

type BlockGenerator struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Block struct {
	ID         string         `json:"ID"`
	Height     int64          `json:"height"`
	Timestamp  int64          `json:"timestamp"`
	Transfers  int64          `json:"transfers"`
	Votes      int64          `json:"votes"`
	GenerateBy BlockGenerator `json:"generateBy"`
	Amount     int64          `json:"amount"`
	Forged     int64          `json:"forged"`
	Size       int64          `json:"size"`
}

type Transfer struct {
	ID        string `json:"ID"`
	Nounce    int64  `json:"nounce"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int64  `json:"amount"`
	Fee       int64  `json:"fee"`
	Timestamp int64  `json:"timestamp"`
	BlockID   string `json:"blockID"`
}

type Vote struct {
	ID        string `json:"ID"`
	Nounce    int64  `json:"nounce"`
	Timestamp int64  `json:"timestamp"`
	Voter     string `json:"voter"`
	Votee     string `json:"votee"`
	BlockID   string `json:"blockID"`
}

type AddressDetails struct {
	Address      string `json:"address"`
	TotalBalance int64  `json:"totalBalance"`
	Nonce        int64  `json:"nonce"`
}

type ConsensusMetrics struct {
	LatestEpoch         int64    `json:"latestEpoch"`
	LatestDelegates     []string `json:"latestDelegates"`
	LatestBlockProducer string   `json:"latestBlockProducer"`
}

type Explorer interface {
	GetBlockchainHeight() (int64, error)
	GetAddressBalance(address string) (int64, error)
	GetAddressDetails(address string) (AddressDetails, error)
	GetLastTransfersByRange(startBlockHeight int64, offset int64, limit int64, showCoinBase bool) ([]Transfer, error)
	GetTransferByID(transferID string) (Transfer, error)
	GetTransfersByAddress(address string, offset int64, limit int64) ([]Transfer, error)
	GetTransfersByBlockID(blkID string, offset int64, limit int64) ([]Transfer, error)
	GetLastVotesByRange(startBlockHeight int64, offset int64, limit int64) ([]Vote, error)
	GetVoteByID(voteID string) (Vote, error)
	GetVotesByAddress(address string, offset int64, limit int64) ([]Vote, error)
	GetVotesByBlockID(blkID string, offset int64, limit int64) ([]Vote, error)
	GetLastBlocksByRange(offset int64, limit int64) ([]Block, error)
	GetBlockByID(blkID string) (Block, error)
	GetCoinStatistic() (CoinStatistic, error)
	GetConsensusMetrics() (ConsensusMetrics, error)
}

func NewExplorerProxy(c barrister.Client) Explorer {
	return ExplorerProxy{c, barrister.MustParseIdlJson([]byte(IdlJsonRaw))}
}

type ExplorerProxy struct {
	client barrister.Client
	idl    *barrister.Idl
}

func (_p ExplorerProxy) GetBlockchainHeight() (int64, error) {
	_res, _err := _p.client.Call("Explorer.getBlockchainHeight")
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getBlockchainHeight").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(int64(0)), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(int64)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getBlockchainHeight returned invalid type: %v", _t)
			return int64(0), &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return int64(0), _err
}

func (_p ExplorerProxy) GetAddressBalance(address string) (int64, error) {
	_res, _err := _p.client.Call("Explorer.getAddressBalance", address)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getAddressBalance").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(int64(0)), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(int64)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getAddressBalance returned invalid type: %v", _t)
			return int64(0), &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return int64(0), _err
}

func (_p ExplorerProxy) GetAddressDetails(address string) (AddressDetails, error) {
	_res, _err := _p.client.Call("Explorer.getAddressDetails", address)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getAddressDetails").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(AddressDetails{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(AddressDetails)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getAddressDetails returned invalid type: %v", _t)
			return AddressDetails{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return AddressDetails{}, _err
}

func (_p ExplorerProxy) GetLastTransfersByRange(startBlockHeight int64, offset int64, limit int64, showCoinBase bool) ([]Transfer, error) {
	_res, _err := _p.client.Call("Explorer.getLastTransfersByRange", startBlockHeight, offset, limit, showCoinBase)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getLastTransfersByRange").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Transfer{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Transfer)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getLastTransfersByRange returned invalid type: %v", _t)
			return []Transfer{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Transfer{}, _err
}

func (_p ExplorerProxy) GetTransferByID(transferID string) (Transfer, error) {
	_res, _err := _p.client.Call("Explorer.getTransferByID", transferID)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getTransferByID").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(Transfer{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(Transfer)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getTransferByID returned invalid type: %v", _t)
			return Transfer{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return Transfer{}, _err
}

func (_p ExplorerProxy) GetTransfersByAddress(address string, offset int64, limit int64) ([]Transfer, error) {
	_res, _err := _p.client.Call("Explorer.getTransfersByAddress", address, offset, limit)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getTransfersByAddress").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Transfer{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Transfer)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getTransfersByAddress returned invalid type: %v", _t)
			return []Transfer{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Transfer{}, _err
}

func (_p ExplorerProxy) GetTransfersByBlockID(blkID string, offset int64, limit int64) ([]Transfer, error) {
	_res, _err := _p.client.Call("Explorer.getTransfersByBlockID", blkID, offset, limit)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getTransfersByBlockID").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Transfer{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Transfer)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getTransfersByBlockID returned invalid type: %v", _t)
			return []Transfer{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Transfer{}, _err
}

func (_p ExplorerProxy) GetLastVotesByRange(startBlockHeight int64, offset int64, limit int64) ([]Vote, error) {
	_res, _err := _p.client.Call("Explorer.getLastVotesByRange", startBlockHeight, offset, limit)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getLastVotesByRange").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Vote{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Vote)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getLastVotesByRange returned invalid type: %v", _t)
			return []Vote{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Vote{}, _err
}

func (_p ExplorerProxy) GetVoteByID(voteID string) (Vote, error) {
	_res, _err := _p.client.Call("Explorer.getVoteByID", voteID)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getVoteByID").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(Vote{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(Vote)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getVoteByID returned invalid type: %v", _t)
			return Vote{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return Vote{}, _err
}

func (_p ExplorerProxy) GetVotesByAddress(address string, offset int64, limit int64) ([]Vote, error) {
	_res, _err := _p.client.Call("Explorer.getVotesByAddress", address, offset, limit)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getVotesByAddress").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Vote{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Vote)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getVotesByAddress returned invalid type: %v", _t)
			return []Vote{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Vote{}, _err
}

func (_p ExplorerProxy) GetVotesByBlockID(blkID string, offset int64, limit int64) ([]Vote, error) {
	_res, _err := _p.client.Call("Explorer.getVotesByBlockID", blkID, offset, limit)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getVotesByBlockID").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Vote{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Vote)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getVotesByBlockID returned invalid type: %v", _t)
			return []Vote{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Vote{}, _err
}

func (_p ExplorerProxy) GetLastBlocksByRange(offset int64, limit int64) ([]Block, error) {
	_res, _err := _p.client.Call("Explorer.getLastBlocksByRange", offset, limit)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getLastBlocksByRange").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Block{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Block)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getLastBlocksByRange returned invalid type: %v", _t)
			return []Block{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Block{}, _err
}

func (_p ExplorerProxy) GetBlockByID(blkID string) (Block, error) {
	_res, _err := _p.client.Call("Explorer.getBlockByID", blkID)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getBlockByID").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(Block{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(Block)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getBlockByID returned invalid type: %v", _t)
			return Block{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return Block{}, _err
}

func (_p ExplorerProxy) GetCoinStatistic() (CoinStatistic, error) {
	_res, _err := _p.client.Call("Explorer.getCoinStatistic")
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getCoinStatistic").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(CoinStatistic{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(CoinStatistic)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getCoinStatistic returned invalid type: %v", _t)
			return CoinStatistic{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return CoinStatistic{}, _err
}

func (_p ExplorerProxy) GetConsensusMetrics() (ConsensusMetrics, error) {
	_res, _err := _p.client.Call("Explorer.getConsensusMetrics")
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getConsensusMetrics").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(ConsensusMetrics{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(ConsensusMetrics)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getConsensusMetrics returned invalid type: %v", _t)
			return ConsensusMetrics{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return ConsensusMetrics{}, _err
}

func NewJSONServer(idl *barrister.Idl, forceASCII bool, explorer Explorer) barrister.Server {
	return NewServer(idl, &barrister.JsonSerializer{forceASCII}, explorer)
}

func NewServer(idl *barrister.Idl, ser barrister.Serializer, explorer Explorer) barrister.Server {
	_svr := barrister.NewServer(idl, ser)
	_svr.AddHandler("Explorer", explorer)
	return _svr
}

var IdlJsonRaw = `[
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "Copyright (c) 2018 IoTeX\nThis is an alpha (internal) release and is not suitable for production. This source code is provided ‘as is’ and no\nwarranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent\npermitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache\nLicense 2.0 that can be found in the LICENSE file.",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "To compile this file:\n1. Install the barrister translator (IDL -\u003e JSON)\nyou need to be root (or use sudo)\npip install barrister",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "2. Install barrister-go\ngo get github.com/coopernurse/barrister-go\ngo install github.com/coopernurse/barrister-go/idl2go",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "3. barrister explorer.idl | $GOPATH/bin/idl2go -i -p explorer",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "CoinStatistic",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "height",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "supply",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "transfers",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "votes",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "aps",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "BlockGenerator",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "name",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "address",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "Block",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "ID",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "height",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "timestamp",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "transfers",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "votes",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "generateBy",
                "type": "BlockGenerator",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "amount",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "forged",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "size",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "Transfer",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "ID",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "nounce",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "sender",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "recipient",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "amount",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "fee",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "timestamp",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "blockID",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "Vote",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "ID",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "nounce",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "timestamp",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "voter",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "votee",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "blockID",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "AddressDetails",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "address",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "totalBalance",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "nonce",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "ConsensusMetrics",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "latestEpoch",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "latestDelegates",
                "type": "string",
                "optional": false,
                "is_array": true,
                "comment": ""
            },
            {
                "name": "latestBlockProducer",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "interface",
        "name": "Explorer",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": [
            {
                "name": "getBlockchainHeight",
                "comment": "get the blockchain tip height",
                "params": [],
                "returns": {
                    "name": "",
                    "type": "int",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getAddressBalance",
                "comment": "get the balance of an address",
                "params": [
                    {
                        "name": "address",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "int",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getAddressDetails",
                "comment": "get the address detail of an iotex address",
                "params": [
                    {
                        "name": "address",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "AddressDetails",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getLastTransfersByRange",
                "comment": "get list of transfers by start block height, transfer offset and limit",
                "params": [
                    {
                        "name": "startBlockHeight",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "offset",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "limit",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "showCoinBase",
                        "type": "bool",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Transfer",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getTransferByID",
                "comment": "get transfers from transaction id",
                "params": [
                    {
                        "name": "transferID",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Transfer",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getTransfersByAddress",
                "comment": "get list of transfer belong to an address",
                "params": [
                    {
                        "name": "address",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "offset",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "limit",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Transfer",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getTransfersByBlockID",
                "comment": "get all transfers in a block",
                "params": [
                    {
                        "name": "blkID",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "offset",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "limit",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Transfer",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getLastVotesByRange",
                "comment": "get list of votes by start block height, vote offset and limit",
                "params": [
                    {
                        "name": "startBlockHeight",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "offset",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "limit",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Vote",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getVoteByID",
                "comment": "get vote from vote id",
                "params": [
                    {
                        "name": "voteID",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Vote",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getVotesByAddress",
                "comment": "get list of vote belong to an address",
                "params": [
                    {
                        "name": "address",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "offset",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "limit",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Vote",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getVotesByBlockID",
                "comment": "get all votes in a block",
                "params": [
                    {
                        "name": "blkID",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "offset",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "limit",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Vote",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getLastBlocksByRange",
                "comment": "get list of blocks by block id offset and limit",
                "params": [
                    {
                        "name": "offset",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "limit",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Block",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getBlockByID",
                "comment": "get block by block id",
                "params": [
                    {
                        "name": "blkID",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Block",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getCoinStatistic",
                "comment": "get statistic of iotx",
                "params": [],
                "returns": {
                    "name": "",
                    "type": "CoinStatistic",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getConsensusMetrics",
                "comment": "get consensus metrics",
                "params": [],
                "returns": {
                    "name": "",
                    "type": "ConsensusMetrics",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            }
        ],
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "meta",
        "name": "",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "0.1.6",
        "date_generated": 1529624748052,
        "checksum": "825d00b6df957edba7b31dbf47f9800f"
    }
]`
