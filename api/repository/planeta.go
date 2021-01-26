package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/jeansferreira/api-b2w-planetas/domain"
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
	var planeta domain.Planeta

	Mongo.Connect()

	c := Mongo.db.C(planeta.CollectionName())
	err := c.Find(bson.M{"_id": id}).One(&planeta)

	if err != nil {
		return nil, err
	}

	return &planeta, nil
}

func (*Planeta) Create(newPlaneta domain.CriarPlaneta) (*domain.Planeta, error) {
	var planeta domain.Planeta

	Mongo.Connect()
	c := Mongo.db.C(planeta.CollectionName())

	// Check if planet already exist
	count, err := c.Find(newPlaneta.Me()).Count()

	if count > 0 {
		return nil, helpers.NewError("Planet already exists")
	}

	err = c.Insert(newPlaneta.ToBson())
	err = c.Find(newPlaneta.Me()).One(&planeta)

	if err != nil {
		return nil, err
	}

	return &planeta, nil
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

	updatedPlaneta = domain.Planeta{
		ID:        bson.ObjectIdHex(id),
		Nome:      planeta.Nome,
		Terreno:   planeta.Terreno,
		Clima:     planeta.Clima,
		CreatedAt: oldVersion.CreatedAt,
		UpdatedAt: &mTime,
	}

	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, updatedPlaneta.ToBson())

	if err != nil {
		return nil, err
	}

	return &updatedPlaneta, nil
}

func (*Planeta) Delete(id string) error {
	var planeta domain.Planeta

	Mongo.Connect()
	c := Mongo.db.C(planeta.CollectionName())
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		fmt.Printf("Error on remove planet %s", id)
		return err
	}

	return nil
}
