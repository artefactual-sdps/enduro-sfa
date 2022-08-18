// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/location"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/pkg"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/predicate"
	"github.com/artefactual-sdps/enduro/internal/storage/types"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeLocation = "Location"
	TypePkg      = "Pkg"
)

// LocationMutation represents an operation that mutates the Location nodes in the graph.
type LocationMutation struct {
	config
	op              Op
	typ             string
	id              *int
	name            *string
	description     *string
	source          *types.LocationSource
	purpose         *types.LocationPurpose
	uuid            *uuid.UUID
	_config         *types.LocationConfig
	clearedFields   map[string]struct{}
	packages        map[int]struct{}
	removedpackages map[int]struct{}
	clearedpackages bool
	done            bool
	oldValue        func(context.Context) (*Location, error)
	predicates      []predicate.Location
}

var _ ent.Mutation = (*LocationMutation)(nil)

// locationOption allows management of the mutation configuration using functional options.
type locationOption func(*LocationMutation)

// newLocationMutation creates new mutation for the Location entity.
func newLocationMutation(c config, op Op, opts ...locationOption) *LocationMutation {
	m := &LocationMutation{
		config:        c,
		op:            op,
		typ:           TypeLocation,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLocationID sets the ID field of the mutation.
func withLocationID(id int) locationOption {
	return func(m *LocationMutation) {
		var (
			err   error
			once  sync.Once
			value *Location
		)
		m.oldValue = func(ctx context.Context) (*Location, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Location.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLocation sets the old Location of the mutation.
func withLocation(node *Location) locationOption {
	return func(m *LocationMutation) {
		m.oldValue = func(context.Context) (*Location, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LocationMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LocationMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("db: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LocationMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *LocationMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Location.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *LocationMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *LocationMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *LocationMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *LocationMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *LocationMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *LocationMutation) ResetDescription() {
	m.description = nil
}

// SetSource sets the "source" field.
func (m *LocationMutation) SetSource(ts types.LocationSource) {
	m.source = &ts
}

// Source returns the value of the "source" field in the mutation.
func (m *LocationMutation) Source() (r types.LocationSource, exists bool) {
	v := m.source
	if v == nil {
		return
	}
	return *v, true
}

// OldSource returns the old "source" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldSource(ctx context.Context) (v types.LocationSource, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSource is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSource requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSource: %w", err)
	}
	return oldValue.Source, nil
}

// ResetSource resets all changes to the "source" field.
func (m *LocationMutation) ResetSource() {
	m.source = nil
}

// SetPurpose sets the "purpose" field.
func (m *LocationMutation) SetPurpose(tp types.LocationPurpose) {
	m.purpose = &tp
}

// Purpose returns the value of the "purpose" field in the mutation.
func (m *LocationMutation) Purpose() (r types.LocationPurpose, exists bool) {
	v := m.purpose
	if v == nil {
		return
	}
	return *v, true
}

// OldPurpose returns the old "purpose" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldPurpose(ctx context.Context) (v types.LocationPurpose, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPurpose is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPurpose requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPurpose: %w", err)
	}
	return oldValue.Purpose, nil
}

// ResetPurpose resets all changes to the "purpose" field.
func (m *LocationMutation) ResetPurpose() {
	m.purpose = nil
}

// SetUUID sets the "uuid" field.
func (m *LocationMutation) SetUUID(u uuid.UUID) {
	m.uuid = &u
}

// UUID returns the value of the "uuid" field in the mutation.
func (m *LocationMutation) UUID() (r uuid.UUID, exists bool) {
	v := m.uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldUUID returns the old "uuid" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUUID: %w", err)
	}
	return oldValue.UUID, nil
}

// ResetUUID resets all changes to the "uuid" field.
func (m *LocationMutation) ResetUUID() {
	m.uuid = nil
}

// SetConfig sets the "config" field.
func (m *LocationMutation) SetConfig(tc types.LocationConfig) {
	m._config = &tc
}

// Config returns the value of the "config" field in the mutation.
func (m *LocationMutation) Config() (r types.LocationConfig, exists bool) {
	v := m._config
	if v == nil {
		return
	}
	return *v, true
}

// OldConfig returns the old "config" field's value of the Location entity.
// If the Location object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LocationMutation) OldConfig(ctx context.Context) (v types.LocationConfig, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldConfig is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldConfig requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfig: %w", err)
	}
	return oldValue.Config, nil
}

// ResetConfig resets all changes to the "config" field.
func (m *LocationMutation) ResetConfig() {
	m._config = nil
}

// AddPackageIDs adds the "packages" edge to the Pkg entity by ids.
func (m *LocationMutation) AddPackageIDs(ids ...int) {
	if m.packages == nil {
		m.packages = make(map[int]struct{})
	}
	for i := range ids {
		m.packages[ids[i]] = struct{}{}
	}
}

// ClearPackages clears the "packages" edge to the Pkg entity.
func (m *LocationMutation) ClearPackages() {
	m.clearedpackages = true
}

// PackagesCleared reports if the "packages" edge to the Pkg entity was cleared.
func (m *LocationMutation) PackagesCleared() bool {
	return m.clearedpackages
}

// RemovePackageIDs removes the "packages" edge to the Pkg entity by IDs.
func (m *LocationMutation) RemovePackageIDs(ids ...int) {
	if m.removedpackages == nil {
		m.removedpackages = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.packages, ids[i])
		m.removedpackages[ids[i]] = struct{}{}
	}
}

// RemovedPackages returns the removed IDs of the "packages" edge to the Pkg entity.
func (m *LocationMutation) RemovedPackagesIDs() (ids []int) {
	for id := range m.removedpackages {
		ids = append(ids, id)
	}
	return
}

// PackagesIDs returns the "packages" edge IDs in the mutation.
func (m *LocationMutation) PackagesIDs() (ids []int) {
	for id := range m.packages {
		ids = append(ids, id)
	}
	return
}

// ResetPackages resets all changes to the "packages" edge.
func (m *LocationMutation) ResetPackages() {
	m.packages = nil
	m.clearedpackages = false
	m.removedpackages = nil
}

// Where appends a list predicates to the LocationMutation builder.
func (m *LocationMutation) Where(ps ...predicate.Location) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *LocationMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Location).
func (m *LocationMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LocationMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.name != nil {
		fields = append(fields, location.FieldName)
	}
	if m.description != nil {
		fields = append(fields, location.FieldDescription)
	}
	if m.source != nil {
		fields = append(fields, location.FieldSource)
	}
	if m.purpose != nil {
		fields = append(fields, location.FieldPurpose)
	}
	if m.uuid != nil {
		fields = append(fields, location.FieldUUID)
	}
	if m._config != nil {
		fields = append(fields, location.FieldConfig)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LocationMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case location.FieldName:
		return m.Name()
	case location.FieldDescription:
		return m.Description()
	case location.FieldSource:
		return m.Source()
	case location.FieldPurpose:
		return m.Purpose()
	case location.FieldUUID:
		return m.UUID()
	case location.FieldConfig:
		return m.Config()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LocationMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case location.FieldName:
		return m.OldName(ctx)
	case location.FieldDescription:
		return m.OldDescription(ctx)
	case location.FieldSource:
		return m.OldSource(ctx)
	case location.FieldPurpose:
		return m.OldPurpose(ctx)
	case location.FieldUUID:
		return m.OldUUID(ctx)
	case location.FieldConfig:
		return m.OldConfig(ctx)
	}
	return nil, fmt.Errorf("unknown Location field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LocationMutation) SetField(name string, value ent.Value) error {
	switch name {
	case location.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case location.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case location.FieldSource:
		v, ok := value.(types.LocationSource)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSource(v)
		return nil
	case location.FieldPurpose:
		v, ok := value.(types.LocationPurpose)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPurpose(v)
		return nil
	case location.FieldUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUUID(v)
		return nil
	case location.FieldConfig:
		v, ok := value.(types.LocationConfig)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfig(v)
		return nil
	}
	return fmt.Errorf("unknown Location field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LocationMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LocationMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LocationMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Location numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LocationMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LocationMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LocationMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Location nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LocationMutation) ResetField(name string) error {
	switch name {
	case location.FieldName:
		m.ResetName()
		return nil
	case location.FieldDescription:
		m.ResetDescription()
		return nil
	case location.FieldSource:
		m.ResetSource()
		return nil
	case location.FieldPurpose:
		m.ResetPurpose()
		return nil
	case location.FieldUUID:
		m.ResetUUID()
		return nil
	case location.FieldConfig:
		m.ResetConfig()
		return nil
	}
	return fmt.Errorf("unknown Location field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LocationMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.packages != nil {
		edges = append(edges, location.EdgePackages)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LocationMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case location.EdgePackages:
		ids := make([]ent.Value, 0, len(m.packages))
		for id := range m.packages {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LocationMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedpackages != nil {
		edges = append(edges, location.EdgePackages)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LocationMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case location.EdgePackages:
		ids := make([]ent.Value, 0, len(m.removedpackages))
		for id := range m.removedpackages {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LocationMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedpackages {
		edges = append(edges, location.EdgePackages)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LocationMutation) EdgeCleared(name string) bool {
	switch name {
	case location.EdgePackages:
		return m.clearedpackages
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LocationMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Location unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LocationMutation) ResetEdge(name string) error {
	switch name {
	case location.EdgePackages:
		m.ResetPackages()
		return nil
	}
	return fmt.Errorf("unknown Location edge %s", name)
}

// PkgMutation represents an operation that mutates the Pkg nodes in the graph.
type PkgMutation struct {
	config
	op              Op
	typ             string
	id              *int
	name            *string
	aip_id          *uuid.UUID
	status          *types.PackageStatus
	object_key      *uuid.UUID
	clearedFields   map[string]struct{}
	location        *int
	clearedlocation bool
	done            bool
	oldValue        func(context.Context) (*Pkg, error)
	predicates      []predicate.Pkg
}

var _ ent.Mutation = (*PkgMutation)(nil)

// pkgOption allows management of the mutation configuration using functional options.
type pkgOption func(*PkgMutation)

// newPkgMutation creates new mutation for the Pkg entity.
func newPkgMutation(c config, op Op, opts ...pkgOption) *PkgMutation {
	m := &PkgMutation{
		config:        c,
		op:            op,
		typ:           TypePkg,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPkgID sets the ID field of the mutation.
func withPkgID(id int) pkgOption {
	return func(m *PkgMutation) {
		var (
			err   error
			once  sync.Once
			value *Pkg
		)
		m.oldValue = func(ctx context.Context) (*Pkg, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Pkg.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPkg sets the old Pkg of the mutation.
func withPkg(node *Pkg) pkgOption {
	return func(m *PkgMutation) {
		m.oldValue = func(context.Context) (*Pkg, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PkgMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PkgMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("db: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PkgMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PkgMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Pkg.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *PkgMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *PkgMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Pkg entity.
// If the Pkg object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PkgMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *PkgMutation) ResetName() {
	m.name = nil
}

// SetAipID sets the "aip_id" field.
func (m *PkgMutation) SetAipID(u uuid.UUID) {
	m.aip_id = &u
}

// AipID returns the value of the "aip_id" field in the mutation.
func (m *PkgMutation) AipID() (r uuid.UUID, exists bool) {
	v := m.aip_id
	if v == nil {
		return
	}
	return *v, true
}

// OldAipID returns the old "aip_id" field's value of the Pkg entity.
// If the Pkg object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PkgMutation) OldAipID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAipID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAipID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAipID: %w", err)
	}
	return oldValue.AipID, nil
}

// ResetAipID resets all changes to the "aip_id" field.
func (m *PkgMutation) ResetAipID() {
	m.aip_id = nil
}

// SetLocationID sets the "location_id" field.
func (m *PkgMutation) SetLocationID(i int) {
	m.location = &i
}

// LocationID returns the value of the "location_id" field in the mutation.
func (m *PkgMutation) LocationID() (r int, exists bool) {
	v := m.location
	if v == nil {
		return
	}
	return *v, true
}

// OldLocationID returns the old "location_id" field's value of the Pkg entity.
// If the Pkg object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PkgMutation) OldLocationID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLocationID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLocationID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLocationID: %w", err)
	}
	return oldValue.LocationID, nil
}

// ClearLocationID clears the value of the "location_id" field.
func (m *PkgMutation) ClearLocationID() {
	m.location = nil
	m.clearedFields[pkg.FieldLocationID] = struct{}{}
}

// LocationIDCleared returns if the "location_id" field was cleared in this mutation.
func (m *PkgMutation) LocationIDCleared() bool {
	_, ok := m.clearedFields[pkg.FieldLocationID]
	return ok
}

// ResetLocationID resets all changes to the "location_id" field.
func (m *PkgMutation) ResetLocationID() {
	m.location = nil
	delete(m.clearedFields, pkg.FieldLocationID)
}

// SetStatus sets the "status" field.
func (m *PkgMutation) SetStatus(ts types.PackageStatus) {
	m.status = &ts
}

// Status returns the value of the "status" field in the mutation.
func (m *PkgMutation) Status() (r types.PackageStatus, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Pkg entity.
// If the Pkg object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PkgMutation) OldStatus(ctx context.Context) (v types.PackageStatus, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *PkgMutation) ResetStatus() {
	m.status = nil
}

// SetObjectKey sets the "object_key" field.
func (m *PkgMutation) SetObjectKey(u uuid.UUID) {
	m.object_key = &u
}

// ObjectKey returns the value of the "object_key" field in the mutation.
func (m *PkgMutation) ObjectKey() (r uuid.UUID, exists bool) {
	v := m.object_key
	if v == nil {
		return
	}
	return *v, true
}

// OldObjectKey returns the old "object_key" field's value of the Pkg entity.
// If the Pkg object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PkgMutation) OldObjectKey(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldObjectKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldObjectKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldObjectKey: %w", err)
	}
	return oldValue.ObjectKey, nil
}

// ResetObjectKey resets all changes to the "object_key" field.
func (m *PkgMutation) ResetObjectKey() {
	m.object_key = nil
}

// ClearLocation clears the "location" edge to the Location entity.
func (m *PkgMutation) ClearLocation() {
	m.clearedlocation = true
}

// LocationCleared reports if the "location" edge to the Location entity was cleared.
func (m *PkgMutation) LocationCleared() bool {
	return m.LocationIDCleared() || m.clearedlocation
}

// LocationIDs returns the "location" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// LocationID instead. It exists only for internal usage by the builders.
func (m *PkgMutation) LocationIDs() (ids []int) {
	if id := m.location; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetLocation resets all changes to the "location" edge.
func (m *PkgMutation) ResetLocation() {
	m.location = nil
	m.clearedlocation = false
}

// Where appends a list predicates to the PkgMutation builder.
func (m *PkgMutation) Where(ps ...predicate.Pkg) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *PkgMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Pkg).
func (m *PkgMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PkgMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, pkg.FieldName)
	}
	if m.aip_id != nil {
		fields = append(fields, pkg.FieldAipID)
	}
	if m.location != nil {
		fields = append(fields, pkg.FieldLocationID)
	}
	if m.status != nil {
		fields = append(fields, pkg.FieldStatus)
	}
	if m.object_key != nil {
		fields = append(fields, pkg.FieldObjectKey)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PkgMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case pkg.FieldName:
		return m.Name()
	case pkg.FieldAipID:
		return m.AipID()
	case pkg.FieldLocationID:
		return m.LocationID()
	case pkg.FieldStatus:
		return m.Status()
	case pkg.FieldObjectKey:
		return m.ObjectKey()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PkgMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case pkg.FieldName:
		return m.OldName(ctx)
	case pkg.FieldAipID:
		return m.OldAipID(ctx)
	case pkg.FieldLocationID:
		return m.OldLocationID(ctx)
	case pkg.FieldStatus:
		return m.OldStatus(ctx)
	case pkg.FieldObjectKey:
		return m.OldObjectKey(ctx)
	}
	return nil, fmt.Errorf("unknown Pkg field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PkgMutation) SetField(name string, value ent.Value) error {
	switch name {
	case pkg.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case pkg.FieldAipID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAipID(v)
		return nil
	case pkg.FieldLocationID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLocationID(v)
		return nil
	case pkg.FieldStatus:
		v, ok := value.(types.PackageStatus)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case pkg.FieldObjectKey:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetObjectKey(v)
		return nil
	}
	return fmt.Errorf("unknown Pkg field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PkgMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PkgMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PkgMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Pkg numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PkgMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(pkg.FieldLocationID) {
		fields = append(fields, pkg.FieldLocationID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PkgMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PkgMutation) ClearField(name string) error {
	switch name {
	case pkg.FieldLocationID:
		m.ClearLocationID()
		return nil
	}
	return fmt.Errorf("unknown Pkg nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PkgMutation) ResetField(name string) error {
	switch name {
	case pkg.FieldName:
		m.ResetName()
		return nil
	case pkg.FieldAipID:
		m.ResetAipID()
		return nil
	case pkg.FieldLocationID:
		m.ResetLocationID()
		return nil
	case pkg.FieldStatus:
		m.ResetStatus()
		return nil
	case pkg.FieldObjectKey:
		m.ResetObjectKey()
		return nil
	}
	return fmt.Errorf("unknown Pkg field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PkgMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.location != nil {
		edges = append(edges, pkg.EdgeLocation)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PkgMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case pkg.EdgeLocation:
		if id := m.location; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PkgMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PkgMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PkgMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedlocation {
		edges = append(edges, pkg.EdgeLocation)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PkgMutation) EdgeCleared(name string) bool {
	switch name {
	case pkg.EdgeLocation:
		return m.clearedlocation
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PkgMutation) ClearEdge(name string) error {
	switch name {
	case pkg.EdgeLocation:
		m.ClearLocation()
		return nil
	}
	return fmt.Errorf("unknown Pkg unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PkgMutation) ResetEdge(name string) error {
	switch name {
	case pkg.EdgeLocation:
		m.ResetLocation()
		return nil
	}
	return fmt.Errorf("unknown Pkg edge %s", name)
}
