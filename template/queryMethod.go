package template

var QueryMethod = `    public static {{.upperStartCamelObject}}Query query() {
        return new {{.upperStartCamelObject}}Query();
    }
`
