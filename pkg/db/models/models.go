package models

import (
	"context"
	"database/sql"
	"errors"

	"github.com/HungryPandaHome/airdrop/pkg/werr"
	"github.com/jmoiron/sqlx"
)

var (
	// ErrNotFound ...
	ErrNotFound = errors.New("no rows")
	// ErrQuery ...
	ErrQuery = errors.New("query error")
)

// Contract ...
type Contract struct {
	ID      int64   `db:"id"`
	Address *string `db:"address_hex"`
}

// ContractModel ...
type ContractModel struct {
	DB *sqlx.DB
}

// ReadContract ...
func (m ContractModel) ReadContext(ctx context.Context) (*Contract, error) {
	var dst []*Contract
	err := m.DB.SelectContext(ctx, &dst, "select * from contract_address where id = 1;")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, werr.Wrap(ErrQuery, "%s", err)
	}
	if len(dst) == 0 {
		return nil, ErrNotFound
	}
	return nil, nil
}
