package central_server

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
	officeI := &db.Mydb{DB: myDB}

	token := handler.Token{Cfg: config}

	centralServerHandler := &handler.CentralServerHandler{UserI: userI, ActivityI: activityI, AuthenticationI: authenticationI, Token: token, OfficeI: officeI}

	e := echo.New()

	e.POST("/api/office/register", centralServerHandler.RegisterOffice)
	e.POST("/api/office/lights", centralServerHandler.HandleLights)
	e.POST("/api/office/lightvalue", centralServerHandler.HandleValue)
	e.POST("/api/office/localserver", centralServerHandler.HandleLocalRequest)

	address := config.CentralServer.Address

	fmt.Print("gi", address)

	if err := e.Start(address); err != nil {
		log.Fatal(err)
	}
}

func Register(root *cobra.Command, cfg config.Config) {
	runServer := &cobra.Command{
		Use:   "central-server",
		Short: "http server for smart office",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(runServer)
}
