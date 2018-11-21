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

type StringSlicer interface {
	Strings() []string
}

var ActionString = [3]string{"REMOVE", "ADD", "UPDATE"}

func Diff(dst StringSlicer, src StringSlicer) map[string]*Action {
	actions := make(map[string]*Action)
	m := make(map[string]bool)

	for id, s := range dst.Strings() {
		actions[s] = &Action{id, Remove}
	}

	for id, s := range src.Strings() {
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
