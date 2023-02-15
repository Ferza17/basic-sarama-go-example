package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetRandomString() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func GetRandomSixNumber() string {
	s := ""
	for i := 0; i < 6; i++ {
		rand.Seed(time.Now().UnixNano())
		min := 0
		max := 9
		s += strconv.Itoa(rand.Intn(max-min+1) + min)
	}
	if strings.HasPrefix(s, "0") {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 9
		s = strings.Replace(s, "0", strconv.Itoa(rand.Intn(max-min+1)+min), 1)
	}
	return s
}
