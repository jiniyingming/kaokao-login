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

// BaseActivityCustomize represents a row from 'aypcddg.base_activity_customize'.
type BaseActivityCustomize struct {
	ID                 int            `json:"id"`                   // id
	BaseActivityID     sql.NullInt64  `json:"base_activity_id"`     // base_activity_id
	Title              sql.NullString `json:"title"`                // title
	Subtitle           sql.NullString `json:"subtitle"`             // subtitle
	Images             sql.NullString `json:"images"`               // images
	Tags               sql.NullString `json:"tags"`                 // tags
	Sort               sql.NullInt64  `json:"sort"`                 // sort
	HeaderImg          sql.NullString `json:"header_img"`           // header_img
	Status             sql.NullInt64  `json:"status"`               // status
	Desc               sql.NullString `json:"desc"`                 // desc
	Logo               sql.NullString `json:"logo"`                 // logo
	PlaceID            sql.NullInt64  `json:"place_id"`             // place_id
	Created            sql.NullInt64  `json:"created"`              // created
	Updated            sql.NullInt64  `json:"updated"`              // updated
	ActivityScheduleID sql.NullInt64  `json:"activity_schedule_id"` // activity_schedule_id
	SetTopDatetime     mysql.NullTime `json:"set_top_datetime"`     // set_top_datetime
	IsTop              sql.NullInt64  `json:"is_top"`               // is_top

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the BaseActivityCustomize exists in the database.
func (bac *BaseActivityCustomize) Exists() bool { //base_activity_customize
	return bac._exists
}

// Deleted provides information if the BaseActivityCustomize has been deleted from the database.
func (bac *BaseActivityCustomize) Deleted() bool {
	return bac._deleted
}

// Get table name
func GetBaseActivityCustomizeTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "base_activity_customize", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the BaseActivityCustomize to the database.
func (bac *BaseActivityCustomize) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if bac._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBaseActivityCustomizeTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`base_activity_id, title, subtitle, images, tags, sort, header_img, status, desc, logo, place_id, created, updated, activity_schedule_id, set_top_datetime, is_top` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bac.BaseActivityID, bac.Title, bac.Subtitle, bac.Images, bac.Tags, bac.Sort, bac.HeaderImg, bac.Status, bac.Desc, bac.Logo, bac.PlaceID, bac.Created, bac.Updated, bac.ActivityScheduleID, bac.SetTopDatetime, bac.IsTop)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, bac.BaseActivityID, bac.Title, bac.Subtitle, bac.Images, bac.Tags, bac.Sort, bac.HeaderImg, bac.Status, bac.Desc, bac.Logo, bac.PlaceID, bac.Created, bac.Updated, bac.ActivityScheduleID, bac.SetTopDatetime, bac.IsTop)
	} else {
		res, err = dbConn.Exec(sqlstr, bac.BaseActivityID, bac.Title, bac.Subtitle, bac.Images, bac.Tags, bac.Sort, bac.HeaderImg, bac.Status, bac.Desc, bac.Logo, bac.PlaceID, bac.Created, bac.Updated, bac.ActivityScheduleID, bac.SetTopDatetime, bac.IsTop)
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
	bac.ID = int(id)
	bac._exists = true

	return nil
}

// Update updates the BaseActivityCustomize in the database.
func (bac *BaseActivityCustomize) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if bac._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBaseActivityCustomizeTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`base_activity_id = ?, title = ?, subtitle = ?, images = ?, tags = ?, sort = ?, header_img = ?, status = ?, desc = ?, logo = ?, place_id = ?, created = ?, updated = ?, activity_schedule_id = ?, set_top_datetime = ?, is_top = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bac.BaseActivityID, bac.Title, bac.Subtitle, bac.Images, bac.Tags, bac.Sort, bac.HeaderImg, bac.Status, bac.Desc, bac.Logo, bac.PlaceID, bac.Created, bac.Updated, bac.ActivityScheduleID, bac.SetTopDatetime, bac.IsTop, bac.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, bac.BaseActivityID, bac.Title, bac.Subtitle, bac.Images, bac.Tags, bac.Sort, bac.HeaderImg, bac.Status, bac.Desc, bac.Logo, bac.PlaceID, bac.Created, bac.Updated, bac.ActivityScheduleID, bac.SetTopDatetime, bac.IsTop, bac.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, bac.BaseActivityID, bac.Title, bac.Subtitle, bac.Images, bac.Tags, bac.Sort, bac.HeaderImg, bac.Status, bac.Desc, bac.Logo, bac.PlaceID, bac.Created, bac.Updated, bac.ActivityScheduleID, bac.SetTopDatetime, bac.IsTop, bac.ID)
	}
	return err
}

// Save saves the BaseActivityCustomize to the database.
func (bac *BaseActivityCustomize) Save(ctx context.Context) error {
	if bac.Exists() {
		return bac.Update(ctx)
	}

	return bac.Insert(ctx)
}

// Delete deletes the BaseActivityCustomize from the database.
func (bac *BaseActivityCustomize) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if bac._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetBaseActivityCustomizeTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, bac.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, bac.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, bac.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	bac._deleted = true

	return nil
}

// BaseActivityCustomizeByID retrieves a row from 'aypcddg.base_activity_customize' as a BaseActivityCustomize.
//
// Generated from index 'base_activity_customize_id_pkey'.
func BaseActivityCustomizeByID(ctx context.Context, id int, key ...interface{}) (*BaseActivityCustomize, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBaseActivityCustomizeTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, base_activity_id, title, subtitle, images, tags, sort, header_img, status, desc, logo, place_id, created, updated, activity_schedule_id, set_top_datetime, is_top ` +
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
	bac := BaseActivityCustomize{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&bac.ID, &bac.BaseActivityID, &bac.Title, &bac.Subtitle, &bac.Images, &bac.Tags, &bac.Sort, &bac.HeaderImg, &bac.Status, &bac.Desc, &bac.Logo, &bac.PlaceID, &bac.Created, &bac.Updated, &bac.ActivityScheduleID, &bac.SetTopDatetime, &bac.IsTop)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&bac.ID, &bac.BaseActivityID, &bac.Title, &bac.Subtitle, &bac.Images, &bac.Tags, &bac.Sort, &bac.HeaderImg, &bac.Status, &bac.Desc, &bac.Logo, &bac.PlaceID, &bac.Created, &bac.Updated, &bac.ActivityScheduleID, &bac.SetTopDatetime, &bac.IsTop)
		if err != nil {
			return nil, err
		}
	}

	return &bac, nil
}

// BaseActivityCustomizesByBaseActivityIDSort retrieves a row from 'aypcddg.base_activity_customize' as a BaseActivityCustomize.
//
// Generated from index 'idx_sort'.
func BaseActivityCustomizesByBaseActivityIDSort(ctx context.Context, baseActivityID sql.NullInt64, sort sql.NullInt64, key ...interface{}) ([]*BaseActivityCustomize, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBaseActivityCustomizeTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, base_activity_id, title, subtitle, images, tags, sort, header_img, status, desc, logo, place_id, created, updated, activity_schedule_id, set_top_datetime, is_top ` +
		`FROM ` + tableName +
		` WHERE base_activity_id = ? AND sort = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, baseActivityID, sort)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, baseActivityID, sort)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, baseActivityID, sort)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*BaseActivityCustomize, 0)
	for queryData.Next() {
		bac := BaseActivityCustomize{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&bac.ID, &bac.BaseActivityID, &bac.Title, &bac.Subtitle, &bac.Images, &bac.Tags, &bac.Sort, &bac.HeaderImg, &bac.Status, &bac.Desc, &bac.Logo, &bac.PlaceID, &bac.Created, &bac.Updated, &bac.ActivityScheduleID, &bac.SetTopDatetime, &bac.IsTop)
		if err != nil {
			return nil, err
		}

		res = append(res, &bac)
	}

	return res, nil
}

// BaseActivityCustomizesByStatus retrieves a row from 'aypcddg.base_activity_customize' as a BaseActivityCustomize.
//
// Generated from index 'idx_status'.
func BaseActivityCustomizesByStatus(ctx context.Context, status sql.NullInt64, key ...interface{}) ([]*BaseActivityCustomize, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetBaseActivityCustomizeTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, base_activity_id, title, subtitle, images, tags, sort, header_img, status, desc, logo, place_id, created, updated, activity_schedule_id, set_top_datetime, is_top ` +
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
	res := make([]*BaseActivityCustomize, 0)
	for queryData.Next() {
		bac := BaseActivityCustomize{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&bac.ID, &bac.BaseActivityID, &bac.Title, &bac.Subtitle, &bac.Images, &bac.Tags, &bac.Sort, &bac.HeaderImg, &bac.Status, &bac.Desc, &bac.Logo, &bac.PlaceID, &bac.Created, &bac.Updated, &bac.ActivityScheduleID, &bac.SetTopDatetime, &bac.IsTop)
		if err != nil {
			return nil, err
		}

		res = append(res, &bac)
	}

	return res, nil
}