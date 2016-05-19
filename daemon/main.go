package main

import "os"
import "log"
import "os/signal"

var hasShutdown = false

func main() {
	daemon, err := NewDaemon(os.Getenv("ARIGATO_ROOT"))
	if err != nil {
		log.Fatalf("Failed to create daemon: %s", err)
	}

	go watch(daemon)
	defer daemon.Shutdown()

	log.Printf("Daemon is now listening on %s", daemon.server.Addr())
	err = daemon.Run()
	if err != nil {
		log.Printf("Error while running daemon: %s", err)
	}
}

func watch(daemon *Daemon) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	s := <-c

	log.Printf("Caught a signal: %s", s)
	shutdown(daemon)
}

func shutdown(daemon *Daemon) {
	err := daemon.Shutdown()
	if err != nil {
		log.Printf("Did not shutdown cleanly, error: %s", err)
	}

	if r := recover(); r != nil {
		log.Printf("Failed shutting down; caught panic: %v", r)
		panic(r)
	}
}