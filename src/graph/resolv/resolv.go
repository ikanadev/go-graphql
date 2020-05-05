package resolv

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v9"
	"github.com/vkevv/go-graphql/src/config"
	"github.com/vkevv/go-graphql/src/db"
	"github.com/vkevv/go-graphql/src/graph/model"
)

// Res handles resolvers logic
type Res struct {
	Conf   config.Config
	UserTx db.UserTx
	TodoTx db.TodoTx
}

// NewRes creates a Resolver
func NewRes(DB *pg.DB, conf config.Config) *Res {
	return &Res{
		Conf:   conf,
		UserTx: db.UserTx{DB: DB},
		TodoTx: db.TodoTx{DB: DB},
	}
}

// GenToken generates a token basen on UserID
func (r *Res) GenToken(ID string) (*model.AuthToken, error) {
	type authClaims struct {
		UserID string `json:"userId"`
		jwt.StandardClaims
	}
	expiredAt := time.Now().Add(time.Hour * 1)

	claims := authClaims{
		UserID: ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(r.Conf.App.JWT))
	if err != nil {
		return nil, err
	}
	return &model.AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil
}
