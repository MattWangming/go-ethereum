package sdksource

import (
	"os/user"
	"testing"
)

func TestGetAccount(t *testing.T) {
	node := "https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	addr := "0x1B37AB8d737B1776d3cC082D246Ee89Ed9693cD2"
	output := GetAccount(node,addr)
	t.Log(output)
}


func TestTransferETH(t *testing.T) {
	usr, _ := user.Current()
	rootDir := usr.HomeDir
	node := "https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	name := "easyzone"
	password := "wm131421"
	toAddr := "0x1B37AB8d737B1776d3cC082D246Ee89Ed9693cD2"
	amount := int64(200000000000000000)
	gasLimit := uint64(21000)
	output := TransferETH(rootDir,node,name,password,toAddr,amount,gasLimit)
	t.Log(output)
}