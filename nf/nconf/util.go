package nconf

import (
	"fmt"
	"strconv"
)

func atoi(v string) int {
	if intV, err := strconv.Atoi(v); err == nil {
		return intV
	}
	panic(fmt.Errorf("The value %s can't be convert to int", v))
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
