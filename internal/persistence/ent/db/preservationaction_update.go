// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/pkg"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/predicate"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/preservationaction"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/preservationtask"
)

// PreservationActionUpdate is the builder for updating PreservationAction entities.
type PreservationActionUpdate struct {
	config
	hooks    []Hook
	mutation *PreservationActionMutation
}

// Where appends a list predicates to the PreservationActionUpdate builder.
func (pau *PreservationActionUpdate) Where(ps ...predicate.PreservationAction) *PreservationActionUpdate {
	pau.mutation.Where(ps...)
	return pau
}

// SetWorkflowID sets the "workflow_id" field.
func (pau *PreservationActionUpdate) SetWorkflowID(s string) *PreservationActionUpdate {
	pau.mutation.SetWorkflowID(s)
	return pau
}

// SetNillableWorkflowID sets the "workflow_id" field if the given value is not nil.
func (pau *PreservationActionUpdate) SetNillableWorkflowID(s *string) *PreservationActionUpdate {
	if s != nil {
		pau.SetWorkflowID(*s)
	}
	return pau
}

// SetType sets the "type" field.
func (pau *PreservationActionUpdate) SetType(i int8) *PreservationActionUpdate {
	pau.mutation.ResetType()
	pau.mutation.SetType(i)
	return pau
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pau *PreservationActionUpdate) SetNillableType(i *int8) *PreservationActionUpdate {
	if i != nil {
		pau.SetType(*i)
	}
	return pau
}

// AddType adds i to the "type" field.
func (pau *PreservationActionUpdate) AddType(i int8) *PreservationActionUpdate {
	pau.mutation.AddType(i)
	return pau
}

// SetStatus sets the "status" field.
func (pau *PreservationActionUpdate) SetStatus(i int8) *PreservationActionUpdate {
	pau.mutation.ResetStatus()
	pau.mutation.SetStatus(i)
	return pau
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pau *PreservationActionUpdate) SetNillableStatus(i *int8) *PreservationActionUpdate {
	if i != nil {
		pau.SetStatus(*i)
	}
	return pau
}

// AddStatus adds i to the "status" field.
func (pau *PreservationActionUpdate) AddStatus(i int8) *PreservationActionUpdate {
	pau.mutation.AddStatus(i)
	return pau
}

// SetStartedAt sets the "started_at" field.
func (pau *PreservationActionUpdate) SetStartedAt(t time.Time) *PreservationActionUpdate {
	pau.mutation.SetStartedAt(t)
	return pau
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (pau *PreservationActionUpdate) SetNillableStartedAt(t *time.Time) *PreservationActionUpdate {
	if t != nil {
		pau.SetStartedAt(*t)
	}
	return pau
}

// ClearStartedAt clears the value of the "started_at" field.
func (pau *PreservationActionUpdate) ClearStartedAt() *PreservationActionUpdate {
	pau.mutation.ClearStartedAt()
	return pau
}

// SetCompletedAt sets the "completed_at" field.
func (pau *PreservationActionUpdate) SetCompletedAt(t time.Time) *PreservationActionUpdate {
	pau.mutation.SetCompletedAt(t)
	return pau
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (pau *PreservationActionUpdate) SetNillableCompletedAt(t *time.Time) *PreservationActionUpdate {
	if t != nil {
		pau.SetCompletedAt(*t)
	}
	return pau
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (pau *PreservationActionUpdate) ClearCompletedAt() *PreservationActionUpdate {
	pau.mutation.ClearCompletedAt()
	return pau
}

// SetPackageID sets the "package_id" field.
func (pau *PreservationActionUpdate) SetPackageID(i int) *PreservationActionUpdate {
	pau.mutation.SetPackageID(i)
	return pau
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (pau *PreservationActionUpdate) SetNillablePackageID(i *int) *PreservationActionUpdate {
	if i != nil {
		pau.SetPackageID(*i)
	}
	return pau
}

// SetPackage sets the "package" edge to the Pkg entity.
func (pau *PreservationActionUpdate) SetPackage(p *Pkg) *PreservationActionUpdate {
	return pau.SetPackageID(p.ID)
}

// AddTaskIDs adds the "tasks" edge to the PreservationTask entity by IDs.
func (pau *PreservationActionUpdate) AddTaskIDs(ids ...int) *PreservationActionUpdate {
	pau.mutation.AddTaskIDs(ids...)
	return pau
}

// AddTasks adds the "tasks" edges to the PreservationTask entity.
func (pau *PreservationActionUpdate) AddTasks(p ...*PreservationTask) *PreservationActionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pau.AddTaskIDs(ids...)
}

// Mutation returns the PreservationActionMutation object of the builder.
func (pau *PreservationActionUpdate) Mutation() *PreservationActionMutation {
	return pau.mutation
}

// ClearPackage clears the "package" edge to the Pkg entity.
func (pau *PreservationActionUpdate) ClearPackage() *PreservationActionUpdate {
	pau.mutation.ClearPackage()
	return pau
}

// ClearTasks clears all "tasks" edges to the PreservationTask entity.
func (pau *PreservationActionUpdate) ClearTasks() *PreservationActionUpdate {
	pau.mutation.ClearTasks()
	return pau
}

// RemoveTaskIDs removes the "tasks" edge to PreservationTask entities by IDs.
func (pau *PreservationActionUpdate) RemoveTaskIDs(ids ...int) *PreservationActionUpdate {
	pau.mutation.RemoveTaskIDs(ids...)
	return pau
}

// RemoveTasks removes "tasks" edges to PreservationTask entities.
func (pau *PreservationActionUpdate) RemoveTasks(p ...*PreservationTask) *PreservationActionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pau.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *PreservationActionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pau.sqlSave, pau.mutation, pau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pau *PreservationActionUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *PreservationActionUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *PreservationActionUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pau *PreservationActionUpdate) check() error {
	if v, ok := pau.mutation.PackageID(); ok {
		if err := preservationaction.PackageIDValidator(v); err != nil {
			return &ValidationError{Name: "package_id", err: fmt.Errorf(`db: validator failed for field "PreservationAction.package_id": %w`, err)}
		}
	}
	if _, ok := pau.mutation.PackageID(); pau.mutation.PackageCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "PreservationAction.package"`)
	}
	return nil
}

func (pau *PreservationActionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(preservationaction.Table, preservationaction.Columns, sqlgraph.NewFieldSpec(preservationaction.FieldID, field.TypeInt))
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.WorkflowID(); ok {
		_spec.SetField(preservationaction.FieldWorkflowID, field.TypeString, value)
	}
	if value, ok := pau.mutation.GetType(); ok {
		_spec.SetField(preservationaction.FieldType, field.TypeInt8, value)
	}
	if value, ok := pau.mutation.AddedType(); ok {
		_spec.AddField(preservationaction.FieldType, field.TypeInt8, value)
	}
	if value, ok := pau.mutation.Status(); ok {
		_spec.SetField(preservationaction.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pau.mutation.AddedStatus(); ok {
		_spec.AddField(preservationaction.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pau.mutation.StartedAt(); ok {
		_spec.SetField(preservationaction.FieldStartedAt, field.TypeTime, value)
	}
	if pau.mutation.StartedAtCleared() {
		_spec.ClearField(preservationaction.FieldStartedAt, field.TypeTime)
	}
	if value, ok := pau.mutation.CompletedAt(); ok {
		_spec.SetField(preservationaction.FieldCompletedAt, field.TypeTime, value)
	}
	if pau.mutation.CompletedAtCleared() {
		_spec.ClearField(preservationaction.FieldCompletedAt, field.TypeTime)
	}
	if pau.mutation.PackageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   preservationaction.PackageTable,
			Columns: []string{preservationaction.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pkg.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.PackageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   preservationaction.PackageTable,
			Columns: []string{preservationaction.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pkg.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pau.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   preservationaction.TasksTable,
			Columns: []string{preservationaction.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preservationtask.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.RemovedTasksIDs(); len(nodes) > 0 && !pau.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   preservationaction.TasksTable,
			Columns: []string{preservationaction.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preservationtask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pau.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   preservationaction.TasksTable,
			Columns: []string{preservationaction.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preservationtask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{preservationaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pau.mutation.done = true
	return n, nil
}

// PreservationActionUpdateOne is the builder for updating a single PreservationAction entity.
type PreservationActionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PreservationActionMutation
}

// SetWorkflowID sets the "workflow_id" field.
func (pauo *PreservationActionUpdateOne) SetWorkflowID(s string) *PreservationActionUpdateOne {
	pauo.mutation.SetWorkflowID(s)
	return pauo
}

// SetNillableWorkflowID sets the "workflow_id" field if the given value is not nil.
func (pauo *PreservationActionUpdateOne) SetNillableWorkflowID(s *string) *PreservationActionUpdateOne {
	if s != nil {
		pauo.SetWorkflowID(*s)
	}
	return pauo
}

// SetType sets the "type" field.
func (pauo *PreservationActionUpdateOne) SetType(i int8) *PreservationActionUpdateOne {
	pauo.mutation.ResetType()
	pauo.mutation.SetType(i)
	return pauo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pauo *PreservationActionUpdateOne) SetNillableType(i *int8) *PreservationActionUpdateOne {
	if i != nil {
		pauo.SetType(*i)
	}
	return pauo
}

// AddType adds i to the "type" field.
func (pauo *PreservationActionUpdateOne) AddType(i int8) *PreservationActionUpdateOne {
	pauo.mutation.AddType(i)
	return pauo
}

// SetStatus sets the "status" field.
func (pauo *PreservationActionUpdateOne) SetStatus(i int8) *PreservationActionUpdateOne {
	pauo.mutation.ResetStatus()
	pauo.mutation.SetStatus(i)
	return pauo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pauo *PreservationActionUpdateOne) SetNillableStatus(i *int8) *PreservationActionUpdateOne {
	if i != nil {
		pauo.SetStatus(*i)
	}
	return pauo
}

// AddStatus adds i to the "status" field.
func (pauo *PreservationActionUpdateOne) AddStatus(i int8) *PreservationActionUpdateOne {
	pauo.mutation.AddStatus(i)
	return pauo
}

// SetStartedAt sets the "started_at" field.
func (pauo *PreservationActionUpdateOne) SetStartedAt(t time.Time) *PreservationActionUpdateOne {
	pauo.mutation.SetStartedAt(t)
	return pauo
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (pauo *PreservationActionUpdateOne) SetNillableStartedAt(t *time.Time) *PreservationActionUpdateOne {
	if t != nil {
		pauo.SetStartedAt(*t)
	}
	return pauo
}

// ClearStartedAt clears the value of the "started_at" field.
func (pauo *PreservationActionUpdateOne) ClearStartedAt() *PreservationActionUpdateOne {
	pauo.mutation.ClearStartedAt()
	return pauo
}

// SetCompletedAt sets the "completed_at" field.
func (pauo *PreservationActionUpdateOne) SetCompletedAt(t time.Time) *PreservationActionUpdateOne {
	pauo.mutation.SetCompletedAt(t)
	return pauo
}

// SetNillableCompletedAt sets the "completed_at" field if the given value is not nil.
func (pauo *PreservationActionUpdateOne) SetNillableCompletedAt(t *time.Time) *PreservationActionUpdateOne {
	if t != nil {
		pauo.SetCompletedAt(*t)
	}
	return pauo
}

// ClearCompletedAt clears the value of the "completed_at" field.
func (pauo *PreservationActionUpdateOne) ClearCompletedAt() *PreservationActionUpdateOne {
	pauo.mutation.ClearCompletedAt()
	return pauo
}

// SetPackageID sets the "package_id" field.
func (pauo *PreservationActionUpdateOne) SetPackageID(i int) *PreservationActionUpdateOne {
	pauo.mutation.SetPackageID(i)
	return pauo
}

// SetNillablePackageID sets the "package_id" field if the given value is not nil.
func (pauo *PreservationActionUpdateOne) SetNillablePackageID(i *int) *PreservationActionUpdateOne {
	if i != nil {
		pauo.SetPackageID(*i)
	}
	return pauo
}

// SetPackage sets the "package" edge to the Pkg entity.
func (pauo *PreservationActionUpdateOne) SetPackage(p *Pkg) *PreservationActionUpdateOne {
	return pauo.SetPackageID(p.ID)
}

// AddTaskIDs adds the "tasks" edge to the PreservationTask entity by IDs.
func (pauo *PreservationActionUpdateOne) AddTaskIDs(ids ...int) *PreservationActionUpdateOne {
	pauo.mutation.AddTaskIDs(ids...)
	return pauo
}

// AddTasks adds the "tasks" edges to the PreservationTask entity.
func (pauo *PreservationActionUpdateOne) AddTasks(p ...*PreservationTask) *PreservationActionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pauo.AddTaskIDs(ids...)
}

// Mutation returns the PreservationActionMutation object of the builder.
func (pauo *PreservationActionUpdateOne) Mutation() *PreservationActionMutation {
	return pauo.mutation
}

// ClearPackage clears the "package" edge to the Pkg entity.
func (pauo *PreservationActionUpdateOne) ClearPackage() *PreservationActionUpdateOne {
	pauo.mutation.ClearPackage()
	return pauo
}

// ClearTasks clears all "tasks" edges to the PreservationTask entity.
func (pauo *PreservationActionUpdateOne) ClearTasks() *PreservationActionUpdateOne {
	pauo.mutation.ClearTasks()
	return pauo
}

// RemoveTaskIDs removes the "tasks" edge to PreservationTask entities by IDs.
func (pauo *PreservationActionUpdateOne) RemoveTaskIDs(ids ...int) *PreservationActionUpdateOne {
	pauo.mutation.RemoveTaskIDs(ids...)
	return pauo
}

// RemoveTasks removes "tasks" edges to PreservationTask entities.
func (pauo *PreservationActionUpdateOne) RemoveTasks(p ...*PreservationTask) *PreservationActionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pauo.RemoveTaskIDs(ids...)
}

// Where appends a list predicates to the PreservationActionUpdate builder.
func (pauo *PreservationActionUpdateOne) Where(ps ...predicate.PreservationAction) *PreservationActionUpdateOne {
	pauo.mutation.Where(ps...)
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *PreservationActionUpdateOne) Select(field string, fields ...string) *PreservationActionUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated PreservationAction entity.
func (pauo *PreservationActionUpdateOne) Save(ctx context.Context) (*PreservationAction, error) {
	return withHooks(ctx, pauo.sqlSave, pauo.mutation, pauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *PreservationActionUpdateOne) SaveX(ctx context.Context) *PreservationAction {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *PreservationActionUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *PreservationActionUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pauo *PreservationActionUpdateOne) check() error {
	if v, ok := pauo.mutation.PackageID(); ok {
		if err := preservationaction.PackageIDValidator(v); err != nil {
			return &ValidationError{Name: "package_id", err: fmt.Errorf(`db: validator failed for field "PreservationAction.package_id": %w`, err)}
		}
	}
	if _, ok := pauo.mutation.PackageID(); pauo.mutation.PackageCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "PreservationAction.package"`)
	}
	return nil
}

func (pauo *PreservationActionUpdateOne) sqlSave(ctx context.Context) (_node *PreservationAction, err error) {
	if err := pauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(preservationaction.Table, preservationaction.Columns, sqlgraph.NewFieldSpec(preservationaction.FieldID, field.TypeInt))
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "PreservationAction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, preservationaction.FieldID)
		for _, f := range fields {
			if !preservationaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != preservationaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.WorkflowID(); ok {
		_spec.SetField(preservationaction.FieldWorkflowID, field.TypeString, value)
	}
	if value, ok := pauo.mutation.GetType(); ok {
		_spec.SetField(preservationaction.FieldType, field.TypeInt8, value)
	}
	if value, ok := pauo.mutation.AddedType(); ok {
		_spec.AddField(preservationaction.FieldType, field.TypeInt8, value)
	}
	if value, ok := pauo.mutation.Status(); ok {
		_spec.SetField(preservationaction.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pauo.mutation.AddedStatus(); ok {
		_spec.AddField(preservationaction.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := pauo.mutation.StartedAt(); ok {
		_spec.SetField(preservationaction.FieldStartedAt, field.TypeTime, value)
	}
	if pauo.mutation.StartedAtCleared() {
		_spec.ClearField(preservationaction.FieldStartedAt, field.TypeTime)
	}
	if value, ok := pauo.mutation.CompletedAt(); ok {
		_spec.SetField(preservationaction.FieldCompletedAt, field.TypeTime, value)
	}
	if pauo.mutation.CompletedAtCleared() {
		_spec.ClearField(preservationaction.FieldCompletedAt, field.TypeTime)
	}
	if pauo.mutation.PackageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   preservationaction.PackageTable,
			Columns: []string{preservationaction.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pkg.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.PackageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   preservationaction.PackageTable,
			Columns: []string{preservationaction.PackageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pkg.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pauo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   preservationaction.TasksTable,
			Columns: []string{preservationaction.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preservationtask.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !pauo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   preservationaction.TasksTable,
			Columns: []string{preservationaction.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preservationtask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pauo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   preservationaction.TasksTable,
			Columns: []string{preservationaction.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preservationtask.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PreservationAction{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{preservationaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pauo.mutation.done = true
	return _node, nil
}
