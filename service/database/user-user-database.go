package database

func (db *appdbimpl) followUser(f FollowAction) (FollowAction, error) {
	_, err := db.c.Exec(`INSERT INTO follows (user_id, followed_id) VALUES (?, ?)`, f.UserID, f.FollowedID)

	if err != nil {
		return f, err
	}

	return f, nil
}

func (db *appdbimpl) unfollowUser(f FollowAction) error {
	res, err := db.c.Exec(`DELETE FROM follows WHERE user_id=? AND followed_d=?`, f.UserID, f.FollowedID)

	if err != nil {
		return err
	}

	rows_aff, err := res.RowsAffected()
	if rows_aff == 0 || err != nil {
		return err
	}
	return err
}

func (db *appdbimpl) banUser(b BanAction) (BanAction, error) {
	_, err := db.c.Exec(`INSERT INTO bans (user_id, banned_id) VALUES (?, ?)`, b.UserID, b.BannedID)

	if err != nil {
		return b, err
	}

	return b, nil
}

func (db *appdbimpl) unbanUser(b BanAction) error {
	res, err := db.c.Exec(`DELETE FROM follows WHERE user_id=? AND banned_id=?`, b.UserID, b.BannedID)

	if err != nil {
		return err
	}

	rows_aff, err := res.RowsAffected()
	if rows_aff == 0 || err != nil {
		return err
	}
	return err
}
