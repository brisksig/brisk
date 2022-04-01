package app01

import (
	"brisk"
	"net/http"
)

func Router() *brisk.Router {
	AppRouter := brisk.NewRouter()
	AppRouter.Add("/api/", http.MethodPost, AppGet)
	return AppRouter
}
