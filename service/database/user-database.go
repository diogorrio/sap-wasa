package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) SetUserID(u User, s string) (User, error) {
	if len(s) == 0 {
		res, err := db.c.Exec("INSERT INTO users(user_id) VALUES (?)")

		if err != nil {
			var user User
			if err := db.c.QueryRow(`SELECT user_id FROM users WHERE user_id = ?`, u.UserID).Scan(&user.UserID); err != nil {
				if err == sql.ErrNoRows {
					return user, errors.New("404: User not found")
				}
			}
			return user, nil
		}
		autoincr, err := res.LastInsertId()
		u.UserID = fmt.Sprintf("User%d", autoincr)
	} else {
		db.c.Exec("INSERT INTO users(user_id) VALUES s")
	}
	return u, nil
}

func (db *appdbimpl) InitSetUserID(u User) (User, error) {
	res, err := db.c.Exec("INSERT INTO users(user_id) VALUES (?)")

	if err != nil {
		var user User
		if err := db.c.QueryRow(`SELECT user_id FROM users WHERE user_id = ?`, u.UserID).Scan(&user.UserID); err != nil {
			if err == sql.ErrNoRows {
				return user, errors.New("404: User not found")
			}
		}
		return user, nil
	}
	autoincr, err := res.LastInsertId()
	u.UserID = fmt.Sprintf("User%d", autoincr)
	return u, nil
}

func (db *appdbimpl) SetUsername(u User, s string) (User, error) {
	res, err := db.c.Exec("INSERT INTO users(user_name) VALUES s", u.UserName)

	rows_aff, err := res.RowsAffected()
	if rows_aff == 0 || err != nil {
		return u, err
	}

	return u, nil
}

func (db *appdbimpl) GetUserStream(u User) ([]Photo, error) {
	var stream []Photo

	// Choose photos from following and unbanned users
	r, err :=
		db.c.Query(`SELECT * FROM photos WHERE user_id IN (SELECT followed_id FROM follows WHERE user_id = ?
                                                                         AND followed_id NOT IN (SELECT user_id FROM bans WHERE banned_id = ?))`, u.UserID)
	if err != nil {
		return stream, errors.New("404: Photo not found")
	}

	if r.Err() != nil {
		return nil, err
	}

	return stream, nil
}

func (db *appdbimpl) GetFollowers(u User) (int, error) {
	var followers_nr int

	if err := db.c.QueryRow(`SELECT COUNT(*) FROM follows WHERE followed_id = ?`, u.UserID).Scan(&followers_nr); err != nil {
		if err == sql.ErrNoRows {
			return followers_nr, errors.New("404: User not found")
		}
	}
	return followers_nr, nil
}

func (db *appdbimpl) GetFollowing(u User) (int, error) {
	var following_nr int

	if err := db.c.QueryRow(`SELECT COUNT(*) FROM follows WHERE user_id = ?`, u.UserID).Scan(&following_nr); err != nil {
		if err == sql.ErrNoRows {
			return following_nr, errors.New("404: User not found")
		}
	}
	return following_nr, nil
}
