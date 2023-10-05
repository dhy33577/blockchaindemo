package main

import (
	"math/big"
	"bytes"
	"math"
	"crypto/sha256"
	"fmt"

)


/**
dinding定义pow的结构体
 */
type Proofofwork struct {
	block *Block
	targetBit *big.Int
}

const targetBits  =20//难度值 16前置4个0，20前置5个0，24前置6个0
func NewProofOfWork(block *Block) *Proofofwork  {
	var IntTarget = big.NewInt(1)//创建00000000000000000000000000001
	IntTarget.Lsh(IntTarget,uint(256-targetBits))//左移256然后右移targetBits 再转成16进制
	return &Proofofwork{block,IntTarget}
}
func (pow *Proofofwork)prepareRawData(nonce int64) []byte  {

	block:=pow.block
	temp:=[][]byte{//二维的[][]byte，需要传进去一个[]byte
		IntToByte(block.Version),
		block.PreBlockHash,
		IntToByte(block.TimeStamp),
		block.MerkelRoot,
		IntToByte(nonce),
		IntToByte(targetBits),
		block.Data}
	data:=bytes.Join(temp,[]byte{})
	return data
}

func (pow *Proofofwork)Run()(int64,[]byte)  {
	var nonce int64
	var hash [32]byte
	var HashInt big.Int

	for nonce<math.MaxInt64 {
		data:= pow.prepareRawData(nonce)
		hash= sha256.Sum256(data)
		HashInt.SetBytes(hash[:])

		if(HashInt.Cmp(pow.targetBit))==-1{
			fmt.Printf("Found hash :%x \n",hash)
			break;
		}else{
			nonce++
		}
	}
	return nonce,hash[:]
}

func (pow *Proofofwork)IsValid()bool  {
	data:= pow.prepareRawData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	var IntHash big.Int
	IntHash.SetBytes(hash[:])
	return IntHash.Cmp(pow.targetBit)==-1

}