package middleware

import (
	"api/domain"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/qor/roles"
)

func ResourcePermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, ok := c.MustGet("current_user").(domain.User)
		if !ok {
			df := 0
			currentUser = domain.User{Role: &df}
		}
		role := changeRole(currentUser.Role)
		permission := ResourcePermission()
		var err error
		switch c.Request.Method {
		case "GET":
			if !permission.HasPermission(roles.Read, role) {
				err = PermissionError()
			}
		case "POST":
			if !permission.HasPermission(roles.Create, role) {
				err = PermissionError()
			}
		case "PATCH", "PUT":
			if !permission.HasPermission(roles.Update, role) {
				err = PermissionError()
			}
		case "DELETE":
			if !permission.HasPermission(roles.Delete, role) {
				err = PermissionError()
			}
		}

		if err != nil {
			c.JSON(401, Error{Message: err.Error()})
			c.Abort()
			return
		}
	}
}

func UserPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, ok := c.MustGet("current_user").(domain.User)
		if !ok {
			c.JSON(401, Error{Message: "Authorization error: You must login."})
			c.Abort()
		}
		role := changeRole(currentUser.Role)
		permission := UserPermission()
		var err error
		switch c.Request.Method {
		case "GET":
			if !permission.HasPermission(roles.Read, role) {
				err = PermissionError()
			}
		case "POST":
			if !permission.HasPermission(roles.Create, role) {
				err = PermissionError()
			}
		case "PATCH", "PUT":
			if !permission.HasPermission(roles.Update, role) {
				err = PermissionError()
			}
		case "DELETE":
			if !permission.HasPermission(roles.Delete, role) {
				err = PermissionError()
			}
		}

		if err != nil {
			c.JSON(401, Error{Message: err.Error()})
			c.Abort()
			return
		}
	}
}

func changeRole(role *int) string {
	switch *role {
	case 1:
		return "writer"
	case 2:
		return "admin"
	default:
		return "user"
	}
}

func PermissionError() error {
	return errors.New("Authorization error: Request is not permit.")
}

func ResourcePermission() *roles.Permission {
	permission := roles.Allow(roles.Read, "user")
	permission.Allow(roles.CRUD, "writer")
	permission.Allow(roles.CRUD, "admin")
	return permission
}

func UserPermission() *roles.Permission {
	permission := roles.Deny(roles.CRUD, "user")
	permission.Allow(roles.Update, "writer")
	permission.Allow(roles.Read, "writer")
	permission.Allow(roles.CRUD, "admin")
	return permission
}
