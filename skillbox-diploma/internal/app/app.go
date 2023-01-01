package app

import "diploma/internal/controller/v1/rest"

func Run() {
	rest.ListenAndServeHTTP()
}
