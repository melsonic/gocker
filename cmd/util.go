package cmd

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
