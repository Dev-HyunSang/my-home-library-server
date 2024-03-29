// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dev-hyunsang/my-home-library-server/ent/book"
	"github.com/dev-hyunsang/my-home-library-server/ent/user"
	"github.com/google/uuid"
)

// BookCreate is the builder for creating a Book entity.
type BookCreate struct {
	config
	mutation *BookMutation
	hooks    []Hook
}

// SetUserUUID sets the "user_uuid" field.
func (bc *BookCreate) SetUserUUID(u uuid.UUID) *BookCreate {
	bc.mutation.SetUserUUID(u)
	return bc
}

// SetTitle sets the "title" field.
func (bc *BookCreate) SetTitle(s string) *BookCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetSubtitle sets the "subtitle" field.
func (bc *BookCreate) SetSubtitle(s string) *BookCreate {
	bc.mutation.SetSubtitle(s)
	return bc
}

// SetPublisher sets the "publisher" field.
func (bc *BookCreate) SetPublisher(s string) *BookCreate {
	bc.mutation.SetPublisher(s)
	return bc
}

// SetPublishingCompany sets the "publishing_company" field.
func (bc *BookCreate) SetPublishingCompany(s string) *BookCreate {
	bc.mutation.SetPublishingCompany(s)
	return bc
}

// SetMemo sets the "memo" field.
func (bc *BookCreate) SetMemo(s string) *BookCreate {
	bc.mutation.SetMemo(s)
	return bc
}

// SetTotalPage sets the "total_page" field.
func (bc *BookCreate) SetTotalPage(i int) *BookCreate {
	bc.mutation.SetTotalPage(i)
	return bc
}

// SetNillableTotalPage sets the "total_page" field if the given value is not nil.
func (bc *BookCreate) SetNillableTotalPage(i *int) *BookCreate {
	if i != nil {
		bc.SetTotalPage(*i)
	}
	return bc
}

// SetCurrentPage sets the "current_page" field.
func (bc *BookCreate) SetCurrentPage(i int) *BookCreate {
	bc.mutation.SetCurrentPage(i)
	return bc
}

// SetNillableCurrentPage sets the "current_page" field if the given value is not nil.
func (bc *BookCreate) SetNillableCurrentPage(i *int) *BookCreate {
	if i != nil {
		bc.SetCurrentPage(*i)
	}
	return bc
}

// SetCreatedAt sets the "created_at" field.
func (bc *BookCreate) SetCreatedAt(t time.Time) *BookCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BookCreate) SetNillableCreatedAt(t *time.Time) *BookCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetEditedAt sets the "edited_at" field.
func (bc *BookCreate) SetEditedAt(t time.Time) *BookCreate {
	bc.mutation.SetEditedAt(t)
	return bc
}

// SetID sets the "id" field.
func (bc *BookCreate) SetID(u uuid.UUID) *BookCreate {
	bc.mutation.SetID(u)
	return bc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bc *BookCreate) SetUserID(id uuid.UUID) *BookCreate {
	bc.mutation.SetUserID(id)
	return bc
}

// SetUser sets the "user" edge to the User entity.
func (bc *BookCreate) SetUser(u *User) *BookCreate {
	return bc.SetUserID(u.ID)
}

// Mutation returns the BookMutation object of the builder.
func (bc *BookCreate) Mutation() *BookMutation {
	return bc.mutation
}

// Save creates the Book in the database.
func (bc *BookCreate) Save(ctx context.Context) (*Book, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookCreate) SaveX(ctx context.Context) *Book {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookCreate) defaults() {
	if _, ok := bc.mutation.TotalPage(); !ok {
		v := book.DefaultTotalPage
		bc.mutation.SetTotalPage(v)
	}
	if _, ok := bc.mutation.CurrentPage(); !ok {
		v := book.DefaultCurrentPage
		bc.mutation.SetCurrentPage(v)
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := book.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookCreate) check() error {
	if _, ok := bc.mutation.UserUUID(); !ok {
		return &ValidationError{Name: "user_uuid", err: errors.New(`ent: missing required field "Book.user_uuid"`)}
	}
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Book.title"`)}
	}
	if _, ok := bc.mutation.Subtitle(); !ok {
		return &ValidationError{Name: "subtitle", err: errors.New(`ent: missing required field "Book.subtitle"`)}
	}
	if _, ok := bc.mutation.Publisher(); !ok {
		return &ValidationError{Name: "publisher", err: errors.New(`ent: missing required field "Book.publisher"`)}
	}
	if _, ok := bc.mutation.PublishingCompany(); !ok {
		return &ValidationError{Name: "publishing_company", err: errors.New(`ent: missing required field "Book.publishing_company"`)}
	}
	if _, ok := bc.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New(`ent: missing required field "Book.memo"`)}
	}
	if _, ok := bc.mutation.TotalPage(); !ok {
		return &ValidationError{Name: "total_page", err: errors.New(`ent: missing required field "Book.total_page"`)}
	}
	if _, ok := bc.mutation.CurrentPage(); !ok {
		return &ValidationError{Name: "current_page", err: errors.New(`ent: missing required field "Book.current_page"`)}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Book.created_at"`)}
	}
	if _, ok := bc.mutation.EditedAt(); !ok {
		return &ValidationError{Name: "edited_at", err: errors.New(`ent: missing required field "Book.edited_at"`)}
	}
	if _, ok := bc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Book.user"`)}
	}
	return nil
}

func (bc *BookCreate) sqlSave(ctx context.Context) (*Book, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookCreate) createSpec() (*Book, *sqlgraph.CreateSpec) {
	var (
		_node = &Book{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(book.Table, sqlgraph.NewFieldSpec(book.FieldID, field.TypeUUID))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Subtitle(); ok {
		_spec.SetField(book.FieldSubtitle, field.TypeString, value)
		_node.Subtitle = value
	}
	if value, ok := bc.mutation.Publisher(); ok {
		_spec.SetField(book.FieldPublisher, field.TypeString, value)
		_node.Publisher = value
	}
	if value, ok := bc.mutation.PublishingCompany(); ok {
		_spec.SetField(book.FieldPublishingCompany, field.TypeString, value)
		_node.PublishingCompany = value
	}
	if value, ok := bc.mutation.Memo(); ok {
		_spec.SetField(book.FieldMemo, field.TypeString, value)
		_node.Memo = value
	}
	if value, ok := bc.mutation.TotalPage(); ok {
		_spec.SetField(book.FieldTotalPage, field.TypeInt, value)
		_node.TotalPage = value
	}
	if value, ok := bc.mutation.CurrentPage(); ok {
		_spec.SetField(book.FieldCurrentPage, field.TypeInt, value)
		_node.CurrentPage = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(book.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.EditedAt(); ok {
		_spec.SetField(book.FieldEditedAt, field.TypeTime, value)
		_node.EditedAt = value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   book.UserTable,
			Columns: []string{book.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookCreateBulk is the builder for creating many Book entities in bulk.
type BookCreateBulk struct {
	config
	err      error
	builders []*BookCreate
}

// Save creates the Book entities in the database.
func (bcb *BookCreateBulk) Save(ctx context.Context) ([]*Book, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Book, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookCreateBulk) SaveX(ctx context.Context) []*Book {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
