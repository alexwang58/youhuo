package controller

import (
	//"errors"
	//	"fmt"
	//"github.com/alexwang58/youhuo/components"
	"net/http"
)

type SiteController struct {
	Controller
}

/*
func (mC *Controller) SiteController(w http.ResponseWriter, r *http.Request) error {
	appc := new(SiteController)
	appc.mController = mC
	if err := mC.Active(w, r, appc); err != nil {
		return nil
	}

	return nil
}
*/
func SiteCtrl(w http.ResponseWriter, r *http.Request) error {
	appc := new(SiteController)
	//appc.mC = new(Controller)
	//appc.mContext = appc.mC.mContext
	if err := appc.Active(w, r, appc); err != nil {
		return nil
	}
	return nil
}

func (appC *SiteController) dispatchError() (string, string, error) {
	return "site", "index", nil
	//return "", "", errors.New("Dispatcher no result")
}

// MUST implement this func
// It`s a action map for this controller
func (appC *SiteController) actionList() map[string]func() {
	return map[string]func(){
		"index": appC.Index,
		"home":  appC.Home,
	}
}

// Action for product/list/xxxx
func (appC *SiteController) Index() {
	appC.render("site/index", rPrm{
		"title": "site/index",
	})
}

// Action for product/list/xxxx
func (appC *SiteController) Home() {
	appC.render("site/home", rPrm{
		"title": "site/home",
	})
}
