package main

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx    context.Context
	client *mongo.Client
)

func MiddlewareContentTypeSet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}

func initDb() {
	mongoUri := os.Getenv("MONGODB_URI")

	if mongoUri == "" {
		mongoUri = "localhost:9100"
	}
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri + "/?connect=direct")

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func initEnforcer(logger *log.Logger) *casbin.Enforcer {
	mongoUri := os.Getenv("MONGODB_URI")
	adapter, err := mongodbadapter.NewAdapter(mongoUri + "/skyriders")
	if err != nil {
		logger.Panicf("Failed to initialize casbin adapter: ", err.Error())
	}

	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		logger.Panicf("Failed to create casbin enforcer: ", err.Error())
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Panicf("Failed to load enforcer policy from the database: ", err.Error())
	}

	configurePolicies(enforcer)

	return enforcer
}

func configurePolicies(enforcer *casbin.Enforcer) {
	_, _ = enforcer.AddPolicy("customer", "logout", "GET")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9000"
	}

	initDb()

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	logger := log.New(os.Stdout, "[api] ", log.LstdFlags)

	router := gin.Default()
	router.Use(MiddlewareContentTypeSet())

	enforcer := initEnforcer(logger)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	logger.Println("Server listening on port", port)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://localhost:9001"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}

	router.Use(cors.New(corsConfig))
	database := client.Database("skyriders")

	routerGroup := router.Group("/api")

	InitializeAllControllers(routerGroup, logger, database, enforcer)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
