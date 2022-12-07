package main

import (
	//alias to logrus can be set by prefixing before package name

	log "github.com/sirupsen/logrus"
)

//example-1
// func main() {
// 	logrus.Println("Hello, world!")
// }

// example-2 : setting loglevel to trace
// func main() {
// 	logrus.SetLevel(logrus.TraceLevel)
// 	logrus.Traceln("Trace Level")
// 	logrus.Debugln("Debug Level")
// 	logrus.Infoln("Info Level")
// 	logrus.Warningln("Warning Level")
// 	logrus.Errorln("Error Level")
// 	logrus.Fatalln("Fatal Level")
// 	logrus.Panicln("Panic Level")
// }

// example-3 : setting loglevel to debug
// func main() {
// 	log.SetLevel(log.DebugLevel)
// 	log.Traceln("Trace Level")
// 	log.Debugln("Debug Level")
// 	log.Infoln("Info Level")
// 	log.Warningln("Warning Level")
// 	log.Errorln("Error Level")
// 	log.Fatalln("Fatal Level")
// 	log.Panicln("Panic Level")
// }

//example-4 : writing logs to console
// func main() {
// 	// Output to stdout instead of the default stderr
// 	log.SetOutput(os.Stdout)

// 	// Only log the debug severity or above
// 	log.SetLevel(log.DebugLevel)

// 	log.Info("Info message")
// 	log.Warn("Warn message")
// 	log.Error("Error message")
// 	log.Fatal("Fatal message")
// }

//example-4 : writing logs to file

// func main() {
// 	logFile := "log.txt"
// 	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("Failed to create logfile" + logFile)
// 		panic(err)
// 	}
// 	defer f.Close()
// 	// Output to stdout instead of the default stderr
// 	log.SetOutput(f)

// 	// Only log the debug severity or above
// 	log.SetLevel(log.DebugLevel)

// 	log.Info("Info message")
// 	log.Warn("Warn message")
// 	log.Error("Error message")
// 	log.Fatal("Fatal message")
// }

//example-4 : writing logs to console and file

// func main() {
// 	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
// 	if err != nil {
// 		fmt.Println("Failed to create logfile" + "log.txt")
// 		panic(err)
// 	}
// 	defer f.Close()

// 	log := &logrus.Logger{
// 		// Log into f file handler and on os.Stdout
// 		Out:   io.MultiWriter(f, os.Stdout),
// 		Level: logrus.DebugLevel,
// 		Formatter: &easy.Formatter{
// 			TimestampFormat: "2006-01-02 15:04:05",
// 			LogFormat:       "[%lvl%]: %time% - %msg%\n",
// 		},
// 	}

// 	log.Trace("Trace message")
// 	log.Info("Info message")
// 	log.Warn("Warn message")
// 	log.Error("Error message")
// 	log.Fatal("Fatal message")
// }

//example 5 : Printing line numbers in the code where logs are getting printed

// func main() {

// 	// Output to stdout instead of the default stderr
// 	log.SetOutput(os.Stdout)

// 	// Only log the debug severity or above
// 	log.SetLevel(log.DebugLevel)

// 	// logrus show line number
// 	log.SetReportCaller(true)

// 	log.Info("Info message")
// 	log.Warn("Warn message")
// 	log.Error("Error message")
// 	log.Fatal("Fatal message")
// }

//example 6 printing the logs with a certain time format

// func main() {
// 	log.SetFormatter(&log.TextFormatter{
// 		FullTimestamp:   true,
// 		TimestampFormat: "2006-01-02T15:04:05.9999999Z07:00",
// 	})
// 	log.SetLevel(log.TraceLevel)
// 	log.Trace("Trace message")
// 	log.Info("Info message")
// 	log.Warn("Warn message")
// 	log.Error("Error message")
// 	log.Fatal("Fatal message")
// }

//example 7 Changing the log format

// func main() {
// 	log := &logrus.Logger{
// 		Out:   os.Stderr,
// 		Level: logrus.DebugLevel,
// 		Formatter: &easy.Formatter{
// 			TimestampFormat: "2006-01-02 15:04:05",
// 			LogFormat:       "[%lvl%]: %time% - %msg%\n",
// 		},
// 	}

// 	log.Trace("Trace message")
// 	log.Info("Info message")
// 	log.Warn("Warn message")
// 	log.Error("Error message")
// 	log.Fatal("Fatal message")
// }

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
	})
	log.SetLevel(log.TraceLevel)
	log.Trace("Trace message")
	log.Info("Info message")
	log.Warn("Warn message")
	log.Error("Error message")
	log.Fatal("Fatal message")
}
