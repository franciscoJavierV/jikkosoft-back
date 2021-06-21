package model

import (
	"gopkg.in/mgo.v2/bson"
)

//user model
type userShema struct {
	Id              bson.ObjectId `bson:":_id" json:"id ,omitempty"`
	UserName        string        `bson:"UserName" json:"UserName ,omitempty"`
	UserDescription string        `bson:"UserDescription" json:"UserDescription ,omitempty"`
	Bussines        bson.ObjectId `bson:"Bussines" json:"Bussines ,omitempty"`
}

type Bussines struct {
	Id              bson.ObjectId `bson:":_id" json:"id ,omitempty"`
	BussinesName    string        `bson:"BussinesName" json:"BussinesName ,omitempty"`
	BussinessAdress string        `bson:"BussinessAdress" json:"BussinessAdress ,omitempty"`
}

//
