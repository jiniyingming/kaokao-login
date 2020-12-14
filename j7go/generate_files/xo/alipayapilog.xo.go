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

// AlipayAPILog represents a row from 'aypcddg.alipay_api_log'.
type AlipayAPILog struct {
	ID        uint64         `json:"id"`         // id
	Title     sql.NullString `json:"title"`      // title
	Content   sql.NullString `json:"content"`    // content
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AlipayAPILog exists in the database.
func (aal *AlipayAPILog) Exists() bool { //alipay_api_log
	return aal._exists
}

// Deleted provides information if the AlipayAPILog has been deleted from the database.
func (aal *AlipayAPILog) Deleted() bool {
	return aal._deleted
}

// Get table name
func GetAlipayAPILogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "alipay_api_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the AlipayAPILog to the database.
func (aal *AlipayAPILog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if aal._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAlipayAPILogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`title, content, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aal.Title, aal.Content, aal.CreatedAt, aal.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, aal.Title, aal.Content, aal.CreatedAt, aal.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, aal.Title, aal.Content, aal.CreatedAt, aal.UpdatedAt)
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
	aal.ID = uint64(id)
	aal._exists = true

	return nil
}

// Update updates the AlipayAPILog in the database.
func (aal *AlipayAPILog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if aal._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAlipayAPILogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`title = ?, content = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aal.Title, aal.Content, aal.CreatedAt, aal.UpdatedAt, aal.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, aal.Title, aal.Content, aal.CreatedAt, aal.UpdatedAt, aal.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, aal.Title, aal.Content, aal.CreatedAt, aal.UpdatedAt, aal.ID)
	}
	return err
}

// Save saves the AlipayAPILog to the database.
func (aal *AlipayAPILog) Save(ctx context.Context) error {
	if aal.Exists() {
		return aal.Update(ctx)
	}

	return aal.Insert(ctx)
}

// Delete deletes the AlipayAPILog from the database.
func (aal *AlipayAPILog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if aal._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAlipayAPILogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aal.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, aal.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, aal.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	aal._deleted = true

	return nil
}

// AlipayAPILogByID retrieves a row from 'aypcddg.alipay_api_log' as a AlipayAPILog.
//
// Generated from index 'alipay_api_log_id_pkey'.
func AlipayAPILogByID(ctx context.Context, id uint64, key ...interface{}) (*AlipayAPILog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAlipayAPILogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, title, content, created_at, updated_at ` +
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
	aal := AlipayAPILog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&aal.ID, &aal.Title, &aal.Content, &aal.CreatedAt, &aal.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&aal.ID, &aal.Title, &aal.Content, &aal.CreatedAt, &aal.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &aal, nil
}