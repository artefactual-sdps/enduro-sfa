package package_

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	goapackage "github.com/artefactual-sdps/enduro/internal/api/gen/package_"
)

type PreservationActionType uint

const (
	ActionTypeUnspecified PreservationActionType = iota
	ActionTypeCreateAIP
	ActionTypeMovePackage
)

func NewPreservationActionType(status string) PreservationActionType {
	var s PreservationActionType

	switch strings.ToLower(status) {
	case "create-aip":
		s = ActionTypeCreateAIP
	case "move-package":
		s = ActionTypeMovePackage
	default:
		s = ActionTypeUnspecified
	}

	return s
}

func (p PreservationActionType) String() string {
	switch p {
	case ActionTypeCreateAIP:
		return "create-aip"
	case ActionTypeMovePackage:
		return "move-package"
	}

	return "unspecified"
}

func (p PreservationActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PreservationActionType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*p = NewPreservationActionType(s)

	return nil
}

type PreservationActionStatus uint

const (
	ActionStatusUnspecified PreservationActionStatus = iota
	ActionStatusInProgress
	ActionStatusDone
	ActionStatusError
	ActionStatusQueued
	ActionStatusPending
)

func NewPreservationActionStatus(status string) PreservationActionStatus {
	var s PreservationActionStatus

	switch strings.ToLower(status) {
	case "in progress":
		s = ActionStatusInProgress
	case "done":
		s = ActionStatusDone
	case "error":
		s = ActionStatusError
	case "queued":
		s = ActionStatusQueued
	case "pending":
		s = ActionStatusPending
	default:
		s = ActionStatusUnspecified
	}

	return s
}

func (p PreservationActionStatus) String() string {
	switch p {
	case ActionStatusInProgress:
		return "in progress"
	case ActionStatusDone:
		return "done"
	case ActionStatusError:
		return "error"
	case ActionStatusQueued:
		return "queued"
	case ActionStatusPending:
		return "pending"
	}

	return "unspecified"
}

func (p PreservationActionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PreservationActionStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*p = NewPreservationActionStatus(s)

	return nil
}

// PreservationAction represents a preservation action in the preservation_action table.
type PreservationAction struct {
	ID          uint                     `db:"id"`
	WorkflowID  string                   `db:"workflow_id"`
	Type        PreservationActionType   `db:"type"`
	Status      PreservationActionStatus `db:"status"`
	StartedAt   sql.NullTime             `db:"started_at"`
	CompletedAt sql.NullTime             `db:"completed_at"`
	PackageID   uint                     `db:"package_id"`
}

type PreservationTaskStatus uint

const (
	TaskStatusUnspecified PreservationTaskStatus = iota
	TaskStatusInProgress
	TaskStatusDone
	TaskStatusError
	TaskStatusQueued
	TaskStatusPending
)

func NewPreservationTaskStatus(status string) PreservationTaskStatus {
	var s PreservationTaskStatus

	switch strings.ToLower(status) {
	case "in progress":
		s = TaskStatusInProgress
	case "done":
		s = TaskStatusDone
	case "error":
		s = TaskStatusError
	case "queued":
		s = TaskStatusQueued
	case "pending":
		s = TaskStatusPending
	default:
		s = TaskStatusUnspecified
	}

	return s
}

func (p PreservationTaskStatus) String() string {
	switch p {
	case TaskStatusInProgress:
		return "in progress"
	case TaskStatusDone:
		return "done"
	case TaskStatusError:
		return "error"
	case TaskStatusQueued:
		return "queued"
	case TaskStatusPending:
		return "pending"
	}

	return "unspecified"
}

func (p PreservationTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *PreservationTaskStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*p = NewPreservationTaskStatus(s)

	return nil
}

// PreservationTask represents a preservation action task in the preservation_task table.
type PreservationTask struct {
	ID                   uint                   `db:"id"`
	TaskID               string                 `db:"task_id"`
	Name                 string                 `db:"name"`
	Status               PreservationTaskStatus `db:"status"`
	StartedAt            sql.NullTime           `db:"started_at"`
	CompletedAt          sql.NullTime           `db:"completed_at"`
	PreservationActionID uint                   `db:"preservation_action_id"`
}

func (w *goaWrapper) PreservationActions(ctx context.Context, payload *goapackage.PreservationActionsPayload) (*goapackage.EnduroPackagePreservationActions, error) {
	goapkg, err := w.Show(ctx, &goapackage.ShowPayload{ID: payload.ID})
	if err != nil {
		return nil, err
	}

	query := "SELECT id, workflow_id, type, status, CONVERT_TZ(started_at, @@session.time_zone, '+00:00') AS started_at, CONVERT_TZ(completed_at, @@session.time_zone, '+00:00') AS completed_at FROM preservation_action WHERE package_id = ? ORDER BY started_at DESC"
	args := []interface{}{goapkg.ID}

	rows, err := w.db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying the database: %w", err)
	}
	defer rows.Close()

	preservation_actions := []*goapackage.EnduroPackagePreservationAction{}
	for rows.Next() {
		pa := PreservationAction{}
		if err := rows.StructScan(&pa); err != nil {
			return nil, fmt.Errorf("error scanning database result: %w", err)
		}
		goapa := &goapackage.EnduroPackagePreservationAction{
			ID:          pa.ID,
			WorkflowID:  pa.WorkflowID,
			Type:        pa.Type.String(),
			Status:      pa.Status.String(),
			StartedAt:   formatTime(pa.StartedAt.Time),
			CompletedAt: formatOptionalTime(pa.CompletedAt),
		}

		ptQuery := "SELECT id, task_id, name, status, CONVERT_TZ(started_at, @@session.time_zone, '+00:00') AS started_at, CONVERT_TZ(completed_at, @@session.time_zone, '+00:00') AS completed_at FROM preservation_task WHERE preservation_action_id = ?"
		ptQueryArgs := []interface{}{pa.ID}

		ptRows, err := w.db.QueryxContext(ctx, ptQuery, ptQueryArgs...)
		if err != nil {
			return nil, fmt.Errorf("error querying the database: %w", err)
		}
		defer ptRows.Close()

		preservation_tasks := []*goapackage.EnduroPackagePreservationTask{}
		for ptRows.Next() {
			pt := PreservationTask{}
			if err := ptRows.StructScan(&pt); err != nil {
				return nil, fmt.Errorf("error scanning database result: %w", err)
			}
			goapt := &goapackage.EnduroPackagePreservationTask{
				ID:          pt.ID,
				TaskID:      pt.TaskID,
				Name:        pt.Name,
				Status:      pt.Status.String(),
				StartedAt:   formatTime(pt.StartedAt.Time),
				CompletedAt: formatOptionalTime(pt.CompletedAt),
			}
			preservation_tasks = append(preservation_tasks, goapt)
		}

		goapa.Tasks = preservation_tasks
		preservation_actions = append(preservation_actions, goapa)
	}

	result := &goapackage.EnduroPackagePreservationActions{
		Actions: preservation_actions,
	}

	return result, nil
}

func (svc *packageImpl) CreatePreservationAction(ctx context.Context, pa *PreservationAction) error {
	startedAt := &pa.StartedAt.Time
	completedAt := &pa.CompletedAt.Time
	if pa.StartedAt.Time.IsZero() {
		startedAt = nil
	}
	if pa.CompletedAt.Time.IsZero() {
		completedAt = nil
	}

	query := `INSERT INTO preservation_action (workflow_id, type, status, started_at, completed_at, package_id) VALUES (?, ?, ?, ?, ?, ?)`
	args := []interface{}{
		pa.WorkflowID,
		pa.Type,
		pa.Status,
		startedAt,
		completedAt,
		pa.PackageID,
	}

	res, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error inserting preservation action: %w", err)
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return fmt.Errorf("error retrieving insert ID: %w", err)
	}

	pa.ID = uint(id)

	// publishEvent(ctx, svc.events, EventTypePackageUpdated, pa.PackageID)

	return nil
}

func (svc *packageImpl) SetPreservationActionStatus(ctx context.Context, ID uint, status PreservationActionStatus) error {
	query := `UPDATE preservation_action SET status = ? WHERE id = ?`
	args := []interface{}{
		status,
		ID,
	}

	_, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error updating preservation action: %w", err)
	}

	return nil
}

func (svc *packageImpl) CompletePreservationAction(ctx context.Context, ID uint, status PreservationActionStatus, completedAt time.Time) error {
	query := `UPDATE preservation_action SET status = ?, completed_at = ? WHERE id = ?`
	args := []interface{}{
		status,
		completedAt,
		ID,
	}

	_, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error updating preservation action: %w", err)
	}

	return nil
}

func (svc *packageImpl) CreatePreservationTask(ctx context.Context, pt *PreservationTask) error {
	startedAt := &pt.StartedAt.Time
	completedAt := &pt.CompletedAt.Time
	if pt.StartedAt.Time.IsZero() {
		startedAt = nil
	}
	if pt.CompletedAt.Time.IsZero() {
		completedAt = nil
	}

	query := `INSERT INTO preservation_task (task_id, name, status, started_at, completed_at, preservation_action_id) VALUES (?, ?, ?, ?, ?, ?)`
	args := []interface{}{
		pt.TaskID,
		pt.Name,
		pt.Status,
		startedAt,
		completedAt,
		pt.PreservationActionID,
	}

	res, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error inserting preservation task: %w", err)
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return fmt.Errorf("error retrieving insert ID: %w", err)
	}

	pt.ID = uint(id)

	return nil
}

func (svc *packageImpl) CompletePreservationTask(ctx context.Context, ID uint, name *string, status PreservationTaskStatus, completedAt time.Time) error {
	var query string
	args := []interface{}{}

	if name != nil {
		query = `UPDATE preservation_task SET name = ?, status = ?, completed_at = ? WHERE id = ?`
		args = append(args, name, status, completedAt, ID)
	} else {
		query = `UPDATE preservation_task SET status = ?, completed_at = ? WHERE id = ?`
		args = append(args, status, completedAt, ID)
	}

	_, err := svc.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error updating preservation task: %w", err)
	}

	return nil
}
