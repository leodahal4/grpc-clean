package main

import (
	"fmt"
	"log"
	"net"

	interfaces "github.com/leodahal4/grpc-clean/pkg/v1"
	repo "github.com/leodahal4/grpc-clean/pkg/v1/repository"
	usecase "github.com/leodahal4/grpc-clean/pkg/v1/usecase"
	"gorm.io/gorm"

	dbConfig "github.com/leodahal4/grpc-clean/internal/db"
	"github.com/leodahal4/grpc-clean/internal/models"
	handler "github.com/leodahal4/grpc-clean/pkg/v1/handler/grpc"
	"google.golang.org/grpc"
)

func main() {
  // connect to the db
  db := dbConfig.DbConn()
  migrations(db)

  // add a listener address
  lis, err := net.Listen("tcp", ":5001")
  if err != nil {
    log.Fatalf("ERROR STARTING THE SERVER : %v", err)
  }

  // start the grpc server
  grpcServer := grpc.NewServer()

  userUseCase := initUserServer(db)
  handler.NewServer(grpcServer, userUseCase)

  // start serving to the address
  log.Fatal(grpcServer.Serve(lis))
}

func initUserServer(db *gorm.DB) interfaces.UseCaseInterface{
  userRepo := repo.New(db)
  return usecase.New(userRepo)
}

func migrations(db *gorm.DB) {
  err := db.AutoMigrate(&models.User{})
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Migrated")
  }
}
