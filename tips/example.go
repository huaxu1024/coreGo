package tips

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 如何获取Go routine的ID
func goId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("无法获得:%v", err))
	}
	return id
}

//  如何判断interface{}的类型
func typeSwitch(x interface{}) {
	switch a := x.(type) {
	case int:
		fmt.Println("%v is int\n", a)
	case string:
		fmt.Println("%v is string\n", a)
	default:
		fmt.Println("%v is interface{}\n", a)
	}
}

// 如何将各种类型转换成字符串？
func convertToString() {
	var text string

	// // int 转换成字符串
	text = strconv.Itoa(1000)
	fmt.Println("text:", text)

	// interface{} 转换成字符串
	var i interface{}
	i = "hello"
	text = i.(string)
	fmt.Println("text:", text)

	// byte[] 转换成字符串
	data := []byte("world")
	text = string(data)
	fmt.Println("text:", text)
}

// 如何生成随机数？如何用某一字符补齐特定长度字符串？
func GetDateNo() string {
	// 生成随机数
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	suffix := strconv.Itoa(r.Intn(1000))

	// 最大长度为4，并且用0来补齐
	if len(suffix) < 4 {
		suffix = fmt.Sprintf("%04s", suffix)
	}

	str := time.Now().Format("060102150405") + suffix
	return str
}

// 如何对map按照key进行排序？
func sortMap(dic map[string]string) {
	var keys []string
	for k := range dic {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		// 这里进行处理
		fmt.Println(k, ":", dic[k])
	}
}

// 如何生成Md5加密字符串？
func getMD5(text string) string {
	data := []byte(text)
	cryptoData := md5.Sum(data)
	return hex.EncodeToString(cryptoData[:])
}

// 如何获取当前程序的运行目录
func getDir() {
	abs, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("path:", abs)
}

// 如何获取当前的GOPATH路径
func getGoPath() {
	fmt.Println(build.Default.GOPATH)
}

// 如何发起HttpGet请求，并以字符串返回结果
func HTTPGet(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf(err.Error(), "HTTPGet")
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf(err.Error(), "HTTPGet")
		return "", err
	}
	return string(body), nil
}

// 如何发起HttpPost请求，并以字符串返回结果
func HTTPPost(url string, body string) (string, error) {
	data := []byte(body)
	reader := bytes.NewReader(data)

	res, err := http.Post(url, "application/x-www-form-urlencoded", reader)
	if err != nil {
		log.Fatalf(err.Error(), "HTTPPost-1")
		return "", err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf(err.Error(), "HTTPPost-2")
		return "", err
	}

	return string(resBody), nil
}

// 如何正确转换日期？
func getDateString() {
	const LayOut = "2006-01-02 15:04:05"

	beijing, _ := time.LoadLocation("Asia/Shanghai")

	// 获取当前日期并转换为字符串打印
	now := time.Now()
	nowStr := now.Format(LayOut)
	fmt.Println(nowStr)

	// 将字符串转换为日期，再转换回去然后打印
	now2, _ := time.ParseInLocation(LayOut, nowStr, beijing)
	now2Str := now2.Format(LayOut)
	fmt.Println(now2Str)

	diff := now.Sub(now2)
	fmt.Println("total :", diff.Hours())
}

// 如何获取当前URL
func getFullURL(r *http.Request) string {
	return "http://" + r.Host + r.RequestURI
	// return r.URL.Scheme + "://" + r.Host + r.RequestURI
}
