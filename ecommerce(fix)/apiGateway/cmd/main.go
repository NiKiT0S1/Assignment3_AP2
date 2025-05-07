// package cmd
package main

import (
	"apiGateway/internal/delivery/handlers"
	"apiGateway/internal/grpc"
	"apiGateway/internal/middleware"
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())

	// gRPC клиент UserService
	userClient, err := grpcDelivery.NewUserClient("localhost:50053")
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}

	// Маршруты для UserService (регистрация и авторизация)
	handlers.RegisterRoutes(r, userClient)

	// Middleware: авторизация через UserService
	// Применяем middleware после регистрации публичных маршрутов
	r.Use(middleware.Auth(userClient))

	// Настраиваем прокси для сервисов
	setupServiceProxies(r)

	log.Println("API Gateway starting on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}
}

// setupServiceProxies настраивает проксирование запросов к микросервисам
func setupServiceProxies(r *gin.Engine) {
	// Inventory Service Proxy
	inventoryURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatalf("Invalid inventory service URL: %v", err)
	}
	inventoryProxy := httputil.NewSingleHostReverseProxy(inventoryURL)

	// Order Service Proxy
	orderURL, err := url.Parse("http://localhost:8082")
	if err != nil {
		log.Fatalf("Invalid order service URL: %v", err)
	}
	orderProxy := httputil.NewSingleHostReverseProxy(orderURL)

	// Настраиваем маршруты для проксирования
	r.Any("/products", gin.WrapH(inventoryProxy))
	r.Any("/products/*proxyPath", gin.WrapH(inventoryProxy))
	r.Any("/orders", gin.WrapH(orderProxy))
	r.Any("/orders/*proxyPath", gin.WrapH(orderProxy))
}
