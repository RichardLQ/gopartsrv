package wxpay

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	wxpay "github.com/wechatpay-apiv3/wechatpay-go/utils"
	"gopartsrv/utils/mini"
	"log"
)


//创建支付服务
func createClient() (*core.Client,error) {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := wxpay.LoadPrivateKeyWithPath(mini.MchPKFileName)
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mini.Mchid, mini.MchCertificateSerialNumber, mchPrivateKey, mini.MchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}
	return client,nil
}

//下单
func CreatOrder(openid string,amount int64) (*jsapi.PrepayWithRequestPaymentResponse,error) {
	client,_:=createClient()
	svc := jsapi.JsapiApiService{Client: client}
	tradeNo := uuid.NewString()[:18]
	resp, _, err := svc.PrepayWithRequestPayment(context.Background(),
		jsapi.PrepayRequest{
			Appid:       core.String(mini.APPID),
			Mchid:       core.String(mini.Mchid),
			Description: core.String(mini.Description),
			OutTradeNo:  core.String(tradeNo),
			Attach:      core.String("龙猫小程序微信支付"),
			NotifyUrl:   core.String(mini.ServerURL),
			Amount: &jsapi.Amount{
				Total: core.Int64(amount),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(openid),
			},
		},
	)
	if err != nil {
		return resp,err
	}
	return resp ,nil
}

type WXPayNotify struct {
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	Appid         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	DeviceInfo    string `xml:"device_info"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	ErrCode       string `xml:"err_code"`
	ErrCodeDes    string `xml:"err_code_des"`
	Openid        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int64  `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       int64  `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	CouponFee     int64  `xml:"coupon_fee"`
	CouponCount   int64  `xml:"coupon_count"`
	CouponID0     string `xml:"coupon_id_0"`
	CouponFee0    int64  `xml:"coupon_fee_0"`
	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

//回调
func CallBack() *notify.Handler{
	mchPrivateKey, err := wxpay.LoadPrivateKeyWithPath(mini.MchPKFileName)
	if err != nil {
		log.Fatal("load merchant private key error")
	}
	ctx := context.Background()
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, mini.MchCertificateSerialNumber, mini.Mchid, mini.MchAPIv3Key)
	if err != nil {
		log.Fatal("load RegisterDownloaderWithPrivateKey error")
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(mini.Mchid)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(mini.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	return handler
}