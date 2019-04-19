package gateways

import (
	"fmt"

	"github.com/ashkarin/ashkarin-pizza-shop/pkg/ingredient"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mgoGateway struct {
	session    *mgo.Session
	db         *mgo.Database
	collection *mgo.Collection
}

// NewMongoDbGateway create a storage gateway to the MongoDB
func NewMongoDbGateway(server, port, username, password, database, collection string) (ingredient.StorageGateway, error) {
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

func (s *mgoGateway) GetByID(id string) (*ingredient.Ingredient, error) {
	ingredient := &ingredient.Ingredient{}
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	if err := s.collection.Find(query).One(&ingredient); err != nil {
		return nil, err
	}
	return ingredient, nil
}

func (s *mgoGateway) DeleteByID(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	return s.collection.Remove(query)
}

func (s *mgoGateway) Store(r *ingredient.Ingredient) error {
	return s.collection.Insert(r)
}

func (s *mgoGateway) Update(r *ingredient.Ingredient) error {
	var ID string
	switch v := r.ID.(type) {
	case string:
		ID = v
	case bson.ObjectId:
		ID = v.Hex()
	}

	query := bson.M{"_id": bson.ObjectIdHex(ID)}
	change := bson.M{"$set": bson.M{
		"name":       r.Name,
		"vegetarian": r.Vegetarian,
		"price":      r.Price,
	}}
	return s.collection.Update(query, change)
}

func (s *mgoGateway) Delete(r *ingredient.Ingredient) error {
	query := bson.M{"_id": bson.ObjectIdHex(r.ID.(string))}
	return s.collection.Remove(query)
}

func (s *mgoGateway) Search(pattern string) ([]*ingredient.Ingredient, error) {
	var ingredients []*ingredient.Ingredient
	regex := bson.M{"$regex": bson.RegEx{Pattern: pattern}}
	if err := s.collection.Find(bson.M{"name": regex}).All(&ingredients); err != nil {
		return nil, err
	}

	for i, ingredient := range ingredients {
		ingredients[i].ID = ingredient.ID.(bson.ObjectId).Hex()
	}
	return ingredients, nil
}

func (s *mgoGateway) Clean() error {
	return s.collection.Remove(bson.M{})
}
