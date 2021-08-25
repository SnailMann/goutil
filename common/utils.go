package common

import "github.com/valyala/fastrand"

// sample to do some thing
func Sample(sample uint32, f func()) {
	if fastrand.Uint32n(sample) == 0 {
		f()
	}
}
