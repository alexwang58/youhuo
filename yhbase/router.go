package yhbase

import (
	//"fmt"
	"github.com/gorilla/mux"
	//"log"
	//"html"
	cctr "github.com/alexwang58/youhuo/controller"
	"net/http"
)

var router = mux.NewRouter()

func initRouter() {
	r := router
	r.StrictSlash(true)
	//mC := new(cctr.Controller)
	//mC.SetRoute(r)

	r.Handle("/", cctr.CtrFunc(cctr.SiteCtrl))
	r.Handle("/product/{action}/{qp:[0-9]*}", cctr.CtrFunc(cctr.ProductCtrl)).Name("ls_prsn_prd")
	r.Handle("/product/{action}", cctr.CtrFunc(cctr.ProductCtrl))
	r.Handle("/product", cctr.CtrFunc(cctr.ProductCtrl))

	r.Handle("/", cctr.CtrFunc(cctr.PersonCtrl))
	r.Handle("/person/{action}/{qp:[0-9]*}", cctr.CtrFunc(cctr.PersonCtrl))
	r.Handle("/person/{action}", cctr.CtrFunc(cctr.PersonCtrl))
	r.Handle("/person", cctr.CtrFunc(cctr.PersonCtrl))

	r.Handle("/site", cctr.CtrFunc(cctr.SiteCtrl))
	r.Handle("/site/{action}", cctr.CtrFunc(cctr.SiteCtrl))

	r.Handle("/{path:.*}", cctr.CtrFunc(cctr.NotFoundHandler))
	http.Handle("/", r)
}
