package main

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type temperature struct {
	ID         string  `json:"id"`
	Celsius    float32 `json:"celsius"`
	Fahrenheit float32 `json:"fahrenheit"`
	Kelvin     float32 `json:"kelvin"`
}

type Temperature struct {
	Celsius    float32 `xml:"celsius"`
	Fahrenheit float32 `xml:"fahrenheit"`
	Kelvin     float32 `xml:"kelvin"`
}
