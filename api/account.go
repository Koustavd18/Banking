package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/Koustavd18/Banking/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountReq struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

type getAccountReq struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listAccReq struct {
	Page int32 `form:"page" binding:"required"`
	Size int32 `form:"size" binding:"required,min=5,max=20"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	acc, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)

}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	acc, err := server.store.GetAccount(ctx, req.ID)
	fmt.Println(acc)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, acc)

}

func (server *Server) getAllAcoounts(ctx *gin.Context) {
	var req listAccReq

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	limit := req.Size
	offset := (req.Page - 1) * limit

	accs, err := server.store.ListAccount(ctx, db.ListAccountParams{Limit: req.Size, Offset: offset})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accs)

}
