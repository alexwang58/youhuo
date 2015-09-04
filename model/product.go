package model

import (
	"github.com/alexwang58/youhuo/components"
	//"gopkg.in/mgo.v2/bson"
	"time"
)

type Product struct {
	Model
	Attribute *ProductAttribute
	// some one this product belongs to
	owner *Person
}

/*
type (prd *Product)Attribute struct {
	_id int64
	// id generate by database automaticly
	uid int64
	// the name of a product
	ptitle string
	// pric this product will be sold
	price int
	// Status of the product
	status int
	// time updated
	lst_tch int
	// time created
	c_time int
}
*/

type ProductAttribute struct {
	// id generate by database automaticly
	Pid int
	// product owner
	Uid int
	// the name of a product
	Ptitle string
	// pric this product will be sold
	Price int
	// images for this product
	Pimg []string
	// Status of the product
	Status int
	// time updated
	Lst_tch int64
	// time created
	C_time int64
}

func NewProduct() *Product {
	dbCon := components.NewMgo("localhost", "27017", "", "")
	dbCnnt := dbCon.C("youhuo", "prd_item")
	return &Product{Attribute: &ProductAttribute{}, Model: Model{dbCon: dbCon, dbCnnt: dbCnnt}}
}

func (prd *Product) Save() {
	//components.NewAutoInc(1)
	prd.Attribute.Pid = components.PRD_AUTOINC.Id()
	prd.BeforeCreate()
	/*
		prd.Attribute = &ProductAttribute{
			components.PRD_AUTOINC.Id(),
			"gongy了4.0",
			4,
			2,
			1231231897,
			234243342,
		}
	*/
	mongo := components.NewMgo("localhost", "27017", "", "")
	err := mongo.C("youhuo", "prd_item").Insert(prd.Attribute)
	if err != nil {
		panic(err)
	}
}

func (prd *Product) BeforeCreate() {
	prd.Attribute.Lst_tch = time.Now().Unix()
	prd.Attribute.C_time = time.Now().Unix()
}

func (prd *Product) UnsetAttribute() {
	prd.Attribute = &ProductAttribute{}
}

func (prd *Product) GetByObjid(objid string) (*ProductAttribute, error) {
	result := new(ProductAttribute)
	err := prd.dbCon.FindByObid(objid, prd.dbCnnt, &result)

	return result, err
}
