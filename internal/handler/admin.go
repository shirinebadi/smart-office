package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/shirinebadi/smart-office/internal/model"
	"github.com/shirinebadi/smart-office/internal/request"
)

func (lsh *LocalServerHandler) UserRegister(c echo.Context) error {
	req := new(request.User)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register User: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	userToDB := &model.LocalUser{Room: req.Room, Card: req.Card}
	central := &model.CentralUser{Password: req.Password, Room: req.Room}

	if err := lsh.UserI.UserRegister(userToDB); err != nil {
		log.Error(err)
	}
	lsh.UserI.UserCentralRegister(central)

	fmt.Print(userToDB)

	return c.NoContent(http.StatusOK)
}

func (lsh *LocalServerHandler) UserActivity(c echo.Context) error {
	req := new(request.User)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register User: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	token := c.Request().Header.Get("Token")
	if token == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	keyVal := make(map[int]string)

	err, activities := lsh.ActivityI.GetAllActivites()

	if err != nil {
		log.Error("Error in Getting Activities", err)
	}

	for _, activity := range activities {
		keyVal[activity.Id] = "Office: " + "8" + " Type: " + activity.Type + " DateTime: " + activity.DateTime.String()
	}

	return c.JSON(http.StatusOK, keyVal)
}

func (lsh *LocalServerHandler) SetLights(c echo.Context) error {
	req := new(request.Lights)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Set Lights: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	url := "http://127.0.0.1" + lsh.Config.CentralServer.Address + "/api/office/lights"
	fmt.Print(url)
	values := map[string]int{"lightsontime": req.LightsOnTime, "lightsofftime": req.LightsOffTime}
	jsonValue, _ := json.Marshal(values)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	fmt.Print(resp.StatusCode)
	if resp.StatusCode == 200 {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusFailedDependency)
}

func (csl *CentralServerHandler) HandleValue(c echo.Context) error {
	req := new(request.Value)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register User: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}
	user := &model.CentralUser{Id: req.ID, Light: req.Light}

	err := csl.UserI.SetUserLight(user)

	if err == nil {
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusBadRequest)

}

func (csl *CentralServerHandler) HandleLights(c echo.Context) error {
	req := new(request.Lights)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register User: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	office, _ := csl.OfficeI.GetOffice(req.ID)
	office.LightsOffTime = req.LightsOffTime
	office.LightsOnTime = req.LightsOnTime

	csl.OfficeI.UpdateLightsTime(office)

	return c.NoContent(http.StatusOK)
}

func (csl *CentralServerHandler) HandleLocalRequest(c echo.Context) error {
	req := new(request.Activity)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register User: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	activity := &model.Activity{
		Id:       req.ID,
		DateTime: time.Now(),
		Type:     "Enter/Exit",
	}

	err := csl.ActivityI.SetActivity(activity)
	if err != nil {
		log.Error("Error in Updating Activity ", err)
	}

	light, err := csl.UserI.GetUserLight(activity.Id)
	if err != nil {
		log.Error("Error in Getting Light ", err)
	}
	fmt.Println("light: ", light)
	return c.JSON(http.StatusOK, map[string]int{
		"Light": light,
	})
}

func (csl *CentralServerHandler) RegisterOffice(c echo.Context) error {
	req := new(request.Office)

	if err := c.Bind(req); err != nil {
		log.Info("Error in Register User: %s", err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	officeToDB := &model.Office{Id: req.ID}
	err := csl.OfficeI.RegisterOffice(officeToDB)
	if err != nil {
		log.Error("Failed to Register Office", err)
	}

	return c.NoContent(http.StatusOK)
}
