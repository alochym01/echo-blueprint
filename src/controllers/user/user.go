package user

import (
	"fmt"
	"helpers"
	"models"
	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	result := fmt.Sprintf("<h1>Show All User with Method %s</h1>", c.Request().Method)
	fmt.Println(c.Path())
	return c.HTML(http.StatusOK, result)
}

func Create(c echo.Context) error {
	result := fmt.Sprintf("<h1>Create User %s with Method %s</h1>", c.Param("id"), c.Request().Method)
	fmt.Println(c.Path())
	return c.HTML(http.StatusOK, result)
}

func Show(c echo.Context) error {
	// result := fmt.Sprintf("<h1>Show User %s with Method %s</h1>", c.Param("id"), c.Request().Method)
	// fmt.Println(c.Path())
	// return c.HTML(http.StatusOK, result)
	user := models.User{}
	// Debug a single operation, show detailed log for this operation
	// helpers.DbGorm.Debug().Where("id = ?", c.Param("id")).First(&user)
	helpers.DbGorm.Debug().First(&user, "id = ?", c.Param("id"))
	return c.JSON(http.StatusOK, user)
}

func Update(c echo.Context) error {
	// result := fmt.Sprintf("<h1>Update User %s with Method %s</h1>", c.Param("id"), c.Request().Method)
	// fmt.Println(c.Path())
	// return c.HTML(http.StatusOK, result)
	// helpers.DbGorm.Debug().Where("id = ?", c.Param("id")).First(&user)
	user := models.User{}
	helpers.DbGorm.Debug().First(&user, "id = ?", c.Param("id"))
	user.Name = "jinzhu 2"
	user.Email = "jinzhu@abc.com"
	helpers.DbGorm.Debug().Save(&user)
	result := helpers.HTTPSuccess{
		Code:    201,
		Message: "Update Success",
	}
	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	result := fmt.Sprintf("<h1>Delete User %s with Method %s</h1>", c.Param("id"), c.Request().Method)
	fmt.Println(c.Path())
	return c.HTML(http.StatusOK, result)
}
