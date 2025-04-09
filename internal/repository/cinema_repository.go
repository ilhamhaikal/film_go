package repository

import (
	"github.com/ilhamhaikal/film_go/internal/models"
	"gorm.io/gorm"
)

type CinemaRepository struct {
	db *gorm.DB
}

func NewCinemaRepository(db *gorm.DB) *CinemaRepository {
	return &CinemaRepository{db: db}
}

func (r *CinemaRepository) Create(cinema *models.Cinema) error {
	return r.db.Create(cinema).Error
}

func (r *CinemaRepository) GetAll() ([]models.Cinema, error) {
	var cinemas []models.Cinema
	err := r.db.Find(&cinemas).Error
	return cinemas, err
}

func (r *CinemaRepository) GetByID(id uint) (*models.Cinema, error) {
	var cinema models.Cinema
	err := r.db.First(&cinema, id).Error
	return &cinema, err
}

func (r *CinemaRepository) Update(cinema *models.Cinema) error {
	return r.db.Save(cinema).Error
}

func (r *CinemaRepository) Delete(id uint) error {
	return r.db.Delete(&models.Cinema{}, id).Error
}

func (r *CinemaRepository) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *CinemaRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}
