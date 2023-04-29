### 1. "注册用户"

1. route definition

- Url: /v1/user/register
- Method: POST
- Request: `RegisterReq`
- Response: `RegisterResp`

2. request definition



```golang
type RegisterReq struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type RegisterResp struct {
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

### 2. "登录"

1. route definition

- Url: /v1/user/login
- Method: POST
- Request: `LoginReq`
- Response: `LoginResp`

2. request definition



```golang
type LoginReq struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginResp struct {
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

### 3. "获取用户详细信息"

1. route definition

- Url: /v1/user/detail
- Method: POST
- Request: `UserInfoReq`
- Response: `UserInfoResp`

2. request definition



```golang
type UserInfoReq struct {
}
```


3. response definition



```golang
type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}

type User struct {
	Id int64 `json:"id"`
	Mobile string `json:"mobile"`
	Name string `json:"Name"`
}
```

### 4. "wechat mini auth"

1. route definition

- Url: /v1/user/wxMiniAuth
- Method: POST
- Request: `WXMiniAuthReq`
- Response: `WXMiniAuthResp`

2. request definition



```golang
type WXMiniAuthReq struct {
	Code string `json:"code"`
	IV string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}
```


3. response definition



```golang
type WXMiniAuthResp struct {
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

