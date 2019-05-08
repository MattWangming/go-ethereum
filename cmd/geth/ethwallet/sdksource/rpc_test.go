package sdksource

import "testing"

func TestGetAccount(t *testing.T) {
	node := "https://kovan.infura.io/v3/ef4fee2bd9954c6c8303854e0dce1ffe"
	addr := "0x1B37AB8d737B1776d3cC082D246Ee89Ed9693cD2"
	output := GetAccount(node,addr)
	t.Log(output)
}
