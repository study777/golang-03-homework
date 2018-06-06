package main

func main() {

	fanIn("cc","gg")

}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s

			}
		}
	}()
	return c
}
