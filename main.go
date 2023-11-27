package main

type TestData struct {
	Num int32
}

func GetTestDataValue() TestData {
	return TestData{}
}

func GetTestDataPtr() *TestData {
	return &TestData{}
}

//go:noinline
func GetTestDataValueNoInline() TestData {
	return TestData{}
}

//go:noinline
func GetTestDataPtrNoInline() *TestData {
	return &TestData{}
}
