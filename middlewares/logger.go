package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logTimestamps := strconv.FormatInt(param.TimeStamp.Unix(), 10)
		var logString = param.Method +","+ param.Latency.String() +","+ logTimestamps +"\n"
		f, err := os.OpenFile("api.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(logString); err != nil {
			log.Println(err)
		}
		return logString
	})
}
