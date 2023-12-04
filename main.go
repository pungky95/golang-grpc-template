package main

import (
	"fmt"
	"net"
	"time"

	"github.com/pungky95/golang-grpc-template/config"
	"github.com/pungky95/golang-grpc-template/repositories"
	"github.com/pungky95/golang-grpc-template/server"
	"github.com/pungky95/golang-grpc-template/services"

	"github.com/pungky95/golang-grpc-template/lib/interceptors"

	pb "github.com/pungky95/golang-grpc-proto-template/generated"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

var log = logrus.New()
var db *gorm.DB

func init() {
	var err error
	err = config.CheckEnv()
	if err != nil {
		log.Fatal(err)
	}
	db, err = config.DB()
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		log.Fatal("failed connect to database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed connect to database")
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
}

func main() {
	port := config.GetEnv("PORT", "")
	apiKey := config.GetEnv("API_KEY", "")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Printf("[API] Failed to connect grpc server due to: %s", err.Error())

		panic(err)
	}
	serverInterceptor := interceptors.NewAPIKeyServerInterceptor(apiKey)
	grpcServer := grpc.NewServer(
		serverInterceptor.WithAPIKeyServerUnaryInterceptor(),
	)
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(
		userRepository,
	)
	userServer := server.NewUserServer(userService)
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	log.Printf("[USER-API] Starting gRPC server on port: %s", port)
	e := grpcServer.Serve(lis)

	if e != nil {
		panic(err)
	}
}
