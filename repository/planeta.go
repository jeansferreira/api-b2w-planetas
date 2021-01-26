package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/jeansferreira/api-b2w-planetas/helpers"
	"gopkg.in/mgo.v2/bson"
)

type Planeta struct{}

func (*Planeta) Search(offset int, limit int, nome string) ([]domain.Planeta, error) {
	var planetas []domain.Planeta
	var d = domain.Planeta{}

	Mongo.Connect()

	c := Mongo.db.C(d.CollectionName())

	if nome == "" {
		err := c.Find(bson.M{}).Skip(offset).Limit(limit).All(&planetas)

		if err != nil {
			return nil, err
		}

	} else {
		err := c.Find(bson.M{"nome": nome}).All(&planetas)

		if err != nil {
			log.Printf("Not found: %s", err)
		}
	}

	return planetas, nil
}

func (*Planeta) Lookup(id bson.ObjectId) (*domain.Planeta, error) {
	var planet domain.Planeta

	Mongo.Connect()

	c := Mongo.db.C(planet.CollectionName())
	err := c.Find(bson.M{"_id": id}).One(&Planeta)

	if err != nil {
		return nil, err
	}

	return &Planeta, nil
}

func (*Planeta) Create(newPlaneta domain.CreatePlaneta) (*domain.Planeta, error) {
	var planeta domain.Planeta

	Mongo.Connect()
	c := Mongo.db.C(Planeta.CollectionName())

	// Check if planet already exist
	count, err := c.Find(newPlaneta.Me()).Count()

	if count > 0 {
		return nil, helpers.NewError("Planet already exists")
	}

	err = c.Insert(newPlaneta.ToBson())
	err = c.Find(newPlaneta.Me()).One(&Planeta)

	if err != nil {
		return nil, err
	}

	return &planet, nil
}

func (*Planeta) Update(planeta domain.Planeta, id string) (*domain.Planeta, error) {
	var updatedPlaneta domain.Planeta

	Mongo.Connect()
	c := Mongo.db.C(planeta.CollectionName())

	var oldVersion domain.Planeta

	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&oldVersion)

	if err != nil {
		return nil, err
	}

	mTime := time.Now()

	updatedPlanet = domain.Planeta{
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
