package model

import (
	"github.com/alexwang58/youhuo/components"
	"gopkg.in/mgo.v2"
)

type Model struct {
	dbCon  *components.DBMongo
	dbCnnt *mgo.Collection
}

/*
func NewModel() *Model {
	return &Model{dbCon: nil, dbCnnt: nil}
}
*/
