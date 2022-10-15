/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	locingester "github.com/tavvfiq/driver-loc-iot/ingester"
	"golang.org/x/net/context"
)

// ingesterCmd represents the ingester command
var ingesterCmd = &cobra.Command{
	Use:   "ingester",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		numOfIngesters := 1
		if args[0] != "" {
			val, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("number of drivers must be a number of integers")
				return
			}
			numOfIngesters = val
		}
		fmt.Println("running ingester")
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		ctx, cancel := context.WithCancel(context.Background())
		done := make(chan bool)
		go func() {
			<-sig
			fmt.Println("cleaning up ingester service")
			cancel()
			done <- true
		}()
		for i := 0; i < numOfIngesters; i++ {
			id := fmt.Sprintf("00%d", i+1)
			d := locingester.NewIngester(id)
			go d.Run(ctx)
		}
		<-done
		time.Sleep(2 * time.Second)
	},
}

func init() {
	rootCmd.AddCommand(ingesterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ingesterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ingesterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
