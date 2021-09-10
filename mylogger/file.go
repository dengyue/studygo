package mylogger

type FileLogger struct {
	level       LogLevel
	filePath    string
	fileName    string
	maxFileSize uint64
}

func NewFileLogger() *FileLogger {

	return &FileLogger{}
}
