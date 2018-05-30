/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	"github.com/seeleteam/dashboard-api/api"
)

var (
	apiConfigFile *string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the api",
	Long: `usage example:
		api.exe start -c cmd\app.json
		start the api.`,

	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		apiConfig, err := LoadConfigFromFile(*apiConfigFile)
		if err != nil {
			fmt.Printf("reading the config file failed: %s\n", err.Error())
			return
		}

		apiServer, err := api.New(apiConfig)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if err := apiServer.Start(); err != nil {
			fmt.Println(err.Error())
			return
		}

		wg.Add(1)
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	apiConfigFile = startCmd.Flags().StringP("config", "c", "", "api config file (required)")
	startCmd.MarkFlagRequired("config")
}
