package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/shirinebadi/smart-office/internal/config"
	"github.com/shirinebadi/smart-office/internal/model"
	"github.com/shirinebadi/smart-office/internal/request"
)

type LocalServerHandler struct {
	UserI           model.UserInterface
	OfficeI         model.OfficeInterface
	AuthenticationI model.AuthenticationInterface
	ActivityI       model.ActivityInterface
	Token           Token
	Config          config.Config
}

type CentralServerHandler struct {
	OfficeI         model.OfficeInterface
	UserI           model.UserInterface
	ActivityI       model.ActivityInterface
	AuthenticationI model.AuthenticationInterface
	Token           Token
}

func (lsh *LocalServerHandler) AdminRegister(c echo.Context) error {
	req := new(request.Authentication)
	fmt.Print(req.Username)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	AdminToDB := model.NewAdmin(req.Username, req.Password)
	if err := lsh.AuthenticationI.AdminRegister(AdminToDB); err != nil {
		log.Error(err)
	}

	token, _ := lsh.Token.GenerateJWT(*&AdminToDB.Username)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (lsh *LocalServerHandler) AdminLogin(c echo.Context) error {
	req := new(request.Authentication)

	if err := c.Bind(req); err != nil {
		log.Error("Error in Login: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	fmt.Print(req.Username)

	admin, err := lsh.AuthenticationI.AdminLogin(req.Username, req.Password)
	if err != nil {
		log.Error("Error in Login: %s", err.Error())
		return c.NoContent(http.StatusUnauthorized)
	}

	fmt.Println(admin)

	token, _ := lsh.Token.GenerateJWT(admin.Username)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
