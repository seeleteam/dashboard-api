/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package handlers

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/seeleteam/dashboard-api/common"
	"github.com/seeleteam/dashboard-api/db/query/origin"
)

// ShowDatabases show database
func ShowDatabases() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := origin.New(common.ShowDatabases)
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("ShowDatabases, err:\n%v\n")
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// ShowRetentionPolices show retention polices
func ShowRetentionPolices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sqlStr bytes.Buffer
		sqlStr.WriteString(common.ShowRetentionPolices + " ")

		db := c.Query(common.RequestDataBase)
		if db != "" {
			sqlStr.WriteString(fmt.Sprintf("on %s ", db))
		}
		query := origin.New(sqlStr.String())
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("ShowRetentionPolices, err:\n%v\n")
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// ShowSeries show series
func ShowSeries() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sqlStr bytes.Buffer
		sqlStr.WriteString(common.ShowSeries + " ")

		db := c.Query(common.RequestDataBase)
		if db != "" {
			sqlStr.WriteString(fmt.Sprintf("on %s ", db))
		}

		measurement := c.Query(common.RequestMeasurement)
		if measurement != "" {
			sqlStr.WriteString(fmt.Sprintf("from \"%s\" ", measurement))
		}

		// format like: [WHERE <tag_key> <operator> [ '<tag_value>' | <regular_expression>]]
		whereExps := c.QueryArray(common.RequestWhereExpressions)
		if whereExps != nil && len(whereExps) != 0 {
			var _count int
			for _, whereExp := range whereExps {
				if whereExp != "" {
					if _count == 0 {
						sqlStr.WriteString(fmt.Sprintf("where %s ", whereExp))

					} else {
						sqlStr.WriteString(fmt.Sprintf("and %s ", whereExp))
					}
				}
			}
		}

		limitVal := c.Query(common.RequestLimit)
		var limit int
		if limitVal != "" {
			limit1, err := strconv.ParseInt(limitVal, 10, 10)
			if err != nil {
				log.Error("ShowSeries, param field %s error:%v\n", common.RequestLimit, limitVal)
				responseData := common.NewResponseData(400, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			limit = int(limit1)
			if limit > 0 {
				sqlStr.WriteString(fmt.Sprintf("limit %d ", limit))

				offsetVal := c.Query(common.RequestOffset)
				var offset int
				if offsetVal != "" {
					offset1, err := strconv.ParseInt(offsetVal, 10, 10)
					if err != nil {
						log.Error("ShowSeries, param field %s error:%v\n", common.RequestOffset, offsetVal)
						responseData := common.NewResponseData(400, err, nil, c.Request.RequestURI)
						ResponseJSON(c, responseData)
						return
					}
					offset = int(offset1)
				}
				if offset > 0 {
					sqlStr.WriteString(fmt.Sprintf("offset %d", offset))
				}
			}
		}

		query := origin.New(sqlStr.String())
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("ShowSeries, err:\n%v\n")
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// ShowMeasurements show measurements(tables name)
func ShowMeasurements() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sqlStr bytes.Buffer
		sqlStr.WriteString(common.ShowMeasurements + " ")

		db := c.Query(common.RequestDataBase)
		if db != "" {
			sqlStr.WriteString(fmt.Sprintf("on %s ", db))
		}

		withExp := c.Query(common.RequestWithExpression)
		if withExp != "" {
			// WITH MEASUREMENT <regular_expression>
			sqlStr.WriteString(fmt.Sprintf("with measurement %s ", withExp))
		}

		// format like: [WHERE <tag_key> <operator> [ '<tag_value>' | <regular_expression>]]
		whereExps := c.QueryArray(common.RequestWhereExpressions)
		if whereExps != nil && len(whereExps) != 0 {
			var _count int
			for _, whereExp := range whereExps {
				if whereExp != "" {
					if _count == 0 {
						sqlStr.WriteString(fmt.Sprintf("where %s ", whereExp))

					} else {
						sqlStr.WriteString(fmt.Sprintf("and %s ", whereExp))
					}
				}
			}
		}

		limitVal := c.Query(common.RequestLimit)
		var limit int
		if limitVal != "" {
			limit1, err := strconv.ParseInt(limitVal, 10, 10)
			if err != nil {
				log.Error("ShowMeasurements, param field %s error:%v\n", common.RequestLimit, limitVal)
				responseData := common.NewResponseData(400, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			limit = int(limit1)
			if limit > 0 {
				sqlStr.WriteString(fmt.Sprintf("limit %d ", limit))

				offsetVal := c.Query(common.RequestOffset)
				var offset int
				if offsetVal != "" {
					offset1, err := strconv.ParseInt(offsetVal, 10, 10)
					if err != nil {
						log.Error("ShowMeasurements, param field %s error:%v\n", common.RequestOffset, offsetVal)
						responseData := common.NewResponseData(400, err, nil, c.Request.RequestURI)
						ResponseJSON(c, responseData)
						return
					}
					offset = int(offset1)
				}
				if offset > 0 {
					sqlStr.WriteString(fmt.Sprintf("offset %d", offset))
				}
			}
		}

		query := origin.New(sqlStr.String())
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("ShowMeasurements, err: %v", err)
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// ShowTagKeys show tag key
func ShowTagKeys() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sqlStr bytes.Buffer
		sqlStr.WriteString(common.ShowTagKeys + " ")

		db := c.Query(common.RequestDataBase)
		if db != "" {
			sqlStr.WriteString(fmt.Sprintf("on %s ", db))
		}

		measurement := c.Query(common.RequestMeasurement)
		if measurement != "" {
			sqlStr.WriteString(fmt.Sprintf("from \"%s\" ", measurement))
		}

		// format like: [WHERE <tag_key> <operator> [ '<tag_value>' | <regular_expression>]]
		whereExps := c.QueryArray(common.RequestWhereExpressions)
		if whereExps != nil && len(whereExps) != 0 {
			var _count int
			for _, whereExp := range whereExps {
				if whereExp != "" {
					if _count == 0 {
						sqlStr.WriteString(fmt.Sprintf("where %s ", whereExp))

					} else {
						sqlStr.WriteString(fmt.Sprintf("and %s ", whereExp))
					}
				}
			}
		}

		limitVal := c.Query(common.RequestLimit)
		var limit int
		if limitVal != "" {
			limit1, err := strconv.ParseInt(limitVal, 10, 10)
			if err != nil {
				log.Error("ShowTagKeys, param field %s error:%v\n", common.RequestLimit, limitVal)
				responseData := common.NewResponseData(400, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			limit = int(limit1)
			if limit > 0 {
				sqlStr.WriteString(fmt.Sprintf("limit %d ", limit))

				offsetVal := c.Query(common.RequestOffset)
				var offset int
				if offsetVal != "" {
					offset1, err := strconv.ParseInt(offsetVal, 10, 10)
					if err != nil {
						log.Error("ShowTagKeys, param field %s error:%v\n", common.RequestOffset, offsetVal)
						responseData := common.NewResponseData(400, err, nil, c.Request.RequestURI)
						ResponseJSON(c, responseData)
						return
					}
					offset = int(offset1)
				}
				if offset > 0 {
					sqlStr.WriteString(fmt.Sprintf("offset %d", offset))
				}
			}
		}

		query := origin.New(sqlStr.String())
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("ShowTagKeys, err:\n%v\n")
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// ShowTagValues show tag values
func ShowTagValues() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sqlStr bytes.Buffer
		sqlStr.WriteString(common.ShowTagValues + " ")

		db := c.Query(common.RequestDataBase)
		if db != "" {
			sqlStr.WriteString(fmt.Sprintf("on %s ", db))
		}

		measurement := c.Query(common.RequestMeasurement)
		if measurement != "" {
			sqlStr.WriteString(fmt.Sprintf("from \"%s\" ", measurement))
		}

		withExp := c.Query(common.RequestWithExpression)
		if withExp == "" {
			errInfo := fmt.Sprintf("ShowTagValues, param field %s required!", common.RequestWithExpression)
			log.Error(errInfo)
			responseData := common.NewResponseData(400, errInfo, nil, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		// WITH KEY [ [<operator> "<tag_key>" | <regular_expression>] | [IN ("<tag_key1>","<tag_key2")]]
		sqlStr.WriteString(fmt.Sprintf("with key %s ", withExp))

		// format like: [WHERE <tag_key> <operator> [ '<tag_value>' | <regular_expression>]]
		whereExps := c.QueryArray(common.RequestWhereExpressions)
		if whereExps != nil && len(whereExps) != 0 {
			var _count int
			for _, whereExp := range whereExps {
				if whereExp != "" {
					if _count == 0 {
						sqlStr.WriteString(fmt.Sprintf("where %s ", whereExp))

					} else {
						sqlStr.WriteString(fmt.Sprintf("and %s ", whereExp))
					}
				}
			}
		}

		limitVal := c.Query(common.RequestLimit)
		var limit int
		if limitVal != "" {
			limit1, err := strconv.ParseInt(limitVal, 10, 10)
			if err != nil {
				log.Error("ShowTagValues, param field %s error:%v\n", common.RequestLimit, limitVal)
				responseData := common.NewResponseData(400, err, nil, c.Request.RequestURI)
				ResponseJSON(c, responseData)
				return
			}
			limit = int(limit1)
			if limit > 0 {
				sqlStr.WriteString(fmt.Sprintf("limit %d ", limit))

				offsetVal := c.Query(common.RequestOffset)
				var offset int
				if offsetVal != "" {
					offset1, err := strconv.ParseInt(offsetVal, 10, 10)
					if err != nil {
						log.Error("ShowTagValues, param field %s error:%v\n", common.RequestOffset, offsetVal)
						responseData := common.NewResponseData(400, err, nil, c.Request.RequestURI)
						ResponseJSON(c, responseData)
						return
					}
					offset = int(offset1)
				}
				if offset > 0 {
					sqlStr.WriteString(fmt.Sprintf("offset %d", offset))
				}
			}
		}

		query := origin.New(sqlStr.String())
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("ShowTagValues, err:\n%v\n")
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}

// ShowFieldKeys show field keys
func ShowFieldKeys() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sqlStr bytes.Buffer
		sqlStr.WriteString(common.ShowFieldKeys + " ")

		db := c.Query(common.RequestDataBase)
		if db != "" {
			sqlStr.WriteString(fmt.Sprintf("on %s ", db))
		}
		measurement := c.Query(common.RequestMeasurement)
		if measurement != "" {
			sqlStr.WriteString(fmt.Sprintf("from \"%s\" ", measurement))
		}

		query := origin.New(sqlStr.String())
		log.Debug("stmt: %v", query.Stmt)

		res, err := query.Query()
		if err != nil {
			log.Error("ShowFieldKeys, err:\n%v\n")
			responseData := common.NewResponseData(500, err, res, c.Request.RequestURI)
			ResponseJSON(c, responseData)
			return
		}
		responseData := common.NewResponseData(200, nil, res, c.Request.RequestURI)
		ResponseJSON(c, responseData)
	}
}
