package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/marcosavieira/go-finance/db/sqlc"
	"github.com/marcosavieira/go-finance/db/util"
)

type createAccountRequest struct {
	UserID      int32     `form:"user_id" json:"user_id" binding:"required"`
	CategoryID  int32     `form:"category_id" json:"category_id" binding:"required"`
	Title       string    `form:"title" json:"title" binding:"required"`
	Type        string    `form:"type" json:"type" binding:"required"`
	Description string    `form:"description" json:"description" binding:"required"`
	Value       int32     `form:"value" json:"value" binding:"required"`
	Date        time.Time `form:"date" json:"date" binding:"required"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	errorToken := util.VerifyBaererToken(ctx)
	if errorToken != nil {
		return
	}
	var req createAccountRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	categoryId := req.CategoryID
	accountType := req.Type

	category, err := server.store.GetCategory(ctx, categoryId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	categoryTypeDifferent := category.Type != accountType

	if categoryTypeDifferent {
		ctx.JSON(http.StatusBadRequest, "Account type different of category type")
		return
	}

	arg := db.CreateAccountParams{
		UserID:      req.UserID,
		CategoryID:  categoryId,
		Title:       req.Title,
		Type:        accountType,
		Description: req.Description,
		Value:       req.Value,
		Date:        req.Date,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	errorToken := util.VerifyBaererToken(ctx)
	if errorToken != nil {
		return
	}
	var req getAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type deleteAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	errorToken := util.VerifyBaererToken(ctx)
	if errorToken != nil {
		return
	}
	var req deleteAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type updateAccountRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title" `
	Description string `json:"description" `
	Value       int32  `json:"value"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	errorToken := util.VerifyBaererToken(ctx)
	if errorToken != nil {
		return
	}
	var req updateAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateAccountParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Value:       req.Value,
	}

	account, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountsRequest struct {
	UserID      int32     `json:"user_id" binding:"required"`
	CategoryID  int32     `json:"category_id"`
	Value       int32     `json:"value" `
	Title       string    `json:"title" `
	Type        string    `json:"type" binding:"required"`
	Description string    `json:"description" `
	Date        time.Time `json:"date"`
}

func (server *Server) getAccounts(ctx *gin.Context) {
	errorToken := util.VerifyBaererToken(ctx)
	if errorToken != nil {
		return
	}
	var req getAccountsRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAccountsParams{
		UserID: req.UserID,
		Type:   req.Type,
		CategoryID: sql.NullInt32{
			Int32: req.CategoryID,
			Valid: req.CategoryID > 0,
		},
		Value: sql.NullInt32{
			Int32: req.Value,
			Valid: req.Value > 0,
		},
		Title:       req.Title,
		Description: req.Description,
		Date: sql.NullTime{
			Time:  req.Date,
			Valid: !req.Date.IsZero(),
		},
	}

	accounts, err := server.store.GetAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type getAccountGraphRequest struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountGraph(ctx *gin.Context) {
	errorToken := util.VerifyBaererToken(ctx)
	if errorToken != nil {
		return
	}
	var req getAccountGraphRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAccountsGraphParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	accountGraph, err := server.store.GetAccountsGraph(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accountGraph)
}

type getAccountReportsRequest struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountReports(ctx *gin.Context) {
	errorToken := util.VerifyBaererToken(ctx)
	if errorToken != nil {
		return
	}
	var req getAccountReportsRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAccountsReportsParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	accountReports, err := server.store.GetAccountsReports(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accountReports)
}
