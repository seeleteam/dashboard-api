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
	_, err := GetConfigFromFile(filepath)
	if err != nil {
		t.Errorf("GetConfigFromFile %s %v", filepath, err)
	}
}

func Test_GetConfigFromFileBad(t *testing.T) {
	filepath := "./config/app-bad.json"
	_, err := GetConfigFromFile(filepath)
	if err == nil {
		t.Errorf("GetConfigFromFile %s %v", filepath, err)
	}
}

func Test_LoadConfigFromFile(t *testing.T) {
	filepath := "./config/app.json"
	_, err := LoadConfigFromFile(filepath)
	if err != nil {
		t.Errorf("LoadConfigFromFile %s %v", filepath, err)
	}
}

func Test_LoadConfigFromFileBad(t *testing.T) {
	filepath := "./config/app-bad.json"
	_, err := LoadConfigFromFile(filepath)
	if err == nil {
		t.Errorf("LoadConfigFromFile %s %v", filepath, err)
	}
}
