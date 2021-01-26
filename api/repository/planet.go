package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/jeansferreira/api-b2w-planetas/domain"
	"github.com/jeansferreira/api-b2w-planetas/helpers"
	"gopkg.in/mgo.v2/bson"
)

type Planet struct{}

func (*Planet) Search(offset int, limit int, name string) ([]domain.Planet, error) {
	var planets []domain.Planet
	var d = domain.Planet{}

	Mongo.Connect()

	c := Mongo.db.C(d.CollectionName())

	if name == "" {
		err := c.Find(bson.M{}).Skip(offset).Limit(limit).All(&planets)

		if err != nil {
			return nil, err
		}

	} else {
		err := c.Find(bson.M{"name": name}).All(&planets)

		if err != nil {
			log.Printf("Not found: %s", err)
		}
	}

	return planets, nil
}

func (*Planet) Lookup(id bson.ObjectId) (*domain.Planet, error) {
	var planet domain.Planet

	Mongo.Connect()

	c := Mongo.db.C(planet.CollectionName())
	err := c.Find(bson.M{"_id": id}).One(&planet)

	if err != nil {
		return nil, err
	}

	return &planet, nil
}

func (*Planet) Create(newPlanet domain.CreatePlanet) (*domain.Planet, error) {
	var planet domain.Planet

	Mongo.Connect()
	c := Mongo.db.C(planet.CollectionName())

	// Check if planet already exist
	count, err := c.Find(newPlanet.Me()).Count()

	if count > 0 {
		return nil, helpers.NewError("Planet already exists")
	}

	err = c.Insert(newPlanet.ToBson())
	err = c.Find(newPlanet.Me()).One(&planet)

	if err != nil {
		return nil, err
	}

	return &planet, nil
}

func (*Planet) Update(planet domain.Planet, id string) (*domain.Planet, error) {
	var updatedPlanet domain.Planet

	Mongo.Connect()
	c := Mongo.db.C(planet.CollectionName())

	var oldVersion domain.Planet

	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&oldVersion)

	if err != nil {
		return nil, err
	}

	mTime := time.Now()

	updatedPlanet = domain.Planet{
		ID:        bson.ObjectIdHex(id),
		Name:      planet.Name,
		Terrain:   planet.Terrain,
		Weather:   planet.Weather,
		CreatedAt: oldVersion.CreatedAt,
		UpdatedAt: &mTime,
	}

	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, updatedPlanet.ToBson())

	if err != nil {
		return nil, err
	}

	return &updatedPlanet, nil
}

func (*Planet) Delete(id string) error {
	var planet domain.Planet

	Mongo.Connect()
	c := Mongo.db.C(planet.CollectionName())
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		fmt.Printf("Error on remove planet %s", id)
		return err
	}

	return nil
}
