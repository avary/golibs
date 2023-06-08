package auth

import (
	"context"
	"time"

	"github.com/skerkour/golibs/db"
	"github.com/skerkour/golibs/uuid"
)

type Account struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateAccount(ctx context.Context, db db.Queryer, accountID uuid.UUID, password string) (err error) {
	panic("TODO")
}

func DeleteAccount(ctx context.Context, db db.Queryer, accountID uuid.UUID) (err error) {
	panic("TODO")
}
