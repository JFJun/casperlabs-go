package client

import (
	"fmt"
	"github.com/JFJun/casperlabs-go/common"
	"github.com/JFJun/casperlabs-go/model"
)

type CasperClient struct {
	url           string
	casper        *common.RpcClient
	eventStoreApi string
}

/*
仅支持http,https
*/
func New(url, eventStoreApi string) *CasperClient {
	cc := new(CasperClient)
	cc.url = url
	cc.eventStoreApi = eventStoreApi

	cc.casper = common.Dial(cc.url, "", "")
	return cc
}

/*
这其实就是根据txid查询交易信息
deployHash就是txid
*/
func (cc *CasperClient) GetDeployByHash(deployHash string) {
	url := fmt.Sprintf("%s/deploy/%s", cc.eventStoreApi, deployHash)
	req := common.HttpGet(url)
	data, err := req.Bytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}

/*
根据区块hash获取区块的信息
*/
func (cc *CasperClient) GetBlockInfoByHash(blockHash string) (*model.ChainBlock, error) {
	var res model.ChainBlock
	params := make(map[string]interface{})
	params["block_identifier"] = map[string]interface{}{
		"Hash": blockHash,
	}
	err := cc.casper.SendRequest("chain_get_block", &res, params)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

/*
根据区块height获取区块的信息
*/
func (cc *CasperClient) GetBlockInfoByHeight(height int64) (*model.ChainBlock, error) {
	var res model.ChainBlock
	params := make(map[string]interface{})
	params["block_identifier"] = map[string]interface{}{
		"Height": height,
	}
	err := cc.casper.SendRequest("chain_get_block", &res, params)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (cc *CasperClient) GetLatestBlockHeight() (int64, error) {
	var res model.ChainBlock
	err := cc.casper.SendRequest("chain_get_block", &res, nil)
	if err != nil {
		return -1, err
	}
	return res.Block.Header.Height, err
}

func (cc *CasperClient) GetStatus() (*model.ChainStatus, error) {
	var status model.ChainStatus
	err := cc.casper.SendRequest("info_get_status", &status, nil)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

/*
params: Account hash or public key
*/
func (cc *CasperClient) GetBalance(ahopk string) {
	//todo
}

func (cc *CasperClient) Transfer() {
	//todo
}
