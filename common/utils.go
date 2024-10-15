package common

import (
	"context"
	"fmt"
	"github.com/valyala/fastrand"
	"reflect"
	"runtime/debug"
	"strings"
)

// Sample do something with a sample rate.
func Sample(sample uint32, f func()) {
	if fastrand.Uint32n(sample) == 0 {
		f()
	}
}

// GetStructName get struct name from instance.
func GetStructName(instance any) string {
	str := reflect.TypeOf(instance).String()
	parts := strings.Split(str, ".")
	right := parts[len(parts)-1]
	return right
}

// Recover recover panic and do something.
func Recover(ctx context.Context, module string, ef func(err any), after func()) {
	if err := recover(); err != nil {
		fmt.Printf(fmt.Sprintf("[%s] module panic: %v, stack: %v", module, err, string(debug.Stack())))
		if ef != nil {
			ef(err)
		}
	}
	if after != nil {
		after()
	}
}
