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
	DailySignCoin int `json:"daily_sign_coin"`
	CouponList    []struct {
		ConditionAmount float64 `json:"condition_amount"`
		DiscountAmount  float64 `json:"discount_amount"`
	} `json:"coupon_list"`
	TeamRewardCoin       int    `json:"team_reward_coin"`
	ShowCreateTeamButton bool   `json:"show_create_team_button"`
	Title                string `json:"title"`
	SubTitle             string `json:"sub_title"`
	RewardExplanation    string `json:"reward_explanation"`
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

type SignInPage struct {
	IsSigned              bool   `json:"is_signed"`             // 是否已签到
	IsOpenRemind          bool   `json:"is_open_remind"`        // 是否开启提醒
	RewardExplanation     string `json:"reward_explanation"`    // 奖励内容
	IsAllowSignSupplement bool   `json:"isAllowSignSupplement"` // 是否允许补签
	SignRecordList        []struct {
		SignDayOrder   int `json:"sign_day_order"`   // 签到第几天
		SignStatus     int `json:"sign_status"`      // 签到状态
		SignRewardCoin int `json:"sign_reward_coin"` // 签到所得积分
	} `json:"sign_record_list"`
	DailySignRewardCoin      int    `json:"daily_sign_reward_coin"`      // 每日签到奖励积分
	ContinuitySignThreeImg   string `json:"continuity_sign_three_img"`   // 连续签到三天图片
	ContinuitySignSevenImg   string `json:"continuity_sign_seven_img"`   // 连续签到七天图片
	TeamRewardCoin           int    `json:"team_reward_coin"`            // 组队奖励积分
	ShowCreateTeamButton     bool   `json:"show_create_team_button"`     // 是否显示组队按钮
	IsShowTwistedEgg         bool   `json:"is_show_twisted_egg"`         // 是否显示扭蛋按钮
	TodaySignOrder           int    `json:"today_sign_order"`            // 今日签到是第几天
	UserNoviceGuidanceStatus int    `json:"user_novice_guidance_status"` // 用户新手指导状态
	CoinGroupExplanation     string `json:"coin_group_explanation"`      // 组团积分规则
}

type QueryDiscountResponse struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    []*Discount `json:"data"`
}

type Discount struct {
	Rule              *DiscountRule `json:"rule"`
	Id                string        `json:"id"`
	Num               string        `json:"num"`
	GetTime           int64         `json:"get_time"`
	IsUsed            bool          `json:"is_used"`
	DateStart         int64         `json:"date_start"`
	DateEnd           int64         `json:"date_end"`
	IsExpired         bool          `json:"is_expired"`
	IsAvailable       bool          `json:"is_available"`
	IsGlobalAvailable interface{}   `json:"is_global_available"`
	IsExpiringSoon    bool          `json:"is_expiring_soon"`
	PeriodOfValidity  string        `json:"period_of_validity"`
	Reason            interface{}   `json:"reason"`
	ReasonV2          interface{}   `json:"reason_v2"`
	LimitProduct      interface{}   `json:"limit_product"`
	ExtLimitProduct   interface{}   `json:"ext_limit_product"`
	DifferenceAmount  float64       `json:"difference_amount"`
}

type DiscountRule struct {
	DiscountId            string   `json:"discount_id"`
	Name                  string   `json:"name"`
	DiscountType          int      `json:"discount_type"`
	ConditionAmount       int      `json:"condition_amount"`
	DiscountAmount        int      `json:"discount_amount"`
	Banner                string   `json:"banner"`
	Content               string   `json:"content"`
	LimitCategoriesStatus bool     `json:"limit_categories_status"`
	LimitProduct          bool     `json:"limit_product"`
	ExcludeSpecials       bool     `json:"exclude_specials"`
	DiscountStatusType    int      `json:"discount_status_type"`
	LimitAreaStatus       bool     `json:"limit_area_status"`
	IsOrderCannotInvoice  bool     `json:"is_order_cannot_invoice"`
	ConsumeRemark         []string `json:"consume_remark"`
	ValidTimeRemark       string   `json:"valid_time_remark"`
	GetType               int      `json:"get_type"`
	AmountPrepaid         int      `json:"amount_prepaid"`
	ConditionAmountDesc   string   `json:"condition_amount_desc"`
	DiscountAmountDesc    string   `json:"discount_amount_desc"`
}
