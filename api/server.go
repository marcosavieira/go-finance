package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/marcosavieira/go-finance/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//*User Routers
	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	router.GET("/user/id/:id", server.getUserById)

	//*Category Routers
	router.POST("/category", server.createCategory)
	router.GET("/category/id/:id", server.getCategory)
	router.GET("/category", server.getCategories)
	router.DELETE("/category/:id", server.deleteCategory)
	router.PUT("/category/:id", server.updateCategory)

	//*Account Routers
	router.POST("/account", server.createAccount)
	router.GET("/account/id/:id", server.getAccount)
	router.GET("/account", server.getAccounts)
	router.GET("/account/graph/:user_id/:type", server.getAccountGraph)
	router.GET("/account/reports/:user_id/:type", server.getAccountReports)
	router.DELETE("/account/:id", server.deleteAccount)
	router.PUT("/account/:id", server.updateAccount)

	//*Login
	router.POST("/login", server.userLogin)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has error response:": err.Error()}
}
