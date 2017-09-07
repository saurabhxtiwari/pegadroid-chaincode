/**
*********************************************************
Copyright (c) 2017 Pegadroid IQ Solutions Private Limited
All rights reserved
*********************************************************
*/

package main

import (
	"fmt"
	"pegadroid-sample-chaincode/com/pegadroid/chaincode/slabber"
)

func main() {
	fmt.Println("Starting Chaincode")
	ch := slabber.Slabber{}
	ch.Start()
}
