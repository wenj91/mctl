package template

var Tag = "`{{if .isDBField}}field:\"{{.field}}\" json:\"{{.json}}\"{{else}}json:\"-\"{{end}}`"
