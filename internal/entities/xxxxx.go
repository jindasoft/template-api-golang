package entities

import "github.com/jindasoft/jinda-platform/xentities"

type Xxxxx struct {
	xentities.MongoBaseModel `json:",inline" bson:",inline"`

	Code        string                `json:"code" bson:"code"`
	Name        xentities.MongoLocale `json:"name" bson:"name"`
	Description xentities.MongoLocale `json:"description" bson:"description"`
	Ord         int16                 `json:"ord" bson:"ord"`
	Icon        string                `json:"icon" bson:"icon"`
}
