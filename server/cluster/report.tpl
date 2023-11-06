{{ define "report" -}}
# Troubleshooting Report
{{ range .Contents  }}
## {{ .Name }}
{{ with .Description }}
{{ . }}
{{ end }}
```console
{{ if .Content }}
{{- .Content -}}
{{ else }}
No Content
{{ end }}
```
{{- with .Error }}
Error during data collection:
```console
{{ . }}
```
{{ end }}
{{ end }}
_Generated using Koor data-control-center (v{{ .Version }}) on {{ now | date "Mon Jan 2 15:04:05 MST 2006" }}_.
{{- end -}}
