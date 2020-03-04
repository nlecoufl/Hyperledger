package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)

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
 

func main() {
  // Open our jsonFile
  jsonFile, err := os.Open("test.json")
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
  fmt.Println(datas)
  // we iterate through every user within our users array and
  // print out the user Type, their name, and their facebook url
  // as just an example
  for i := 0; i < len(datas.Data.Stations); i++ {
//      fmt.Println("station_id: " + strconv.Itoa(datas.data.Stations[i].station_id))
//      fmt.Println("stationCode: " + datas.data.Stations[i].stationCode)
        fmt.Println("StationCode: " + datas.Data.Stations[i].StationCode)
        fmt.Println("Station_id: " + strconv.Itoa(datas.Data.Stations[i].Station_id))
        for j := 0; j < 2; j++ {
          if j==0 {
            fmt.Println("Mechanical: " + strconv.Itoa(datas.Data.Stations[i].Num_bikes_available_types[j].Mechanical))
          }
          if j==1 {
            fmt.Println("Ebike: " + strconv.Itoa(datas.Data.Stations[i].Num_bikes_available_types[j].Ebike))
          }

        }
  }
}