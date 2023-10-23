package db

import (
	"context"
	"strconv"
	"strings"
)

func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	poolConf, err := pgxpool.ParseConfig(getConnectionString(cfg))
	if err != nil {
		return nil, err
	}

	poolConf.MaxConnLifetime = cfg.ConnMaxLifetime
	poolConf.MaxConns = cfg.MaxOpenConns
	if cfg.PreferSimpleProtocol {
		poolConf.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConf)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func getConnectionString(cfg config.DBConfig) string {
	var builder strings.Builder

	builder.WriteString("sslmode=disable")
	builder.WriteString(" dbname=")
	builder.WriteString(cfg.Database)
	builder.WriteString(" user=")
	builder.WriteString(cfg.Login)
	builder.WriteString(" password=")
	builder.WriteString(cfg.Password)
	builder.WriteString(" host=")
	builder.WriteString(cfg.Host.Name)
	builder.WriteString(" port=")
	builder.WriteString(strconv.Itoa(cfg.Host.Port))
	builder.WriteString(" connect_timeout=")
	builder.WriteString(strconv.FormatUint(cfg.ConnectTimeout, 10))

	return builder.String()
}
