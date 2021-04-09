package controllers

import (
	"lottery/models"
	"lottery/conf"
	"lottery/services"
)

func (api *LuckyApi) prize(prizeCode int, limitBlack bool) *models.ObjGiftPrize {
	var prizeGift *models.ObjGiftPrize
	//获得所有有库存的可发放的奖品列表
	giftList := services.NewGiftService().GetAllUse(true)
	for _, gift := range giftList {
		//如果 获奖概率在其中一个奖品的中奖概率区间内
		//则发放奖品，并停止循环
		if gift.PrizeCodeA <= prizeCode &&
			gift.PrizeCodeB >= prizeCode {
			// 中奖编码区间满足条件，说明可以中奖
			// 如果是黑名单用户并且抽到小奖的话，可以发奖
			// 如果不是黑名单用户，可以直接发奖
			if !limitBlack || gift.Gtype < conf.GtypeGiftSmall {
				prizeGift = &gift
				break
			}
		}
	}
	return prizeGift
}
