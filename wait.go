package goutil

func WaitForever() {
	forever := make(chan bool)
	<-forever
}
