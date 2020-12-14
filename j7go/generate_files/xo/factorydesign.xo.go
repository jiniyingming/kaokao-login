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

// FactoryDesign represents a row from 'aypcddg.factory_design'.
type FactoryDesign struct {
	ID      uint64         `json:"id"`      // id
	Fid     int64          `json:"fid"`     // fid
	Version int            `json:"version"` // version
	Code    sql.NullString `json:"code"`    // code
	Key     sql.NullString `json:"key"`     // key
	Type    sql.NullInt64  `json:"type"`    // type

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryDesign exists in the database.
func (fd *FactoryDesign) Exists() bool { //factory_design
	return fd._exists
}

// Deleted provides information if the FactoryDesign has been deleted from the database.
func (fd *FactoryDesign) Deleted() bool {
	return fd._deleted
}

// Get table name
func GetFactoryDesignTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "factory_design", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryDesign to the database.
func (fd *FactoryDesign) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if fd._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryDesignTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`fid, version, code, key, type` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fd.Fid, fd.Version, fd.Code, fd.Key, fd.Type)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, fd.Fid, fd.Version, fd.Code, fd.Key, fd.Type)
	} else {
		res, err = dbConn.Exec(sqlstr, fd.Fid, fd.Version, fd.Code, fd.Key, fd.Type)
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
	fd.ID = uint64(id)
	fd._exists = true

	return nil
}

// Update updates the FactoryDesign in the database.
func (fd *FactoryDesign) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fd._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryDesignTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`fid = ?, version = ?, code = ?, key = ?, type = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fd.Fid, fd.Version, fd.Code, fd.Key, fd.Type, fd.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fd.Fid, fd.Version, fd.Code, fd.Key, fd.Type, fd.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, fd.Fid, fd.Version, fd.Code, fd.Key, fd.Type, fd.ID)
	}
	return err
}

// Save saves the FactoryDesign to the database.
func (fd *FactoryDesign) Save(ctx context.Context) error {
	if fd.Exists() {
		return fd.Update(ctx)
	}

	return fd.Insert(ctx)
}

// Delete deletes the FactoryDesign from the database.
func (fd *FactoryDesign) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fd._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryDesignTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fd.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fd.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, fd.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	fd._deleted = true

	return nil
}

// FactoryDesignByID retrieves a row from 'aypcddg.factory_design' as a FactoryDesign.
//
// Generated from index 'factory_design_id_pkey'.
func FactoryDesignByID(ctx context.Context, id uint64, key ...interface{}) (*FactoryDesign, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryDesignTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, fid, version, code, key, type ` +
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
	fd := FactoryDesign{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&fd.ID, &fd.Fid, &fd.Version, &fd.Code, &fd.Key, &fd.Type)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&fd.ID, &fd.Fid, &fd.Version, &fd.Code, &fd.Key, &fd.Type)
		if err != nil {
			return nil, err
		}
	}

	return &fd, nil
}