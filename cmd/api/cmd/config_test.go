/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package cmd

import (
	"testing"
)

func Test_GetConfigFromFile(t *testing.T) {
	filepath := "./config/app.json"
	config, err := GetConfigFromFile(filepath)
	if err != nil {
		t.Fatalf("GetConfigFromFile %s %v", filepath, err)
	}
}

func Test_GetConfigFromFileBad(t *testing.T) {
	filepath := "./config/app-bad.json"
	config, err := GetConfigFromFile(filepath)
	if err == nil {
		t.Fatalf("GetConfigFromFile %s %v", filepath, err)
	}
}

func Test_LoadConfigFromFile(t *testing.T) {
	filepath := "./config/app.json"
	config, err := LoadConfigFromFile(filepath)
	if err != nil {
		t.Fatalf("LoadConfigFromFile %s %v", filepath, err)
	}
}

func Test_LoadConfigFromFileBad(t *testing.T) {
	filepath := "./config/app-bad.json"
	config, err := LoadConfigFromFile(filepath)
	if err == nil {
		t.Fatalf("LoadConfigFromFile %s %v", filepath, err)
	}
}
