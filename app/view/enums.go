package view

const ResultTemplateTextInput = `
{{- print .Prompt " " (Foreground "206"  (Mask .FinalValue)) "\n" -}}
`

const ResultTemplateSelect = `
{{- print .Prompt " " (Foreground "206"  (alias .FinalChoice)) "\n" -}}
`

const ResultChoiceName = "alias"
