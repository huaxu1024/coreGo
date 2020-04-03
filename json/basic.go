package main


type Account struct {
	Email 		string `json:"email"`
	Nickname 	string `json:"nickname,omitempty"` //当其有值的时候就输出，而没有值(零值)的时候就不输出
	Password 	string `json:"-"` //ignore
	Money 		float64 `json:"money,string"`// 类型转换
	Extra 		[]interface{}
}

type User struct {
	Name    string `json:"name"`
	Age     int `json:"age"`
	Roles   []string `json:"roles"`
	Skill   map[string]float64 `json:"skill"`
	MAccount Account `json:account`
	Level map[string]interface{}
}

type LoginUser struct {
	UserName interface{} `json:"username"`
	Password string      `json:"password"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Mixed struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	city    string `json:"city"`
	Country string `json:"country"`
}