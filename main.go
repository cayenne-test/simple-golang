package lib

import "github.com/kimh/simple-golang/dep"

func String() string {
	return "Hello World!"
}

func Number() int {
	return 1
}

func Converted() int {
	return dep.Convert(1)
}
