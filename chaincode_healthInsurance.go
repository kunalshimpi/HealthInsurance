/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleHealthChaincode example simple Chaincode implementation
type SimpleHealthChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleHealthChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleHealthChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("**********Inside Init*******");
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	err:=stub.CreateTable("InsuranceAmount", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name:"Owner",Type: shim.ColumnDefinition_Bytes, Key: true},
		&shim.ColumnDefinition{Name:"Amount",Type:shim.ColumnDefinition_INT32, Key: false},
	})
	if err!= nil {
		return nil, errors.New("Error in Creating InsuranceAmount Table!")
	}

	adminCert, err := stub.GetCallerMetadata()

	if err!= nil{
		return nil, errors.New("Error Getting Metadata")
	}
	if len(adminCert) == 0 {
		return nil, errors.New("Admin Certificate is Empty!")
	}
	stub.PutState("admin", adminCert)

	fmt.Println("Admin is [%x] : ", adminCert)
	
	

	fmt.Println("Init Finished!")

	return nil, nil
}
func (t *SimpleHealthChaincode) assign(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("assign is running " + function)
	owner := arg[0];
err:= stub.InsertRow("InsuranceAmount", shim.Row{
		Columns: []*shim.Column {
			&shim.Column{Value: &shim.Column_String_{String_:owner}},
			&shim.Column{Value: &shim.Column_Int32{Int32:1000}},
		}
	})
}
// Invoke is our entry point to invoke a chaincode function
func (t *SimpleHealthChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleHealthChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "dummy_query" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return nil, nil;
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}
