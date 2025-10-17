{{/* template for license-dependencies.md, see readme for details */}}
### dependency licenses
below are the licenses of every dependency module.

| module | license |
| --- | --- |
{{- range . }}
| {{ .Name }} | [{{ .LicenseName }}]({{ .LicenseURL }}) |
{{- end }}