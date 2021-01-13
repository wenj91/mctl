package template

var ToStr = `
func (m *{{.upperStartCamelObject}}) ToString() string {
	str := ""

	bs, err := json.Marshal(m)
	if nil == err {
		str = string(bs)
	}

	return str
}
`
