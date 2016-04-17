package main

import (
	"flag"

	"github.com/gin-gonic/gin"
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
	DefaultPort = "8080"
)

// AccessKey access key
var AccessKey string
var mongoDBUrl, mongoDBName, port string

func main() {
	router := gin.Default()

	flag.StringVar(&mongoDBUrl, MongoDBUrlKey, MongoDBUrlDefaultValue, "Mongo DB URL.")
	flag.StringVar(&mongoDBName, MongoDBNameKey, MongoDBNameDefaultValue, "Mongo DB database name")
	flag.StringVar(&AccessKey, AccessKeyKey, "", "Custom authentication token")
	flag.StringVar(&port, PortKey, DefaultPort, "Port")
	flag.Parse()

	router.Use(CheckHeaders())
	routes := router.Group("/api")
	{
		APIRoutes(routes)
	}

	router.Run(":" + port)
}
