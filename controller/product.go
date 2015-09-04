package controller

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/alexwang58/youhuo/components"
	"github.com/alexwang58/youhuo/model"
	"mime/multipart"
	"net/http"
	"path"
	"strconv"
	"strings"
	//	"reflect"
)

type ProductController struct {
	Controller
}

func ProductCtrl(w http.ResponseWriter, r *http.Request) error {
	appc := new(ProductController)
	if err := appc.Active(w, r, appc); err != nil {
		return nil
	}
	return nil

}

// MUST implement this func
// It`s a action map for this controller
func (appC *ProductController) actionList() map[string]func() {
	return map[string]func(){
		"index": appC.Index,
		"list":  appC.List,
		"show":  appC.Show,
		"add":   appC.Add,
	}
}

func (appC *ProductController) dispatchError() (string, string, error) {
	return "site", "index", nil
}

// Action for product/list/xxxx
func (appC *ProductController) Index() {
	appC.render("product/index", rPrm{
		"title": "Procuct/Index",
		//"person": appC.mController.param["qp"],
	})
}

// Action for product/list/xxxx
func (appC *ProductController) List() {
	prdt := model.NewProduct()
	result, err := prdt.GetByObjid("5583a5a91db504e763888030")
	if err != nil {
		fmt.Println("Fuckk")
	}
	fmt.Println(result)
	appC.render("product/list", rPrm{
		"title":  "Procuct/List",
		"person": appC.mContext.GetParam("qp"),
	})
}

// Create a product
func (appC *ProductController) Show() {
	appC.render("product/show", rPrm{
		"title":  "Procuct/Show",
		"person": appC.mContext.GetParam("qp"),
	})
}

// Create a new product
func (appC *ProductController) Add() {
	prdt := model.NewProduct()
	//model.Save()
	ptitle := appC.mContext.GetPost("ptitle")
	price := appC.mContext.GetPost("price")
	picname, picpath, picerr := components.Upload(appC.mContext.GetResquest(), "picture", appC)
	if ptitle != "" && price != "" {
		uid, ok := appC.yhSc.Get(components.KeySessionUid).(int)
		if !ok {
			fmt.Println("Uid is NOT type of string")
			appC.Redirect("person", "signin")
			return
		}
		prdt.Attribute.Uid = uid
		prdt.Attribute.Ptitle = ptitle
		_price, err := strconv.ParseInt(price, 10, 32)
		if err == nil {
			prdt.Attribute.Price = int(_price)
			prdt.Save()
		} else {
			fmt.Println(err)
		}
		if picerr != nil {
			fmt.Println(picerr)
		} else {
			fmt.Println("Up to :" + picpath + "/" + picname)
			prdimg := model.NewPrdimg()
			filepath := picpath + "/" + picname
			prdimgAttr, _ := prdimg.ParseImgAttribute(filepath, picname)
			/*
				if err != nil {
					fmt.Println(err)
				} else {
			*/
			fmt.Println(prdimgAttr)
			spfn := strings.Split(picname, ".")
			prdimgAttr.Imghash = spfn[0]
			fmt.Println(prdimgAttr)
			prdimg.BeforeCreate()
			fmt.Println("Hhh")
			prdimg.Attribute = prdimgAttr
			fmt.Println("bbb")
			if drr := prdimg.Save(); drr == nil {
				fmt.Println("ccc")
				prdt.UnsetAttribute()
				fmt.Println("ddd")
				prdt.Attribute.Pimg = []string{spfn[0]}
				fmt.Println("eee")
				fmt.Println(prdt.Attribute)
				fmt.Println("fff")
			} else {
				fmt.Println(drr)
			}

			/*}*/
		}
	} else {
		fmt.Println("Paramate NOT set")
	}

	appC.render("product/add", rPrm{
		"title": "Procuct/Add",
		//"person": appC.mContext.GetParam("qp"),
	})
}

func (appC *ProductController) Target(fHeader *multipart.FileHeader, bytes *[]byte) (filename string, basepath string, err error) {
	filename = fHeader.Filename
	if filename == "" {
		return "", "", errors.New("FileHeader can not get filename")
	}

	extname := path.Ext(filename)
	sd5 := md5.Sum(*bytes)
	md5sum := hex.EncodeToString(sd5[:])
	filename = md5sum + extname
	hashPath := components.PrdimgHashPath(md5sum)
	basepath = model.PRDIMG_IMG_BASEPATH + "/" + hashPath
	err = nil
	return
}
