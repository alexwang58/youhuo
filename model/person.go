package model

import (
	//"fmt"
	"github.com/alexwang58/youhuo/components"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	Model
	Attribute *PersonAttribute
}

type PersonAttribute struct {
	// id generate by database automaticly
	Uid int
	// the name of the user
	Nick string
	// the phone number which can contact with whom
	Phone string
	// password
	Passwd string
	// time updated
	Lst_tch int64
	// time created
	C_time int64
}

func NewPerson() *Person {
	dbCon := components.NewMgo("localhost", "27017", "", "")
	dbCnnt := dbCon.C("youhuo", "profile")
	return &Person{Attribute: &PersonAttribute{}, Model: Model{dbCon: dbCon, dbCnnt: dbCnnt}}
}

func (psn *Person) AccessExists(name string, pwd string) (*PersonAttribute, error) {
	result := new(PersonAttribute)
	query := bson.M{"nick": name, "passwd": components.EncryptPassword(pwd)}
	err := psn.dbCon.FindOne(psn.dbCnnt, query, &result)
	return result, err
}

func (psn *Person) GetByUid(uid int) (*PersonAttribute, error) {
	result := new(PersonAttribute)
	query := bson.M{"uid": uid}
	err := psn.dbCon.FindOne(psn.dbCnnt, query, &result)
	return result, err
}

func (psn *Person) Save() {
	//components.NewAutoInc(1)
	/*
		psn.attribute = &PersonAttribute{
			components.PSN_AUTOINC.Id(),
			name,
			phone,
			pwd,
			1231231897,
			234243342,
		}
	*/
	mongo := components.NewMgo("localhost", "27017", "", "")
	err := mongo.C("youhuo", "profile").Insert(psn.Attribute)
	if err != nil {
		panic(err)
	}
}

func (psn *Person) beforeCreate() {
	psn.Attribute.Lst_tch = time.Now().Unix()
	psn.Attribute.C_time = time.Now().Unix()
}
