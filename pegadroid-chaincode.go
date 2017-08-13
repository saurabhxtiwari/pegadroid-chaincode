/**
*********************************************************
Copyright (c) 2017 Pegadroid IQ Solutions Private Limited
All rights reserved
*********************************************************
*/

package main

import (
	"fmt"
	"pegadroid-chaincode/com/pegadroid/chaincode/chatter"
)

func main() {
	fmt.Println("Starting Chaincode")
	ch := chatter.Chatter{}
	ch.Start()
}
