package handlers

import (
	"context"
	"order-service/databases"
	"order-service/utils"
	"time"

	orderv1 "github.com/ol1mov-dev/protos/pkg/order/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderServer struct {
	orderv1.UnimplementedOrderV1ServiceServer
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (*orderv1.CreateOrderResponse, error) {
	DB, err := databases.ConnectPostgreSQL()

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to connect to database")
	}

	defer DB.Close()

	query := `INSERT INTO orders (public_order_number, user_id, product_id, quantity, total_amount, warehouse_id)
				VALUES ($1, $2, $3, $4, $5, $6)`

	publicOrderNumber := utils.GeneratePublicOrderNumber()

	_, err = DB.Exec(
		query,
		publicOrderNumber,
		req.UserId,
		req.ProductId,
		req.Quantity,
		req.TotalAmount,
		req.WarehouseId,
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create order")
	}

	return &orderv1.CreateOrderResponse{PublicOrderNumber: publicOrderNumber}, nil
}

func (s *OrderServer) GetAllOrdersByUserId(ctx context.Context, req *orderv1.GetOrdersByUserIdRequest) (*orderv1.GetOrdersByUserIdResponse, error) {
	DB, err := databases.ConnectPostgreSQL()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to connect to database")
	}

	defer DB.Close()

	query := `SELECT id, public_order_number, user_id, product_id, quantity, status, total_amount, warehouse_id, created_at, updated_at FROM orders WHERE user_id=$1`
	rows, err := DB.Query(query, req.UserId)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get orders")
	}

	var orders []*orderv1.Order
	var createdAt time.Time
	var updatedAt time.Time

	defer rows.Close()
	for rows.Next() {
		order := &orderv1.Order{}

		err = rows.Scan(
			&order.Id,
			&order.PublicOrderNumber,
			&order.UserId,
			&order.ProductId,
			&order.Quantity,
			&order.Status,
			&order.TotalAmount,
			&order.WarehouseId,
			&createdAt,
			&updatedAt,
		)
		order.CreatedAt = timestamppb.New(createdAt)
		order.UpdatedAt = timestamppb.New(updatedAt)

		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, status.Error(codes.Internal, "failed to iterate orders")
	}

	return &orderv1.GetOrdersByUserIdResponse{Orders: orders}, nil
}

func (s *OrderServer) GetOrderByPublicOrderNumber(ctx context.Context, req *orderv1.GetOrderByPublicOrderNumberRequest) (*orderv1.GetOrderByPublicOrderNumberResponse, error) {
	DB, err := databases.ConnectPostgreSQL()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to connect to database")
	}

	defer DB.Close()

	query := `SELECT id, public_order_number, user_id, product_id, quantity, status, total_amount, warehouse_id, created_at, updated_at FROM orders WHERE public_order_number=$1`
	rows, err := DB.Query(query, req.PublicOrderNumber)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get orders")
	}

	var orders []*orderv1.Order
	var createdAt time.Time
	var updatedAt time.Time

	defer rows.Close()
	for rows.Next() {
		order := &orderv1.Order{}

		err = rows.Scan(
			&order.Id,
			&order.PublicOrderNumber,
			&order.UserId,
			&order.ProductId,
			&order.Quantity,
			&order.Status,
			&order.TotalAmount,
			&order.WarehouseId,
			&createdAt,
			&updatedAt,
		)
		order.CreatedAt = timestamppb.New(createdAt)
		order.UpdatedAt = timestamppb.New(updatedAt)

		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, status.Error(codes.Internal, "failed to iterate orders")
	}

	return &orderv1.GetOrderByPublicOrderNumberResponse{Orders: orders}, nil
}

func (s *OrderServer) GetAllOrdersByFilters(ctx context.Context, req *orderv1.GetOrdersByFiltersRequest) (*orderv1.GetOrdersByFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByFilters not implemented")
}

func (s *OrderServer) UpdateOrder(ctx context.Context, req *orderv1.UpdateOrderRequest) (*orderv1.UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}

func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *orderv1.UpdateOrderStatusRequest) (*orderv1.UpdateOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}

func (s *OrderServer) CancelOrder(ctx context.Context, req *orderv1.CancelOrderRequest) (*orderv1.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
