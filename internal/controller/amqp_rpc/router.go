package amqprpc

import (
	"github.com/nguyenthanhworkspace/golang-starter/internal/usecase"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
