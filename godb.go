package godb

import (
    "fmt"
    "gopkg.in/mgo.v2"
)

var database string;

var conn *MongoConnection

func init() {
    conn = new(MongoConnection)
}

func GetMongoConnection() *MongoConnection {
    return conn
}

func GetDatabase() (*mgo.Database, *mgo.Session) {
    s := conn.session.Copy()
    return s.DB(database), s
}

type MongoConnection struct {
    session *mgo.Session
}

func (m *MongoConnection) Connect(user, password string, host, database string) {
    database = database;
    var uri string;
    if len(user) > 0 && len(password) > 0 {
        uri = fmt.Sprintf("mongodb://%s:%s@%s:27017/%s", user, password, host, database)
    } else {
        uri = fmt.Sprintf("mongodb://%s:27017/%s", host, database)
    }

    session, err := mgo.Dial(uri)
    if err != nil {
        fmt.Println("Authentication failed!")
        panic(err)
    }

    session.SetMode(mgo.Monotonic, true)
    m.session = session
}

func (m *MongoConnection) Close() {
    m.session.Close()
}
