package main

// Simple static server of Angular compiled dist/project folder.
// A bit improved version.
import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	folder := ""
	port := "3000"	// Default port
	paramCount := len(os.Args)

	if paramCount == 1 {
		// Nothing to do no folder & port are given
		log.Fatal(`
Simple static server of Angular compiled dist/project folder.
Run "ng build --watch" then in another terminal
use dist/project folder as parameter for this utility.
Usage: stic STATIC_FOLDER_TO_SERVE [PORT_NUMBER]
Default port number: 3000
Examples:
stic .
stic ~/Projects/ng/ultima12/dist/ultima
stic ~/Projects/ng/ultima12/dist/ultima 4000`)

	} else if paramCount > 2 {
		// Folder & port are given so use them
		folder = os.Args[1]
		port = os.Args[2]

	} else {
		// Only folder is given use default port
		folder = os.Args[1]
	}

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		log.Fatal("Folder does not exist.")
	}

	// Serve static folder
	http.Handle("/", http.FileServer(http.Dir(folder)))
	log.Printf("\nServing static folder: %s\nListening on port: %s\nPress Ctrl-C to stop server\n", folder, port)

	// Catch the Ctrl-C and SIGTERM from kill command
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		signalType := <-ch
		signal.Stop(ch)
		log.Println("Exit command received. Exiting...")
		log.Println("Terminate signal type: ", signalType)

		//*********************************************
		// Note: here call your app Close() method to
		// properly close resources before exiting.
		//*********************************************

		os.Exit(0)
	}()

	port = ":" + port
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Http server fatal panic error: ", err)
	}
}
