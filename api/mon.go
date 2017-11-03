package api

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DbHandler struct {
	*mgo.Session
}

type URL struct {
	ID      int    `bson:"_id"`
	LongURL string `bson:"long_url"`
	Clicks  int `bson:"clicks"`
}

type Counter struct {
	ID	int `bson:"_id"`
	Seq int `bson:"seq"`
}

type URLDbHandler interface {
	GetURLById(int) (URL, error)
	GetURLByLongURL(string) (URL, error)
	GetNextId() (int, error)
	UpdateClickCount(int) error
	AddURL(URL) error
}

func NewDbHandler(databaseEndpoint string) (*DbHandler, error) {
	session, err := mgo.Dial(databaseEndpoint)
	return &DbHandler{
		Session: session,
	}, err
}

func (handler *DbHandler) GetURLById(id int) (URL, error) {
	session := handler.Session.Copy()
	defer session.Close()
	url := URL{}
	err := session.DB("sahajr-website").C("urls").Find(bson.M{"_id": id}).One(&url)
	return url, err
}

func (handler *DbHandler) GetURLByLongURL(longURL string) (URL, error) {
	session := handler.Session.Copy()
	defer session.Close()
	url := URL{}
	err := session.DB("sahajr-website").C("urls").Find(bson.M{"long_url": longURL}).One(&url)
	return url, err
}

func (handler *DbHandler) AddURL(url URL) error {
	session := handler.Session.Copy()
	defer session.Close()
	return session.DB("sahajr-website").C("urls").Insert(url)
}

func (handler *DbHandler) GetNextId() (int, error) {
	session := handler.Session.Copy()
	defer session.Close()
	counter := Counter{}
	increment := mgo.Change{
		Update: bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
	}
	_, err := session.DB("sahajr-website").C("counters").Find(bson.M{"_id": "url_count"}).Apply(increment, &counter)
	if err != nil {
		return -1, err
	}
	return counter.Seq, nil
}

func (handler *DbHandler) UpdateClickCount(documentID int) error {
	session := handler.Session.Copy()
	defer session.Close()
	url := URL{}
	increment := mgo.Change{
		Update: bson.M{"$inc": bson.M{"clicks": 1}},
		ReturnNew: true,
	}
	_, err := session.DB("sahajr-website").C("urls").Find(bson.M{"_id": documentID}).Apply(increment, &url)
	return err
}