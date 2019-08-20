package models

import (
	"fmt"

	"github.com/mrz1836/go-logger"
)

var test string

func init() {
	test = "this"
	logger.Data(2, logger.DEBUG, fmt.Sprintf("Testing! %s", test))
}
