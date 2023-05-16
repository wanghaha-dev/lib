package ghelper

func Loops(count int, f func()) {
	for i := 0; i < count; i++ {
		f()
	}
}
