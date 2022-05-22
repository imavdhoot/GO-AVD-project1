package model

import (
	"github.com/imavdhoot/GO-AVD-project1/constant"
	"log"
)

type NewMemberReq struct {
	Name       string `json:"name" validate:"min=3,max=64,regexp=^[a-zA-Z0-9_ ]*$"`
	Email      string `json:"email" validate:"min=6,max=256,regexp=^[a-zA-Z0-9\\-._@]*$"`
	MerchantId string `json:"merchantId" validate:"min=3,max=64,regexp=^[a-zA-Z0-9_-]*$"`
}

type UpdateMemberReq struct {
	Email string `json:"email" validate:"min=6,max=256,regexp=^[a-zA-Z0-9\\-._@]*$"`
}

type Member struct {
	ID         int    `json:"id" gorm:"primaryKey;column:id"`
	Name       string `json:"name" gorm:"column:name"`
	Email      string `json:"email" gorm:"unique;column:email"`
	MerchantId string `json:"merchantId" gorm:"column:merchant_id"`
}

func CreateMember(newMemb NewMemberReq) (Member, error) {
	memb := Member{Name: newMemb.Name, Email: newMemb.Email, MerchantId: newMemb.MerchantId}

	result := goDB.Select("Name", "Email", "MerchantId").Create(&memb)

	log.Printf("[CreateMember] member:: %+v", memb)
	log.Printf("[CreateMember] Result:: %+v", result)

	if result.Error != nil {
		log.Printf("[CreateMember] Error:: %s", result.Error)
	}

	return memb, result.Error
}

func UpdateMember(id int, newMemb UpdateMemberReq) (int, error) {
	merc := Member{ID: id}

	result := goDB.Model(&merc).Select("Email").Updates(Member{Email: newMemb.Email})

	log.Printf("[UpdateMember] Result:: %+v", result)
	if result.Error != nil {
		log.Printf("[UpdateMember] Error:: %s", result.Error)
	}

	return id, result.Error
}

func FetchMember(id int) (Member, error) {
	memb := Member{ID: id}

	result := goDB.First(&memb)

	log.Printf("[FetchMember] Member:: %+v", memb)
	log.Printf("[FetchMember] Result:: %+v", result)
	if result.Error != nil {
		log.Printf("[FetchMember] Error:: %s", result.Error)
	}

	return memb, result.Error
}

func DeleteMember(id int) (int, error) {
	merc := Member{ID: id}

	result := goDB.Delete(&merc)

	log.Printf("[DeleteMember] merc:: %+v", merc)
	log.Printf("[DeleteMember] Result:: %+v", result)
	if result.Error != nil {
		log.Printf("[DeleteMember] Error:: %s", result.Error)
	}

	return id, result.Error
}

func MemberListByMerchant(merchantId string, pageNo int) ([]Member, error) {
	var membs []Member

	offset := (pageNo - 1) * constant.PageSize

	log.Printf("[MemberListByMerchant] offset:: %d", offset)

	result := goDB.Raw("SELECT * FROM members WHERE merchant_id = ? LIMIT ?, ?",
		merchantId, offset, constant.PageSize).Scan(&membs)

	log.Printf("[MemberListByMerchant] Member:: %+v", membs)
	log.Printf("[MemberListByMerchant] Result:: %+v", result)
	if result.Error != nil {
		log.Printf("[MemberListByMerchant] Error:: %s", result.Error)
	}

	return membs, result.Error
}
