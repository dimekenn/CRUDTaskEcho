package app

import (
	"context"
	"crudTaskEcho/configs"
	"crudTaskEcho/internal/app/handler"
	"crudTaskEcho/internal/app/repository"
	"crudTaskEcho/internal/app/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"time"
)

func StartHTTPServer(ctx context.Context, errCh chan<- error, cfg *configs.Configs) {
	app := echo.New()

	pool, poolErr := initDb(ctx, cfg.DbUrl)
	if poolErr != nil {
		errCh <- poolErr
		return
	}
	repo := repository.NewCRUDRepository(pool)
	crudService := service.NewCRUDService(repo)
	crudHandler := handler.NewCRUDHandler(crudService)

	app.GET("/user/:id", crudHandler.GetUserByUUID)
	app.POST("/user", crudHandler.CreateUser)
	app.PUT("/user", crudHandler.UpdateUserByUUID)

	errCh <- app.Start(cfg.Port)
}

func initDb(ctx context.Context, url string) (*pgxpool.Pool, error) {
	conf, cfgErr := pgxpool.ParseConfig(url)
	if cfgErr != nil {
		return nil, cfgErr
	}
	conf.MaxConns = 20
	conf.MinConns = 10
	conf.MaxConnIdleTime = 10 * time.Second

	conn, connErr := pgxpool.ConnectConfig(ctx, conf)
	if connErr != nil {
		return nil, connErr
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}
	return conn, nil
}
