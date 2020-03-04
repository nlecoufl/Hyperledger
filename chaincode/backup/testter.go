package main

import (
    "encoding/json"
    "fmt"

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

type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

func main() {
    var b = []byte(`{"lastUpdatedOther":1579850797,"ttl":3600,"data":{"stations":[{"stationCode":"16107","station_id":213688169,"num_bikes_available":2,"numBikesAvailable":2,"num_bikes_available_types":[{"mechanical":0},{"ebike":2}],"num_docks_available":33,"numDocksAvailable":33,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850732},{"stationCode":"6015","station_id":99950133,"num_bikes_available":9,"numBikesAvailable":9,"num_bikes_available_types":[{"mechanical":6},{"ebike":3}],"num_docks_available":46,"numDocksAvailable":46,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850255},{"stationCode":"9020","station_id":36255,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":21,"numDocksAvailable":21,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850484},{"stationCode":"12109","station_id":37815204,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":1},{"ebike":3}],"num_docks_available":26,"numDocksAvailable":26,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850753},{"stationCode":"5001","station_id":100769544,"num_bikes_available":8,"numBikesAvailable":8,"num_bikes_available_types":[{"mechanical":4},{"ebike":4}],"num_docks_available":30,"numDocksAvailable":30,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850786},{"stationCode":"14014","station_id":85002689,"num_bikes_available":11,"numBikesAvailable":11,"num_bikes_available_types":[{"mechanical":9},{"ebike":2}],"num_docks_available":47,"numDocksAvailable":47,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850734},{"stationCode":"17026","station_id":54000559,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":1},{"ebike":3}],"num_docks_available":36,"numDocksAvailable":36,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850765},{"stationCode":"17041","station_id":85043758,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":2},{"ebike":2}],"num_docks_available":35,"numDocksAvailable":35,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850729},{"stationCode":"10013","station_id":123095125,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":2},{"ebike":2}],"num_docks_available":56,"numDocksAvailable":56,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850678},{"stationCode":"9104","station_id":129026597,"num_bikes_available":7,"numBikesAvailable":7,"num_bikes_available_types":[{"mechanical":4},{"ebike":3}],"num_docks_available":15,"numDocksAvailable":15,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850421},{"stationCode":"14111","station_id":251039991,"num_bikes_available":2,"numBikesAvailable":2,"num_bikes_available_types":[{"mechanical":2},{"ebike":0}],"num_docks_available":23,"numDocksAvailable":23,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850746},{"stationCode":"6003","station_id":37874517,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":21,"numDocksAvailable":21,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850219},{"stationCode":"13007","station_id":66491398,"num_bikes_available":2,"numBikesAvailable":2,"num_bikes_available_types":[{"mechanical":0},{"ebike":2}],"num_docks_available":45,"numDocksAvailable":45,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850720},{"stationCode":"5110","station_id":210403080,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":23,"numDocksAvailable":23,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850689},{"stationCode":"6108","station_id":210561800,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":4},{"ebike":0}],"num_docks_available":13,"numDocksAvailable":13,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850370},{"stationCode":"33006","station_id":209063434,"num_bikes_available":9,"numBikesAvailable":9,"num_bikes_available_types":[{"mechanical":2},{"ebike":7}],"num_docks_available":23,"numDocksAvailable":23,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850316},{"stationCode":"42016","station_id":94555589,"num_bikes_available":3,"numBikesAvailable":3,"num_bikes_available_types":[{"mechanical":1},{"ebike":2}],"num_docks_available":24,"numDocksAvailable":24,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850779},{"stationCode":"41301","station_id":244498842,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":1},{"ebike":3}],"num_docks_available":46,"numDocksAvailable":46,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850627},{"stationCode":"10105","station_id":210738704,"num_bikes_available":0,"numBikesAvailable":0,"num_bikes_available_types":[{"mechanical":0},{"ebike":0}],"num_docks_available":25,"numDocksAvailable":25,"is_installed":1,"is_returning":0,"is_renting":0,"last_reported":1579850576},{"stationCode":"21010","station_id":43195240,"num_bikes_available":4,"numBikesAvailable":4,"num_bikes_available_types":[{"mechanical":2},{"ebike":2}],"num_docks_available":21,"numDocksAvailable":21,"is_installed":1,"is_returning":1,"is_renting":1,"last_reported":1579850515}]}}`)
    
    var j Datas
    json.Unmarshal(b, &j)
    printJson(j)

	i := 0
	for i < len(j.Data.Stations) {
		fmt.Println("i is ", i)
		jAsBytes, _ := json.Marshal(j.Data.Stations[i])
    fmt.Println(jAsBytes)
    fmt.Println("Added", j.Data.Stations[i])
		i = i + 1
	}
}

func printJson(j Datas) {
    fmt.Printf("LastUpdatedOther %v\n", j.LastUpdatedOther)
    fmt.Printf("Ttl %v\n", j.Ttl)
    fmt.Printf("Datas %v\n", j.Data)
}