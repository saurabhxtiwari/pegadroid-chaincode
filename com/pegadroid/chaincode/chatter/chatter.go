/**
*********************************************************
Copyright (c) 2017 Pegadroid IQ Solutions Private Limited
All rights reserved
*********************************************************
*/

package chatter

import (
	"encoding/json"
	"fmt"
	"math/rand"
	assets "pegadroid-chaincode/com/pegadroid/chaincode/chatter/assets"
	pgError "pegadroid-chaincode/com/pegadroid/chaincode/error"
	"time"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	peer "github.com/hyperledger/fabric/protos/peer"
)

// Chatter struct is our chaincode interface implementation
type Chatter struct {
}

// Init chaincode interface method
func (chatter *Chatter) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Init Success"))
}

//Invoke chaincode interface method
func (chatter *Chatter) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	sysError := pgError.TransactionError{ErrorCode: -1, ErrorMessage: ""}
	var args []string
	var funcName string
	funcName, args = stub.GetFunctionAndParameters()

	if len(args) != 1 {
		sysError = pgError.TransactionError{ErrorCode: pgError.InvalidArguments, ErrorMessage: "Invalid argument"}
		return shim.Error(sysError.Error())
	}

	switch funcName {
	case "createPerson":
		person := assets.Person{}
		if unmarshalError := json.Unmarshal([]byte(args[0]), &person); unmarshalError != nil {
			sysError = pgError.TransactionError{ErrorCode: pgError.UnmarshalError, ErrorMessage: "Invalid argument " + unmarshalError.Error()}
			return shim.Error(sysError.Error())
		}

		createPersonResponse, createPersonError := createPerson(stub, &person)
		if createPersonError != nil {
			return shim.Error(createPersonError.Error())
		}

		return shim.Success([]byte("Invoke Success : " + createPersonResponse))
	case "queryPerson":
		return shim.Success([]byte("Invoke Success"))
	default:
		return shim.Success([]byte("Invoke Success"))
	}
}

func createPerson(stub shim.ChaincodeStubInterface, person *assets.Person) (string, error) {
	sysError := pgError.TransactionError{ErrorCode: -1, ErrorMessage: ""}
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	person.SetID(randomGenerator.Int())

	data, marshalError := json.Marshal(person)
	if marshalError != nil {
		sysError = pgError.TransactionError{ErrorCode: pgError.MarshalError, ErrorMessage: "Marshal Error " + marshalError.Error()}
		return "", &sysError
	}

	createPersonError := stub.PutState(person.GetEmail(), data)
	if createPersonError != nil {
		sysError = pgError.TransactionError{ErrorCode: pgError.PutStateError, ErrorMessage: "Create person failed - " + createPersonError.Error()}
		return "", &sysError
	}

	response := fmt.Sprintf("Tx ID: %s Data: %s", stub.GetTxID(), string(data))
	return response, nil
}

// Start function to start the chaincode container
func (chatter *Chatter) Start() {
	if error := shim.Start(chatter); error != nil {
		e := pgError.TransactionError{ErrorCode: 1, ErrorMessage: "Unprecedented error."}
		fmt.Printf("Error starting chaincode container %s", e.Error())
	}
}
