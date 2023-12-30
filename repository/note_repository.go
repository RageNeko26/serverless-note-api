package repository

import "main/entity"

type NoteRepository interface {
	FindNote(id string) *entity.NoteEntity
	CreateNote(*entity.NoteEntity) error
}
