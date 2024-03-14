package authentication

import (
	"github.com/nekizz/final-project/backend/authentication/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Login(username string, password string) (*model.Account, error) {
	var account model.Account

	query := r.db.Model(&model.Account{}).Where("username = ? AND password = ?", username, password).Find(&account)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &account, nil
}

func (r *Repository) GetCustomerIdByAccountId(accountID int) (int, error) {
	var customerID int

	query := r.db.Model(&model.Account{}).Select("customers.id").Joins("INNER JOIN users ON users.account_id = accounts.id").Joins("INNER JOIN customers ON customers.user_id = users.id").Where("accounts.id = ?", accountID).First(&customerID)
	if err := query.Error; nil != err {
		return 0, err
	}

	return customerID, nil
}

func (r *Repository) UpdateOne(c *model.Account) error {
	query := r.db.Model(&model.Account{}).Where("id = ?", c.ID).Update("password", c.Password)
	if err := query.Error; nil != err {
		return err
	}

	return nil
}

func (r *Repository) CreatOne(e *model.Account) (*model.Account, error) {
	query := r.db.Model(&model.Account{}).Create(e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}

func (r *Repository) checkExistUsername(username string) (int64, error) {
	var count int64

	query := r.db.Model(&model.Account{}).Where("username = ?", username).Count(&count)
	if err := query.Error; nil != err {
		return 0, err
	}

	return count, nil
}
