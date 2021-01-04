package compare

type Diff interface {
	Diffs() []string
	Append(...string)
}
