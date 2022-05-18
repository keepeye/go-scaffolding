package boost

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func CustomHttpLogger(tags []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			logFields := logrus.Fields{
				"module": "accesslog",
			}

			reqPath := req.URL.Path
			if reqPath == "" {
				reqPath = "/"
			}

			for _, tag := range tags {
				switch tag {

				case "id":
					id := req.Header.Get(echo.HeaderXRequestID)
					if id == "" {
						id = res.Header().Get(echo.HeaderXRequestID)
					}
					logFields["id"] = id
				case "remote_ip":
					logFields["remote_ip"] = c.RealIP()
				case "host":
					logFields["host"] = req.Host
				case "uri":
					logFields["uri"] = req.RequestURI
				case "method":
					logFields["method"] = req.Method
				case "path":
					logFields["path"] = reqPath
				case "protocol":
					logFields["protocol"] = req.Proto
				case "referer":
					logFields["referer"] = req.Referer()
				case "user_agent":
					logFields["user_agent"] = req.UserAgent()
				case "status":
					n := res.Status
					logFields["status"] = n
				case "error":
					if err != nil {
						// Error may contain invalid JSON e.g. `"`
						b, _ := json.Marshal(err.Error())
						b = b[1 : len(b)-1]
						logFields["error"] = string(b)
					}
				case "latency":
					l := stop.Sub(start)
					logFields["latency"] = strconv.FormatInt(int64(l), 10)
				case "latency_human":
					logFields["latency_human"] = stop.Sub(start).String()
				case "bytes_in":
					cl := req.Header.Get(echo.HeaderContentLength)
					if cl == "" {
						cl = "0"
					}
					logFields["bytes_in"] = cl
				case "bytes_out":
					logFields["bytes_out"] = strconv.FormatInt(res.Size, 10)
				default:
					switch {
					case strings.HasPrefix(tag, "header:"):
						logFields[tag[7:]] = c.Request().Header.Get(tag[7:])
					case strings.HasPrefix(tag, "query:"):
						logFields[tag[6:]] = c.QueryParam(tag[6:])
					case strings.HasPrefix(tag, "form:"):
						logFields[tag[5:]] = c.FormValue(tag[5:])
					case strings.HasPrefix(tag, "cookie:"):
						cookie, err := c.Cookie(tag[7:])
						if err == nil {
							logFields[tag[7:]] = cookie.Value
						}
					}
				}
			}
			logrus.WithFields(logFields).Infof("%s %s", req.Method, req.RequestURI)
			return
		}
	}
}
