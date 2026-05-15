package postJob

type PromotionItem struct {
	Icon string `json:"icon"`
	Text string `json:"text"`
}

type PostJobPromotion struct {
	Promotion    PromotionItem `json:"promotion"`
	SecPromotion PromotionItem `json:"secPromotion"`
}
