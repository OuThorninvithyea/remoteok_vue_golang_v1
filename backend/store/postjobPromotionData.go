package store

import "backend/model/postJob"

var PostjobPromotionData = postJob.PostJobPromotion{
	Promotion: postJob.PromotionItem{
		Icon: "🏷",
		Text: "A discount of 10% with code CAF11F815FF8E1043683838170244255 is applied on checkout",
	},
	SecPromotion: postJob.PromotionItem{
		Icon: "⛳️",
		Text: "Save up to 50% when buying multiple job posts — Buy a bundle →",
	},
}
