package main

import "fmt"

func main()  {
	bc := NewBlockChain()
	bc.AddBlock("A to B send 1btc")
	bc.AddBlock("B to C send 3btc")
	for i, block := range bc.blocks  {
		fmt.Println("========block num",i)
		fmt.Printf("Data : %s \n",block.Data)
		fmt.Println("Version :",block.Version)
		fmt.Printf("hash        :%x \n",block.Hash)
		fmt.Printf("PreBlockHash:%x \n",block.PreBlockHash)
		fmt.Println("TimeStamp : ",block.TimeStamp)
		fmt.Println("TargetBits : ",block.TargetBits)
		fmt.Printf("Nonce : %d \n",block.Nonce)
		fmt.Printf("MerkelRoot : %x \n",block.MerkelRoot)
		pow:=NewProofOfWork(block)
		fmt.Printf("Isvalid : %v \n",pow.IsValid())
	}
}
