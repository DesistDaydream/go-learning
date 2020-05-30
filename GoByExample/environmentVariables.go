// 环境变量
package main
import (
	"os"
	"strings"
	"fmt"
)
func main() {
	os.Setenv("FOO", "1")
}