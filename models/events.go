package models

import (
	"time"
)

type Event struct {
	DomainID    string    `bson:domain_id`
	TypeID      string    `bson:type_id`
	CreatedAt   time.Time `bson:created_at`
	Description string    `bson:description`
	// Source
}

type Type struct {
	Name      string `bson:name`
	Shortname string `bson:short_name`
	Color     string `bson:color` // Show a color for the type of event
}
