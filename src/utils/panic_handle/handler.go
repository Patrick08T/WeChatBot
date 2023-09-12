package panic_handle

import (
	"ZhuRong/backend/src/util/log"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
	"runtime/debug"
)

func HandlePanic(f func(x interface{})) {
	if x := recover(); x != nil {
		f(x)
	}
}

func PCall(f func()) {
	defer HandlePanic(func(x interface{}) {
		PanicLog(x)
	})
	f()
}

func PanicLog(x interface{}) {
	stack := string(debug.Stack())
	log.ERRORF("PCall panic recover! p=%v stack=%v", x, stack)
}

func ExitSystem() {
	stack := string(debug.Stack())
	log.ERROR(stack)
	// TODO 确保日志落地
	os.Exit(1)
}

func HandleEchoPanic() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			if middleware.DefaultRecoverConfig.Skipper(context) {
				return next(context)
			}
			defer func() {
				if x := recover(); x != nil {
					func(x interface{}) {
						err, ok := x.(error)
						if !ok {
							err = fmt.Errorf("%v", x)
						}
						if !middleware.DefaultRecoverConfig.DisablePrintStack {
							PanicLog(x)
						}
						context.Error(err)
					}(x)
				}
			}()
			return next(context)
		}
	}
}
