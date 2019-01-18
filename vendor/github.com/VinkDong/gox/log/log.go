package log

import (
	"fmt"
	"log"
	"os"
	"sync"
	"io"
	"time"
	"runtime"
)

const ClrEnd = "\x1b[0m"
const ClrDEBUG = "\x1b[37;2m"
const ClrSuccess = "\x1b[32;2m"

var debug = false
var std = New(os.Stdout, "", log.LstdFlags)

type Logger struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	prefix string     // prefix to write at beginning of each line
	flag   int        // properties
	out    io.Writer  // destination for output
	buf    []byte     // for accumulating text to write
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{out: out, prefix: prefix, flag: flag}
}

// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

// formatHeader writes log header to buf in following order:
//   * l.prefix (if it's not blank),
//   * date and/or time (if corresponding flags are provided),
//   * file and line number (if corresponding flags are provided).
func (l *Logger) formatHeader(buf *[]byte, t time.Time, file string, line int) {
	*buf = append(*buf, l.prefix...)
	if l.flag&(log.Ldate|log.Ltime|log.Lmicroseconds) != 0 {
		if l.flag&log.LUTC != 0 {
			t = t.UTC()
		}
		if l.flag&log.Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if l.flag&(log.Ltime|log.Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if l.flag&log.Lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}
	if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
		if l.flag&log.Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, ": "...)
	}
}

// Output writes the output for a logging event. The string s contains
// the text to print after the prefix specified by the flags of the
// Logger. A newline is appended if the last character of s is not
// already a newline. Calldepth is used to recover the PC and is
// provided for generality, although at the moment on all pre-defined
// paths it will be 2.
func (l *Logger) Output(calldepth int, s string, color string) error {
	now := time.Now() // get this early.
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
		// Release lock while getting caller info - it's expensive.
		l.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		l.mu.Lock()
	}
	l.buf = l.buf[:0]
	if len(color) > 0 {
		l.buf = []byte(color)
	}
	l.formatHeader(&l.buf, now, file, line)
	l.buf = append(l.buf, s...)
	if len(color) > 0 {
		l.buf = append(l.buf, ClrEnd...)
	}
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, err := l.out.Write(l.buf)
	return err
}

func EnableDebug() {
	debug = true
}

func Info(l interface{}) {
	info := fmt.Sprintf("[%s] %v", "INFO", l)
	std.Output(2, info, "")
}

func Infof(l string, a ...interface{}) {
	tmp := fmt.Sprintf(l, a...)
	info := fmt.Sprintf("[%s] %s", "INFO", tmp)
	std.Output(2, info, "")
}

func Error(l interface{}) {
	err := fmt.Sprintf("[%s] %v", "Error", l)
	log.Println(err)
}

func Errorf(format string, a ...interface{}) {
	tmp := fmt.Sprintf(format, a...)
	err := fmt.Sprintf("[%s] %s", "Error", tmp)
	log.Println(err)
}

func Debug(l interface{}) {
	if debug {
		debug := fmt.Sprintf("[%s] %v", "DEBUG", l)
		std.Output(2, debug, ClrDEBUG)
	}
}

func Debugf(format string, a ...interface{}) {
	if debug {
		tmp := fmt.Sprintf(format, a...)
		debug := fmt.Sprintf("[%s] %s", "DEBUG", tmp)
		std.Output(2, debug, ClrDEBUG)
	}
}

func Lock()  {
	std.mu.Lock()
}

func Unlock()  {
	std.mu.Unlock()
}

func Success(l interface{}) {
	success := fmt.Sprintf("[INFO] %v", l)
	std.Output(2, success, ClrSuccess)
}

func Successf(format string, a ...interface{}) {
	tmp := fmt.Sprintf(format, a...)
	success := fmt.Sprintf("[INFO] %v", tmp)
	std.Output(2, success, ClrSuccess)
}