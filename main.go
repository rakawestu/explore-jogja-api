package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rakawestu/explore-jogja-api/middlewares"
	"github.com/rakawestu/explore-jogja-api/orm"
	route "github.com/rakawestu/explore-jogja-api/routes"
)

const (
	// MongoDBUrlKey is flag key for Mongo DB URL
	MongoDBUrlKey = "MongoDBUrl"
	// MongoDBNameKey is flag key for Mongo DB name
	MongoDBNameKey = "MongoDBName"
	// MongoDBUrlDefaultValue is Mongo DB default URL
	MongoDBUrlDefaultValue = "localhost"
	// MongoDBNameDefaultValue is Mongo DB default name
	MongoDBNameDefaultValue = "explore-jogja-api"
	// AccessKeyKey is a header key for AccessKey
	AccessKeyKey = "Explore-Jogja-Auth"
	// PortKey is key for port value
	PortKey = "Port"
	// DefaultPort is default value for port variable
	DefaultPort = "5000"
)

// AccessKey access key
var AccessKey string

// MongoDBUrl url of mongo db
var MongoDBUrl string

// MongoDBName database name of mongo db
var MongoDBName string
var port string

func main() {
	router := gin.New()
	port := os.Getenv("PORT")
	flag.StringVar(&MongoDBUrl, MongoDBUrlKey, MongoDBUrlDefaultValue, "Mongo DB URL.")
	flag.StringVar(&MongoDBName, MongoDBNameKey, MongoDBNameDefaultValue, "Mongo DB database name")
	flag.StringVar(&AccessKey, AccessKeyKey, "", "Custom authentication token")
	flag.Parse()

	orm.MongoDBName = MongoDBName
	orm.MongoDBUrl = MongoDBUrl
	middlewares.AccessKey = AccessKey

	router.Use(middlewares.CheckHeaders())
	routes := router.Group("/api")
	{
		route.PlaceRoutes(routes)
	}

	router.Run(":" + port)
}
