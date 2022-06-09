package httproutes

import (
	"github.com/gorilla/mux"
	"github.com/iitheo/theofetchrewards/pkg/api/controllers/v1/pointscontroller"
	"github.com/iitheo/theofetchrewards/pkg/api/middleware"
	"github.com/urfave/negroni"
)

//Router for all routes
func Router() *negroni.Negroni {
	route := mux.NewRouter()

	n := negroni.Classic()
	n.Use(middleware.Cors())
	n.UseHandler(route)

	//***************************************
	// Points  ROUTES
	//***************************************

	pointRoute := route.PathPrefix("/v1/points").Subrouter()
	pointRoute.HandleFunc("/getall", pointscontroller.GetAllPoints).Methods("GET")
	pointRoute.HandleFunc("/add", pointscontroller.AddPoints).Methods("POST")
	pointRoute.HandleFunc("/spend", pointscontroller.SpendPoints).Methods("POST")

	return n
}
