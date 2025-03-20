package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string    `gorm:"type:char(36);primary_key" json:"id"`
	Email        string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"` // "-" excludes from JSON
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	Vaults       []Vault   `gorm:"foreignKey:UserID" json:"vaults,omitempty"`
	Devices      []Device  `gorm:"foreignKey:UserID" json:"devices,omitempty"`
}

type Vault struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id"`
	UserID    string    `gorm:"type:char(36);not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	Secrets   []Secret  `gorm:"foreignKey:VaultID" json:"secrets,omitempty"`
}

type Secret struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id"`
	VaultID   string    `gorm:"type:char(36);not null" json:"vault_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Data      string    `gorm:"type:text;not null" json:"data"` // Encrypted JSON data
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null" json:"updated_at"`
}

type Device struct {
	ID       string    `gorm:"type:char(36);primary_key" json:"id"`
	UserID   string    `gorm:"type:char(36);not null" json:"user_id"`
	DeviceID string    `gorm:"type:varchar(255);unique;not null" json:"device_id"`
	LastSync time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"last_sync"`
}

type AuditLog struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id"`
	UserID    string    `gorm:"type:char(36);not null" json:"user_id"`
	Action    string    `gorm:"type:varchar(255);not null" json:"action"`
	Timestamp time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"timestamp"`
	Metadata  string    `gorm:"type:json" json:"metadata,omitempty"`
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

func (v *Vault) BeforeCreate(_ *gorm.DB) error {
	if v.ID == "" {
		v.ID = uuid.New().String()
	}
	return nil
}

func (s *Secret) BeforeCreate(_ *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

func (d *Device) BeforeCreate(_ *gorm.DB) error {
	if d.ID == "" {
		d.ID = uuid.New().String()
	}
	return nil
}

func (a *AuditLog) BeforeCreate(_ *gorm.DB) error {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return nil
}
