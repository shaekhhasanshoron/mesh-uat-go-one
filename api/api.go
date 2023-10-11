package api

import (
	"github.com/labstack/echo/v4"
	_type "mesh-uat-go-one/types"
	"mesh-uat-go-one/utils"
	"net/http"
)

type HomeInf interface {
	Index(c echo.Context) error
	ApiIndex(c echo.Context) error
	Health(c echo.Context) error
	HelloApiOne(c echo.Context) error
	HelloApiTwo(c echo.Context) error
	MeshUatTwoApiOne(c echo.Context) error
	MeshUatTwoApiTwo(c echo.Context) error
}

type home struct{}

func MeshApi() HomeInf {
	return &home{}
}

func (h *home) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"appName": "Mesh-UAT-Go-One",
	})
}

func (h *home) ApiIndex(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "This is Mesh-UAT-Go-One Canary Api",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *home) Health(c echo.Context) error {
	return c.String(http.StatusOK, "Mesh-UAT-Go-One I am live!")
}

func (h *home) HelloApiOne(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "Hello From Mesh-UAT-Go-One Canary API One",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *home) HelloApiTwo(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "Hello From Mesh-UAT-Go-One Canary API Two",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *home) MeshUatTwoApiOne(c echo.Context) error {
	resp, err := utils.MeshUatTwoAppApiMeshOne()
	if err != nil {
		return c.JSON(http.StatusBadRequest, _type.Response().Error("Error Occurred Mesh-UAT-Go-Two Api One: "+err.Error()))
	}

	response := _type.ResponseDto{
		Message: "Mesh-UAT-Go-One API --TO-- Mesh-UAT-Go-Two Api One",
		Data:    resp,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *home) MeshUatTwoApiTwo(c echo.Context) error {
	resp, err := utils.MeshUatTwoAppApiMeshTwo()
	if err != nil {
		return c.JSON(http.StatusBadRequest, _type.Response().Error("Error Occurred Mesh-UAT-Go-Two Api Two: "+err.Error()))
	}

	response := _type.ResponseDto{
		Message: "Mesh-UAT-Go-One API --TO-- Mesh-UAT-Go-Two Api Two",
		Data:    resp,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}
