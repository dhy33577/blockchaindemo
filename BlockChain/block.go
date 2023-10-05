package main

import (
	"time"
)

type Block struct {
	Version int64//区块版本号
	PreBlockHash []byte
	Hash []byte//为了方便而做了一些简化，正常比特币区块是不包含自己的hash值的
	TimeStamp int64
	TargetBits int64
	Nonce int64 //随机值
	MerkelRoot []byte

	Data []byte//交易本来是一个单独的数据结构,这里未来简便，便放在和区块头一起
}
/*
创建一个新的区块
*/
func NewBlock(data string,prevBlockHash []byte) *Block  {
	block:= &Block{
		Version:1,
		PreBlockHash:prevBlockHash,
		TimeStamp:time.Now().Unix(),
		TargetBits:targetBits,
		//Nonce:5,
		MerkelRoot:[]byte{},
		Data:[]byte(data)}
	//block.SetHash()
	pow:=NewProofOfWork(block)
	nonce,hash := pow.Run()
	block.Nonce=nonce
	block.Hash = hash
	return block
}
/**
设置区块的hash值
 */
 /*
func (block *Block)SetHash()  {
	//方法一

	//temp:=[][]byte{//二维的[][]byte，需要传进去一个[]byte
	//	IntToByte(block.Version),
	//	block.PreBlockHash,
	//	IntToByte(block.TimeStamp),
	//	block.MerkelRoot,
	//	IntToByte(block.Nonce),
	//	block.Data}
	////将区块的各个字段连接成以一个切片，使用[]byte{}进行连接，目的是避免污染原区块的信息
	//data:=bytes.Join(temp,[]byte{})
	////对区块进行sha256哈希算法，返回值为[32]byte 数组，不是切片
	//hash := sha256.Sum256(data)

	//方法二
	timestamp := []byte(strconv.FormatInt(block.TimeStamp, 10))
	headers := bytes.Join([][]byte{block.PreBlockHash, block.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	block.Hash= hash[:]
}
*/

/**
创建比特币创世区块，即它的第一个区块，它的前一个区块的hash值为空
 */
func NewGenesisBlock() *Block {
	return NewBlock("Create Genesis Block",[]byte{})
}

