package http

import (
	"io"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

type Storage interface {
	Get(key string) (string, error)
	Set(key string, value string)
	Delete(key string) error
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) Handler {
	return Handler{storage: storage}
}

func (h Handler) GetValue(ctx echo.Context) error {
	key, err := validateKey(ctx.Param("key"))

	if err != nil {
		log.Warn("GetValue: invalid key", "key", key)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	value, err := h.storage.Get(key)

	if err != nil {
		log.Warn("GetValue: value not found", "key", key)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	log.Info("GetValue", "key", key, "value", value)
	return ctx.String(http.StatusOK, value)
}

func (h Handler) SetValue(ctx echo.Context) error {
	key, err := validateKey(ctx.Param("key"))

	if err != nil {
		log.Warn("SetValue: invalid key", "key", key)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	value, err := io.ReadAll(ctx.Request().Body)
	stringValue := string(value)

	if err != nil {
		log.Error("SetValue: error reading body", "key", key, "err", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "what do you mean")
	}

	h.storage.Set(key, stringValue)
	log.Info("SetValue", "key", key, "value", stringValue)

	return ctx.String(http.StatusOK, stringValue)
}

func (h Handler) DeleteValue(ctx echo.Context) error {
	key, err := validateKey(ctx.Param("key"))

	if err != nil {
		log.Warn("DeleteValue: invalid key", "key", key)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.storage.Delete(key); err != nil {
		log.Warn("DeleteValue: value not found", "key", key)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	log.Info("DeleteValue", "key", key)
	return ctx.NoContent(http.StatusOK)
}
