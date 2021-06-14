package db

import (
	"embed"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/HungryPandaHome/airdrop/pkg/utils"
	"github.com/HungryPandaHome/airdrop/pkg/werr"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
)

const (
	// CurrentSchemaVersion ...
	CurrentSchemaVersion = 1
)

var (
	// ErrOpen ...
	ErrOpen = errors.New("failed to open url")
	// ErrConn ...
	ErrConn = errors.New("connection failed")
	// ErrUri ...
	ErrUri = errors.New("bad uri")
	// ErrSources ...
	ErrSources = errors.New("bad sources")

	// ErrMigrate ...
	ErrMigrate = errors.New("migration init error")
)

//go:embed migrations/*
var migrations embed.FS

// Sources ...
func Sources() (source.Driver, error) {
	return httpfs.New(http.FS(migrations), "migrations")
}

// Open ...
func Open(uri string) (*sqlx.DB, error) {
	// wait for postgresql connection
	db, err := sqlx.Open("postgres", uri)
	if err != nil {
		return nil, werr.Wrap(ErrUri, "%s", err)
	}
	if err = utils.Retry(db.Ping, 5); err != nil {
		return nil, werr.Wrap(ErrConn, "%s", err)
	}

	return db, err
}

// GetMigrationInstance ...
func GetMigrationInstance(db *sqlx.DB, uri string) (*migrate.Migrate, error) {
	pu, err := url.Parse(uri)
	if err != nil {
		return nil, werr.Wrap(ErrUri, "%s", err)
	}
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, werr.Wrap(ErrMigrate, "%s", err)
	}
	sources, err := Sources()
	if err != nil {
		return nil, werr.Wrap(ErrSources, "%s", err)
	}

	ss := strings.Split(pu.Path, "/")
	dbName := ss[len(ss)-1]
	m, err := migrate.NewWithInstance("httpfs", sources,
		dbName, driver)
	if err != nil {
		return nil, werr.Wrap(ErrMigrate, "%s", err)
	}
	return m, nil
}

// Migrate ...
func Migrate(db *sqlx.DB, uri string, version int) error {
	mr, err := GetMigrationInstance(db, uri)
	if err != nil {
		return err
	}
	if err := mr.Migrate(CurrentSchemaVersion); err != nil &&
		!errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
