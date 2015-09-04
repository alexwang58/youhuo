package controller

import (
	//	"errors"
	//"fmt
	//"github.com/alexwang58/youhuo/components"
	//"github.com/alexwang58/youhuo/model"
	"net/http"
	//	"reflect"
)

type AuthController struct {
	Controller
}

func AuthCtrl(w http.ResponseWriter, r *http.Request) error {
	appc := new(AuthController)
	if err := appc.Active(w, r, appc); err != nil {
		return nil
	}
	return nil
}

// MUST implement this func
// It`s a action map for this controller
func (appC *AuthController) actionList() map[string]func() {
	return map[string]func(){
		"auth":   appC.Auth,
		"check":  appC.Check,
		"expire": appC.Expire,
	}
}

func (appC *AuthController) dispatchError() (string, string, error) {
	return "site", "index", nil
}

// Action for product/list/xxxx
func (appC *AuthController) Auth() {
	appC.render("person/auth", rPrm{
		"title": "Auth/Auth",
		//"person": appC.mController.param["qp"],
	})
}

// Action for product/list/xxxx
func (appC *AuthController) Check() {
	appC.render("person/check", rPrm{
		"title": "Auth/Check",
		//"person": appC.mController.param["qp"],
	})
}

// Action for product/list/xxxx
func (appC *AuthController) Expire() {
	appC.render("person/expire", rPrm{
		"title": "Auth/Expire",
		//"person": appC.mController.param["qp"],
	})
}
