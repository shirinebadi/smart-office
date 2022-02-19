package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/shirinebadi/smart-office/internal/request"
)

func (lsh *LocalServerHandler) Login(c echo.Context) error {
	req := new(request.Authentication)

	if err := c.Bind(req); err != nil {
		log.Error("Error in Login: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	user, err := lsh.AuthenticationI.UserLogin(req.Username, req.Password)
	if err != nil {
		log.Error("Error in Login: %s", err.Error())
		return c.NoContent(http.StatusUnauthorized)
	}

	token, _ := lsh.Token.GenerateJWT(user.Id)

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (lsh *LocalServerHandler) LightValue(c echo.Context) error {

	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	id := c.Param("id")
	light := c.QueryParam("light")

	url := "http://127.0.0.1" + lsh.Config.CentralServer.Address + "/api/office/lightvalue"
	fmt.Print(url)
	intID, _ := strconv.Atoi(id)
	intLight, _ := strconv.Atoi(light)
	values := map[string]int{"id": intID, "light": intLight}
	jsonValue, _ := json.Marshal(values)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	if resp.StatusCode == 200 {
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusBadRequest)
}
