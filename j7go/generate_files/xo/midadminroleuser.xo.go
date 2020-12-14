// Package xo contains the types for schema 'aypcddg'.
package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"fmt"
	"j7go/components"
	"j7go/utils"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// MidAdminRoleUser represents a row from 'aypcddg.mid_admin_role_users'.
type MidAdminRoleUser struct {
	RoleID    int            `json:"role_id"`    // role_id
	UserID    int            `json:"user_id"`    // user_id
	CreatedAt mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt mysql.NullTime `json:"updated_at"` // updated_at
}

// MidAdminRoleUsersByRoleIDUserID retrieves a row from 'aypcddg.mid_admin_role_users' as a MidAdminRoleUser.
//
// Generated from index 'mid_admin_role_users_role_id_user_id_index'.
func MidAdminRoleUsersByRoleIDUserID(ctx context.Context, roleID int, userID int, key ...interface{}) ([]*MidAdminRoleUser, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminRoleUserTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`role_id, user_id, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE role_id = ? AND user_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, roleID, userID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, roleID, userID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, roleID, userID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*MidAdminRoleUser, 0)
	for queryData.Next() {
		maru := MidAdminRoleUser{}

		// scan
		err = queryData.Scan(&maru.RoleID, &maru.UserID, &maru.CreatedAt, &maru.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &maru)
	}

	return res, nil
}