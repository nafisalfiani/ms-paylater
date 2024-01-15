package repository

import (
	"ms-paylater/entity"

	"gorm.io/gorm"
)

type transaction struct {
	db *gorm.DB
}

type TransactionInterface interface {
	List(userId int) ([]entity.Transaction, error)
	Create(transaction entity.Transaction) (entity.Transaction, error)
	Update(transaction entity.Transaction) (entity.Transaction, error)
	Delete(transactionId int) error
}

// initTransaction create Transaction repository
func initTransaction(db *gorm.DB) TransactionInterface {
	return &transaction{
		db: db,
	}
}

func (t *transaction) List(userId int) ([]entity.Transaction, error) {
	transactions := []entity.Transaction{}
	if err := t.db.Find(&transactions, t.db.Where("user_id = ?", userId)).Error; err != nil {
		return transactions, errorAlias(err)
	}

	return transactions, nil
}

func (t *transaction) Create(transaction entity.Transaction) (entity.Transaction, error) {
	if err := t.db.Transaction(func(tx *gorm.DB) error {
		loan := entity.Loan{}
		if err := tx.First(&loan, tx.Where("user_id = ?", transaction.UserId)).Error; err != nil {
			return err
		}

		transaction.BalancePrev, transaction.BalanceAfter = transaction.CalculateBalance(loan.Balance, loan.Limit)
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		loan.Balance = loan.CalculateBalance(transaction.Type, transaction.Amount)
		if err := tx.Save(&loan).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return transaction, errorAlias(err)
	}

	return transaction, nil
}

func (t *transaction) Update(transaction entity.Transaction) (entity.Transaction, error) {
	if err := t.db.Save(&transaction).Error; err != nil {
		return transaction, errorAlias(err)
	}

	return transaction, nil
}

func (t *transaction) Delete(transactionId int) error {
	if err := t.db.Delete(&entity.Transaction{Id: transactionId}).Error; err != nil {
		return errorAlias(err)
	}

	return nil
}
