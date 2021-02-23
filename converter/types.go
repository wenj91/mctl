package converter

import (
	"fmt"
	"strings"
)

var (
	commonMysqlDataTypeMap = map[string]string{
		// For consistency, all integer types are converted to int64
		// number
		"bool":      "int64",
		"boolean":   "int64",
		"tinyint":   "int64",
		"smallint":  "int64",
		"mediumint": "int64",
		"int":       "int64",
		"integer":   "int64",
		"bigint":    "int64",
		"float":     "float64",
		"double":    "float64",
		"decimal":   "float64",
		// date&time
		"date":      "time.Time",
		"datetime":  "time.Time",
		"timestamp": "time.Time",
		"time":      "string",
		"year":      "int64",
		// string
		"char":       "string",
		"varchar":    "string",
		"binary":     "string",
		"varbinary":  "string",
		"tinytext":   "string",
		"text":       "string",
		"mediumtext": "string",
		"longtext":   "string",
		"enum":       "string",
		"set":        "string",
		"json":       "string",
	}

	entMysqlDataTypeMap = map[string]string{
		// For consistency, all integer types are converted to int64
		// number
		"bool":      "field.Bool(\"%s\")",
		"boolean":   "field.Bool(\"%s\")",
		"tinyint":   "field.Int(\"%s\")",
		"smallint":  "field.Int(\"%s\")",
		"mediumint": "field.Int(\"%s\")",
		"int":       "field.Int(\"%s\")",
		"integer":   "field.Int(\"%s\")",
		"bigint":    "field.Int(\"%s\")",
		"float":     "field.Float(\"%s\")",
		"double":    "field.Float(\"%s\")",
		"decimal":   "field.Float(\"%s\")",
		// date&time
		"date":      "field.Time(\"%s\")",
		"datetime":  "field.Time(\"%s\")",
		"timestamp": "field.Time(\"%s\")",
		"time":      "field.String(\"%s\")",
		"year":      "field.Int(\"%s\")",
		// string
		"char":       "field.String(\"%s\")",
		"varchar":    "field.String(\"%s\")",
		"binary":     "field.String(\"%s\")",
		"varbinary":  "field.String(\"%s\")",
		"tinytext":   "field.String(\"%s\")",
		"text":       "field.String(\"%s\")",
		"mediumtext": "field.String(\"%s\")",
		"longtext":   "field.String(\"%s\")",
		"enum":       "field.Enum(\"%s\").Values(%s)",
		"set":        "field.String(\"%s\")",
		"json":       "field.JSON(\"%s\", %s)",
	}
)

func ConvertDataTypeToEntType(dataBaseType string, name string, enumOrJson ...interface{}) (string, error) {
	tp, ok := entMysqlDataTypeMap[strings.ToLower(dataBaseType)]
	if !ok {
		return "", fmt.Errorf("unexpected database type: %s", dataBaseType)
	}

	if dataBaseType == "enum" || dataBaseType == "json" {
		return fmt.Sprintf(tp, name, enumOrJson), nil
	}

	return fmt.Sprintf(tp, name), nil
}

func ConvertDataType(dataBaseType string, isDefaultNull bool) (string, error) {
	tp, ok := commonMysqlDataTypeMap[strings.ToLower(dataBaseType)]
	if !ok {
		return "", fmt.Errorf("unexpected database type: %s", dataBaseType)
	}

	return mayConvertNullType(tp, isDefaultNull), nil
}

func mayConvertNullType(goDataType string, isDefaultNull bool) string {
	if !isDefaultNull {
		return goDataType
	}

	switch goDataType {
	case "int64":
		return "*int64"
	case "int32":
		return "*int64"
	case "float64":
		return "*float64"
	case "bool":
		return "*bool"
	case "string":
		return "*string"
	case "time.Time":
		return "*time.Time"
	default:
		return goDataType
	}
}
