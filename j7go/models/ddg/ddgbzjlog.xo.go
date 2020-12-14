// Package xo contains the types for schema 'ddg_local'.
package ddg

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"j7go/components"
	"j7go/utils"

	"go.uber.org/zap"
)

// DdgBzjLog represents a row from 'ddg_local.ddg_bzj_log'.
type DdgBzjLog struct {
	ID           int     `json:"id"`            // id
	Type         int64   `json:"type"`          // type
	OperatorID   uint64  `json:"operator_id"`   // operator_id
	OperatorName string  `json:"operator_name"` // operator_name
	RelatedID    int64   `json:"related_id"`    // related_id
	Price        float64 `json:"price"`         // price
	Created      int64   `json:"created"`       // created
	Updated      int64   `json:"updated"`       // updated
	Action       int64   `json:"action"`        // action
	BzjID        int     `json:"bzj_id"`        // bzj_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DdgBzjLog exists in the database.
func (dbl *DdgBzjLog) Exists() bool { //ddg_bzj_log
	return dbl._exists
}

// Deleted provides information if the DdgBzjLog has been deleted from the database.
func (dbl *DdgBzjLog) Deleted() bool {
	return dbl._deleted
}

// Get table name
func GetDdgBzjLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "ddg_bzj_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DdgBzjLog to the database.
func (dbl *DdgBzjLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if dbl._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDdgBzjLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`type, operator_id, operator_name, related_id, price, created, updated, action, bzj_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dbl.Type, dbl.OperatorID, dbl.OperatorName, dbl.RelatedID, dbl.Price, dbl.Created, dbl.Updated, dbl.Action, dbl.BzjID)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, dbl.Type, dbl.OperatorID, dbl.OperatorName, dbl.RelatedID, dbl.Price, dbl.Created, dbl.Updated, dbl.Action, dbl.BzjID)
	} else {
		res, err = dbConn.Exec(sqlstr, dbl.Type, dbl.OperatorID, dbl.OperatorName, dbl.RelatedID, dbl.Price, dbl.Created, dbl.Updated, dbl.Action, dbl.BzjID)
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
	dbl.ID = int(id)
	dbl._exists = true

	return nil
}

// Update updates the DdgBzjLog in the database.
func (dbl *DdgBzjLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dbl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	if !dbl._exists {
		return errors.New("记录不存在")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDdgBzjLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`type = ?, operator_id = ?, operator_name = ?, related_id = ?, price = ?, created = ?, updated = ?, action = ?, bzj_id = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dbl.Type, dbl.OperatorID, dbl.OperatorName, dbl.RelatedID, dbl.Price, dbl.Created, dbl.Updated, dbl.Action, dbl.BzjID, dbl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dbl.Type, dbl.OperatorID, dbl.OperatorName, dbl.RelatedID, dbl.Price, dbl.Created, dbl.Updated, dbl.Action, dbl.BzjID, dbl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dbl.Type, dbl.OperatorID, dbl.OperatorName, dbl.RelatedID, dbl.Price, dbl.Created, dbl.Updated, dbl.Action, dbl.BzjID, dbl.ID)
	}
	return err
}

// Save saves the DdgBzjLog to the database.
func (dbl *DdgBzjLog) Save(ctx context.Context) error {
	if dbl.Exists() {
		return dbl.Update(ctx)
	}

	return dbl.Insert(ctx)
}

// Delete deletes the DdgBzjLog from the database.
func (dbl *DdgBzjLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dbl._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDdgBzjLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dbl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dbl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dbl.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	dbl._deleted = true

	return nil
}

// DdgBzjLogByID retrieves a row from 'ddg_local.ddg_bzj_log' as a DdgBzjLog.
//
// Generated from index 'ddg_bzj_log_id_pkey'.
func DdgBzjLogByID(ctx context.Context, id int, key ...interface{}) (*DdgBzjLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDdgBzjLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, type, operator_id, operator_name, related_id, price, created, updated, action, bzj_id ` +
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
	dbl := DdgBzjLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&dbl.ID, &dbl.Type, &dbl.OperatorID, &dbl.OperatorName, &dbl.RelatedID, &dbl.Price, &dbl.Created, &dbl.Updated, &dbl.Action, &dbl.BzjID)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&dbl.ID, &dbl.Type, &dbl.OperatorID, &dbl.OperatorName, &dbl.RelatedID, &dbl.Price, &dbl.Created, &dbl.Updated, &dbl.Action, &dbl.BzjID)
		if err != nil {
			return nil, err
		}
	}

	return &dbl, nil
}