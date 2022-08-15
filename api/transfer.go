package api

import (
	"net/http"

	db "github.com/althof3/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	FromAccountID    int64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountID    int64 `json:"to_account_id" binding:"required,min=1"`
	Amount   int64 `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req TransferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	args := db.TransferTxParams{
		FromAccountID:  req.FromAccountID,
        ToAccountID: req.ToAccountID,
        Amount: req.Amount,
	}

	result, err := server.store.TransferTx(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}