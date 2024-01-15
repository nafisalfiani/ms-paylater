package repository

import (
	"ms-paylater/entity"

	"gorm.io/gorm"
)

type loan struct {
	db *gorm.DB
}

type LoanInterface interface {
	Get(userId int) (entity.Loan, error)
	Create(loan entity.Loan) (entity.Loan, error)
	Update(loan entity.Loan) (entity.Loan, error)
	Delete(loanId int) error
}

// initLoan create loan repository
func initLoan(db *gorm.DB) LoanInterface {
	return &loan{
		db: db,
	}
}

func (l *loan) Get(userId int) (entity.Loan, error) {
	loan := entity.Loan{}
	if err := l.db.First(&loan, l.db.Where("user_id = ?", userId)).Error; err != nil {
		return loan, errorAlias(err)
	}

	return loan, nil
}

func (l *loan) Create(loan entity.Loan) (entity.Loan, error) {
	if err := l.db.Create(&loan).Error; err != nil {
		return loan, errorAlias(err)
	}

	return loan, nil
}

func (l *loan) Update(loan entity.Loan) (entity.Loan, error) {
	if err := l.db.Save(&loan).Error; err != nil {
		return loan, errorAlias(err)
	}

	return loan, nil
}

func (l *loan) Delete(loanId int) error {
	if err := l.db.Delete(&entity.Loan{Id: loanId}).Error; err != nil {
		return errorAlias(err)
	}

	return nil
}
