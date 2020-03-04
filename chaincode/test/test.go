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
	 "io/ioutil"
	 "os"
 
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

	// Open our jsonFile
	jsonFile, err := os.Open("../../test.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened test.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
  
	// we initialize our Datas array
	var datas Datas
  
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'datas' which we defined above
	json.Unmarshal(byteValue, &datas)
 
	 i := 0
	 for i < len(datas.Data.Stations) {
		 fmt.Println("i is ", i)
		 dataAsBytes, _ := json.Marshal(datas.Data.Stations[i])
		 APIstub.PutState("DATA"+strconv.Itoa(i),dataAsBytes)
		 fmt.Println("Added",datas.Data.Stations[i])
		 i = i + 1
	 }
 
	 return shim.Success(nil)
 }
 
 
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
 