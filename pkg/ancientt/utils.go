package ancientt

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
)

func (r *Runner) prepareConfigAndOutputDir(p *RunParams) (string, error) {
	r.data = &TplData{
		OutputFormat: p.OutputFormat,
		HostNetwork:  p.HostNetwork,
	}

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "dir")
	if err != nil {
		return "", err
	}
	r.data.OutputDir = tmpDir

	testdefsFilePath := path.Join(tmpDir, TestDefinitionsFileName)
	testdefsFile, err := os.Create(testdefsFilePath)
	if err != nil {
		return "", err
	}
	defer testdefsFile.Close()

	// Render ancientt test definitions config
	if err := renderConfigFile(testdefsFile, &TplData{
		HostNetwork:  p.HostNetwork,
		OutputFormat: p.OutputFormat,
		OutputDir:    tmpDir,
		OutputFile:   fmt.Sprintf("ancientt.%s", OutputFormatToFileExtensionMap[p.OutputFormat]),
	}); err != nil {
		return "", err
	}

	return testdefsFilePath, nil
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
