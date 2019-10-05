package models

// Location data type representing the info of a location.
type Location struct {
	Lat      string
	Long     string
	Locality string
	Country  string
}

// Circuit data type representing the info of a f1 race circuit.
type Circuit struct {
	CircuitID   string
	URL         string
	CircuitName string
	Location    Location
}

// Race data type representing the info of a f1 race.
type Race struct {
	Season   string
	Round    int
	URL      string
	RaceName string
	Circuit  Circuit
	Date     string
	Time     string
}
