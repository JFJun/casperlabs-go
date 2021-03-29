package client

import "github.com/JFJun/casperlabs-go/common"

type CasperClient struct {
	url, user, password string
	casper              *common.RpcClient
}

/*
仅支持http,https
*/
func New(url, user, password string) *CasperClient {
	cc := new(CasperClient)
	cc.url = url
	cc.user = user
	cc.password = password
	cc.casper = common.Dial(cc.url, cc.user, cc.password)
	return cc
}

/*
这其实就是根据txid查询交易信息
deployHash就是txid
*/
func (cc *CasperClient) GetDeployByHash(deployHash string) {
	//todo

}

/*
根据区块hash获取区块的信息
*/
func (cc *CasperClient) GetBlockInfo(blockHash string) {
	//todo
}

/*
根据区块height获取区块的信息
*/
func (cc *CasperClient) GetBlockInfoByHeight(height int64) {
	//todo
}

func (cc *CasperClient) GetLatestBlockInfo() {
	//todo
}

func (cc *CasperClient) GetStatus() {
	//todo
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
