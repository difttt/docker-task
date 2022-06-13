# EmailSystem

The system can send email to a specific user according different params of message

## 1. Start server

Run follwing commands

```
./email
```

Started successfully

```
starting server at port 3030
```

## 2. Test case 1

**2.1** **Send request via postman**

using postman to send post request, and the message is json format, such as:

target url: http://127.0.0.1:3030/form

message body:

```
{
"name": "shiyivei",
"address": "shiyivei@outlook.com",
"subject": "Contract Expiration Reminder",
"content":"Your contract is about to expire, please be aware of investment risks"
}`
```

**2.2** **Check the email**

Afer sending request, check the email and terminal to verify if everything is ok

## 3. Test case 2

**3.1 Send request via API**

```
func main() {

	//1. 调用url
	postUrl:= "http://127.0.0.1:3030/form"

	//2. 提供一组json格式的参数,样例如下

	body := []byte (
		`{
		"name": "xxx",
		"address": "xxx@outlook.com",
		"subject": "Contract Expiration Reminder",
		"content":"Your contract is about to expire, please be aware of investment risks"
	 	}`,
	)

	//3. 发送post请求（测试业务逻辑可以直接使用 postman 发送json格式的post请求）
	r,err := http.NewRequest("POST",postUrl,bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res,err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}
	fmt.Println("Email param provided!")
}
```

**3.2 Check the email**

Afer sending request, check the email and terminal to verify if everything is ok
