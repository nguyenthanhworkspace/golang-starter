// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/mysql"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/nguyenthanhworkspace/golang-starter/config"
	amqprpc "github.com/nguyenthanhworkspace/golang-starter/internal/controller/amqp_rpc"
	v1 "github.com/nguyenthanhworkspace/golang-starter/internal/controller/http/v1"
	"github.com/nguyenthanhworkspace/golang-starter/internal/usecase"
	"github.com/nguyenthanhworkspace/golang-starter/internal/usecase/repo"
	"github.com/nguyenthanhworkspace/golang-starter/internal/usecase/webapi"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/httpserver"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/logger"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/postgres"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	dbSql, err := mysql.New(cfg.MYSQL.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mysql.New: %w", err))
	}
	defer dbSql.Close()

	// Use case
	translationUseCase := usecase.New(
		repo.New(pg),
		webapi.New(),
	)
	userUseCase := usecase.NewUserUseCase(
		repo.NewUserRepo(dbSql),
	)

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(translationUseCase)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// HTTP Server
	handler := gin.New()
	useCaseList := v1.NewUseCaseList(translationUseCase, userUseCase)
	v1.NewRouter(handler, l, useCaseList)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
