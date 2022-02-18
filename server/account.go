package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createAccount(ctx *gin.Context) {
	ctx.String(http.StatusOK, "duckhue01")
}

func (server *Server) listAccounts(ctx *gin.Context) {
	ctx.String(http.StatusOK, "duckhue01")

}

func (server *Server) getAccount(ctx *gin.Context) {
	ctx.String(http.StatusOK, "duckhue01")

}
