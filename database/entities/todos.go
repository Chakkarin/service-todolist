package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todos struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"" json:"description"`
	DueDate     time.Time `gorm:"not null;index" json:"due_date"`
	Priority    string    `gorm:"not null" json:"priority"` // LOW, MEDIUM, HIGH
	Completed   bool      `gorm:"not null;default:false" json:"completed"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
