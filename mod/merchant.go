package mod

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

type NewMerchant struct {
	Name    string `json:"name" validate:"min=3,max=64,regexp=^[a-zA-Z0-9_ ]*$"`
	Address string `json:"address" validate:"min=6,max=256,regexp=^[a-zA-Z0-9_ \\,]*$"`
}

func AddMerchant(ctx *gin.Context) {

	var merchant NewMerchant

	if err := ctx.ShouldBindJSON(&merchant); err != nil {
		log.Println("[AddMerchant] error from json reading:: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Printf("[AddMerchant] New merchant:: %+v\n", merchant)

	if validateErr := validator.Validate(merchant); validateErr != nil {
		log.Println("[AddMerchant] validateErr :: ", validateErr)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": validateErr.Error(),
		})
		return
	}

	merchantId := uuid.New()
	log.Println("[AddMerchant] Generated merchantId:: ", merchantId)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "merchant added successfully",
	})
}
