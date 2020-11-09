/*
 SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

//SmartContract is the data structure which represents this contract and on which  various contract lifecycle functions are attached
type SmartContract struct {
}

//Insurance info
type Insurance struct {
	ObjectType      string `json:"Type"`
	PatientsID      string `json:"patientsID"`
	InsuranceIDno   string `json:"insuranceIDNo"`
	PatientName     string `json:"patientName"`
	InsuranceStatus string `json:"insuranceStatus"`
	ClaimedBy       string `json:"claimedBy"`
	ItemFee         ItemFee
	Details         string `json:"details"`
	PrescriberSign  string `json:"prescriberSign"`
}

//ItemFee by Hospital
type ItemFee struct {
	TotalFee      string `json:"totalFee"`
	CoveredAmount string `json:"coveredAmount"`
}

//Init method
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("Init Firing!")
	return shim.Success(nil)
}

//Invoke functions
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Chaincode Invoke Is Running " + function)
	if function == "addInsuranceInfo" {
		return t.addInsuranceInfo(stub, args)
	}
	if function == "queryInsuranceInfo" {
		return t.queryInsuranceInfo(stub, args)
	}
	fmt.Println("Invoke did not find specified function " + function)
	return shim.Error("Invoke did not find specified function " + function)
}

func (t *SmartContract) addInsuranceInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 9 {
		return shim.Error("Incorrect Number of Aruments. Expecting 19")
	}

	fmt.Println("Adding new Insurance Info")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th Argument Must be a Non-Empty String")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th Argument Must be a Non-Empty String")
	}
	if len(args[7]) <= 0 {
		return shim.Error("8th Argument Must be a Non-Empty String")
	}
	if len(args[8]) <= 0 {
		return shim.Error("9th Argument Must be a Non-Empty String")
	}

	patientsID := args[0]
	insuranceIDNo := args[1]
	patientName := args[2]
	insuranceStatus := args[3]
	claimedBy := args[4]
	totalFee := args[5]
	coveredAmount := args[6]
	details := args[7]
	prescriberSign := args[8]

	// ======Check if PurchaseRequest Already exists

	InsuranceAsBytes, err := stub.GetState(patientsID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if InsuranceAsBytes != nil {
		return shim.Error("The Inserted Patient ID already Exists: " + patientsID)
	}

	// ===== Create Item Object and Marshal to JSON

	objectType := "Insurance"
	Insurance := &Insurance{objectType, patientsID, insuranceIDNo, patientName, insuranceStatus, claimedBy, ItemFee{totalFee, coveredAmount}, details, prescriberSign}
	InsuranceJSONasBytes, err := json.Marshal(Insurance)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save PurchaseRequest to State

	err = stub.PutState(patientsID, InsuranceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved Insurance Info")
	return shim.Success(nil)
}

func (t *SmartContract) queryInsuranceInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientsID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"Insurance\",\"patientsID\":\"%s\"}}", patientsID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//Main Function starts up the Chaincode
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Smart Contract could not be run. Error Occured: %s", err)
	} else {
		fmt.Println("Smart Contract successfully Initiated")
	}
}
