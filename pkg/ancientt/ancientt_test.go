package ancientt

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderConfigFile(t *testing.T) {
	tests := []struct {
		Data     *TplData
		Expected string
		Error    bool
	}{
		{
			// No data given, error is expected
			Data:     nil,
			Expected: "",
			Error:    true,
		},
		{
			// Sane default values
			Data: &TplData{
				HostNetwork:  false,
				OutputFormat: CSVOutputFormat,
				OutputDir:    "/tmp",
				OutputFile:   "ancientt.csv",
			},
			Expected: `version: '0'
runner:
  name: kubernetes
  kubernetes:
    inClusterConfig: true
    image: 'quay.io/galexrt/container-toolbox:v20231109-123917-402'
    namespace: koor-ancientt
    timeouts:
      deleteTimeout: 20
      runningTimeout: 60
      succeedTimeout: 60
    hosts:
      ignoreSchedulingDisabled: true
      tolerations: []
    hostNetwork: false
tests:
- name: iperf3-all-to-all-hosts
  type: iperf3
  transformations:
  - source: "bits_per_second"
    destination: "gigabits_per_second"
    action: "add"
    modifier: 100000000
    modifierAction: "division"
  outputs:

  - name: csv
    csv:
      filePath: '/tmp'
      namePattern: 'ancientt.csv'

  runOptions:
    continueOnError: true
    rounds: 1
    interval: 10s
    mode: "sequential"
    parallelcount: 1
  hosts:
    clients:
    - name: all-servers
      #all: true
      random: true
      count: 1
    servers:
    - name: all-servers
      #all: true
      random: true
      count: 1
  iperf3:
    udp: false
    duration: 10
    interval: 1
    additionalFlags:
      clients: []
      server: []
`,
			Error: false,
		},
	}

	for _, test := range tests {
		buf := &bytes.Buffer{}
		err := renderConfigFile(buf, test.Data)
		if test.Error {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		if test.Expected != "" {
			assert.Equal(t, test.Expected, buf.String())
		}
	}
}
