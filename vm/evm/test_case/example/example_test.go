package test

import (
	"testing"
	"nkn-core/vm/evm/test_case"
)

const (
	ABI = `[{"constant":false,"inputs":[],"name":"kill","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_newgreeting","type":"string"}],"name":"setGreeting","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"greet","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"},{"inputs":[{"name":"_greeting","type":"string"}],"payable":false,"type":"constructor"}]`
	BIN = `6060604052341561000c57fe5b6040516104cf3803806104cf833981016040528080518201919050505b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060019080519060200190610080929190610088565b505b5061012d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c957805160ff19168380011785556100f7565b828001600101855582156100f7579182015b828111156100f65782518255916020019190600101906100db565b5b5090506101049190610108565b5090565b61012a91905b8082111561012657600081600090555060010161010e565b5090565b90565b6103938061013c6000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806341c0e1b514610051578063a413686214610063578063cfae3217146100bd575bfe5b341561005957fe5b610061610156565b005b341561006b57fe5b6100bb600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506101ea565b005b34156100c557fe5b6100cd610205565b604051808060200182810382528381815181526020019150805190602001908083836000831461011c575b80518252602083111561011c576020820191506020810190506020830392506100f8565b505050905090810190601f1680156101485780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156101e757600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b5b565b80600190805190602001906102009291906102ae565b505b50565b61020d61032e565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102a35780601f10610278576101008083540402835291602001916102a3565b820191906000526020600020905b81548152906001019060200180831161028657829003601f168201915b505050505090505b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106102ef57805160ff191683800117855561031d565b8280016001018555821561031d579182015b8281111561031c578251825591602001919060010190610301565b5b50905061032a9190610342565b5090565b602060405190810160405280600081525090565b61036491905b80821115610360576000816000905550600101610348565b5090565b905600a165627a7a72305820a89120798f8b367b08eefd82299ea98351bfcca35faaa1e4010fed675a54348e0029`
)

func TestGreet(t *testing.T) {
	t.Log("testing greet start")
	codeHash, account, evm, parsed, err := test_case.NewEngine(ABI, BIN, "testing")
	if err != nil { t.Errorf("new engine error:%v", err)}
	input, err := parsed.Pack("greet")
	if err != nil { t.Errorf("pack error:%v", err)}
	ret, err := evm.Call(*account, *codeHash, input)
	ret0 := new(string)
	err = parsed.Unpack(ret0, "greet", ret)
	if err != nil { t.Errorf("unpack error:%v", err)}
	t.Log("ret0:", *ret0)
	t.Log("testing greet end")

}

func TestSetGreet(t *testing.T) {
	t.Log("testing set greet start")

	codeHash, account, evm, parsed, err := test_case.NewEngine(ABI, BIN, "testing")
	if err != nil { t.Errorf("new engine error:%v", err)}

	input, err := parsed.Pack("setGreeting", "hello world!")
	if err != nil { t.Errorf("pack error:%v", err)}
	ret, err := evm.Call(*account, *codeHash, input)

	input, err = parsed.Pack("greet")
	if err != nil { t.Errorf("pack error:%v", err)}
	ret, err = evm.Call(*account, *codeHash, input)
	ret0 := new(string)
	err = parsed.Unpack(ret0, "greet", ret)
	if err != nil { t.Errorf("unpack error:%v", err)}
	t.Log("ret0:", *ret0)

	t.Log("testing set greet end")

}
