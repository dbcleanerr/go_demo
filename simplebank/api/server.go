package api

import (
	"github.com/gin-gonic/gin"
	db "go_demo/simplebank/db/sqlc"
)

// Server HTTP requests for our bank service.
type Server struct {
	// db
	store *db.Store
	// gin
	route *gin.Engine
}

// NewServer create a new http server and setup routing.
// 传入 store :db
// 传出 Server : Server struct
func NewServer(store *db.Store) *Server {
	route := gin.Default()

	// 用指针
	server := &Server{
		store: store,
		route: route,
	}

	// route
	route.POST("/accounts", server.createAccount)
	route.GET("/accounts/:id", server.getAccount)
	route.GET("/accounts", server.listAccount)
	route.POST("/transfer", server.transfer)
	route.POST("/users", server.createUser)

	return server
}

func (s Server) Start(address string) error {
	return s.route.Run(address)
}

// HTTP请求错误处理
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
