package database

import (
	"context"
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
	sqlConn.SetMaxIdleConns(d.cfg.SetMaxIdleConns)
	sqlConn.SetMaxOpenConns(d.cfg.SetMaxOpenConns)
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

type SessionTxKey struct{}

type Session interface {
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Create(interface{}) (tx *gorm.DB)
	CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB)
	Update(column string, value interface{}) (tx *gorm.DB)
	UpdateColumns(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Exec(sql string, values ...interface{}) (tx *gorm.DB)
	WithContext(ctx context.Context) *gorm.DB
}

func (d *Database) Create(i interface{}) (tx *gorm.DB) {
	return d.Conn.Create(i)
}

func (d *Database) CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB) {
	return d.Conn.CreateInBatches(value, batchSize)
}

func (d *Database) Update(column string, value interface{}) (tx *gorm.DB) {
	return d.Conn.Update(column, value)
}

func (d *Database) UpdateColumns(values interface{}) (tx *gorm.DB) {
	return d.UpdateColumns(values)
}

func (d *Database) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
	return d.Conn.Delete(value, conds)
}

func (d *Database) Exec(sql string, values ...interface{}) (tx *gorm.DB) {
	return d.Conn.Exec(sql, values)
}

func (d *Database) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return d.Conn.Find(dest, conds)
}

func (d *Database) WithContext(ctx context.Context) *gorm.DB {
	return d.Conn.Session(&gorm.Session{Context: ctx})
}

func (d *Database) GetSession(ctx context.Context) Session {
	if value := ctx.Value(SessionTxKey{}); value != nil {
		if v, ok := value.(*gorm.DB); ok {
			return v
		}
	}

	return d
}

func (d *Database) Session(ctx context.Context) (context.Context, func(error), error) {
	if value := ctx.Value(SessionTxKey{}); value != nil {
		if v, ok := value.(*gorm.DB); ok {
			return ctx, save(v), nil
		}
	}
	tx := d.Conn.Begin(nil)
	ctx = context.WithValue(ctx, SessionTxKey{}, tx)

	return ctx, save(tx), nil
}

func save(tx *gorm.DB) func(error) {
	return func(err error) {
		if tx == nil {
			return
		}
		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}
}
