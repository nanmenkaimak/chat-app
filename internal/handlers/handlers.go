package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/nanmenkaimak/chat-app/internal/repository"
	"github.com/nanmenkaimak/chat-app/internal/repository/dbrepo"
)

// Repo is repository used by handlers
var Repo *Repository

// Repository is repository type
type Repository struct {
	DB repository.DatabaseRepo
}

// NewRepo creates new repository
func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db),
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
