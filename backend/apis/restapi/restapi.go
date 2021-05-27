package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/noirstar/autotrading/backend/models"
	"github.com/noirstar/autotrading/backend/utils/env"
	"github.com/noirstar/autotrading/backend/utils/jwt"
	"github.com/noirstar/autotrading/backend/utils/myerr"
)

var baseURL = env.GetEnv("UPBIT_BASE_URL")

// GetAccount 전체 계좌 조회
func GetAccount(accessKey string, secretKey string) []byte {

	reqURL := baseURL + "/v1/accounts"
	tokenString := jwt.GetJwtToken(accessKey, secretKey)

	return RequestToServer(reqURL, "GET", tokenString, nil)
}

// GetOrderChance 주문 가능 정보 - 마켓별 주문 가능 정보 확인
func GetOrderChance(accessKey string, secretKey string, query *models.ReqChance) []byte {

	reqURL := baseURL + "/v1/orders/chance"
	queryMap := ConvertStructToMap(query)
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, queryMap)

	return RequestToServer(reqURL, "GET", tokenString, queryMap)
}

// GetOrderSearch 개별 주문 조회 - 주문 UUID 를 통해 개별 주문건을 조회
func GetOrderSearch(accessKey string, secretKey string, query *models.ReqOrderSearch) []byte {

	reqURL := baseURL + "/v1/order"
	queryMap := ConvertStructToMap(query)
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, queryMap)

	return RequestToServer(reqURL, "GET", tokenString, queryMap)
}

// GetOrdersSearch 주문 리스트 조회 - 주문 리스트를 조회
func GetOrdersSearch(accessKey string, secretKey string, query *models.ReqOrdersSearch) []byte {

	reqURL := baseURL + "/v1/orders"
	queryMap := ConvertStructToMap(query)
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, queryMap)

	return RequestToServer(reqURL, "GET", tokenString, queryMap)
}

// DeleteOrder 주문 취소 접수 - 주문 UUID를 통해 해당 주문에 대한 취소 접수
func DeleteOrder(accessKey string, secretKey string, query *models.ReqDeleteOrder) []byte {

	reqURL := baseURL + "/v1/order"
	queryMap := ConvertStructToMap(query)
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, queryMap)

	return RequestToServer(reqURL, "DELETE", tokenString, queryMap)
}

// PostOrder 주문하기
func PostOrder(accessKey string, secretKey string, query *models.ReqOrders) []byte {

	reqURL := baseURL + "/v1/orders"
	queryMap := ConvertStructToMap(query)
	tokenString := jwt.GetJwtTokenWithQuery(accessKey, secretKey, queryMap)

	return RequestToServer(reqURL, "POST", tokenString, queryMap)
}

// GetMarketCode 마켓 코드 조회 - 업비트에서 거래 가능한 마켓 목록
func GetMarketCode() []byte {

	reqURL := baseURL + "/v1/market/all?isDetails=true"

	return RequestToServerSimple(reqURL, "GET", nil)
}

// GetMinuteCandles 분 캔들 조회
func GetMinuteCandles(query *models.ReqMinuteCandles, unit string) []byte {

	reqURL := baseURL + "/v1/candles/minutes/" + unit

	return RequestToServerSimple(reqURL, "GET", ConvertStructToMap(query))
}

// GetDayCandles 일 캔들 조회
func GetDayCandles(query *models.ReqDayCandles) []byte {

	reqURL := baseURL + "/v1/candles/days"

	return RequestToServerSimple(reqURL, "GET", ConvertStructToMap(query))
}

// GetWeekCandles 주 캔들 조회
func GetWeekCandles(query *models.ReqWeekCandles) []byte {

	reqURL := baseURL + "/v1/candles/weeks"

	return RequestToServerSimple(reqURL, "GET", ConvertStructToMap(query))
}

// GetMonthsCandles 달 캔들 조회
func GetMonthsCandles(query *models.ReqMonthCandles) []byte {

	reqURL := baseURL + "/v1/candles/months"

	return RequestToServerSimple(reqURL, "GET", ConvertStructToMap(query))
}

// RequestToServer 업비트 서버로 요청
func RequestToServer(reqURL string, method string, tokenString string, query map[string]interface{}) []byte {

	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, nil)
	myerr.CheckErr(err)

	q := url.Values{}

	for key, value := range query {
		switch val := value.(type) {
		case string:
			q.Add(key, value.(string))
		case int, uint32, uint64:
			q.Add(key, value.(string))
		case []string:
			for _, v := range val {
				q.Add(key, v)
			}
		case []interface{}:
			for _, v := range val {
				q.Add(key, v.(string))
			}
		}
	}
	req.URL.RawQuery = q.Encode()

	fmt.Println(q.Encode())

	req.Header.Add("Authorization", "Bearer "+tokenString)
	res, err := client.Do(req)
	myerr.CheckErr(err)
	bytes, err := ioutil.ReadAll(res.Body)
	myerr.CheckErr(err)
	defer res.Body.Close()

	return bytes
}

// RequestToServerSimple 토큰 미포함 요청
func RequestToServerSimple(reqURL string, method string, query map[string]interface{}) []byte {

	client := &http.Client{}
	req, err := http.NewRequest(method, reqURL, nil)
	myerr.CheckErr(err)

	q := req.URL.Query()

	for key, value := range query {
		switch val := value.(type) {
		case string:
			q.Add(key, value.(string))
		case int, uint32, uint64:
			q.Add(key, value.(string))
		case []string:
			for _, v := range val {
				q.Add(key, v)
			}
		case []interface{}:
			for _, v := range val {
				q.Add(key, v.(string))
			}
		}
	}
	req.URL.RawQuery = q.Encode()

	fmt.Println(q.Encode())

	res, err := client.Do(req)
	myerr.CheckErr(err)
	bytes, err := ioutil.ReadAll(res.Body)
	myerr.CheckErr(err)
	defer res.Body.Close()

	return bytes

}

// ConvertStructToMap struct -> map[string]interface{}
func ConvertStructToMap(object interface{}) map[string]interface{} {
	conv := make(map[string]interface{})
	tmp, err := json.Marshal(object)
	myerr.CheckErr(err)
	json.Unmarshal(tmp, &conv)
	return conv
}
