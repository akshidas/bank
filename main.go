package main

func main() {
	server := newApiServer(":5000")
	server.Run()

}
