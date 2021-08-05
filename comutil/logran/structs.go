package logran

type Configuration struct {
}

type Property struct {
}

type Appender struct {
}

type Root struct {
}

type Logger struct {
	Func string
}

type AppenderRef struct {
	Ref string `xml:"ref,attr"`
}
