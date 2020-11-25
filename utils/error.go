package utils

func panicOnNonNil(err error) {
	if err != nil {
		panic(err)
	}
}
