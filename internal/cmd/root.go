package cmd

import (
	"github.com/shirinebadi/smart-office/internal/cmd/central_server"
	"github.com/shirinebadi/smart-office/internal/cmd/local_server"
	"github.com/shirinebadi/smart-office/internal/cmd/mqtt_worker"
	"github.com/shirinebadi/smart-office/internal/config"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	var root = &cobra.Command{
		Use: "smart-office",
	}
	cfg := config.Init()

	local_server.Register(root, cfg)
	central_server.Register(root, cfg)
	mqtt_worker.Register(root, cfg)

	return root
}
