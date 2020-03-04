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
    LastUpdatedOther string `json:"lastUpdatedOther"`
    Ttl string `json:"ttl"`
    Data Stations `json:"data"`
}

type Stations struct {
    Stations []Data `json:"stations"`
}

type Data struct {
    StationCode   string `json:"stationCode"`
    Station_id  string `json:"station_id"`
    Num_bikes_available string `json:"num_bikes_available"`
    NumBikesAvailable string  `json:"numBikesAvailable"`
    Num_bikes_available_types  []Biketype `json:"num_bikes_available_types"`
    Num_docks_available  string `json:"num_docks_available"`
    NumDocksAvailable string `json:"numDocksAvailable"`
    Is_intalled string `json:"is_intalled"`
    Is_returning string `json:"is_returning"`
    Is_renting string `json:"is_renting"`
    Last_reported string `json:"last_reported"`
}

type Biketype struct {
    Mechanical string `json:"mechanical,omitempty"`
    Ebike string `json:"ebike,omitempty"`
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
	 } else if function == "createDatas" {
		 return s.createDatas(APIstub, args)
	 }  
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 DataAsBytes, _ := APIstub.GetState(args[0])
	 return shim.Success(DataAsBytes)
 }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	var b = []byte(`{"lastUpdatedOther":1579850797,"ttl":3600,"data":{"stations":[{"stationCode":"16107","station_id":213688169,"num_bikes_available":2,"numBikesAvailable":2,"num_bikes_available_types":[{"mechanical":0},{"ebike":2}],"num_docks_available":33,"numDocksAvailable":33,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850732},{"stationCode":"6015","station_id":99950133,"num_bikes_available":9,"numBikesAvailable":9,"num_bikes_available_types":[{"mechanical":6},{"ebike":3}],"num_docks_available":46,"numDocksAvailable":46,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850255},{"stationCode":"9020","station_id":36255,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":21,"numDocksAvailable":21,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850484},{"stationCode":"12109","station_id":37815204,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":1},{"ebike":3}],"num_docks_available":26,"numDocksAvailable":26,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850753},{"stationCode":"5001","station_id":100769544,"num_bikes_available":8,"numBikesAvailable":8,"num_bikes_available_types":[{"mechanical":4},{"ebike":4}],"num_docks_available":30,"numDocksAvailable":30,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850786},{"stationCode":"14014","station_id":85002689,"num_bikes_available":11,"numBikesAvailable":11,"num_bikes_available_types":[{"mechanical":9},{"ebike":2}],"num_docks_available":47,"numDocksAvailable":47,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850734},{"stationCode":"17026","station_id":54000559,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":1},{"ebike":3}],"num_docks_available":36,"numDocksAvailable":36,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850765},{"stationCode":"17041","station_id":85043758,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":2},{"ebike":2}],"num_docks_available":35,"numDocksAvailable":35,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850729},{"stationCode":"10013","station_id":123095125,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":2},{"ebike":2}],"num_docks_available":56,"numDocksAvailable":56,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850678},{"stationCode":"9104","station_id":129026597,"num_bikes_available":7,"numBikesAvailable":7,"num_bikes_available_types":[{"mechanical":4},{"ebike":3}],"num_docks_available":15,"numDocksAvailable":15,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850421},{"stationCode":"14111","station_id":251039991,"num_bikes_available":2,"numBikesAvailable":2,"num_bikes_available_types":[{"mechanical":2},{"ebike":0}],"num_docks_available":23,"numDocksAvailable":23,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850746},{"stationCode":"6003","station_id":37874517,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":21,"numDocksAvailable":21,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850219},{"stationCode":"13007","station_id":66491398,"num_bikes_available":2,"numBikesAvailable":2,"num_bikes_available_types":[{"mechanical":0},{"ebike":2}],"num_docks_available":45,"numDocksAvailable":45,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850720},{"stationCode":"5110","station_id":210403080,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":23,"numDocksAvailable":23,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850689},{"stationCode":"6108","station_id":210561800,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":4},{"ebike":0}],"num_docks_available":13,"numDocksAvailable":13,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850370},{"stationCode":"33006","station_id":209063434,"num_bikes_available":9,"numBikesAvailable":9,"num_bikes_available_types":[{"mechanical":2},{"ebike":7}],"num_docks_available":23,"numDocksAvailable":23,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850316},{"stationCode":"42016","station_id":94555589,"num_bikes_available":3,"numBikesAvailable":3,"num_bikes_available_types":[{"mechanical":1},{"ebike":2}],"num_docks_available":24,"numDocksAvailable":24,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850779},{"stationCode":"41301","station_id":244498842,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":1},{"ebike":3}],"num_docks_available":46,"numDocksAvailable":46,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850627},{"stationCode":"10105","station_id":210738704,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":25,"numDocksAvailable":25,"is_installed":1,"is_returning":0,"is_renting":0,"last_reported":1579850576},{"stationCode":"21010","station_id":43195240,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":2},{"ebike":2}],"num_docks_available":21,"numDocksAvailable":21,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850515}]}}`)
    
	var j Datas
    json.Unmarshal(b,&j)

	i := 0
	for i < len(j.Data.Stations) {
		fmt.Println("i is ", i)
		jAsBytes, _ := json.Marshal(j.Data.Stations[i])
		APIstub.PutState("DATA"+strconv.Itoa(i), jAsBytes)
		fmt.Println("Added", j.Data.Stations[i])
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

 
func (s *SmartContract) createDatas(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 13 {
		return shim.Error("Incorrect number of arguments. Expecting 12")
	}

	var biketype1 = Biketype{Mechanical: args[5], Ebike: "0"}
	var biketype2 = Biketype{Mechanical: "0", Ebike: args[6]}
	var data = Data{StationCode: args[1], Station_id: args[2], Num_bikes_available: args[3], NumBikesAvailable: args[4], Num_bikes_available_types: []Biketype{biketype1,biketype2}, Num_docks_available: args[7], NumDocksAvailable: args[8], Is_intalled: args[9], Is_returning: args[10], Is_renting: args[11], Last_reported: args[12]}

	dataAsBytes, _ := json.Marshal(data)
	APIstub.PutState(args[0], dataAsBytes)

	return shim.Success(nil)
}
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
 