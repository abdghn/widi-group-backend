package seed

import (
	"log"

	"product-order-be/api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
		Role:     "User",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
		Role:     "Admin",
	},
}

var posts = []models.Order{
	models.Order{
		Name:     "Test1",
		Category: "Test1",
		Type:     "Test1",
		Total:    "20",
		Price:    "5000",
	},
	models.Order{
		Name:     "Test2",
		Category: "Test2",
		Type:     "Test2",
		Total:    "30",
		Price:    "600000",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Order{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Order{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Order{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].UserID = users[i].ID

		err = db.Debug().Model(&models.Order{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
