package poamanager

import (
	"github.com/linkchain/common/util/log"
	poameta "github.com/linkchain/poa/meta"
	"github.com/linkchain/common/math"

	"github.com/linkchain/meta/block"
)

var mainchain  []poameta.POABlock
var bestHeight int


type POAChainManager struct {

}

func (m *POAChainManager) Init(i interface{}) bool{
	log.Info("POAChainManager init...");
	mainchain = make([]poameta.POABlock,1)
	bestHeight=0;
	return true
}

func (m *POAChainManager) Start() bool{
	log.Info("POAChainManager start...");
	return true
}

func (s *POAChainManager) Stop(){
	log.Info("POAChainManager stop...");
}

func (m *POAChainManager) GetBestBlock() block.IBlock  {
	return &mainchain[bestHeight];
}

func (m *POAChainManager) GetMainChain() block.IBlock  {
	return &mainchain[bestHeight];
}

func (m *POAChainManager) GetBlockByHash(h math.Hash) block.IBlock  {
	return &mainchain[bestHeight];
}

func (m *POAChainManager) GetBlockByHeight(height uint32) block.IBlock  {
	return &mainchain[height];
}

func (m *POAChainManager) GetBlockChainInfo() string  {
	return "this is poa chain";
}

func (m *POAChainManager) AddBlock(block block.IBlock)  {
	//mainchain = append(mainchain,block)
	bestHeight++
}

func (m *POAChainManager) UpdateChain() bool  {
	return true
}