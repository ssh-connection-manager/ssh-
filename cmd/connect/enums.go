package connect

const (
	UseCommand       = `connect`
	ShortDescription = `Вывод подключенных ssh соединений`
	LongDescription  = `Данная команда выводит список ваших добавленных соедениний по ssh,
в случае его отсутвия попросит добавить подключения`
)

const (
	FilterPlaceholder = "Введите или выбирите название подключения"
	SelectionPrompt   = "Выбор подключения"
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
		{{- print (Foreground "201" (Bold "▸ ")) (Selected $choice) "\n" }}
	{{- else }}
		{{- print " " (Unselected $choice) "\n" }}
	{{- end }}
{{- end}}`
)
