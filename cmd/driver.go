/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/tavvfiq/driver-loc-iot/driver"
)

// driverCmd represents the driver command
var driverCmd = &cobra.Command{
	Use:   "driver",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		numOfDrivers := 1
		if args[0] != "" {
			val, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("number of drivers must be a number of integers")
				return
			}
			numOfDrivers = val
		}
		done := make(chan bool)
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			defer cancel()
			<-sig
			fmt.Println("cleaning up driver services...")
			done <- true
		}()
		fmt.Println("running simulate driver")
		for i := 0; i < numOfDrivers; i++ {
			id := fmt.Sprintf("00%d", i+1)
			d := driver.NewDriver(id)
			go d.Run(ctx)
		}
		<-done
		time.Sleep(2 * time.Second)
	},
}

func init() {
	rootCmd.AddCommand(driverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// driverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// driverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
