package main

import (
	"fmt"
	"github.com/iitheo/theofetchrewards/pkg/api/routes/v1/httproutes"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	env  string
	port string
)

func init() {

	env = strings.TrimSpace(strings.ToLower(os.Getenv("SERVER")))

	if env == "" {
		env = "development"
		os.Setenv("PORT", "8081")
	}
	log.Printf("%s environment started", env)

}

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.SetOutput(logger.Writer())
	port = os.Getenv("PORT")

	log.Printf("server running at %s", port)
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           httproutes.Router(),
		TLSConfig:         nil,
		ReadTimeout:       90 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      90 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server - %s\n", err.Error())
	}
}
