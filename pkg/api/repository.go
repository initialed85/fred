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

type Repository struct {
	ID                                    uuid.UUID  `json:"id"`
	CreatedAt                             time.Time  `json:"created_at"`
	UpdatedAt                             time.Time  `json:"updated_at"`
	DeletedAt                             *time.Time `json:"deleted_at"`
	LastSynced                            time.Time  `json:"last_synced"`
	URL                                   string     `json:"url"`
	Username                              *string    `json:"username"`
	Password                              *string    `json:"password"`
	SSHKey                                *string    `json:"ssh_key"`
	ReferencedByRuleRepositoryIDObjects   []*Rule    `json:"referenced_by_rule_repository_id_objects"`
	ReferencedByChangeRepositoryIDObjects []*Change  `json:"referenced_by_change_repository_id_objects"`
}

var RepositoryTable = "repository"

var RepositoryTableNamespaceID int32 = 1337 + 5

var (
	RepositoryTableIDColumn         = "id"
	RepositoryTableCreatedAtColumn  = "created_at"
	RepositoryTableUpdatedAtColumn  = "updated_at"
	RepositoryTableDeletedAtColumn  = "deleted_at"
	RepositoryTableLastSyncedColumn = "last_synced"
	RepositoryTableURLColumn        = "url"
	RepositoryTableUsernameColumn   = "username"
	RepositoryTablePasswordColumn   = "password"
	RepositoryTableSSHKeyColumn     = "ssh_key"
)

var (
	RepositoryTableIDColumnWithTypeCast         = `"id" AS id`
	RepositoryTableCreatedAtColumnWithTypeCast  = `"created_at" AS created_at`
	RepositoryTableUpdatedAtColumnWithTypeCast  = `"updated_at" AS updated_at`
	RepositoryTableDeletedAtColumnWithTypeCast  = `"deleted_at" AS deleted_at`
	RepositoryTableLastSyncedColumnWithTypeCast = `"last_synced" AS last_synced`
	RepositoryTableURLColumnWithTypeCast        = `"url" AS url`
	RepositoryTableUsernameColumnWithTypeCast   = `"username" AS username`
	RepositoryTablePasswordColumnWithTypeCast   = `"password" AS password`
	RepositoryTableSSHKeyColumnWithTypeCast     = `"ssh_key" AS ssh_key`
)

var RepositoryTableColumns = []string{
	RepositoryTableIDColumn,
	RepositoryTableCreatedAtColumn,
	RepositoryTableUpdatedAtColumn,
	RepositoryTableDeletedAtColumn,
	RepositoryTableLastSyncedColumn,
	RepositoryTableURLColumn,
	RepositoryTableUsernameColumn,
	RepositoryTablePasswordColumn,
	RepositoryTableSSHKeyColumn,
}

var RepositoryTableColumnsWithTypeCasts = []string{
	RepositoryTableIDColumnWithTypeCast,
	RepositoryTableCreatedAtColumnWithTypeCast,
	RepositoryTableUpdatedAtColumnWithTypeCast,
	RepositoryTableDeletedAtColumnWithTypeCast,
	RepositoryTableLastSyncedColumnWithTypeCast,
	RepositoryTableURLColumnWithTypeCast,
	RepositoryTableUsernameColumnWithTypeCast,
	RepositoryTablePasswordColumnWithTypeCast,
	RepositoryTableSSHKeyColumnWithTypeCast,
}

var RepositoryIntrospectedTable *introspect.Table

var RepositoryTableColumnLookup map[string]*introspect.Column

var (
	RepositoryTablePrimaryKeyColumn = RepositoryTableIDColumn
)

func init() {
	RepositoryIntrospectedTable = tableByName[RepositoryTable]

	/* only needed during templating */
	if RepositoryIntrospectedTable == nil {
		RepositoryIntrospectedTable = &introspect.Table{}
	}

	RepositoryTableColumnLookup = RepositoryIntrospectedTable.ColumnByName
}

type RepositoryOnePathParams struct {
	PrimaryKey uuid.UUID `json:"primaryKey"`
}

type RepositoryLoadQueryParams struct {
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

func (m *Repository) GetPrimaryKeyColumn() string {
	return RepositoryTablePrimaryKeyColumn
}

func (m *Repository) GetPrimaryKeyValue() any {
	return m.ID
}

func (m *Repository) FromItem(item map[string]any) error {
	if item == nil {
		return fmt.Errorf(
			"item unexpectedly nil during RepositoryFromItem",
		)
	}

	if len(item) == 0 {
		return fmt.Errorf(
			"item unexpectedly empty during RepositoryFromItem",
		)
	}

	wrapError := func(k string, v any, err error) error {
		return fmt.Errorf("%v: %#+v; error; %v", k, v, err)
	}

	for k, v := range item {
		_, ok := RepositoryTableColumnLookup[k]
		if !ok {
			return fmt.Errorf(
				"item contained unexpected key %#+v during RepositoryFromItem; item: %#+v",
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

		case "last_synced":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uulast_synced.UUID", temp1))
				}
			}

			m.LastSynced = temp2

		case "url":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uuurl.UUID", temp1))
				}
			}

			m.URL = temp2

		case "username":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uuusername.UUID", temp1))
				}
			}

			m.Username = &temp2

		case "password":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uupassword.UUID", temp1))
				}
			}

			m.Password = &temp2

		case "ssh_key":
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
					return wrapError(k, v, fmt.Errorf("failed to cast %#+v to uussh_key.UUID", temp1))
				}
			}

			m.SSHKey = &temp2

		}
	}

	return nil
}

func (m *Repository) Reload(ctx context.Context, tx pgx.Tx, includeDeleteds ...bool) error {
	extraWhere := ""
	if len(includeDeleteds) > 0 && includeDeleteds[0] {
		if slices.Contains(RepositoryTableColumns, "deleted_at") {
			extraWhere = "\n    AND (deleted_at IS null OR deleted_at IS NOT null)"
		}
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	o, _, _, _, _, err := SelectRepository(
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
	m.LastSynced = o.LastSynced
	m.URL = o.URL
	m.Username = o.Username
	m.Password = o.Password
	m.SSHKey = o.SSHKey
	m.ReferencedByRuleRepositoryIDObjects = o.ReferencedByRuleRepositoryIDObjects
	m.ReferencedByChangeRepositoryIDObjects = o.ReferencedByChangeRepositoryIDObjects

	return nil
}

func (m *Repository) Insert(ctx context.Context, tx pgx.Tx, setPrimaryKey bool, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setPrimaryKey && (setZeroValues || !types.IsZeroUUID(m.ID) || slices.Contains(forceSetValuesForFields, RepositoryTableIDColumn) || isRequired(RepositoryTableColumnLookup, RepositoryTableIDColumn)) {
		columns = append(columns, RepositoryTableIDColumn)

		v, err := types.FormatUUID(m.ID)
		if err != nil {
			return fmt.Errorf("failed to handle m.ID; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, RepositoryTableCreatedAtColumn) || isRequired(RepositoryTableColumnLookup, RepositoryTableCreatedAtColumn) {
		columns = append(columns, RepositoryTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, RepositoryTableUpdatedAtColumn) || isRequired(RepositoryTableColumnLookup, RepositoryTableUpdatedAtColumn) {
		columns = append(columns, RepositoryTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, RepositoryTableDeletedAtColumn) || isRequired(RepositoryTableColumnLookup, RepositoryTableDeletedAtColumn) {
		columns = append(columns, RepositoryTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.LastSynced) || slices.Contains(forceSetValuesForFields, RepositoryTableLastSyncedColumn) || isRequired(RepositoryTableColumnLookup, RepositoryTableLastSyncedColumn) {
		columns = append(columns, RepositoryTableLastSyncedColumn)

		v, err := types.FormatTime(m.LastSynced)
		if err != nil {
			return fmt.Errorf("failed to handle m.LastSynced; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.URL) || slices.Contains(forceSetValuesForFields, RepositoryTableURLColumn) || isRequired(RepositoryTableColumnLookup, RepositoryTableURLColumn) {
		columns = append(columns, RepositoryTableURLColumn)

		v, err := types.FormatString(m.URL)
		if err != nil {
			return fmt.Errorf("failed to handle m.URL; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.Username) || slices.Contains(forceSetValuesForFields, RepositoryTableUsernameColumn) || isRequired(RepositoryTableColumnLookup, RepositoryTableUsernameColumn) {
		columns = append(columns, RepositoryTableUsernameColumn)

		v, err := types.FormatString(m.Username)
		if err != nil {
			return fmt.Errorf("failed to handle m.Username; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.Password) || slices.Contains(forceSetValuesForFields, RepositoryTablePasswordColumn) || isRequired(RepositoryTableColumnLookup, RepositoryTablePasswordColumn) {
		columns = append(columns, RepositoryTablePasswordColumn)

		v, err := types.FormatString(m.Password)
		if err != nil {
			return fmt.Errorf("failed to handle m.Password; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.SSHKey) || slices.Contains(forceSetValuesForFields, RepositoryTableSSHKeyColumn) || isRequired(RepositoryTableColumnLookup, RepositoryTableSSHKeyColumn) {
		columns = append(columns, RepositoryTableSSHKeyColumn)

		v, err := types.FormatString(m.SSHKey)
		if err != nil {
			return fmt.Errorf("failed to handle m.SSHKey; %v", err)
		}

		values = append(values, v)
	}

	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	item, err := query.Insert(
		ctx,
		tx,
		RepositoryTable,
		columns,
		nil,
		false,
		false,
		RepositoryTableColumns,
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to insert %#+v; %v", m, err)
	}
	v := (*item)[RepositoryTableIDColumn]

	if v == nil {
		return fmt.Errorf("failed to find %v in %#+v", RepositoryTableIDColumn, item)
	}

	wrapError := func(err error) error {
		return fmt.Errorf(
			"failed to treat %v: %#+v as uuid.UUID: %v",
			RepositoryTableIDColumn,
			(*item)[RepositoryTableIDColumn],
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

func (m *Repository) Update(ctx context.Context, tx pgx.Tx, setZeroValues bool, forceSetValuesForFields ...string) error {
	columns := make([]string, 0)
	values := make([]any, 0)

	if setZeroValues || !types.IsZeroTime(m.CreatedAt) || slices.Contains(forceSetValuesForFields, RepositoryTableCreatedAtColumn) {
		columns = append(columns, RepositoryTableCreatedAtColumn)

		v, err := types.FormatTime(m.CreatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.CreatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.UpdatedAt) || slices.Contains(forceSetValuesForFields, RepositoryTableUpdatedAtColumn) {
		columns = append(columns, RepositoryTableUpdatedAtColumn)

		v, err := types.FormatTime(m.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.UpdatedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.DeletedAt) || slices.Contains(forceSetValuesForFields, RepositoryTableDeletedAtColumn) {
		columns = append(columns, RepositoryTableDeletedAtColumn)

		v, err := types.FormatTime(m.DeletedAt)
		if err != nil {
			return fmt.Errorf("failed to handle m.DeletedAt; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroTime(m.LastSynced) || slices.Contains(forceSetValuesForFields, RepositoryTableLastSyncedColumn) {
		columns = append(columns, RepositoryTableLastSyncedColumn)

		v, err := types.FormatTime(m.LastSynced)
		if err != nil {
			return fmt.Errorf("failed to handle m.LastSynced; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.URL) || slices.Contains(forceSetValuesForFields, RepositoryTableURLColumn) {
		columns = append(columns, RepositoryTableURLColumn)

		v, err := types.FormatString(m.URL)
		if err != nil {
			return fmt.Errorf("failed to handle m.URL; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.Username) || slices.Contains(forceSetValuesForFields, RepositoryTableUsernameColumn) {
		columns = append(columns, RepositoryTableUsernameColumn)

		v, err := types.FormatString(m.Username)
		if err != nil {
			return fmt.Errorf("failed to handle m.Username; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.Password) || slices.Contains(forceSetValuesForFields, RepositoryTablePasswordColumn) {
		columns = append(columns, RepositoryTablePasswordColumn)

		v, err := types.FormatString(m.Password)
		if err != nil {
			return fmt.Errorf("failed to handle m.Password; %v", err)
		}

		values = append(values, v)
	}

	if setZeroValues || !types.IsZeroString(m.SSHKey) || slices.Contains(forceSetValuesForFields, RepositoryTableSSHKeyColumn) {
		columns = append(columns, RepositoryTableSSHKeyColumn)

		v, err := types.FormatString(m.SSHKey)
		if err != nil {
			return fmt.Errorf("failed to handle m.SSHKey; %v", err)
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
		RepositoryTable,
		columns,
		fmt.Sprintf("%v = $$??", RepositoryTableIDColumn),
		RepositoryTableColumns,
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

func (m *Repository) Delete(ctx context.Context, tx pgx.Tx, hardDeletes ...bool) error {
	hardDelete := false
	if len(hardDeletes) > 0 {
		hardDelete = hardDeletes[0]
	}

	if !hardDelete && slices.Contains(RepositoryTableColumns, "deleted_at") {
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
		RepositoryTable,
		fmt.Sprintf("%v = $$??", RepositoryTableIDColumn),
		values...,
	)
	if err != nil {
		return fmt.Errorf("failed to delete %#+v; %v", m, err)
	}

	_ = m.Reload(ctx, tx, true)

	return nil
}

func (m *Repository) LockTable(ctx context.Context, tx pgx.Tx, timeouts ...time.Duration) error {
	return query.LockTable(ctx, tx, RepositoryTable, timeouts...)
}

func (m *Repository) LockTableWithRetries(ctx context.Context, tx pgx.Tx, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.LockTableWithRetries(ctx, tx, RepositoryTable, overallTimeout, individualAttempttimeout)
}

func (m *Repository) AdvisoryLock(ctx context.Context, tx pgx.Tx, key int32, timeouts ...time.Duration) error {
	return query.AdvisoryLock(ctx, tx, RepositoryTableNamespaceID, key, timeouts...)
}

func (m *Repository) AdvisoryLockWithRetries(ctx context.Context, tx pgx.Tx, key int32, overallTimeout time.Duration, individualAttempttimeout time.Duration) error {
	return query.AdvisoryLockWithRetries(ctx, tx, RepositoryTableNamespaceID, key, overallTimeout, individualAttempttimeout)
}

func SelectRepositories(ctx context.Context, tx pgx.Tx, where string, orderBy *string, limit *int, offset *int, values ...any) ([]*Repository, int64, int64, int64, int64, error) {
	before := time.Now()

	if config.Debug() {
		log.Printf("entered SelectRepositories")

		defer func() {
			log.Printf("exited SelectRepositories in %s", time.Since(before))
		}()
	}
	if slices.Contains(RepositoryTableColumns, "deleted_at") {
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
	ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("%s{%v}", RepositoryTable, nil), !isLoadQuery)
	if !ok {
		return []*Repository{}, 0, 0, 0, 0, nil
	}

	items, count, totalCount, page, totalPages, err := query.Select(
		ctx,
		tx,
		RepositoryTableColumnsWithTypeCasts,
		RepositoryTable,
		where,
		orderBy,
		limit,
		offset,
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectRepositorys; %v", err)
	}

	objects := make([]*Repository, 0)

	for _, item := range *items {
		object := &Repository{}

		err = object.FromItem(item)
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		err = func() error {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("__ReferencedBy__%s{%v}", RepositoryTable, object.GetPrimaryKeyValue()), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectRepositories->SelectRules for object.ReferencedByRuleRepositoryIDObjects")
				}

				object.ReferencedByRuleRepositoryIDObjects, _, _, _, _, err = SelectRules(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", RuleTableRepositoryIDColumn),
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
					log.Printf("loaded SelectRepositories->SelectRules for object.ReferencedByRuleRepositoryIDObjects in %s", time.Since(thisBefore))
				}

			}

			return nil
		}()
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}

		err = func() error {
			ctx, ok := query.HandleQueryPathGraphCycles(ctx, fmt.Sprintf("__ReferencedBy__%s{%v}", RepositoryTable, object.GetPrimaryKeyValue()), true)
			if ok {
				thisBefore := time.Now()

				if config.Debug() {
					log.Printf("loading SelectRepositories->SelectChanges for object.ReferencedByChangeRepositoryIDObjects")
				}

				object.ReferencedByChangeRepositoryIDObjects, _, _, _, _, err = SelectChanges(
					ctx,
					tx,
					fmt.Sprintf("%v = $1", ChangeTableRepositoryIDColumn),
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
					log.Printf("loaded SelectRepositories->SelectChanges for object.ReferencedByChangeRepositoryIDObjects in %s", time.Since(thisBefore))
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

func SelectRepository(ctx context.Context, tx pgx.Tx, where string, values ...any) (*Repository, int64, int64, int64, int64, error) {
	ctx, cleanup := query.WithQueryID(ctx)
	defer cleanup()

	ctx = query.WithMaxDepth(ctx, nil)

	objects, _, _, _, _, err := SelectRepositories(
		ctx,
		tx,
		where,
		nil,
		helpers.Ptr(2),
		helpers.Ptr(0),
		values...,
	)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf("failed to call SelectRepository; %v", err)
	}

	if len(objects) > 1 {
		return nil, 0, 0, 0, 0, fmt.Errorf("attempt to call SelectRepository returned more than 1 row")
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

func handleGetRepositories(arguments *server.SelectManyArguments, db *pgxpool.Pool) ([]*Repository, int64, int64, int64, int64, error) {
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

	objects, count, totalCount, page, totalPages, err := SelectRepositories(arguments.Ctx, tx, arguments.Where, arguments.OrderBy, arguments.Limit, arguments.Offset, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return objects, count, totalCount, page, totalPages, nil
}

func handleGetRepository(arguments *server.SelectOneArguments, db *pgxpool.Pool, primaryKey uuid.UUID) ([]*Repository, int64, int64, int64, int64, error) {
	tx, err := db.Begin(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	defer func() {
		_ = tx.Rollback(arguments.Ctx)
	}()

	object, count, totalCount, page, totalPages, err := SelectRepository(arguments.Ctx, tx, arguments.Where, arguments.Values...)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err = tx.Commit(arguments.Ctx)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	return []*Repository{object}, count, totalCount, page, totalPages, nil
}

func handlePostRepositorys(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, objects []*Repository, forceSetValuesForFieldsByObjectIndex [][]string) ([]*Repository, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.INSERT}, RepositoryTable, xid)
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

func handlePutRepository(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Repository) ([]*Repository, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, RepositoryTable, xid)
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

	return []*Repository{object}, count, totalCount, page, totalPages, nil
}

func handlePatchRepository(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Repository, forceSetValuesForFields []string) ([]*Repository, int64, int64, int64, int64, error) {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.UPDATE, stream.SOFT_DELETE, stream.SOFT_RESTORE, stream.SOFT_UPDATE}, RepositoryTable, xid)
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

	return []*Repository{object}, count, totalCount, page, totalPages, nil
}

func handleDeleteRepository(arguments *server.LoadArguments, db *pgxpool.Pool, waitForChange server.WaitForChange, object *Repository) error {
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
		_, err := waitForChange(arguments.Ctx, []stream.Action{stream.DELETE, stream.SOFT_DELETE}, RepositoryTable, xid)
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

func GetRepositoryRouter(db *pgxpool.Pool, redisPool *redis.Pool, httpMiddlewares []server.HTTPMiddleware, objectMiddlewares []server.ObjectMiddleware, waitForChange server.WaitForChange) chi.Router {
	r := chi.NewRouter()

	for _, m := range httpMiddlewares {
		r.Use(m)
	}

	func() {
		getManyHandler, err := getHTTPHandler(
			http.MethodGet,
			"/repositories",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams map[string]any,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[Repository], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectManyArguments(ctx, queryParams, RepositoryIntrospectedTable, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Repository]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Repository]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[Repository]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[Repository]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetRepositories(arguments, db)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Repository]{}, err
				}

				limit := int64(0)
				if arguments.Limit != nil {
					limit = int64(*arguments.Limit)
				}

				offset := int64(0)
				if arguments.Offset != nil {
					offset = int64(*arguments.Offset)
				}

				response := server.Response[Repository]{
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

					return server.Response[Repository]{}, err
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
			Repository{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getManyHandler.PathWithinRouter, getManyHandler.ServeHTTP)
	}()

	func() {
		getOneHandler, err := getHTTPHandler(
			http.MethodGet,
			"/repositories/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams RepositoryOnePathParams,
				queryParams RepositoryLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.Response[Repository], error) {
				before := time.Now()

				redisConn := redisPool.Get()
				defer func() {
					_ = redisConn.Close()
				}()

				arguments, err := server.GetSelectOneArguments(ctx, queryParams.Depth, RepositoryIntrospectedTable, pathParams.PrimaryKey, nil, nil)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache not yet reached; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Repository]{}, err
				}

				cachedResponseAsJSON, cacheHit, err := server.GetCachedResponseAsJSON(arguments.RequestHash, redisConn)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache failed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Repository]{}, err
				}

				if cacheHit {
					var cachedResponse server.Response[Repository]

					/* TODO: it'd be nice to be able to avoid this (i.e. just pass straight through) */
					err = json.Unmarshal(cachedResponseAsJSON, &cachedResponse)
					if err != nil {
						if config.Debug() {
							log.Printf("request cache hit but failed unmarshal; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
						}

						return server.Response[Repository]{}, err
					}

					if config.Debug() {
						log.Printf("request cache hit; request succeeded in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return cachedResponse, nil
				}

				objects, count, totalCount, _, _, err := handleGetRepository(arguments, db, pathParams.PrimaryKey)
				if err != nil {
					if config.Debug() {
						log.Printf("request cache missed; request failed in %s %s path: %#+v query: %#+v req: %#+v", time.Since(before), http.MethodGet, pathParams, queryParams, req)
					}

					return server.Response[Repository]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				response := server.Response[Repository]{
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

					return server.Response[Repository]{}, err
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
			Repository{},
		)
		if err != nil {
			panic(err)
		}
		r.Get(getOneHandler.PathWithinRouter, getOneHandler.ServeHTTP)
	}()

	func() {
		postHandler, err := getHTTPHandler(
			http.MethodPost,
			"/repositories",
			http.StatusCreated,
			func(
				ctx context.Context,
				pathParams server.EmptyPathParams,
				queryParams RepositoryLoadQueryParams,
				req []*Repository,
				rawReq any,
			) (server.Response[Repository], error) {
				allRawItems, ok := rawReq.([]any)
				if !ok {
					return server.Response[Repository]{}, fmt.Errorf("failed to cast %#+v to []map[string]any", rawReq)
				}

				allItems := make([]map[string]any, 0)
				for _, rawItem := range allRawItems {
					item, ok := rawItem.(map[string]any)
					if !ok {
						return server.Response[Repository]{}, fmt.Errorf("failed to cast %#+v to map[string]any", rawItem)
					}

					allItems = append(allItems, item)
				}

				forceSetValuesForFieldsByObjectIndex := make([][]string, 0)
				for _, item := range allItems {
					forceSetValuesForFields := make([]string, 0)
					for _, possibleField := range maps.Keys(item) {
						if !slices.Contains(RepositoryTableColumns, possibleField) {
							continue
						}

						forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
					}
					forceSetValuesForFieldsByObjectIndex = append(forceSetValuesForFieldsByObjectIndex, forceSetValuesForFields)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Repository]{}, err
				}

				objects, count, totalCount, _, _, err := handlePostRepositorys(arguments, db, waitForChange, req, forceSetValuesForFieldsByObjectIndex)
				if err != nil {
					return server.Response[Repository]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Repository]{
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
			Repository{},
		)
		if err != nil {
			panic(err)
		}
		r.Post(postHandler.PathWithinRouter, postHandler.ServeHTTP)
	}()

	func() {
		putHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/repositories/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams RepositoryOnePathParams,
				queryParams RepositoryLoadQueryParams,
				req Repository,
				rawReq any,
			) (server.Response[Repository], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[Repository]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Repository]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePutRepository(arguments, db, waitForChange, object)
				if err != nil {
					return server.Response[Repository]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Repository]{
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
			Repository{},
		)
		if err != nil {
			panic(err)
		}
		r.Put(putHandler.PathWithinRouter, putHandler.ServeHTTP)
	}()

	func() {
		patchHandler, err := getHTTPHandler(
			http.MethodPatch,
			"/repositories/{primaryKey}",
			http.StatusOK,
			func(
				ctx context.Context,
				pathParams RepositoryOnePathParams,
				queryParams RepositoryLoadQueryParams,
				req Repository,
				rawReq any,
			) (server.Response[Repository], error) {
				item, ok := rawReq.(map[string]any)
				if !ok {
					return server.Response[Repository]{}, fmt.Errorf("failed to cast %#+v to map[string]any", item)
				}

				forceSetValuesForFields := make([]string, 0)
				for _, possibleField := range maps.Keys(item) {
					if !slices.Contains(RepositoryTableColumns, possibleField) {
						continue
					}

					forceSetValuesForFields = append(forceSetValuesForFields, possibleField)
				}

				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.Response[Repository]{}, err
				}

				object := &req
				object.ID = pathParams.PrimaryKey

				objects, count, totalCount, _, _, err := handlePatchRepository(arguments, db, waitForChange, object, forceSetValuesForFields)
				if err != nil {
					return server.Response[Repository]{}, err
				}

				limit := int64(0)

				offset := int64(0)

				return server.Response[Repository]{
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
			Repository{},
		)
		if err != nil {
			panic(err)
		}
		r.Patch(patchHandler.PathWithinRouter, patchHandler.ServeHTTP)
	}()

	func() {
		deleteHandler, err := getHTTPHandler(
			http.MethodDelete,
			"/repositories/{primaryKey}",
			http.StatusNoContent,
			func(
				ctx context.Context,
				pathParams RepositoryOnePathParams,
				queryParams RepositoryLoadQueryParams,
				req server.EmptyRequest,
				rawReq any,
			) (server.EmptyResponse, error) {
				arguments, err := server.GetLoadArguments(ctx, queryParams.Depth)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				object := &Repository{}
				object.ID = pathParams.PrimaryKey

				err = handleDeleteRepository(arguments, db, waitForChange, object)
				if err != nil {
					return server.EmptyResponse{}, err
				}

				return server.EmptyResponse{}, nil
			},
			Repository{},
		)
		if err != nil {
			panic(err)
		}
		r.Delete(deleteHandler.PathWithinRouter, deleteHandler.ServeHTTP)
	}()

	return r
}

func NewRepositoryFromItem(item map[string]any) (any, error) {
	object := &Repository{}

	err := object.FromItem(item)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func init() {
	register(
		RepositoryTable,
		Repository{},
		NewRepositoryFromItem,
		"/repositories",
		GetRepositoryRouter,
	)
}
