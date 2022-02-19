package local_server

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/shirinebadi/smart-office/internal/config"
	"github.com/shirinebadi/smart-office/internal/handler"
	"github.com/shirinebadi/smart-office/pkg/data/db"
	"github.com/spf13/cobra"
)

func main(config config.Config) {
	myDB, err := db.Init()
	if err != nil {
		log.Fatal("Failed to setup db: ", err.Error())
	}

	userI := &db.Mydb{DB: myDB}
	activityI := &db.Mydb{DB: myDB}
	authenticationI := &db.Mydb{DB: myDB}

	token := handler.Token{Cfg: config}

	localServerHandler := &handler.LocalServerHandler{AuthenticationI: authenticationI, Token: token, UserI: userI, ActivityI: activityI, Config: config}
	//centralServerHandler := &handler.CentralServerHandler{UserI: userI, ActivityI: activityI, AuthenticationI: authenticationI, Token: token}

	e := echo.New()

	e.POST("/api/admin/login", localServerHandler.AdminLogin)
	e.POST("/api/admin/register", localServerHandler.AdminRegister)
	e.POST("/api/admin/user/register", localServerHandler.UserRegister)
	e.GET("/api/admin/activities", localServerHandler.UserActivity)
	e.POST("/api/user/login", localServerHandler.Login)
	e.POST("/api/admin/setlights", localServerHandler.SetLights)
	e.GET("/api/user/:id", localServerHandler.LightValue)

	address := config.LocalServer.Address

	fmt.Print("gi", address)

	if err := e.Start(address); err != nil {
		log.Fatal(err)
	}
}

func Register(root *cobra.Command, cfg config.Config) {
	runServer := &cobra.Command{
		Use:   "local-server",
		Short: "http server for smart office",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runServer)
}
