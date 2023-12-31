// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/nonchan7720/go-storage-to-messenger/pkg/ent/outbox"
	"github.com/nonchan7720/go-storage-to-messenger/pkg/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeOutbox = "Outbox"
)

// OutboxMutation represents an operation that mutates the Outbox nodes in the graph.
type OutboxMutation struct {
	config
	op             Op
	typ            string
	id             *int64
	aggregate_type *string
	aggregate_id   *string
	event          *string
	payload        *[]byte
	retry_at       *time.Time
	retry_count    *int
	addretry_count *int
	clearedFields  map[string]struct{}
	done           bool
	oldValue       func(context.Context) (*Outbox, error)
	predicates     []predicate.Outbox
}

var _ ent.Mutation = (*OutboxMutation)(nil)

// outboxOption allows management of the mutation configuration using functional options.
type outboxOption func(*OutboxMutation)

// newOutboxMutation creates new mutation for the Outbox entity.
func newOutboxMutation(c config, op Op, opts ...outboxOption) *OutboxMutation {
	m := &OutboxMutation{
		config:        c,
		op:            op,
		typ:           TypeOutbox,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withOutboxID sets the ID field of the mutation.
func withOutboxID(id int64) outboxOption {
	return func(m *OutboxMutation) {
		var (
			err   error
			once  sync.Once
			value *Outbox
		)
		m.oldValue = func(ctx context.Context) (*Outbox, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Outbox.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withOutbox sets the old Outbox of the mutation.
func withOutbox(node *Outbox) outboxOption {
	return func(m *OutboxMutation) {
		m.oldValue = func(context.Context) (*Outbox, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m OutboxMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m OutboxMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Outbox entities.
func (m *OutboxMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *OutboxMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *OutboxMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Outbox.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAggregateType sets the "aggregate_type" field.
func (m *OutboxMutation) SetAggregateType(s string) {
	m.aggregate_type = &s
}

// AggregateType returns the value of the "aggregate_type" field in the mutation.
func (m *OutboxMutation) AggregateType() (r string, exists bool) {
	v := m.aggregate_type
	if v == nil {
		return
	}
	return *v, true
}

// OldAggregateType returns the old "aggregate_type" field's value of the Outbox entity.
// If the Outbox object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OutboxMutation) OldAggregateType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAggregateType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAggregateType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAggregateType: %w", err)
	}
	return oldValue.AggregateType, nil
}

// ResetAggregateType resets all changes to the "aggregate_type" field.
func (m *OutboxMutation) ResetAggregateType() {
	m.aggregate_type = nil
}

// SetAggregateID sets the "aggregate_id" field.
func (m *OutboxMutation) SetAggregateID(s string) {
	m.aggregate_id = &s
}

// AggregateID returns the value of the "aggregate_id" field in the mutation.
func (m *OutboxMutation) AggregateID() (r string, exists bool) {
	v := m.aggregate_id
	if v == nil {
		return
	}
	return *v, true
}

// OldAggregateID returns the old "aggregate_id" field's value of the Outbox entity.
// If the Outbox object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OutboxMutation) OldAggregateID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAggregateID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAggregateID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAggregateID: %w", err)
	}
	return oldValue.AggregateID, nil
}

// ResetAggregateID resets all changes to the "aggregate_id" field.
func (m *OutboxMutation) ResetAggregateID() {
	m.aggregate_id = nil
}

// SetEvent sets the "event" field.
func (m *OutboxMutation) SetEvent(s string) {
	m.event = &s
}

// Event returns the value of the "event" field in the mutation.
func (m *OutboxMutation) Event() (r string, exists bool) {
	v := m.event
	if v == nil {
		return
	}
	return *v, true
}

// OldEvent returns the old "event" field's value of the Outbox entity.
// If the Outbox object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OutboxMutation) OldEvent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEvent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEvent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEvent: %w", err)
	}
	return oldValue.Event, nil
}

// ResetEvent resets all changes to the "event" field.
func (m *OutboxMutation) ResetEvent() {
	m.event = nil
}

// SetPayload sets the "payload" field.
func (m *OutboxMutation) SetPayload(b []byte) {
	m.payload = &b
}

// Payload returns the value of the "payload" field in the mutation.
func (m *OutboxMutation) Payload() (r []byte, exists bool) {
	v := m.payload
	if v == nil {
		return
	}
	return *v, true
}

// OldPayload returns the old "payload" field's value of the Outbox entity.
// If the Outbox object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OutboxMutation) OldPayload(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPayload is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPayload requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPayload: %w", err)
	}
	return oldValue.Payload, nil
}

// ResetPayload resets all changes to the "payload" field.
func (m *OutboxMutation) ResetPayload() {
	m.payload = nil
}

// SetRetryAt sets the "retry_at" field.
func (m *OutboxMutation) SetRetryAt(t time.Time) {
	m.retry_at = &t
}

// RetryAt returns the value of the "retry_at" field in the mutation.
func (m *OutboxMutation) RetryAt() (r time.Time, exists bool) {
	v := m.retry_at
	if v == nil {
		return
	}
	return *v, true
}

// OldRetryAt returns the old "retry_at" field's value of the Outbox entity.
// If the Outbox object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OutboxMutation) OldRetryAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRetryAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRetryAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRetryAt: %w", err)
	}
	return oldValue.RetryAt, nil
}

// ResetRetryAt resets all changes to the "retry_at" field.
func (m *OutboxMutation) ResetRetryAt() {
	m.retry_at = nil
}

// SetRetryCount sets the "retry_count" field.
func (m *OutboxMutation) SetRetryCount(i int) {
	m.retry_count = &i
	m.addretry_count = nil
}

// RetryCount returns the value of the "retry_count" field in the mutation.
func (m *OutboxMutation) RetryCount() (r int, exists bool) {
	v := m.retry_count
	if v == nil {
		return
	}
	return *v, true
}

// OldRetryCount returns the old "retry_count" field's value of the Outbox entity.
// If the Outbox object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *OutboxMutation) OldRetryCount(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRetryCount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRetryCount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRetryCount: %w", err)
	}
	return oldValue.RetryCount, nil
}

// AddRetryCount adds i to the "retry_count" field.
func (m *OutboxMutation) AddRetryCount(i int) {
	if m.addretry_count != nil {
		*m.addretry_count += i
	} else {
		m.addretry_count = &i
	}
}

// AddedRetryCount returns the value that was added to the "retry_count" field in this mutation.
func (m *OutboxMutation) AddedRetryCount() (r int, exists bool) {
	v := m.addretry_count
	if v == nil {
		return
	}
	return *v, true
}

// ClearRetryCount clears the value of the "retry_count" field.
func (m *OutboxMutation) ClearRetryCount() {
	m.retry_count = nil
	m.addretry_count = nil
	m.clearedFields[outbox.FieldRetryCount] = struct{}{}
}

// RetryCountCleared returns if the "retry_count" field was cleared in this mutation.
func (m *OutboxMutation) RetryCountCleared() bool {
	_, ok := m.clearedFields[outbox.FieldRetryCount]
	return ok
}

// ResetRetryCount resets all changes to the "retry_count" field.
func (m *OutboxMutation) ResetRetryCount() {
	m.retry_count = nil
	m.addretry_count = nil
	delete(m.clearedFields, outbox.FieldRetryCount)
}

// Where appends a list predicates to the OutboxMutation builder.
func (m *OutboxMutation) Where(ps ...predicate.Outbox) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the OutboxMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *OutboxMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Outbox, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *OutboxMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *OutboxMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Outbox).
func (m *OutboxMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *OutboxMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.aggregate_type != nil {
		fields = append(fields, outbox.FieldAggregateType)
	}
	if m.aggregate_id != nil {
		fields = append(fields, outbox.FieldAggregateID)
	}
	if m.event != nil {
		fields = append(fields, outbox.FieldEvent)
	}
	if m.payload != nil {
		fields = append(fields, outbox.FieldPayload)
	}
	if m.retry_at != nil {
		fields = append(fields, outbox.FieldRetryAt)
	}
	if m.retry_count != nil {
		fields = append(fields, outbox.FieldRetryCount)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *OutboxMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case outbox.FieldAggregateType:
		return m.AggregateType()
	case outbox.FieldAggregateID:
		return m.AggregateID()
	case outbox.FieldEvent:
		return m.Event()
	case outbox.FieldPayload:
		return m.Payload()
	case outbox.FieldRetryAt:
		return m.RetryAt()
	case outbox.FieldRetryCount:
		return m.RetryCount()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *OutboxMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case outbox.FieldAggregateType:
		return m.OldAggregateType(ctx)
	case outbox.FieldAggregateID:
		return m.OldAggregateID(ctx)
	case outbox.FieldEvent:
		return m.OldEvent(ctx)
	case outbox.FieldPayload:
		return m.OldPayload(ctx)
	case outbox.FieldRetryAt:
		return m.OldRetryAt(ctx)
	case outbox.FieldRetryCount:
		return m.OldRetryCount(ctx)
	}
	return nil, fmt.Errorf("unknown Outbox field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OutboxMutation) SetField(name string, value ent.Value) error {
	switch name {
	case outbox.FieldAggregateType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAggregateType(v)
		return nil
	case outbox.FieldAggregateID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAggregateID(v)
		return nil
	case outbox.FieldEvent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEvent(v)
		return nil
	case outbox.FieldPayload:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPayload(v)
		return nil
	case outbox.FieldRetryAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRetryAt(v)
		return nil
	case outbox.FieldRetryCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRetryCount(v)
		return nil
	}
	return fmt.Errorf("unknown Outbox field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *OutboxMutation) AddedFields() []string {
	var fields []string
	if m.addretry_count != nil {
		fields = append(fields, outbox.FieldRetryCount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *OutboxMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case outbox.FieldRetryCount:
		return m.AddedRetryCount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *OutboxMutation) AddField(name string, value ent.Value) error {
	switch name {
	case outbox.FieldRetryCount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRetryCount(v)
		return nil
	}
	return fmt.Errorf("unknown Outbox numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *OutboxMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(outbox.FieldRetryCount) {
		fields = append(fields, outbox.FieldRetryCount)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *OutboxMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *OutboxMutation) ClearField(name string) error {
	switch name {
	case outbox.FieldRetryCount:
		m.ClearRetryCount()
		return nil
	}
	return fmt.Errorf("unknown Outbox nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *OutboxMutation) ResetField(name string) error {
	switch name {
	case outbox.FieldAggregateType:
		m.ResetAggregateType()
		return nil
	case outbox.FieldAggregateID:
		m.ResetAggregateID()
		return nil
	case outbox.FieldEvent:
		m.ResetEvent()
		return nil
	case outbox.FieldPayload:
		m.ResetPayload()
		return nil
	case outbox.FieldRetryAt:
		m.ResetRetryAt()
		return nil
	case outbox.FieldRetryCount:
		m.ResetRetryCount()
		return nil
	}
	return fmt.Errorf("unknown Outbox field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *OutboxMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *OutboxMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *OutboxMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *OutboxMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *OutboxMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *OutboxMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *OutboxMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Outbox unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *OutboxMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Outbox edge %s", name)
}
