package models

import (
	"fmt"
	"time"

	"rest-api.com/restapi/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// later: add it to a database
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	e.ID = id
	events = append(events, e)
}

func GetAllEvents() ([]Event, error) {

	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

// Update updates an existing event in the database
func (event *Event) Update() error {
	query := `
    UPDATE events
    SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
    WHERE id = ?`

	fmt.Println("SQL Query:", query)                 // Log the query
	fmt.Println("Updating event with ID:", event.ID) // Log the event ID

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing statement:", err) // Log the error
		return fmt.Errorf("error preparing update statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	if err != nil {
		fmt.Println("Error executing statement:", err) // Log the error
		return fmt.Errorf("error executing update statement: %v", err)
	}

	fmt.Println("Event updated successfully in the database.") // Log successful update
	return nil
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err

}
