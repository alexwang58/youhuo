package controller

import (
	//	"errors"
	"fmt"
	"github.com/alexwang58/youhuo/components"
	"github.com/alexwang58/youhuo/model"
	"net/http"
	"strconv"
	//	"reflect"
)

type PersonController struct {
	Controller
}

func PersonCtrl(w http.ResponseWriter, r *http.Request) error {
	appc := new(PersonController)
	if err := appc.Active(w, r, appc); err != nil {
		return nil
	}
	return nil
}

// MUST implement this func
// It`s a action map for this controller
func (appC *PersonController) actionList() map[string]func() {
	return map[string]func(){
		"create": appC.Create,
		"signin": appC.Signin,
		"logout": appC.Logout,
		"home":   appC.Home,
	}
}

func (appC *PersonController) dispatchError() (string, string, error) {
	return "site", "index", nil
}

// Action for product/list/xxxx
func (appC *PersonController) Home() {
	/*
		psn := model.NewPerson()
		phone := appC.mContext.GetPost("phone")
		paswd := appC.mContext.GetPost("paswd")
		person, err := psn.AccessExists(phone, paswd)
		fmt.Println(person)
	*/
	uid, ok := appC.yhSc.Get(components.KeySessionUid).(int)
	if !ok {
		fmt.Println("Uid is NOT type of string")
		appC.Redirect("person", "signin")
		return
	}

	psn := model.NewPerson()
	person, err := psn.GetByUid(uid)
	if err != nil {
		fmt.Println("User not found")
		appC.Redirect("person", "signin")
		return
	}

	appC.render("person/home", rPrm{
		"title":   "Person/Home",
		"Profile": person,
	})
}

// Action for product/list/xxxx
func (appC *PersonController) Signin() {
	fmt.Println("SIGNING")
	psn := model.NewPerson()
	phone := appC.mContext.GetPost("phone")
	paswd := appC.mContext.GetPost("paswd")
	if phone != "" && paswd != "" {
		fmt.Println(phone + "||" + paswd)
		person, err := psn.AccessExists(phone, paswd)
		fmt.Println(person)
		if err == nil {
			fmt.Println(components.KeySessionUid + "/" + strconv.Itoa(person.Uid))
			appC.yhSc.Set(components.KeySessionUid, person.Uid)
			appC.Redirect("person", "home")
			return
		}
	}

	appC.render("person/signin", rPrm{
		"title": "Person/Signin",
		// "uid":   person.Uid,
		// "nick":  person.Nick,
	})
}

// Action for product/list/xxxx
func (appC *PersonController) Logout() {
	appC.yhSc.Del(components.KeySessionUid)
	appC.render("person/home", rPrm{
		"title": "Person/home",
		//"person": appC.mController.param["qp"],
	})
}

// Action for product/list/xxxx
func (appC *PersonController) Create() {
	phone := appC.mContext.GetPost("phone")
	paswd := appC.mContext.GetPost("paswd")
	fmt.Println("ParamResived:")
	fmt.Println(phone)
	fmt.Println(paswd)

	psn := model.NewPerson()
	psn.Attribute.Uid = components.PSN_AUTOINC.Id()
	// nick initialed as phone
	psn.Attribute.Nick = phone
	psn.Attribute.Phone = phone
	psn.Attribute.Passwd = components.EncryptPassword(paswd)
	psn.Attribute.Lst_tch = 1231231897
	psn.Attribute.C_time = 1231231897
	psn.Save()

	appC.render("person/create", rPrm{
		"title": "Person/Create",
		//"person": appC.mController.param["qp"],
	})
}
