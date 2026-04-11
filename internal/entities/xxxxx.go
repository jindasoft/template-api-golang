package entities

import "github.com/jindasoft/jinda-platforms/xentities"

type Xxxxx struct {
	Code string                `json:"code" bson:"code"`
	Name xentities.MongoLocale `json:"name" bson:"name"`

	xentities.MongoBase `json:",inline" bson:",inline"`
}
