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
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodGet,
		Func:                  controllers.GetOneUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodPut,
		Func:                  controllers.UpdateUser,
		RequireAuthentication: false,
	},
	{
		Uri:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Func:                  controllers.DeleteUser,
		RequireAuthentication: false,
	},
}
