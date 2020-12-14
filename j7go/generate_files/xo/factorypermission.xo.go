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

	"go.uber.org/zap"
)

// FactoryPermission represents a row from 'aypcddg.factory_permissions'.
type FactoryPermission struct {
	Fpid        uint           `json:"fpid"`         // fpid
	Name        sql.NullString `json:"name"`         // name
	ActionClass sql.NullString `json:"action_class"` // action_class
	Path        sql.NullString `json:"path"`         // path

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryPermission exists in the database.
func (fp *FactoryPermission) Exists() bool { //factory_permissions
	return fp._exists
}

// Deleted provides information if the FactoryPermission has been deleted from the database.
func (fp *FactoryPermission) Deleted() bool {
	return fp._deleted
}

// Get table name
func GetFactoryPermissionTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "factory_permissions", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryPermission to the database.
func (fp *FactoryPermission) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if fp._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryPermissionTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`fpid, name, action_class, path` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fp.Fpid, fp.Name, fp.ActionClass, fp.Path)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, fp.Fpid, fp.Name, fp.ActionClass, fp.Path)
	} else {
		res, err = dbConn.Exec(sqlstr, fp.Fpid, fp.Name, fp.ActionClass, fp.Path)
	}

	if err != nil {
		return err
	}

	// set existence
	fp._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	fp.Fpid = uint(id)
	fp._exists = true

	return nil
}

// Update updates the FactoryPermission in the database.
func (fp *FactoryPermission) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fp._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryPermissionTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`name = ?, action_class = ?, path = ?` +
		` WHERE fpid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fp.Name, fp.ActionClass, fp.Path, fp.Fpid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fp.Name, fp.ActionClass, fp.Path, fp.Fpid)
	} else {
		_, err = dbConn.Exec(sqlstr, fp.Name, fp.ActionClass, fp.Path, fp.Fpid)
	}
	return err
}

// Save saves the FactoryPermission to the database.
func (fp *FactoryPermission) Save(ctx context.Context) error {
	if fp.Exists() {
		return fp.Update(ctx)
	}

	return fp.Insert(ctx)
}

// Delete deletes the FactoryPermission from the database.
func (fp *FactoryPermission) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fp._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryPermissionTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE fpid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fp.Fpid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fp.Fpid)
	} else {
		_, err = dbConn.Exec(sqlstr, fp.Fpid)
	}

	if err != nil {
		return err
	}

	// set deleted
	fp._deleted = true

	return nil
}

// FactoryPermissionByFpid retrieves a row from 'aypcddg.factory_permissions' as a FactoryPermission.
//
// Generated from index 'factory_permissions_fpid_pkey'.
func FactoryPermissionByFpid(ctx context.Context, fpid uint, key ...interface{}) (*FactoryPermission, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryPermissionTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`fpid, name, action_class, path ` +
		`FROM ` + tableName +
		` WHERE fpid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fpid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	fp := FactoryPermission{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fpid).Scan(&fp.Fpid, &fp.Name, &fp.ActionClass, &fp.Path)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fpid).Scan(&fp.Fpid, &fp.Name, &fp.ActionClass, &fp.Path)
		if err != nil {
			return nil, err
		}
	}

	return &fp, nil
}