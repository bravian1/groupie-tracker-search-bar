package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations  string `json:"relations"`
	Dates	string `json:"dates"`
	Locations	string `json:"locations"`
}
type NewArtist struct {
	ID           int      
	Image        string   
	Name         string   
	Members      []string 
	CreationDate int      
	FirstAlbum   string   
	Relations  string 
	Dates	string 
	Locations	[]string 
}

type WrapedLocation struct {
	Index []Location `json:"index"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates string	`json:"dates"`
}

type WrapedDate struct {
	Index []Date `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Id             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
