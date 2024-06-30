package cli

type Flag struct {
	Name string // like `-tag`
	Help string // help message
	MinValues int // minimum values count
	MaxValues int // maximum values count. if 0, this flag peforms bool flag.
	DefaultValues []string
	Workdir string // default value is `.`
}

func (f *Flag) Has() bool {
	for _, a := range Args {
		if a == f.Name {
			return true
		}
	}
	return false
}