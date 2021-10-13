package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fghwett/pupu/config"
	"github.com/fghwett/pupu/util"
)

var conf *config.Conf

type Task struct {
	config *config.Config
	client *http.Client
	result []string
}

func New(c *config.Conf) *Task {
	conf = c
	return &Task{
		config: c.Config,
		client: &http.Client{},
		result: []string{"==== 朴朴超市任务 ===="},
	}
}

func (t *Task) Do() {
	if err := t.getToken(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【刷新登陆状态】：失败 %s", err))
		return
	}

	if err := t.signTask(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【签到任务】：失败 %s", err))
		return
	}

	if err := t.shareTask(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【分享任务】：失败 %s", err))
		return
	}

	if err := t.queryPointTask(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【查询积分任务】：失败 %s", err))
		return
	}
}

func (t *Task) getToken() error {
	if time.Now().UnixNano()/1e6 < t.config.ExpiredAt {
		t.result = append(t.result, fmt.Sprintf("【刷新登陆状态】：已经是最新了"))
		return nil
	}

	reqData, err := json.Marshal(&RefreshTokenRequest{
		RefreshToken: t.config.RefreshToken,
	})
	if err != nil {
		return fmt.Errorf("请求参数序列化失败 %s", err)
	}

	reqUrl := "https://cauth.pupuapi.com/clientauth/user/refresh_token"
	req, err := http.NewRequest(http.MethodPut, reqUrl, bytes.NewReader(reqData))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Pupumall/2.7.0;iOS 15.0;245D5D8A-3860-490A-BD61-9B3E10221D4E")
	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Add("pp-userid", "a32193f0-de63-4182-a1d3-8919cb92fc6f")
	req.Header.Add("pp-os", "20")
	req.Header.Add("pp-storeid", "fa701852-eb6e-405e-9321-aa2a12d0fb09")

	resp, err := t.client.Do(req)

	response := &Response{Data: &RefreshTokenResponse{}}
	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	if response.ErrCode != 0 {
		return fmt.Errorf(response.ErrMsg)
	}

	respData := response.Data.(*RefreshTokenResponse)
	t.config = &config.Config{
		RefreshToken: respData.RefreshToken,
		AccessToken:  fmt.Sprintf("Bearer %s", respData.AccessToken),
		ExpiredAt:    respData.ExpiresIn / 1e3,
	}

	if err := conf.Update(t.config); err != nil {
		return err
	}

	t.result = append(t.result, "【刷新登陆状态】：成功")

	return nil
}

func (t *Task) signTask() error {
	reqUrl := "https://j1.pupuapi.com/client/game/sign?city_zip=350100&challenge="
	req, _ := http.NewRequest(http.MethodPost, reqUrl, nil)
	req.Header.Set("Authorization", t.config.AccessToken)
	req.Header.Set("User-Agent", "Pupumall/2.7.0;iOS 15.0;245D5D8A-3860-490A-BD61-9B3E10221D4E")
	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Add("pp-userid", "a32193f0-de63-4182-a1d3-8919cb92fc6f")
	req.Header.Add("pp-os", "20")
	req.Header.Add("pp-storeid", "fa701852-eb6e-405e-9321-aa2a12d0fb09")

	resp, err := t.client.Do(req)

	response := &Response{Data: &SignInResponse{}}
	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	if response.ErrCode != 0 {
		return fmt.Errorf(response.ErrMsg)
	}

	signInResp := response.Data.(*SignInResponse)
	result := fmt.Sprintf("【签到任务】：签到成功 获得%d积分", signInResp.IncreasedScore)

	if signInResp.RewardCouponList != nil && len(signInResp.RewardCouponList) > 0 {
		for _, x := range signInResp.RewardCouponList {
			result += fmt.Sprintf(" 获得满%.2f减%.2f优惠券", x.ConditionAmount/100, x.DiscountAmount/100)
		}
	}

	t.result = append(t.result, result)

	return nil
}

func (t *Task) shareTask() error {
	reqUrl := "https://j1.pupuapi.com/client/game/sign/share"

	req, _ := http.NewRequest(http.MethodPost, reqUrl, nil)
	req.Header.Set("Authorization", t.config.AccessToken)
	req.Header.Set("User-Agent", "Pupumall/2.7.0;iOS 15.0;245D5D8A-3860-490A-BD61-9B3E10221D4E")
	req.Header.Set("content-type", "application/json; charset=utf-8")
	req.Header.Add("pp-userid", "a32193f0-de63-4182-a1d3-8919cb92fc6f")
	req.Header.Add("pp-os", "20")
	req.Header.Add("pp-storeid", "fa701852-eb6e-405e-9321-aa2a12d0fb09")

	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}

	response := &Response{}
	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	if response.ErrCode != 0 {
		return fmt.Errorf(response.ErrMsg)
	}

	t.result = append(t.result, fmt.Sprintf("【分享任务】：完成 获取%f积分", response.Data.(float64)))

	return nil
}

func (t *Task) queryPointTask() error {
	reqUrl := "https://j1.pupuapi.com/client/coin"

	req, _ := http.NewRequest(http.MethodGet, reqUrl, nil)
	req.Header.Set("Authorization", t.config.AccessToken)
	req.Header.Set("User-Agent", "Pupumall/2.7.0;iOS 15.0;245D5D8A-3860-490A-BD61-9B3E10221D4E")
	req.Header.Add("pp-userid", "a32193f0-de63-4182-a1d3-8919cb92fc6f")
	req.Header.Add("pp-os", "20")
	req.Header.Add("pp-storeid", "fa701852-eb6e-405e-9321-aa2a12d0fb09")

	resp, err := t.client.Do(req)
	if err != nil {
		return err
	}

	response := &Response{Data: &QueryPointResponse{}}
	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	if response.ErrCode != 0 {
		return fmt.Errorf(response.ErrMsg)
	}

	respData := response.Data.(*QueryPointResponse)

	t.result = append(t.result, fmt.Sprintf("【查询积分任务】：朴分%d，即将过期朴分%d，过期时间%s", respData.Balance, respData.ExpiringCoin, time.Unix(respData.ExpireTime/1000, 0).Format("2006-01-02 15:04:05")))

	return nil
}

func (t *Task) GetResult() string {
	return strings.Join(t.result, " \n\n ")
}
