package delete

const (
	UseCommand       = `delete`
	ShortDescription = `Removing an alias connection`
	LongDescription  = `This command allows you to remove an alias connection`
)

const (
	FilterPlaceholder = "Enter or select a name to delete the connection"
	SelectionPrompt   = "Selecting whether to remove the connection"
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
		{{- print (Foreground "206" (Bold "× ")) (Selected $choice) "\n" }}
	{{- else }}
		{{- print " " (Unselected $choice) "\n" }}
	{{- end }}
{{- end}}`
)
