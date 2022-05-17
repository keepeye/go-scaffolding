package boost

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

// Controller 注册控制器路由
// 用法：Controller(r, "/users", new(User))
// - GET /users/list --> *User.GetList
// - GET /users/profile/:name/:age --> *User.GetProfileByNameByAge
// - POST /users --> *Users.Post
func RegisterController(router *echo.Group, structObj interface{}) {
	snake := regexp.MustCompile("([A-Z])")
	reflectValue := reflect.ValueOf(structObj)
	reflectType := reflectValue.Type()
	for i := 0; i < reflectType.NumMethod(); i++ {
		method := reflectValue.Method(i)
		methodType := reflectType.Method(i)
		methodName := methodType.Name
		// 方法名必须以 Get Post Put Delete 开头
		verb := ""
		for _, v := range []string{"Get", "Post", "Put", "Delete"} {
			if strings.HasPrefix(methodName, v) {
				verb = v
				break
			}
		}
		if "" == verb {
			continue
		}
		methodName = methodName[len(verb):]
		segments := make([]string, 0)
		if methodName != "" {
			// 将方法名转换为下划线小写格式
			methodName = snake.ReplaceAllString(methodName, "_${1}")
			methodName = strings.ToLower(methodName)
			if methodName[0] == '_' {
				methodName = methodName[1:]
			}
			// 将方法名按下划线分割
			segments = strings.Split(methodName, "_")
		}
		k := 0
		pathes := make([]string, 0, len(segments))
		for k < len(segments) {
			cur := segments[k]
			if cur == "by" && k < len(segments)-1 {
				k++
				pathes = append(pathes, ":"+segments[k])
			} else {
				pathes = append(pathes, cur)
			}
			k++
		}
		var path = ""
		if len(pathes) > 0 {
			path = "/" + strings.Join(pathes, "/")
		}
		verb = strings.ToUpper(verb)
		handler := method.Interface().(func(echo.Context) error)
		router.Add(verb, path, handler).Name = fmt.Sprintf("%s.%s", reflectType.String(), methodType.Name)
	}
}
