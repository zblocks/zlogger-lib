package zlogger_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/zblocks/zlogger-lib"
)

func TestUtils(t *testing.T) {
  t.Run("Test Create logger name", func(t *testing.T) {
    loggername := zlogger.CreateLoggerName("svc_name", "pkg_name", "file_name", "function_name")
    assert.Equal(t, "svc_name.pkg_name.file_name.function_name", loggername)
  })
}