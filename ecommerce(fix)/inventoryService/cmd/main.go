// package cmd
package main

import (
	"log"
	"net"

	grpcDelivery "inventoryService/internal/delivery/grpc"
	pb "inventoryService/internal/delivery/grpc/pb"
	"inventoryService/internal/message"
	"inventoryService/internal/repository"
	"inventoryService/internal/usecase"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=0000 dbname=ecommerce sslmode=disable")
	if err != nil {
		log.Fatalln("Failed to connect DB:", err)
	}

	// Инициализация RabbitMQ клиента
	rabbitClient, err := message.NewRabbitMQClient("amqp://guest:guest@localhost:5672/", "order_events")
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer rabbitClient.Close()

	// Инициализация сервисов
	productRepo := repository.NewProductRepo(db)
	productUC := usecase.NewProductUsecase(productRepo)

	// Запуск потребителя сообщений
	consumer := message.NewMessageConsumer(productUC, rabbitClient)
	if err := consumer.Start(); err != nil {
		log.Fatalf("Failed to start message consumer: %v", err)
	}
	log.Println("RabbitMQ consumer started successfully")

	// Запуск gRPC сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpcDelivery.NewInventoryHandler(productUC)
	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, server)

	log.Println("InventoryService gRPC started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
