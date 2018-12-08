package controller

// Predefined const error strings.
const (
	ErrInputData    = "数据输入错误"
	ErrDatabase     = "数据库操作错误"
	ErrDupUser      = "用户信息已存在"
	ErrNoUser       = "用户信息不存在"
	ErrPass         = "密码不正确"
	ErrNoUserPass   = "用户信息不存在或密码不正确"
	ErrNoUserChange = "用户信息不存在或数据未改变"
	ErrInvalidUser  = "用户信息不正确"
	ErrOpenFile     = "打开文件出错"
	ErrWriteFile    = "写文件出错"
	ErrSystem       = "操作系统错误"
)

// UserData definition.
type UserSuccessLoginData struct {
	AccessToken string `json:"access_token"`
	UserName    string `json:"user_name"`
}

// CreateDevice definition.
type CreateObjectData struct {
	Id int `json:"id"`
}

// GetDevices definition.
type GetDeviceData struct {
	TotalCount int64       `json:"total_count"`
	Devices    interface{} `json:"devices"`
}

// GetAirAds definition.
type GetAirAdData struct {
	TotalCount int64       `json:"total_count"`
	AirAds     interface{} `json:"airads"`
}

// Predefined controller error/success values.

// BaseController definition.
//type BaseController struct {
//	beego.Controller
//

var sqlOp = map[string]string{
	"eq": "=",
	"ne": "<>",
	"gt": ">",
	"ge": ">=",
	"lt": "<",
	"le": "<=",
}

// ParseQueryParm parse query parameters.
//   query=col1:op1:val1,col2:op2:val2,...
//   op: one of eq, ne, gt, ge, lt, le
//func (base *BaseController) ParseQueryParameter() (v map[string]string, o map[string]string, err error) {
//	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
//	queryVal := make(map[string]string)
//	queryOp := make(map[string]string)
//
//	query := base.GetString("query")
//	if query == "" {
//		return queryVal, queryOp, nil
//	}
//
//	for _, cond := range strings.Split(query, ",") {
//		kov := strings.Split(cond, ":")
//		if len(kov) != 3 {
//			return queryVal, queryOp, errors.New("Query format != k:o:v")
//		}
//
//		var key string
//		var value string
//		var operator string
//		if !nameRule.MatchString(kov[0]) {
//			return queryVal, queryOp, errors.New("Query key format is wrong")
//		}
//		key = kov[0]
//		if op, ok := sqlOp[kov[1]]; ok {
//			operator = op
//		} else {
//			return queryVal, queryOp, errors.New("Query operator is wrong")
//		}
//		value = strings.Replace(kov[2], "'", "\\'", -1)
//
//		queryVal[key] = value
//		queryOp[key] = operator
//	}
//
//	return queryVal, queryOp, nil
//}

// ParseOrderParameter parse order parameters.
//   order=col1:asc|desc,col2:asc|esc,...
//func (base *BaseController) ParseOrderParameter() (o map[string]string, err error) {
//	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
//	order := make(map[string]string)
//
//	v := base.GetString("order")
//	if v == "" {
//		return order, nil
//	}
//
//	for _, cond := range strings.Split(v, ",") {
//		kv := strings.Split(cond, ":")
//		if len(kv) != 2 {
//			return order, errors.New("Order format != k:v")
//		}
//		if !nameRule.MatchString(kv[0]) {
//			return order, errors.New("Order key format is wrong")
//		}
//		if kv[1] != "asc" && kv[1] != "desc" {
//			return order, errors.New("Order val isn't asc/desc")
//		}
//
//		order[kv[0]] = kv[1]
//	}
//
//	return order, nil
//}

// ParseLimitParameter parse limit parameter.
//   limit=n
//func (base *BaseController) ParseLimitParameter() (l int64, err error) {
//	if v, err := base.GetInt64("limit"); err != nil {
//		return 10, err
//	} else if v > 0 {
//		return v, nil
//	} else {
//		return 10, nil
//	}
//}

// ParseOffsetParameter parse offset parameter.
//   offset=n
//func (base *BaseController) ParseOffsetParameter() (o int64, err error) {
//	if v, err := base.GetInt64("offset"); err != nil {
//		return 0, err
//	} else if v > 0 {
//		return v, nil
//	} else {
//		return 0, nil
//	}
//}

// VerifyForm use validation to verify input parameters.
//func (base *BaseController) VerifyForm(obj interface{}) (err error) {
//	valid := validation.Validation{}
//	ok, err := valid.Valid(obj)
//	if err != nil {
//		return err
//	}
//	if !ok {
//		str := ""
//		for _, err := range valid.Errors {
//			str += err.Key + ":" + err.Message + ";"
//		}
//		return errors.New(str)
//	}
//
//	return nil
//}

// ParseToken parse JWT token in http header.
//func (base *BaseController) ParseToken() (t *jwt.Token, e *ControllerError) {
//	authString := base.Ctx.Input.Header("Authorization")
//	beego.Debug("AuthString:", authString)
//
//	kv := strings.Split(authString, " ")
//	if len(kv) != 2 || kv[0] != "Bearer" {
//		beego.Error("AuthString invalid:", authString)
//		return nil, errInputData
//	}
//	tokenString := kv[1]
//
//	// Parse token
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte("secret"), nil
//	})
//	if err != nil {
//		beego.Error("Parse token:", err)
//		if ve, ok := err.(*jwt.ValidationError); ok {
//			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
//				// That's not even a token
//				return nil, errInputData
//			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
//				// Token is either expired or not active yet
//				return nil, errExpired
//			} else {
//				// Couldn't handle this token
//				return nil, errInputData
//			}
//		} else {
//			// Couldn't handle this token
//			return nil, errInputData
//		}
//	}
//	if !token.Valid {
//		beego.Error("Token invalid:", tokenString)
//		return nil, errInputData
//	}
//	beego.Debug("Token:", token)
//
//	return token, nil
//}
