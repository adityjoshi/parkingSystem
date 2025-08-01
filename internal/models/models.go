package models

import "time"

type Vehicle struct {
	ID          uint   `gorm:"primaryKey"`
	NumberPlate string `gorm:"unique"`
	VehicleType string
	CreatedAt   time.Time
}

type ParkingSlot struct {
	ID         uint   `gorm:"primaryKey"`
	SlotNumber string `gorm:"unique"`
	SlotType   string
	Status     string
}

type ParkingSession struct {
	ID            uint `gorm:"primaryKey"`
	VehiclePlate  string
	SlotID        uint
	EntryTime     time.Time
	ExitTime      *time.Time
	Status        string
	BillingType   string
	BillingAmount float64
}
