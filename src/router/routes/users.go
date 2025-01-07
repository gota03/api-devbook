package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		Uri:                   "/users",
		Method:                http.MethodPost,
		Func:                  controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users",
		Method:                http.MethodGet,
		Func:                  controllers.GetUsers,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodGet,
		Func:                  controllers.GetOneUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodPut,
		Func:                  controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Func:                  controllers.DeleteUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}/follow",
		Method:                http.MethodPost,
		Func:                  controllers.FollowUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}/unfollow",
		Method:                http.MethodPost,
		Func:                  controllers.UnfollowUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}/followers",
		Method:                http.MethodGet,
		Func:                  controllers.SearchFollowers,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{userId}/following",
		Method:                http.MethodGet,
		Func:                  controllers.SearchFollowing,
		RequireAuthentication: true,
	},
}
