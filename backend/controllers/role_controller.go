package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B6102647/app/ent"
	"github.com/B6102647/app/ent/role"
	"github.com/gin-gonic/gin"
)

// RoleController defines the struct for the Role controller
type RoleController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateRole handles POST requests for adding Role entities
// @Summary Create Role
// @Description Create Role
// @ID create-Role
// @Accept   json
// @Produce  json
// @Param Role body ent.Role true "Role entity"
// @Success 200 {object} ent.Role
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Role [post]
func (ctl *RoleController) CreateRole(c *gin.Context) {
	obj := ent.Role{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Role binding failed",
		})
		return
	}

	ro, err := ctl.client.Role.
		Create().
		SetROLEID(obj.ROLEID).
		SetROLENAME(obj.ROLENAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ro)
}

// GetRole handles GET requests to retrieve a Role entity
// @Summary Get a Role entity by ID
// @Description get Role by ID
// @ID get-Role
// @Produce  json
// @Param id path int true "Role ID"
// @Success 200 {object} ent.Role
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Roles/{id} [get]
func (ctl *RoleController) GetRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ro, err := ctl.client.Role.
		Query().
		Where(role.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ro)
}

// ListRole handles request to get a list of Role entities
// @Summary List Role entities
// @Description list Role entities
// @ID list-Role
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Role
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Role [get]
func (ctl *RoleController) ListRole(c *gin.Context) {
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

	role, err := ctl.client.Role.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, role)
}

// DeleteRole handles DELETE requests to delete a Role entity
// @Summary Delete a Role entity by ID
// @Description get Role by ID
// @ID delete-PRole
// @Produce  json
// @Param id path int true "Role ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Role/{id} [delete]
func (ctl *RoleController) DeleteRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Role.
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

// UpdateRole handles PUT requests to update a Role entity
// @Summary Update a Role entity by ID
// @Description update Role by ID
// @ID update-Role
// @Accept   json
// @Produce  json
// @Param id path int true "Role ID"
// @Param Role body ent.Role true "Role entity"
// @Success 200 {object} ent.Role
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Role/{id} [put]
func (ctl *RoleController) UpdateRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Role{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Role binding failed",
		})
		return
	}
	obj.ID = int(id)
	ro, err := ctl.client.Role.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ro)
}

// NewRoleController creates and registers handles for the Role controller
func NewRoleController(router gin.IRouter, client *ent.Client) *RoleController {
	pp := &RoleController{
		client: client,
		router: router,
	}
	pp.register()
	return pp
}

// InitRoleController registers routes to the main engine
func (ctl *RoleController) register() {
	Role := ctl.router.Group("/Role")

	Role.GET("", ctl.ListRole)

	// CRUD
	Role.POST("", ctl.CreateRole)
	Role.GET(":id", ctl.GetRole)
	Role.PUT(":id", ctl.UpdateRole)
	Role.DELETE(":id", ctl.DeleteRole)
}
