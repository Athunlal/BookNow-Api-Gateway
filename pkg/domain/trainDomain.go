package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Train struct {
	TrainId        string             `json:"trainid,omitempty" bson:"trainid,omitempty"`
	TrainNumber    uint               `json:"trainNumber" bson:"trainNumber,omitempty"`
	TrainName      string             `json:"trainName" bson:"trainName,omitempty" validate:"required,min=2,max=50"`
	Route          primitive.ObjectID `json:"route,omitempty" bson:"route,omitempty"`
	StartingTime   string             `json:"startingtime,omitempty" bson:"startingtime,omitempty"`
	EndingtingTime string             `json:"endingtingtime,omitempty" bson:"endingtingtime,omitempty"`
	TrainType      string             `json:"traintype" bson:"traintype,omitempty"`
	Seatsid        primitive.ObjectID `json:"seatid,omitempty" bson:"_id,omitempty"`
	Compartment    []Compartment      `json:"compartment,omitempty" bson:"compartment,omitempty"`
}
type Compartment struct {
	Seatid primitive.ObjectID `json:"seatid,omitempty" bson:"_id,omitempty"`
}

type Station struct {
	StationName string `json:"stationname" bson:"stationname,omitempty"`
	City        string `json:"city" bson:"city,omitempty"`
	StationType string `json:"stationtype" bson:"stationtype,omitempty"`
}
type Route struct {
	RouteId   primitive.ObjectID `json:"routeid" bson:"_id,omitempty"`
	RouteName string             `json:"routename" bson:"routename,omitempty"`
	RouteMap  []RouteStation     `json:"routemap,omitempty" bson:"routemap,omitempty"`
}

type RouteStation struct {
	StationId   primitive.ObjectID `json:"stationid" bson:"stationid,omitempty"`
	StationName string             `json:"stationname" bson:"stationname,omitempty"`
	Distance    float32            `json:"distance,omitempty" bson:"distance,omitempty"`
	Time        time.Time          `json:"time" bson:"time,omitempty"`
}

type SearchingTrainRequstedData struct {
	Date                 string             `json:"data" bson:"data,omitempty"`
	SourceStationid      primitive.ObjectID `json:"sourcestationid,omitempty" bson:"sourcestationid,omitempty"`
	DestinationStationid primitive.ObjectID `json:"destinationstationid,omitempty" bson:"destinationstationid,omitempty"`
}

type SearchingTrainResponseData struct {
	SearcheResponse []Train `json:"searcheresponse,omitempty" bson:"searcheresponse,omitempty"`
}

type SeatDetails struct {
	SeatNumber     string `json:"seatnumber,omitempty" bson:"seatnumber,omitempty"`
	SeatType       string `json:"seattype,omitempty" bson:"seattype,omitempty"`
	HasPowerOutlet bool   `json:"haspoweroutlet,omitempty" bson:"haspoweroutlet,omitempty"`
}

type Seats struct {
	Price        int           `json:"price,omitempty" bson:"price,omitempty"`
	Availability bool          `json:"availability,omitempty" bson:"availability,omitempty"`
	Compartment  string        `json:"compartment,omitempty" bson:"compartment,omitempty"`
	SeatDetails  []SeatDetails `json:"seatDetails,omitempty" bson:"seatDetails,omitempty"`
}

type SeatAddingData struct {
	Price         float32 `json:"price"`
	NumbserOfSeat int     `json:"number_of_seat"`
	Compartment   string  `json:"compartment"`
	TypeOfSeat    string  `json:"type_of_seate"`
}
type BookingData struct {
	TrainId       string `json:"trainid" binding:"required"`
	CompartmentId string `json:"compartmentid" binding:"required"`
}

type Travelers struct {
	Travelername string `json:"travelername"`
}

type Payment struct {
	Trainname   string      `json:"trainname,omitempty" bson:"price,omitempty"`
	TrainNumber int64       `json:"trainnumber"`
	UserName    string      `json:"username"`
	Travelers   []Travelers `json:"travelers"`
}

type Wallet struct {
	Username string
	Amount   float32
}
