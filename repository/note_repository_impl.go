package repository

type NoteRepository interface {
	FindNote(id string)
	CreateNote() error
}
