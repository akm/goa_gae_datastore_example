//go:generate goagen bootstrap -d github.com/akm/goa_gae_datastore_example/design
// +build appengine

package server

import (
	"net/http"

	"github.com/akm/goa_gae_datastore_example/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/akm/goa_gae_datastore_example/controller"
)

func main() {
	// Create service
	service := goa.New("appengine")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "User" controller
	c := controller.NewUserController(service)
	app.MountUserController(service, c)

	// // Start service
	// if err := service.ListenAndServe(":8080"); err != nil {
	// 	service.LogError("startup", "err", err)
	// }

	// Start service
	http.HandleFunc("/", service.Mux.ServeHTTP)
}
