package db

import (
	"context"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	poolConf, err := pgxpool.ParseConfig(getConnectionString())
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(ctx, poolConf)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func getConnectionString() string {
	var builder strings.Builder

	builder.WriteString("sslmode=disable")
	builder.WriteString(" dbname=")
	builder.WriteString("admin")
	builder.WriteString(" user=")
	builder.WriteString("admin")
	builder.WriteString(" password=")
	builder.WriteString("admin")
	builder.WriteString(" host=")
	builder.WriteString("51.250.105.228")
	builder.WriteString(" port=")
	builder.WriteString(strconv.Itoa(5432))
	builder.WriteString(" connect_timeout=")
	builder.WriteString(strconv.FormatUint(10, 10))

	return builder.String()
}
