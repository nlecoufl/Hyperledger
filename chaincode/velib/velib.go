/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"


	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Datas struct {
    LastUpdatedOther int `json:"lastUpdatedOther"`
    Ttl int `json:"ttl"`
    Data Stations `json:"data"`
}

type Stations struct {
    Stations []Data `json:"stations"`
}

type Data struct {
    StationCode   string `json:"stationCode"`
    Station_id  int `json:"station_id"`
    Num_bikes_available int `json:"num_bikes_available"`
    NumBikesAvailable int  `json:"numBikesAvailable"`
    Num_bikes_available_types  []Biketype `json:"num_bikes_available_types"`
    Num_docks_available  int `json:"num_docks_available"`
    NumDocksAvailable int `json:"numDocksAvailable"`
    Is_intalled int `json:"is_intalled"`
    Is_returning int `json:"is_returning"`
    Is_renting int `json:"is_renting"`
    Last_reported int `json:"last_reported"`
}

type Biketype struct {
    Mechanical int `json:"mechanical,omitempty"`
    Ebike int `json:"ebike,omitempty"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryData" {
		return s.queryData(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "queryAllDatas" {
		return s.queryAllDatas(APIstub)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	dataAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(dataAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	datas := []Data{ 
		Data{StationCode:"16107", Station_id:213688169, Num_bikes_available:2, NumBikesAvailable:2, Num_bikes_available_types:[]Biketype{Biketype{Mechanical:0, Ebike:0},Biketype{Mechanical:0,Ebike:2},}, Num_docks_available:33, NumDocksAvailable:33, Is_intalled:1, Is_returning:1, Is_renting:1, Last_reported:1579850732},
		Data{StationCode:"6015", Station_id:99950133, Num_bikes_available:9, NumBikesAvailable:9, Num_bikes_available_types:[]Biketype{Biketype{Mechanical:6, Ebike:0},Biketype{Mechanical:0,Ebike:3},}, Num_docks_available:46, NumDocksAvailable:46, Is_intalled:1, Is_returning:1, Is_renting:1, Last_reported:1579850255},
	}

	i := 0
	for i < len(datas) {
		fmt.Println("i is ", i)
		dataAsBytes, _ := json.Marshal(datas[i])
		APIstub.PutState("DATA"+strconv.Itoa(i), dataAsBytes)
		fmt.Println("Added", datas[i])
		i = i + 1
	}
	return shim.Success(nil)
}

/* func (s *SmartContract) createData(APIstub shim.ChaincodeStubInterface, args []int) sc.Response {

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var data = Data{StationCode: args[1], Station_id: args[2], Num_bikes_available: args[3], NumBikesAvailable: args[4], Num_bikes_available_types: args[5], Num_docks_available: args[6], NumDocksAvailable: args[7], Is_intalled: args[8], Is_returning:args[8], Is_renting: args[9], Last_reported: args[10] }
	dataAsBytes, _ := json.Marshal(data)
	APIstub.PutState(args[0], dataAsBytes)

	return shim.Success(nil)
}*/

func (s *SmartContract) queryAllDatas(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "DATA0"
	endKey := "DATA999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
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

	fmt.Printf("- queryAllDatas:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
