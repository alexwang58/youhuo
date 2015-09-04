package components

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gorilla/sessions"
	"net/http"
)

func EncryptPassword(pwd string) string {
	data := []byte(pwd)
	sd5 := md5.Sum(data)
	return hex.EncodeToString(sd5[:])
}

type YouhuoSCession struct {
	CookieStore *sessions.CookieStore
	Session     *sessions.Session
	w           http.ResponseWriter
	r           *http.Request
}

var sessionSecret = "ek^golala"
var youhuoSessionName = "youhuo.com"
var KeySessionUid = "lkwen_luisj"
var CookieStore *sessions.CookieStore

func NewYouhuoSession(w http.ResponseWriter, r *http.Request) (*YouhuoSCession, error) {
	yhSc := &YouhuoSCession{
		CookieStore: sessions.NewCookieStore([]byte(sessionSecret)),
		w:           w,
		r:           r,
	}

	sssn, err := yhSc.CookieStore.Get(r, youhuoSessionName)
	if err != nil {
		return nil, err
	}
	yhSc.Session = sssn

	return yhSc, nil
}

func (this *YouhuoSCession) Set(key string, value interface{}) error {
	this.Session.Values[key] = value
	return this.Save()
}

func (this *YouhuoSCession) Get(key string) interface{} {
	return this.Session.Values[key]
}

func (this *YouhuoSCession) Save() error {
	return this.CookieStore.Save(this.r, this.w, this.Session)
}

func (this *YouhuoSCession) Del(key string) {
	delete(this.Session.Values, key)
	this.CookieStore.Save(this.r, this.w, this.Session)
}

/*
func SessionStart() {
	if CookieStore == nil {
		CookieStore = sessions.NewCookieStore([]byte(sessionSecret))
	}
}

func GetProfileSession(r *http.Request) (*sessions.Session, error) {
	if CookieStore == nil {
		SessionStart()
	}
	return CookieStore.Get(r, youhuoSessionName)
}

func SetCookie(w http.ResponseWriter, r *http.Request) error {
	CookieStore.Save(r, w, session)
}
*/
