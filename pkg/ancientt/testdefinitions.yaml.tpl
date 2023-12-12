version: '0'
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
    hostNetwork: {{ .HostNetwork }}
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
{{ if eq .OutputFormat "excelize" }}
  - name: excelize
    excelize:
      filePath: '{{ .OutputDir }}'
      namePattern: "ancientt.xlsx"
      saveAfterRows: 100
{{ else }}
  - name: csv
    csv:
      filePath: '{{ .OutputDir }}'
      namePattern: 'ancientt.csv'
{{ end }}
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
