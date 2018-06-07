/**
*
*
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/db"
	"github.com/seeleteam/dashboard-api/db/query/origin"
	"github.com/seeleteam/dashboard-api/db/query/param"
)

// SelectBySQLs get data from influxdb by multi influxdb sql
func SelectBySQLs() gin.HandlerFunc {
	return func(c *gin.Context) {
		sqls := c.QueryArray(common.RequestSQL)

		if sqls == nil || len(sqls) == 0 {
			errInfo := fmt.Sprintf("param field sql(array) required!")
			log.Error(errInfo)
			responseData := common.NewResponseData(400, errors.New(errInfo), nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		// be separated by semicolon
		var sqlStr bytes.Buffer
		for _, sql := range sqls {
			if sql != "" {
				if strings.HasSuffix(sql, ";") {
					sqlStr.WriteString(sql)
				} else {
					sqlStr.WriteString(sql + ";")
				}
			}
		}
		if sqlStr.String() == "" {
			errInfo := fmt.Sprintf("param field sqls content error")
			log.Error(errInfo)
			responseData := common.NewResponseData(500, errors.New(errInfo), nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		query := origin.New(sqlStr.String())
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("SelectBySQLs, err:\n%v\n")
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// SelectWithParams select with params(generate sql)
func SelectWithParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := c.QueryArray(common.RequestFields)
		tableName := c.Query(common.RequestMeasurement)
		if tableName == "" {
			errInfo := fmt.Sprintf("param %s blank", common.RequestMeasurement)
			log.Error(errInfo)
			responseData := common.NewResponseData(404, errInfo, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		whereExpressions := c.QueryArray(common.RequestWhereExpressions)
		startTime := c.Query(common.RequestStartTime)
		endTime := c.Query(common.RequestEndTime)
		timeSince := c.Query(common.RequestTimeSince)

		intervals := c.Query(common.RequestIntervals)
		intervalsOffset := c.Query(common.RequestIntervalsOffset)
		tags := c.QueryArray(common.RequestTags)

		fillOption := c.Query(common.RequestFillOption)
		orderBy := c.Query(common.RequestOrderBy)

		limit := 0
		limitVal := c.Query(common.RequestLimit)
		if limitVal != "" {
			limit1, err := strconv.ParseInt(limitVal, 10, 10)
			if err != nil {
				log.Error(err)
				responseData := common.NewResponseData(404, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			limit = int(limit1)
		}

		offset := 0
		offsetVal := c.Query(common.RequestOffset)
		if offsetVal != "" {
			offset1, err := strconv.ParseInt(limitVal, 10, 10)
			if err != nil {
				log.Error(err)
				responseData := common.NewResponseData(404, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			offset = int(offset1)
		}

		slimit := 0
		slimitVal := c.Query(common.RequestSLimit)
		if slimitVal != "" {
			slimit1, err := strconv.ParseInt(slimitVal, 10, 10)
			if err != nil {
				log.Error(err)
				responseData := common.NewResponseData(404, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			slimit = int(slimit1)
		}

		soffset := 0
		soffsetVal := c.Query(common.RequestSOffset)
		if soffsetVal != "" {
			soffset1, err := strconv.ParseInt(slimitVal, 10, 10)
			if err != nil {
				log.Error(err)
				responseData := common.NewResponseData(404, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			soffset = int(soffset1)
		}

		timeZone := c.Query(common.RequestTimeZone)

		condition := &db.Condition{
			Fields:           fields,
			Measurement:      tableName,
			WhereExpressions: whereExpressions,
			StartTime:        startTime,
			EndTime:          endTime,
			TimeSince:        timeSince,
			Intervals:        intervals,
			IntervalsOffset:  intervalsOffset,
			Tags:             tags,
			FillOption:       fillOption,
			OrderBy:          orderBy,
			Limit:            limit,
			Offset:           offset,
			SLimit:           slimit,
			SOffset:          soffset,
			TimeZone:         timeZone,
		}

		paramQuery, err := param.New(condition)
		if err != nil {
			log.Error("%v", err)
			responseData := common.NewResponseData(500, err, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		log.Debug("stmt: %v\n", paramQuery.Stmt)

		res, err := paramQuery.Query()
		if err != nil {
			log.Error("%v", err)
			responseData := common.NewResponseData(500, err, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, err, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// SelectNodeInfo select node info with params(generate sql)
func SelectNodeInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		field := c.Query("field")
		if field == "" {
			field = "last(value)"
		}
		tableName := c.Query("measurement")
		if tableName == "" {
			tableName = "runtime.memory.alloc.gauge"
		}
		group := c.Query("group")
		if group == "" {
			group = "shardid,coinbase,networkid,nodename"
		}

		nodeInfoQuery := &param.Query{
			Stmt: fmt.Sprintf("select %s from \"%s\" group by %s fill(null)", field, tableName, group),
		}

		log.Debug("stmt: %v\n", nodeInfoQuery)

		res, err := nodeInfoQuery.Query()
		if err != nil {
			log.Error("%v", err)
			responseData := common.NewResponseData(500, err, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}

		var nodeInfo []map[string]string

		for _, result := range res {
			nodeInfo = make([]map[string]string, 0)
			for _, series := range result.Series {
				nodeInfo = append(nodeInfo, series.Tags)
			}
		}

		responseData := common.NewResponseData(200, err, &nodeInfo, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}
