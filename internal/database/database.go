package database


import (
	"math"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/TomDev24/MusicServiceGo/pkg/pagination"
	"github.com/TomDev24/MusicServiceGo/internal/config"
)

type Database struct {
	Conn *gorm.DB
}

func Init(cfg *config.Config) (*Database, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Db.User,
		cfg.Db.Pass,
		cfg.Db.Adress,
		cfg.Db.Port,
		cfg.Db.Name,
	)

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
    if err != nil {
		return nil, err
    }

	return &Database{
		Conn: db,
	}, nil
}

func Paginate(value interface{}, pagination *pagination.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
    var totalRows int64
    db.Model(value).Count(&totalRows)

    pagination.TotalRows = totalRows
    totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
    pagination.TotalPages = totalPages

    return func(db *gorm.DB) *gorm.DB {
        return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
    }
}
