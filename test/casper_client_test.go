package test

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/casperlabs-go/client"
	"testing"
)

var (
	casper = client.New("http://173.212.247.216:7777/rpc", "", "")
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
	block, err := casper.GetBlockInfoByHeight(18719)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
}

func Test_GetBlockInfoByHash(t *testing.T) {
	block, err := casper.GetBlockInfoByHash("8eb9a991af69f92be6f92f1ff3cf93345af550ff82b602805a1cb9fc0a8fbde5")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
}
