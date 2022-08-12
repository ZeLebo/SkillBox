package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/create":
			log.Info("Created new user")
		case "/make_friends":
			log.Info("Made two users friends")
		case "user":
			log.Info("Deleted user")
		case "/{id:[0-9]+}":
			log.Info("Changed the age of the user")
		default:
			log.Info("Requested path: ", r.URL.Path, "\n",
				"Method: ", r.Method, "\n",
				"Body: ", r.Body, "\n")
		}
		next(w, r)
	}
}
