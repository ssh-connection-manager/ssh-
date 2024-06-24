package connect

const (
	UseCommand       = `connect`
	ShortDescription = `Output connected ssh connections`
	LongDescription  = `This command displays a list of your added ssh connections,
if none are available, it will ask you to add connections`
)

const (
	FilterPlaceholder = "Enter or select the name of the connection"
	SelectionPrompt   = "Connection selection"
	FilterPrompt      = "Search by alias:"
	PageSize          = "5"
)

const (
	Template = `
{{- if .Prompt -}}
	{{ Bold .Prompt }}
{{ end -}}

{{ if .IsFiltered }}
	{{- print .FilterPrompt " " .FilterInput }}
{{ end }}

{{- range  $i, $choice := .Choices }}
	{{- if IsScrollUpHintPosition $i }}
		{{- print "⇡ " -}}
	{{- else if IsScrollDownHintPosition $i -}}
		{{- print "⇣ " -}}
	{{- else -}}
		{{- print "  " -}}
	{{- end -}}
	
	{{- if eq $.SelectedIndex $i }}
		{{- print (Foreground "201" (Bold "▸ ")) (Selected $choice) "\n" }}
	{{- else }}
		{{- print " " (Unselected $choice) "\n" }}
	{{- end }}
{{- end}}`
)
