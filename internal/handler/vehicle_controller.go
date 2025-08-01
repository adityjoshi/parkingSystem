package controllers

import (
	"net/http"
	"time"

	"github.com/adityjoshi/parkingSystem.git/internal/db"
	"github.com/adityjoshi/parkingSystem.git/internal/models"

	"github.com/gin-gonic/gin"
)

type RegisterVehicleRequest struct {
	NumberPlate string `json:"number_plate" binding:"required"`
	VehicleType string `json:"vehicle_type" binding:"required"` // e.g., "car", "bike"
	BillingType string `json:"billing_type" binding:"required"` // e.g., "hourly", "daypass"
}

func RegisterVehicle(c *gin.Context) {
	var req RegisterVehicleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := db.DB

	var vehicle models.Vehicle
	if err := db.Where("number_plate = ?", req.NumberPlate).First(&vehicle).Error; err != nil {

		vehicle = models.Vehicle{
			NumberPlate: req.NumberPlate,
			VehicleType: req.VehicleType,
			CreatedAt:   time.Now(),
		}
		if err := db.Create(&vehicle).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register vehicle"})
			return
		}
	}

	var slot models.ParkingSlot
	if err := db.Where("slot_type = ? AND status = ?", req.VehicleType, "available").First(&slot).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No available slot found for this vehicle type"})
		return
	}

	slot.Status = "occupied"
	db.Save(&slot)

	session := models.ParkingSession{
		VehiclePlate:  vehicle.NumberPlate,
		SlotID:        slot.ID,
		EntryTime:     time.Now(),
		Status:        "active",
		BillingType:   req.BillingType,
		BillingAmount: 0.0,
	}

	if err := db.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start parking session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Vehicle registered and slot assigned",
		"slot_number":   slot.SlotNumber,
		"vehicle_plate": vehicle.NumberPlate,
		"session_id":    session.ID,
		"entry_time":    session.EntryTime,
	})
}
