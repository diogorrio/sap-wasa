/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	PhotoNr     int    `json:"photo_nr"`
	FollowersNr int    `json:"followers_nr"`
	FollowingNr int    `json:"following_nr"`
}

type Stream struct {
	UserID      string  `json:"user_id"`
	StreamID    string  `json:"stream_id"`
	PhotoStream []Photo `json:"stream"`
}

type Photo struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	PhotoID   string `json:"photo_id"`
	PhotoData string `json:"photo_data"`
	PhotoTime string `json:"photo_time"`
	LikeNr    int    `json:"like_nr"`
	Liked     bool   `json:"like"`
	CommentNr int    `json:"comment_nr"`
}

type FollowAction struct {
	UserID     string `json:"user_id"`
	FollowedID string `json:"followed_id"`
}

type BanAction struct {
	UserID   string `json:"user_id"`
	BannedID string `json:"banned_id"`
}

type LikeAction struct {
	UserID  string `json:"user_id"`
	LikedID string `json:"liked_id"`
	PhotoID string `json:"photo_id"`
	LikeID  string `json:"like_id"`
}

type CommentAction struct {
	UserID      string    `json:"user_id"`
	CommentedID string    `json:"commented_id"`
	PhotoID     string    `json:"photo_id"`
	CommentArr  []Comment `json:"comment_array"`
}

type Comment struct {
	CommentBody string `json:"content"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetName(name string) error
	// TODO: ALTER
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	// TODO: ALTER

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
