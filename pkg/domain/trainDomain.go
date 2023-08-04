package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Train struct {
	TrainNumber uint               `json:"trainNumber" bson:"trainNumber,omitempty"`
	TrainName   string             `json:"trainName" bson:"trainName,omitempty" validate:"required,min=2,max=50"`
	Route       primitive.ObjectID `json:"route,omitempty" bson:"route,omitempty"`
	TrainType   string             `json:"traintype" bson:"traintype,omitempty"`
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
