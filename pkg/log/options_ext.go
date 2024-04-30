package log

import "strings"

func LevelToString(level int32) (string, bool) {
	l := Level(level)
	s, ok := levelToString[l]
	return s, ok
}

func LevelToStringExt(level Level) (string, bool) {
	s, ok := levelToString[level]
	return s, ok
}

func FormatLevelString(level string) string {
	return strings.ToLower(level)
}

func LevelStringCheck(level string) bool {
	_, ok := stringToLevel[FormatLevelString(level)]
	return ok
}

func StringToLevel(level string) Level {
	l, _ := stringToLevel[FormatLevelString(level)]
	return l
}
