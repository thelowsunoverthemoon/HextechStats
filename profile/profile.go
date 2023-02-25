package profile

/* Database functions for each profile (SQLite) */

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Profile struct {
	Name   string    `json:"name"`
	Server string `json:"server"`
	Date   string `json:"date"`
}

var DB *sql.DB
// Path to your database here
var DATA_BASE = ""

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", DATA_BASE)
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func GetProfile(name string, server string) (string, error) {

	stmt, err := DB.Prepare("SELECT data from profile WHERE name = ? AND server = ?")

	if err != nil {
		return "", err
	}

	var data string
	sqlErr := stmt.QueryRow(name, server).Scan(&data)

	if sqlErr != nil {
        // special case b/c error is not fatal
		if sqlErr == sql.ErrNoRows {
			return "", nil
		}
		return "", sqlErr
	}
	return data, nil
    
    
}

func GetProfiles() ([]Profile, error) {

	rows, err := DB.Query("SELECT name, server, date from profile")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	profiles := make([]Profile, 0)

    // get all rows by continually scanning
	for rows.Next() {
		profile := Profile{}
        err = rows.Scan(&profile.Name, &profile.Server, &profile.Date)

		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return profiles, err
}


func DeleteProfile(name string, server string) (bool, error) {

	tx, err := DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("DELETE from profile WHERE name = ? AND server = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(name, server)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}



func AddProfile(name string, server string, date string, data string) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO profile (name, server, date, data) VALUES (?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(name, server, date, data)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpdateProfile(name string, server string, date string, data string) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE profile SET date = ?, data = ? WHERE name = ? AND server = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(date, data, name, server)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
