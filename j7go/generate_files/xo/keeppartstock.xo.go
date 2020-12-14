// Package xo contains the types for schema 'aypcddg'.
package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"j7go/components"
	"j7go/utils"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// KeepPartStock represents a row from 'aypcddg.keep_part_stock'.
type KeepPartStock struct {
	ID               uint64         `json:"id"`                 // id
	Gid              int            `json:"gid"`                // gid
	KeepStock        int            `json:"keep_stock"`         // keep_stock
	OriginTotalStock int            `json:"origin_total_stock"` // origin_total_stock
	KeepScale        int            `json:"keep_scale"`         // keep_scale
	PutStock         int            `json:"put_stock"`          // put_stock
	CreatedAt        mysql.NullTime `json:"created_at"`         // created_at
	UpdatedAt        mysql.NullTime `json:"updated_at"`         // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the KeepPartStock exists in the database.
func (kps *KeepPartStock) Exists() bool { //keep_part_stock
	return kps._exists
}

// Deleted provides information if the KeepPartStock has been deleted from the database.
func (kps *KeepPartStock) Deleted() bool {
	return kps._deleted
}

// Get table name
func GetKeepPartStockTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "keep_part_stock", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the KeepPartStock to the database.
func (kps *KeepPartStock) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if kps._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetKeepPartStockTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`gid, keep_stock, origin_total_stock, keep_scale, put_stock, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, kps.Gid, kps.KeepStock, kps.OriginTotalStock, kps.KeepScale, kps.PutStock, kps.CreatedAt, kps.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, kps.Gid, kps.KeepStock, kps.OriginTotalStock, kps.KeepScale, kps.PutStock, kps.CreatedAt, kps.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, kps.Gid, kps.KeepStock, kps.OriginTotalStock, kps.KeepScale, kps.PutStock, kps.CreatedAt, kps.UpdatedAt)
	}

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	kps.ID = uint64(id)
	kps._exists = true

	return nil
}

// Update updates the KeepPartStock in the database.
func (kps *KeepPartStock) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if kps._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetKeepPartStockTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`gid = ?, keep_stock = ?, origin_total_stock = ?, keep_scale = ?, put_stock = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, kps.Gid, kps.KeepStock, kps.OriginTotalStock, kps.KeepScale, kps.PutStock, kps.CreatedAt, kps.UpdatedAt, kps.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, kps.Gid, kps.KeepStock, kps.OriginTotalStock, kps.KeepScale, kps.PutStock, kps.CreatedAt, kps.UpdatedAt, kps.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, kps.Gid, kps.KeepStock, kps.OriginTotalStock, kps.KeepScale, kps.PutStock, kps.CreatedAt, kps.UpdatedAt, kps.ID)
	}
	return err
}

// Save saves the KeepPartStock to the database.
func (kps *KeepPartStock) Save(ctx context.Context) error {
	if kps.Exists() {
		return kps.Update(ctx)
	}

	return kps.Insert(ctx)
}

// Delete deletes the KeepPartStock from the database.
func (kps *KeepPartStock) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if kps._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetKeepPartStockTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, kps.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, kps.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, kps.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	kps._deleted = true

	return nil
}

// KeepPartStockByID retrieves a row from 'aypcddg.keep_part_stock' as a KeepPartStock.
//
// Generated from index 'keep_part_stock_id_pkey'.
func KeepPartStockByID(ctx context.Context, id uint64, key ...interface{}) (*KeepPartStock, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetKeepPartStockTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, gid, keep_stock, origin_total_stock, keep_scale, put_stock, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, id)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	kps := KeepPartStock{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&kps.ID, &kps.Gid, &kps.KeepStock, &kps.OriginTotalStock, &kps.KeepScale, &kps.PutStock, &kps.CreatedAt, &kps.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&kps.ID, &kps.Gid, &kps.KeepStock, &kps.OriginTotalStock, &kps.KeepScale, &kps.PutStock, &kps.CreatedAt, &kps.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &kps, nil
}