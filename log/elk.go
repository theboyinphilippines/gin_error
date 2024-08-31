package logx

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// 通过TCP直接将日志写入logstash
type ElkLogger struct {
	Logger *LogStash
}

type LogStash struct {
	hostname string
	port     int
	Conn     *net.TCPConn
	TimeOut  int
}

func newLogStash(hostname string, port int, timeout int) *LogStash {
	return &LogStash{
		hostname: hostname,
		port:     port,
		Conn:     nil,
		TimeOut:  timeout,
	}
}

func (l *LogStash) Connect() (*net.TCPConn, error) {
	var conn *net.TCPConn
	service := fmt.Sprintf("%s:%d", l.hostname, l.port)
	addr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		return nil, err
	}
	conn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}
	if conn != nil {
		l.Conn = conn
		l.Conn.SetKeepAlive(true)
		l.Conn.SetKeepAlivePeriod(time.Duration(5) * time.Second)
		l.setTimeouts()
	}
	return conn, nil
}

func (l *LogStash) setTimeouts() {
	deadline := time.Now().Add(time.Duration(l.TimeOut) * time.Millisecond)
	_ = l.Conn.SetDeadline(deadline)
	_ = l.Conn.SetReadDeadline(deadline)
	_ = l.Conn.SetWriteDeadline(deadline)
}

// Write. message: json
func (l *LogStash) Write(message string) (err error) {
	message = fmt.Sprintf("%s\n", message)
	if l.Conn != nil {
		_, err = l.Conn.Write([]byte(message))
		if err != nil {
			l.Connect()
			return
		} else {
			l.setTimeouts()
			return nil
		}
	}
	return
}

func NewElkLogger(hostname string, port int, timeout int) *ElkLogger {
	return &ElkLogger{Logger: newLogStash(hostname, port, timeout)}
}

var Logger *ElkLogger

type ElkMsg struct {
	Service string
	Level   string
	Msg     string
}

func (log *ElkLogger) Errorf(format string, args ...interface{}) {
	msg := &ElkMsg{
		Service: "ceph-backend",
		Level:   "error",
		Msg:     fmt.Sprintf(format, args...),
	}
	var (
		bytes []byte
		err   error
	)
	if bytes, err = json.Marshal(msg); err != nil {
		fmt.Println("elk msg marshal error")
		return
	}
	log.Logger.Write(string(bytes))
}

func (log *ElkLogger) Infof(format string, args ...interface{}) {
	msg := &ElkMsg{
		Service: "ceph-backend",
		Level:   "info",
		Msg:     fmt.Sprintf(format, args...),
	}
	var (
		bytes []byte
		err   error
	)
	if bytes, err = json.Marshal(msg); err != nil {
		fmt.Println("elk msg marshal error")
		return
	}
	log.Logger.Write(string(bytes))
}

func (log *ElkLogger) Warnf(format string, args ...interface{}) {
	msg := &ElkMsg{
		Service: "ceph-backend",
		Level:   "warn",
		Msg:     fmt.Sprintf(format, args...),
	}
	var (
		bytes []byte
		err   error
	)
	if bytes, err = json.Marshal(msg); err != nil {
		fmt.Println("elk msg marshal error")
		return
	}
	log.Logger.Write(string(bytes))
}
