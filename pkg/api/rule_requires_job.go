package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/netip"
	"slices"
	"strings"
	"time"

	"github.com/cridenour/go-postgis"
	"github.com/go-chi/chi/v5"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/initialed85/djangolang/pkg/config"
	"github.com/initialed85/djangolang/pkg/helpers"
	"github.com/initialed85/djangolang/pkg/introspect"
	"github.com/initialed85/djangolang/pkg/query"
	"github.com/initialed85/djangolang/pkg/server"
	"github.com/initialed85/djangolang/pkg/stream"
	"github.com/initialed85/djangolang/pkg/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/maps"
)

type RuleRequiresJob struct {
	ID           uuid.UUID  `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	RuleID       uuid.UUID  `json:"rule_id"`
	RuleIDObject *Rule      `json:"rule_id_object"`
	JobID        uuid.UUID  `json:"job_id"`
	JobIDObject  *Job       `json:"job_id_object"`
}

var RuleRequiresJobTable = "rule_requires_job"

var RuleRequiresJobTableNamespaceID int32 = 1337 + 7

var (
	RuleRequiresJobTableIDColumn        = "id"
	RuleRequiresJobTableCreatedAtColumn = "created_at"
	RuleRequiresJobTableUpdatedAtColumn = "updated_at"
	RuleRequiresJobTableDeletedAtColumn = "deleted_at"
	RuleRequiresJobTableRuleIDColumn    = "rule_id"
	RuleRequiresJobTableJobIDColumn     = "job_id"
)

var (
	RuleRequiresJobTableIDColumnWithTypeCast        = `"id" AS id`
	RuleRequiresJobTableCreatedAtColumnWithTypeCast = `"created_at" AS created_at`
	RuleRequiresJobTableUpdatedAtColumnWithTypeCast = `"updated_at" AS updated_at`
	RuleRequiresJobTableDeletedAtColumnWithTypeCast = `"deleted_at" AS deleted_at`
	RuleRequiresJobTableRuleIDColumnWithTypeCast    = `"rule_id" AS rule_id`
	RuleRequiresJobTableJobIDColumnWithTypeCast     = `"job_id" AS job_id`
)

var RuleRequiresJobTableColumns = []string{
	RuleRequiresJobTableIDColumn,
	RuleRequiresJobTableCreatedAtColumn,
	RuleRequiresJobTableUpdatedAtColumn,
	RuleRequiresJobTableDeletedAtColumn,
	RuleRequiresJobTableRuleIDColumn,
	RuleRequiresJobTableJobIDColumn,
}

var RuleRequiresJobTableColumnsWithTypeCasts = []string{
	RuleRequiresJobTableIDColumnWithTypeCast,
	RuleRequiresJobTableCreatedAtColumnWithTypeCast,
	RuleRequiresJobTableUpdatedAtColumnWithTypeCast,
	RuleRequiresJobTableDeletedAtColumnWithTypeCast,
	RuleRequiresJobTableRuleIDColumnWithTypeCast,
	RuleRequiresJobTableJobIDColumnWithTypeCast,
}

var RuleRequiresJobIntrospectedTable *introspect.Table

var RuleRequiresJobTableColumnLookup map[string]*introspect.Column

var (
	RuleRequiresJobTablePrimaryKeyColumn = RuleRequiresJobTableIDColumn
)

func init() {
	RuleRequiresJobIntrospectedTable = tableByName[RuleRequiresJobTable]

	/* only needed during templating */
	if RuleRequiresJobIntrospectedTable == nil {
		RuleRequiresJobIntrospectedTable = &introspect.Table{}
	}

	RuleRequiresJobTableColumnLookup = RuleRequiresJobIntrospectedTable.ColumnByName
}

type RuleRequiresJobOnePathParams struct {
	PrimaryKey uuid.UUID `json:"primaryKey"`
}

type RuleRequiresJobLoadQueryParams struct {
	Depth *int `json:"depth"`
}

/*
TODO: find a way to not need this- there is a piece in the templating logic
that uses goimports but pending where the code is built, it may resolve
the packages to import to the wrong ones (causing odd failures)
these are just here to ensure we don't get unused imports
*/
var _ = []any{
	time.Time{},
	uuid.UUID{},
	pgtype.Hstore{},
	postgis.PointZ{},
	netip.Prefix{},
	errors.Is,
	sql.ErrNoRows,
}

func (m *RuleRequiresJob) GetPrimaryKeyColumn() string {
	return RuleRequiresJobTablePrimaryKeyColumn
}

func (m *RuleRequiresJob) GetPrimaryKeyValue() any {
	return m.ID
}

func (m *RuleRequiresJob) FromItem(item map[string]any) error {
	if item == nil {
		return fmt.Errorf(
			"item unexpectedly nil during RuleRequiresJobFromItem",
		)
	}

	if len(item) == 0 {
		return fmt.Errorf(
			"item unexpectedly empty during RuleRequiresJobFromItem",
		)
	}

	wrapError := func(k string, v any, err error) error {
		return fmt.Errorf("%v: %#+v; error; %v", k, v, err)
	}

	for k, v := range item {
		_, ok := RuleRequiresJobTableColumnLookup[k]
		if !ok {
			return fmt.Errorf(
				"item contained unexpected key %#+v during RuleRequiresJobFromItem; item: %#+v",
				k, item,
			)
		}

		switch k {
		case "id":
			if v == nil {
				continue
			}

			temp1, err := types.ParseUUID(v)
			if err != nil {
				return wrapError(k, v, err)
			}

			temp2, ok := temp1.(uuid.UUID)
			if !ok {
				if temp1 != nil {
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uuid.UUID", temp1))
				}
			}

			m.ID = temp2

		case "created_at":
			if v == nil {
				continue
			}

			temp1, err := types.ParseTime(v)
			if err != nil {
				return wrapError(k, v, err)
			}

			temp2, ok := temp1.(time.Time)
			if !ok {
				if temp1 != nil {
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uucreated_at.UUID", temp1))
				}
			}

			m.CreatedAt = temp2

		case "updated_at":
			if v == nil {
				continue
			}

			temp1, err := types.ParseTime(v)
			if err != nil {
				return wrapError(k, v, err)
			}

			temp2, ok := temp1.(time.Time)
			if !ok {
				if temp1 != nil {
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uuupdated_at.UUID", temp1))
				}
			}

			m.UpdatedAt = temp2

		case "deleted_at":
			if v == nil {
				continue
			}

			temp1, err := types.ParseTime(v)
			if err != nil {
				return wrapError(k, v, err)
			}

			temp2, ok := temp1.(time.Time)
			if !ok {
				if temp1 != nil {
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uudeleted_at.UUID", temp1))
				}
			}

			m.DeletedAt = &temp2

		case "rule_id":
			if v == nil {
				continue
			}

			temp1, err := types.ParseUUID(v)
			if err != nil {
				return wrapError(k, v, err)
			}

			temp2, ok := temp1.(uuid.UUID)
			if !ok {
				if temp1 != nil {
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uurule_id.UUID", temp1))
				}
			}

			m.RuleID = temp2

		case "job_id":
			if v == nil {
				continue
			}

			temp1, err := types.ParseUUID(v)
			if err != nil {
				return wrapError(k, v, err)
			}

			temp2, ok := temp1.(uuid.UUID)
			if !ok {
				if temp1 != nil {
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uujob_id.UUID", temp1))
				}
			}

			m.JobID = temp2

		}
	}

	return nil
}

func (m *RuleRequiresJob) Reload(ctx context.Context, tx pgx.Tx, includeDeleteds ...bool) error {
	extraWhere := ""
	if len(includeDeleteds) > 0 && includeDeleteds[0] {
		if slices.Contains(RuleRequiresJobTableColumns, "deleted_at") {
			extraWhere = "\n    AND (deleted_at IS null OR deleted_at IS NOT null)"
		}
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	o, _, _, _, _, err := SelectRuleRequiresJob(
		ctx,
		tx,
		fmt.Sprintf("%v = $1%v", m.GetPrimaryKeyColumn(), extraWhere),
		m.GetPrimaryKeyValue(),
	)
	if err != nil {
		return err
	}

	m.ID = o.ID
	m.CreatedAt = o.CreatedAt
	m.UpdatedAt = o.UpdatedAt
	m.DeletedAt = o.DeletedAt
	m.RuleID = o.RuleID
	m.RuleIDObject = o.RuleIDObject
	m.JobID = o.JobID
	m.JobIDObject = o.JobIDObject

	return nil
}

func (m *RuleRequiresJob) Insert(ctx context.Context, tx pgx.Tx, setPrimaryKey bool, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setPrimaryKey && (setZeroValues || !types.IsZeroUUID(m.ID) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableIDColumn) || isRequired(RuleRequiresJobTableColumnLookup, RuleRequiresJobTableIDColumn)) {
		columns = append(columns, RuleRequiresJobTableIDColumn)

		v, err := types.FormatUUID(m.ID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableCreatedAtColumn) || isRequired(RuleRequiresJobTableColumnLookup, RuleRequiresJobTableCreatedAtColumn) {
		columns = append(columns, RuleRequiresJobTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableUpdatedAtColumn) || isRequired(RuleRequiresJobTableColumnLookup, RuleRequiresJobTableUpdatedAtColumn) {
		columns = append(columns, RuleRequiresJobTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableDeletedAtColumn) || isRequired(RuleRequiresJobTableColumnLookup, RuleRequiresJobTableDeletedAtColumn) {
		columns = append(columns, RuleRequiresJobTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.RuleID) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableRuleIDColumn) || isRequired(RuleRequiresJobTableColumnLookup, RuleRequiresJobTableRuleIDColumn) {
		columns = append(columns, RuleRequiresJobTableRuleIDColumn)

		v, err := types.FormatUUID(m.RuleID)
		if err != nil {
			return fmt.Errorf("failed to handle m.RuleID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.JobID) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableJobIDColumn) || isRequired(RuleRequiresJobTableColumnLookup, RuleRequiresJobTableJobIDColumn) {
		columns = append(columns, RuleRequiresJobTableJobIDColumn)

		v, err := types.FormatUUID(m.JobID)
		if err != nil {
			return fmt.Errorf("failed to handle m.JobID; %v", err)
		}

		values = append(values, v)
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	item, err := query.Insert(
		ctx,
		tx,
		RuleRequiresJobTable,
		columns,
		nil,
		false,
		false,
		RuleRequiresJobTableColumns,
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to insert %#+v; %v", m, err)
	}
	v := (*item)[RuleRequiresJobTableIDColumn]

	if v == nil {
		return fmt.Errorf("failed to find %v in %#+v", RuleRequiresJobTableIDColumn, item)
	}

	wrapError := func(err error) error {
		return fmt.Errorf(
			"failed to treat %v: %#+v as uuid.UUID: %v",
			RuleRequiresJobTableIDColumn,
			(*item)[RuleRequiresJobTableIDColumn],
			err,
		)
	}

	temp1, err := types.ParseUUID(v)
	if err != nil {
		return wrapError(err)
	}

	temp2, ok := temp1.(uuid.UUID)
	if !ok {
		return wrapError(fmt.Errorf("failed to cast to uuid.UUID"))
	}

	m.ID = temp2

	err = m.Reload(ctx, tx, slices.Contains(forceSetValuesForFields, "deleted_at"))
	if err != nil {
		return fmt.Errorf("failed to reload after insert; %v", err)
	}

	return nil
}

func (m *RuleRequiresJob) Update(ctx context.Context, tx pgx.Tx, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableCreatedAtColumn) {
		columns = append(columns, RuleRequiresJobTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableUpdatedAtColumn) {
		columns = append(columns, RuleRequiresJobTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableDeletedAtColumn) {
		columns = append(columns, RuleRequiresJobTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.RuleID) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableRuleIDColumn) {
		columns = append(columns, RuleRequiresJobTableRuleIDColumn)

		v, err := types.FormatUUID(m.RuleID)
		if err != nil {
			return fmt.Errorf("failed to handle m.RuleID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.JobID) || slices.Contains(forceSetValuesForFields, RuleRequiresJobTableJobIDColumn) {
		columns = append(columns, RuleRequiresJobTableJobIDColumn)

		v, err := types.FormatUUID(m.JobID)
		if err != nil {
			return fmt.Errorf("failed to handle m.JobID; %v", err)
		}

		values = append(values, v)
	}

	v, err := types.FormatUUID(m.ID)
	if err != nil {
		return fmt.Errorf("failed to handle m.ID; %v", err)
	}

	values = append(values, v)

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	_, err = query.Update(
		ctx,
		tx,
		RuleRequiresJobTable,
		columns,
		fmt.Sprintf("%v = $$??", RuleRequiresJobTableIDColumn),
		RuleRequiresJobTableColumns,
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to update %#+v; %v", m, err)
	}

	err = m.Reload(ctx, tx, slices.Contains(forceSetValuesForFields, "deleted_at"))
	if err != nil {
		return fmt.Errorf("failed to reload after update")
	}

	return nil
}

func (m *RuleRequiresJob) Delete(ctx context.Context, tx pgx.Tx, hardDeletes ...bool) error {
	hardDelete := false
	if len(hardDeletes) > 0 {
		hardDelete = hardDeletes[0]
	}

	if !hardDelete && slices.Contains(RuleRequiresJobTableColumns, "deleted_at") {
		m.DeletedAt = helpers.Ptr(time.Now().UTC())
		err := m.Update(ctx, tx, false, "deleted_at")
		if err != nil {
			return fmt.Errorf("failed to soft-delete (update) %#+v; %v", m, err)
		}
	}

	values := make([]any, 0)
	v, err := types.FormatUUID(m.ID)
	if err != nil {
		return fmt.Errorf("failed to handle m.ID; %v", err)
	}

	values = append(values, v)

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	err = query.Delete(
		ctx,
		tx,
		RuleRequiresJobTable,
		fmt.Sprintf("%v = $$??", RuleRequiresJobTableIDColumn),
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to delete %#+v; %v", m, err)
	}

	_ = m.Reload(ctx, tx, true)

	return nil
}

func (m *RuleRequiresJob) LockTable(ctx context.Context, tx pgx.Tx, timeouts ...time.Duration) error {
	return query.LockTable(ctx, tx, RuleRequiresJobTable, timeouts...)
}

func (m *RuleRequiresJob) LockTableWithRetries(ctx context.Context, tx pgx.Tx, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.LockTableWithRetries(ctx, tx, RuleRequiresJobTable, overallTimeout, individualAttempttimeout)
}

func (m *RuleRequiresJob) AdvisoryLock(ctx context.Context, tx pgx.Tx, key int32, timeouts ...time.Duration) error {
	return query.AdvisoryLock(ctx, tx, RuleRequiresJobTableNamespaceID, key, timeouts...)
}

func (m *RuleRequiresJob) AdvisoryLockWithRetries(ctx context.Context, tx pgx.Tx, key int32, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.AdvisoryLockWithRetries(ctx, tx, RuleRequiresJobTableNamespaceID, key, overallTimeout, individualAttempttimeout)
}

func SelectRuleRequiresJobs(ctx context.Context, tx pgx.Tx, where string, orderBy *string, limit *int, offset *int, values ...any) ([]*RuleRequiresJob, int64, int64, int64, int64, error) {
	before := time.Now()

	if config.Debug() {
		log.Printf("entered SelectRuleRequiresJobs")

		defer func() {
			log.Printf("exited SelectRuleRequiresJobs in %s", time.Since(before))
		}()
	}
	if slices.Contains(RuleRequiresJobTableColumns, "deleted_at") {
		if !strings.Contains(where, "deleted_at") {
			if where != "" {
				where += "\n    AND "
			}

			where += "deleted_at IS null"
		}
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	possiblePathValue := query.GetCurrentPathValue(ctx)
	isLoadQuery := possiblePathValue != nil && len(possiblePathValue.VisitedTableNames) > 0
	ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", RuleRequiresJobTable, nil), !isLoadQuery)
	if !ok {
		return []*RuleRequiresJob{}, 0, 0, 0, 0, nil
	}

	items, count, totalCount, page, totalPages, err := query.Select(
		ctx,
		tx,
		RuleRequiresJobTableColumnsWithTypeCasts,
		RuleRequiresJobTable,
		where,
		orderBy,
		limit,
		offset,
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectRuleRequiresJobs; %v", err)
	}

	objects := make([]*RuleRequiresJob, 0)

	for _, item := range *items {
		object := &RuleRequiresJob{}

		err = object.FromItem(item)
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		if !types.IsZeroUUID(object.RuleID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", RuleTable, object.RuleID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectRuleRequiresJobs->SelectRule for object.RuleIDObject")
				}

				object.RuleIDObject, _, _, _, _, err = SelectRule(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", RuleTablePrimaryKeyColumn),
					object.RuleID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectRuleRequiresJobs->SelectRule for object.RuleIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.JobID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", JobTable, object.JobID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectRuleRequiresJobs->SelectJob for object.JobIDObject")
				}

				object.JobIDObject, _, _, _, _, err = SelectJob(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", JobTablePrimaryKeyColumn),
					object.JobID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectRuleRequiresJobs->SelectJob for object.JobIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		objects = append(objects, object)
	}

	return objects, count, totalCount, page, totalPages, nil
}

func SelectRuleRequiresJob(ctx context.Context, tx pgx.Tx, where string, values ...any) (*RuleRequiresJob, int64, int64, int64, int64, error) {
	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	objects, _, _, _, _, err := SelectRuleRequiresJobs(
		ctx,
		tx,
		where,
		nil,
		helpers.Ptr(2),
		helpers.Ptr(0),
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectRuleRequiresJob; %v", err)
	}

	if len(objects) > 1 {
		return nil, 0, 0, 0, 0, fmt.Errorf("attempt to call SelectRuleRequiresJob returned more than 1 row")
	}

	if len(objects) < 1 {
		return nil, 0, 0, 0, 0, sql.ErrNoRows
	}

	object := objects[0]

	count := int64(1)
	totalCount := count
	page := int64(1)
	totalPages := page

	return object, count, totalCount, page, totalPages, nil
}

func handleGetRuleRequiresJobs(arguments *server.SelectManyArguments, db *pgxpool.Pool) ([]*RuleRequiresJob, int64, int64, int64, int64, error) {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		if config.Debug() {
			log.Printf("")
		}

		return nil, 0, 0, 0, 0, err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	objects, count, totalCount, page, totalPages, err := SelectRuleRequiresJobs(arguments.Ctx, tx, arguments.Where, arguments.OrderBy, arguments.Limit, arguments.Offset, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return objects, count, totalCount, page, totalPages, nil
}

func handleGetRuleRequiresJob(arguments *server.SelectOneArguments, db *pgxpool.Pool, primaryKey uuid.UUID) ([]*RuleRequiresJob, int64, int64, int64, int64, error) {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	object, count, totalCount, page, totalPages, err := SelectRuleRequiresJob(arguments.Ctx, tx, arguments.Where, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return []*RuleRequiresJob{object}, count, totalCount, page, totalPages, nil
}

func handlePostRuleRequiresJobs(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, objects []*RuleRequiresJob, forceSetValuesForFieldsByObjectIndex [][]string) ([]*RuleRequiresJob, int64, int64, int64, int64, error) {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		err = fmt.Errorf("failed to begin DB transaction; %v", err)
		return nil, 0, 0, 0, 0, err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	xid, err := query.GetXid(arguments.Ctx, tx)
	if err != nil {
		err = fmt.Errorf("failed to get xid; %v", err)
		return nil, 0, 0, 0, 0, err
	}
	_ = xid

	for i, object := range objects {
		err = object.Insert(arguments.Ctx, tx, false, false, forceSetValuesForFieldsByObjectIndex[i]...)
		if err != nil {
			err = fmt.Errorf("failed to insert %#+v; %v", object, err)
			return nil, 0, 0, 0, 0, err
		}

		objects[i] = object
	}

	errs := make(chan error, 1)
	go func() {
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.INSERT}, RuleRequiresJobTable, xid)
		if err != nil {
			err = fmt.Errorf("failed to wait for change; %v", err)
			errs <- err
			return
		}

		errs <- nil
	}()

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		err = fmt.Errorf("failed to commit DB transaction; %v", err)
		return nil, 0, 0, 0, 0, err
	}

	select {
	case <-arguments.Ctx.Done():
		err = fmt.Errorf("context canceled")
		return nil, 0, 0, 0, 0, err
	case err = <-errs:
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}
	}

	count := int64(len(objects))
	totalCount := count
	page := int64(1)
	totalPages := page

	return objects, count, totalCount, page, totalPages, nil
}

func handlePutRuleRequiresJob(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *RuleRequiresJob) ([]*RuleRequiresJob, int64, int64, int64, int64, error) {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		err = fmt.Errorf("failed to begin DB transaction; %v", err)
		return nil, 0, 0, 0, 0, err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	xid, err := query.GetXid(arguments.Ctx, tx)
	if err != nil {
		err = fmt.Errorf("failed to get xid; %v", err)
		return nil, 0, 0, 0, 0, err
	}
	_ = xid

	err = object.Update(arguments.Ctx, tx, true)
	if err != nil {
		err = fmt.Errorf("failed to update %#+v; %v", object, err)
		return nil, 0, 0, 0, 0, err
	}

	errs := make(chan error, 1)
	go func() {
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, RuleRequiresJobTable, xid)
		if err != nil {
			err = fmt.Errorf("failed to wait for change; %v", err)
			errs <- err
			return
		}

		errs <- nil
	}()

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		err = fmt.Errorf("failed to commit DB transaction; %v", err)
		return nil, 0, 0, 0, 0, err
	}

	select {
	case <-arguments.Ctx.Done():
		err = fmt.Errorf("context canceled")
		return nil, 0, 0, 0, 0, err
	case err = <-errs:
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}
	}

	count := int64(1)
	totalCount := count
	page := int64(1)
	totalPages := page

	return []*RuleRequiresJob{object}, count, totalCount, page, totalPages, nil
}

func handlePatchRuleRequiresJob(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *RuleRequiresJob, forceSetValuesForFields []string) ([]*RuleRequiresJob, int64, int64, int64, int64, error) {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		err = fmt.Errorf("failed to begin DB transaction; %v", err)
		return nil, 0, 0, 0, 0, err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	xid, err := query.GetXid(arguments.Ctx, tx)
	if err != nil {
		err = fmt.Errorf("failed to get xid; %v", err)
		return nil, 0, 0, 0, 0, err
	}
	_ = xid

	err = object.Update(arguments.Ctx, tx, false, forceSetValuesForFields...)
	if err != nil {
		err = fmt.Errorf("failed to update %#+v; %v", object, err)
		return nil, 0, 0, 0, 0, err
	}

	errs := make(chan error, 1)
	go func() {
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, RuleRequiresJobTable, xid)
		if err != nil {
			err = fmt.Errorf("failed to wait for change; %v", err)
			errs <- err
			return
		}

		errs <- nil
	}()

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		err = fmt.Errorf("failed to commit DB transaction; %v", err)
		return nil, 0, 0, 0, 0, err
	}

	select {
	case <-arguments.Ctx.Done():
		err = fmt.Errorf("context canceled")
		return nil, 0, 0, 0, 0, err
	case err = <-errs:
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}
	}

	count := int64(1)
	totalCount := count
	page := int64(1)
	totalPages := page

	return []*RuleRequiresJob{object}, count, totalCount, page, totalPages, nil
}

func handleDeleteRuleRequiresJob(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *RuleRequiresJob) error {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		err = fmt.Errorf("failed to begin DB transaction; %v", err)
		return err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	xid, err := query.GetXid(arguments.Ctx, tx)
	if err != nil {
		err = fmt.Errorf("failed to get xid; %v", err)
		return err
	}
	_ = xid

	err = object.Delete(arguments.Ctx, tx)
	if err != nil {
		err = fmt.Errorf("failed to delete %#+v; %v", object, err)
		return err
	}

	errs := make(chan error, 1)
	go func() {
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.DELETE, stream.SOFT_DELETE}, RuleRequiresJobTable, xid)
		if err != nil {
			err = fmt.Errorf("failed to wait for change; %v", err)
			errs <- err
			return
		}

		errs <- nil
	}()

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		err = fmt.Errorf("failed to commit DB transaction; %v", err)
		return err
	}

	select {
	case <-arguments.Ctx.Done():
		err = fmt.Errorf("context canceled")
		return err
	case err = <-errs:
		if err != nil {
			return err
		}
	}

	return nil
}

func GetRuleRequiresJobRouter(db *pgxpool.Pool, redisPool *redis.Pool, httpMiddlewares []server.HTTPMiddleware, objectMiddlewares []server.ObjectMiddleware, waitForChange server.WaitForChange) chi.Router {
	r := chi.NewRouter()

	for _, m := range httpMiddlewares {
		r.Use(m)
	}

	func() {
		getManyHandler, err := getHTTPHandler(
			http.MethodGet,
			"/rule-requires-jobs",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams map[string]any,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[RuleRequiresJob], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectManyArguments(ctx, queryParams, RuleRequiresJobIntrospectedTable, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[RuleRequiresJob]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[RuleRequiresJob]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[RuleRequiresJob]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[RuleRequiresJob]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetRuleRequiresJobs(arguments, db)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[RuleRequiresJob]{}, err
				}

				limit := int64(0)
				if arguments.Limit != nil {
					limit = int64(*arguments.Limit)
				}

				offset := int64(0)
				if arguments.Offset != nil {
					offset = int64(*arguments.Offset)
				}

				response := server.Response[RuleRequiresJob]{
					Status:     http.StatusOK,
					Success:    true,
					Error:      nil,
					Objects:    objects,
					Count:      count,
					TotalCount: totalCount,
					Limit:      limit,
					Offset:     offset,
				}

				/* TODO: it'd be nice to be able to avoid this (i.e. just marshal once, further out) */
				responseAsJSON, err := json.Marshal(response)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[RuleRequiresJob]{}, err
				}

				err = server.StoreCachedResponse(arguments.RequestHash, redisConn, responseAsJSON)
				if err != nil {
					log.Printf("warning; %v", err)
				}

				if config.Debug() {
					log.Printf("request cache missed; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
				}

				return response, nil
			},
			RuleRequiresJob{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getManyHandler.PathWithinRouter, getManyHandler.ServeHTTP)
	}()

	func() {
		getOneHandler, err := getHTTPHandler(
			http.MethodGet,
			"/rule-requires-jobs/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams RuleRequiresJobOnePathParams,
				queryParams RuleRequiresJobLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[RuleRequiresJob], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectOneArguments(ctx, queryParams.Depth, RuleRequiresJobIntrospectedTable, pathParams.PrimaryKey, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[RuleRequiresJob]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[RuleRequiresJob]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[RuleRequiresJob]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[RuleRequiresJob]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetRuleRequiresJob(arguments, db, pathParams.PrimaryKey)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[RuleRequiresJob]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				response := server.Response[RuleRequiresJob]{
					Status:     http.StatusOK,
					Success:    true,
					Error:      nil,
					Objects:    objects,
					Count:      count,
					TotalCount: totalCount,
					Limit:      limit,
					Offset:     offset,
				}

				/* TODO: it'd be nice to be able to avoid this (i.e. just marshal once, further out) */
				responseAsJSON, err := json.Marshal(response)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[RuleRequiresJob]{}, err
				}

				err = server.StoreCachedResponse(arguments.RequestHash, redisConn, responseAsJSON)
				if err != nil {
					log.Printf("warning; %v", err)
				}

				if config.Debug() {
					log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
				}

				return response, nil
			},
			RuleRequiresJob{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getOneHandler.PathWithinRouter, getOneHandler.ServeHTTP)
	}()

	func() {
		postHandler, err := getHTTPHandler(
			http.MethodPost,
			"/rule-requires-jobs",
			http.StatusCreated,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams RuleRequiresJobLoadQueryParams,
				req []*RuleRequiresJob,
				rawReq any,
			) (server.Response[RuleRequiresJob], error) {
				allRawItems, ok := rawReq.([]any)
				if !ok {
					return server.Response[RuleRequiresJob]{}, fmt.Errorf("failed to cast %#+v to []map[string]any", rawReq)
				}

				allItems := make([]map[string]any, 0)
				for _, rawItem := range allRawItems {
					item, ok := rawItem.(map[string]any)
					if !ok {
						return server.Response[RuleRequiresJob]{}, fmt.Errorf("failed to cast %#+v to map[string]any", rawItem)
					}

					allItems = append(allItems, item)
				}

				forceSetValuesForFieldsByObjectIndex := make([][]string, 0)
				for _, item := range allItems {
					forceSetValuesForFields := make([]string, 0)
					for _, possibleField := range maps.Keys(item) {
						if !slices.Contains(RuleRequiresJobTableColumns, possibleField) {
							continue
						}

						forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
					}
					forceSetValuesForFieldsByObjectIndex = append(forceSetValuesForFieldsByObjectIndex, forceSetValuesForFields)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[RuleRequiresJob]{}, err
				}

				objects, count, totalCount, _, _, err := handlePostRuleRequiresJobs(arguments, db, waitForChange, req, forceSetValuesForFieldsByObjectIndex)
				if err != nil {
					return server.Response[RuleRequiresJob]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[RuleRequiresJob]{
					Status:     http.StatusOK,
					Success:    true,
					Error:      nil,
					Objects:    objects,
					Count:      count,
					TotalCount: totalCount,
					Limit:      limit,
					Offset:     offset,
				}, nil
			},
			RuleRequiresJob{},
		)
		if err != nil {
			panic(err)
		}
		r.Post(postHandler.PathWithinRouter, postHandler.ServeHTTP)
	}()

	func() {
		putHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/rule-requires-jobs/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams RuleRequiresJobOnePathParams,
				queryParams RuleRequiresJobLoadQueryParams,
				req RuleRequiresJob,
				rawReq any,
			) (server.Response[RuleRequiresJob], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[RuleRequiresJob]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[RuleRequiresJob]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePutRuleRequiresJob(arguments, db, waitForChange, object)
				if err != nil {
					return server.Response[RuleRequiresJob]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[RuleRequiresJob]{
					Status:     http.StatusOK,
					Success:    true,
					Error:      nil,
					Objects:    objects,
					Count:      count,
					TotalCount: totalCount,
					Limit:      limit,
					Offset:     offset,
				}, nil
			},
			RuleRequiresJob{},
		)
		if err != nil {
			panic(err)
		}
		r.Put(putHandler.PathWithinRouter, putHandler.ServeHTTP)
	}()

	func() {
		patchHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/rule-requires-jobs/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams RuleRequiresJobOnePathParams,
				queryParams RuleRequiresJobLoadQueryParams,
				req RuleRequiresJob,
				rawReq any,
			) (server.Response[RuleRequiresJob], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[RuleRequiresJob]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				forceSetValuesForFields := make([]string, 0)
				for _, possibleField := range maps.Keys(item) {
					if !slices.Contains(RuleRequiresJobTableColumns, possibleField) {
						continue
					}

					forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[RuleRequiresJob]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePatchRuleRequiresJob(arguments, db, waitForChange, object, forceSetValuesForFields)
				if err != nil {
					return server.Response[RuleRequiresJob]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[RuleRequiresJob]{
					Status:     http.StatusOK,
					Success:    true,
					Error:      nil,
					Objects:    objects,
					Count:      count,
					TotalCount: totalCount,
					Limit:      limit,
					Offset:     offset,
				}, nil
			},
			RuleRequiresJob{},
		)
		if err != nil {
			panic(err)
		}
		r.Patch(patchHandler.PathWithinRouter, patchHandler.ServeHTTP)
	}()

	func() {
		deleteHandler, err := getHTTPHandler(
			http.MethodDelete,
			"/rule-requires-jobs/{primaryKey}",
			http.StatusNoContent,
			func(
				ctx context.Context,
				pathParams RuleRequiresJobOnePathParams,
				queryParams RuleRequiresJobLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.EmptyResponse, error) {
				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				object := &RuleRequiresJob{}
				object.ID = pathParams.PrimaryKey

				err = handleDeleteRuleRequiresJob(arguments, db, waitForChange, object)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				return server.EmptyResponse{}, nil
			},
			RuleRequiresJob{},
		)
		if err != nil {
			panic(err)
		}
		r.Delete(deleteHandler.PathWithinRouter, deleteHandler.ServeHTTP)
	}()

	return r
}

func NewRuleRequiresJobFromItem(item map[string]any) (any, error) {
	object := &RuleRequiresJob{}

	err := object.FromItem(item)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func init() {
	register(
		RuleRequiresJobTable,
		RuleRequiresJob{},
		NewRuleRequiresJobFromItem,
		"/rule-requires-jobs",
		GetRuleRequiresJobRouter,
	)
}