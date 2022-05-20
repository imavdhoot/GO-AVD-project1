package model

import (
	"log"
)

//type Merchants struct {
//	ID        string    `gorm:"primaryKey;column:id"`
//	Name      string    `gorm:"unique;column:name"`
//	Address   string    `gorm:"column:address"`
//	CreatedAt time.Time `gorm:"column:created_at"`
//	UpdatedAt time.Time `gorm:"column:updated_at"`
//}

type NewMerchantReq struct {
	Name    string `json:"name" validate:"min=3,max=64,regexp=^[a-zA-Z0-9_ ]*$"`
	Address string `json:"address" validate:"min=6,max=256,regexp=^[a-zA-Z0-9_ \\,]*$"`
}

type UpdateMerchantReq struct {
	Address string `json:"address" validate:"min=6,max=256,regexp=^[a-zA-Z0-9_ \\,]*$"`
}

type Merchant struct {
	ID      string `gorm:"primaryKey;column:id"`
	Name    string `gorm:"unique;column:name"`
	Address string `gorm:"column:address"`
}

func CreateMerchant(id string, newMerc NewMerchantReq) (string, error) {
	merc := Merchant{ID: id, Name: newMerc.Name, Address: newMerc.Address}

	result := goDB.Select("ID", "Name", "Address").Create(&merc)

	log.Printf("[CreateMerchant] Result:: %+v", result)

	if result.Error != nil {
		log.Printf("[CreateMerchant] Error:: %s", result.Error)
	}

	return id, result.Error
}

func UpdateMerchant(id string, newMerc UpdateMerchantReq) (string, error) {
	merc := Merchant{ID: id}

	result := goDB.Model(&merc).Select("Address").Updates(Merchant{Address: newMerc.Address})

	log.Printf("[UpdateMerchant] Result:: %+v", result)
	if result.Error != nil {
		log.Printf("[UpdateMerchant] Error:: %s", result.Error)
	}

	return id, result.Error
}

func FetchMerchant(id string) (Merchant, error) {
	merc := Merchant{ID: id}

	result := goDB.First(&merc)

	log.Printf("[FetchMerchant] merc:: %+v", merc)
	log.Printf("[FetchMerchant] Result:: %+v", result)
	if result.Error != nil {
		log.Printf("[FetchMerchant] Error:: %s", result.Error)
	}

	return merc, result.Error
}

func DeleteMerchant(id string) (string, error) {
	merc := Merchant{ID: id}

	result := goDB.Delete(&merc)

	log.Printf("[DeleteMerchant] merc:: %+v", merc)
	log.Printf("[DeleteMerchant] Result:: %+v", result)
	if result.Error != nil {
		log.Printf("[DeleteMerchant] Error:: %s", result.Error)
	}

	return id, result.Error
}
