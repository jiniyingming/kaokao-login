// Package xo contains the types for schema 'ddg_local'.
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

// AdminDeptRole represents a row from 'ddg_local.admin_dept_role'.
type AdminDeptRole struct {
	Adrid int16          `json:"adrid"` // adrid
	Title sql.NullString `json:"title"` // title
	Adid  int16          `json:"adid"`  // adid

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AdminDeptRole exists in the database.
func (adr *AdminDeptRole) Exists() bool { //admin_dept_role
	return adr._exists
}

// Deleted provides information if the AdminDeptRole has been deleted from the database.
func (adr *AdminDeptRole) Deleted() bool {
	return adr._deleted
}

// Get table name
func GetAdminDeptRoleTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "admin_dept_role", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the AdminDeptRole to the database.
func (adr *AdminDeptRole) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if adr._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdminDeptRoleTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`title, adid` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, adr.Title, adr.Adid)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, adr.Title, adr.Adid)
	} else {
		res, err = dbConn.Exec(sqlstr, adr.Title, adr.Adid)
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
	adr.Adrid = int16(id)
	adr._exists = true

	return nil
}

// Update updates the AdminDeptRole in the database.
func (adr *AdminDeptRole) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if adr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdminDeptRoleTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`title = ?, adid = ?` +
		` WHERE adrid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, adr.Title, adr.Adid, adr.Adrid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, adr.Title, adr.Adid, adr.Adrid)
	} else {
		_, err = dbConn.Exec(sqlstr, adr.Title, adr.Adid, adr.Adrid)
	}
	return err
}

// Save saves the AdminDeptRole to the database.
func (adr *AdminDeptRole) Save(ctx context.Context) error {
	if adr.Exists() {
		return adr.Update(ctx)
	}

	return adr.Insert(ctx)
}

// Delete deletes the AdminDeptRole from the database.
func (adr *AdminDeptRole) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if adr._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdminDeptRoleTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE adrid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, adr.Adrid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, adr.Adrid)
	} else {
		_, err = dbConn.Exec(sqlstr, adr.Adrid)
	}

	if err != nil {
		return err
	}

	// set deleted
	adr._deleted = true

	return nil
}

// AdminDeptRoleByAdrid retrieves a row from 'ddg_local.admin_dept_role' as a AdminDeptRole.
//
// Generated from index 'admin_dept_role_adrid_pkey'.
func AdminDeptRoleByAdrid(ctx context.Context, adrid int16, key ...interface{}) (*AdminDeptRole, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAdminDeptRoleTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`adrid, title, adid ` +
		`FROM ` + tableName +
		` WHERE adrid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, adrid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	adr := AdminDeptRole{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, adrid).Scan(&adr.Adrid, &adr.Title, &adr.Adid)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, adrid).Scan(&adr.Adrid, &adr.Title, &adr.Adid)
		if err != nil {
			return nil, err
		}
	}

	return &adr, nil
}