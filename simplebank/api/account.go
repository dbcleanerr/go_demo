package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "go_demo/simplebank/db/sqlc"
	"net/http"
)

type createAccountRequest struct {
	// binding对应下面的 bindJson类似的方法
	// required 必真字段
	// oneof 多选一
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=RMB EUR"`
}

func (s Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	// 取请求的json参数
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 返回
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}

	// 用server里指定的store
	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	// url的方式
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := s.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=10"`
}

func (s Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := s.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
