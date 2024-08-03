package organization

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrganizationController struct {
	Service *OrganizationService
}

func (c *OrganizationController) CreateOrganization(ctx *gin.Context) {
	var org Organization
	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := c.Service.CreateOrganization(org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

func (c *OrganizationController) GetOrganization(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	org, err := c.Service.GetOrganization(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}
	ctx.JSON(http.StatusOK, org)
}

func (c *OrganizationController) UpdateOrganization(ctx *gin.Context) {
	var org Organization
	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org.ID, _ = strconv.Atoi(ctx.Param("id"))
	err := c.Service.UpdateOrganization(org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Organization updated"})
}

func (c *OrganizationController) DeleteOrganization(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.Service.DeleteOrganization(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Organization deleted"})
}

func (c *OrganizationController) GetAllOrganizations(ctx *gin.Context) {
	orgs, err := c.Service.GetAllOrganizations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orgs)
}
