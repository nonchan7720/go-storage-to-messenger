// Code generated by ent, DO NOT EDIT.

package outbox

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/nonchan7720/go-storage-to-messenger/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Outbox {
	return predicate.Outbox(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Outbox {
	return predicate.Outbox(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Outbox {
	return predicate.Outbox(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Outbox {
	return predicate.Outbox(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Outbox {
	return predicate.Outbox(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Outbox {
	return predicate.Outbox(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Outbox {
	return predicate.Outbox(sql.FieldLTE(FieldID, id))
}

// AggregateType applies equality check predicate on the "aggregate_type" field. It's identical to AggregateTypeEQ.
func AggregateType(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldAggregateType, v))
}

// AggregateID applies equality check predicate on the "aggregate_id" field. It's identical to AggregateIDEQ.
func AggregateID(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldAggregateID, v))
}

// Event applies equality check predicate on the "event" field. It's identical to EventEQ.
func Event(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldEvent, v))
}

// Payload applies equality check predicate on the "payload" field. It's identical to PayloadEQ.
func Payload(v []byte) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldPayload, v))
}

// RetryAt applies equality check predicate on the "retry_at" field. It's identical to RetryAtEQ.
func RetryAt(v time.Time) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldRetryAt, v))
}

// RetryCount applies equality check predicate on the "retry_count" field. It's identical to RetryCountEQ.
func RetryCount(v int) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldRetryCount, v))
}

// AggregateTypeEQ applies the EQ predicate on the "aggregate_type" field.
func AggregateTypeEQ(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldAggregateType, v))
}

// AggregateTypeNEQ applies the NEQ predicate on the "aggregate_type" field.
func AggregateTypeNEQ(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldNEQ(FieldAggregateType, v))
}

// AggregateTypeIn applies the In predicate on the "aggregate_type" field.
func AggregateTypeIn(vs ...string) predicate.Outbox {
	return predicate.Outbox(sql.FieldIn(FieldAggregateType, vs...))
}

// AggregateTypeNotIn applies the NotIn predicate on the "aggregate_type" field.
func AggregateTypeNotIn(vs ...string) predicate.Outbox {
	return predicate.Outbox(sql.FieldNotIn(FieldAggregateType, vs...))
}

// AggregateTypeGT applies the GT predicate on the "aggregate_type" field.
func AggregateTypeGT(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldGT(FieldAggregateType, v))
}

// AggregateTypeGTE applies the GTE predicate on the "aggregate_type" field.
func AggregateTypeGTE(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldGTE(FieldAggregateType, v))
}

// AggregateTypeLT applies the LT predicate on the "aggregate_type" field.
func AggregateTypeLT(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldLT(FieldAggregateType, v))
}

// AggregateTypeLTE applies the LTE predicate on the "aggregate_type" field.
func AggregateTypeLTE(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldLTE(FieldAggregateType, v))
}

// AggregateTypeContains applies the Contains predicate on the "aggregate_type" field.
func AggregateTypeContains(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldContains(FieldAggregateType, v))
}

// AggregateTypeHasPrefix applies the HasPrefix predicate on the "aggregate_type" field.
func AggregateTypeHasPrefix(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldHasPrefix(FieldAggregateType, v))
}

// AggregateTypeHasSuffix applies the HasSuffix predicate on the "aggregate_type" field.
func AggregateTypeHasSuffix(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldHasSuffix(FieldAggregateType, v))
}

// AggregateTypeEqualFold applies the EqualFold predicate on the "aggregate_type" field.
func AggregateTypeEqualFold(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldEqualFold(FieldAggregateType, v))
}

// AggregateTypeContainsFold applies the ContainsFold predicate on the "aggregate_type" field.
func AggregateTypeContainsFold(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldContainsFold(FieldAggregateType, v))
}

// AggregateIDEQ applies the EQ predicate on the "aggregate_id" field.
func AggregateIDEQ(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldAggregateID, v))
}

// AggregateIDNEQ applies the NEQ predicate on the "aggregate_id" field.
func AggregateIDNEQ(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldNEQ(FieldAggregateID, v))
}

// AggregateIDIn applies the In predicate on the "aggregate_id" field.
func AggregateIDIn(vs ...string) predicate.Outbox {
	return predicate.Outbox(sql.FieldIn(FieldAggregateID, vs...))
}

// AggregateIDNotIn applies the NotIn predicate on the "aggregate_id" field.
func AggregateIDNotIn(vs ...string) predicate.Outbox {
	return predicate.Outbox(sql.FieldNotIn(FieldAggregateID, vs...))
}

// AggregateIDGT applies the GT predicate on the "aggregate_id" field.
func AggregateIDGT(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldGT(FieldAggregateID, v))
}

// AggregateIDGTE applies the GTE predicate on the "aggregate_id" field.
func AggregateIDGTE(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldGTE(FieldAggregateID, v))
}

// AggregateIDLT applies the LT predicate on the "aggregate_id" field.
func AggregateIDLT(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldLT(FieldAggregateID, v))
}

// AggregateIDLTE applies the LTE predicate on the "aggregate_id" field.
func AggregateIDLTE(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldLTE(FieldAggregateID, v))
}

// AggregateIDContains applies the Contains predicate on the "aggregate_id" field.
func AggregateIDContains(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldContains(FieldAggregateID, v))
}

// AggregateIDHasPrefix applies the HasPrefix predicate on the "aggregate_id" field.
func AggregateIDHasPrefix(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldHasPrefix(FieldAggregateID, v))
}

// AggregateIDHasSuffix applies the HasSuffix predicate on the "aggregate_id" field.
func AggregateIDHasSuffix(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldHasSuffix(FieldAggregateID, v))
}

// AggregateIDEqualFold applies the EqualFold predicate on the "aggregate_id" field.
func AggregateIDEqualFold(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldEqualFold(FieldAggregateID, v))
}

// AggregateIDContainsFold applies the ContainsFold predicate on the "aggregate_id" field.
func AggregateIDContainsFold(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldContainsFold(FieldAggregateID, v))
}

// EventEQ applies the EQ predicate on the "event" field.
func EventEQ(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldEvent, v))
}

// EventNEQ applies the NEQ predicate on the "event" field.
func EventNEQ(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldNEQ(FieldEvent, v))
}

// EventIn applies the In predicate on the "event" field.
func EventIn(vs ...string) predicate.Outbox {
	return predicate.Outbox(sql.FieldIn(FieldEvent, vs...))
}

// EventNotIn applies the NotIn predicate on the "event" field.
func EventNotIn(vs ...string) predicate.Outbox {
	return predicate.Outbox(sql.FieldNotIn(FieldEvent, vs...))
}

// EventGT applies the GT predicate on the "event" field.
func EventGT(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldGT(FieldEvent, v))
}

// EventGTE applies the GTE predicate on the "event" field.
func EventGTE(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldGTE(FieldEvent, v))
}

// EventLT applies the LT predicate on the "event" field.
func EventLT(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldLT(FieldEvent, v))
}

// EventLTE applies the LTE predicate on the "event" field.
func EventLTE(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldLTE(FieldEvent, v))
}

// EventContains applies the Contains predicate on the "event" field.
func EventContains(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldContains(FieldEvent, v))
}

// EventHasPrefix applies the HasPrefix predicate on the "event" field.
func EventHasPrefix(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldHasPrefix(FieldEvent, v))
}

// EventHasSuffix applies the HasSuffix predicate on the "event" field.
func EventHasSuffix(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldHasSuffix(FieldEvent, v))
}

// EventEqualFold applies the EqualFold predicate on the "event" field.
func EventEqualFold(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldEqualFold(FieldEvent, v))
}

// EventContainsFold applies the ContainsFold predicate on the "event" field.
func EventContainsFold(v string) predicate.Outbox {
	return predicate.Outbox(sql.FieldContainsFold(FieldEvent, v))
}

// PayloadEQ applies the EQ predicate on the "payload" field.
func PayloadEQ(v []byte) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldPayload, v))
}

// PayloadNEQ applies the NEQ predicate on the "payload" field.
func PayloadNEQ(v []byte) predicate.Outbox {
	return predicate.Outbox(sql.FieldNEQ(FieldPayload, v))
}

// PayloadIn applies the In predicate on the "payload" field.
func PayloadIn(vs ...[]byte) predicate.Outbox {
	return predicate.Outbox(sql.FieldIn(FieldPayload, vs...))
}

// PayloadNotIn applies the NotIn predicate on the "payload" field.
func PayloadNotIn(vs ...[]byte) predicate.Outbox {
	return predicate.Outbox(sql.FieldNotIn(FieldPayload, vs...))
}

// PayloadGT applies the GT predicate on the "payload" field.
func PayloadGT(v []byte) predicate.Outbox {
	return predicate.Outbox(sql.FieldGT(FieldPayload, v))
}

// PayloadGTE applies the GTE predicate on the "payload" field.
func PayloadGTE(v []byte) predicate.Outbox {
	return predicate.Outbox(sql.FieldGTE(FieldPayload, v))
}

// PayloadLT applies the LT predicate on the "payload" field.
func PayloadLT(v []byte) predicate.Outbox {
	return predicate.Outbox(sql.FieldLT(FieldPayload, v))
}

// PayloadLTE applies the LTE predicate on the "payload" field.
func PayloadLTE(v []byte) predicate.Outbox {
	return predicate.Outbox(sql.FieldLTE(FieldPayload, v))
}

// RetryAtEQ applies the EQ predicate on the "retry_at" field.
func RetryAtEQ(v time.Time) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldRetryAt, v))
}

// RetryAtNEQ applies the NEQ predicate on the "retry_at" field.
func RetryAtNEQ(v time.Time) predicate.Outbox {
	return predicate.Outbox(sql.FieldNEQ(FieldRetryAt, v))
}

// RetryAtIn applies the In predicate on the "retry_at" field.
func RetryAtIn(vs ...time.Time) predicate.Outbox {
	return predicate.Outbox(sql.FieldIn(FieldRetryAt, vs...))
}

// RetryAtNotIn applies the NotIn predicate on the "retry_at" field.
func RetryAtNotIn(vs ...time.Time) predicate.Outbox {
	return predicate.Outbox(sql.FieldNotIn(FieldRetryAt, vs...))
}

// RetryAtGT applies the GT predicate on the "retry_at" field.
func RetryAtGT(v time.Time) predicate.Outbox {
	return predicate.Outbox(sql.FieldGT(FieldRetryAt, v))
}

// RetryAtGTE applies the GTE predicate on the "retry_at" field.
func RetryAtGTE(v time.Time) predicate.Outbox {
	return predicate.Outbox(sql.FieldGTE(FieldRetryAt, v))
}

// RetryAtLT applies the LT predicate on the "retry_at" field.
func RetryAtLT(v time.Time) predicate.Outbox {
	return predicate.Outbox(sql.FieldLT(FieldRetryAt, v))
}

// RetryAtLTE applies the LTE predicate on the "retry_at" field.
func RetryAtLTE(v time.Time) predicate.Outbox {
	return predicate.Outbox(sql.FieldLTE(FieldRetryAt, v))
}

// RetryCountEQ applies the EQ predicate on the "retry_count" field.
func RetryCountEQ(v int) predicate.Outbox {
	return predicate.Outbox(sql.FieldEQ(FieldRetryCount, v))
}

// RetryCountNEQ applies the NEQ predicate on the "retry_count" field.
func RetryCountNEQ(v int) predicate.Outbox {
	return predicate.Outbox(sql.FieldNEQ(FieldRetryCount, v))
}

// RetryCountIn applies the In predicate on the "retry_count" field.
func RetryCountIn(vs ...int) predicate.Outbox {
	return predicate.Outbox(sql.FieldIn(FieldRetryCount, vs...))
}

// RetryCountNotIn applies the NotIn predicate on the "retry_count" field.
func RetryCountNotIn(vs ...int) predicate.Outbox {
	return predicate.Outbox(sql.FieldNotIn(FieldRetryCount, vs...))
}

// RetryCountGT applies the GT predicate on the "retry_count" field.
func RetryCountGT(v int) predicate.Outbox {
	return predicate.Outbox(sql.FieldGT(FieldRetryCount, v))
}

// RetryCountGTE applies the GTE predicate on the "retry_count" field.
func RetryCountGTE(v int) predicate.Outbox {
	return predicate.Outbox(sql.FieldGTE(FieldRetryCount, v))
}

// RetryCountLT applies the LT predicate on the "retry_count" field.
func RetryCountLT(v int) predicate.Outbox {
	return predicate.Outbox(sql.FieldLT(FieldRetryCount, v))
}

// RetryCountLTE applies the LTE predicate on the "retry_count" field.
func RetryCountLTE(v int) predicate.Outbox {
	return predicate.Outbox(sql.FieldLTE(FieldRetryCount, v))
}

// RetryCountIsNil applies the IsNil predicate on the "retry_count" field.
func RetryCountIsNil() predicate.Outbox {
	return predicate.Outbox(sql.FieldIsNull(FieldRetryCount))
}

// RetryCountNotNil applies the NotNil predicate on the "retry_count" field.
func RetryCountNotNil() predicate.Outbox {
	return predicate.Outbox(sql.FieldNotNull(FieldRetryCount))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Outbox) predicate.Outbox {
	return predicate.Outbox(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Outbox) predicate.Outbox {
	return predicate.Outbox(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Outbox) predicate.Outbox {
	return predicate.Outbox(sql.NotPredicates(p))
}
