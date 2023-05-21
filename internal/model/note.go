package model

type Note struct {
	Id   int
	Text string
}

type NoteInfo struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}
