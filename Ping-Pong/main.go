package pingpong

// Ping-Pong Using Channels: Implement a program where two goroutines send "ping" and "pong" messages back and forth using channels.
func PingPong(n int) []string {
	sender := make(chan string)
	receiver := make(chan string)
	results := make([]string, 0, n)

	// Start goroutines for ping and pong
	go func() {
		for i := 0; i < n; i++ {
			sender <- "ping"
		}
		close(sender)
	}()
	go func() {
		for msg := range sender {
			if msg == "ping" {
				receiver <- "pong"
			}
		}
		close(receiver)
	}()

	// Collect results
	for msg := range receiver {
		results = append(results, msg)
	}
	return results
}
