package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
)

func MiddlewareContentTypeSet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}

func initDb(mongoUri string) {
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

func initEnforcer(logger *log.Logger, mongoUri string) *casbin.Enforcer {
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

	configurePolicies(enforcer, logger)

	return enforcer
}

func configurePolicies(enforcer *casbin.Enforcer, logger *log.Logger) {

	if hasPolicy := enforcer.HasGroupingPolicy("6425bd9edb1ff9554c5621da", "admin"); !hasPolicy {
		_, err := enforcer.AddGroupingPolicy("6425bd9edb1ff9554c5621da", "admin")
		if err != nil {
			logger.Println("Failed to add admin group policy")
		}
	}

	if hasPolicy := enforcer.HasPolicy("customer", "logout", "GET"); !hasPolicy {
		_, err := enforcer.AddPolicy("customer", "logout", "GET")
		if err != nil {
			logger.Println("Failed to add logout policy for customers [GET]")
		}
	}
	if hasPolicy := enforcer.HasPolicy("admin", "logout", "GET"); !hasPolicy {
		_, err := enforcer.AddPolicy("admin", "logout", "GET")
		if err != nil {
			logger.Println("Failed to add logout policy for admins [GET]")
		}
	}

	if hasPolicy := enforcer.HasPolicy("admin", "flight", "POST"); !hasPolicy {
		_, err := enforcer.AddPolicy("admin", "flight", "POST")
		if err != nil {
			logger.Println("Failed to add flight policy for admins [POST]")
		}
	}
	if hasPolicy := enforcer.HasPolicy("admin", "flight", "DELETE"); !hasPolicy {
		_, err := enforcer.AddPolicy("admin", "flight", "DELETE")
		if err != nil {
			logger.Println("Failed to add flight policy for admins [DELETE]")
		}
	}
	if hasPolicy := enforcer.HasPolicy("customer", "tickets", "POST"); !hasPolicy {
		_, err := enforcer.AddPolicy("customer", "tickets", "POST")
		if err != nil {
			logger.Println("Failed to add tickets policy for customers [POST]")
		}
	}
	if hasPolicy := enforcer.HasPolicy("customer", "tickets", "GET"); !hasPolicy {
		_, err := enforcer.AddPolicy("customer", "tickets", "GET")
		if err != nil {
			logger.Println("Failed to add tickets policy for customers [GET]")
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9000"
	}
	mongoUri := os.Getenv("MONGODB_URI")
	if mongoUri == "" {
		mongoUri = "localhost:9100"
	}

	initDb(mongoUri)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	logger := log.New(os.Stdout, "[api] ", log.LstdFlags)

	router := gin.Default()
	router.Use(MiddlewareContentTypeSet())

	enforcer := initEnforcer(logger, mongoUri)

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
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
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
