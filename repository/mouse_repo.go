package repository

import (
	"fiber3-app/backend/database"
	"fiber3-app/backend/domain"

	"gorm.io/gorm"
)

type MouseRepository struct {
	db *gorm.DB
}

func NewMouseRepository() *MouseRepository {
	return &MouseRepository{db: database.DB}
}

func (r *MouseRepository) GetAll() []domain.Mouse {
	var mice []domain.Mouse

	r.db.Raw(`
		SELECT UUID_TO_CHAR(ID) ID, AVATAR, "NAME", AMOUNT, VENDOR, STATUS, RATING
		FROM MICE`).Scan(&mice)

	return mice
}

func (r *MouseRepository) GetByID(id string) *domain.Mouse {
	var mouse domain.Mouse

	r.db.Raw(`
		SELECT UUID_TO_CHAR(ID) ID, AVATAR, "NAME", AMOUNT, VENDOR, STATUS, RATING
		FROM MICE
		WHERE ID = CHAR_TO_UUID(?)`, id).Scan(&mouse)

	return &mouse
}

func (r *MouseRepository) Create(m *domain.Mouse) error {
	qry := `
		INSERT INTO MICE (ID, AVATAR, "NAME", AMOUNT, VENDOR, STATUS, RATING)
		VALUES (CHAR_TO_UUID(?), ?, ?, ?, ?, ?, ?)`
	result := r.db.Exec(qry, m.ID, m.Avatar, m.Name, m.Amount, m.Vendor, m.Status, m.Rating)

	return result.Error
}

func (r *MouseRepository) Update(m *domain.Mouse) error {
	qry := `
		UPDATE PROFILE SET 
			AVATAR = ?, 
			"NAME" = ?, 
			AMOUNT = ?, 
			VENDOR = ?, 
			STATUS = ?, 
			RATING = ?
		WHERE ID = CHAR_TO_UUID(?)`
	result := r.db.Raw(qry, m.Avatar, m.Name, m.Amount, m.Vendor, m.Status, m.Rating, m.ID).Row()

	return result.Err()

}

func (r *MouseRepository) Delete(id string) string {
	qry := `DELETE FROM MICE WHERE ID = CHAR_TO_UUID(?)`

	result := r.db.Raw(qry, id).Row()

	return result.Err().Error()
}
