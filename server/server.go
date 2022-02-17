package server

import (
	"github.com/duckhue01/back-end/store/store"
	"github.com/duckhue01/back-end/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  store.Store
	config *util.Config
}

func NewServer(conf *util.Config, store store.Store) (*Server, error) {
	server := &Server{
		router: &gin.Engine{},
		store:  store,
		config: conf,
	}

	server.setupRouter()

	return server, nil

}

func (server *Server) setupRouter() {
	router := gin.Default()

	// router.POST("/users", server.createUser)
	// router.POST("/users/login", server.loginUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	// router.POST("/transfers", server.createTransfer)

	server.router = router

}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
