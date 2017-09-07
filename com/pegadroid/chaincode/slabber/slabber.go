/**
*********************************************************
Copyright (c) 2017 Pegadroid IQ Solutions Private Limited
All rights reserved
*********************************************************
*/

package slabber

import (
	"encoding/json"
	"fmt"
	"math/rand"
	pgError "pegadroid-sample-chaincode/com/pegadroid/chaincode/error"
	assets "pegadroid-sample-chaincode/com/pegadroid/chaincode/slabber/assets"
	"time"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	peer "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("SLABBER-CHAINCODE-LOGGER")

// Slabber struct is our chaincode interface implementation
type Slabber struct {
}

// Init chaincode interface method
func (slabber *Slabber) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Init Success"))
}

//Invoke chaincode interface method
func (slabber *Slabber) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	sysError := pgError.TransactionError{ErrorCode: -1, ErrorMessage: ""}
	var args []string
	var funcName string
	funcName, args = stub.GetFunctionAndParameters()

	argsBytesMatrix := stub.GetArgs()
	for i := range argsBytesMatrix {
		logger.Info(argsBytesMatrix[i])
		//fmt.Println(argsBytesMatrix[i])
	}

	argsSlice, _ := stub.GetArgsSlice()
	logger.Info(argsSlice)
	//fmt.Println(argsSlice)

	/*if len(args) != 1 {
		sysError = pgError.TransactionError{ErrorCode: pgError.InvalidArguments, ErrorMessage: "Invalid argument"}
		return shim.Error(sysError.Error())
	}*/

	for _, value := range args {
		logger.Info(value)
	}

	switch funcName {
	case "createPerson":
		logger.Info("Inside create person")
		person := assets.Person{}
		if unmarshalError := json.Unmarshal([]byte(args[0]), &person); unmarshalError != nil {
			sysError = pgError.TransactionError{ErrorCode: pgError.UnmarshalError, ErrorMessage: "Invalid argument " + unmarshalError.Error()}
			return shim.Error(sysError.Error())
		}

		createPersonResponse, createPersonError := createPerson(stub, &person)
		if createPersonError != nil {
			return shim.Error(createPersonError.Error())
		}

		logger.Info("Invoke Success : " + createPersonResponse)
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
func (slabber *Slabber) Start() {
	logger.SetLevel(shim.LogInfo)
	shim.SetLoggingLevel(shim.LogInfo)
	if error := shim.Start(slabber); error != nil {
		e := pgError.TransactionError{ErrorCode: 1, ErrorMessage: "Unprecedented error."}
		fmt.Printf("Error starting chaincode container %s", e.Error())
	}
}
