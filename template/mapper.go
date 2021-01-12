package template

var Mapper = `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE mapper PUBLIC "gobatis"
        "https://raw.githubusercontent.com/wenj91/gobatis/master/gobatis.dtd">
<mapper namespace="{{.mapper}}Mapper">
  {{.stmts}}
</mapper>
`
