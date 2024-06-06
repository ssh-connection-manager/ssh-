package change

const (
	UseCommand       = `change`
	ShortDescription = `Изменения подключения`
	LongDescription  = `Данная команда изменяет выбранное подключения по ssh`
)

const (
	FilterPlaceholder = "Введите или выбирите название для изменения подключения"
	SelectionPrompt   = "Выбор изменения подключения"
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

const (
	Placeholder = " не должен быть пустым"
)

var (
	NameAlias    = "Алиас"
	NameAddress  = "Адресс"
	NameLogin    = "Логин"
	NamePassword = "Пароль"
)
