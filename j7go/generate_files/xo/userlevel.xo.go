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

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// UserLevel represents a row from 'ddg_local.user_level'.
type UserLevel struct {
	ID            int64          `json:"id"`              // id
	UserID        sql.NullInt64  `json:"user_id"`         // user_id
	UserType      sql.NullInt64  `json:"user_type"`       // user_type
	CreatedAt     mysql.NullTime `json:"created_at"`      // created_at
	UpdateAt      mysql.NullTime `json:"update_at"`       // update_at
	LevelConfigID sql.NullInt64  `json:"level_config_id"` // level_config_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserLevel exists in the database.
func (ul *UserLevel) Exists() bool { //user_level
	return ul._exists
}

// Deleted provides information if the UserLevel has been deleted from the database.
func (ul *UserLevel) Deleted() bool {
	return ul._deleted
}

// Get table name
func GetUserLevelTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "user_level", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserLevel to the database.
func (ul *UserLevel) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ul._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserLevelTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`user_id, user_type, created_at, update_at, level_config_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ul.UserID, ul.UserType, ul.CreatedAt, ul.UpdateAt, ul.LevelConfigID)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ul.UserID, ul.UserType, ul.CreatedAt, ul.UpdateAt, ul.LevelConfigID)
	} else {
		res, err = dbConn.Exec(sqlstr, ul.UserID, ul.UserType, ul.CreatedAt, ul.UpdateAt, ul.LevelConfigID)
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
	ul.ID = int64(id)
	ul._exists = true

	return nil
}

// Update updates the UserLevel in the database.
func (ul *UserLevel) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ul._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserLevelTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`user_id = ?, user_type = ?, created_at = ?, update_at = ?, level_config_id = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ul.UserID, ul.UserType, ul.CreatedAt, ul.UpdateAt, ul.LevelConfigID, ul.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ul.UserID, ul.UserType, ul.CreatedAt, ul.UpdateAt, ul.LevelConfigID, ul.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ul.UserID, ul.UserType, ul.CreatedAt, ul.UpdateAt, ul.LevelConfigID, ul.ID)
	}
	return err
}

// Save saves the UserLevel to the database.
func (ul *UserLevel) Save(ctx context.Context) error {
	if ul.Exists() {
		return ul.Update(ctx)
	}

	return ul.Insert(ctx)
}

// Delete deletes the UserLevel from the database.
func (ul *UserLevel) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ul._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserLevelTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ul.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ul.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ul.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ul._deleted = true

	return nil
}

// UserLevelByUserID retrieves a row from 'ddg_local.user_level' as a UserLevel.
//
// Generated from index 'user_id'.
func UserLevelByUserID(ctx context.Context, userID sql.NullInt64, key ...interface{}) (*UserLevel, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserLevelTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, user_id, user_type, created_at, update_at, level_config_id ` +
		`FROM ` + tableName +
		` WHERE user_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, userID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ul := UserLevel{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, userID).Scan(&ul.ID, &ul.UserID, &ul.UserType, &ul.CreatedAt, &ul.UpdateAt, &ul.LevelConfigID)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, userID).Scan(&ul.ID, &ul.UserID, &ul.UserType, &ul.CreatedAt, &ul.UpdateAt, &ul.LevelConfigID)
		if err != nil {
			return nil, err
		}
	}

	return &ul, nil
}

// UserLevelByID retrieves a row from 'ddg_local.user_level' as a UserLevel.
//
// Generated from index 'user_level_id_pkey'.
func UserLevelByID(ctx context.Context, id int64, key ...interface{}) (*UserLevel, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserLevelTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, user_id, user_type, created_at, update_at, level_config_id ` +
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
	ul := UserLevel{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&ul.ID, &ul.UserID, &ul.UserType, &ul.CreatedAt, &ul.UpdateAt, &ul.LevelConfigID)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&ul.ID, &ul.UserID, &ul.UserType, &ul.CreatedAt, &ul.UpdateAt, &ul.LevelConfigID)
		if err != nil {
			return nil, err
		}
	}

	return &ul, nil
}