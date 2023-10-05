package main

import (
	"os"
	"bytes"
	"encoding/binary"
	"fmt"
)

//func main() {
//	fmt.Println("start do IntToByte...")
//	bt:=IntToByte(10)
//	fmt.Println(string(bt))
//}

/*
将int转化为byte的工具方法
*/
func IntToByte(num int64) []byte  {
	var buffer bytes.Buffer
	err:=binary.Write(&buffer,binary.BigEndian,num)
	CheckErr("IntToByte",err)
	return buffer.Bytes()
}
/**
打印错误信息
 */
func CheckErr(info string ,err error)  {
	if err!=nil{
		fmt.Println("err info :",info,err)
		os.Exit(1)
	}

}
