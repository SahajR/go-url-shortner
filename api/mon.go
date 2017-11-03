package api

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DbHandler struct {
	DatabaseName string
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

func NewDbHandler(databaseEndpoint string, databaseName string) (*DbHandler, error) {
	session, err := mgo.Dial(databaseEndpoint)
	return &DbHandler{
		DatabaseName:databaseName,
		Session: session,
	}, err
}

// Get the URL document associated with the given ID in the mongo collection
func (handler *DbHandler) GetURLById(id int) (URL, error) {
	session := handler.Session.Copy()
	defer session.Close()
	url := URL{}
	err := session.DB(handler.DatabaseName).C("urls").Find(bson.M{"_id": id}).One(&url)
	return url, err
}

// Get the URL document associated with the given long URL in the mongo collection
func (handler *DbHandler) GetURLByLongURL(longURL string) (URL, error) {
	session := handler.Session.Copy()
	defer session.Close()
	url := URL{}
	err := session.DB(handler.DatabaseName).C("urls").Find(bson.M{"long_url": longURL}).One(&url)
	return url, err
}

// Adds the long URL to the database
func (handler *DbHandler) AddURL(url URL) error {
	session := handler.Session.Copy()
	defer session.Close()
	return session.DB(handler.DatabaseName).C("urls").Insert(url)
}

// Get the next ID in the database
func (handler *DbHandler) GetNextId() (int, error) {
	session := handler.Session.Copy()
	defer session.Close()
	counter := Counter{}
	increment := mgo.Change{
		Update: bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
	}
	_, err := session.DB(handler.DatabaseName).C("counters").Find(bson.M{"_id": "url_count"}).Apply(increment, &counter)
	if err != nil {
		return -1, err
	}
	return counter.Seq, nil
}

// Updates the click count of the link associated with the short code.
func (handler *DbHandler) UpdateClickCount(documentID int) error {
	session := handler.Session.Copy()
	defer session.Close()
	url := URL{}
	increment := mgo.Change{
		Update: bson.M{"$inc": bson.M{"clicks": 1}},
		ReturnNew: true,
	}
	_, err := session.DB(handler.DatabaseName).C("urls").Find(bson.M{"_id": documentID}).Apply(increment, &url)
	return err
}
