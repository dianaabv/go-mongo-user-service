package main

import (
	"context"
	// "database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"gokit-example/account"
	"go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO os.Getenv
const (
	dbsource = "mongodb://localhost:27017"
	dbsourcewithcred = "mongodb://admin:abc123@localhost:27017"
	hosts      = "localhost:27017"
	database   = "buddyApp"
	username   = "admin"
	password   = "abc123"
	collection = "goUsers"
)
func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	clientOptions := options.Client().ApplyURI(dbsource)

    // Connect to MongoDB 
    db, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
	    level.Error(logger).Log("exit", err)
	}

	err = db.Ping(context.TODO(), nil)
	if err != nil {
		level.Error(logger).Log("exit", err)
	}
	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	{	
		repository := account.NewRepo(db, logger)

		srv = account.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := account.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
