package config

import "fmt"

type LogArgs struct {
	Dir  string
	FileName string
	empty bool
}

func EmptyLog() *LogArgs {
	return &LogArgs{empty: true}
}

func (l *LogArgs) Empty() (bool, error) {
	if l.empty {
		return true, fmt.Errorf("log is empty")
	}
	return false, nil
}

func (l *LogArgs) ExistsDir() bool {
	return len(l.Dir) > 0
}

func (l *LogArgs) ExistsFile() bool {
	return len(l.FileName) > 0
}
