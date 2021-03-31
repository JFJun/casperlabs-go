package client

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/JFJun/casperlabs-go/common"
	"github.com/JFJun/casperlabs-go/keys"
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

func (cc *CasperClient) GetLatestBlockInfo() (*model.ChainBlock, error) {
	var res model.ChainBlock
	err := cc.casper.SendRequest("chain_get_block", &res, nil)
	if err != nil {
		return nil, err
	}
	return &res, err
}
func (cc *CasperClient) GetLatestBlockHeight() (int64, error) {
	var res model.ChainBlock
	err := cc.casper.SendRequest("chain_get_block", &res, nil)
	if err != nil {
		return -1, err
	}
	return res.Block.Header.Height, err
}

func (cc *CasperClient) GetBlockTransferByHeight(height int64) (*model.BlockTransfer, error) {
	var res model.BlockTransfer
	params := make(map[string]interface{})
	params["block_identifier"] = map[string]interface{}{
		"Height": height,
	}
	err := cc.casper.SendRequest("chain_get_block_transfers", &res, params)
	if err != nil {
		return nil, fmt.Errorf("rpc chain_get_block_transfers error: %v", err)
	}
	return &res, nil
}

func (cc *CasperClient) GetStatus() (*model.ChainStatus, error) {
	var status model.ChainStatus
	err := cc.casper.SendRequest("info_get_status", &status, nil)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (cc *CasperClient) GetBalance(address string) (string, error) {
	return cc.GetBalanceWithHeight(address, -1)
}
func (cc *CasperClient) GetBalanceWithHeight(address string, height int64) (balance string, err error) {
	var lb *model.ChainBlock
	if height < 0 {
		lb, err = cc.GetLatestBlockInfo()
	} else {
		lb, err = cc.GetBlockInfoByHeight(height)
	}
	if err != nil || lb == nil {
		return "", fmt.Errorf("balance get state root hash error")
	}
	accountHash, err := keys.AddressToAccountHash(address)
	if err != nil {
		return "", err
	}
	accountHashStr := hex.EncodeToString(accountHash[:])
	stateRootHash := lb.Block.Header.StateRootHash
	key := fmt.Sprintf("account-hash-%s", accountHashStr)
	bs, err := cc.GetBlockState(stateRootHash, key, nil)
	if err != nil {
		return "", err
	}
	balanceUref := bs.StoredValue.Account.MainPurse
	if balanceUref == "" {
		return "", errors.New("balance uref is null")
	}
	var ab model.AccountBalance
	bp := map[string]interface{}{
		"state_root_hash": stateRootHash,
		"purse_uref":      balanceUref,
	}
	err = cc.casper.SendRequest("state_get_balance", &ab, bp)
	if err != nil {
		return "", fmt.Errorf("rpc state_get_balance error: %v", err)
	}

	return ab.BalanceValue, err
}

func (cc *CasperClient) GetBlockState(stateRootHash, key string, path []string) (*model.BlockState, error) {
	var res model.BlockState
	params := make(map[string]interface{})
	params["state_root_hash"] = stateRootHash
	params["key"] = key
	if path == nil {
		params["path"] = []interface{}{}
	} else {
		params["path"] = path
	}
	err := cc.casper.SendRequest("state_get_item", &res, params)
	if err != nil {
		return nil, fmt.Errorf("rpc state_get_item error: %v", err)
	}
	return &res, nil
}

func (cc *CasperClient) Transfer() {
	//todo
}
