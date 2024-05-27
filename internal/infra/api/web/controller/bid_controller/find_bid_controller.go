package bid_controller

import (
	"context"
	"fullcycle-auction_go/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (u *BidController) FindBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	bidOutputList, err := u.bidUseCase.FindBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, bidOutputList)
}
