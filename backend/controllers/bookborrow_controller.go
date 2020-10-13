package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B6102647/app/ent"
	"github.com/B6102647/app/ent/bookborrow"
	"github.com/gin-gonic/gin"
)

// BookBorrowController defines the struct for the BookBorrow controller
type BookBorrowController struct {
	client *ent.Client
	router gin.IRouter
}

type BookBorrow struct {
	Owner   int
	Book    int
	Perpose int
	Added   string
}

// CreateBookBorrow handles POST requests for adding BookBorrow entities
// @Summary Create BookBorrow
// @Description Create BookBorrow
// @ID create-BookBorrow
// @Accept   json
// @Produce  json
// @Param BookBorrow body ent.BookBorrow true "BookBorrow entity"
// @Success 200 {object} ent.BookBorrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /BookBorrow [post]
func (ctl *BookBorrowController) CreateBookBorrow(c *gin.Context) {
	obj := ent.BookBorrow{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "BookBorrow binding failed",
		})
		return
	}

	bb, err := ctl.client.BookBorrow.
		Create().
		SetADDEDTIME(obj.ADDEDTIME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bb)
}

// GetBookBorrow handles GET requests to retrieve a BookBorrow entity
// @Summary Get a BookBorrow entity by ID
// @Description get BookBorrow by ID
// @ID get-BookBorrow
// @Produce  json
// @Param id path int true "BookBorrow ID"
// @Success 200 {object} ent.BookBorrow
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /BookBorrow/{id} [get]
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

// ListBookBorrow handles request to get a list of BookBorrow entities
// @Summary List BookBorrow entities
// @Description list BookBorrow entities
// @ID list-BookBorrow
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.BookBorrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /BookBorrow [get]
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

	bookborrow, err := ctl.client.BookBorrow.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bookborrow)
}

// DeleteBookBorrow handles DELETE requests to delete a BookBorrow entity
// @Summary Delete a BookBorrow entity by ID
// @Description get BookBorrow by ID
// @ID delete-BookBorrow
// @Produce  json
// @Param id path int true "BookBorrow ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /BookBorrow/{id} [delete]
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

// UpdateBookBorrow handles PUT requests to update a BookBorrow entity
// @Summary Update a BookBorrow entity by ID
// @Description update BookBorrow by ID
// @ID update-BookBorrow
// @Accept   json
// @Produce  json
// @Param id path int true "BookBorrow ID"
// @Param BookBorrow body ent.BookBorrow true "BookBorrow entity"
// @Success 200 {object} ent.BookBorrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /BookBorrow/{id} [put]
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

// NewBookBorrowController creates and registers handles for the BookBorrow controller
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
	BookBorrow := ctl.router.Group("/bookborrow")

	BookBorrow.GET("", ctl.ListBookBorrow)

	// CRUD
	BookBorrow.POST("", ctl.CreateBookBorrow)
	BookBorrow.GET(":id", ctl.GetBookBorrow)
	BookBorrow.PUT(":id", ctl.UpdateBookBorrow)
	BookBorrow.DELETE(":id", ctl.DeleteBookBorrow)
}
