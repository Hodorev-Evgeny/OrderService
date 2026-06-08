package main

import (
	core_logger "OrderService/internal/core/logger"
	"OrderService/internal/core/transport/server"
	"context"
	"fmt"
	"os/signal"
	"syscall"

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

	server_config := server.MustGetServerConfig()
	server := server.NewServer(
		server_config,
	)

	logger.Info("Starting server")
	if err := server.Start(ctx); err != nil {
		logger.Error("Error starting server", zap.Error(err))
	}
}
