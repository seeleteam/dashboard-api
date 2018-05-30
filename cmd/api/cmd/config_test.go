/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package cmd

import (
	"fmt"
	"testing"
)

func Test_GetConfigFromFile(t *testing.T) {
	filepath := "./app.json"
	config, err := GetConfigFromFile(filepath)
	if err != nil {
		fmt.Printf("app.json %v", config)
	}
	fmt.Printf("%#v", config)
}

func Test_LoadConfigFromFile(t *testing.T) {
	filepath := "./app.json"
	config, err := LoadConfigFromFile(filepath)
	if err != nil {
		fmt.Printf("app.json %v", config)
	}
	fmt.Printf("%#v", *config)
}
