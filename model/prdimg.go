package model

import (
	"errors"
	"github.com/alexwang58/youhuo/components"
	"image"
	"os"
	"time"
)

const (
	PRDIMG_IMG_BASEPATH = "./assets/product_pic"
)

type Prdimg struct {
	Model
	Attribute *PrdimgAttribute
}

type PrdimgAttribute struct {
	// md5sum from image
	Imghash string
	// image width image.Config.Width
	Width int
	// image height image.Config.Height
	Height int
	// time updated
	Lst_tch int64
	// time created
	C_time int64
}

func NewPrdimg() *Prdimg {
	dbCon := components.NewMgo("localhost", "27017", "", "")
	dbCnnt := dbCon.C("youhuo", "prd_img")
	return &Prdimg{Attribute: &PrdimgAttribute{}, Model: Model{dbCon: dbCon, dbCnnt: dbCnnt}}
}

func (this *Prdimg) ParseImgAttribute(filepath string, md5sum string) (*PrdimgAttribute, error) {
	if reader, err := os.Open(filepath); err == nil {
		defer reader.Close()
		config, _, err := image.DecodeConfig(reader)
		if err != nil {
			return &PrdimgAttribute{Width: 0, Height: 0}, errors.New("Can not decodeConfig for: " + filepath + "\n" + err.Error())
		}

		return &PrdimgAttribute{Width: config.Width, Height: config.Height}, nil
	} else {
		return nil, errors.New("Can not open file: " + filepath + "\n" + err.Error())
	}

}

func (this *Prdimg) BeforeCreate() {
	this.Attribute.Lst_tch = time.Now().Unix()
	this.Attribute.C_time = time.Now().Unix()
}

func (this *Prdimg) Save() error {
	//components.NewAutoInc(1)
	if this.Attribute.Imghash == "" {
		return errors.New("Imghash Can not be empty")
	}
	this.BeforeCreate()
	mongo := components.NewMgo("localhost", "27017", "", "")
	err := mongo.C("youhuo", "prd_img").Insert(this.Attribute)
	if err != nil {
		return err
	}

	return nil
}
