package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Puskesmas struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;" json:"name"`
	Wilayah   string    `gorm:"size:255;" json:"Wilayah"`
	User      User      `json:"user"`
	UserID    uint32    `json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Puskesmas) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Wilayah = html.EscapeString(strings.TrimSpace(p.Wilayah))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Puskesmas) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Wilayah == "" {
		return errors.New("Required Wilayah")
	}
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *Puskesmas) SavePuskesmas(db *gorm.DB) (*Puskesmas, error) {
	var err error
	err = db.Debug().Model(&Puskesmas{}).Create(&p).Error
	if err != nil {
		return &Puskesmas{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Puskesmas{}, err
		}
	}
	return p, nil
}

func (p *Puskesmas) FindAllPuskesmas(db *gorm.DB) (*[]Puskesmas, error) {
	var err error
	dataPuskesmas := []Puskesmas{}
	err = db.Debug().Model(&Order{}).Limit(100).Find(&dataPuskesmas).Error
	if err != nil {
		return &[]Puskesmas{}, err
	}
	if len(dataPuskesmas) > 0 {
		for i, _ := range dataPuskesmas {
			err := db.Debug().Model(&User{}).Where("id = ?", dataPuskesmas[i].UserID).Take(&dataPuskesmas[i].User).Error
			if err != nil {
				return &[]Puskesmas{}, err
			}
		}
	}
	return &dataPuskesmas, nil
}

func (p *Puskesmas) FindPuskesmasByID(db *gorm.DB, pid uint64) (*Puskesmas, error) {
	var err error
	err = db.Debug().Model(&Puskesmas{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Puskesmas{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Puskesmas{}, err
		}
	}
	return p, nil
}
func (p *Puskesmas) FindPuskesmasByUserId(db *gorm.DB, pid uint64) (*[]Puskesmas, error) {
	var err error
	dataPuskesmas := []Puskesmas{}
	// err = db.Debug().Model(&Order{}).Limit(100).Where("user_id = ?", pid).Take(&orders).Error
	err = db.Debug().Model(&Puskesmas{}).Where("user_id = ?", pid).Find(&dataPuskesmas).Error
	if err != nil {
		return &[]Puskesmas{}, err
	}
	if len(dataPuskesmas) > 0 {
		for i, _ := range dataPuskesmas {
			err := db.Debug().Model(&User{}).Where("id = ?", dataPuskesmas[i].UserID).Take(&dataPuskesmas[i].User).Error
			if err != nil {
				return &[]Puskesmas{}, err
			}
		}
	}
	return &dataPuskesmas, nil
}

func (p *Puskesmas) UpdateAPuskesmas(db *gorm.DB) (*Puskesmas, error) {

	var err error

	err = db.Debug().Model(&Puskesmas{}).Where("id = ?", p.ID).Updates(Puskesmas{Name: p.Name, Wilayah: p.Wilayah, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Puskesmas{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Puskesmas{}, err
		}
	}
	return p, nil
}

func (p *Puskesmas) DeleteAPuskesmas(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Puskesmas{}).Where("id = ?", pid).Take(&Puskesmas{}).Delete(&Puskesmas{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Puskesmas not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
