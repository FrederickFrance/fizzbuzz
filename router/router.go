package router

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"test.com/fizzbuzz/computation"
	"test.com/fizzbuzz/logger"
	"test.com/fizzbuzz/metrics"
	"test.com/fizzbuzz/utils"

	"github.com/gin-gonic/gin"
)

func FillRouter() {
	r := gin.Default()
	r.GET("/compute", func(c *gin.Context) {

		errs := make([]error, 0)

		// TODO: Can be factorized
		int1, ok := c.GetQuery("int1")
		uint1 := uint64(0)
		if !ok {
			errs = append(errs, errors.New("int1 is missing"))
		} else {
			if value, err := strconv.ParseUint(int1, 10, 64); err != nil {
				errs = append(errs, err)
			} else if value == 0 {
				errs = append(errs, errors.New("int1 must be greater than 0"))
			} else {
				uint1 = value
			}
		}

		int2, ok := c.GetQuery("int2")
		uint2 := uint64(0)
		if !ok {
			errs = append(errs, errors.New("int2 is missing"))
		} else {

			if value, err := strconv.ParseUint(int2, 10, 64); err != nil {
				errs = append(errs, err)
			} else if value == 0 {
				errs = append(errs, errors.New("int2 must be greater than 0"))
			} else {
				uint2 = value
			}
		}

		limit, ok := c.GetQuery("limit")
		ulimit := uint64(0)
		if !ok {
			errs = append(errs, errors.New("limit is missing"))
		} else {
			if value, err := strconv.ParseUint(limit, 10, 64); err != nil {
				errs = append(errs, err)
			} else {
				ulimit = value
			}
		}

		str1, ok := c.GetQuery("str1")
		if !ok {
			errs = append(errs, errors.New("str1 is missing"))
		}

		str2, ok := c.GetQuery("str2")
		if !ok {
			errs = append(errs, errors.New("str2 is missing"))
		}

		if len(errs) > 0 {
			for _, e := range errs {
				logger.Logger.Error(e)
			}
			c.String(http.StatusBadRequest, fmt.Sprint(errs))
		} else {
			param := metrics.Parameters{
				Int1:  uint1,
				Int2:  uint2,
				Limit: ulimit,
				Str1:  str1,
				Str2:  str2,
			}

			if value, ok := metrics.UsedParameters[param]; !ok {
				metrics.UsedParameters[param] = 1
			} else {
				metrics.UsedParameters[param] = value + 1
			}

			stringsArray, errorsArray := computation.Compute(param)
			if len(errorsArray) > 0 {
				for _, e := range errorsArray {
					logger.Logger.Error(e)
				}
				c.String(http.StatusBadRequest, fmt.Sprint(errorsArray))
			} else {
				c.String(http.StatusOK, fmt.Sprint(stringsArray))
			}
		}

	})
	r.GET("/metrics", func(c *gin.Context) {
		result := utils.MostUsed(metrics.UsedParameters)
		if result == nil {
			c.Status(http.StatusNoContent)
		} else {
			c.JSON(http.StatusOK, *result)
		}
	})
	if err := r.Run(); err != nil {
		logger.Logger.WithField("error", err).Panic("Panic during launching")
	}
}
