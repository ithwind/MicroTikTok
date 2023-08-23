package util

import (
	"fmt"
	"testing"
	"time"
)

func TestConvertTimeFormat(t *testing.T) {

	clock := time.Unix(1692774720, 0)
	fmt.Println(clock)
}
