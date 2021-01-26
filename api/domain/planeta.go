package domain

import (
	"time"

	"github.com/jeansferreira/star-wars-planets-api/helpers"
	"gopkg.in/mgo.v2/bson"
)

// The representation of a created planet
type Planeta struct {
	ID        bson.ObjectId `bson:"_id"          json:"id,omitempty"`
	Nome      string        `bson:"nome"         json:"nome,omitempty"`
	Clima     string        `bson:"clima"      	 json:"clima,omitempty"`
	Terreno   string        `bson:"terreno"      json:"terreno,omitempty"`
	Count     int           `bson:"count"        json:"count"`
	CreatedAt *time.Time    `bson:"created_at"   json:"created_at,omitempty"`
	UpdatedAt *time.Time    `bson:"updated_at"   json:"updated_at,omitempty"`
	DeletedAt *time.Time    `bson:"deleted_at"   json:"deleted_at,omitempty"`
}

// The representation of a potential planet
type CriarPlaneta struct {
	Nome    string `bson:"nome"         json:"nome"`
	Clima   string `bson:"clima"      	json:"clima"`
	Terreno string `bson:"terreno"      json:"terreno"`
}

// Warning about timezone issues:
// https://stackoverflow.com/questions/44873825/how-to-get-timestamp-of-utc-time-with-golang
func (c *CriarPlaneta) ToBson() bson.M {
	return bson.M{
		"nome":       c.Nome,
		"clima":      c.Clima,
		"terreno":    c.Terreno,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
}

func (c *Planeta) ToBson() bson.M {
	return bson.M{
		"nome":       c.Nome,
		"clima":      c.Clima,
		"terreno":    c.Terreno,
		"created_at": c.CreatedAt,
		"updated_at": time.Now(),
	}
}

func (c *CriarPlaneta) Me() bson.M {
	return bson.M{
		"nome": c.Nome,
	}
}

func (c *Planeta) Me() bson.M {
	return bson.M{
		"_id": c.ID,
	}
}

func (c *CriarPlaneta) IsValid() (bool, error) {
	if c.Nome == "" {
		return false, helpers.NewError("Planets must have a name")
	}

	if c.Terreno == "" {
		return false, helpers.NewError("Planets must have a terrain")
	}

	if c.Clima == "" {
		return false, helpers.NewError("Planets must have a weather")
	}

	return true, nil
}

func (c *Planeta) IsValid() (bool, error) {
	if c.Nome == "" {
		return false, helpers.NewError("Planets must have a name")
	}

	if c.Terreno == "" {
		return false, helpers.NewError("Planets must have a terrain")
	}

	if c.Clima == "" {
		return false, helpers.NewError("Planets must have a weather")
	}

	return true, nil
}

// Same idea from gorm, when "tableName" implements the gorm.Tabler interface
func (*Planeta) CollectionName() string {
	return "planetas"
}
