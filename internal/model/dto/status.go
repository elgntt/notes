package dto

type Enum interface {
	IsValid() bool
}

func (t TaskStatus) IsValid() bool {
	switch t {
	case "backlog", "open", "progress", "review", "completed":
		return true
	}

	return false
}

type TaskStatus string
