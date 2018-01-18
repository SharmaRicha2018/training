package filewriter

import (
	"fmt"
	"github.com/jabong/floRest/src/common/logger/formatter"
	"github.com/jabong/floRest/src/common/logger/message"
	"log"
	"os"
	"time"
)

//FileWriter is a file logger structure
type FileWriter struct {
	//filename is the log file name with absolute path
	filename string

	//file is the file handle for log filename
	file *os.File

	// formatter
	myFormat formatter.FormatInterface
}

//GetNewObj returns a file logger with log file name fname, having configuration
//specified in conf and allowedLogLevel specifies the log level that are actually to
//be logged
func GetNewObj(fname string) (*FileWriter, error) {
	obj := new(FileWriter)
	obj.filename = fname

	err := obj.setLogFileHandle()
	if err != nil {
		return nil, err
	}
	return obj, nil
}

//SetOutput sets the file handle where to write the log
func (fw *FileWriter) setOutput() {
	logName := fw.getLogFileExt()
	log.SetFlags(0)
	if _, err := os.Stat(logName); err == nil {
		log.SetOutput(fw.file)
		return
	}
	if fw.file != nil {
		fw.file.Close()
	}
	err := fw.setLogFileHandle()
	if err != nil {
		fmt.Printf("\nError In Settting log file handle %s: %+v\n", fw.filename, err)
		return
	}
	log.SetOutput(fw.file)
}

//setLogFileHandle opens a file and assigns the file handle to file in fileLogger
func (fw *FileWriter) setLogFileHandle() error {
	logName := fw.getLogFileExt()
	f, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("\nError In Opening log file %s: %+v\n", fw.filename, err)
		return err
	}
	fw.file = f
	return nil
}

//getLogFileExt returns the log file ext name after appending
//today's date and ".log" to the filename
func (fw *FileWriter) getLogFileExt() string {
	t := time.Now().Local()
	tf := t.Format("2006-01-02")
	return fw.filename + tf + ".log"
}

// Write write message to file
func (fw *FileWriter) Write(msg *message.LogMsg) {
	str := fw.myFormat.GetFormattedLog(msg)
	fw.setOutput()
	log.Println(str)
}

// SetFormatter get formatted object
func (fw *FileWriter) SetFormatter(format formatter.FormatInterface) {
	fw.myFormat = format
}
