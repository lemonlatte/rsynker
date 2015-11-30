package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var cmdServer *cobra.Command
var port string

func runServer(cmd *cobra.Command, args []string) {
	fmt.Println("Start a rsync server at port:", port)

	rsyncCmd := exec.Command("rsync",
		"--no-detach", "--daemon",
		"--port", port,
		"--config", configPath)

	rsyncCmd.Stdout = os.Stdout
	rsyncCmd.Stderr = os.Stderr

	if err := rsyncCmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := rsyncCmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cmdServer = &cobra.Command{
		Use:   "server",
		Short: "Run a rsync server",
		Long:  "Run a rsync server",
		Run:   runServer,
	}

	cmdServer.Flags().StringVarP(&port, "port", "p", "873", "port of rsync server")
}
