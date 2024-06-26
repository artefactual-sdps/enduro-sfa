// Code generated by goa v3.15.2, DO NOT EDIT.
//
// package views
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package views

import (
	"github.com/google/uuid"
	goa "goa.design/goa/v3/pkg"
)

// EnduroStoredPackage is the viewed result type that is projected based on a
// view.
type EnduroStoredPackage struct {
	// Type to project
	Projected *EnduroStoredPackageView
	// View to render
	View string
}

// EnduroPackagePreservationActions is the viewed result type that is projected
// based on a view.
type EnduroPackagePreservationActions struct {
	// Type to project
	Projected *EnduroPackagePreservationActionsView
	// View to render
	View string
}

// EnduroStoredPackageView is a type that runs validations on a projected type.
type EnduroStoredPackageView struct {
	// Identifier of package
	ID *uint
	// Name of the package
	Name *string
	// Identifier of storage location
	LocationID *uuid.UUID
	// Status of the package
	Status *string
	// Identifier of processing workflow
	WorkflowID *string
	// Identifier of latest processing workflow run
	RunID *string
	// Identifier of AIP
	AipID *string
	// Creation datetime
	CreatedAt *string
	// Start datetime
	StartedAt *string
	// Completion datetime
	CompletedAt *string
}

// EnduroPackagePreservationActionsView is a type that runs validations on a
// projected type.
type EnduroPackagePreservationActionsView struct {
	Actions EnduroPackagePreservationActionCollectionView
}

// EnduroPackagePreservationActionCollectionView is a type that runs
// validations on a projected type.
type EnduroPackagePreservationActionCollectionView []*EnduroPackagePreservationActionView

// EnduroPackagePreservationActionView is a type that runs validations on a
// projected type.
type EnduroPackagePreservationActionView struct {
	ID          *uint
	WorkflowID  *string
	Type        *string
	Status      *string
	StartedAt   *string
	CompletedAt *string
	Tasks       EnduroPackagePreservationTaskCollectionView
	PackageID   *uint
}

// EnduroPackagePreservationTaskCollectionView is a type that runs validations
// on a projected type.
type EnduroPackagePreservationTaskCollectionView []*EnduroPackagePreservationTaskView

// EnduroPackagePreservationTaskView is a type that runs validations on a
// projected type.
type EnduroPackagePreservationTaskView struct {
	ID                   *uint
	TaskID               *string
	Name                 *string
	Status               *string
	StartedAt            *string
	CompletedAt          *string
	Note                 *string
	PreservationActionID *uint
}

var (
	// EnduroStoredPackageMap is a map indexing the attribute names of
	// EnduroStoredPackage by view name.
	EnduroStoredPackageMap = map[string][]string{
		"default": {
			"id",
			"name",
			"location_id",
			"status",
			"workflow_id",
			"run_id",
			"aip_id",
			"created_at",
			"started_at",
			"completed_at",
		},
	}
	// EnduroPackagePreservationActionsMap is a map indexing the attribute names of
	// EnduroPackagePreservationActions by view name.
	EnduroPackagePreservationActionsMap = map[string][]string{
		"default": {
			"actions",
		},
	}
	// EnduroPackagePreservationActionCollectionMap is a map indexing the attribute
	// names of EnduroPackagePreservationActionCollection by view name.
	EnduroPackagePreservationActionCollectionMap = map[string][]string{
		"simple": {
			"id",
			"workflow_id",
			"type",
			"status",
			"started_at",
			"completed_at",
			"package_id",
		},
		"default": {
			"id",
			"workflow_id",
			"type",
			"status",
			"started_at",
			"completed_at",
			"tasks",
			"package_id",
		},
	}
	// EnduroPackagePreservationActionMap is a map indexing the attribute names of
	// EnduroPackagePreservationAction by view name.
	EnduroPackagePreservationActionMap = map[string][]string{
		"simple": {
			"id",
			"workflow_id",
			"type",
			"status",
			"started_at",
			"completed_at",
			"package_id",
		},
		"default": {
			"id",
			"workflow_id",
			"type",
			"status",
			"started_at",
			"completed_at",
			"tasks",
			"package_id",
		},
	}
	// EnduroPackagePreservationTaskCollectionMap is a map indexing the attribute
	// names of EnduroPackagePreservationTaskCollection by view name.
	EnduroPackagePreservationTaskCollectionMap = map[string][]string{
		"default": {
			"id",
			"task_id",
			"name",
			"status",
			"started_at",
			"completed_at",
			"note",
			"preservation_action_id",
		},
	}
	// EnduroPackagePreservationTaskMap is a map indexing the attribute names of
	// EnduroPackagePreservationTask by view name.
	EnduroPackagePreservationTaskMap = map[string][]string{
		"default": {
			"id",
			"task_id",
			"name",
			"status",
			"started_at",
			"completed_at",
			"note",
			"preservation_action_id",
		},
	}
)

// ValidateEnduroStoredPackage runs the validations defined on the viewed
// result type EnduroStoredPackage.
func ValidateEnduroStoredPackage(result *EnduroStoredPackage) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroStoredPackageView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateEnduroPackagePreservationActions runs the validations defined on the
// viewed result type EnduroPackagePreservationActions.
func ValidateEnduroPackagePreservationActions(result *EnduroPackagePreservationActions) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroPackagePreservationActionsView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateEnduroStoredPackageView runs the validations defined on
// EnduroStoredPackageView using the "default" view.
func ValidateEnduroStoredPackageView(result *EnduroStoredPackageView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "new" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "unknown" || *result.Status == "queued" || *result.Status == "pending" || *result.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
	}
	if result.WorkflowID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.workflow_id", *result.WorkflowID, goa.FormatUUID))
	}
	if result.RunID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.run_id", *result.RunID, goa.FormatUUID))
	}
	if result.AipID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.aip_id", *result.AipID, goa.FormatUUID))
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	return
}

// ValidateEnduroPackagePreservationActionsView runs the validations defined on
// EnduroPackagePreservationActionsView using the "default" view.
func ValidateEnduroPackagePreservationActionsView(result *EnduroPackagePreservationActionsView) (err error) {

	if result.Actions != nil {
		if err2 := ValidateEnduroPackagePreservationActionCollectionView(result.Actions); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationActionCollectionViewSimple runs the
// validations defined on EnduroPackagePreservationActionCollectionView using
// the "simple" view.
func ValidateEnduroPackagePreservationActionCollectionViewSimple(result EnduroPackagePreservationActionCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateEnduroPackagePreservationActionViewSimple(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationActionCollectionView runs the validations
// defined on EnduroPackagePreservationActionCollectionView using the "default"
// view.
func ValidateEnduroPackagePreservationActionCollectionView(result EnduroPackagePreservationActionCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateEnduroPackagePreservationActionView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationActionViewSimple runs the validations
// defined on EnduroPackagePreservationActionView using the "simple" view.
func ValidateEnduroPackagePreservationActionViewSimple(result *EnduroPackagePreservationActionView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.WorkflowID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("workflow_id", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.StartedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("started_at", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "create-aip" || *result.Type == "create-and-review-aip" || *result.Type == "move-package") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []any{"create-aip", "create-and-review-aip", "move-package"}))
		}
	}
	if result.Status != nil {
		if !(*result.Status == "unspecified" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "queued" || *result.Status == "pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unspecified", "in progress", "done", "error", "queued", "pending"}))
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	return
}

// ValidateEnduroPackagePreservationActionView runs the validations defined on
// EnduroPackagePreservationActionView using the "default" view.
func ValidateEnduroPackagePreservationActionView(result *EnduroPackagePreservationActionView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.WorkflowID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("workflow_id", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.StartedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("started_at", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "create-aip" || *result.Type == "create-and-review-aip" || *result.Type == "move-package") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []any{"create-aip", "create-and-review-aip", "move-package"}))
		}
	}
	if result.Status != nil {
		if !(*result.Status == "unspecified" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "queued" || *result.Status == "pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unspecified", "in progress", "done", "error", "queued", "pending"}))
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	if result.Tasks != nil {
		if err2 := ValidateEnduroPackagePreservationTaskCollectionView(result.Tasks); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationTaskCollectionView runs the validations
// defined on EnduroPackagePreservationTaskCollectionView using the "default"
// view.
func ValidateEnduroPackagePreservationTaskCollectionView(result EnduroPackagePreservationTaskCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateEnduroPackagePreservationTaskView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroPackagePreservationTaskView runs the validations defined on
// EnduroPackagePreservationTaskView using the "default" view.
func ValidateEnduroPackagePreservationTaskView(result *EnduroPackagePreservationTaskView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.TaskID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("task_id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.StartedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("started_at", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "unspecified" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "queued" || *result.Status == "pending") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []any{"unspecified", "in progress", "done", "error", "queued", "pending"}))
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	return
}
