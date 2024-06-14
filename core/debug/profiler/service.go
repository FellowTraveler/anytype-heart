package profiler

import (
	"archive/zip"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"

	"github.com/anyproto/any-sync/app"

	"github.com/anyproto/anytype-heart/pkg/lib/logging"
	"github.com/anyproto/anytype-heart/util/debug"
)

var log = logging.Logger("profiler")

type Service interface {
	app.ComponentRunnable

	RunProfiler(ctx context.Context, seconds int) (string, error)
}

type service struct {
	closeCh chan struct{}

	timesHighMemoryUsageDetected int
	previousHighMemoryDetected   uint64
}

func New() Service {
	return &service{
		closeCh: make(chan struct{}),
	}
}

func (s *service) Init(a *app.App) (err error) {
	return nil
}

func (s *service) Name() (name string) {
	return "profiler"
}

func (s *service) Run(ctx context.Context) (err error) {
	go s.run()

	return nil
}

func (s *service) RunProfiler(ctx context.Context, seconds int) (string, error) {
	var tracerBuf bytes.Buffer
	err := trace.Start(&tracerBuf)
	if err != nil {
		return "", fmt.Errorf("start tracer: %w", err)
	}

	var cpuProfileBuf bytes.Buffer
	err = pprof.StartCPUProfile(&cpuProfileBuf)
	if err != nil {
		return "", fmt.Errorf("start cpu profile: %w", err)
	}

	var heapStartBuf bytes.Buffer
	err = pprof.WriteHeapProfile(&heapStartBuf)
	if err != nil {
		return "", fmt.Errorf("write starting heap profile: %w", err)
	}

	goroutinesStart := debug.Stack(true)

	select {
	case <-time.After(time.Duration(seconds) * time.Second):
	case <-ctx.Done():
	case <-s.closeCh:
	}

	pprof.StopCPUProfile()
	trace.Stop()
	var heapEndBuf bytes.Buffer
	err = pprof.WriteHeapProfile(&heapEndBuf)
	if err != nil {
		return "", fmt.Errorf("write ending heap profile: %w", err)
	}
	goroutinesEnd := debug.Stack(true)

	f, err := os.CreateTemp("", "anytype_profile.*.zip")
	if err != nil {
		return "", fmt.Errorf("create temp file: %w", err)
	}
	err = createZipArchive(f, []zipFile{
		{name: "trace", data: &tracerBuf},
		{name: "cpu_profile", data: &cpuProfileBuf},
		{name: "heap_start", data: &heapStartBuf},
		{name: "heap_end", data: &heapEndBuf},
		{name: "goroutines_start.txt", data: bytes.NewReader(goroutinesStart)},
		{name: "goroutines_end.txt", data: bytes.NewReader(goroutinesEnd)},
	})
	if err != nil {
		return "", errors.Join(fmt.Errorf("create zip archive: %w", err), f.Close())
	}
	return f.Name(), f.Close()
}

type zipFile struct {
	name string
	data io.Reader
}

func createZipArchive(w io.Writer, files []zipFile) error {
	zipw := zip.NewWriter(w)
	err := func() error {
		for _, file := range files {
			f, err := zipw.Create(file.name)
			if err != nil {
				return fmt.Errorf("create file in zip archive: %w", err)
			}
			_, err = io.Copy(f, file.data)
			if err != nil {
				return fmt.Errorf("copy data to file: %w", err)
			}
		}
		return nil
	}()
	return errors.Join(err, zipw.Close())
}

func (s *service) Close(ctx context.Context) (err error) {
	close(s.closeCh)
	return nil
}
