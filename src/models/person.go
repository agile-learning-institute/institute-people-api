package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserName    string             `json:"userName,omitempty"`
	FirstName   string             `json:"firstName,omitempty"`
	LastName    string             `json:"lastName,omitempty"`
	Rolls       []string           `json:"rolls,omitempty"`
	Description string             `json:"description,omitempty"`
	Status      string             `json:"status,omitempty"`
	Member      bool               `json:"member,omitempty"`
	Mentor      bool               `json:"mentor,omitempty"`
	Donor       bool               `json:"donor,omitempty"`
	Contact     bool               `json:"contact,omitempty"`
	Title       string             `json:"title,omitempty"`
	Cadence     string             `json:"cadence,omitempty"`
	Email       string             `json:"eMail,omitempty"`
	GitHub      string             `json:"gitHub,omitempty"`
	Phone       string             `json:"phone,omitempty"`
	Device      string             `json:"device,omitempty"`
	Location    string             `json:"location,omitempty"`
	MentorId    primitive.ObjectID `json:"mentorId,omitempty"`
	PartnerId   primitive.ObjectID `json:"partnerId,omitempty"`
	LastSaved   *BreadCrumb        `json:"lastSaved,omitempty"`
}
