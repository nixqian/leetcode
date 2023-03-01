package godemo

func testChannel() {
	c := make(chan struct{})
	c <- struct{}{}
	var a rune
}
