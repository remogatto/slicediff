package slicediff

const (
	Remove = iota + 1
	Add
	Update
)

type Action struct {
	Id   int
	Type int
}

var ActionString = [3]string{"REMOVE", "ADD", "UPDATE"}

func Diff(dst func() []string, src func() []string) map[string]*Action {
	actions := make(map[string]*Action)
	m := make(map[string]bool)

	for id, s := range dst() {
		actions[s] = &Action{id, Remove}
	}

	for id, s := range src() {
		if _, ok := m[s]; !ok {
			m[s] = true
			_, ok := actions[s]
			if !ok {
				actions[s] = &Action{id, Add}
				continue
			}
			actions[s] = &Action{id, Update}
		}
	}

	return actions
}
