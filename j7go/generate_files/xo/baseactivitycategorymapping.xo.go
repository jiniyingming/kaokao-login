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

// BaseActivityCategoryMapping represents a row from 'aypcddg.base_activity_category_mapping'.
type BaseActivityCategoryMapping struct {
	ID         int           `json:"id"`          // id
	CategoryID sql.NullInt64 `json:"category_id"` // category_id
	ActivityID sql.NullInt64 `json:"activity_id"` // activity_id
	Class1     sql.NullInt64 `json:"class1"`      // class1
	Class2     sql.NullInt64 `json:"class2"`      // class2
	Class3     sql.NullInt64 `json:"class3"`      // class3
	GoodsNum   sql.NullInt64 `json:"goods_num"`   // goods_num
	Created    sql.NullInt64 `json:"created"`     // created
	Updated    sql.NullInt64 `json:"updated"`     // updated
	Status     sql.NullInt64 `json:"status"`      // status
	ShowType   int8          `json:"show_type"`   // show_type

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the BaseActivityCategoryMapping exists in the database.
func (bacm *BaseActivityCategoryMapping) Exists() bool { //base_activity_category_mapping
	return bacm._exists
}

// Deleted provides information if the BaseActivityCategoryMapping has been deleted from the database.
func (bacm *BaseActivityCategoryMapping) Deleted() bool {
	return bacm._deleted
}

// Get table name
func GetBaseActivityCategoryMappingTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "base_activity_category_mapping", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the BaseActivityCategoryMapping to the database.
func (bacm *BaseActivityCategoryMapping) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if bacm._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBaseActivityCategoryMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`category_id, activity_id, class1, class2, class3, goods_num, created, updated, status, show_type` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bacm.CategoryID, bacm.ActivityID, bacm.Class1, bacm.Class2, bacm.Class3, bacm.GoodsNum, bacm.Created, bacm.Updated, bacm.Status, bacm.ShowType)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, bacm.CategoryID, bacm.ActivityID, bacm.Class1, bacm.Class2, bacm.Class3, bacm.GoodsNum, bacm.Created, bacm.Updated, bacm.Status, bacm.ShowType)
	} else {
		res, err = dbConn.Exec(sqlstr, bacm.CategoryID, bacm.ActivityID, bacm.Class1, bacm.Class2, bacm.Class3, bacm.GoodsNum, bacm.Created, bacm.Updated, bacm.Status, bacm.ShowType)
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
	bacm.ID = int(id)
	bacm._exists = true

	return nil
}

// Update updates the BaseActivityCategoryMapping in the database.
func (bacm *BaseActivityCategoryMapping) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if bacm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBaseActivityCategoryMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`category_id = ?, activity_id = ?, class1 = ?, class2 = ?, class3 = ?, goods_num = ?, created = ?, updated = ?, status = ?, show_type = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bacm.CategoryID, bacm.ActivityID, bacm.Class1, bacm.Class2, bacm.Class3, bacm.GoodsNum, bacm.Created, bacm.Updated, bacm.Status, bacm.ShowType, bacm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, bacm.CategoryID, bacm.ActivityID, bacm.Class1, bacm.Class2, bacm.Class3, bacm.GoodsNum, bacm.Created, bacm.Updated, bacm.Status, bacm.ShowType, bacm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, bacm.CategoryID, bacm.ActivityID, bacm.Class1, bacm.Class2, bacm.Class3, bacm.GoodsNum, bacm.Created, bacm.Updated, bacm.Status, bacm.ShowType, bacm.ID)
	}
	return err
}

// Save saves the BaseActivityCategoryMapping to the database.
func (bacm *BaseActivityCategoryMapping) Save(ctx context.Context) error {
	if bacm.Exists() {
		return bacm.Update(ctx)
	}

	return bacm.Insert(ctx)
}

// Delete deletes the BaseActivityCategoryMapping from the database.
func (bacm *BaseActivityCategoryMapping) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if bacm._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBaseActivityCategoryMappingTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bacm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, bacm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, bacm.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	bacm._deleted = true

	return nil
}

// BaseActivityCategoryMappingByID retrieves a row from 'aypcddg.base_activity_category_mapping' as a BaseActivityCategoryMapping.
//
// Generated from index 'base_activity_category_mapping_id_pkey'.
func BaseActivityCategoryMappingByID(ctx context.Context, id int, key ...interface{}) (*BaseActivityCategoryMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBaseActivityCategoryMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, category_id, activity_id, class1, class2, class3, goods_num, created, updated, status, show_type ` +
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
	bacm := BaseActivityCategoryMapping{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&bacm.ID, &bacm.CategoryID, &bacm.ActivityID, &bacm.Class1, &bacm.Class2, &bacm.Class3, &bacm.GoodsNum, &bacm.Created, &bacm.Updated, &bacm.Status, &bacm.ShowType)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&bacm.ID, &bacm.CategoryID, &bacm.ActivityID, &bacm.Class1, &bacm.Class2, &bacm.Class3, &bacm.GoodsNum, &bacm.Created, &bacm.Updated, &bacm.Status, &bacm.ShowType)
		if err != nil {
			return nil, err
		}
	}

	return &bacm, nil
}

// BaseActivityCategoryMappingsByStatusCategoryID retrieves a row from 'aypcddg.base_activity_category_mapping' as a BaseActivityCategoryMapping.
//
// Generated from index 'idx_category_id'.
func BaseActivityCategoryMappingsByStatusCategoryID(ctx context.Context, status sql.NullInt64, categoryID sql.NullInt64, key ...interface{}) ([]*BaseActivityCategoryMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBaseActivityCategoryMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, category_id, activity_id, class1, class2, class3, goods_num, created, updated, status, show_type ` +
		`FROM ` + tableName +
		` WHERE status = ? AND category_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, status, categoryID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, status, categoryID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, status, categoryID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*BaseActivityCategoryMapping, 0)
	for queryData.Next() {
		bacm := BaseActivityCategoryMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&bacm.ID, &bacm.CategoryID, &bacm.ActivityID, &bacm.Class1, &bacm.Class2, &bacm.Class3, &bacm.GoodsNum, &bacm.Created, &bacm.Updated, &bacm.Status, &bacm.ShowType)
		if err != nil {
			return nil, err
		}

		res = append(res, &bacm)
	}

	return res, nil
}

// BaseActivityCategoryMappingsByClass1ActivityID retrieves a row from 'aypcddg.base_activity_category_mapping' as a BaseActivityCategoryMapping.
//
// Generated from index 'idx_class1_activity_id'.
func BaseActivityCategoryMappingsByClass1ActivityID(ctx context.Context, class1 sql.NullInt64, activityID sql.NullInt64, key ...interface{}) ([]*BaseActivityCategoryMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBaseActivityCategoryMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, category_id, activity_id, class1, class2, class3, goods_num, created, updated, status, show_type ` +
		`FROM ` + tableName +
		` WHERE class1 = ? AND activity_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, class1, activityID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, class1, activityID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, class1, activityID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*BaseActivityCategoryMapping, 0)
	for queryData.Next() {
		bacm := BaseActivityCategoryMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&bacm.ID, &bacm.CategoryID, &bacm.ActivityID, &bacm.Class1, &bacm.Class2, &bacm.Class3, &bacm.GoodsNum, &bacm.Created, &bacm.Updated, &bacm.Status, &bacm.ShowType)
		if err != nil {
			return nil, err
		}

		res = append(res, &bacm)
	}

	return res, nil
}

// BaseActivityCategoryMappingsByShowType retrieves a row from 'aypcddg.base_activity_category_mapping' as a BaseActivityCategoryMapping.
//
// Generated from index 'idx_show_type'.
func BaseActivityCategoryMappingsByShowType(ctx context.Context, showType int8, key ...interface{}) ([]*BaseActivityCategoryMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBaseActivityCategoryMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, category_id, activity_id, class1, class2, class3, goods_num, created, updated, status, show_type ` +
		`FROM ` + tableName +
		` WHERE show_type = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, showType)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, showType)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, showType)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*BaseActivityCategoryMapping, 0)
	for queryData.Next() {
		bacm := BaseActivityCategoryMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&bacm.ID, &bacm.CategoryID, &bacm.ActivityID, &bacm.Class1, &bacm.Class2, &bacm.Class3, &bacm.GoodsNum, &bacm.Created, &bacm.Updated, &bacm.Status, &bacm.ShowType)
		if err != nil {
			return nil, err
		}

		res = append(res, &bacm)
	}

	return res, nil
}

// BaseActivityCategoryMappingsByStatusClass1ActivityID retrieves a row from 'aypcddg.base_activity_category_mapping' as a BaseActivityCategoryMapping.
//
// Generated from index 'idx_status'.
func BaseActivityCategoryMappingsByStatusClass1ActivityID(ctx context.Context, status sql.NullInt64, class1 sql.NullInt64, activityID sql.NullInt64, key ...interface{}) ([]*BaseActivityCategoryMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBaseActivityCategoryMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, category_id, activity_id, class1, class2, class3, goods_num, created, updated, status, show_type ` +
		`FROM ` + tableName +
		` WHERE status = ? AND class1 = ? AND activity_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, status, class1, activityID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, status, class1, activityID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, status, class1, activityID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*BaseActivityCategoryMapping, 0)
	for queryData.Next() {
		bacm := BaseActivityCategoryMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&bacm.ID, &bacm.CategoryID, &bacm.ActivityID, &bacm.Class1, &bacm.Class2, &bacm.Class3, &bacm.GoodsNum, &bacm.Created, &bacm.Updated, &bacm.Status, &bacm.ShowType)
		if err != nil {
			return nil, err
		}

		res = append(res, &bacm)
	}

	return res, nil
}
