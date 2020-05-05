package graph

import (
	"github.com/go-pg/pg/v9"
	"github.com/vkevv/go-graphql/src/config"
	"github.com/vkevv/go-graphql/src/graph/resolv"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver the main struct for resolvers
type Resolver struct {
	Res  *resolv.Res
	Conf config.Config
}

// NewResolver Returns a new Resolver
func NewResolver(DB *pg.DB, conf config.Config) *Resolver {
	return &Resolver{
		Res:  resolv.NewRes(DB, conf),
		Conf: conf,
	}
}
