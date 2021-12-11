package model

type Privilege struct {
	//ID   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name string `json:"name"`
}
