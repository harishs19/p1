package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	config "registration/config"
	"registration/logger"
	repo "registration/repo"
	r "registration/route"
	"syscall"
	"time"
)

func init() {
	config.Load()
}
func LoggerInit(loglevel string) *logger.Logger {
	log := logger.New(loglevel)
	return log
}
func main() {

	log := LoggerInit(os.Getenv("LOG_LEVEL"))

	listenAddr := ":" + os.Getenv("HTTP_PORT")
	ctx := context.Background()
	db, err := repo.NewDB(ctx)

	if err != nil {
		log.Warn("error in db connection %s", err)
	}
	defer db.Close()
	log.Info("Successfully connected to the database %s", os.Getenv("DB_CONNECTION"))


	router, err := r.Routes(db, log)

	if err != nil {
		log.Warn("Error initializing router %s", err)
		os.Exit(1)

	}

	
	log.Info("Starting the HTTP server: %s", listenAddr)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("HTTP_PORT"),
		Handler: router,
	}

	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Info("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Info("Received signal...%s", sig)

	duration, err := time.ParseDuration(os.Getenv("SHUTDOWN_TIME"))
	if err != nil {
		log.Fatal("Error in parsing duration", err)

	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown error:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Info("Server exiting")

}
