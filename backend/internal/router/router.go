package router

import (
	"net/http"

	"github.com/Anand-S23/mailbox/internal/controller"
	"github.com/gorilla/handlers"
)

func NewRouter(c *controller.Controller) *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc("GET /ping", Fn(c.Ping))

    return router
}

func NewCorsRouter(router *http.ServeMux, allowedOrigin string) http.Handler {
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000", allowedOrigin}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

    return corsHandler(router)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func Fn(fn apiFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := fn(w, r)
        if err != nil {
            errMsg := controller.ErrorMessage(err.Error())
            controller.WriteJSON(w, http.StatusInternalServerError, errMsg)
        }
    }
}

