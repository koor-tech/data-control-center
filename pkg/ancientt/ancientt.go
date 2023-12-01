package ancientt

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"path"
	"sync"
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

var OutputFormatToFileExtensionMap = map[OutputFormat]string{
	CSVOutputFormat:      "csv",
	ExcelizeOutputFormat: "xlsx",
}

type TplData struct {
	OutputDir  string
	OutputFile string

	// Ancientt Options/Toggles
	HostNetwork  bool
	OutputFormat OutputFormat
}

type Runner struct {
	mutex sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc

	cmd *exec.Cmd

	outputDir  string
	outputFile string
}

func New() *Runner {
	return &Runner{
		mutex: sync.Mutex{},
	}
}

func (r *Runner) IsRunning() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.cmd != nil
}

func (r *Runner) Start(outputFormat OutputFormat) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	outputDir, testDefsFile, err := prepareConfigAndOutputDir(outputFormat)
	if err != nil {
		return err
	}

	r.outputDir = outputDir

	ctx, cancel := context.WithCancel(context.Background())

	r.ctx = ctx
	r.cancel = cancel

	args := []string{"ancientt", "--testdefinition=" + testDefsFile, "-y"}
	cmd := exec.CommandContext(ctx, "ancientt", args...)

	if err := cmd.Start(); err != nil {
		return err
	}

	r.cmd = cmd

	return nil
}

func (r *Runner) Cancel() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.cancel()

	if err := r.cmd.Wait(); err != nil {
		return err
	}

	r.cmd = nil
	r.cancel = nil
	r.ctx = nil

	r.outputDir = ""
	r.outputFile = ""

	return nil
}

func (r *Runner) Cleanup() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.outputDir == "" {
		return nil
	}

	return os.RemoveAll(r.outputDir)
}

func (r *Runner) GetResultFile() (string, []byte, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	resultFile := fmt.Sprintf("ancientt.%s", CSVOutputFormat)
	_ = resultFile

	// TODO

	fileContent := []byte{}

	return r.outputFile, fileContent, nil
}

func prepareConfigAndOutputDir(outputFormat OutputFormat) (string, string, error) {
	// create a temporary directory
	tmpDir, err := os.MkdirTemp("", "dir") // in Go version older than 1.17 you can use ioutil.TempDir
	if err != nil {
		return "", "", err
	}

	testdefsFilePath := path.Join(tmpDir, TestDefinitionsFileName)
	testdefsFile, err := os.Create(testdefsFilePath)
	if err != nil {
		return "", "", err
	}
	defer testdefsFile.Close()

	if err := renderConfigFile(testdefsFile, &TplData{
		HostNetwork:  false,
		OutputFormat: outputFormat,
		OutputDir:    tmpDir,
		OutputFile:   fmt.Sprintf("ancientt.%s", outputFormat),
	}); err != nil {
		return "", "", err
	}

	return tmpDir, testdefsFilePath, nil
}

func renderConfigFile(out io.Writer, tplData *TplData) error {
	t, err := template.New("").ParseFS(testdefsTemplates, "**.tpl")
	if err != nil {
		return err
	}

	if err := t.ExecuteTemplate(out, "testdefinitions.yaml.tpl", tplData); err != nil {
		return err
	}

	return nil
}
