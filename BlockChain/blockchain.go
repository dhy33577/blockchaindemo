package main

import "os"

type BlockChain struct {
	blocks []*Block
}
/**
创建区块链
 */
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
/**
添加区块
 */
func (bc *BlockChain)AddBlock(data string)  {
	if len(bc.blocks)<=0{
		os.Exit(1)
	}
	lastBlock := bc.blocks[len(bc.blocks)-1]//获取链上的最后一个区块
	prevHash := lastBlock.Hash//当前最后一个区块的hash值就是需要创建的区块的prevHash
	block := NewBlock(data,prevHash)//调用方法创建区块
	bc.blocks = append(bc.blocks,block)//将创建的新区块添加到链上
}
