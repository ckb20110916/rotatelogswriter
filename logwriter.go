package rotatelogswriter

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func New(folder, filename string, maxAge, rotateTime time.Duration) (logWriter io.Writer) {
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return
	}
	pathname := filepath.ToSlash(path.Join(folder, filename))
	linkName := strings.Join([]string{pathname, ".log"}, "")
	logWriter, err = rotatelogs.New(
		strings.Join([]string{pathname, ".%Y%m%d%H%M", ".log"}, ""),
		rotatelogs.WithLinkName(linkName),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotateTime),
	)
	if err != nil {
		return
	}
	return
}
