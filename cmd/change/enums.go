package change

const (
	UseCommand       = `change`
	ShortDescription = `Connection changes`
	LongDescription  = `This command modifies the selected ssh connection`
)

const (
	FilterPlaceholder = "Enter or select a name to change the connection"
	SelectionPrompt   = "Selecting a connection change"
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
		{{- print (Foreground "206" (Bold "▸ ")) (Selected $choice) "\n" }}
	{{- else }}
		{{- print " " (Unselected $choice) "\n" }}
	{{- end }}
{{- end}}`
)

const (
	Placeholder = " не должен быть пустым"
)

var (
	NameAlias    = "Alias"
	NameAddress  = "Address"
	NameLogin    = "Login"
	NamePassword = "Password"
)
