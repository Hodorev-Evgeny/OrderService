package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	core_logger "github.com/Hodorev-Evgeny/OrderService/internal/core/logger"
	core_pgx_pool "github.com/Hodorev-Evgeny/OrderService/internal/core/repository/postgres/pgx"
	"github.com/Hodorev-Evgeny/OrderService/internal/core/transport/server"
	feature_order_repository "github.com/Hodorev-Evgeny/OrderService/internal/feature/orders/repository"
	feature_order_service "github.com/Hodorev-Evgeny/OrderService/internal/feature/orders/service"
	feature_order_transport "github.com/Hodorev-Evgeny/OrderService/internal/feature/orders/transport"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	logger_config := core_logger.MustNewConfig()
	logger, err := core_logger.NewLogger(logger_config)
	if err != nil {
		fmt.Println("Error initializing logger")
	}
	ctx = core_logger.ToContext(ctx, logger)

	configPool := core_pgx_pool.MustPostgresConfig()
	pool := core_pgx_pool.CreatePoolMust(ctx, configPool)

	orderCaseRepository := feature_order_repository.NewOrderRepository(pool)
	orderCaseServise := feature_order_service.NewOrderService(orderCaseRepository)
	orderCase := feature_order_transport.NewOrderTransport(orderCaseServise)

	server_config := server.MustGetServerConfig()
	server := server.NewServer(
		server_config,
		orderCase,
	)

	logger.Info("Starting server")
	if err := server.Start(ctx); err != nil {
		logger.Error("Error starting server", zap.Error(err))
	}
}
