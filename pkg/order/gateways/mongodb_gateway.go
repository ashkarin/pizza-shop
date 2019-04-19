package gateways

import (
	"fmt"

	m "github.com/ashkarin/ashkarin-pizza-shop/pkg/order"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mgoGateway struct {
	session    *mgo.Session
	db         *mgo.Database
	collection *mgo.Collection
}

// NewMongoDbGateway create a storage gateway to the MongoDB
func NewMongoDbGateway(server, port, username, password, database, collection string) (m.StorageGateway, error) {
	url := fmt.Sprintf("%s:%s", server, port)
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(database)
	gw := &mgoGateway{
		session:    session,
		db:         db,
		collection: db.C(collection),
	}
	return gw, nil
}

func (s *mgoGateway) GetByID(id string) (*m.Order, error) {
	order := &m.Order{}
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	if err := s.collection.Find(query).One(&order); err != nil {
		return nil, err
	}
	return order, nil
}

func (s *mgoGateway) DeleteByID(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	return s.collection.Remove(query)
}

func (s *mgoGateway) Place(r *m.Order) (string, error) {
	r.ID = bson.NewObjectId()
	err := s.collection.Insert(r)
	return r.ID.(bson.ObjectId).Hex(), err
}

func (s *mgoGateway) Update(r *m.Order) error {
	var ID string
	switch v := r.ID.(type) {
	case string:
		ID = v
	case bson.ObjectId:
		ID = v.Hex()
	}

	query := bson.M{"_id": bson.ObjectIdHex(ID)}
	change := bson.M{"$set": bson.M{
		"acceptedAt":  r.AcceptedAt,
		"completedAt": r.CompletedAt,
		"content":     r.Content,
		"status":      r.Status,
	}}
	return s.collection.Update(query, change)
}

func (s *mgoGateway) Delete(r *m.Order) error {
	query := bson.M{"_id": bson.ObjectIdHex(r.ID.(string))}
	return s.collection.Remove(query)
}

func (s *mgoGateway) Clean() error {
	return s.collection.Remove(bson.M{})
}
