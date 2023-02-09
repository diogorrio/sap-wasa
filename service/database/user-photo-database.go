package database

func (db *appdbimpl) UploadPhoto(p Photo) (Photo, error) {
	_, err := db.c.Exec(`INSERT INTO photos (userId, user_name, photo_id, photo_data, photo_time, like_nr, comment_nr, liked)
			VALUES (?, ?, ?, ?)`, p.UserID, p.UserName, p.PhotoID, p.PhotoData, p.PhotoTime, p.LikeNr, p.CommentNr, p.Liked)

	if err != nil {
		return p, err
	}
	return p, nil
}

func (db *appdbimpl) DeletePhoto(s string) error {
	_, err_photos := db.c.Exec(`DELETE FROM photos WHERE photo_data=?`, s)
	if err_photos != nil {
		return err_photos
	}
	_, err_likes := db.c.Exec(`DELETE FROM likes WHERE photo_data=?`, s)
	if err_likes != nil {
		return err_likes
	}
	_, err_comments := db.c.Exec(`DELETE FROM comments WHERE photo_data=?`, s)
	if err_comments != nil {
		return err_comments
	}

	return nil
}

func (db *appdbimpl) AddLike(l LikeAction) (LikeAction, error) {
	_, err := db.c.Exec(`INSERT INTO likes (user_id, liked_id, photo_id, like_id)
			VALUES (?, ?, ?, ?)`, l.UserID, l.LikedID, l.PhotoID, l.LikeID)

	if err != nil {
		return l, err
	}
	return l, nil
}

func (db *appdbimpl) RemoveLike(l LikeAction) error {
	res, err := db.c.Exec(`DELETE FROM likes WHERE like_id=?`, l.LikeID)

	if err != nil {
		return err
	}

	rows_aff, err := res.RowsAffected()
	if rows_aff == 0 || err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) AddComment(c CommentAction) (CommentAction, error) {
	_, err := db.c.Exec(`INSERT INTO comments (user_id, commented_id, photo_id, comment_arr)
			VALUES (?, ?, ?, ?)`, c.UserID, c.CommentedID, c.PhotoID, c.CommentArr)

	if err != nil {
		return c, err
	}
	return c, nil
}

func (db *appdbimpl) RemoveComment(s string) error {
	res, err := db.c.Exec(`DELETE FROM comments WHERE comment_arr=?`, s)

	if err != nil {
		return err
	}

	rows_aff, err := res.RowsAffected()
	if rows_aff == 0 || err != nil {
		return err
	}
	return nil
}
