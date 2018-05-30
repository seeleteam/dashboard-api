/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package db

import (
	"fmt"
	"testing"
)

func Test_Query(t *testing.T) {
	showDB := "show databases"
	res, err := Query(showDB)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v:\n%#v\n", showDB, res)
}
