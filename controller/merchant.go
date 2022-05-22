package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/imavdhoot/GO-AVD-project1/constant"
	"github.com/imavdhoot/GO-AVD-project1/model"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

func AddMerchant(ctx *gin.Context) {

	var merchant model.NewMerchantReq

	if err := ctx.ShouldBindJSON(&merchant); err != nil {
		log.Println("[AddMerchant] error from json reading:: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  err.Error(),
		})
		return
	}
	log.Printf("[AddMerchant] New merchant:: %+v\n", merchant)

	if validateErr := validator.Validate(merchant); validateErr != nil {
		log.Println("[AddMerchant] validateErr :: ", validateErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  validateErr.Error(),
		})
		return
	}

	merchantId := uuid.New().String()
	log.Println("[AddMerchant] Generated merchantId:: ", merchantId)

	_, createErr := model.CreateMerchant(merchantId, merchant)
	if createErr != nil {
		log.Println("[AddMerchant] error in creating new merchant:: ", createErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  createErr.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     constant.StatusOK,
		"merchantId": merchantId,
		"name":       merchant.Name,
		"address":    merchant.Address,
		"message":    "merchant added successfully",
	})
}

func UpdateMerchant(ctx *gin.Context) {

	var merchant model.UpdateMerchantReq

	merchantId := ctx.Param("id")
	if merchantId == "" {
		log.Println("[UpdateMerchant] invalid/empty merchantId for update:: ", merchantId)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  constant.ErrMerchantIdInvalid.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&merchant); err != nil {
		log.Println("[UpdateMerchant] error from json reading:: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  err.Error(),
		})
		return
	}
	log.Printf("[UpdateMerchant] merchantId:: %s", merchantId)
	log.Printf("[UpdateMerchant] merchant:: %+v\n", merchant)

	if validateErr := validator.Validate(merchant); validateErr != nil {
		log.Println("[UpdateMerchant] validateErr :: ", validateErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  validateErr.Error(),
		})
		return
	}

	_, getErr := model.FetchMerchant(merchantId)
	if getErr != nil {
		log.Println("[UpdateMerchant] error in getting merchant:: ", getErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  getErr.Error(),
		})
		return
	}

	_, updateErr := model.UpdateMerchant(merchantId, merchant)
	if updateErr != nil {
		log.Println("[UpdateMerchant] error in updating merchant:: ", updateErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  updateErr.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     constant.StatusOK,
		"merchantId": merchantId,
		"message":    "merchant updated successfully",
	})
}

func GetMerchant(ctx *gin.Context) {

	merchantId := ctx.Param("id")
	if merchantId == "" {
		log.Println("[GetMerchant] invalid/empty merchantId for update:: ")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  constant.ErrMerchantIdInvalid.Error(),
		})
		return
	}

	log.Printf("[GetMerchant] merchantId:: %s", merchantId)

	getRes, getErr := model.FetchMerchant(merchantId)
	if getErr != nil {
		log.Println("[GetMerchant] error in getting merchant:: ", getErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  getErr.Error(),
		})
		return
	}

	log.Printf("[GetMerchant] merchant:: %+v", getRes)

	ctx.JSON(http.StatusOK, gin.H{
		"status":     constant.StatusOK,
		"merchantId": getRes.ID,
		"name":       getRes.Name,
		"address":    getRes.Address,
		"message":    "merchant fetched successfully",
	})
}

func DeleteMerchant(ctx *gin.Context) {

	merchantId := ctx.Param("id")
	if merchantId == "" {
		log.Println("[DeleteMerchant] invalid/empty merchantId for update:: ")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  constant.ErrMerchantIdInvalid.Error(),
		})
		return
	}

	log.Printf("[DeleteMerchant] merchantId:: %s", merchantId)

	_, getErr := model.FetchMerchant(merchantId)
	if getErr != nil {
		log.Println("[DeleteMerchant] error in getting merchant:: ", getErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  getErr.Error(),
		})
		return
	}

	DeleteRes, DeleteErr := model.DeleteMerchant(merchantId)
	if DeleteErr != nil {
		log.Println("[DeleteMerchant] error in deleting merchant:: ", DeleteErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.StatusErr,
			"error":  DeleteErr.Error(),
		})
		return
	}

	log.Printf("[DeleteMerchant] merchantId:: %+v", DeleteRes)

	ctx.JSON(http.StatusOK, gin.H{
		"status":     constant.StatusOK,
		"merchantId": DeleteRes,
		"message":    "merchant deleted successfully",
	})
}
