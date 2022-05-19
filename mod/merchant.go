package mod

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
		"status":       constant.StatusOK,
		"merchantCode": merchantId,
		"name":         merchant.Name,
		"address":      merchant.Address,
		"message":      "merchant added successfully",
	})
}
