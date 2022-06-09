// Code generated by goa v3.5.5, DO NOT EDIT.
//
// package service
//
// Command:
// $ goa-v3.5.5 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package package_

import (
	"context"

	package_views "github.com/artefactual-labs/enduro/internal/api/gen/package_/views"
	goa "goa.design/goa/v3/pkg"
)

// The package service manages packages being transferred to a3m.
type Service interface {
	// Monitor implements monitor.
	Monitor(context.Context, MonitorServerStream) (err error)
	// List all stored packages
	List(context.Context, *ListPayload) (res *ListResult, err error)
	// Show package by ID
	Show(context.Context, *ShowPayload) (res *EnduroStoredPackage, err error)
	// Delete package by ID
	Delete(context.Context, *DeletePayload) (err error)
	// Cancel package processing by ID
	Cancel(context.Context, *CancelPayload) (err error)
	// Retry package processing by ID
	Retry(context.Context, *RetryPayload) (err error)
	// Retrieve workflow status by ID
	Workflow(context.Context, *WorkflowPayload) (res *EnduroPackageWorkflowStatus, err error)
	// Download package by ID
	Download(context.Context, *DownloadPayload) (res []byte, err error)
	// Bulk operations (retry, cancel...).
	Bulk(context.Context, *BulkPayload) (res *BulkResult, err error)
	// Retrieve status of current bulk operation.
	BulkStatus(context.Context) (res *BulkStatusResult, err error)
	// List all preservation actions by ID
	PreservationActions(context.Context, *PreservationActionsPayload) (res *EnduroPackagePreservationActions, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "package"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [11]string{"monitor", "list", "show", "delete", "cancel", "retry", "workflow", "download", "bulk", "bulk_status", "preservation-actions"}

// MonitorServerStream is the interface a "monitor" endpoint server stream must
// satisfy.
type MonitorServerStream interface {
	// Send streams instances of "EnduroMonitorUpdate".
	Send(*EnduroMonitorUpdate) error
	// Close closes the stream.
	Close() error
}

// MonitorClientStream is the interface a "monitor" endpoint client stream must
// satisfy.
type MonitorClientStream interface {
	// Recv reads instances of "EnduroMonitorUpdate" from the stream.
	Recv() (*EnduroMonitorUpdate, error)
}

// EnduroMonitorUpdate is the result type of the package service monitor method.
type EnduroMonitorUpdate struct {
	// Identifier of package
	ID uint
	// Type of the event
	Type string
	// Package
	Item *EnduroStoredPackage
}

// ListPayload is the payload type of the package service list method.
type ListPayload struct {
	Name                *string
	AipID               *string
	EarliestCreatedTime *string
	LatestCreatedTime   *string
	Status              *string
	// Pagination cursor
	Cursor *string
}

// ListResult is the result type of the package service list method.
type ListResult struct {
	Items      EnduroStoredPackageCollection
	NextCursor *string
}

// ShowPayload is the payload type of the package service show method.
type ShowPayload struct {
	// Identifier of package to show
	ID uint
}

// EnduroStoredPackage is the result type of the package service show method.
type EnduroStoredPackage struct {
	// Identifier of package
	ID uint
	// Name of the package
	Name *string
	// Status of the package
	Status string
	// Identifier of processing workflow
	WorkflowID *string
	// Identifier of latest processing workflow run
	RunID *string
	// Identifier of Archivematica AIP
	AipID *string
	// Creation datetime
	CreatedAt string
	// Start datetime
	StartedAt *string
	// Completion datetime
	CompletedAt *string
}

// DeletePayload is the payload type of the package service delete method.
type DeletePayload struct {
	// Identifier of package to delete
	ID uint
}

// CancelPayload is the payload type of the package service cancel method.
type CancelPayload struct {
	// Identifier of package to remove
	ID uint
}

// RetryPayload is the payload type of the package service retry method.
type RetryPayload struct {
	// Identifier of package to retry
	ID uint
}

// WorkflowPayload is the payload type of the package service workflow method.
type WorkflowPayload struct {
	// Identifier of package to look up
	ID uint
}

// EnduroPackageWorkflowStatus is the result type of the package service
// workflow method.
type EnduroPackageWorkflowStatus struct {
	Status  *string
	History EnduroPackageWorkflowHistoryCollection
}

// DownloadPayload is the payload type of the package service download method.
type DownloadPayload struct {
	// Identifier of package to look up
	ID uint
}

// BulkPayload is the payload type of the package service bulk method.
type BulkPayload struct {
	Operation string
	Status    string
	Size      uint
}

// BulkResult is the result type of the package service bulk method.
type BulkResult struct {
	WorkflowID string
	RunID      string
}

// BulkStatusResult is the result type of the package service bulk_status
// method.
type BulkStatusResult struct {
	Running    bool
	StartedAt  *string
	ClosedAt   *string
	Status     *string
	WorkflowID *string
	RunID      *string
}

// PreservationActionsPayload is the payload type of the package service
// preservation-actions method.
type PreservationActionsPayload struct {
	// Identifier of package to look up
	ID uint
}

// EnduroPackagePreservationActions is the result type of the package service
// preservation-actions method.
type EnduroPackagePreservationActions struct {
	Actions EnduroPackagePreservationActionsActionCollection
}

type EnduroStoredPackageCollection []*EnduroStoredPackage

type EnduroPackageWorkflowHistoryCollection []*EnduroPackageWorkflowHistory

// WorkflowHistoryEvent describes a history event in Temporal.
type EnduroPackageWorkflowHistory struct {
	// Identifier of package
	ID *uint
	// Type of the event
	Type *string
	// Contents of the event
	Details interface{}
}

type EnduroPackagePreservationActionsActionCollection []*EnduroPackagePreservationActionsAction

// PreservationAction describes a preservation action.
type EnduroPackagePreservationActionsAction struct {
	ID        uint
	ActionID  string
	Name      string
	Status    string
	StartedAt string
}

// Package not found.
type PackageNotfound struct {
	// Message of error
	Message string
	// Identifier of missing package
	ID uint
}

// Error returns an error description.
func (e *PackageNotfound) Error() string {
	return "Package not found."
}

// ErrorName returns "PackageNotfound".
func (e *PackageNotfound) ErrorName() string {
	return e.Message
}

// MakeNotRunning builds a goa.ServiceError from an error.
func MakeNotRunning(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_running",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotAvailable builds a goa.ServiceError from an error.
func MakeNotAvailable(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_available",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotValid builds a goa.ServiceError from an error.
func MakeNotValid(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_valid",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewEnduroMonitorUpdate initializes result type EnduroMonitorUpdate from
// viewed result type EnduroMonitorUpdate.
func NewEnduroMonitorUpdate(vres *package_views.EnduroMonitorUpdate) *EnduroMonitorUpdate {
	return newEnduroMonitorUpdate(vres.Projected)
}

// NewViewedEnduroMonitorUpdate initializes viewed result type
// EnduroMonitorUpdate from result type EnduroMonitorUpdate using the given
// view.
func NewViewedEnduroMonitorUpdate(res *EnduroMonitorUpdate, view string) *package_views.EnduroMonitorUpdate {
	p := newEnduroMonitorUpdateView(res)
	return &package_views.EnduroMonitorUpdate{Projected: p, View: "default"}
}

// NewEnduroStoredPackage initializes result type EnduroStoredPackage from
// viewed result type EnduroStoredPackage.
func NewEnduroStoredPackage(vres *package_views.EnduroStoredPackage) *EnduroStoredPackage {
	return newEnduroStoredPackage(vres.Projected)
}

// NewViewedEnduroStoredPackage initializes viewed result type
// EnduroStoredPackage from result type EnduroStoredPackage using the given
// view.
func NewViewedEnduroStoredPackage(res *EnduroStoredPackage, view string) *package_views.EnduroStoredPackage {
	p := newEnduroStoredPackageView(res)
	return &package_views.EnduroStoredPackage{Projected: p, View: "default"}
}

// NewEnduroPackageWorkflowStatus initializes result type
// EnduroPackageWorkflowStatus from viewed result type
// EnduroPackageWorkflowStatus.
func NewEnduroPackageWorkflowStatus(vres *package_views.EnduroPackageWorkflowStatus) *EnduroPackageWorkflowStatus {
	return newEnduroPackageWorkflowStatus(vres.Projected)
}

// NewViewedEnduroPackageWorkflowStatus initializes viewed result type
// EnduroPackageWorkflowStatus from result type EnduroPackageWorkflowStatus
// using the given view.
func NewViewedEnduroPackageWorkflowStatus(res *EnduroPackageWorkflowStatus, view string) *package_views.EnduroPackageWorkflowStatus {
	p := newEnduroPackageWorkflowStatusView(res)
	return &package_views.EnduroPackageWorkflowStatus{Projected: p, View: "default"}
}

// NewEnduroPackagePreservationActions initializes result type
// EnduroPackagePreservationActions from viewed result type
// EnduroPackagePreservationActions.
func NewEnduroPackagePreservationActions(vres *package_views.EnduroPackagePreservationActions) *EnduroPackagePreservationActions {
	return newEnduroPackagePreservationActions(vres.Projected)
}

// NewViewedEnduroPackagePreservationActions initializes viewed result type
// EnduroPackagePreservationActions from result type
// EnduroPackagePreservationActions using the given view.
func NewViewedEnduroPackagePreservationActions(res *EnduroPackagePreservationActions, view string) *package_views.EnduroPackagePreservationActions {
	p := newEnduroPackagePreservationActionsView(res)
	return &package_views.EnduroPackagePreservationActions{Projected: p, View: "default"}
}

// newEnduroMonitorUpdate converts projected type EnduroMonitorUpdate to
// service type EnduroMonitorUpdate.
func newEnduroMonitorUpdate(vres *package_views.EnduroMonitorUpdateView) *EnduroMonitorUpdate {
	res := &EnduroMonitorUpdate{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Item != nil {
		res.Item = newEnduroStoredPackage(vres.Item)
	}
	return res
}

// newEnduroMonitorUpdateView projects result type EnduroMonitorUpdate to
// projected type EnduroMonitorUpdateView using the "default" view.
func newEnduroMonitorUpdateView(res *EnduroMonitorUpdate) *package_views.EnduroMonitorUpdateView {
	vres := &package_views.EnduroMonitorUpdateView{
		ID:   &res.ID,
		Type: &res.Type,
	}
	if res.Item != nil {
		vres.Item = newEnduroStoredPackageView(res.Item)
	}
	return vres
}

// newEnduroStoredPackage converts projected type EnduroStoredPackage to
// service type EnduroStoredPackage.
func newEnduroStoredPackage(vres *package_views.EnduroStoredPackageView) *EnduroStoredPackage {
	res := &EnduroStoredPackage{
		Name:        vres.Name,
		WorkflowID:  vres.WorkflowID,
		RunID:       vres.RunID,
		AipID:       vres.AipID,
		StartedAt:   vres.StartedAt,
		CompletedAt: vres.CompletedAt,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.CreatedAt != nil {
		res.CreatedAt = *vres.CreatedAt
	}
	if vres.Status == nil {
		res.Status = "new"
	}
	return res
}

// newEnduroStoredPackageView projects result type EnduroStoredPackage to
// projected type EnduroStoredPackageView using the "default" view.
func newEnduroStoredPackageView(res *EnduroStoredPackage) *package_views.EnduroStoredPackageView {
	vres := &package_views.EnduroStoredPackageView{
		ID:          &res.ID,
		Name:        res.Name,
		Status:      &res.Status,
		WorkflowID:  res.WorkflowID,
		RunID:       res.RunID,
		AipID:       res.AipID,
		CreatedAt:   &res.CreatedAt,
		StartedAt:   res.StartedAt,
		CompletedAt: res.CompletedAt,
	}
	return vres
}

// newEnduroPackageWorkflowStatus converts projected type
// EnduroPackageWorkflowStatus to service type EnduroPackageWorkflowStatus.
func newEnduroPackageWorkflowStatus(vres *package_views.EnduroPackageWorkflowStatusView) *EnduroPackageWorkflowStatus {
	res := &EnduroPackageWorkflowStatus{
		Status: vres.Status,
	}
	if vres.History != nil {
		res.History = newEnduroPackageWorkflowHistoryCollection(vres.History)
	}
	return res
}

// newEnduroPackageWorkflowStatusView projects result type
// EnduroPackageWorkflowStatus to projected type
// EnduroPackageWorkflowStatusView using the "default" view.
func newEnduroPackageWorkflowStatusView(res *EnduroPackageWorkflowStatus) *package_views.EnduroPackageWorkflowStatusView {
	vres := &package_views.EnduroPackageWorkflowStatusView{
		Status: res.Status,
	}
	if res.History != nil {
		vres.History = newEnduroPackageWorkflowHistoryCollectionView(res.History)
	}
	return vres
}

// newEnduroPackageWorkflowHistoryCollection converts projected type
// EnduroPackageWorkflowHistoryCollection to service type
// EnduroPackageWorkflowHistoryCollection.
func newEnduroPackageWorkflowHistoryCollection(vres package_views.EnduroPackageWorkflowHistoryCollectionView) EnduroPackageWorkflowHistoryCollection {
	res := make(EnduroPackageWorkflowHistoryCollection, len(vres))
	for i, n := range vres {
		res[i] = newEnduroPackageWorkflowHistory(n)
	}
	return res
}

// newEnduroPackageWorkflowHistoryCollectionView projects result type
// EnduroPackageWorkflowHistoryCollection to projected type
// EnduroPackageWorkflowHistoryCollectionView using the "default" view.
func newEnduroPackageWorkflowHistoryCollectionView(res EnduroPackageWorkflowHistoryCollection) package_views.EnduroPackageWorkflowHistoryCollectionView {
	vres := make(package_views.EnduroPackageWorkflowHistoryCollectionView, len(res))
	for i, n := range res {
		vres[i] = newEnduroPackageWorkflowHistoryView(n)
	}
	return vres
}

// newEnduroPackageWorkflowHistory converts projected type
// EnduroPackageWorkflowHistory to service type EnduroPackageWorkflowHistory.
func newEnduroPackageWorkflowHistory(vres *package_views.EnduroPackageWorkflowHistoryView) *EnduroPackageWorkflowHistory {
	res := &EnduroPackageWorkflowHistory{
		ID:      vres.ID,
		Type:    vres.Type,
		Details: vres.Details,
	}
	return res
}

// newEnduroPackageWorkflowHistoryView projects result type
// EnduroPackageWorkflowHistory to projected type
// EnduroPackageWorkflowHistoryView using the "default" view.
func newEnduroPackageWorkflowHistoryView(res *EnduroPackageWorkflowHistory) *package_views.EnduroPackageWorkflowHistoryView {
	vres := &package_views.EnduroPackageWorkflowHistoryView{
		ID:      res.ID,
		Type:    res.Type,
		Details: res.Details,
	}
	return vres
}

// newEnduroPackagePreservationActions converts projected type
// EnduroPackagePreservationActions to service type
// EnduroPackagePreservationActions.
func newEnduroPackagePreservationActions(vres *package_views.EnduroPackagePreservationActionsView) *EnduroPackagePreservationActions {
	res := &EnduroPackagePreservationActions{}
	if vres.Actions != nil {
		res.Actions = newEnduroPackagePreservationActionsActionCollection(vres.Actions)
	}
	return res
}

// newEnduroPackagePreservationActionsView projects result type
// EnduroPackagePreservationActions to projected type
// EnduroPackagePreservationActionsView using the "default" view.
func newEnduroPackagePreservationActionsView(res *EnduroPackagePreservationActions) *package_views.EnduroPackagePreservationActionsView {
	vres := &package_views.EnduroPackagePreservationActionsView{}
	if res.Actions != nil {
		vres.Actions = newEnduroPackagePreservationActionsActionCollectionView(res.Actions)
	}
	return vres
}

// newEnduroPackagePreservationActionsActionCollection converts projected type
// EnduroPackagePreservationActionsActionCollection to service type
// EnduroPackagePreservationActionsActionCollection.
func newEnduroPackagePreservationActionsActionCollection(vres package_views.EnduroPackagePreservationActionsActionCollectionView) EnduroPackagePreservationActionsActionCollection {
	res := make(EnduroPackagePreservationActionsActionCollection, len(vres))
	for i, n := range vres {
		res[i] = newEnduroPackagePreservationActionsAction(n)
	}
	return res
}

// newEnduroPackagePreservationActionsActionCollectionView projects result type
// EnduroPackagePreservationActionsActionCollection to projected type
// EnduroPackagePreservationActionsActionCollectionView using the "default"
// view.
func newEnduroPackagePreservationActionsActionCollectionView(res EnduroPackagePreservationActionsActionCollection) package_views.EnduroPackagePreservationActionsActionCollectionView {
	vres := make(package_views.EnduroPackagePreservationActionsActionCollectionView, len(res))
	for i, n := range res {
		vres[i] = newEnduroPackagePreservationActionsActionView(n)
	}
	return vres
}

// newEnduroPackagePreservationActionsAction converts projected type
// EnduroPackagePreservationActionsAction to service type
// EnduroPackagePreservationActionsAction.
func newEnduroPackagePreservationActionsAction(vres *package_views.EnduroPackagePreservationActionsActionView) *EnduroPackagePreservationActionsAction {
	res := &EnduroPackagePreservationActionsAction{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.ActionID != nil {
		res.ActionID = *vres.ActionID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Status != nil {
		res.Status = *vres.Status
	}
	if vres.StartedAt != nil {
		res.StartedAt = *vres.StartedAt
	}
	return res
}

// newEnduroPackagePreservationActionsActionView projects result type
// EnduroPackagePreservationActionsAction to projected type
// EnduroPackagePreservationActionsActionView using the "default" view.
func newEnduroPackagePreservationActionsActionView(res *EnduroPackagePreservationActionsAction) *package_views.EnduroPackagePreservationActionsActionView {
	vres := &package_views.EnduroPackagePreservationActionsActionView{
		ID:        &res.ID,
		ActionID:  &res.ActionID,
		Name:      &res.Name,
		Status:    &res.Status,
		StartedAt: &res.StartedAt,
	}
	return vres
}
