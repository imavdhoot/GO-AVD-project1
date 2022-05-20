package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/imavdhoot/GO-AVD-project1/constant"
	"github.com/imavdhoot/GO-AVD-project1/model"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"net/mail"
	"strconv"
)

func AddMember(ctx *gin.Context) {

	var member model.NewMemberReq

	if err := ctx.ShouldBindJSON(&member); err != nil {
		log.Println("[AddMember] error from json reading:: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  err.Error(),
		})
		return
	}
	log.Printf("[AddMember] New member:: %+v\n", member)

	if validateErr := validator.Validate(member); validateErr != nil {
		log.Println("[AddMember] validateErr :: ", validateErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  validateErr.Error(),
		})
		return
	}

	if _, emailErr := mail.ParseAddress(member.Email); emailErr != nil {
		log.Printf("[AddMember] Error validating email address:: %s error:: %s", member.Email, emailErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  constant.ErrInvalidEmailAddress.Error(),
		})
		return
	}

	_, GetMerchantErr := model.FetchMerchant(member.MerchantId)
	if GetMerchantErr != nil {
		log.Println("[AddMember] error in getting merchant:: ", GetMerchantErr)
		var errMessage = GetMerchantErr
		if GetMerchantErr.Error() == constant.ErrRecNotFound {
			errMessage = constant.ErrMerchantNotFound
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  errMessage.Error(),
		})
		return
	}

	createRes, createErr := model.CreateMember(member)
	if createErr != nil {
		log.Println("[AddMember] error in creating new Member:: ", createErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  createErr.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     constant.StatusOK,
		"memberId":   createRes.ID,
		"name":       member.Name,
		"email":      member.Email,
		"merchantId": member.MerchantId,
		"message":    "Member added successfully",
	})
}

func UpdateMember(ctx *gin.Context) {

	var Member model.UpdateMemberReq

	memberId, memberIdErr := strconv.Atoi(ctx.Param("id"))
	if memberIdErr != nil {
		log.Println("[UpdateMember] invalid/empty memberId for update:: ")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  constant.ErrMemberIdInvalid.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&Member); err != nil {
		log.Println("[UpdateMember] error from json reading:: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  err.Error(),
		})
		return
	}
	log.Printf("[UpdateMember] memberId:: %s", memberId)
	log.Printf("[UpdateMember] Member:: %+v\n", Member)

	if validateErr := validator.Validate(Member); validateErr != nil {
		log.Println("[UpdateMember] validateErr :: ", validateErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  validateErr.Error(),
		})
		return
	}

	_, getErr := model.FetchMember(memberId)
	if getErr != nil {
		log.Println("[UpdateMember] error in getting Member:: ", getErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  getErr.Error(),
		})
		return
	}

	_, updateErr := model.UpdateMember(memberId, Member)
	if updateErr != nil {
		log.Println("[UpdateMember] error in updating Member:: ", updateErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  updateErr.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   constant.StatusOK,
		"memberId": memberId,
		"message":  "Member updated successfully",
	})
}

func GetMember(ctx *gin.Context) {

	memberId, memberIdErr := strconv.Atoi(ctx.Param("id"))
	if memberIdErr != nil {
		log.Println("[GetMember] invalid/empty memberId for update:: ")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  constant.ErrMemberIdInvalid.Error(),
		})
		return
	}

	log.Printf("[GetMember] memberId:: %s", memberId)

	getRes, getErr := model.FetchMember(memberId)
	if getErr != nil {
		log.Println("[GetMember] error in getting Member:: ", getErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  getErr.Error(),
		})
		return
	}

	log.Printf("[GetMember] Member:: %+v", getRes)

	ctx.JSON(http.StatusOK, gin.H{
		"status":     constant.StatusOK,
		"memberId":   getRes.ID,
		"name":       getRes.Name,
		"email":      getRes.Email,
		"merchantId": getRes.MerchantId,
		"message":    "Member fetched successfully",
	})
}

func DeleteMember(ctx *gin.Context) {

	memberId, memberIdErr := strconv.Atoi(ctx.Param("id"))
	if memberIdErr != nil {
		log.Println("[DeleteMember] invalid/empty memberId for update:: ")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  constant.ErrMemberIdInvalid.Error(),
		})
		return
	}

	log.Printf("[DeleteMember] memberId:: %s", memberId)

	_, getErr := model.FetchMember(memberId)
	if getErr != nil {
		log.Println("[DeleteMember] error in getting Member:: ", getErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  getErr.Error(),
		})
		return
	}

	DeleteRes, DeleteErr := model.DeleteMember(memberId)
	if DeleteErr != nil {
		log.Println("[DeleteMember] error in deleting Member:: ", DeleteErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  DeleteErr.Error(),
		})
		return
	}

	log.Printf("[DeleteMember] memberId:: %+v", DeleteRes)

	ctx.JSON(http.StatusOK, gin.H{
		"status":   constant.StatusOK,
		"memberId": DeleteRes,
		"message":  "Member deleted successfully",
	})
}
