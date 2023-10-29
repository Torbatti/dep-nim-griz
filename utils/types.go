package utils

import (
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

type App struct {
	Router *chi.Mux
	Db     *gorm.DB
}
