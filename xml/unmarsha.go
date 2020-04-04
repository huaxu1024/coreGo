package main

import (
	"encoding/xml"
	"fmt"
)

type Root struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	ResultCode string   `xml:"result_code"`
	ErrCode    string   `xml:"err_code"`
	ErrCodeDes string   `xml:"err_code_des"`
	CodeURL    string   `xml:"code_url"`
	AppId      string   `xml:"appid"`
	MchId      string   `xml:"mch_id"`
	NonceStr   string   `xml:"nonce_str"`
}

func (c Root) String() string {
	return fmt.Sprintf("root-[AppId:%v, MchId:%v, NonceStr:%v]", c.AppId, c.MchId, c.NonceStr)
}

func testXMl() {
	xmlStr := `<xml>
    <return_code><![CDATA[SUCCESS]]></return_code>
    <return_msg><![CDATA[OK]]></return_msg>
    <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
    <mch_id><![CDATA[10000100]]></mch_id>
    <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
    <openid><![CDATA[oUpF8uMuAJO_M2pxb1Q9zNjWeS6o]]></openid>
    <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
    <result_code><![CDATA[SUCCESS]]></result_code>
    <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
    <trade_type><![CDATA[Native]]></trade_type>
    <code_url><![CDATA[weixin://wxpay/s/An4baqw]]></code_url>
 </xml>`

	xmlByte := []byte(xmlStr)
	var root Root
	err := xml.Unmarshal(xmlByte, &root)
	if err != nil {
		fmt.Println("xml parse error", err.Error())
	}
	fmt.Printf("%v", root)
}
