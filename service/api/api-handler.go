package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//// Register routes
	//rt.router.GET("/", rt.getHelloWorld)
	//rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Login Tag Related
	rt.router.POST("/session", rt.login)

	// User Tag Related
	rt.router.GET("/user/:user_id/get_user_profile", rt.getUserProfile)
	rt.router.PUT("/user/:user_id/set_user_id", rt.setUserID)
	rt.router.PUT("/user/:user_id/set_user_name", rt.setUsername)
	rt.router.GET("/user/:user_id/get_user_stream", rt.getUserStream)
	rt.router.GET("/user/:user_id/get_followers", rt.getFollowers)
	rt.router.GET("/user/:user_id/get_following", rt.getFollowing)

	// User-Photo Interaction Related
	rt.router.POST("user/:user_id/photo/:photo_id", rt.uploadPhoto)
	rt.router.DELETE("user/:user_id/photo/:photo_id", rt.deletePhoto)

	rt.router.PUT("/user/:user_id/photo/:photo_id/like_photo/:like_id", rt.addLike)
	rt.router.DELETE("/user/:user_id/photo/:photo_id/like_photo/:like_id", rt.removeLike)

	rt.router.POST("/user/:user_id/photo/:photo_id/comment_photo/:comment_id", rt.addComment)
	rt.router.DELETE("/user/:user_id/photo/:photo_id/comment_photo/:comment_id", rt.removeComment)

	// User-User Interaction Related
	rt.router.PUT("/user/:user_id/follow_user/:follow_id", rt.followUser)
	rt.router.DELETE("/user/:user_id/follow_user/:follow_id", rt.unfollowUser)

	rt.router.PUT("/user/:user_id/ban_user/:ban_id", rt.banUser)
	rt.router.DELETE("/user/:user_id/ban_user/:ban_id", rt.unbanUser)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
