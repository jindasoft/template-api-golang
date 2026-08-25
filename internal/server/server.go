package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"template-api-golang/configs"

	"github.com/jindasoft/jinda-platform/xdb"
	"github.com/jindasoft/jinda-platform/xlogger"
	"github.com/jinzhu/copier"
)

type server struct {
	port   int
	config *configs.Configs
	// redis    xcache.RedisService
	mongo xdb.MongoService
	// postgres xdb.PostgresService
	// rabbit   xmq.RabbitMQService
}

func NewServer() *http.Server {
	conf := configs.GetConfig()
	sect := configs.GetSecret()

	xlogger.Init(
		conf.Server.LogLevel,
		conf.App.Name,
		conf.Server.PrettyPrint,
	)
	ctx := context.Background()

	// Redis
	// var redisConfig xcache.RedisConfig
	// if err := copier.Copy(&redisConfig, sect.Redis); err != nil {
	// 	log.Fatalf("Failed to copy Redis config: %s", err)
	// }
	// redisService, err := xcache.NewRedis(ctx, &redisConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Redis: %s", err)
	// }

	// MongoDB
	var mongoConfig xdb.MongoConfig
	if err := copier.Copy(&mongoConfig, sect.MongoDB); err != nil {
		log.Fatalf("Failed to copy mongo config: %s", err)
	}
	mongoService, err := xdb.NewMongoService(ctx, &mongoConfig)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err)
	}

	// PostgreSQL
	// var postgresConfig xdb.PostgresConfig
	// if err := copier.Copy(&postgresConfig, sect.Postgres); err != nil {
	// 	log.Fatalf("Failed to copy PostgreSQL config: %s", err)
	// }
	// postgresService, err := xdb.NewPostgresService(ctx, &postgresConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to PostgreSQL: %s", err)
	// }

	// PostgreSQL Migrations
	// if err := migrations.RunMigrations(postgresService); err != nil {
	// 	log.Fatalf("Failed to run migrations: %s", err)
	// }

	// RabbitMQ
	// var rabbitConfig xmq.RabbitMQConfig
	// if err := copier.Copy(&rabbitConfig, sect.RabbitMQ); err != nil {
	// 	log.Fatalf("Failed to copy RabbitMQ config: %s", err)
	// }
	// rabbitMQService, err := xmq.NewRabbitMQ(ctx, &rabbitConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	// }

	// RabbitMQ Consumer
	// consumer := NewConsumer(rabbitMQService, postgresService)
	// go consumer.StartConsumers(ctx)

	newServer := &server{
		port:   conf.App.Port,
		config: conf,
		// redis:    redisService,
		mongo: mongoService,
		// postgres: postgresService,
		// rabbit:   rabbitMQService,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  conf.Server.IdleTimeout,
		ReadTimeout:  conf.Server.ReadTimeout,
		WriteTimeout: conf.Server.WriteTimeout,
	}

	xlogger.SysInfof("Port: %d", conf.App.Port)

	return server
}
