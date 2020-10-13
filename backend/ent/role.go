// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/B6102647/app/ent/role"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ROLEID holds the value of the "ROLE_ID" field.
	ROLEID int `json:"ROLE_ID,omitempty"`
	// ROLENAME holds the value of the "ROLE_NAME" field.
	ROLENAME string `json:"ROLE_NAME,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges RoleEdges `json:"edges"`
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Role holds the value of the Role edge.
	Role []*User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) RoleOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "Role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // ROLE_ID
		&sql.NullString{}, // ROLE_NAME
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(values ...interface{}) error {
	if m, n := len(values), len(role.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field ROLE_ID", values[0])
	} else if value.Valid {
		r.ROLEID = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field ROLE_NAME", values[1])
	} else if value.Valid {
		r.ROLENAME = value.String
	}
	return nil
}

// QueryRole queries the Role edge of the Role.
func (r *Role) QueryRole() *UserQuery {
	return (&RoleClient{config: r.config}).QueryRole(r)
}

// Update returns a builder for updating this Role.
// Note that, you need to call Role.Unwrap() before calling this method, if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return (&RoleClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", ROLE_ID=")
	builder.WriteString(fmt.Sprintf("%v", r.ROLEID))
	builder.WriteString(", ROLE_NAME=")
	builder.WriteString(r.ROLENAME)
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role

func (r Roles) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
