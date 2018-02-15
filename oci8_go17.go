// +build !go1.8

package oci8

func handleOutput(vv interface{}) (outValue, bool) {
	return outValue{}, false
}
