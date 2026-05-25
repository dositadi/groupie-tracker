package jsonlog

type Color struct {}

func (c *Color) Get(color string) string {
	ansiiColor := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
		"orange": "\033[38;5;166m",
		"reset":  "\033[0m",
	}

	if val, ok := ansiiColor[color]; ok {
		return val
	}
	return ""
}
