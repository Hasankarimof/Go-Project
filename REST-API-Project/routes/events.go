package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest-api.com/restapi/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID format"})
		return
	}

	fmt.Println("Event ID:", eventId) // Log the event ID

	existingEvent, err := models.GetEventByID(eventId)
	if err != nil {
		fmt.Println("Error fetching event:", err) // Log the error
		if err.Error() == "event not found" {
			context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching event"})
		}
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println("Error parsing JSON:", err) // Log the error
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event data"})
		return
	}

	fmt.Println("Updated Event Data:", updatedEvent) // Log the updated event data

	// Update only fields that are provided
	if updatedEvent.Name != "" {
		existingEvent.Name = updatedEvent.Name
	}
	if updatedEvent.Description != "" {
		existingEvent.Description = updatedEvent.Description
	}
	if updatedEvent.Location != "" {
		existingEvent.Location = updatedEvent.Location
	}
	if !updatedEvent.DateTime.IsZero() {
		existingEvent.DateTime = updatedEvent.DateTime
	}

	fmt.Println("Merged Event:", existingEvent) // Log the merged event data

	err = existingEvent.Update()
	if err != nil {
		fmt.Println("Error updating event:", err) // Log the error
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID format"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}
