package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	Conn *gorm.DB
	cfg  *DatabaseConfig
}

func NewDatabase(cfg *DatabaseConfig) (*Database, error) {
	db := &Database{
		cfg: cfg,
	}

	return db, nil
}

func (d *Database) Open() error {
	cfg := &gorm.Config{
		SkipDefaultTransaction: d.cfg.SkipDefaultTransaction,
	}
	conn, err := gorm.Open(postgres.Open(d.cfg.getDsn()), cfg)
	if err != nil {
		return err
	}

	sqlConn, err := conn.DB()
	// SetMaxIdleConns устанавливает максимальное количество соединений в пуле бездействия.
	sqlConn.SetMaxIdleConns(d.cfg.SetMaxIdleConns)
	// SetMaxOpenConns устанавливает максимальное количество открытых соединений с БД.
	sqlConn.SetMaxOpenConns(d.cfg.SetMaxOpenConns)
	// SetConnMaxLifetime устанавливает максимальное время повторного использования соединения.
	sqlConn.SetConnMaxLifetime(time.Duration(d.cfg.SetConnMaxLifetime) * time.Second)

	err = sqlConn.Ping()
	if err != nil {
		return err
	}

	d.Conn = conn

	return nil
}

func (d *Database) Close() error {
	db, err := d.Conn.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}
