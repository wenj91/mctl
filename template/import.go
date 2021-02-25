package template

var Imports = `import cn.zzstc.sbp.common.db.AbstractQuery;
import cn.zzstc.sbp.common.db.Cond;
import cn.zzstc.sbp.common.db.Op;

{{if .bigdecimal}}import java.math.BigDecimal;{{end}}
`
