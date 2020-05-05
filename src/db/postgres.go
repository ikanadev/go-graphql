package db

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// Logger struct to log database actions
type Logger struct{}

// BeforeQuery trigger to capture query before execution
func (d Logger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

// AfterQuery trigger to capture query after execution
func (d Logger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

// New creates a new instance of postgres DB
func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	return db
}

// CreateTables create tables based on interfaces
func CreateTables(db *pg.DB, structs ...interface{}) error {
	for _, model := range structs {
		err := db.CreateTable(&model, &orm.CreateTableOptions{
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
