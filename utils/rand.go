package utils

//生成随机数
import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandInt(n int) string {
	s := fmt.Sprintf("%12v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000))
	return strings.TrimSpace(s)
}
