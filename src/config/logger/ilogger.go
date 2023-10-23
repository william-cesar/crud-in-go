package logger

type Options string
type TprefixOption map[Options]string

const (
	INFO  Options = "INFO"
	ERROR Options = "ERROR"
	WARN  Options = "WARN"
)

var prefix = TprefixOption{
	INFO:  "[INFO] - ",
	ERROR: "[ERROR] - ",
	WARN:  "[WARN] - ",
}

type TLogger struct {
	Prefix  Options
	Journey string
	Message string
	Error   error
}
