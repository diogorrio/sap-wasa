package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

// ** Main schema instance structuring **

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

// ** Main schema methods **

func (u *User) fromDatabase(user database.User) {
	u.UserID = user.UserID
	u.UserName = user.UserName
	u.PhotoNr = user.PhotoNr
	u.FollowingNr = user.FollowingNr
	u.FollowersNr = user.FollowersNr
}

func (u *User) toDatabase() database.User {
	return database.User{
		UserID:      u.UserID,
		UserName:    u.UserName,
		PhotoNr:     u.PhotoNr,
		FollowingNr: u.FollowingNr,
		FollowersNr: u.FollowersNr,
	}
}

// Due to the non-convertibility of array instances, I had to drop this method
//func (s *Stream) streamFromDatabase(stream database.Stream) {
//	s.UserID = stream.UserID
//	s.StreamID = stream.StreamID
//	s.PhotoStream = stream.PhotoStream
//}
//
//func (s *Stream) streamToDatabase() database.Stream {
//	return database.Stream{
//		UserID:      s.UserID,
//		StreamID:    s.StreamID,
//		PhotoStream: s.PhotoStream,
//	}
//
//}

func (p *Photo) photoFromDatabase(photo database.Photo) {
	p.UserID = photo.UserID
	p.UserName = photo.UserName
	p.PhotoID = photo.PhotoID
	p.PhotoData = photo.PhotoData
	p.PhotoTime = photo.PhotoTime
	p.LikeNr = photo.LikeNr
	p.Liked = photo.Liked
	p.CommentNr = photo.CommentNr
}

func (p *Photo) photoToDatabase() database.Photo {
	return database.Photo{
		UserID:    p.UserID,
		UserName:  p.UserName,
		PhotoID:   p.PhotoID,
		PhotoData: p.PhotoData,
		PhotoTime: p.PhotoTime,
		LikeNr:    p.LikeNr,
		Liked:     p.Liked,
		CommentNr: p.CommentNr,
	}
}

func (f *FollowAction) followActionFromDatabase(followAction database.FollowAction) {
	f.UserID = followAction.UserID
	f.FollowedID = followAction.FollowedID
}

func (f *FollowAction) followActionToDatabase() database.FollowAction {
	return database.FollowAction{
		UserID:     f.UserID,
		FollowedID: f.FollowedID,
	}
}

func (b *BanAction) banActionFromDatabase(banAction database.BanAction) {
	b.UserID = banAction.UserID
	b.BannedID = banAction.BannedID
}

func (b *BanAction) banActionToDatabase() database.BanAction {
	return database.BanAction{
		UserID:   b.UserID,
		BannedID: b.BannedID,
	}
}

func (l *LikeAction) likeActionFromDatabase(likeAction database.LikeAction) {
	l.UserID = likeAction.UserID
	l.LikedID = likeAction.LikedID
	l.PhotoID = likeAction.PhotoID
	l.LikeID = likeAction.LikeID
}

func (l *LikeAction) likeActionToDatabase() database.LikeAction {
	return database.LikeAction{
		UserID:  l.UserID,
		LikedID: l.LikedID,
		PhotoID: l.PhotoID,
		LikeID:  l.LikeID,
	}
}

// Same reasoning as per the stream.
//func (c *CommentAction) commentActionFromDatabase(commentAction database.CommentAction) {
//	c.UserID = commentAction.UserID
//	c.CommentedID = commentAction.CommentedID
//	c.PhotoID = commentAction.PhotoID
//	c.CommentArr = commentAction.CommentArr
//}
//
//func (c *CommentAction) commentActionToDatabase() database.CommentAction {
//	return database.CommentAction{
//		UserID:      c.UserID,
//		CommentedID: c.CommentedID,
//		PhotoID:     c.PhotoID,
//		CommentArr:  c.CommentArr,
//	}
//}

func (cb *Comment) commentFromDatabase(commentBodyAction database.Comment) {
	cb.CommentBody = commentBodyAction.CommentBody
}

func (cb *Comment) commentToDatabase() database.Comment {
	return database.Comment{
		CommentBody: cb.CommentBody,
	}
}
