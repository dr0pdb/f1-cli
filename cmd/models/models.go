package models

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// ErgastDuration represents time.Duration
type ErgastDuration struct {
	time.Duration
}

// ErgastDate represents time.Time
type ErgastDate struct {
	time.Time
}

// ErgastTime represents time.Time
type ErgastTime struct {
	time.Time
}

// Mrdata represents the Mrdata field of the response.
type Mrdata struct {
	XMLName xml.Name `xml:"MRData"`
	Races   []Race   `xml:"RaceTable>Race"`
}

// Race represents the properties of a f1 race.
type Race struct {
	NoResults         bool
	Circuit           Circuit
	Date              ErgastDate
	Time              ErgastTime
	RaceName          string
	Season            int                `xml:"season,attr"`
	Round             int                `xml:"round,attr"`
	Results           []Result           `xml:"ResultsList>Result"`
	QualifyingResults []QualifyingResult `xml:"QualifyingList>QualifyingResult"`
}

// QualifyingResult represents the results of qualifying round before the f1 race.
type QualifyingResult struct {
	Driver      Driver
	Constructor Constructor
	Q1          ErgastDuration
	Q2          ErgastDuration
	Q3          ErgastDuration
	Position    int `xml:"position,attr"`
}

// Circuit represents the properties of a f1 circuit.
type Circuit struct {
	CircuitName string
}

// Result represents the final results
type Result struct {
	Constructor Constructor
	Driver      Driver
	Laps        int
	Grid        int
	StatusID    int `xml:"statusId"`
	Status      string
	FastestLap  Lap
	Number      int `xml:"number,attr"`
	Position    int `xml:"position,attr"`
	Points      int `xml:"points,attr"`
}

// Lap represents the properties of a lap.
type Lap struct {
	Time              ErgastDuration
	Rank              int `xml:"rank,attr"`
	Lap               int `xml:"lap,attr"`
	AverageSpeed      float64
	AverageSpeedUnits string `xml:"units,attr"`
}

// Driver represents a f1 driver.
type Driver struct {
	DriverID        string `xml:"driverId,attr"`
	Code            string `xml:"code,attr"`
	PermanentNumber int
	GivenName       string
	FamilyName      string
	Nationality     string
	DateOfBirth     ErgastDate
}

// Constructor represents a f1 team.
type Constructor struct {
	ConstructorID string `xml:"constructorId,attr"`
	Name          string
	Nationality   string
}

// UnmarshalXML helps parse the ErgastDuration out of the XML response.
func (e *ErgastDuration) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string

	d.DecodeElement(&v, &start)

	p, err := parseErgastDuration(v)
	if err != nil {
		return err
	}

	*e = ErgastDuration{p}
	return nil
}

func parseErgastDuration(v string) (time.Duration, error) {
	parts := strings.Split(v, ":")

	// Minutes will be parts[0]

	m := parts[0]

	parts = strings.Split(parts[1], ".")

	// Seconds will be parts[0], Microseconds will be parts[1]

	s := parts[0]
	ms := parts[1]

	return time.ParseDuration(fmt.Sprintf("%vm%vs%vms", m, s, ms))
}

// UnmarshalXML decodes a ErgastDate from xml. source: https://stackoverflow.com/questions/17301149/golang-xml-unmarshal-and-time-time-fields
func (e *ErgastDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "2006-01-02"
	var v string

	d.DecodeElement(&v, &start)

	parse, err := time.Parse(format, v)
	if err != nil {
		return err
	}

	*e = ErgastDate{parse}
	return nil
}

// UnmarshalXML decodes ErgastTime from xml.
func (e *ErgastTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const format = "15:04:05Z"
	var v string

	d.DecodeElement(&v, &start)

	parse, err := time.Parse(format, v)
	if err != nil {
		return err
	}

	*e = ErgastTime{parse}
	return nil
}
