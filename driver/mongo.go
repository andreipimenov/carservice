package driver

import mgo "gopkg.in/mgo.v2"

type Mongo struct {
	Session *mgo.Session
	DB      string
}

func NewMongo(url string, db string) (*Mongo, error) {
	s, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	m := &Mongo{
		Session: s,
		DB:      db,
	}
	return m, nil
}

func (m *Mongo) Get(c string, q interface{}, r interface{}) error {
	s := m.Session.Copy()
	defer s.Close()
	err := s.DB(m.DB).C(c).Find(q).One(r)
	return err
}
