package template

var Fields = `    @Override
    public String[] fields() {
        return new String[]{{print "{" .fields print "}"}};
    }`
