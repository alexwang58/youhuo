package components

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBMongo struct {
	user    string
	pass    string
	host    string
	port    string
	session *mgo.Session
}

// New returns an DBMongo object
func NewMgo(_host string, _port string, _user string, _pass string) *DBMongo {
	return &DBMongo{host: _host, port: _port, user: _user, pass: _pass}
}

func (dbM *DBMongo) connect() error {
	var url string
	if dbM.user == "" || dbM.pass == "" {
		url = "mongodb://" + dbM.host + ":" + dbM.port
	} else {
		url = "mongodb://" + dbM.user + ":" + dbM.pass + "&" + dbM.host + ":" + dbM.port
	}
	var err error
	if dbM.session, err = mgo.Dial(url); err != nil {
		return err
	}

	return nil
}

func (dbM *DBMongo) C(dbs string, collenction string) *mgo.Collection {
	if dbM.session == nil {
		dbM.connect()
	}
	return dbM.session.DB(dbs).C(collenction)
}

func (dbM *DBMongo) Close() {
	dbM.session.Close()
}

func (dbM *DBMongo) FindOne(c *mgo.Collection, query map[string]interface{}, result interface{}) error {
	return c.Find(bson.M(query)).One(result)
}

func (dbM *DBMongo) FindByObid(objid string, c *mgo.Collection, result interface{}) error {
	query := bson.M{"_id": bson.ObjectIdHex(objid)}
	return dbM.FindOne(c, query, result)
}

/*
func (dbM *DBMongo) Update(c *mgo.Collection, condition map[string]interface{}, dset interface{}) error {
	return c.Find(bson.M(query)).One(result)
}
*/
