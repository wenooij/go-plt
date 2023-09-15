package plt

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/golang/glog"
)

type seriesKey = string

func runtimeSeriesKey(skip int) seriesKey {
	_, file, line, _ := runtime.Caller(skip)
	file = strings.TrimSuffix(file, filepath.Ext(file))
	key := fmt.Sprintf("%s:%d", file, line)
	if glog.V(1) {
		glog.Infof("runtimeSeriesKey(%d) created: %q", skip, key)
	}
	return key
}

func numericSeriesKey(n uint) seriesKey {
	return fmt.Sprintf("/series/%d", n)
}

func seriesName(s seriesKey) string {
	const maxSeriesNameLen = 12
	s = filepath.Base(s)
	if n := len(s); n > maxSeriesNameLen {
		s = fmt.Sprint(s[:maxSeriesNameLen/2], "...", s[n-maxSeriesNameLen/2:])
	}
	return s
}
