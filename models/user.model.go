package models // A model is a class / struct containing the attributes that represent columns in the DB table

import (
	"github.com/google/uuid"
	"time"
)

// Create a Go struct that will be transformed into SQL tables:
//   - Create a User struct to represent the SQL table in the Postgres database.
//   - Use UUID (Universally Unique Identifier) as a default value for the ID column.
//   - Then, we specified tags with backtick annotation on the fields that need modification.
//   - Lastly, we added a unique constraint on the Email column to ensure that no two users have the same email addresses.
type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"` // uuid_generate_v4() is a default value on ID column
	// ID     uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
