package task

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type Response struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

type SignInResponse struct {
	IncreasedScore   int `json:"increased_score"`
	RewardCouponList []struct {
		ConditionAmount float64 `json:"condition_amount"`
		DiscountAmount  float64 `json:"discount_amount"`
	} `json:"reward_coupon_list"`
}

type QueryPointResponse struct {
	Balance      int   `json:"balance"`
	ExpiringCoin int   `json:"expiring_coin"`
	ExpireTime   int64 `json:"expire_time"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	IsBindPhone  bool   `json:"is_bind_phone"`
	UserId       string `json:"user_id"`
	NickName     string `json:"nickName"`
	IsNewUser    bool   `json:"isNewUser"`
}
