// package main


// import (
	
//     "os"
// 	// "log"
	
// 	log "github.com/sirupsen/logrus"
// )

// func main() {





// 	f, err := os.OpenFile("text.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	// logger := log.New(f, "prefix", log.LstdFlags)
// 	// logger.Println("text to append hhhhhhhhh")
// 	// logger.Println("more text to append new hhhhhhhhhh")

// 	 // Log as JSON instead of the default ASCII formatter.
// 	 log.SetFormatter(&log.JSONFormatter{})

// 	log.SetOutput(f)
// 	log.Println("more text to append new hhhhhhhhhh")
// 	// log.Output(1, "this is an event")
// }



// TRY 2

package main

import (
    "os"
    "fmt"
    log "github.com/sirupsen/logrus"
    // "LOGRUS-TEST/middlewares"
)

func init() {

    // open a file
    f, err := os.OpenFile("testlogrus.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
        fmt.Printf("error opening file: %v", err)
    }

    // don't forget to close it
    // defer f.Close()

    // Log as JSON instead of the default ASCII formatter.
    log.SetFormatter(&log.JSONFormatter{})

    // Output to stderr instead of stdout, could also be a file.
    log.SetOutput(f)

    // Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	// don't forget to close it
    // defer f.Close()
}


func main() {

    // log.WithFields(log.Fields{
    //     "Animal": "Logrus",
    // }).Info("A logrus appears")


}