package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/casperlabs-go/client"
	"testing"
)

const (
	eventStoreApi = "https://event-store-api-clarity-delta.make.services"
	RpcUrl        = "https://node-clarity-delta.make.services/rpc"
)

var (
	casper = client.New(RpcUrl, eventStoreApi)
)

func Test_GetStatus(t *testing.T) {

	status, err := casper.GetStatus()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(status.ApiVersion)
}

func Test_GetLatestBlockInfo(t *testing.T) {

	height, err := casper.GetLatestBlockHeight()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(height)

}

func Test_GetBlockInfoByHeight(t *testing.T) {
	block, err := casper.GetBlockInfoByHeight(18584)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
}

func Test_GetBlockInfoByHash(t *testing.T) {
	block, err := casper.GetBlockInfoByHash("6f168ef1d9bfcca97146b4925924e9594dba03a3fe30952653867ecc5fda5746")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
}

func Test_GetDeployByDeployHash(t *testing.T) {
	txid := "20f1190d4ddc06246e07d5fd0454d90f3b509936e3d2584350239104e183a000"
	casper.GetDeployByHash(txid)
}

func Test_GetBlockTransfer(t *testing.T) {
	casper.GetBlockTransferByHeight(18584)
}

func Test_GetBalance(t *testing.T) {
	balance, err := casper.GetBalance("018fee8e44d228eb323e72ff977699dab804d8250868bf6533d3a48b82fe46631d")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(balance)
}
