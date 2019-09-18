package main

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	templates, err := GetTemplates()
	panicIfError(err)

	methods, err := ParseMethods()
	panicIfError(err)

	panicIfError(WriteFiles(templates, methods))
}
