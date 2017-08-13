/**
*********************************************************
Copyright (c) 2017 Pegadroid IQ Solutions Private Limited
All rights reserved
*********************************************************
*/

package error

import (
	"fmt"
)

const (
	// InvalidArguments - Error code for invalid arguments to invoke function
	InvalidArguments = 1
	// UnmarshalError - Error code for invalid argument to unmarshal function
	UnmarshalError = 2
	//MarshalError - Error code while marshaling an object
	MarshalError = 3
	// PutStateError - Error code for while saving state of an asset
	PutStateError = 4
)

// TransactionError is the generic error object
// for any chaincode transactions. It implements the Go Error interface
type TransactionError struct {
	ErrorCode    int
	ErrorMessage string
}

func (e *TransactionError) Error() string {
	return fmt.Sprintf("%d : %s", e.ErrorCode, e.ErrorMessage)
}
