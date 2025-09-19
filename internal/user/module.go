package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	Service *Service
	Handler *Handler
}

func InitModule(db *pgxpool.Pool) *Module {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	return &Module{Service: service, Handler: handler}
}

func (m *Module) RegisterRoutes(r *gin.Engine) {
	m.Handler.RegisterRoutes(r)
}
