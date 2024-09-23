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

type Execution struct {
	ID                                                uuid.UUID              `json:"id"`
	CreatedAt                                         time.Time              `json:"created_at"`
	UpdatedAt                                         time.Time              `json:"updated_at"`
	DeletedAt                                         *time.Time             `json:"deleted_at"`
	TriggerID                                         uuid.UUID              `json:"trigger_id"`
	TriggerIDObject                                   *Trigger               `json:"trigger_id_object"`
	BuildOutputID                                     *uuid.UUID             `json:"build_output_id"`
	BuildOutputIDObject                               *Output                `json:"build_output_id_object"`
	TestOutputID                                      *uuid.UUID             `json:"test_output_id"`
	TestOutputIDObject                                *Output                `json:"test_output_id_object"`
	PublishOutputID                                   *uuid.UUID             `json:"publish_output_id"`
	PublishOutputIDObject                             *Output                `json:"publish_output_id_object"`
	DeployOutputID                                    *uuid.UUID             `json:"deploy_output_id"`
	DeployOutputIDObject                              *Output                `json:"deploy_output_id_object"`
	ValidateOutputID                                  *uuid.UUID             `json:"validate_output_id"`
	ValidateOutputIDObject                            *Output                `json:"validate_output_id_object"`
	JobID                                             uuid.UUID              `json:"job_id"`
	JobIDObject                                       *Job                   `json:"job_id_object"`
	ReferencedByTriggerHasExecutionExecutionIDObjects []*TriggerHasExecution `json:"referenced_by_trigger_has_execution_execution_id_objects"`
}

var ExecutionTable = "execution"

var ExecutionTableNamespaceID int32 = 1337 + 2

var (
	ExecutionTableIDColumn               = "id"
	ExecutionTableCreatedAtColumn        = "created_at"
	ExecutionTableUpdatedAtColumn        = "updated_at"
	ExecutionTableDeletedAtColumn        = "deleted_at"
	ExecutionTableTriggerIDColumn        = "trigger_id"
	ExecutionTableBuildOutputIDColumn    = "build_output_id"
	ExecutionTableTestOutputIDColumn     = "test_output_id"
	ExecutionTablePublishOutputIDColumn  = "publish_output_id"
	ExecutionTableDeployOutputIDColumn   = "deploy_output_id"
	ExecutionTableValidateOutputIDColumn = "validate_output_id"
	ExecutionTableJobIDColumn            = "job_id"
)

var (
	ExecutionTableIDColumnWithTypeCast               = `"id" AS id`
	ExecutionTableCreatedAtColumnWithTypeCast        = `"created_at" AS created_at`
	ExecutionTableUpdatedAtColumnWithTypeCast        = `"updated_at" AS updated_at`
	ExecutionTableDeletedAtColumnWithTypeCast        = `"deleted_at" AS deleted_at`
	ExecutionTableTriggerIDColumnWithTypeCast        = `"trigger_id" AS trigger_id`
	ExecutionTableBuildOutputIDColumnWithTypeCast    = `"build_output_id" AS build_output_id`
	ExecutionTableTestOutputIDColumnWithTypeCast     = `"test_output_id" AS test_output_id`
	ExecutionTablePublishOutputIDColumnWithTypeCast  = `"publish_output_id" AS publish_output_id`
	ExecutionTableDeployOutputIDColumnWithTypeCast   = `"deploy_output_id" AS deploy_output_id`
	ExecutionTableValidateOutputIDColumnWithTypeCast = `"validate_output_id" AS validate_output_id`
	ExecutionTableJobIDColumnWithTypeCast            = `"job_id" AS job_id`
)

var ExecutionTableColumns = []string{
	ExecutionTableIDColumn,
	ExecutionTableCreatedAtColumn,
	ExecutionTableUpdatedAtColumn,
	ExecutionTableDeletedAtColumn,
	ExecutionTableTriggerIDColumn,
	ExecutionTableBuildOutputIDColumn,
	ExecutionTableTestOutputIDColumn,
	ExecutionTablePublishOutputIDColumn,
	ExecutionTableDeployOutputIDColumn,
	ExecutionTableValidateOutputIDColumn,
	ExecutionTableJobIDColumn,
}

var ExecutionTableColumnsWithTypeCasts = []string{
	ExecutionTableIDColumnWithTypeCast,
	ExecutionTableCreatedAtColumnWithTypeCast,
	ExecutionTableUpdatedAtColumnWithTypeCast,
	ExecutionTableDeletedAtColumnWithTypeCast,
	ExecutionTableTriggerIDColumnWithTypeCast,
	ExecutionTableBuildOutputIDColumnWithTypeCast,
	ExecutionTableTestOutputIDColumnWithTypeCast,
	ExecutionTablePublishOutputIDColumnWithTypeCast,
	ExecutionTableDeployOutputIDColumnWithTypeCast,
	ExecutionTableValidateOutputIDColumnWithTypeCast,
	ExecutionTableJobIDColumnWithTypeCast,
}

var ExecutionIntrospectedTable *introspect.Table

var ExecutionTableColumnLookup map[string]*introspect.Column

var (
	ExecutionTablePrimaryKeyColumn = ExecutionTableIDColumn
)

func init() {
	ExecutionIntrospectedTable = tableByName[ExecutionTable]

	/* only needed during templating */
	if ExecutionIntrospectedTable == nil {
		ExecutionIntrospectedTable = &introspect.Table{}
	}

	ExecutionTableColumnLookup = ExecutionIntrospectedTable.ColumnByName
}

type ExecutionOnePathParams struct {
	PrimaryKey uuid.UUID `json:"primaryKey"`
}

type ExecutionLoadQueryParams struct {
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

func (m *Execution) GetPrimaryKeyColumn() string {
	return ExecutionTablePrimaryKeyColumn
}

func (m *Execution) GetPrimaryKeyValue() any {
	return m.ID
}

func (m *Execution) FromItem(item map[string]any) error {
	if item == nil {
		return fmt.Errorf(
			"item unexpectedly nil during ExecutionFromItem",
		)
	}

	if len(item) == 0 {
		return fmt.Errorf(
			"item unexpectedly empty during ExecutionFromItem",
		)
	}

	wrapError := func(k string, v any, err error) error {
		return fmt.Errorf("%v: %#+v; error; %v", k, v, err)
	}

	for k, v := range item {
		_, ok := ExecutionTableColumnLookup[k]
		if !ok {
			return fmt.Errorf(
				"item contained unexpected key %#+v during ExecutionFromItem; item: %#+v",
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

		case "trigger_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uutrigger_id.UUID", temp1))
				}
			}

			m.TriggerID = temp2

		case "build_output_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uubuild_output_id.UUID", temp1))
				}
			}

			m.BuildOutputID = &temp2

		case "test_output_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uutest_output_id.UUID", temp1))
				}
			}

			m.TestOutputID = &temp2

		case "publish_output_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uupublish_output_id.UUID", temp1))
				}
			}

			m.PublishOutputID = &temp2

		case "deploy_output_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uudeploy_output_id.UUID", temp1))
				}
			}

			m.DeployOutputID = &temp2

		case "validate_output_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uuvalidate_output_id.UUID", temp1))
				}
			}

			m.ValidateOutputID = &temp2

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

func (m *Execution) Reload(ctx context.Context, tx pgx.Tx, includeDeleteds ...bool) error {
	extraWhere := ""
	if len(includeDeleteds) > 0 && includeDeleteds[0] {
		if slices.Contains(ExecutionTableColumns, "deleted_at") {
			extraWhere = "\n    AND (deleted_at IS null OR deleted_at IS NOT null)"
		}
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	o, _, _, _, _, err := SelectExecution(
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
	m.TriggerID = o.TriggerID
	m.TriggerIDObject = o.TriggerIDObject
	m.BuildOutputID = o.BuildOutputID
	m.BuildOutputIDObject = o.BuildOutputIDObject
	m.TestOutputID = o.TestOutputID
	m.TestOutputIDObject = o.TestOutputIDObject
	m.PublishOutputID = o.PublishOutputID
	m.PublishOutputIDObject = o.PublishOutputIDObject
	m.DeployOutputID = o.DeployOutputID
	m.DeployOutputIDObject = o.DeployOutputIDObject
	m.ValidateOutputID = o.ValidateOutputID
	m.ValidateOutputIDObject = o.ValidateOutputIDObject
	m.JobID = o.JobID
	m.JobIDObject = o.JobIDObject
	m.ReferencedByTriggerHasExecutionExecutionIDObjects = o.ReferencedByTriggerHasExecutionExecutionIDObjects

	return nil
}

func (m *Execution) Insert(ctx context.Context, tx pgx.Tx, setPrimaryKey bool, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setPrimaryKey && (setZeroValues || !types.IsZeroUUID(m.ID) || slices.Contains(forceSetValuesForFields, ExecutionTableIDColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableIDColumn)) {
		columns = append(columns, ExecutionTableIDColumn)

		v, err := types.FormatUUID(m.ID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, ExecutionTableCreatedAtColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableCreatedAtColumn) {
		columns = append(columns, ExecutionTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, ExecutionTableUpdatedAtColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableUpdatedAtColumn) {
		columns = append(columns, ExecutionTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, ExecutionTableDeletedAtColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableDeletedAtColumn) {
		columns = append(columns, ExecutionTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.TriggerID) || slices.Contains(forceSetValuesForFields, ExecutionTableTriggerIDColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableTriggerIDColumn) {
		columns = append(columns, ExecutionTableTriggerIDColumn)

		v, err := types.FormatUUID(m.TriggerID)
		if err != nil {
			return fmt.Errorf("failed to handle m.TriggerID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.BuildOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTableBuildOutputIDColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableBuildOutputIDColumn) {
		columns = append(columns, ExecutionTableBuildOutputIDColumn)

		v, err := types.FormatUUID(m.BuildOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.BuildOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.TestOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTableTestOutputIDColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableTestOutputIDColumn) {
		columns = append(columns, ExecutionTableTestOutputIDColumn)

		v, err := types.FormatUUID(m.TestOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.TestOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.PublishOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTablePublishOutputIDColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTablePublishOutputIDColumn) {
		columns = append(columns, ExecutionTablePublishOutputIDColumn)

		v, err := types.FormatUUID(m.PublishOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.PublishOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.DeployOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTableDeployOutputIDColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableDeployOutputIDColumn) {
		columns = append(columns, ExecutionTableDeployOutputIDColumn)

		v, err := types.FormatUUID(m.DeployOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeployOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.ValidateOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTableValidateOutputIDColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableValidateOutputIDColumn) {
		columns = append(columns, ExecutionTableValidateOutputIDColumn)

		v, err := types.FormatUUID(m.ValidateOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ValidateOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.JobID) || slices.Contains(forceSetValuesForFields, ExecutionTableJobIDColumn) || isRequired(ExecutionTableColumnLookup, ExecutionTableJobIDColumn) {
		columns = append(columns, ExecutionTableJobIDColumn)

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
		ExecutionTable,
		columns,
		nil,
		false,
		false,
		ExecutionTableColumns,
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to insert %#+v; %v", m, err)
	}
	v := (*item)[ExecutionTableIDColumn]

	if v == nil {
		return fmt.Errorf("failed to find %v in %#+v", ExecutionTableIDColumn, item)
	}

	wrapError := func(err error) error {
		return fmt.Errorf(
			"failed to treat %v: %#+v as uuid.UUID: %v",
			ExecutionTableIDColumn,
			(*item)[ExecutionTableIDColumn],
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

func (m *Execution) Update(ctx context.Context, tx pgx.Tx, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, ExecutionTableCreatedAtColumn) {
		columns = append(columns, ExecutionTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, ExecutionTableUpdatedAtColumn) {
		columns = append(columns, ExecutionTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, ExecutionTableDeletedAtColumn) {
		columns = append(columns, ExecutionTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.TriggerID) || slices.Contains(forceSetValuesForFields, ExecutionTableTriggerIDColumn) {
		columns = append(columns, ExecutionTableTriggerIDColumn)

		v, err := types.FormatUUID(m.TriggerID)
		if err != nil {
			return fmt.Errorf("failed to handle m.TriggerID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.BuildOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTableBuildOutputIDColumn) {
		columns = append(columns, ExecutionTableBuildOutputIDColumn)

		v, err := types.FormatUUID(m.BuildOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.BuildOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.TestOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTableTestOutputIDColumn) {
		columns = append(columns, ExecutionTableTestOutputIDColumn)

		v, err := types.FormatUUID(m.TestOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.TestOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.PublishOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTablePublishOutputIDColumn) {
		columns = append(columns, ExecutionTablePublishOutputIDColumn)

		v, err := types.FormatUUID(m.PublishOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.PublishOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.DeployOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTableDeployOutputIDColumn) {
		columns = append(columns, ExecutionTableDeployOutputIDColumn)

		v, err := types.FormatUUID(m.DeployOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeployOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.ValidateOutputID) || slices.Contains(forceSetValuesForFields, ExecutionTableValidateOutputIDColumn) {
		columns = append(columns, ExecutionTableValidateOutputIDColumn)

		v, err := types.FormatUUID(m.ValidateOutputID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ValidateOutputID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.JobID) || slices.Contains(forceSetValuesForFields, ExecutionTableJobIDColumn) {
		columns = append(columns, ExecutionTableJobIDColumn)

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
		ExecutionTable,
		columns,
		fmt.Sprintf("%v = $$??", ExecutionTableIDColumn),
		ExecutionTableColumns,
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

func (m *Execution) Delete(ctx context.Context, tx pgx.Tx, hardDeletes ...bool) error {
	hardDelete := false
	if len(hardDeletes) > 0 {
		hardDelete = hardDeletes[0]
	}

	if !hardDelete && slices.Contains(ExecutionTableColumns, "deleted_at") {
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
		ExecutionTable,
		fmt.Sprintf("%v = $$??", ExecutionTableIDColumn),
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to delete %#+v; %v", m, err)
	}

	_ = m.Reload(ctx, tx, true)

	return nil
}

func (m *Execution) LockTable(ctx context.Context, tx pgx.Tx, timeouts ...time.Duration) error {
	return query.LockTable(ctx, tx, ExecutionTable, timeouts...)
}

func (m *Execution) LockTableWithRetries(ctx context.Context, tx pgx.Tx, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.LockTableWithRetries(ctx, tx, ExecutionTable, overallTimeout, individualAttempttimeout)
}

func (m *Execution) AdvisoryLock(ctx context.Context, tx pgx.Tx, key int32, timeouts ...time.Duration) error {
	return query.AdvisoryLock(ctx, tx, ExecutionTableNamespaceID, key, timeouts...)
}

func (m *Execution) AdvisoryLockWithRetries(ctx context.Context, tx pgx.Tx, key int32, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.AdvisoryLockWithRetries(ctx, tx, ExecutionTableNamespaceID, key, overallTimeout, individualAttempttimeout)
}

func SelectExecutions(ctx context.Context, tx pgx.Tx, where string, orderBy *string, limit *int, offset *int, values ...any) ([]*Execution, int64, int64, int64, int64, error) {
	before := time.Now()

	if config.Debug() {
		log.Printf("entered SelectExecutions")

		defer func() {
			log.Printf("exited SelectExecutions in %s", time.Since(before))
		}()
	}
	if slices.Contains(ExecutionTableColumns, "deleted_at") {
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
	ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", ExecutionTable, nil), !isLoadQuery)
	if !ok {
		return []*Execution{}, 0, 0, 0, 0, nil
	}

	items, count, totalCount, page, totalPages, err := query.Select(
		ctx,
		tx,
		ExecutionTableColumnsWithTypeCasts,
		ExecutionTable,
		where,
		orderBy,
		limit,
		offset,
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectExecutions; %v", err)
	}

	objects := make([]*Execution, 0)

	for _, item := range *items {
		object := &Execution{}

		err = object.FromItem(item)
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		if !types.IsZeroUUID(object.TriggerID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", TriggerTable, object.TriggerID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectExecutions->SelectTrigger for object.TriggerIDObject")
				}

				object.TriggerIDObject, _, _, _, _, err = SelectTrigger(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", TriggerTablePrimaryKeyColumn),
					object.TriggerID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectExecutions->SelectTrigger for object.TriggerIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.BuildOutputID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", OutputTable, object.BuildOutputID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectExecutions->SelectOutput for object.BuildOutputIDObject")
				}

				object.BuildOutputIDObject, _, _, _, _, err = SelectOutput(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", OutputTablePrimaryKeyColumn),
					object.BuildOutputID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectExecutions->SelectOutput for object.BuildOutputIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.TestOutputID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", OutputTable, object.TestOutputID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectExecutions->SelectOutput for object.TestOutputIDObject")
				}

				object.TestOutputIDObject, _, _, _, _, err = SelectOutput(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", OutputTablePrimaryKeyColumn),
					object.TestOutputID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectExecutions->SelectOutput for object.TestOutputIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.PublishOutputID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", OutputTable, object.PublishOutputID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectExecutions->SelectOutput for object.PublishOutputIDObject")
				}

				object.PublishOutputIDObject, _, _, _, _, err = SelectOutput(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", OutputTablePrimaryKeyColumn),
					object.PublishOutputID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectExecutions->SelectOutput for object.PublishOutputIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.DeployOutputID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", OutputTable, object.DeployOutputID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectExecutions->SelectOutput for object.DeployOutputIDObject")
				}

				object.DeployOutputIDObject, _, _, _, _, err = SelectOutput(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", OutputTablePrimaryKeyColumn),
					object.DeployOutputID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectExecutions->SelectOutput for object.DeployOutputIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.ValidateOutputID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", OutputTable, object.ValidateOutputID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectExecutions->SelectOutput for object.ValidateOutputIDObject")
				}

				object.ValidateOutputIDObject, _, _, _, _, err = SelectOutput(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", OutputTablePrimaryKeyColumn),
					object.ValidateOutputID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectExecutions->SelectOutput for object.ValidateOutputIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.JobID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", JobTable, object.JobID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectExecutions->SelectJob for object.JobIDObject")
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
					log.Printf("loaded SelectExecutions->SelectJob for object.JobIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		err = func() error {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("__ReferencedBy__%s{%v}", ExecutionTable, object.GetPrimaryKeyValue()), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectExecutions->SelectTriggerHasExecutions for object.ReferencedByTriggerHasExecutionExecutionIDObjects")
				}

				object.ReferencedByTriggerHasExecutionExecutionIDObjects, _, _, _, _, err = SelectTriggerHasExecutions(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", TriggerHasExecutionTableExecutionIDColumn),
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
					log.Printf("loaded SelectExecutions->SelectTriggerHasExecutions for object.ReferencedByTriggerHasExecutionExecutionIDObjects in %s", time.Since(thisBefore))
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

func SelectExecution(ctx context.Context, tx pgx.Tx, where string, values ...any) (*Execution, int64, int64, int64, int64, error) {
	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	objects, _, _, _, _, err := SelectExecutions(
		ctx,
		tx,
		where,
		nil,
		helpers.Ptr(2),
		helpers.Ptr(0),
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectExecution; %v", err)
	}

	if len(objects) > 1 {
		return nil, 0, 0, 0, 0, fmt.Errorf("attempt to call SelectExecution returned more than 1 row")
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

func handleGetExecutions(arguments *server.SelectManyArguments, db *pgxpool.Pool) ([]*Execution, int64, int64, int64, int64, error) {
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

	objects, count, totalCount, page, totalPages, err := SelectExecutions(arguments.Ctx, tx, arguments.Where, arguments.OrderBy, arguments.Limit, arguments.Offset, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return objects, count, totalCount, page, totalPages, nil
}

func handleGetExecution(arguments *server.SelectOneArguments, db *pgxpool.Pool, primaryKey uuid.UUID) ([]*Execution, int64, int64, int64, int64, error) {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	object, count, totalCount, page, totalPages, err := SelectExecution(arguments.Ctx, tx, arguments.Where, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return []*Execution{object}, count, totalCount, page, totalPages, nil
}

func handlePostExecutions(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, objects []*Execution, forceSetValuesForFieldsByObjectIndex [][]string) ([]*Execution, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.INSERT}, ExecutionTable, xid)
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

func handlePutExecution(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Execution) ([]*Execution, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, ExecutionTable, xid)
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

	return []*Execution{object}, count, totalCount, page, totalPages, nil
}

func handlePatchExecution(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Execution, forceSetValuesForFields []string) ([]*Execution, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, ExecutionTable, xid)
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

	return []*Execution{object}, count, totalCount, page, totalPages, nil
}

func handleDeleteExecution(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Execution) error {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.DELETE, stream.SOFT_DELETE}, ExecutionTable, xid)
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

func GetExecutionRouter(db *pgxpool.Pool, redisPool *redis.Pool, httpMiddlewares []server.HTTPMiddleware, objectMiddlewares []server.ObjectMiddleware, waitForChange server.WaitForChange) chi.Router {
	r := chi.NewRouter()

	for _, m := range httpMiddlewares {
		r.Use(m)
	}

	func() {
		getManyHandler, err := getHTTPHandler(
			http.MethodGet,
			"/executions",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams map[string]any,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[Execution], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectManyArguments(ctx, queryParams, ExecutionIntrospectedTable, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Execution]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Execution]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[Execution]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[Execution]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetExecutions(arguments, db)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Execution]{}, err
				}

				limit := int64(0)
				if arguments.Limit != nil {
					limit = int64(*arguments.Limit)
				}

				offset := int64(0)
				if arguments.Offset != nil {
					offset = int64(*arguments.Offset)
				}

				response := server.Response[Execution]{
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

					return server.Response[Execution]{}, err
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
			Execution{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getManyHandler.PathWithinRouter, getManyHandler.ServeHTTP)
	}()

	func() {
		getOneHandler, err := getHTTPHandler(
			http.MethodGet,
			"/executions/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams ExecutionOnePathParams,
				queryParams ExecutionLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[Execution], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectOneArguments(ctx, queryParams.Depth, ExecutionIntrospectedTable, pathParams.PrimaryKey, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Execution]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Execution]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[Execution]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[Execution]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetExecution(arguments, db, pathParams.PrimaryKey)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Execution]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				response := server.Response[Execution]{
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

					return server.Response[Execution]{}, err
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
			Execution{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getOneHandler.PathWithinRouter, getOneHandler.ServeHTTP)
	}()

	func() {
		postHandler, err := getHTTPHandler(
			http.MethodPost,
			"/executions",
			http.StatusCreated,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams ExecutionLoadQueryParams,
				req []*Execution,
				rawReq any,
			) (server.Response[Execution], error) {
				allRawItems, ok := rawReq.([]any)
				if !ok {
					return server.Response[Execution]{}, fmt.Errorf("failed to cast %#+v to []map[string]any", rawReq)
				}

				allItems := make([]map[string]any, 0)
				for _, rawItem := range allRawItems {
					item, ok := rawItem.(map[string]any)
					if !ok {
						return server.Response[Execution]{}, fmt.Errorf("failed to cast %#+v to map[string]any", rawItem)
					}

					allItems = append(allItems, item)
				}

				forceSetValuesForFieldsByObjectIndex := make([][]string, 0)
				for _, item := range allItems {
					forceSetValuesForFields := make([]string, 0)
					for _, possibleField := range maps.Keys(item) {
						if !slices.Contains(ExecutionTableColumns, possibleField) {
							continue
						}

						forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
					}
					forceSetValuesForFieldsByObjectIndex = append(forceSetValuesForFieldsByObjectIndex, forceSetValuesForFields)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Execution]{}, err
				}

				objects, count, totalCount, _, _, err := handlePostExecutions(arguments, db, waitForChange, req, forceSetValuesForFieldsByObjectIndex)
				if err != nil {
					return server.Response[Execution]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Execution]{
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
			Execution{},
		)
		if err != nil {
			panic(err)
		}
		r.Post(postHandler.PathWithinRouter, postHandler.ServeHTTP)
	}()

	func() {
		putHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/executions/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams ExecutionOnePathParams,
				queryParams ExecutionLoadQueryParams,
				req Execution,
				rawReq any,
			) (server.Response[Execution], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[Execution]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Execution]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePutExecution(arguments, db, waitForChange, object)
				if err != nil {
					return server.Response[Execution]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Execution]{
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
			Execution{},
		)
		if err != nil {
			panic(err)
		}
		r.Put(putHandler.PathWithinRouter, putHandler.ServeHTTP)
	}()

	func() {
		patchHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/executions/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams ExecutionOnePathParams,
				queryParams ExecutionLoadQueryParams,
				req Execution,
				rawReq any,
			) (server.Response[Execution], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[Execution]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				forceSetValuesForFields := make([]string, 0)
				for _, possibleField := range maps.Keys(item) {
					if !slices.Contains(ExecutionTableColumns, possibleField) {
						continue
					}

					forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Execution]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePatchExecution(arguments, db, waitForChange, object, forceSetValuesForFields)
				if err != nil {
					return server.Response[Execution]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Execution]{
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
			Execution{},
		)
		if err != nil {
			panic(err)
		}
		r.Patch(patchHandler.PathWithinRouter, patchHandler.ServeHTTP)
	}()

	func() {
		deleteHandler, err := getHTTPHandler(
			http.MethodDelete,
			"/executions/{primaryKey}",
			http.StatusNoContent,
			func(
				ctx context.Context,
				pathParams ExecutionOnePathParams,
				queryParams ExecutionLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.EmptyResponse, error) {
				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				object := &Execution{}
				object.ID = pathParams.PrimaryKey

				err = handleDeleteExecution(arguments, db, waitForChange, object)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				return server.EmptyResponse{}, nil
			},
			Execution{},
		)
		if err != nil {
			panic(err)
		}
		r.Delete(deleteHandler.PathWithinRouter, deleteHandler.ServeHTTP)
	}()

	return r
}

func NewExecutionFromItem(item map[string]any) (any, error) {
	object := &Execution{}

	err := object.FromItem(item)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func init() {
	register(
		ExecutionTable,
		Execution{},
		NewExecutionFromItem,
		"/executions",
		GetExecutionRouter,
	)
}
