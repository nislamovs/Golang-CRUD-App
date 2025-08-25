package models

import (
	"database/sql"
	"fmt"
	"golang-crud-app/db"
	_ "golang-crud-app/db"
	"strconv"
	"time"
)

type Event struct {
	ID          int64     `db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Location    string    `json:"location" db:"location"`
	DateTime    time.Time `json:"dateTime" db:"dateTime"`
	UserID      int64     `db:"user_id"`
}

//var events []Event = []Event{}

func (e *Event) Save() (string, error) {
	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Could not close stmt.")
		}
	}(stmt)

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	e.ID = id
	//events = append(events, e)

	return strconv.Itoa(int(id)), nil
}

func (e Event) Update() (string, error) {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ?
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Could not close stmt.")
		}
	}(stmt)

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return "", err
	}

	lastId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf(" \n last inserted id : %d", lastId)
	fmt.Printf(" \n rows affected : %d \n", rowsAffected)

	return strconv.Itoa(int(e.ID)), nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic("Could not close rows.")
		}
	}(rows)

	var responseEvents []Event
	for rows.Next() {
		var event Event
		var timestamp string
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &timestamp, &event.UserID)
		if err != nil {
			return nil, err
		}
		event.DateTime, _ = time.Parse(time.RFC3339, timestamp)
		responseEvents = append(responseEvents, event)
	}

	return responseEvents, nil
}

func GetEventById(id int64) (Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return Event{}, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic("Could not close rows.")
		}
	}(stmt)

	var responseEvent Event
	var timestamp string
	err = stmt.QueryRow(id).Scan(&responseEvent.ID, &responseEvent.Name, &responseEvent.Description, &responseEvent.Location, &timestamp, &responseEvent.UserID)
	if err != nil {
		return Event{}, err
	}
	responseEvent.DateTime, _ = time.Parse(time.RFC3339, timestamp)
	return responseEvent, nil
}

func DeleteById(id int64) error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic("Could not close rows.")
		}
	}(stmt)

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (e Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations (event_id, user_id)
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Could not close stmt.")
		}
	}(stmt)

	_, err = stmt.Exec(e.ID, e.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM registrations
	WHERE event_id = ? AND user_id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Could not close stmt.")
		}
	}(stmt)

	_, err = stmt.Exec(e.ID, e.UserID)
	if err != nil {
		return err
	}

	return nil
}
