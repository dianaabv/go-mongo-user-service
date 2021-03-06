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
	"gokit-example/activity"
	"gokit-example/account/config"
	"go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/joho/godotenv"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        fmt.Print("No .env file found")
    }
}
func main() {
	conf := config.New()
	var httpAddr = flag.String("http", ":" + conf.AppConfig.Defaultport, "http listen address")
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

	clientOptions := options.Client().ApplyURI(conf.Database.DBSource)
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
	// add email index to user on start of application
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
	coll := db.Database(conf.Database.Database).Collection(conf.Database.CollectionUsers)
	if _, err := coll.Indexes().CreateOne(ctx, index); err != nil {
	  fmt.Println("Could not create index:", err)
	}
	// end  of it 
	var srv account.Service
	{	
		repository := account.NewRepo(db, logger)

		srv = account.NewService(repository, logger)
	}

	var actv activity.Service
	{	
		repository := activity.NewRepo(db, logger)

		actv = activity.NewService(repository, logger)
	}


	mux := http.NewServeMux()
	// activity.NewHTTPServer(ctx, activityEndpoints)
	accountEndpoints := account.MakeEndpoints(srv)
	activityEndpoints := activity.MakeEndpoints(actv)
	mux.Handle("/account/v1/", account.NewHTTPServer(ctx, accountEndpoints))
	mux.Handle("/activity/v1/", activity.NewHTTPServer(ctx, activityEndpoints))
	http.Handle("/", accessControl(mux))
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()


	go func() {
		fmt.Println("listening on port", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()

	level.Error(logger).Log("exit", <-errs)
}
func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}