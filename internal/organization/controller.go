package organization

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	auth "github.com/sittaendah/aegis/internal/auth"
)

type OrganizationController struct {
	Service *OrganizationService
}

func (c *OrganizationController) CreateOrganization(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	claims, err := auth.ParseToken(tokenStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var org Organization
	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org.CreatedBy = claims.Username

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
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	var org Organization
	claims, err := auth.ParseToken(tokenStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	org.UpdatedBy = claims.Username

	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org.ID, _ = strconv.Atoi(ctx.Param("id"))

	err = c.Service.UpdateOrganization(org, claims.Username)
	if err != nil {
		if err.Error() == "insufficient permissions" || err.Error() == "user not authorized" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Organization updated"})
}

func (c *OrganizationController) DeleteOrganization(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")
	if tokenStr == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	claims, err := auth.ParseToken(tokenStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	err = c.Service.DeleteOrganization(id, claims.Username)
	if err != nil {
		if err.Error() == "insufficient permissions" || err.Error() == "user not authorized" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
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
