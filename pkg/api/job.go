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

type Job struct {
	ID                                      uuid.UUID          `json:"id"`
	CreatedAt                               time.Time          `json:"created_at"`
	UpdatedAt                               time.Time          `json:"updated_at"`
	DeletedAt                               *time.Time         `json:"deleted_at"`
	Name                                    string             `json:"name"`
	JobExecutorClaimedUntil                 time.Time          `json:"job_executor_claimed_until"`
	RuleID                                  uuid.UUID          `json:"rule_id"`
	RuleIDObject                            *Rule              `json:"rule_id_object"`
	BuildTaskID                             *uuid.UUID         `json:"build_task_id"`
	BuildTaskIDObject                       *Task              `json:"build_task_id_object"`
	TestTaskID                              *uuid.UUID         `json:"test_task_id"`
	TestTaskIDObject                        *Task              `json:"test_task_id_object"`
	PublishTaskID                           *uuid.UUID         `json:"publish_task_id"`
	PublishTaskIDObject                     *Task              `json:"publish_task_id_object"`
	DeployTaskID                            *uuid.UUID         `json:"deploy_task_id"`
	DeployTaskIDObject                      *Task              `json:"deploy_task_id_object"`
	ValidateTaskID                          *uuid.UUID         `json:"validate_task_id"`
	ValidateTaskIDObject                    *Task              `json:"validate_task_id_object"`
	ReferencedByExecutionJobIDObjects       []*Execution       `json:"referenced_by_execution_job_id_objects"`
	ReferencedByRuleRequiresJobJobIDObjects []*RuleRequiresJob `json:"referenced_by_rule_requires_job_job_id_objects"`
}

var JobTable = "job"

var JobTableNamespaceID int32 = 1337 + 3

var (
	JobTableIDColumn                      = "id"
	JobTableCreatedAtColumn               = "created_at"
	JobTableUpdatedAtColumn               = "updated_at"
	JobTableDeletedAtColumn               = "deleted_at"
	JobTableNameColumn                    = "name"
	JobTableJobExecutorClaimedUntilColumn = "job_executor_claimed_until"
	JobTableRuleIDColumn                  = "rule_id"
	JobTableBuildTaskIDColumn             = "build_task_id"
	JobTableTestTaskIDColumn              = "test_task_id"
	JobTablePublishTaskIDColumn           = "publish_task_id"
	JobTableDeployTaskIDColumn            = "deploy_task_id"
	JobTableValidateTaskIDColumn          = "validate_task_id"
)

var (
	JobTableIDColumnWithTypeCast                      = `"id" AS id`
	JobTableCreatedAtColumnWithTypeCast               = `"created_at" AS created_at`
	JobTableUpdatedAtColumnWithTypeCast               = `"updated_at" AS updated_at`
	JobTableDeletedAtColumnWithTypeCast               = `"deleted_at" AS deleted_at`
	JobTableNameColumnWithTypeCast                    = `"name" AS name`
	JobTableJobExecutorClaimedUntilColumnWithTypeCast = `"job_executor_claimed_until" AS job_executor_claimed_until`
	JobTableRuleIDColumnWithTypeCast                  = `"rule_id" AS rule_id`
	JobTableBuildTaskIDColumnWithTypeCast             = `"build_task_id" AS build_task_id`
	JobTableTestTaskIDColumnWithTypeCast              = `"test_task_id" AS test_task_id`
	JobTablePublishTaskIDColumnWithTypeCast           = `"publish_task_id" AS publish_task_id`
	JobTableDeployTaskIDColumnWithTypeCast            = `"deploy_task_id" AS deploy_task_id`
	JobTableValidateTaskIDColumnWithTypeCast          = `"validate_task_id" AS validate_task_id`
)

var JobTableColumns = []string{
	JobTableIDColumn,
	JobTableCreatedAtColumn,
	JobTableUpdatedAtColumn,
	JobTableDeletedAtColumn,
	JobTableNameColumn,
	JobTableJobExecutorClaimedUntilColumn,
	JobTableRuleIDColumn,
	JobTableBuildTaskIDColumn,
	JobTableTestTaskIDColumn,
	JobTablePublishTaskIDColumn,
	JobTableDeployTaskIDColumn,
	JobTableValidateTaskIDColumn,
}

var JobTableColumnsWithTypeCasts = []string{
	JobTableIDColumnWithTypeCast,
	JobTableCreatedAtColumnWithTypeCast,
	JobTableUpdatedAtColumnWithTypeCast,
	JobTableDeletedAtColumnWithTypeCast,
	JobTableNameColumnWithTypeCast,
	JobTableJobExecutorClaimedUntilColumnWithTypeCast,
	JobTableRuleIDColumnWithTypeCast,
	JobTableBuildTaskIDColumnWithTypeCast,
	JobTableTestTaskIDColumnWithTypeCast,
	JobTablePublishTaskIDColumnWithTypeCast,
	JobTableDeployTaskIDColumnWithTypeCast,
	JobTableValidateTaskIDColumnWithTypeCast,
}

var JobIntrospectedTable *introspect.Table

var JobTableColumnLookup map[string]*introspect.Column

var (
	JobTablePrimaryKeyColumn = JobTableIDColumn
)

func init() {
	JobIntrospectedTable = tableByName[JobTable]

	/* only needed during templating */
	if JobIntrospectedTable == nil {
		JobIntrospectedTable = &introspect.Table{}
	}

	JobTableColumnLookup = JobIntrospectedTable.ColumnByName
}

type JobOnePathParams struct {
	PrimaryKey uuid.UUID `json:"primaryKey"`
}

type JobLoadQueryParams struct {
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

func (m *Job) GetPrimaryKeyColumn() string {
	return JobTablePrimaryKeyColumn
}

func (m *Job) GetPrimaryKeyValue() any {
	return m.ID
}

func (m *Job) FromItem(item map[string]any) error {
	if item == nil {
		return fmt.Errorf(
			"item unexpectedly nil during JobFromItem",
		)
	}

	if len(item) == 0 {
		return fmt.Errorf(
			"item unexpectedly empty during JobFromItem",
		)
	}

	wrapError := func(k string, v any, err error) error {
		return fmt.Errorf("%v: %#+v; error; %v", k, v, err)
	}

	for k, v := range item {
		_, ok := JobTableColumnLookup[k]
		if !ok {
			return fmt.Errorf(
				"item contained unexpected key %#+v during JobFromItem; item: %#+v",
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

		case "name":
			if v == nil {
				continue
			}

			temp1, err := types.ParseString(v)
			if err != nil {
				return wrapError(k, v, err)
			}

			temp2, ok := temp1.(string)
			if !ok {
				if temp1 != nil {
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uuname.UUID", temp1))
				}
			}

			m.Name = temp2

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

		case "build_task_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uubuild_task_id.UUID", temp1))
				}
			}

			m.BuildTaskID = &temp2

		case "test_task_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uutest_task_id.UUID", temp1))
				}
			}

			m.TestTaskID = &temp2

		case "publish_task_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uupublish_task_id.UUID", temp1))
				}
			}

			m.PublishTaskID = &temp2

		case "deploy_task_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uudeploy_task_id.UUID", temp1))
				}
			}

			m.DeployTaskID = &temp2

		case "validate_task_id":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uuvalidate_task_id.UUID", temp1))
				}
			}

			m.ValidateTaskID = &temp2

		}
	}

	return nil
}

func (m *Job) Reload(ctx context.Context, tx pgx.Tx, includeDeleteds ...bool) error {
	extraWhere := ""
	if len(includeDeleteds) > 0 && includeDeleteds[0] {
		if slices.Contains(JobTableColumns, "deleted_at") {
			extraWhere = "\n    AND (deleted_at IS null OR deleted_at IS NOT null)"
		}
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	o, _, _, _, _, err := SelectJob(
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
	m.Name = o.Name
	m.JobExecutorClaimedUntil = o.JobExecutorClaimedUntil
	m.RuleID = o.RuleID
	m.RuleIDObject = o.RuleIDObject
	m.BuildTaskID = o.BuildTaskID
	m.BuildTaskIDObject = o.BuildTaskIDObject
	m.TestTaskID = o.TestTaskID
	m.TestTaskIDObject = o.TestTaskIDObject
	m.PublishTaskID = o.PublishTaskID
	m.PublishTaskIDObject = o.PublishTaskIDObject
	m.DeployTaskID = o.DeployTaskID
	m.DeployTaskIDObject = o.DeployTaskIDObject
	m.ValidateTaskID = o.ValidateTaskID
	m.ValidateTaskIDObject = o.ValidateTaskIDObject
	m.ReferencedByExecutionJobIDObjects = o.ReferencedByExecutionJobIDObjects
	m.ReferencedByRuleRequiresJobJobIDObjects = o.ReferencedByRuleRequiresJobJobIDObjects

	return nil
}

func (m *Job) Insert(ctx context.Context, tx pgx.Tx, setPrimaryKey bool, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setPrimaryKey && (setZeroValues || !types.IsZeroUUID(m.ID) || slices.Contains(forceSetValuesForFields, JobTableIDColumn) || isRequired(JobTableColumnLookup, JobTableIDColumn)) {
		columns = append(columns, JobTableIDColumn)

		v, err := types.FormatUUID(m.ID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, JobTableCreatedAtColumn) || isRequired(JobTableColumnLookup, JobTableCreatedAtColumn) {
		columns = append(columns, JobTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, JobTableUpdatedAtColumn) || isRequired(JobTableColumnLookup, JobTableUpdatedAtColumn) {
		columns = append(columns, JobTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, JobTableDeletedAtColumn) || isRequired(JobTableColumnLookup, JobTableDeletedAtColumn) {
		columns = append(columns, JobTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.Name) || slices.Contains(forceSetValuesForFields, JobTableNameColumn) || isRequired(JobTableColumnLookup, JobTableNameColumn) {
		columns = append(columns, JobTableNameColumn)

		v, err := types.FormatString(m.Name)
		if err != nil {
			return fmt.Errorf("failed to handle m.Name; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.JobExecutorClaimedUntil) || slices.Contains(forceSetValuesForFields, JobTableJobExecutorClaimedUntilColumn) || isRequired(JobTableColumnLookup, JobTableJobExecutorClaimedUntilColumn) {
		columns = append(columns, JobTableJobExecutorClaimedUntilColumn)

		v, err := types.FormatTime(m.JobExecutorClaimedUntil)
		if err != nil {
			return fmt.Errorf("failed to handle m.JobExecutorClaimedUntil; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.RuleID) || slices.Contains(forceSetValuesForFields, JobTableRuleIDColumn) || isRequired(JobTableColumnLookup, JobTableRuleIDColumn) {
		columns = append(columns, JobTableRuleIDColumn)

		v, err := types.FormatUUID(m.RuleID)
		if err != nil {
			return fmt.Errorf("failed to handle m.RuleID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.BuildTaskID) || slices.Contains(forceSetValuesForFields, JobTableBuildTaskIDColumn) || isRequired(JobTableColumnLookup, JobTableBuildTaskIDColumn) {
		columns = append(columns, JobTableBuildTaskIDColumn)

		v, err := types.FormatUUID(m.BuildTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.BuildTaskID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.TestTaskID) || slices.Contains(forceSetValuesForFields, JobTableTestTaskIDColumn) || isRequired(JobTableColumnLookup, JobTableTestTaskIDColumn) {
		columns = append(columns, JobTableTestTaskIDColumn)

		v, err := types.FormatUUID(m.TestTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.TestTaskID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.PublishTaskID) || slices.Contains(forceSetValuesForFields, JobTablePublishTaskIDColumn) || isRequired(JobTableColumnLookup, JobTablePublishTaskIDColumn) {
		columns = append(columns, JobTablePublishTaskIDColumn)

		v, err := types.FormatUUID(m.PublishTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.PublishTaskID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.DeployTaskID) || slices.Contains(forceSetValuesForFields, JobTableDeployTaskIDColumn) || isRequired(JobTableColumnLookup, JobTableDeployTaskIDColumn) {
		columns = append(columns, JobTableDeployTaskIDColumn)

		v, err := types.FormatUUID(m.DeployTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeployTaskID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.ValidateTaskID) || slices.Contains(forceSetValuesForFields, JobTableValidateTaskIDColumn) || isRequired(JobTableColumnLookup, JobTableValidateTaskIDColumn) {
		columns = append(columns, JobTableValidateTaskIDColumn)

		v, err := types.FormatUUID(m.ValidateTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ValidateTaskID; %v", err)
		}

		values = append(values, v)
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	item, err := query.Insert(
		ctx,
		tx,
		JobTable,
		columns,
		nil,
		false,
		false,
		JobTableColumns,
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to insert %#+v; %v", m, err)
	}
	v := (*item)[JobTableIDColumn]

	if v == nil {
		return fmt.Errorf("failed to find %v in %#+v", JobTableIDColumn, item)
	}

	wrapError := func(err error) error {
		return fmt.Errorf(
			"failed to treat %v: %#+v as uuid.UUID: %v",
			JobTableIDColumn,
			(*item)[JobTableIDColumn],
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

func (m *Job) Update(ctx context.Context, tx pgx.Tx, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, JobTableCreatedAtColumn) {
		columns = append(columns, JobTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, JobTableUpdatedAtColumn) {
		columns = append(columns, JobTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, JobTableDeletedAtColumn) {
		columns = append(columns, JobTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.Name) || slices.Contains(forceSetValuesForFields, JobTableNameColumn) {
		columns = append(columns, JobTableNameColumn)

		v, err := types.FormatString(m.Name)
		if err != nil {
			return fmt.Errorf("failed to handle m.Name; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.JobExecutorClaimedUntil) || slices.Contains(forceSetValuesForFields, JobTableJobExecutorClaimedUntilColumn) {
		columns = append(columns, JobTableJobExecutorClaimedUntilColumn)

		v, err := types.FormatTime(m.JobExecutorClaimedUntil)
		if err != nil {
			return fmt.Errorf("failed to handle m.JobExecutorClaimedUntil; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.RuleID) || slices.Contains(forceSetValuesForFields, JobTableRuleIDColumn) {
		columns = append(columns, JobTableRuleIDColumn)

		v, err := types.FormatUUID(m.RuleID)
		if err != nil {
			return fmt.Errorf("failed to handle m.RuleID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.BuildTaskID) || slices.Contains(forceSetValuesForFields, JobTableBuildTaskIDColumn) {
		columns = append(columns, JobTableBuildTaskIDColumn)

		v, err := types.FormatUUID(m.BuildTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.BuildTaskID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.TestTaskID) || slices.Contains(forceSetValuesForFields, JobTableTestTaskIDColumn) {
		columns = append(columns, JobTableTestTaskIDColumn)

		v, err := types.FormatUUID(m.TestTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.TestTaskID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.PublishTaskID) || slices.Contains(forceSetValuesForFields, JobTablePublishTaskIDColumn) {
		columns = append(columns, JobTablePublishTaskIDColumn)

		v, err := types.FormatUUID(m.PublishTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.PublishTaskID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.DeployTaskID) || slices.Contains(forceSetValuesForFields, JobTableDeployTaskIDColumn) {
		columns = append(columns, JobTableDeployTaskIDColumn)

		v, err := types.FormatUUID(m.DeployTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeployTaskID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroUUID(m.ValidateTaskID) || slices.Contains(forceSetValuesForFields, JobTableValidateTaskIDColumn) {
		columns = append(columns, JobTableValidateTaskIDColumn)

		v, err := types.FormatUUID(m.ValidateTaskID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ValidateTaskID; %v", err)
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
		JobTable,
		columns,
		fmt.Sprintf("%v = $$??", JobTableIDColumn),
		JobTableColumns,
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

func (m *Job) Delete(ctx context.Context, tx pgx.Tx, hardDeletes ...bool) error {
	hardDelete := false
	if len(hardDeletes) > 0 {
		hardDelete = hardDeletes[0]
	}

	if !hardDelete && slices.Contains(JobTableColumns, "deleted_at") {
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
		JobTable,
		fmt.Sprintf("%v = $$??", JobTableIDColumn),
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to delete %#+v; %v", m, err)
	}

	_ = m.Reload(ctx, tx, true)

	return nil
}

func (m *Job) LockTable(ctx context.Context, tx pgx.Tx, timeouts ...time.Duration) error {
	return query.LockTable(ctx, tx, JobTable, timeouts...)
}

func (m *Job) LockTableWithRetries(ctx context.Context, tx pgx.Tx, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.LockTableWithRetries(ctx, tx, JobTable, overallTimeout, individualAttempttimeout)
}

func (m *Job) AdvisoryLock(ctx context.Context, tx pgx.Tx, key int32, timeouts ...time.Duration) error {
	return query.AdvisoryLock(ctx, tx, JobTableNamespaceID, key, timeouts...)
}

func (m *Job) AdvisoryLockWithRetries(ctx context.Context, tx pgx.Tx, key int32, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.AdvisoryLockWithRetries(ctx, tx, JobTableNamespaceID, key, overallTimeout, individualAttempttimeout)
}

func SelectJobs(ctx context.Context, tx pgx.Tx, where string, orderBy *string, limit *int, offset *int, values ...any) ([]*Job, int64, int64, int64, int64, error) {
	before := time.Now()

	if config.Debug() {
		log.Printf("entered SelectJobs")

		defer func() {
			log.Printf("exited SelectJobs in %s", time.Since(before))
		}()
	}
	if slices.Contains(JobTableColumns, "deleted_at") {
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
	ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", JobTable, nil), !isLoadQuery)
	if !ok {
		return []*Job{}, 0, 0, 0, 0, nil
	}

	items, count, totalCount, page, totalPages, err := query.Select(
		ctx,
		tx,
		JobTableColumnsWithTypeCasts,
		JobTable,
		where,
		orderBy,
		limit,
		offset,
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectJobs; %v", err)
	}

	objects := make([]*Job, 0)

	for _, item := range *items {
		object := &Job{}

		err = object.FromItem(item)
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		if !types.IsZeroUUID(object.RuleID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", RuleTable, object.RuleID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectJobs->SelectRule for object.RuleIDObject")
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
					log.Printf("loaded SelectJobs->SelectRule for object.RuleIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.BuildTaskID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", TaskTable, object.BuildTaskID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectJobs->SelectTask for object.BuildTaskIDObject")
				}

				object.BuildTaskIDObject, _, _, _, _, err = SelectTask(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", TaskTablePrimaryKeyColumn),
					object.BuildTaskID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectJobs->SelectTask for object.BuildTaskIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.TestTaskID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", TaskTable, object.TestTaskID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectJobs->SelectTask for object.TestTaskIDObject")
				}

				object.TestTaskIDObject, _, _, _, _, err = SelectTask(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", TaskTablePrimaryKeyColumn),
					object.TestTaskID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectJobs->SelectTask for object.TestTaskIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.PublishTaskID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", TaskTable, object.PublishTaskID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectJobs->SelectTask for object.PublishTaskIDObject")
				}

				object.PublishTaskIDObject, _, _, _, _, err = SelectTask(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", TaskTablePrimaryKeyColumn),
					object.PublishTaskID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectJobs->SelectTask for object.PublishTaskIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.DeployTaskID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", TaskTable, object.DeployTaskID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectJobs->SelectTask for object.DeployTaskIDObject")
				}

				object.DeployTaskIDObject, _, _, _, _, err = SelectTask(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", TaskTablePrimaryKeyColumn),
					object.DeployTaskID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectJobs->SelectTask for object.DeployTaskIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		if !types.IsZeroUUID(object.ValidateTaskID) {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", TaskTable, object.ValidateTaskID), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectJobs->SelectTask for object.ValidateTaskIDObject")
				}

				object.ValidateTaskIDObject, _, _, _, _, err = SelectTask(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", TaskTablePrimaryKeyColumn),
					object.ValidateTaskID,
				)
				if err != nil {
					if !errors.Is(err, sql.ErrNoRows) {
						return nil, 0, 0, 0, 0, err
					}
				}

				if config.Debug() {
					log.Printf("loaded SelectJobs->SelectTask for object.ValidateTaskIDObject in %s", time.Since(thisBefore))
				}
			}
		}

		err = func() error {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("__ReferencedBy__%s{%v}", JobTable, object.GetPrimaryKeyValue()), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectJobs->SelectExecutions for object.ReferencedByExecutionJobIDObjects")
				}

				object.ReferencedByExecutionJobIDObjects, _, _, _, _, err = SelectExecutions(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", ExecutionTableJobIDColumn),
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
					log.Printf("loaded SelectJobs->SelectExecutions for object.ReferencedByExecutionJobIDObjects in %s", time.Since(thisBefore))
				}

			}

			return nil
		}()
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		err = func() error {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("__ReferencedBy__%s{%v}", JobTable, object.GetPrimaryKeyValue()), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectJobs->SelectRuleRequiresJobs for object.ReferencedByRuleRequiresJobJobIDObjects")
				}

				object.ReferencedByRuleRequiresJobJobIDObjects, _, _, _, _, err = SelectRuleRequiresJobs(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", RuleRequiresJobTableJobIDColumn),
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
					log.Printf("loaded SelectJobs->SelectRuleRequiresJobs for object.ReferencedByRuleRequiresJobJobIDObjects in %s", time.Since(thisBefore))
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

func SelectJob(ctx context.Context, tx pgx.Tx, where string, values ...any) (*Job, int64, int64, int64, int64, error) {
	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	objects, _, _, _, _, err := SelectJobs(
		ctx,
		tx,
		where,
		nil,
		helpers.Ptr(2),
		helpers.Ptr(0),
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectJob; %v", err)
	}

	if len(objects) > 1 {
		return nil, 0, 0, 0, 0, fmt.Errorf("attempt to call SelectJob returned more than 1 row")
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

func handleGetJobs(arguments *server.SelectManyArguments, db *pgxpool.Pool) ([]*Job, int64, int64, int64, int64, error) {
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

	objects, count, totalCount, page, totalPages, err := SelectJobs(arguments.Ctx, tx, arguments.Where, arguments.OrderBy, arguments.Limit, arguments.Offset, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return objects, count, totalCount, page, totalPages, nil
}

func handleGetJob(arguments *server.SelectOneArguments, db *pgxpool.Pool, primaryKey uuid.UUID) ([]*Job, int64, int64, int64, int64, error) {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	object, count, totalCount, page, totalPages, err := SelectJob(arguments.Ctx, tx, arguments.Where, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return []*Job{object}, count, totalCount, page, totalPages, nil
}

func handlePostJobs(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, objects []*Job, forceSetValuesForFieldsByObjectIndex [][]string) ([]*Job, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.INSERT}, JobTable, xid)
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

func handlePutJob(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Job) ([]*Job, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, JobTable, xid)
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

	return []*Job{object}, count, totalCount, page, totalPages, nil
}

func handlePatchJob(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Job, forceSetValuesForFields []string) ([]*Job, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, JobTable, xid)
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

	return []*Job{object}, count, totalCount, page, totalPages, nil
}

func handleDeleteJob(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Job) error {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.DELETE, stream.SOFT_DELETE}, JobTable, xid)
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

func GetJobRouter(db *pgxpool.Pool, redisPool *redis.Pool, httpMiddlewares []server.HTTPMiddleware, objectMiddlewares []server.ObjectMiddleware, waitForChange server.WaitForChange) chi.Router {
	r := chi.NewRouter()

	for _, m := range httpMiddlewares {
		r.Use(m)
	}

	func() {
		getManyHandler, err := getHTTPHandler(
			http.MethodGet,
			"/jobs",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams map[string]any,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[Job], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectManyArguments(ctx, queryParams, JobIntrospectedTable, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Job]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Job]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[Job]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[Job]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetJobs(arguments, db)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Job]{}, err
				}

				limit := int64(0)
				if arguments.Limit != nil {
					limit = int64(*arguments.Limit)
				}

				offset := int64(0)
				if arguments.Offset != nil {
					offset = int64(*arguments.Offset)
				}

				response := server.Response[Job]{
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

					return server.Response[Job]{}, err
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
			Job{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getManyHandler.PathWithinRouter, getManyHandler.ServeHTTP)
	}()

	func() {
		getOneHandler, err := getHTTPHandler(
			http.MethodGet,
			"/jobs/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams JobOnePathParams,
				queryParams JobLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[Job], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectOneArguments(ctx, queryParams.Depth, JobIntrospectedTable, pathParams.PrimaryKey, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Job]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Job]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[Job]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[Job]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetJob(arguments, db, pathParams.PrimaryKey)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Job]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				response := server.Response[Job]{
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

					return server.Response[Job]{}, err
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
			Job{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getOneHandler.PathWithinRouter, getOneHandler.ServeHTTP)
	}()

	func() {
		postHandler, err := getHTTPHandler(
			http.MethodPost,
			"/jobs",
			http.StatusCreated,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams JobLoadQueryParams,
				req []*Job,
				rawReq any,
			) (server.Response[Job], error) {
				allRawItems, ok := rawReq.([]any)
				if !ok {
					return server.Response[Job]{}, fmt.Errorf("failed to cast %#+v to []map[string]any", rawReq)
				}

				allItems := make([]map[string]any, 0)
				for _, rawItem := range allRawItems {
					item, ok := rawItem.(map[string]any)
					if !ok {
						return server.Response[Job]{}, fmt.Errorf("failed to cast %#+v to map[string]any", rawItem)
					}

					allItems = append(allItems, item)
				}

				forceSetValuesForFieldsByObjectIndex := make([][]string, 0)
				for _, item := range allItems {
					forceSetValuesForFields := make([]string, 0)
					for _, possibleField := range maps.Keys(item) {
						if !slices.Contains(JobTableColumns, possibleField) {
							continue
						}

						forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
					}
					forceSetValuesForFieldsByObjectIndex = append(forceSetValuesForFieldsByObjectIndex, forceSetValuesForFields)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Job]{}, err
				}

				objects, count, totalCount, _, _, err := handlePostJobs(arguments, db, waitForChange, req, forceSetValuesForFieldsByObjectIndex)
				if err != nil {
					return server.Response[Job]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Job]{
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
			Job{},
		)
		if err != nil {
			panic(err)
		}
		r.Post(postHandler.PathWithinRouter, postHandler.ServeHTTP)
	}()

	func() {
		putHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/jobs/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams JobOnePathParams,
				queryParams JobLoadQueryParams,
				req Job,
				rawReq any,
			) (server.Response[Job], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[Job]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Job]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePutJob(arguments, db, waitForChange, object)
				if err != nil {
					return server.Response[Job]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Job]{
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
			Job{},
		)
		if err != nil {
			panic(err)
		}
		r.Put(putHandler.PathWithinRouter, putHandler.ServeHTTP)
	}()

	func() {
		patchHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/jobs/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams JobOnePathParams,
				queryParams JobLoadQueryParams,
				req Job,
				rawReq any,
			) (server.Response[Job], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[Job]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				forceSetValuesForFields := make([]string, 0)
				for _, possibleField := range maps.Keys(item) {
					if !slices.Contains(JobTableColumns, possibleField) {
						continue
					}

					forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Job]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePatchJob(arguments, db, waitForChange, object, forceSetValuesForFields)
				if err != nil {
					return server.Response[Job]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Job]{
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
			Job{},
		)
		if err != nil {
			panic(err)
		}
		r.Patch(patchHandler.PathWithinRouter, patchHandler.ServeHTTP)
	}()

	func() {
		deleteHandler, err := getHTTPHandler(
			http.MethodDelete,
			"/jobs/{primaryKey}",
			http.StatusNoContent,
			func(
				ctx context.Context,
				pathParams JobOnePathParams,
				queryParams JobLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.EmptyResponse, error) {
				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				object := &Job{}
				object.ID = pathParams.PrimaryKey

				err = handleDeleteJob(arguments, db, waitForChange, object)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				return server.EmptyResponse{}, nil
			},
			Job{},
		)
		if err != nil {
			panic(err)
		}
		r.Delete(deleteHandler.PathWithinRouter, deleteHandler.ServeHTTP)
	}()

	return r
}

func NewJobFromItem(item map[string]any) (any, error) {
	object := &Job{}

	err := object.FromItem(item)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func init() {
	register(
		JobTable,
		Job{},
		NewJobFromItem,
		"/jobs",
		GetJobRouter,
	)
}
