package util

import (
	"fmt"
	"testing"
	"time"
)

func TestConvertTimeFormat(t *testing.T) {
	clock := time.Unix(1690848000, 0)
	fmt.Println(clock)
}
