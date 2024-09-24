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

type Trigger struct {
	ID                                              uuid.UUID              `json:"id"`
	CreatedAt                                       time.Time              `json:"created_at"`
	UpdatedAt                                       time.Time              `json:"updated_at"`
	DeletedAt                                       *time.Time             `json:"deleted_at"`
	JobExecutorClaimedUntil                         time.Time              `json:"job_executor_claimed_until"`
	JobExecutionStartedAt                           *time.Time             `json:"job_execution_started_at"`
	RuleID                                          uuid.UUID              `json:"rule_id"`
	RuleIDObject                                    *Rule                  `json:"rule_id_object"`
	ChangeID                                        uuid.UUID              `json:"change_id"`
	ChangeIDObject                                  *Change                `json:"change_id_object"`
	ReferencedByExecutionTriggerIDObjects           []*Execution           `json:"referenced_by_execution_trigger_id_objects"`
	ReferencedByTriggerHasExecutionTriggerIDObjects []*TriggerHasExecution `json:"referenced_by_trigger_has_execution_trigger_id_objects"`
}

var TriggerTable = "trigger"

var TriggerTableNamespaceID int32 = 1337 + 10

var (
	TriggerTableIDColumn                      = "id"
	TriggerTableCreatedAtColumn               = "created_at"
	TriggerTableUpdatedAtColumn               = "updated_at"
	TriggerTableDeletedAtColumn               = "deleted_at"
	TriggerTableJobExecutorClaimedUntilColumn = "job_executor_claimed_until"
	TriggerTableJobExecutionStartedAtColumn   = "job_execution_started_at"
	TriggerTableRuleIDColumn                  = "rule_id"
	TriggerTableChangeIDColumn                = "change_id"
)

var (
	TriggerTableIDColumnWithTypeCast                      = `"id" AS id`
	TriggerTableCreatedAtColumnWithTypeCast               = `"created_at" AS created_at`
	TriggerTableUpdatedAtColumnWithTypeCast               = `"updated_at" AS updated_at`
	TriggerTableDeletedAtColumnWithTypeCast               = `"deleted_at" AS deleted_at`
	TriggerTableJobExecutorClaimedUntilColumnWithTypeCast = `"job_executor_claimed_until" AS job_executor_claimed_until`
	TriggerTableJobExecutionStartedAtColumnWithTypeCast   = `"job_execution_started_at" AS job_execution_started_at`
	TriggerTableRuleIDColumnWithTypeCast                  = `"rule_id" AS rule_id`
	TriggerTableChangeIDColumnWithTypeCast                = `"change_id" AS change_id`
)

var TriggerTableColumns = []string{
	TriggerTableIDColumn,
	TriggerTableCreatedAtColumn,
	TriggerTableUpdatedAtColumn,
	TriggerTableDeletedAtColumn,
	TriggerTableJobExecutorClaimedUntilColumn,
	TriggerTableJobExecutionStartedAtColumn,
	TriggerTableRuleIDColumn,
	TriggerTableChangeIDColumn,
}

var TriggerTableColumnsWithTypeCasts = []string{
	TriggerTableIDColumnWithTypeCast,
	TriggerTableCreatedAtColumnWithTypeCast,
	TriggerTableUpdatedAtColumnWithTypeCast,
	TriggerTableDeletedAtColumnWithTypeCast,
	TriggerTableJobExecutorClaimedUntilColumnWithTypeCast,
	TriggerTableJobExecutionStartedAtColumnWithTypeCast,
	TriggerTableRuleIDColumnWithTypeCast,
	TriggerTableChangeIDColumnWithTypeCast,
}

var TriggerIntrospectedTable *introspect.Table

var TriggerTableColumnLookup map[string]*introspect.Column

var (
	TriggerTablePrimaryKeyColumn = TriggerTableIDColumn
)

func init() {
	TriggerIntrospectedTable = tableByName[TriggerTable]

	/* only needed during templating */
	if TriggerIntrospectedTable == nil {
		TriggerIntrospectedTable = &introspect.Table{}
	}

	TriggerTableColumnLookup = TriggerIntrospectedTable.ColumnByName
}

type TriggerOnePathParams struct {
	PrimaryKey uuid.UUID `json:"primaryKey"`
}

type TriggerLoadQueryParams struct {
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

func (m *Trigger) GetPrimaryKeyColumn() string {
	return TriggerTablePrimaryKeyColumn
}

func (m *Trigger) GetPrimaryKeyValue() any {
	return m.ID
}

func (m *Trigger) FromItem(item map[string]any) error {
	if item == nil {
		return fmt.Errorf(
			"item unexpectedly nil during TriggerFromItem",
		)
	}

	if len(item) == 0 {
		return fmt.Errorf(
			"item unexpectedly empty during TriggerFromItem",
		)
	}

	wrapError := func(k string, v any, err error) error {
		return fmt.Errorf("%v: %#+v; error; %v", k, v, err)
	}

	for k, v := range item {
		_, ok := TriggerTableColumnLookup[k]
		if !ok {
			return fmt.Errorf(
				"item contained unexpected key %#+v during TriggerFromItem; item: %#+v",
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

		case "job_executor_claimed_until":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uujob_executor_claimed_until.UUID", temp1))
				}
			}

			m.JobExecutorClaimedUntil = temp2

		case "job_execution_started_at":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uujob_execution_started_at.UUID", temp1))
				}
			}

			m.JobExecutionStartedAt = &temp2

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

		case "change_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uuchange_id.UUID", temp1))
				}
			}

			m.ChangeID = temp2

		}
	}

	return nil
}

func (m *Trigger) Reload(ctx context.Context, tx pgx.Tx, includeDeleteds ...bool) error {
	extraWhere := ""
	if len(includeDeleteds) > 0 && includeDeleteds[0] {
		if slices.Contains(TriggerTableColumns, "deleted_at") {
			extraWhere = "\n    AND (deleted_at IS null OR deleted_at IS NOT null)"
		}
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	o, _, _, _, _, err := SelectTrigger(
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
	m.JobExecutorClaimedUntil = o.JobExecutorClaimedUntil
	m.JobExecutionStartedAt = o.JobExecutionStartedAt
	m.RuleID = o.RuleID
	m.RuleIDObject = o.RuleIDObject
	m.ChangeID = o.ChangeID
	m.ChangeIDObject = o.ChangeIDObject
	m.ReferencedByExecutionTriggerIDObjects = o.ReferencedByExecutionTriggerIDObjects
	m.ReferencedByTriggerHasExecutionTriggerIDObjects = o.ReferencedByTriggerHasExecutionTriggerIDObjects

	return nil
}

func (m *Trigger) Insert(ctx context.Context, tx pgx.Tx, setPrimaryKey bool, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setPrimaryKey && (setZeroValues || !types.IsZeroUUID(m.ID) || slices.Contains(forceSetValuesForFields, TriggerTableIDColumn) || isRequired(TriggerTableColumnLookup, TriggerTableIDColumn)) {
		columns = append(columns, TriggerTableIDColumn)

		v, err := types.FormatUUID(m.ID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, TriggerTableCreatedAtColumn) || isRequired(TriggerTableColumnLookup, TriggerTableCreatedAtColumn) {
		columns = append(columns, TriggerTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, TriggerTableUpdatedAtColumn) || isRequired(TriggerTableColumnLookup, TriggerTableUpdatedAtColumn) {
		columns = append(columns, TriggerTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, TriggerTableDeletedAtColumn) || isRequired(TriggerTableColumnLookup, TriggerTableDeletedAtColumn) {
		columns = append(columns, TriggerTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.JobExecutorClaimedUntil) || slices.Contains(forceSetValuesForFields, TriggerTableJobExecutorClaimedUntilColumn) || isRequired(TriggerTableColumnLookup, TriggerTableJobExecutorClaimedUntilColumn) {
		columns = append(columns, TriggerTableJobExecutorClaimedUntilColumn)

		v, err := types.FormatTime(m.JobExecutorClaimedUntil)
		if err != nil {
			return fmt.Errorf("failed to handle m.JobExecutorClaimedUntil; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.JobExecutionStartedAt) || slices.Contains(forceSetValuesForFields, TriggerTableJobExecutionStartedAtColumn) || isRequired(TriggerTableColumnLookup, TriggerTableJobExecutionStartedAtColumn) {
		columns = append(columns, TriggerTableJobExecutionStartedAtColumn)

		v, err := types.FormatTime(m.JobExecutionStartedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.JobExecutionStartedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.RuleID) || slices.Contains(forceSetValuesForFields, TriggerTableRuleIDColumn) || isRequired(TriggerTableColumnLookup, TriggerTableRuleIDColumn) {
		columns = append(columns, TriggerTableRuleIDColumn)

		v, err := types.FormatUUID(m.RuleID)
		if err != nil {
			return fmt.Errorf("failed to handle m.RuleID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.ChangeID) || slices.Contains(forceSetValuesForFields, TriggerTableChangeIDColumn) || isRequired(TriggerTableColumnLookup, TriggerTableChangeIDColumn) {
		columns = append(columns, TriggerTableChangeIDColumn)

		v, err := types.FormatUUID(m.ChangeID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ChangeID; %v", err)
		}

		values = append(values, v)
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	item, err := query.Insert(
		ctx,
		tx,
		TriggerTable,
		columns,
		nil,
		false,
		false,
		TriggerTableColumns,
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to insert %#+v; %v", m, err)
	}
	v := (*item)[TriggerTableIDColumn]

	if v == nil {
		return fmt.Errorf("failed to find %v in %#+v", TriggerTableIDColumn, item)
	}

	wrapError := func(err error) error {
		return fmt.Errorf(
			"failed to treat %v: %#+v as uuid.UUID: %v",
			TriggerTableIDColumn,
			(*item)[TriggerTableIDColumn],
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

func (m *Trigger) Update(ctx context.Context, tx pgx.Tx, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, TriggerTableCreatedAtColumn) {
		columns = append(columns, TriggerTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, TriggerTableUpdatedAtColumn) {
		columns = append(columns, TriggerTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, TriggerTableDeletedAtColumn) {
		columns = append(columns, TriggerTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.JobExecutorClaimedUntil) || slices.Contains(forceSetValuesForFields, TriggerTableJobExecutorClaimedUntilColumn) {
		columns = append(columns, TriggerTableJobExecutorClaimedUntilColumn)

		v, err := types.FormatTime(m.JobExecutorClaimedUntil)
		if err != nil {
			return fmt.Errorf("failed to handle m.JobExecutorClaimedUntil; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.JobExecutionStartedAt) || slices.Contains(forceSetValuesForFields, TriggerTableJobExecutionStartedAtColumn) {
		columns = append(columns, TriggerTableJobExecutionStartedAtColumn)

		v, err := types.FormatTime(m.JobExecutionStartedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.JobExecutionStartedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.RuleID) || slices.Contains(forceSetValuesForFields, TriggerTableRuleIDColumn) {
		columns = append(columns, TriggerTableRuleIDColumn)

		v, err := types.FormatUUID(m.RuleID)
		if err != nil {
			return fmt.Errorf("failed to handle m.RuleID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.ChangeID) || slices.Contains(forceSetValuesForFields, TriggerTableChangeIDColumn) {
		columns = append(columns, TriggerTableChangeIDColumn)

		v, err := types.FormatUUID(m.ChangeID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ChangeID; %v", err)
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
		TriggerTable,
		columns,
		fmt.Sprintf("%v = $$??", TriggerTableIDColumn),
		TriggerTableColumns,
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

func (m *Trigger) Delete(ctx context.Context, tx pgx.Tx, hardDeletes ...bool) error {
	hardDelete := false
	if len(hardDeletes) > 0 {
		hardDelete = hardDeletes[0]
	}

	if !hardDelete && slices.Contains(TriggerTableColumns, "deleted_at") {
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
		TriggerTable,
		fmt.Sprintf("%v = $$??", TriggerTableIDColumn),
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to delete %#+v; %v", m, err)
	}

	_ = m.Reload(ctx, tx, true)

	return nil
}

func (m *Trigger) LockTable(ctx context.Context, tx pgx.Tx, timeouts ...time.Duration) error {
	return query.LockTable(ctx, tx, TriggerTable, timeouts...)
}

func (m *Trigger) LockTableWithRetries(ctx context.Context, tx pgx.Tx, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.LockTableWithRetries(ctx, tx, TriggerTable, overallTimeout, individualAttempttimeout)
}

func (m *Trigger) AdvisoryLock(ctx context.Context, tx pgx.Tx, key int32, timeouts ...time.Duration) error {
	return query.AdvisoryLock(ctx, tx, TriggerTableNamespaceID, key, timeouts...)
}

func (m *Trigger) AdvisoryLockWithRetries(ctx context.Context, tx pgx.Tx, key int32, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.AdvisoryLockWithRetries(ctx, tx, TriggerTableNamespaceID, key, overallTimeout, individualAttempttimeout)
}

func SelectTriggers(ctx context.Context, tx pgx.Tx, where string, orderBy *string, limit *int, offset *int, values ...any) ([]*Trigger, int64, int64, int64, int64, error) {
	before := time.Now()

	if config.Debug() {
		log.Printf("entered SelectTriggers")

		defer func() {
			log.Printf("exited SelectTriggers in %s", time.Since(before))
		}()
	}
	if slices.Contains(TriggerTableColumns, "deleted_at") {
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
	ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", TriggerTable, nil), !isLoadQuery)
	if !ok {
		return []*Trigger{}, 0, 0, 0, 0, nil
	}

	items, count, totalCount, page, totalPages, err := query.Select(
		ctx,
		tx,
		TriggerTableColumnsWithTypeCasts,
		TriggerTable,
		where,
		orderBy,
		limit,
		offset,
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectTriggers; %v", err)
	}

	objects := make([]*Trigger, 0)

	for _, item := range *items {
		object := &Trigger{}

		err = object.FromItem(item)
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		if !types.IsZeroUUID(object.RuleID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", RuleTable, object.RuleID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectTriggers->SelectRule for object.RuleIDObject")
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
					log.Printf("loaded SelectTriggers->SelectRule for object.RuleIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.ChangeID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", ChangeTable, object.ChangeID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectTriggers->SelectChange for object.ChangeIDObject")
				}

				object.ChangeIDObject, _, _, _, _, err = SelectChange(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", ChangeTablePrimaryKeyColumn),
					object.ChangeID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectTriggers->SelectChange for object.ChangeIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		err = func() error {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("__ReferencedBy__%s{%v}", TriggerTable, object.GetPrimaryKeyValue()), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectTriggers->SelectExecutions for object.ReferencedByExecutionTriggerIDObjects")
				}

				object.ReferencedByExecutionTriggerIDObjects, _, _, _, _, err = SelectExecutions(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", ExecutionTableTriggerIDColumn),
					nil,
					nil,
					nil,
					object.GetPrimaryKeyValue(),
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectTriggers->SelectExecutions for object.ReferencedByExecutionTriggerIDObjects in %s", time.Since(thisBefore))
				}

			}

			return nil
		}()
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		err = func() error {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("__ReferencedBy__%s{%v}", TriggerTable, object.GetPrimaryKeyValue()), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectTriggers->SelectTriggerHasExecutions for object.ReferencedByTriggerHasExecutionTriggerIDObjects")
				}

				object.ReferencedByTriggerHasExecutionTriggerIDObjects, _, _, _, _, err = SelectTriggerHasExecutions(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", TriggerHasExecutionTableTriggerIDColumn),
					nil,
					nil,
					nil,
					object.GetPrimaryKeyValue(),
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectTriggers->SelectTriggerHasExecutions for object.ReferencedByTriggerHasExecutionTriggerIDObjects in %s", time.Since(thisBefore))
				}

			}

			return nil
		}()
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		objects = append(objects, object)
	}

	return objects, count, totalCount, page, totalPages, nil
}

func SelectTrigger(ctx context.Context, tx pgx.Tx, where string, values ...any) (*Trigger, int64, int64, int64, int64, error) {
	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	objects, _, _, _, _, err := SelectTriggers(
		ctx,
		tx,
		where,
		nil,
		helpers.Ptr(2),
		helpers.Ptr(0),
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectTrigger; %v", err)
	}

	if len(objects) > 1 {
		return nil, 0, 0, 0, 0, fmt.Errorf("attempt to call SelectTrigger returned more than 1 row")
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

func handleGetTriggers(arguments *server.SelectManyArguments, db *pgxpool.Pool) ([]*Trigger, int64, int64, int64, int64, error) {
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

	objects, count, totalCount, page, totalPages, err := SelectTriggers(arguments.Ctx, tx, arguments.Where, arguments.OrderBy, arguments.Limit, arguments.Offset, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return objects, count, totalCount, page, totalPages, nil
}

func handleGetTrigger(arguments *server.SelectOneArguments, db *pgxpool.Pool, primaryKey uuid.UUID) ([]*Trigger, int64, int64, int64, int64, error) {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	object, count, totalCount, page, totalPages, err := SelectTrigger(arguments.Ctx, tx, arguments.Where, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return []*Trigger{object}, count, totalCount, page, totalPages, nil
}

func handlePostTriggers(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, objects []*Trigger, forceSetValuesForFieldsByObjectIndex [][]string) ([]*Trigger, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.INSERT}, TriggerTable, xid)
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

func handlePutTrigger(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Trigger) ([]*Trigger, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, TriggerTable, xid)
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

	return []*Trigger{object}, count, totalCount, page, totalPages, nil
}

func handlePatchTrigger(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Trigger, forceSetValuesForFields []string) ([]*Trigger, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, TriggerTable, xid)
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

	return []*Trigger{object}, count, totalCount, page, totalPages, nil
}

func handleDeleteTrigger(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Trigger) error {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.DELETE, stream.SOFT_DELETE}, TriggerTable, xid)
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

func GetTriggerRouter(db *pgxpool.Pool, redisPool *redis.Pool, httpMiddlewares []server.HTTPMiddleware, objectMiddlewares []server.ObjectMiddleware, waitForChange server.WaitForChange) chi.Router {
	r := chi.NewRouter()

	for _, m := range httpMiddlewares {
		r.Use(m)
	}

	func() {
		getManyHandler, err := getHTTPHandler(
			http.MethodGet,
			"/triggers",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams map[string]any,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[Trigger], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectManyArguments(ctx, queryParams, TriggerIntrospectedTable, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Trigger]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Trigger]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[Trigger]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[Trigger]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetTriggers(arguments, db)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Trigger]{}, err
				}

				limit := int64(0)
				if arguments.Limit != nil {
					limit = int64(*arguments.Limit)
				}

				offset := int64(0)
				if arguments.Offset != nil {
					offset = int64(*arguments.Offset)
				}

				response := server.Response[Trigger]{
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

					return server.Response[Trigger]{}, err
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
			Trigger{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getManyHandler.PathWithinRouter, getManyHandler.ServeHTTP)
	}()

	func() {
		getOneHandler, err := getHTTPHandler(
			http.MethodGet,
			"/triggers/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams TriggerOnePathParams,
				queryParams TriggerLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[Trigger], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectOneArguments(ctx, queryParams.Depth, TriggerIntrospectedTable, pathParams.PrimaryKey, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Trigger]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Trigger]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[Trigger]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[Trigger]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetTrigger(arguments, db, pathParams.PrimaryKey)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Trigger]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				response := server.Response[Trigger]{
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

					return server.Response[Trigger]{}, err
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
			Trigger{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getOneHandler.PathWithinRouter, getOneHandler.ServeHTTP)
	}()

	func() {
		postHandler, err := getHTTPHandler(
			http.MethodPost,
			"/triggers",
			http.StatusCreated,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams TriggerLoadQueryParams,
				req []*Trigger,
				rawReq any,
			) (server.Response[Trigger], error) {
				allRawItems, ok := rawReq.([]any)
				if !ok {
					return server.Response[Trigger]{}, fmt.Errorf("failed to cast %#+v to []map[string]any", rawReq)
				}

				allItems := make([]map[string]any, 0)
				for _, rawItem := range allRawItems {
					item, ok := rawItem.(map[string]any)
					if !ok {
						return server.Response[Trigger]{}, fmt.Errorf("failed to cast %#+v to map[string]any", rawItem)
					}

					allItems = append(allItems, item)
				}

				forceSetValuesForFieldsByObjectIndex := make([][]string, 0)
				for _, item := range allItems {
					forceSetValuesForFields := make([]string, 0)
					for _, possibleField := range maps.Keys(item) {
						if !slices.Contains(TriggerTableColumns, possibleField) {
							continue
						}

						forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
					}
					forceSetValuesForFieldsByObjectIndex = append(forceSetValuesForFieldsByObjectIndex, forceSetValuesForFields)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Trigger]{}, err
				}

				objects, count, totalCount, _, _, err := handlePostTriggers(arguments, db, waitForChange, req, forceSetValuesForFieldsByObjectIndex)
				if err != nil {
					return server.Response[Trigger]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Trigger]{
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
			Trigger{},
		)
		if err != nil {
			panic(err)
		}
		r.Post(postHandler.PathWithinRouter, postHandler.ServeHTTP)
	}()

	func() {
		putHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/triggers/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams TriggerOnePathParams,
				queryParams TriggerLoadQueryParams,
				req Trigger,
				rawReq any,
			) (server.Response[Trigger], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[Trigger]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Trigger]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePutTrigger(arguments, db, waitForChange, object)
				if err != nil {
					return server.Response[Trigger]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Trigger]{
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
			Trigger{},
		)
		if err != nil {
			panic(err)
		}
		r.Put(putHandler.PathWithinRouter, putHandler.ServeHTTP)
	}()

	func() {
		patchHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/triggers/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams TriggerOnePathParams,
				queryParams TriggerLoadQueryParams,
				req Trigger,
				rawReq any,
			) (server.Response[Trigger], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[Trigger]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				forceSetValuesForFields := make([]string, 0)
				for _, possibleField := range maps.Keys(item) {
					if !slices.Contains(TriggerTableColumns, possibleField) {
						continue
					}

					forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Trigger]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePatchTrigger(arguments, db, waitForChange, object, forceSetValuesForFields)
				if err != nil {
					return server.Response[Trigger]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Trigger]{
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
			Trigger{},
		)
		if err != nil {
			panic(err)
		}
		r.Patch(patchHandler.PathWithinRouter, patchHandler.ServeHTTP)
	}()

	func() {
		deleteHandler, err := getHTTPHandler(
			http.MethodDelete,
			"/triggers/{primaryKey}",
			http.StatusNoContent,
			func(
				ctx context.Context,
				pathParams TriggerOnePathParams,
				queryParams TriggerLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.EmptyResponse, error) {
				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				object := &Trigger{}
				object.ID = pathParams.PrimaryKey

				err = handleDeleteTrigger(arguments, db, waitForChange, object)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				return server.EmptyResponse{}, nil
			},
			Trigger{},
		)
		if err != nil {
			panic(err)
		}
		r.Delete(deleteHandler.PathWithinRouter, deleteHandler.ServeHTTP)
	}()

	return r
}

func NewTriggerFromItem(item map[string]any) (any, error) {
	object := &Trigger{}

	err := object.FromItem(item)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func init() {
	register(
		TriggerTable,
		Trigger{},
		NewTriggerFromItem,
		"/triggers",
		GetTriggerRouter,
	)
}
