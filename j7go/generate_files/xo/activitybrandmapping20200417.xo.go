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

// ActivityBrandMapping20200417 represents a row from 'ddg_local.activity_brand_mapping_20200417'.
type ActivityBrandMapping20200417 struct {
	ID                uint           `json:"id"`                  // id
	ActivityID        int            `json:"activity_id"`         // activity_id
	FactoryActivityID int            `json:"factory_activity_id"` // factory_activity_id
	BrandID           int            `json:"brand_id"`            // brand_id
	FactoryBrandID    int            `json:"factory_brand_id"`    // factory_brand_id
	Status            int8           `json:"status"`              // status
	CreatedAt         mysql.NullTime `json:"created_at"`          // created_at
	UpdatedAt         mysql.NullTime `json:"updated_at"`          // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ActivityBrandMapping20200417 exists in the database.
func (abm *ActivityBrandMapping20200417) Exists() bool { //activity_brand_mapping_20200417
	return abm._exists
}

// Deleted provides information if the ActivityBrandMapping20200417 has been deleted from the database.
func (abm *ActivityBrandMapping20200417) Deleted() bool {
	return abm._deleted
}

// Get table name
func GetActivityBrandMapping20200417TableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "activity_brand_mapping_20200417", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the ActivityBrandMapping20200417 to the database.
func (abm *ActivityBrandMapping20200417) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if abm._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityBrandMapping20200417TableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`activity_id, factory_activity_id, brand_id, factory_brand_id, status, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, abm.ActivityID, abm.FactoryActivityID, abm.BrandID, abm.FactoryBrandID, abm.Status, abm.CreatedAt, abm.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, abm.ActivityID, abm.FactoryActivityID, abm.BrandID, abm.FactoryBrandID, abm.Status, abm.CreatedAt, abm.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, abm.ActivityID, abm.FactoryActivityID, abm.BrandID, abm.FactoryBrandID, abm.Status, abm.CreatedAt, abm.UpdatedAt)
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
	abm.ID = uint(id)
	abm._exists = true

	return nil
}

// Update updates the ActivityBrandMapping20200417 in the database.
func (abm *ActivityBrandMapping20200417) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if abm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityBrandMapping20200417TableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`activity_id = ?, factory_activity_id = ?, brand_id = ?, factory_brand_id = ?, status = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, abm.ActivityID, abm.FactoryActivityID, abm.BrandID, abm.FactoryBrandID, abm.Status, abm.CreatedAt, abm.UpdatedAt, abm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, abm.ActivityID, abm.FactoryActivityID, abm.BrandID, abm.FactoryBrandID, abm.Status, abm.CreatedAt, abm.UpdatedAt, abm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, abm.ActivityID, abm.FactoryActivityID, abm.BrandID, abm.FactoryBrandID, abm.Status, abm.CreatedAt, abm.UpdatedAt, abm.ID)
	}
	return err
}

// Save saves the ActivityBrandMapping20200417 to the database.
func (abm *ActivityBrandMapping20200417) Save(ctx context.Context) error {
	if abm.Exists() {
		return abm.Update(ctx)
	}

	return abm.Insert(ctx)
}

// Delete deletes the ActivityBrandMapping20200417 from the database.
func (abm *ActivityBrandMapping20200417) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if abm._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityBrandMapping20200417TableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, abm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, abm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, abm.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	abm._deleted = true

	return nil
}

// ActivityBrandMapping20200417ByID retrieves a row from 'ddg_local.activity_brand_mapping_20200417' as a ActivityBrandMapping20200417.
//
// Generated from index 'activity_brand_mapping_20200417_id_pkey'.
func ActivityBrandMapping20200417ByID(ctx context.Context, id uint, key ...interface{}) (*ActivityBrandMapping20200417, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityBrandMapping20200417TableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, brand_id, factory_brand_id, status, created_at, updated_at ` +
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
	abm := ActivityBrandMapping20200417{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&abm.ID, &abm.ActivityID, &abm.FactoryActivityID, &abm.BrandID, &abm.FactoryBrandID, &abm.Status, &abm.CreatedAt, &abm.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&abm.ID, &abm.ActivityID, &abm.FactoryActivityID, &abm.BrandID, &abm.FactoryBrandID, &abm.Status, &abm.CreatedAt, &abm.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &abm, nil
}

// ActivityBrandMapping20200417sByActivityID retrieves a row from 'ddg_local.activity_brand_mapping_20200417' as a ActivityBrandMapping20200417.
//
// Generated from index 'activity_brand_mapping_activity_id_index'.
func ActivityBrandMapping20200417sByActivityID(ctx context.Context, activityID int, key ...interface{}) ([]*ActivityBrandMapping20200417, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityBrandMapping20200417TableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, brand_id, factory_brand_id, status, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE activity_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, activityID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, activityID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, activityID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*ActivityBrandMapping20200417, 0)
	for queryData.Next() {
		abm := ActivityBrandMapping20200417{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&abm.ID, &abm.ActivityID, &abm.FactoryActivityID, &abm.BrandID, &abm.FactoryBrandID, &abm.Status, &abm.CreatedAt, &abm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &abm)
	}

	return res, nil
}

// ActivityBrandMapping20200417sByBrandID retrieves a row from 'ddg_local.activity_brand_mapping_20200417' as a ActivityBrandMapping20200417.
//
// Generated from index 'activity_brand_mapping_brand_id_index'.
func ActivityBrandMapping20200417sByBrandID(ctx context.Context, brandID int, key ...interface{}) ([]*ActivityBrandMapping20200417, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityBrandMapping20200417TableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, brand_id, factory_brand_id, status, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE brand_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, brandID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, brandID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, brandID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*ActivityBrandMapping20200417, 0)
	for queryData.Next() {
		abm := ActivityBrandMapping20200417{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&abm.ID, &abm.ActivityID, &abm.FactoryActivityID, &abm.BrandID, &abm.FactoryBrandID, &abm.Status, &abm.CreatedAt, &abm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &abm)
	}

	return res, nil
}

// ActivityBrandMapping20200417sByFactoryActivityID retrieves a row from 'ddg_local.activity_brand_mapping_20200417' as a ActivityBrandMapping20200417.
//
// Generated from index 'activity_brand_mapping_factory_activity_id_index'.
func ActivityBrandMapping20200417sByFactoryActivityID(ctx context.Context, factoryActivityID int, key ...interface{}) ([]*ActivityBrandMapping20200417, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityBrandMapping20200417TableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, brand_id, factory_brand_id, status, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE factory_activity_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, factoryActivityID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, factoryActivityID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, factoryActivityID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*ActivityBrandMapping20200417, 0)
	for queryData.Next() {
		abm := ActivityBrandMapping20200417{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&abm.ID, &abm.ActivityID, &abm.FactoryActivityID, &abm.BrandID, &abm.FactoryBrandID, &abm.Status, &abm.CreatedAt, &abm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &abm)
	}

	return res, nil
}

// ActivityBrandMapping20200417sByStatus retrieves a row from 'ddg_local.activity_brand_mapping_20200417' as a ActivityBrandMapping20200417.
//
// Generated from index 'activity_brand_mapping_status_index'.
func ActivityBrandMapping20200417sByStatus(ctx context.Context, status int8, key ...interface{}) ([]*ActivityBrandMapping20200417, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityBrandMapping20200417TableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, brand_id, factory_brand_id, status, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE status = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, status)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*ActivityBrandMapping20200417, 0)
	for queryData.Next() {
		abm := ActivityBrandMapping20200417{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&abm.ID, &abm.ActivityID, &abm.FactoryActivityID, &abm.BrandID, &abm.FactoryBrandID, &abm.Status, &abm.CreatedAt, &abm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &abm)
	}

	return res, nil
}
