package routes

import (
	"golearn/first-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Path parameter 'id' has an invalid value."})
		return
	}

	event, err := model.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := model.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func postEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "the event was successfully created", "event": event})
}

func putEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Path parameter 'id' has an invalid value."})
		return
	}

  _, err = model.GetEvent(id)

  if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
    return
  }

  var updatedEvent model.Event
  err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body is malformed."})
		return
	}

  updatedEvent.ID = id
  err = updatedEvent.Update()

  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body is malformed."})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Path parameter 'id' has an invalid value."})
		return
	}

  event, err := model.GetEvent(id)

  if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
    return
  }

  err = event.Delete()

  if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An unspecified error occurred."})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
