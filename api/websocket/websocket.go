package websocket

import (
	"bytes"

	"github.com/nknorg/nkn/api/httprestful/common"
	Err "github.com/nknorg/nkn/api/httprestful/error"
	"github.com/nknorg/nkn/api/websocket/server"
	. "github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/core/ledger"
	"github.com/nknorg/nkn/events"
	. "github.com/nknorg/nkn/net/protocol"
	. "github.com/nknorg/nkn/util/config"
	"github.com/nknorg/nkn/vault"
)

var ws *server.WsServer

var (
	pushBlockFlag    bool = false
	pushRawBlockFlag bool = false
	pushBlockTxsFlag bool = false
)

func StartServer(n Noder, w vault.Wallet) {
	//	common.SetNode(n)
	ledger.DefaultLedger.Blockchain.BCEvents.Subscribe(events.EventBlockPersistCompleted, SendBlock2WSclient)
	go func() {
		ws = server.InitWsServer(n, w)
		ws.Start()
	}()
}

func SendBlock2WSclient(v interface{}) {
	go func() {
		PushSigChainBlockHash(v)
	}()
	if Parameters.HttpWsPort != 0 && pushBlockFlag {
		go func() {
			PushBlock(v)
		}()
	}
	if Parameters.HttpWsPort != 0 && pushBlockTxsFlag {
		go func() {
			PushBlockTransactions(v)
		}()
	}
}

func Stop() {
	if ws == nil {
		return
	}
	ws.Stop()
}

func ReStartServer() {
	// TODO
	//	if ws == nil {
	//		n := common.GetNode()
	//		ws = server.InitWsServer(n)
	//		ws.Start()
	//		return
	//	}
	//	ws.Restart()
}

func GetWsPushBlockFlag() bool {
	return pushBlockFlag
}

func SetWsPushBlockFlag(b bool) {
	pushBlockFlag = b
}

func GetPushRawBlockFlag() bool {
	return pushRawBlockFlag
}

func SetPushRawBlockFlag(b bool) {
	pushRawBlockFlag = b
}

func GetPushBlockTxsFlag() bool {
	return pushBlockTxsFlag
}

func SetPushBlockTxsFlag(b bool) {
	pushBlockTxsFlag = b
}

func SetTxHashMap(txhash string, sessionid string) {
	if ws == nil {
		return
	}
	ws.SetTxHashMap(txhash, sessionid)
}

func PushResult(txHash Uint256, errcode int64, action string, result interface{}) {
	if ws != nil {
		resp := common.ResponsePack(Err.SUCCESS)
		resp["Result"] = result
		resp["Error"] = errcode
		resp["Action"] = action
		resp["Desc"] = Err.ErrMap[resp["Error"].(int64)]
		ws.PushTxResult(BytesToHexString(txHash.ToArrayReverse()), resp)
	}
}

func PushSmartCodeInvokeResult(txHash Uint256, errcode int64, result interface{}) {
	if ws == nil {
		return
	}
	resp := common.ResponsePack(Err.SUCCESS)
	var Result = make(map[string]interface{})
	txHashStr := BytesToHexString(txHash.ToArray())
	Result["TxHash"] = txHashStr
	Result["ExecResult"] = result

	resp["Result"] = Result
	resp["Action"] = "sendsmartcodeinvoke"
	resp["Error"] = errcode
	resp["Desc"] = Err.ErrMap[errcode]
	ws.PushTxResult(txHashStr, resp)
}

func PushBlock(v interface{}) {
	if ws == nil {
		return
	}
	resp := common.ResponsePack(Err.SUCCESS)
	if block, ok := v.(*ledger.Block); ok {
		if pushRawBlockFlag {
			w := bytes.NewBuffer(nil)
			block.Serialize(w)
			resp["Result"] = BytesToHexString(w.Bytes())
		} else {
			resp["Result"] = common.GetBlockInfo(block)
		}
		resp["Action"] = "sendRawBlock"
		ws.PushResult(resp)
	}
}

func PushBlockTransactions(v interface{}) {
	if ws == nil {
		return
	}
	resp := common.ResponsePack(Err.SUCCESS)
	if block, ok := v.(*ledger.Block); ok {
		if pushBlockTxsFlag {
			resp["Result"] = common.GetBlockTransactions(block)
		}
		resp["Action"] = "sendblocktransactions"
		ws.PushResult(resp)
	}
}

func PushSigChainBlockHash(v interface{}) {
	if ws == nil {
		return
	}
	resp := common.ResponsePack(Err.SUCCESS)
	if block, ok := v.(*ledger.Block); ok {
		resp["Action"] = "updateSigChainBlockHash"
		resp["Result"] = common.GetBlockInfo(block).BlockData.PrevBlockHash
		ws.PushResult(resp)
	}
}

func GetServer() *server.WsServer {
	return ws
}