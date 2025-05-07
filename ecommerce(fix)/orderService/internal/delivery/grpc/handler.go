package grpc

import (
	"context"
	"log"
	pb "orderService/internal/delivery/grpc/pb"
	"orderService/internal/domain"
)

type OrderHandler struct {
	orderUC domain.OrderUsecase
	pb.UnimplementedOrderServiceServer
}

func NewOrderHandler(uc domain.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUC: uc}
}

func (h *OrderHandler) CreateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	log.Printf("[gRPC] Received CreateOrder request for user %d", req.UserId)

	// Преобразование из gRPC модели в доменную
	domainOrder := &domain.Order{
		UserID: int(req.UserId),
		Status: req.Status,
	}

	for _, item := range req.Items {
		domainOrder.Items = append(domainOrder.Items, domain.OrderItem{
			ProductID: int(item.ProductId),
			Quantity:  int(item.Quantity),
		})
	}

	// Создание заказа (также вызовет публикацию сообщения)
	err := h.orderUC.Create(domainOrder)
	if err != nil {
		log.Printf("[gRPC] Error creating order: %v", err)
		return nil, err
	}

	// Преобразование обратно в gRPC модель для ответа
	resp := &pb.Order{
		Id:     int32(domainOrder.ID),
		UserId: int32(domainOrder.UserID),
		Status: domainOrder.Status,
	}

	for _, item := range domainOrder.Items {
		resp.Items = append(resp.Items, &pb.OrderItem{
			Id:        int32(item.ID),
			OrderId:   int32(item.OrderID),
			ProductId: int32(item.ProductID),
			Quantity:  int32(item.Quantity),
		})
	}

	log.Printf("[gRPC] Successfully created order %d", domainOrder.ID)
	return resp, nil
}

func (h *OrderHandler) GetOrder(ctx context.Context, req *pb.OrderID) (*pb.Order, error) {
	order, err := h.orderUC.GetByID(int(req.Id))
	if err != nil {
		return nil, err
	}

	// Преобразование модели домена в gRPC ответ
	resp := &pb.Order{
		Id:     int32(order.ID),
		UserId: int32(order.UserID),
		Status: order.Status,
	}

	for _, item := range order.Items {
		resp.Items = append(resp.Items, &pb.OrderItem{
			Id:        int32(item.ID),
			OrderId:   int32(item.OrderID),
			ProductId: int32(item.ProductID),
			Quantity:  int32(item.Quantity),
		})
	}

	return resp, nil
}

func (h *OrderHandler) UpdateOrderStatus(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	err := h.orderUC.UpdateStatus(int(req.Id), req.Status)
	if err != nil {
		return nil, err
	}

	order, err := h.orderUC.GetByID(int(req.Id))
	if err != nil {
		return nil, err
	}

	// Преобразование модели домена в gRPC ответ
	resp := &pb.Order{
		Id:     int32(order.ID),
		UserId: int32(order.UserID),
		Status: order.Status,
	}

	for _, item := range order.Items {
		resp.Items = append(resp.Items, &pb.OrderItem{
			Id:        int32(item.ID),
			OrderId:   int32(item.OrderID),
			ProductId: int32(item.ProductID),
			Quantity:  int32(item.Quantity),
		})
	}

	return resp, nil
}

func (h *OrderHandler) ListOrdersByUser(ctx context.Context, req *pb.ListOrdersRequest) (*pb.OrderList, error) {
	orders, err := h.orderUC.ListByUser(int(req.UserId))
	if err != nil {
		return nil, err
	}

	resp := &pb.OrderList{}
	for _, order := range orders {
		pbOrder := &pb.Order{
			Id:     int32(order.ID),
			UserId: int32(order.UserID),
			Status: order.Status,
		}

		for _, item := range order.Items {
			pbOrder.Items = append(pbOrder.Items, &pb.OrderItem{
				Id:        int32(item.ID),
				OrderId:   int32(item.OrderID),
				ProductId: int32(item.ProductID),
				Quantity:  int32(item.Quantity),
			})
		}

		resp.Orders = append(resp.Orders, pbOrder)
	}

	return resp, nil
}
