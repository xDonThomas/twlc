package twlc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	Info    = "INFO"
	Success = "SUCCESS"
	Warning = "WARNING"
	Error   = "ERROR"
	Debug   = "DEBUG"
	Trace   = "TRACE"
)

type Twlc struct {
	SaveInLogFile bool
	ShowInConsole bool
	ColorMessages bool
	BGColor       bool
	FGColor       bool
	WithTime      bool
	LogDir        string
	LogFilePath   string
}

func (t *Twlc) WriteLog(messageType string, message string) {
	if t.SaveInLogFile {
		date := time.Now().Format("20060102")
		t.LogFilePath = filepath.Join(t.LogDir, "twlc_"+date+".log")
		// Create the log file if it doesn't exist
		t.createLogFile()
		// Open the log file for appending
		file, err := os.OpenFile(t.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		defer file.Close()

		logger := log.New(file, "", log.LstdFlags)
		if t.WithTime {
			logger.SetFlags(log.LstdFlags | log.Lshortfile)
		}
		logger.Printf("[%s] %s", messageType, message)
	}

	if t.ColorMessages {
		messageType, message = t.setColor(messageType, message)
	}

	if t.ShowInConsole {
		if t.WithTime {
			log.Printf("[%s] %s", messageType, message)
		} else {
			fmt.Printf("[%s] %s\n", messageType, message)
		}
	}
}

func (t *Twlc) setColor(messageType string, message string) (string, string) {
	saveMessageType := messageType
	saveMessage := message

	switch messageType {
	case Info:
		message = "\033[34m" + message + "\033[0m"
		messageType = "\033[44;44m" + messageType + "\033[0m"
	case Success:
		message = "\033[32m" + message + "\033[0m"
		messageType = "\033[42;32m" + messageType + "\033[0m"
	case Warning:
		message = "\033[33m" + message + "\033[0m"
		messageType = "\033[43;33m" + messageType + "\033[0m"
	case Error:
		message = "\033[31m" + message + "\033[0m"
		messageType = "\033[41;31m" + messageType + "\033[0m"
	case Debug:
		message = "\033[35m" + message + "\033[0m"
		messageType = "\033[45;35m" + messageType + "\033[0m"
	case Trace:
		message = "\033[36m" + message + "\033[0m"
		messageType = "\033[46;36m" + messageType + "\033[0m"
	default:
		return messageType, message
	}

	if !t.BGColor {
		messageType = saveMessageType
	}

	if !t.FGColor {
		message = saveMessage
	}

	return messageType, message
}

func (t *Twlc) WriteError(message string) {
	t.WriteLog(Error, message)
}

func (t *Twlc) WriteWarning(message string) {
	t.WriteLog(Warning, message)
}

func (t *Twlc) WriteInfo(message string) {
	t.WriteLog(Info, message)
}

func (t *Twlc) WriteSuccess(message string) {
	t.WriteLog(Success, message)
}

func (t *Twlc) WriteDebug(message string) {
	t.WriteLog(Debug, message)
}

func (t *Twlc) WriteTrace(message string) {
	t.WriteLog(Trace, message)
}

// StructToString converts a struct to a string representation.
// It uses the twlc.StructToString function.
// %+v displays the struct with field names.
// input: Animal{"Dog", 5}
// output: {Name:Dog Age:5}
// %v displays the struct without field names.
// input: Animal{"Dog", 5}
// output: {Dog 5}
// %#v displays the struct with additional details, including the type name.
// input: Animal{"Dog", 5}
// output: main.Animal{Name:"Dog", Age:5}
func (t *Twlc) StructToString(_struct interface{}, simple bool) string {
	if simple {
		return fmt.Sprintf("%+v", _struct)
	}
	return fmt.Sprintf("%#v", _struct)
}

// StructToJson converts a struct to a JSON string representation.
// It uses the json.MarshalIndent function to format the JSON output with indentation.
// The function returns the JSON string or an error message if the conversion fails.
// The JSON output is indented for better readability.
func (t *Twlc) StructToJson(_struct interface{}) (string, error) {
	jsonData, err := json.MarshalIndent(_struct, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to convert struct to JSON: %v", err)
	}
	return string(jsonData), nil
}

func (t *Twlc) createLogFile() {
	if _, err := os.Stat(t.LogFilePath); os.IsNotExist(err) {
		file, err := os.Create(t.LogFilePath)
		if err != nil {
			log.Fatalf("Failed to create log file: %v", err)
		}
		file.Close()
	}
}

func NewTwlc(saveInLogFile, ShowInConsole, colorMessages, bgColor, fgColor, withTime bool, logDir string) *Twlc {
	createLogDir(logDir)

	return &Twlc{
		SaveInLogFile: saveInLogFile,
		ShowInConsole: ShowInConsole,
		WithTime:      withTime,
		ColorMessages: colorMessages,
		BGColor:       bgColor,
		FGColor:       fgColor,
		LogDir:        logDir,
	}
}

func DefaultTwlc() *Twlc {
	exeDir, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exeDir = filepath.Dir(exeDir)

	logDir := exeDir + "/logs/"

	createLogDir(logDir)

	return &Twlc{true, true, true, true, true, true, logDir, ""}
}

func createLogDir(logDir string) {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}
}
