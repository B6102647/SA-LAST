// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/B6102647/app/ent/migrate"

	"github.com/B6102647/app/ent/book"
	"github.com/B6102647/app/ent/bookborrow"
	"github.com/B6102647/app/ent/purpose"
	"github.com/B6102647/app/ent/role"
	"github.com/B6102647/app/ent/status"
	"github.com/B6102647/app/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Book is the client for interacting with the Book builders.
	Book *BookClient
	// BookBorrow is the client for interacting with the BookBorrow builders.
	BookBorrow *BookBorrowClient
	// Purpose is the client for interacting with the Purpose builders.
	Purpose *PurposeClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// Status is the client for interacting with the Status builders.
	Status *StatusClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Book = NewBookClient(c.config)
	c.BookBorrow = NewBookBorrowClient(c.config)
	c.Purpose = NewPurposeClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.Status = NewStatusClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Book:       NewBookClient(cfg),
		BookBorrow: NewBookBorrowClient(cfg),
		Purpose:    NewPurposeClient(cfg),
		Role:       NewRoleClient(cfg),
		Status:     NewStatusClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:     cfg,
		Book:       NewBookClient(cfg),
		BookBorrow: NewBookBorrowClient(cfg),
		Purpose:    NewPurposeClient(cfg),
		Role:       NewRoleClient(cfg),
		Status:     NewStatusClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Book.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Book.Use(hooks...)
	c.BookBorrow.Use(hooks...)
	c.Purpose.Use(hooks...)
	c.Role.Use(hooks...)
	c.Status.Use(hooks...)
	c.User.Use(hooks...)
}

// BookClient is a client for the Book schema.
type BookClient struct {
	config
}

// NewBookClient returns a client for the Book from the given config.
func NewBookClient(c config) *BookClient {
	return &BookClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `book.Hooks(f(g(h())))`.
func (c *BookClient) Use(hooks ...Hook) {
	c.hooks.Book = append(c.hooks.Book, hooks...)
}

// Create returns a create builder for Book.
func (c *BookClient) Create() *BookCreate {
	mutation := newBookMutation(c.config, OpCreate)
	return &BookCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Book.
func (c *BookClient) Update() *BookUpdate {
	mutation := newBookMutation(c.config, OpUpdate)
	return &BookUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookClient) UpdateOne(b *Book) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBook(b))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookClient) UpdateOneID(id int) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBookID(id))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Book.
func (c *BookClient) Delete() *BookDelete {
	mutation := newBookMutation(c.config, OpDelete)
	return &BookDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BookClient) DeleteOne(b *Book) *BookDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BookClient) DeleteOneID(id int) *BookDeleteOne {
	builder := c.Delete().Where(book.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookDeleteOne{builder}
}

// Create returns a query builder for Book.
func (c *BookClient) Query() *BookQuery {
	return &BookQuery{config: c.config}
}

// Get returns a Book entity by its id.
func (c *BookClient) Get(ctx context.Context, id int) (*Book, error) {
	return c.Query().Where(book.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookClient) GetX(ctx context.Context, id int) *Book {
	b, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return b
}

// QueryBooklist queries the Booklist edge of a Book.
func (c *BookClient) QueryBooklist(b *Book) *BookBorrowQuery {
	query := &BookBorrowQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, id),
			sqlgraph.To(bookborrow.Table, bookborrow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, book.BooklistTable, book.BooklistColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStatus queries the Status edge of a Book.
func (c *BookClient) QueryStatus(b *Book) *StatusQuery {
	query := &StatusQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, id),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, book.StatusTable, book.StatusColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BookClient) Hooks() []Hook {
	return c.hooks.Book
}

// BookBorrowClient is a client for the BookBorrow schema.
type BookBorrowClient struct {
	config
}

// NewBookBorrowClient returns a client for the BookBorrow from the given config.
func NewBookBorrowClient(c config) *BookBorrowClient {
	return &BookBorrowClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bookborrow.Hooks(f(g(h())))`.
func (c *BookBorrowClient) Use(hooks ...Hook) {
	c.hooks.BookBorrow = append(c.hooks.BookBorrow, hooks...)
}

// Create returns a create builder for BookBorrow.
func (c *BookBorrowClient) Create() *BookBorrowCreate {
	mutation := newBookBorrowMutation(c.config, OpCreate)
	return &BookBorrowCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for BookBorrow.
func (c *BookBorrowClient) Update() *BookBorrowUpdate {
	mutation := newBookBorrowMutation(c.config, OpUpdate)
	return &BookBorrowUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookBorrowClient) UpdateOne(bb *BookBorrow) *BookBorrowUpdateOne {
	mutation := newBookBorrowMutation(c.config, OpUpdateOne, withBookBorrow(bb))
	return &BookBorrowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookBorrowClient) UpdateOneID(id int) *BookBorrowUpdateOne {
	mutation := newBookBorrowMutation(c.config, OpUpdateOne, withBookBorrowID(id))
	return &BookBorrowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BookBorrow.
func (c *BookBorrowClient) Delete() *BookBorrowDelete {
	mutation := newBookBorrowMutation(c.config, OpDelete)
	return &BookBorrowDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BookBorrowClient) DeleteOne(bb *BookBorrow) *BookBorrowDeleteOne {
	return c.DeleteOneID(bb.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BookBorrowClient) DeleteOneID(id int) *BookBorrowDeleteOne {
	builder := c.Delete().Where(bookborrow.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookBorrowDeleteOne{builder}
}

// Create returns a query builder for BookBorrow.
func (c *BookBorrowClient) Query() *BookBorrowQuery {
	return &BookBorrowQuery{config: c.config}
}

// Get returns a BookBorrow entity by its id.
func (c *BookBorrowClient) Get(ctx context.Context, id int) (*BookBorrow, error) {
	return c.Query().Where(bookborrow.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookBorrowClient) GetX(ctx context.Context, id int) *BookBorrow {
	bb, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return bb
}

// QueryUSER queries the USER edge of a BookBorrow.
func (c *BookBorrowClient) QueryUSER(bb *BookBorrow) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bb.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bookborrow.Table, bookborrow.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookborrow.USERTable, bookborrow.USERColumn),
		)
		fromV = sqlgraph.Neighbors(bb.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBOOK queries the BOOK edge of a BookBorrow.
func (c *BookBorrowClient) QueryBOOK(bb *BookBorrow) *BookQuery {
	query := &BookQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bb.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bookborrow.Table, bookborrow.FieldID, id),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookborrow.BOOKTable, bookborrow.BOOKColumn),
		)
		fromV = sqlgraph.Neighbors(bb.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPURPOSE queries the PURPOSE edge of a BookBorrow.
func (c *BookBorrowClient) QueryPURPOSE(bb *BookBorrow) *PurposeQuery {
	query := &PurposeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bb.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bookborrow.Table, bookborrow.FieldID, id),
			sqlgraph.To(purpose.Table, purpose.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookborrow.PURPOSETable, bookborrow.PURPOSEColumn),
		)
		fromV = sqlgraph.Neighbors(bb.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BookBorrowClient) Hooks() []Hook {
	return c.hooks.BookBorrow
}

// PurposeClient is a client for the Purpose schema.
type PurposeClient struct {
	config
}

// NewPurposeClient returns a client for the Purpose from the given config.
func NewPurposeClient(c config) *PurposeClient {
	return &PurposeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `purpose.Hooks(f(g(h())))`.
func (c *PurposeClient) Use(hooks ...Hook) {
	c.hooks.Purpose = append(c.hooks.Purpose, hooks...)
}

// Create returns a create builder for Purpose.
func (c *PurposeClient) Create() *PurposeCreate {
	mutation := newPurposeMutation(c.config, OpCreate)
	return &PurposeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Purpose.
func (c *PurposeClient) Update() *PurposeUpdate {
	mutation := newPurposeMutation(c.config, OpUpdate)
	return &PurposeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PurposeClient) UpdateOne(pu *Purpose) *PurposeUpdateOne {
	mutation := newPurposeMutation(c.config, OpUpdateOne, withPurpose(pu))
	return &PurposeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PurposeClient) UpdateOneID(id int) *PurposeUpdateOne {
	mutation := newPurposeMutation(c.config, OpUpdateOne, withPurposeID(id))
	return &PurposeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Purpose.
func (c *PurposeClient) Delete() *PurposeDelete {
	mutation := newPurposeMutation(c.config, OpDelete)
	return &PurposeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PurposeClient) DeleteOne(pu *Purpose) *PurposeDeleteOne {
	return c.DeleteOneID(pu.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PurposeClient) DeleteOneID(id int) *PurposeDeleteOne {
	builder := c.Delete().Where(purpose.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PurposeDeleteOne{builder}
}

// Create returns a query builder for Purpose.
func (c *PurposeClient) Query() *PurposeQuery {
	return &PurposeQuery{config: c.config}
}

// Get returns a Purpose entity by its id.
func (c *PurposeClient) Get(ctx context.Context, id int) (*Purpose, error) {
	return c.Query().Where(purpose.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PurposeClient) GetX(ctx context.Context, id int) *Purpose {
	pu, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pu
}

// QueryFor queries the for edge of a Purpose.
func (c *PurposeClient) QueryFor(pu *Purpose) *BookBorrowQuery {
	query := &BookBorrowQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(purpose.Table, purpose.FieldID, id),
			sqlgraph.To(bookborrow.Table, bookborrow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, purpose.ForTable, purpose.ForColumn),
		)
		fromV = sqlgraph.Neighbors(pu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PurposeClient) Hooks() []Hook {
	return c.hooks.Purpose
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Create returns a create builder for Role.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *RoleClient) DeleteOneID(id int) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Create returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{config: c.config}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int) *Role {
	r, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return r
}

// QueryRole queries the role edge of a Role.
func (c *RoleClient) QueryRole(r *Role) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, role.RoleTable, role.RoleColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// StatusClient is a client for the Status schema.
type StatusClient struct {
	config
}

// NewStatusClient returns a client for the Status from the given config.
func NewStatusClient(c config) *StatusClient {
	return &StatusClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `status.Hooks(f(g(h())))`.
func (c *StatusClient) Use(hooks ...Hook) {
	c.hooks.Status = append(c.hooks.Status, hooks...)
}

// Create returns a create builder for Status.
func (c *StatusClient) Create() *StatusCreate {
	mutation := newStatusMutation(c.config, OpCreate)
	return &StatusCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Status.
func (c *StatusClient) Update() *StatusUpdate {
	mutation := newStatusMutation(c.config, OpUpdate)
	return &StatusUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StatusClient) UpdateOne(s *Status) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatus(s))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StatusClient) UpdateOneID(id int) *StatusUpdateOne {
	mutation := newStatusMutation(c.config, OpUpdateOne, withStatusID(id))
	return &StatusUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Status.
func (c *StatusClient) Delete() *StatusDelete {
	mutation := newStatusMutation(c.config, OpDelete)
	return &StatusDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StatusClient) DeleteOne(s *Status) *StatusDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StatusClient) DeleteOneID(id int) *StatusDeleteOne {
	builder := c.Delete().Where(status.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StatusDeleteOne{builder}
}

// Create returns a query builder for Status.
func (c *StatusClient) Query() *StatusQuery {
	return &StatusQuery{config: c.config}
}

// Get returns a Status entity by its id.
func (c *StatusClient) Get(ctx context.Context, id int) (*Status, error) {
	return c.Query().Where(status.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StatusClient) GetX(ctx context.Context, id int) *Status {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryStatus queries the status edge of a Status.
func (c *StatusClient) QueryStatus(s *Status) *BookQuery {
	query := &BookQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(status.Table, status.FieldID, id),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, status.StatusTable, status.StatusColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StatusClient) Hooks() []Hook {
	return c.hooks.Status
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryBorrow queries the Borrow edge of a User.
func (c *UserClient) QueryBorrow(u *User) *BookBorrowQuery {
	query := &BookBorrowQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(bookborrow.Table, bookborrow.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.BorrowTable, user.BorrowColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRolePlay queries the RolePlay edge of a User.
func (c *UserClient) QueryRolePlay(u *User) *RoleQuery {
	query := &RoleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.RolePlayTable, user.RolePlayColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
