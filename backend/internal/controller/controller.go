package controller

import (
	"net/http"
)

type Controller struct {
    production bool
}

func NewController(production bool) *Controller {
    return &Controller {
        production: production,
    }
}

func (c *Controller) Ping(w http.ResponseWriter, r *http.Request) error {
    return WriteJSON(w, http.StatusOK, "Pong")
}
