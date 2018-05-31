/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package cmd

import (
	"testing"
)

func Test_GetConfigFromFile(t *testing.T) {
	filepath := "./app.json"
	config, err := GetConfigFromFile(filepath)
	if err != nil {
		t.Errorf("app.json %v", config)
	}
}

func Test_LoadConfigFromFile(t *testing.T) {
	filepath := "./app.json"
	config, err := LoadConfigFromFile(filepath)
	if err != nil {
		t.Errorf("app.json %v", config)
	}
}
