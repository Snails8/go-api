package middleware

import "log"

type Logger struct {
	Default   *log.Logger
	Debug     *log.Logger
	Warning   *log.Logger
	Error     *log.Logger
	Critical  *log.Logger
}