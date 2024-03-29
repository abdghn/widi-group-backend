package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:255;" json:"name"`
	Category  string `gorm:"size:255;" json:"category"`
	Type      string `gorm:"size:255;" json:"type"`
	Total     string `gorm:"size:255;" json:"total"`
	Price     string `gorm:"size:255;" json:"price"`
	Status    string `gorm:"size:255;" json:"status"`
	Image     string `gorm:"size:255;" json:"image"`
	Puskesmas string `gorm:"size:255;" json:"puskesmas"`
	User      User   `json:"user"`
	UserID    uint32 `json:"user_id"`
	// Puskesmas   User      `json:"puskesmas"`
	// PuskesmasID uint32    `json:"puskesmas_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Order) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Category = html.EscapeString(strings.TrimSpace(p.Category))
	p.Type = html.EscapeString(strings.TrimSpace(p.Type))
	p.Total = html.EscapeString(strings.TrimSpace(p.Total))
	p.Price = html.EscapeString(strings.TrimSpace(p.Price))
	p.Status = html.EscapeString(strings.TrimSpace(p.Status))
	p.Image = html.EscapeString(strings.TrimSpace(p.Image))
	p.Puskesmas = html.EscapeString(strings.TrimSpace(p.Puskesmas))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Order) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Category == "" {
		return errors.New("Required Category")
	}
	if p.Type == "" {
		return errors.New("Required Type")
	}
	if p.Total == "" {
		return errors.New("Required Total")
	}
	if p.Price == "" {
		return errors.New("Required Price")
	}
	if p.Puskesmas == "" {
		return errors.New("Required Puskesmas")
	}
	// if p.Image == "" {
	// 	return errors.New("Required Image")
	// }
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *Order) SaveOrder(db *gorm.DB) (*Order, error) {
	var err error
	err = db.Debug().Model(&Order{}).Create(&p).Error
	if err != nil {
		return &Order{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Order{}, err
		}
		// err = db.Debug().Model(&Puskesmas{}).Where("id = ?", p.PuskesmasID).Take(&p.Puskesmas).Error
		// if err != nil {
		// 	return &Order{}, err
		// }
	}
	return p, nil
}

func (p *Order) FindAllOrders(db *gorm.DB) (*[]Order, error) {
	var err error
	orders := []Order{}
	err = db.Debug().Model(&Order{}).Limit(100).Find(&orders).Error
	if err != nil {
		return &[]Order{}, err
	}
	if len(orders) > 0 {
		for i, _ := range orders {
			err := db.Debug().Model(&User{}).Where("id = ?", orders[i].UserID).Take(&orders[i].User).Error
			if err != nil {
				return &[]Order{}, err
			}
		}
	}
	return &orders, nil
}

func (p *Order) FindorderByID(db *gorm.DB, pid uint64) (*Order, error) {
	var err error
	err = db.Debug().Model(&Order{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Order{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Order{}, err
		}
	}
	return p, nil
}
func (p *Order) FindOrdersByUserId(db *gorm.DB, pid uint64) (*[]Order, error) {
	var err error
	orders := []Order{}
	// err = db.Debug().Model(&Order{}).Limit(100).Where("user_id = ?", pid).Take(&orders).Error
	err = db.Debug().Model(&Order{}).Where("user_id = ?", pid).Find(&orders).Error
	if err != nil {
		return &[]Order{}, err
	}
	if len(orders) > 0 {
		for i, _ := range orders {
			err := db.Debug().Model(&User{}).Where("id = ?", orders[i].UserID).Take(&orders[i].User).Error
			if err != nil {
				return &[]Order{}, err
			}
		}
	}
	return &orders, nil
}

func (p *Order) UpdateAOrder(db *gorm.DB) (*Order, error) {

	var err error

	err = db.Debug().Model(&Order{}).Where("id = ?", p.ID).Updates(Order{Name: p.Name, Category: p.Category, Type: p.Type, Total: p.Total, Status: p.Status, Price: p.Price, Image: p.Image, Puskesmas: p.Puskesmas, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Order{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Order{}, err
		}
	}
	return p, nil
}

func (p *Order) DeleteAOrder(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Order{}).Where("id = ?", pid).Take(&Order{}).Delete(&Order{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Order not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
