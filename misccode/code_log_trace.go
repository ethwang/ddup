package misccode

import (
	"context"
	"fmt"
)

func DEBUG(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
func DEBUGT(ctx context.Context, format string, a ...interface{}) {
	if ctx == nil {
		DEBUG(format, a)
		return
	}
	format = "[TRACEID:%d] " + format
	fmt.Println("format: ", format)

	a = append([]interface{}{ctx.Value("TRACEID")}, a...)
	fmt.Printf(format, a...)
}

func INFOT(ctx context.Context, format string, a ...interface{}) {
	format = "[TRACEID:%d]" + format
	a = append([]interface{}{ctx.Value("TRACEID").(int64)}, a...)
	fmt.Printf(format, a...)
}

func WARNT(ctx context.Context, format string, a ...interface{}) {
	format = "[TRACEID:%d]" + format
	a = append([]interface{}{ctx.Value("TRACEID").(int64)}, a...)
	fmt.Printf(format, a...)
}

func ERRORT(ctx context.Context, format string, a ...interface{}) {
	format = "[TRACEID:%d]" + format
	a = append([]interface{}{ctx.Value("TRACEID").(int64)}, a...)
	fmt.Printf(format, a...)
}
