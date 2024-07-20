package routes

import (
	"net/http"

	"WbTest/internal/middleware"
	"WbTest/internal/mock/delivery"
	"github.com/gorilla/mux"
)

func GetRouter(mockHandlers *delivery.MockDelivery, mw *middleware.Middleware) *mux.Router {
	router := mux.NewRouter()
	assignRoutes(router, mockHandlers)
	assignMiddleware(router, mw)
	return router
}

func assignRoutes(router *mux.Router, mockHandlers *delivery.MockDelivery) {
	router.HandleFunc("/mocks", mockHandlers.CreateMock).Methods(http.MethodPost)
	router.HandleFunc("/mocks", mockHandlers.GetAllMocks).Methods(http.MethodGet)
	router.HandleFunc("/mocks/{method}/{url:.*}", mockHandlers.GetMock).Methods(http.MethodGet)
}

func assignMiddleware(router *mux.Router, mw *middleware.Middleware) {
	router.Use(mw.AccessLog)
	//router.Use(mw.Auth)
}

func serveHTMLFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../templates/index.html")
}
