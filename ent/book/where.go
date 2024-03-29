// Code generated by ent, DO NOT EDIT.

package book

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/dev-hyunsang/my-home-library-server/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldID, id))
}

// UserUUID applies equality check predicate on the "user_uuid" field. It's identical to UserUUIDEQ.
func UserUUID(v uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUserUUID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// Subtitle applies equality check predicate on the "subtitle" field. It's identical to SubtitleEQ.
func Subtitle(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldSubtitle, v))
}

// Publisher applies equality check predicate on the "publisher" field. It's identical to PublisherEQ.
func Publisher(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublisher, v))
}

// PublishingCompany applies equality check predicate on the "publishing_company" field. It's identical to PublishingCompanyEQ.
func PublishingCompany(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublishingCompany, v))
}

// Memo applies equality check predicate on the "memo" field. It's identical to MemoEQ.
func Memo(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldMemo, v))
}

// TotalPage applies equality check predicate on the "total_page" field. It's identical to TotalPageEQ.
func TotalPage(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTotalPage, v))
}

// CurrentPage applies equality check predicate on the "current_page" field. It's identical to CurrentPageEQ.
func CurrentPage(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCurrentPage, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedAt, v))
}

// EditedAt applies equality check predicate on the "edited_at" field. It's identical to EditedAtEQ.
func EditedAt(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldEditedAt, v))
}

// UserUUIDEQ applies the EQ predicate on the "user_uuid" field.
func UserUUIDEQ(v uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldUserUUID, v))
}

// UserUUIDNEQ applies the NEQ predicate on the "user_uuid" field.
func UserUUIDNEQ(v uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldUserUUID, v))
}

// UserUUIDIn applies the In predicate on the "user_uuid" field.
func UserUUIDIn(vs ...uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldUserUUID, vs...))
}

// UserUUIDNotIn applies the NotIn predicate on the "user_uuid" field.
func UserUUIDNotIn(vs ...uuid.UUID) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldUserUUID, vs...))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldTitle, v))
}

// SubtitleEQ applies the EQ predicate on the "subtitle" field.
func SubtitleEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldSubtitle, v))
}

// SubtitleNEQ applies the NEQ predicate on the "subtitle" field.
func SubtitleNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldSubtitle, v))
}

// SubtitleIn applies the In predicate on the "subtitle" field.
func SubtitleIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldSubtitle, vs...))
}

// SubtitleNotIn applies the NotIn predicate on the "subtitle" field.
func SubtitleNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldSubtitle, vs...))
}

// SubtitleGT applies the GT predicate on the "subtitle" field.
func SubtitleGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldSubtitle, v))
}

// SubtitleGTE applies the GTE predicate on the "subtitle" field.
func SubtitleGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldSubtitle, v))
}

// SubtitleLT applies the LT predicate on the "subtitle" field.
func SubtitleLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldSubtitle, v))
}

// SubtitleLTE applies the LTE predicate on the "subtitle" field.
func SubtitleLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldSubtitle, v))
}

// SubtitleContains applies the Contains predicate on the "subtitle" field.
func SubtitleContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldSubtitle, v))
}

// SubtitleHasPrefix applies the HasPrefix predicate on the "subtitle" field.
func SubtitleHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldSubtitle, v))
}

// SubtitleHasSuffix applies the HasSuffix predicate on the "subtitle" field.
func SubtitleHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldSubtitle, v))
}

// SubtitleEqualFold applies the EqualFold predicate on the "subtitle" field.
func SubtitleEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldSubtitle, v))
}

// SubtitleContainsFold applies the ContainsFold predicate on the "subtitle" field.
func SubtitleContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldSubtitle, v))
}

// PublisherEQ applies the EQ predicate on the "publisher" field.
func PublisherEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublisher, v))
}

// PublisherNEQ applies the NEQ predicate on the "publisher" field.
func PublisherNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldPublisher, v))
}

// PublisherIn applies the In predicate on the "publisher" field.
func PublisherIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldPublisher, vs...))
}

// PublisherNotIn applies the NotIn predicate on the "publisher" field.
func PublisherNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldPublisher, vs...))
}

// PublisherGT applies the GT predicate on the "publisher" field.
func PublisherGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldPublisher, v))
}

// PublisherGTE applies the GTE predicate on the "publisher" field.
func PublisherGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldPublisher, v))
}

// PublisherLT applies the LT predicate on the "publisher" field.
func PublisherLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldPublisher, v))
}

// PublisherLTE applies the LTE predicate on the "publisher" field.
func PublisherLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldPublisher, v))
}

// PublisherContains applies the Contains predicate on the "publisher" field.
func PublisherContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldPublisher, v))
}

// PublisherHasPrefix applies the HasPrefix predicate on the "publisher" field.
func PublisherHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldPublisher, v))
}

// PublisherHasSuffix applies the HasSuffix predicate on the "publisher" field.
func PublisherHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldPublisher, v))
}

// PublisherEqualFold applies the EqualFold predicate on the "publisher" field.
func PublisherEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldPublisher, v))
}

// PublisherContainsFold applies the ContainsFold predicate on the "publisher" field.
func PublisherContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldPublisher, v))
}

// PublishingCompanyEQ applies the EQ predicate on the "publishing_company" field.
func PublishingCompanyEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldPublishingCompany, v))
}

// PublishingCompanyNEQ applies the NEQ predicate on the "publishing_company" field.
func PublishingCompanyNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldPublishingCompany, v))
}

// PublishingCompanyIn applies the In predicate on the "publishing_company" field.
func PublishingCompanyIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldPublishingCompany, vs...))
}

// PublishingCompanyNotIn applies the NotIn predicate on the "publishing_company" field.
func PublishingCompanyNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldPublishingCompany, vs...))
}

// PublishingCompanyGT applies the GT predicate on the "publishing_company" field.
func PublishingCompanyGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldPublishingCompany, v))
}

// PublishingCompanyGTE applies the GTE predicate on the "publishing_company" field.
func PublishingCompanyGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldPublishingCompany, v))
}

// PublishingCompanyLT applies the LT predicate on the "publishing_company" field.
func PublishingCompanyLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldPublishingCompany, v))
}

// PublishingCompanyLTE applies the LTE predicate on the "publishing_company" field.
func PublishingCompanyLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldPublishingCompany, v))
}

// PublishingCompanyContains applies the Contains predicate on the "publishing_company" field.
func PublishingCompanyContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldPublishingCompany, v))
}

// PublishingCompanyHasPrefix applies the HasPrefix predicate on the "publishing_company" field.
func PublishingCompanyHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldPublishingCompany, v))
}

// PublishingCompanyHasSuffix applies the HasSuffix predicate on the "publishing_company" field.
func PublishingCompanyHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldPublishingCompany, v))
}

// PublishingCompanyEqualFold applies the EqualFold predicate on the "publishing_company" field.
func PublishingCompanyEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldPublishingCompany, v))
}

// PublishingCompanyContainsFold applies the ContainsFold predicate on the "publishing_company" field.
func PublishingCompanyContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldPublishingCompany, v))
}

// MemoEQ applies the EQ predicate on the "memo" field.
func MemoEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldMemo, v))
}

// MemoNEQ applies the NEQ predicate on the "memo" field.
func MemoNEQ(v string) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldMemo, v))
}

// MemoIn applies the In predicate on the "memo" field.
func MemoIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldMemo, vs...))
}

// MemoNotIn applies the NotIn predicate on the "memo" field.
func MemoNotIn(vs ...string) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldMemo, vs...))
}

// MemoGT applies the GT predicate on the "memo" field.
func MemoGT(v string) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldMemo, v))
}

// MemoGTE applies the GTE predicate on the "memo" field.
func MemoGTE(v string) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldMemo, v))
}

// MemoLT applies the LT predicate on the "memo" field.
func MemoLT(v string) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldMemo, v))
}

// MemoLTE applies the LTE predicate on the "memo" field.
func MemoLTE(v string) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldMemo, v))
}

// MemoContains applies the Contains predicate on the "memo" field.
func MemoContains(v string) predicate.Book {
	return predicate.Book(sql.FieldContains(FieldMemo, v))
}

// MemoHasPrefix applies the HasPrefix predicate on the "memo" field.
func MemoHasPrefix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasPrefix(FieldMemo, v))
}

// MemoHasSuffix applies the HasSuffix predicate on the "memo" field.
func MemoHasSuffix(v string) predicate.Book {
	return predicate.Book(sql.FieldHasSuffix(FieldMemo, v))
}

// MemoEqualFold applies the EqualFold predicate on the "memo" field.
func MemoEqualFold(v string) predicate.Book {
	return predicate.Book(sql.FieldEqualFold(FieldMemo, v))
}

// MemoContainsFold applies the ContainsFold predicate on the "memo" field.
func MemoContainsFold(v string) predicate.Book {
	return predicate.Book(sql.FieldContainsFold(FieldMemo, v))
}

// TotalPageEQ applies the EQ predicate on the "total_page" field.
func TotalPageEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldTotalPage, v))
}

// TotalPageNEQ applies the NEQ predicate on the "total_page" field.
func TotalPageNEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldTotalPage, v))
}

// TotalPageIn applies the In predicate on the "total_page" field.
func TotalPageIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldTotalPage, vs...))
}

// TotalPageNotIn applies the NotIn predicate on the "total_page" field.
func TotalPageNotIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldTotalPage, vs...))
}

// TotalPageGT applies the GT predicate on the "total_page" field.
func TotalPageGT(v int) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldTotalPage, v))
}

// TotalPageGTE applies the GTE predicate on the "total_page" field.
func TotalPageGTE(v int) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldTotalPage, v))
}

// TotalPageLT applies the LT predicate on the "total_page" field.
func TotalPageLT(v int) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldTotalPage, v))
}

// TotalPageLTE applies the LTE predicate on the "total_page" field.
func TotalPageLTE(v int) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldTotalPage, v))
}

// CurrentPageEQ applies the EQ predicate on the "current_page" field.
func CurrentPageEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCurrentPage, v))
}

// CurrentPageNEQ applies the NEQ predicate on the "current_page" field.
func CurrentPageNEQ(v int) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCurrentPage, v))
}

// CurrentPageIn applies the In predicate on the "current_page" field.
func CurrentPageIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCurrentPage, vs...))
}

// CurrentPageNotIn applies the NotIn predicate on the "current_page" field.
func CurrentPageNotIn(vs ...int) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCurrentPage, vs...))
}

// CurrentPageGT applies the GT predicate on the "current_page" field.
func CurrentPageGT(v int) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCurrentPage, v))
}

// CurrentPageGTE applies the GTE predicate on the "current_page" field.
func CurrentPageGTE(v int) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCurrentPage, v))
}

// CurrentPageLT applies the LT predicate on the "current_page" field.
func CurrentPageLT(v int) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCurrentPage, v))
}

// CurrentPageLTE applies the LTE predicate on the "current_page" field.
func CurrentPageLTE(v int) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCurrentPage, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldCreatedAt, v))
}

// EditedAtEQ applies the EQ predicate on the "edited_at" field.
func EditedAtEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldEQ(FieldEditedAt, v))
}

// EditedAtNEQ applies the NEQ predicate on the "edited_at" field.
func EditedAtNEQ(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldNEQ(FieldEditedAt, v))
}

// EditedAtIn applies the In predicate on the "edited_at" field.
func EditedAtIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldIn(FieldEditedAt, vs...))
}

// EditedAtNotIn applies the NotIn predicate on the "edited_at" field.
func EditedAtNotIn(vs ...time.Time) predicate.Book {
	return predicate.Book(sql.FieldNotIn(FieldEditedAt, vs...))
}

// EditedAtGT applies the GT predicate on the "edited_at" field.
func EditedAtGT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGT(FieldEditedAt, v))
}

// EditedAtGTE applies the GTE predicate on the "edited_at" field.
func EditedAtGTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldGTE(FieldEditedAt, v))
}

// EditedAtLT applies the LT predicate on the "edited_at" field.
func EditedAtLT(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLT(FieldEditedAt, v))
}

// EditedAtLTE applies the LTE predicate on the "edited_at" field.
func EditedAtLTE(v time.Time) predicate.Book {
	return predicate.Book(sql.FieldLTE(FieldEditedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Book) predicate.Book {
	return predicate.Book(sql.NotPredicates(p))
}
