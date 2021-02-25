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

	javaMysqlDataTypeMap = map[string]string{
		// For consistency, all integer types are converted to Integer
		// number
		"bool":      "Integer",
		"boolean":   "Integer",
		"tinyint":   "Integer",
		"smallint":  "Integer",
		"mediumint": "Integer",
		"int":       "Integer",
		"integer":   "Integer",
		"bigint":    "Integer",
		"float":     "Float",
		"double":    "Double",
		"decimal":   "BigDecimal",
		// date&time
		"date":      "String",
		"datetime":  "String",
		"timestamp": "String",
		"time":      "String",
		"year":      "Integer",
		// String
		"char":       "String",
		"varchar":    "String",
		"binary":     "String",
		"varbinary":  "String",
		"tinytext":   "String",
		"text":       "String",
		"mediumtext": "String",
		"longtext":   "String",
		"enum":       "String",
		"set":        "String",
		"json":       "String",
	}

	entMysqlDataTypeMap = map[string]string{
		// For consistency, all integer types are converted to int64
		// number
		"bool":      "field.Bool(\"%s\").StorageKey(\"%s\")",
		"boolean":   "field.Bool(\"%s\").StorageKey(\"%s\")",
		"tinyint":   "field.Int(\"%s\").StorageKey(\"%s\")",
		"smallint":  "field.Int(\"%s\").StorageKey(\"%s\")",
		"mediumint": "field.Int(\"%s\").StorageKey(\"%s\")",
		"int":       "field.Int(\"%s\").StorageKey(\"%s\")",
		"integer":   "field.Int(\"%s\").StorageKey(\"%s\")",
		"bigint":    "field.Int(\"%s\").StorageKey(\"%s\")",
		"float":     "field.Float(\"%s\").StorageKey(\"%s\")",
		"double":    "field.Float(\"%s\").StorageKey(\"%s\")",
		"decimal":   "field.Float(\"%s\").StorageKey(\"%s\")",
		// date&time
		"date":      "field.Time(\"%s\").StorageKey(\"%s\")",
		"datetime":  "field.Time(\"%s\").StorageKey(\"%s\")",
		"timestamp": "field.Time(\"%s\").StorageKey(\"%s\")",
		"time":      "field.String(\"%s\").StorageKey(\"%s\")",
		"year":      "field.Int(\"%s\").StorageKey(\"%s\")",
		// string
		"char":       "field.String(\"%s\").StorageKey(\"%s\")",
		"varchar":    "field.String(\"%s\").StorageKey(\"%s\")",
		"binary":     "field.String(\"%s\").StorageKey(\"%s\")",
		"varbinary":  "field.String(\"%s\").StorageKey(\"%s\")",
		"tinytext":   "field.String(\"%s\").StorageKey(\"%s\")",
		"text":       "field.String(\"%s\").StorageKey(\"%s\")",
		"mediumtext": "field.String(\"%s\").StorageKey(\"%s\")",
		"longtext":   "field.String(\"%s\").StorageKey(\"%s\")",
		"enum":       "field.Enum(\"%s\").Values(%s).StorageKey(\"%s\")",
		"set":        "field.String(\"%s\").StorageKey(\"%s\")",
		"json":       "field.JSON(\"%s\", %s).StorageKey(\"%s\")",
	}
)

func ConvertDateTypeToJavaType(dataBaseType string) (string, error) {
	tp, ok := javaMysqlDataTypeMap[strings.ToLower(dataBaseType)]
	if !ok {
		return "", fmt.Errorf("unexpected database type: %s", dataBaseType)
	}

	return tp, nil
}

func ConvertDataTypeToEntType(dataBaseType string, name string, colName string, enumOrJson ...string) (string, error) {
	tp, ok := entMysqlDataTypeMap[strings.ToLower(dataBaseType)]
	if !ok {
		return "", fmt.Errorf("unexpected database type: %s", dataBaseType)
	}

	if dataBaseType == "enum" || dataBaseType == "json" {
		ej := ""
		if len(enumOrJson) > 0 {
			ej = enumOrJson[0]
		}

		if ej == "" && dataBaseType == "enum" {
			ej = "\"\""
		}

		if ej == "" && dataBaseType == "json" {
			ej = "[]byte{}"
		}

		return fmt.Sprintf(tp, name, ej, colName), nil
	}

	return fmt.Sprintf(tp, name, colName), nil
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
