package cache

import (
	"fmt"
	"strconv"
)

var (
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:prodcut:%s", strconv.Itoa(int(id)))
}
