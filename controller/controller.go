package controller

import (
	"errors"
	"fmt"
	"github.com/alexwang58/youhuo/components"
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	"net/http"
	"strings"
)

type Controller struct {
	mContext *components.Context
	router   *mux.Router
	status   int
	yhSc     *components.YouhuoSCession
}

type AppController interface {
	actionList() map[string]func()
	dispatchError() (string, string, error)
}

func NewController() *Controller {
	return &Controller{}
}

// child controller Must implement function:
// func (mC *Controller) actionList() *actionList{}
// this structure defined for the return value
//type actionList map[string]func()

// this type defined for parameter data for mC.render()
type rPrm map[string]interface{}

// handlerFunc adapts a function to an http.Handler.
type CtrFunc func(http.ResponseWriter, *http.Request) error

func (f CtrFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServeHTTP")
	/*
		if session, err := components.GetProfileSession(r); err == nil {
			fmt.Println(session)
			components.CookieStore.Save(r, w, session)
		} else {
			fmt.Println("[Error]SessionSaveFaild!!!")
		}
	*/
	err := f(w, r)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

// CtrFunc adapts a function to an http.Handler. Use in router.go
func (mC *Controller) Active(w http.ResponseWriter, r *http.Request, appC AppController) error {
	fmt.Println("Active")
	var err error
	mC.mContext = components.NewContext()
	mC.mContext.SetHost(r.Host)
	mC.mContext.SetResponseWriter(w)
	//r.ParseForm()
	mC.mContext.SetResquest(r)
	fmt.Println("BeginRun")
	yhsc, err := components.NewYouhuoSession(w, r)
	if err == nil {
		mC.yhSc = yhsc
	} else {
		fmt.Println("[Error]SessionSaveFaild!!!")
	}
	mC.mContext.SetLayout("main")
	fmt.Println("RequestUri:" + r.RequestURI)
	mC.setStatus(http.StatusOK)
	mC.mContext.SetQueryParam(mC.getVars(r))
	mC.mContext.SetRemoteAddr(r.RemoteAddr)
	fmt.Println("[ActionRes] Begin: " + r.RequestURI)
	_name, _action, err := parseAction(r.RequestURI)
	fmt.Printf("[ActionRes] App: %v, Action: %v\n", _name, _action)
	mC.mContext.SetName(_name)
	mC.mContext.SetAction(_action)
	if err != nil {
		ErrorHandler(w, r, rPrm{
			"mesg": err.Error(),
		})
		return err
	}
	if err = mC.dispatchAction(mC.mContext.GetName(), mC.mContext.GetAction(), appC); err != nil {
		if _name, _action, _err := appC.dispatchError(); _err == nil {
			_rdUrl := "http://" + mC.mContext.GetHost() + "/" + _name + "/" + _action
			mC.mContext.SetName(_name)
			mC.mContext.SetAction(_action)
			fmt.Println("[Redirect] :" + _rdUrl)
			http.Redirect(mC.mContext.GetResponseWriter(), mC.mContext.GetResquest(), _rdUrl, http.StatusFound)
			return nil
		}
		mC.render("error", rPrm{
			"mesg": err.Error(),
		})
	}
	return nil
}

// Default template layout is main.html
// Call this function to change it
// Example: mC.setLayout("newlayout")

func (mC *Controller) setLayout(layout string) {
	components.Layout = layout
}

// Default view path is YH_DOCROOT."/view"
// Call this function to change it
// Example: mC.setViewDir("newdir")
func (mC *Controller) setViewDir(dirname string) {
	components.ViewDir = dirname
}

func (mC *Controller) setStatus(status int) {
	mC.status = status
}

func (mC *Controller) SetRoute(r *mux.Router) {
	mC.router = r
}

// controller can get vars from mux.Vars
func (mC *Controller) getVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

// Dispatch action to controllers
func (mC *Controller) dispatchAction(name string, action string, appC AppController) error {
	if funcs, ok := (appC.actionList())[action]; ok {
		funcs()
		return nil
	}
	return errors.New("Action Not Exists: " + name + "/" + action)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) error {
	return components.RenderTemplate(w, http.StatusNotFound, "site/404", nil)
}

// When handle error page, paramete data["mesg"] MUST be set
func ErrorHandler(w http.ResponseWriter, r *http.Request, data interface{}) error {
	/*
		if _, ok := data["mesg"]; !ok {
			data["mesg"] = "Message Not Set"
		}
	*/
	return components.RenderTemplate(w, http.StatusInternalServerError, "site/error", data)
}

func (mC *Controller) Redirect(name string, action string) {
	url := "http://" + mC.mContext.GetHost() + "/" + name + "/" + action
	http.Redirect(mC.mContext.GetResponseWriter(), mC.mContext.GetResquest(), url, http.StatusFound)
}

func (mC *Controller) render(name string, data interface{}) error {
	fmt.Printf("Render[%v]: %v\n", mC.status, name)
	fmt.Println(data)
	err := components.RenderTemplate(mC.mContext.GetResponseWriter(), mC.status, name, data)
	if err != nil {
		fmt.Errorf("[RenderError] " + err.Error())
		return err
	}
	return nil
}

// parse uri, return controllerName , action
// normal format looks like controller/action
// if uri is "/" will return site/index
// if uri is "controler", do not contain action name, return controller/index
func parseAction(uri string) (string, string, error) {
	// "http://example.com/" OR "http://example.com/"
	// render to site/index
	if uri == "/" {
		fmt.Println("Redict /")
		return "site", "index", nil
	}

	uriComp := strings.Split(uri, "/")
	//fmt.Println(uriComp[1], uriComp[2])
	if len(uriComp) >= 3 && len(uriComp[1]) > 0 && len(uriComp[2]) > 0 {
		fmt.Println("Normal ", uriComp[1], uriComp[2])
		return uriComp[1], uriComp[2], nil
	}

	// "http://example.com/mail", render to mail/index
	if len(uriComp[1]) > 0 {
		fmt.Println("Redict ", uriComp[1], "index")
		return uriComp[1], "index", nil
	}

	return "", "", errors.New("Invalid Router: " + uri)
}
