package db

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Test_Query(t *testing.T) {
	//todo test, 1. add temp db data, 2. query 3.
	//where time > now() - 1h
	//stddev

	columnsVal := "stddev(count)"
	dbVal := "chain.block.insert.meter"

	timeStartAgo := "3h"
	// TIME ZONE Z0
	// timeUnitVal := "1m"
	intervals := "10m"
	offsetTime := ",-15s"
	// offsetTime := ",1m"
	// offsetTime := ",-1m"
	// tagsVal1 := "coinBase"
	tagsVal1 := "addr"
	// tagsVal1 := ""
	tagsVal := ""
	if tagsVal1 == "" {
		tagsVal = ""
	} else {
		tagsVal = "," + tagsVal1
	}

	// fillVal := "none" // linear, none, null(default val), previous
	fillVal := "0" // null
	limitVal := 20
	zoneVal := "tz('Asia/Shanghai')"
	// zoneVal := ""
	queryCount := fmt.Sprintf("select %s from \"%s\" where time >= now() - %s group by time(%s%s) %s fill(%s) limit %d %s;",
		columnsVal, dbVal, timeStartAgo, intervals, offsetTime, tagsVal, fillVal, limitVal, zoneVal)
	res, err := Query(queryCount)

	// res, err := Query("select stddev(count) from \"chain.block.insert.meter\" where time >= now() - 1h group by time(5m),addr fill(null) limit 2000 tz('Asia/Shanghai');")

	// res, err := Query("select mean(count) from \"chain.block.insert.meter\" where time >= now() - 1h group by time(1m),addr fill(null) limit 20 tz('Asia/Shanghai');")

	// res, err := Query("select mean(count) from \"chain.block.insert.meter\" where time >= now() - 1h group by time(5m),addr fill(null) limit 20 tz('Asia/Shanghai');")

	// res, err := Query("select mean(count) from \"chain.block.insert.meter\" where time >= now() - 1h group by time(5m),addr fill(0) limit 20 tz('Asia/Shanghai');")

	// res, err := Query("select mean(count) from \"chain.block.insert.meter\" group by addr,time(30s);" +
	// "select mean(count) from \"chain.block.insert.meter\" group by addr,time(60s)") // group occur tags in result
	// res, err := Query("select mean(count) from \"chain.block.insert.meter\" group by addr,time(30s) ") // group occur tags in result

	// res, err := Query("select mean(count) from \"chain.block.insert.meter\" group by addr,time(30s) ") // group occur tags in result
	// res, err := Query("select count, time, addr from \"chain.block.insert.meter\" group by addr,coinBase limit 20") // group occur tags in result
	// res, err := Query("select count, time, addr from \"chain.block.insert.meter\" group by addr,coinBase order by desc limit 20") // group occur tags in result
	// res, err := Query("select count, time, addr from \"chain.block.insert.meter\"  order by desc limit 20") // multip lines should use group by
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Test_Query %v\n", res)
	// two statement will generate two resultSet

	for index, v := range res {
		resultErr := v.Err
		//time addr coinBase count m1 m15 m5 mean name networkId version
		// resultSeries := v.Series
		// resultMessages := v.Messages
		// t.Logf("index: %d, content: {err=%v, series=%v, messages=%v}\n", index, resultErr, resultSeries, resultMessages)
		// fmt.Printf("index: %d, content: {err=%v, series=%v, messages=%v}\n", index, resultErr, resultSeries, resultMessages)
		fmt.Printf("index:%v, errMsg:%v\n", index, resultErr)
		// group by generate the series, mark the diff lines data
		linesCount := len(v.Series)
		fmt.Printf("total lines is %v\n", linesCount)
		dataLines := make([][]map[string]interface{}, linesCount)
		dataFields := make([]string, linesCount)

		//[]interface{}
		for i, val := range v.Series {
			name := val.Name
			columns := val.Columns
			tags := val.Tags
			vals := val.Values
			fmt.Printf("index:%v, name:%v, columns:%v,tags:%v\n", i, name, columns, tags)
			fmt.Printf("index:%v, name:%v, columns:%v,tags:%v\n", i, name, columns, tags[tagsVal1])
			dataFields[i] = tags[tagsVal1]
			//todo
			dataSet1 := make([]map[string]interface{}, 1)

			for _, valj := range vals {
				// fmt.Printf("row %d, val %v\n", j, valj)
				lindData := map[string]interface{}{
					"time": valj[0],
					"tag":  tags[tagsVal1],
					"val":  valj[1],
				}

				dataSet1 = append(dataSet1, lindData)
			}
			dataLines[i] = dataSet1

		}

		// fmt.Printf("dataLines is %v\n", dataLines)

		validLinesData := mergeMap(dataLines[:])
		// fmt.Printf("valid data:\n%v\n", validLinesData)
		jsonByte, err := json.Marshal(validLinesData)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("valid dataFields:\n%v\n", dataFields)
		fmt.Printf("valid data:\n%v\n", string(jsonByte))

	}
}

// all data is exist in the time point, if not exist use 0 or other data fill
func mergeMap(dataSet [][]map[string]interface{}) []map[string]interface{} {
	lineCount := len(dataSet)
	if lineCount == 0 {
		return nil
	}
	fmt.Printf("dataSet is %v\n", dataSet)

	dataSize := len(dataSet[0])
	dataSetsNew := make([]map[string]interface{}, 0)
continueL:
	for i := 0; i < dataSize; i++ {
		innerMap := make(map[string]interface{})
		for j := 0; j < lineCount; j++ {
			data := dataSet[j][i]
			if data == nil {
				continue continueL
			}
			// fmt.Printf("data is %#v\n", data)
			tag := data["tag"].(string)
			innerMap[tag] = data["val"]
		}
		innerMap["time"] = dataSet[0][i]["time"].(string)
		dataSetsNew = append(dataSetsNew, innerMap)
	}
	// fmt.Printf("valid dataSetsNew:\n%v\n", dataSetsNew)
	return dataSetsNew
}
