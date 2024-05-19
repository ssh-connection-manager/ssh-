package delete

const (
	UseCommand       = `delete`
	ShortDescription = `Удаление подключение по алиасу`
	LongDescription  = `Данная команда позволяет удалить подключение по алиасу`
)

const (
	FilterPlaceholder = "Введите или выбирите название для удаление подключения"
	SelectionPrompt   = "Выбор удаление подключения"
	FilterPrompt      = "Поиск по алиуса:"
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
		{{- print (Foreground "201" (Bold "× ")) (Selected $choice) "\n" }}
	{{- else }}
		{{- print " " (Unselected $choice) "\n" }}
	{{- end }}
{{- end}}`
)
