package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

func ConvertFlightFilterData(data map[string][]string) bson.M {
	converted := bson.M{}
	for key, el := range data {
		if key == "dateSource" || key == "dateDestination" {
			n, _ := strconv.Atoi(el[0])
			t := time.Unix(int64(n)/1000, 0)
			converted[key] = primitive.NewDateTimeFromTime(t)
		} else if key == "boughtTickets" || key == "totalTickets" {
			converted[key], _ = strconv.Atoi(el[0])
		} else {
			converted[key] = el[0]
		}
	}
	return converted
}
