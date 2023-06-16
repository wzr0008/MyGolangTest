package Shape

type Log struct {
	msg string
}

func (l *Log) Add(str string) {
	l.msg += "\n" + str
}
func (l *Log) String() string {
	return l.msg
}

type Customer struct {
	Log
	Name string
}
