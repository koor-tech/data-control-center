package ancientt

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/koor-tech/data-control-center/pkg/config"
	"go.uber.org/fx"
)

var (
	//go:embed *.tpl
	testdefsTemplates embed.FS
)

const (
	TestDefinitionsFileName = "testdefinitions.yaml"
)

type OutputFormat string

const (
	CSVOutputFormat      OutputFormat = "csv"
	ExcelizeOutputFormat OutputFormat = "excelize"
)

var (
	OutputFormatToFileExtensionMap = map[OutputFormat]string{
		CSVOutputFormat:      "csv",
		ExcelizeOutputFormat: "xlsx",
	}
	OutputFormatToMetaMap = map[OutputFormat]string{
		CSVOutputFormat:      "text/csv",
		ExcelizeOutputFormat: "application/vnd.ms-excel",
	}
)

type TplData struct {
	OutputDir  string
	OutputFile string

	// Ancientt Options/Toggles
	HostNetwork  bool
	OutputFormat OutputFormat
}

type RunParams struct {
	HostNetwork  bool
	OutputFormat OutputFormat
}

type Runner struct {
	mutex sync.Mutex

	cmdPath string

	run  IRun
	data *TplData
}

func New(lc fx.Lifecycle, cfg *config.Config) *Runner {
	command := "ancientt"

	if cfg.AncienttCmd != "" {
		command = cfg.AncienttCmd
	}

	r := &Runner{
		mutex: sync.Mutex{},

		cmdPath: command,
	}

	lc.Append(fx.StopHook(func(_ context.Context) error {
		return r.Cancel()
	}))

	return r
}

func (r *Runner) IsRunning() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.run != nil && r.run.IsRunning()
}

func (r *Runner) IsComplete() (bool, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.run == nil {
		return false, nil
	}

	logs, err := r.run.GetLogs()
	if err != nil {
		return false, err
	}

	return strings.Contains(logs, "done with tests"), nil
}

func (r *Runner) Start(p *RunParams) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	testDefsFile, err := r.prepareConfigAndOutputDir(p)
	if err != nil {
		return err
	}

	args := []string{r.cmdPath, "--testdefinition=" + testDefsFile, "--yes"}

	run, err := NewRun(r.data.OutputDir, "ancientt", args)
	if err != nil {
		return err
	}

	r.run = run

	if err := r.run.Start(); err != nil {
		return err
	}

	return nil
}

func (r *Runner) Cancel() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := r.run.Stop(); err != nil {
		return err
	}

	r.run = nil

	if err := os.RemoveAll(r.data.OutputDir); err != nil {
		return err
	}

	r.data = nil

	return nil
}

func (r *Runner) Stop() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.run == nil {
		return nil
	}

	if err := r.run.Stop(); err != nil {
		return err
	}

	return nil
}

func (r *Runner) GetLogs() (string, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.run.GetLogs()
}

func (r *Runner) GetResultFile() (string, string, []byte, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.run == nil || r.data == nil {
		return "", "", []byte{}, fmt.Errorf("no run data found")
	}

	resultFilePath := path.Join(r.data.OutputDir, r.data.OutputFile)
	out, err := os.ReadFile(resultFilePath)
	if err != nil {
		return "", "", []byte{}, err
	}

	fileMeta := OutputFormatToMetaMap[r.data.OutputFormat]
	return fmt.Sprintf("ancientt-%s.csv", time.Now().Format("2006-01-02")), fileMeta, out, nil
}
