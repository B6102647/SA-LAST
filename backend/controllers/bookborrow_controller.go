package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/B6102647/app/ent"
	"github.com/B6102647/app/ent/book"
	"github.com/B6102647/app/ent/bookborrow"
	"github.com/B6102647/app/ent/purpose"
	"github.com/B6102647/app/ent/user"
	"github.com/gin-gonic/gin"
)

// BookBorrowController defines the struct for the bookborrow controller
type BookBorrowController struct {
	client *ent.Client
	router gin.IRouter
}

//BookBorrow struct
type BookBorrow struct {
	User    int
	Book    int
	Purpose int
	Added   string
}

// CreateBookBorrow handles POST requests for adding bookborrow entities
// @Summary Create bookborrow
// @Description Create bookborrow
// @ID create-bookborrow
// @Accept   json
// @Produce  json
// @Param bookborrow body BookBorrow true "BookBorrow entity"
// @Success 200 {object} ent.BookBorrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows [post]
func (ctl *BookBorrowController) CreateBookBorrow(c *gin.Context) {
	obj := BookBorrow{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bookBorrow binding failed",
		})
		return
	}
	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	bk, err := ctl.client.Book.
		Query().
		Where(book.IDEQ(int(obj.Book))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Book not found",
		})
		return
	}
	pp, err := ctl.client.Purpose.
		Query().
		Where(purpose.IDEQ(int(obj.Purpose))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "purpose not found",
		})
		return
	}
	times, err := time.Parse(time.RFC3339, obj.Added)
	bb, err := ctl.client.BookBorrow.
		Create().
		SetBOOK(bk).
		SetUSER(u).
		SetPURPOSE(pp).
		SetADDEDTIME(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bb)
}

// GetBookBorrow handles GET requests to retrieve a bookborrow entity
// @Summary Get a bookborrow entity by ID
// @Description get bookborrow by ID
// @ID get-bookborrow
// @Produce  json
// @Param id path int true "BookBorrow ID"
// @Success 200 {object} ent.BookBorrow
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows/{id} [get]
func (ctl *BookBorrowController) GetBookBorrow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	bb, err := ctl.client.BookBorrow.
		Query().
		Where(bookborrow.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bb)
}

// ListBookBorrow handles request to get a list of bookborrow entities
// @Summary List bookborrow entities
// @Description list bookborrow entities
// @ID list-bookborrow
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.BookBorrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows [get]
func (ctl *BookBorrowController) ListBookBorrow(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	bookborrows, err := ctl.client.BookBorrow.
		Query().
		WithBOOK().
		WithUSER().
		WithPURPOSE().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bookborrows)
}

// DeleteBookBorrow handles DELETE requests to delete a bookborrow entity
// @Summary Delete a bookborrow entity by ID
// @Description get bookborrow by ID
// @ID delete-bookborrow
// @Produce  json
// @Param id path int true "BookBorrow ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows/{id} [delete]
func (ctl *BookBorrowController) DeleteBookBorrow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.BookBorrow.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateBookBorrow handles PUT requests to update a bookBorrow entity
// @Summary Update a bookborrow entity by ID
// @Description update bookborrow by ID
// @ID update-bookborrow
// @Accept   json
// @Produce  json
// @Param id path int true "BookBorrow ID"
// @Param bookborrow body ent.BookBorrow true "BookBorrow entity"
// @Success 200 {object} ent.BookBorrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows/{id} [put]
func (ctl *BookBorrowController) UpdateBookBorrow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.BookBorrow{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "BookBorrow binding failed",
		})
		return
	}
	obj.ID = int(id)
	bb, err := ctl.client.BookBorrow.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, bb)
}

// NewBookBorrowController creates and registers handles for the bookborrow controller
func NewBookBorrowController(router gin.IRouter, client *ent.Client) *BookBorrowController {
	bb := &BookBorrowController{
		client: client,
		router: router,
	}
	bb.register()
	return bb
}

// InitBookBorrowController registers routes to the main engine
func (ctl *BookBorrowController) register() {
	bookborrows := ctl.router.Group("/bookborrows")

	bookborrows.GET("", ctl.ListBookBorrow)

	// CRUD
	bookborrows.POST("", ctl.CreateBookBorrow)
	bookborrows.GET(":id", ctl.GetBookBorrow)
	bookborrows.PUT(":id", ctl.UpdateBookBorrow)
	bookborrows.DELETE(":id", ctl.DeleteBookBorrow)
}
