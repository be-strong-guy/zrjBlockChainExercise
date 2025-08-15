package task1

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
*
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
transactions 表（包含字段 id 主键， from_account_id 转出账户ID，
to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
如果余额不足，则回滚事务。
*/
type Accounts struct {
	gorm.Model
	ID      uint `gorm:"primary_key;auto_increment" json:"id"`
	Balance int  `gorm:"default:0" json:"balance"`
}
type Transactions struct {
	gorm.Model
	Accounts      Accounts `gorm:"foreignkey:ID" json:"accounts"`
	ID            uint     `gorm:"primary_key;auto_increment" json:"id"`
	Amount        int      `gorm:"default:0" json:"amount"`
	FromAccountID uint
	ToAccountID   uint
}

func TransferAccounts(db *gorm.DB) {
	account1 := Accounts{Balance: 1000}
	account2 := Accounts{Balance: 1000}
	err := db.AutoMigrate(&Accounts{})
	err = db.AutoMigrate(&Transactions{})
	if err != nil {
		panic(err)
	}
	db.Create(&account1)
	db.Create(&account2)
	//开始转账
	accountAID := uint(1)
	accountBID := uint(2)
	amount := 100
	err = db.Transaction(func(tx *gorm.DB) error {
		var accountA, accountB Accounts
		// 1. 锁定账户1
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&accountA, accountAID).Error; err != nil {
			return err
		}
		// 2. 检查余额
		if accountA.Balance < amount {
			return errors.New("余额不足")
		}
		// 3. 锁定账户B
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&accountB, accountBID).Error; err != nil {
			return err
		}

		// 4. 更新账户余额
		if err := tx.Model(&Accounts{}).
			Where("id = ?", accountAID).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}

		if err := tx.Model(&Accounts{}).
			Where("id = ?", accountBID).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}
		// 5. 记录到交易表Transaction
		txRecord := Transactions{
			FromAccountID: accountAID,
			ToAccountID:   accountBID,
			Amount:        amount,
		}
		tx.Create(&txRecord)
		return nil
	})
	if err != nil {
		fmt.Println("转账失败", err)
		panic(err)
	} else {
		fmt.Println("转账成功")
	}
}
