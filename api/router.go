package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Route struct {
	Name    string
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

type Routes []Route

func NewRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(gin.Recovery())
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Path, route.Handler)
		case http.MethodPost:
			router.POST(route.Path, route.Handler)
		case http.MethodPut:
			router.PUT(route.Path, route.Handler)
		}
	}
	return router
}

var routes = Routes{
	{
		"GetPerson",
		"/api/person/:id",
		http.MethodGet,
		GetPerson,
	},
}
