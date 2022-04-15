package data

import (
	_ "embed"

	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteStorage struct {
	db *gorm.DB
}

//go:embed sqlite_schema.sql
var sqliteSchema string

func NewSqliteStorage(dbFile string) (Storage, error) {
	db, err := gorm.Open(sqlite.Open(dbFile))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// init database
	_ = db.Exec(sqliteSchema)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &sqliteStorage{
		db: db,
	}, nil
}

func (s *sqliteStorage) Close() error {
	return nil
}

func (s *sqliteStorage) SavePlayers(data []Player) error {
	return s.db.Save(data).Error
}

func (s *sqliteStorage) LoadPlayers(config *BatchConfig) ([]Player, error) {
	var res []Player

	err := s.db.
		Limit(config.Limit).
		Offset(config.Offset).
		Find(&res).
		Error

	return res, err
}
func (s *sqliteStorage) SaveRating(data []Rating) error {
	return s.db.Save(data).Error
}

func (s *sqliteStorage) LoadRating(config *BatchConfig) ([]Rating, error) {
	var res []Rating

	err := s.db.
		Limit(config.Limit).
		Offset(config.Offset).
		Find(&res).
		Error

	return res, err
}
