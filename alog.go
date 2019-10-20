package alog

import (
	"io"
)

type alog struct {
}

// Print creates informational logs based on the inputs
func (l *alog) Print(v ...interface{}) {

}

// Println prints the data coming in on individual lines
func (l *alog) Println(v ...interface{}) {

}

// Printf creates an informational log using the format and values
func (l *alog) Printf(format string, v ...interface{}) {

}

// Warn creates a warning log using the error passed in along with the
// values passed in
func (l *alog) Warn(err error, v ...interface{}) {

}

// Warnln creates a warning log using the error and values passed in.
// Each error and value is printed on a different line
func (l *alog) Warnln(err error, v ...interface{}) {

}

// Warnf creates a warning log using the error passed in, along with the string
// formatting and values
func (l *alog) Warnf(err error, format string, v ...interface{}) {

}

// Error creates an error log using the error and other values passed in
func (l *alog) Error(err error, v ...interface{}) {

}

// Error creates error logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Errorln(err error, v ...interface{}) {

}

// Errorf creates an error log using the error passed in, along with the string
// formatting and values
func (l *alog) Errorf(err error, format string, v ...interface{}) {

}

// Crit creates critical logs using the error and other values passed in
func (l *alog) Crit(err error, v ...interface{}) {

}

// Critln creates critical logs using the error and other values passed in.
// Each error and value is printed on a different line
func (l *alog) Critln(err error, v ...interface{}) {

}

// Critf creates a critical log using the error passed in, along with the string
// formatting and values
func (l *alog) Critf(err error, format string, v ...interface{}) {

}

// Fatal creates a fatal log using the error and values passed into the method
// After logging the fatal log the Fatal method throws a panic to crash the application
func (l *alog) Fatal(err error, v ...interface{}) {

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// Fatalln creates fatal logs using the error and other values passed in.
// Each error and value is printed on a different line
// After logging the fatal log the Fatalln method throws a panic to crash the application
func (l *alog) Fatalln(err error, v ...interface{}) {

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// Fatalf creates an error log using the error passed in, along with the string
// formatting and values
// After logging the fatal log the Fatalf method throws a panic to crash the application
func (l *alog) Fatalf(err error, format string, v ...interface{}) {

	// TODO: Update panic to include information about the fatal, as well as stack trace information
	panic(err)
}

// AddOutput adds an additional logging source to the logger which
// will be added to the different logging outputs for this logger
func (l *alog) AddOutput(out io.Writer) {

}