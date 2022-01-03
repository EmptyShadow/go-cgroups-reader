package cgroupsreader_test

import (
	"path/filepath"
	"runtime"
	"testing"
	"time"

	cgroupsreader "github.com/EmptyShadow/go-cgroups-reader/v1"
	"github.com/stretchr/testify/require"
)

func Test_ReadResources_NotExists(t *testing.T) {
	requiring := require.New(t)

	_, enabled, err := cgroupsreader.ReadResources("/", "")
	requiring.NoError(err)
	requiring.False(enabled)
}

func Test_ReadResources(t *testing.T) {
	requiring := require.New(t)
	_, file, _, _ := runtime.Caller(0)
	dir, _ := filepath.Split(file)

	testdata := filepath.Join(dir, "testdata")

	recources, enabled, err := cgroupsreader.ReadResources(testdata, "test")
	requiring.NoError(err)
	requiring.True(enabled)
	requiring.Equal(recources.CPU.PeriodUS, time.Millisecond*100)
	requiring.Equal(recources.CPU.QuotaUS, time.Microsecond*-1)
	requiring.Equal(recources.Memory.LimitInBytes, uint64(9223372036854771712))
	requiring.Equal(recources.Memory.MemSwapLimitInBytes, uint64(9223372036854771712))
}

func Test_ReadCPUResources_NotExists(t *testing.T) {
	requiring := require.New(t)

	_, enabled, err := cgroupsreader.ReadCPUResources("/", "")
	requiring.NoError(err)
	requiring.False(enabled)
}

func Test_ReadCPUResources(t *testing.T) {
	requiring := require.New(t)
	_, file, _, _ := runtime.Caller(0)
	dir, _ := filepath.Split(file)

	testdata := filepath.Join(dir, "testdata")

	recources, enabled, err := cgroupsreader.ReadCPUResources(testdata, "test")
	requiring.NoError(err)
	requiring.True(enabled)
	requiring.Equal(recources.PeriodUS, time.Millisecond*100)
	requiring.Equal(recources.QuotaUS, time.Microsecond*-1)
}

func Test_ReadMemoryResources_NotExists(t *testing.T) {
	requiring := require.New(t)

	_, enabled, err := cgroupsreader.ReadMemoryResources("/", "")
	requiring.NoError(err)
	requiring.False(enabled)
}

func Test_ReadMemoryResources(t *testing.T) {
	requiring := require.New(t)
	_, file, _, _ := runtime.Caller(0)
	dir, _ := filepath.Split(file)

	testdata := filepath.Join(dir, "testdata")

	recources, enabled, err := cgroupsreader.ReadMemoryResources(testdata, "test")
	requiring.NoError(err)
	requiring.True(enabled)
	requiring.Equal(recources.LimitInBytes, uint64(9223372036854771712))
	requiring.Equal(recources.MemSwapLimitInBytes, uint64(9223372036854771712))
}
