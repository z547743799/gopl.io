package main

type client chan<- string

var (
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	messages <- "---"
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		}
	}
}
func main() {

	go broadcaster()
	for input.Scan() {
		messages <- input.Text()
	}

}
