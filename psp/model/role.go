package model

type Role struct {
	//ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name       string      `json:"name"`
	Privileges []Privilege `json:"privileges"`
}
