package handlers

import (
	"go-rest-api/internal/repository"
	"net/http"

	"github.com/labstack/echo"
)

type TransactionHandler struct {
	txnDB repository.TransactionRepository
}

// NewTransactionHandler returns a new instance of TransactionHandler
func NewTransactionHandler(txnDB repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{
		txnDB: txnDB,
	}
}

// InsertTransaction Inserts a new Transaction into the database
func (h *TransactionHandler) InsertTransaction(c echo.Context) error {
	var txn repository.Transaction
	if err := c.Bind(&txn); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	if err := c.Validate(&txn); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}
	if err := h.txnDB.Insert(c.Request().Context(), &txn); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to insert transaction")
	} else {
		return c.JSON(http.StatusOK, "Success")
	}
}
