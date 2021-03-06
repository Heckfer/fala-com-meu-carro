package model

import (
	"regexp"
	"time"

	"github.com/drborges/appx"

	"appengine/datastore"
)

type Post struct {
	appx.Model

	CarPlate  string    `json:"car_plate"`
	Message   string    `json:"message"`
	UserId    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	Flagged   bool      `json:"flagged"`
	Deleted   bool      `json:"deleted"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
}

func (post *Post) KeySpec() *appx.KeySpec {
	return &appx.KeySpec{
		Kind:       "Post",
		Incomplete: true,
	}
}

var Posts = struct {
	All           func(string) *datastore.Query
	AllByCarPlate func(string, string) *datastore.Query
}{
	All: func(country string) *datastore.Query {
		return datastore.NewQuery(new(Post).
			KeySpec().Kind).
			Filter("Deleted=", false).
			Filter("Country=", country).
			Order("-CreatedAt")
	},
	AllByCarPlate: func(carPlate string, country string) *datastore.Query {
		return datastore.NewQuery(new(Post).
			KeySpec().Kind).
			Filter("CarPlate=", carPlate).
			Filter("Deleted=", false).
			Filter("Country=", country).
			Order("-CreatedAt")
	},
}

func (post Post) IsValid(country string) (bool, []string) {
	errors := []string{}

	if post.CarPlate == "" {
		errors = append(errors, "Placa não pode ser vazia.")
	}

	if country == "br" && !post.isPlateValid() {
		errors = append(errors, "Placa inválida")
	}

	if post.Message == "" {
		errors = append(errors, "Mensagem não pode ser vazia")
	}

	if post.UserId == "" {
		errors = append(errors, "ID do usuário não pode ser vazio")
	}

	if post.UserName == "" {
		errors = append(errors, "Nome de usuário não pode ser vazio")
	}

	return len(errors) == 0, errors
}

func (post *Post) isPlateValid() bool {
	valid, _ := regexp.MatchString("[A-Z]{3}-[0-9]{4}", post.CarPlate)
	return valid
}
