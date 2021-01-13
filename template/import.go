package template

var (
	Imports = `import (
	"encoding/json"
	"database/sql"
	{{if .time}}"time"{{end}}

	"github.com/wenj91/gobatis"
)
`
	ImportsNoCache = `import (
	"encoding/json"
	"database/sql"
	{{if .time}}"time"{{end}}

	"github.com/wenj91/gobatis"
)
`
)
