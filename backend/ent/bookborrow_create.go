// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/B6102647/app/ent/book"
	"github.com/B6102647/app/ent/bookborrow"
	"github.com/B6102647/app/ent/purpose"
	"github.com/B6102647/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// BookBorrowCreate is the builder for creating a BookBorrow entity.
type BookBorrowCreate struct {
	config
	mutation *BookBorrowMutation
	hooks    []Hook
}

// SetADDEDTIME sets the ADDED_TIME field.
func (bbc *BookBorrowCreate) SetADDEDTIME(t time.Time) *BookBorrowCreate {
	bbc.mutation.SetADDEDTIME(t)
	return bbc
}

// SetUSERID sets the USER edge to User by id.
func (bbc *BookBorrowCreate) SetUSERID(id int) *BookBorrowCreate {
	bbc.mutation.SetUSERID(id)
	return bbc
}

// SetNillableUSERID sets the USER edge to User by id if the given value is not nil.
func (bbc *BookBorrowCreate) SetNillableUSERID(id *int) *BookBorrowCreate {
	if id != nil {
		bbc = bbc.SetUSERID(*id)
	}
	return bbc
}

// SetUSER sets the USER edge to User.
func (bbc *BookBorrowCreate) SetUSER(u *User) *BookBorrowCreate {
	return bbc.SetUSERID(u.ID)
}

// SetBOOKID sets the BOOK edge to Book by id.
func (bbc *BookBorrowCreate) SetBOOKID(id int) *BookBorrowCreate {
	bbc.mutation.SetBOOKID(id)
	return bbc
}

// SetNillableBOOKID sets the BOOK edge to Book by id if the given value is not nil.
func (bbc *BookBorrowCreate) SetNillableBOOKID(id *int) *BookBorrowCreate {
	if id != nil {
		bbc = bbc.SetBOOKID(*id)
	}
	return bbc
}

// SetBOOK sets the BOOK edge to Book.
func (bbc *BookBorrowCreate) SetBOOK(b *Book) *BookBorrowCreate {
	return bbc.SetBOOKID(b.ID)
}

// SetPURPOSEID sets the PURPOSE edge to Purpose by id.
func (bbc *BookBorrowCreate) SetPURPOSEID(id int) *BookBorrowCreate {
	bbc.mutation.SetPURPOSEID(id)
	return bbc
}

// SetNillablePURPOSEID sets the PURPOSE edge to Purpose by id if the given value is not nil.
func (bbc *BookBorrowCreate) SetNillablePURPOSEID(id *int) *BookBorrowCreate {
	if id != nil {
		bbc = bbc.SetPURPOSEID(*id)
	}
	return bbc
}

// SetPURPOSE sets the PURPOSE edge to Purpose.
func (bbc *BookBorrowCreate) SetPURPOSE(p *Purpose) *BookBorrowCreate {
	return bbc.SetPURPOSEID(p.ID)
}

// Mutation returns the BookBorrowMutation object of the builder.
func (bbc *BookBorrowCreate) Mutation() *BookBorrowMutation {
	return bbc.mutation
}

// Save creates the BookBorrow in the database.
func (bbc *BookBorrowCreate) Save(ctx context.Context) (*BookBorrow, error) {
	if _, ok := bbc.mutation.ADDEDTIME(); !ok {
		return nil, &ValidationError{Name: "ADDED_TIME", err: errors.New("ent: missing required field \"ADDED_TIME\"")}
	}
	var (
		err  error
		node *BookBorrow
	)
	if len(bbc.hooks) == 0 {
		node, err = bbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookBorrowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bbc.mutation = mutation
			node, err = bbc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bbc.hooks) - 1; i >= 0; i-- {
			mut = bbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bbc *BookBorrowCreate) SaveX(ctx context.Context) *BookBorrow {
	v, err := bbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bbc *BookBorrowCreate) sqlSave(ctx context.Context) (*BookBorrow, error) {
	bb, _spec := bbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bbc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	bb.ID = int(id)
	return bb, nil
}

func (bbc *BookBorrowCreate) createSpec() (*BookBorrow, *sqlgraph.CreateSpec) {
	var (
		bb    = &BookBorrow{config: bbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bookborrow.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookborrow.FieldID,
			},
		}
	)
	if value, ok := bbc.mutation.ADDEDTIME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bookborrow.FieldADDEDTIME,
		})
		bb.ADDEDTIME = value
	}
	if nodes := bbc.mutation.USERIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.USERTable,
			Columns: []string{bookborrow.USERColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bbc.mutation.BOOKIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.BOOKTable,
			Columns: []string{bookborrow.BOOKColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: book.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bbc.mutation.PURPOSEIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bookborrow.PURPOSETable,
			Columns: []string{bookborrow.PURPOSEColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: purpose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return bb, _spec
}
