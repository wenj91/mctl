package template

var (
	Imports = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	jsoniter "github.com/json-iterator/go"
	"github.com/wenj91/gobatis"
)
`
	ImportsNoCache = `import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	jsoniter "github.com/json-iterator/go"
	"github.com/wenj91/gobatis"
)
`
)
