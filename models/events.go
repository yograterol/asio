package models

import "time"

type EventDomain struct {
	DomainID    string    `bson:"domain_id"`
	TypeID      string    `bson:"type_id"`
	CreatedAt   time.Time `bson:"created_at"`
	Description string    `bson:"description"`
	// Source
}

// For the future
type EventServices struct { // Services = Daemons
	DaemonID    string    `bson:"daemon_id"`
	TypeID      string    `bson:"type_id"`
	CreatedAt   time.Time `bson:"created_at"`
	Description string    `bson:"description"`
	Environment string    `bson:"env"` // Shoul be a JSON field with all env data where run the daemon
	// Source
}

type Type struct {
	Name      string `bson:"name"`
	Shortname string `bson:"short_name"`
	Color     string `bson:"color"` // Show a color for the type of event
}
