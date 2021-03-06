package models

import (
	"airad/common/support"
	"airad/common/util"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

func (u *User) TableName() string {
	return "user"
}

type User struct {
	Id          int    `json:"id" gorm:"column:id;pk;unique;auto_increment;int(11)"`
	Username    string `json:"username" gorm:"column:username;unique;size(32)"`
	Password    string `json:"password" gorm:"column:password;size(128)"`
	Avatar      string `json:"avatar, omitempty" gorm:"column:avatar;varbinary"`
	Salt        string `json:"salt" gorm:"column:salt;size(128)"`
	Token       string `json:"token" gorm:"column:token;size(256)"`
	Gender      int    `json:"gender" gorm:"column:gender;size(1)"` // 0:Male, 1: Female, 2: undefined
	Age         int    `json:"age" gorm:"column:age:size(3)"`
	Address     string `json:"address" gorm:"column:address;size(50)"`
	Email       string `json:"email" gorm:"column:email;size(50)"`
	LastLogin   int64  `json:"last_login" gorm:"column:last_login;type:integer;size(11)"`
	Status      int    `json:"status" gorm:"column:status;size(1)"` // 0: enabled, 1:disabled
	CreatedAt   int64  `json:"created_at" gorm:"column:created_at;size(11)"`
	UpdatedAt   int64  `json:"updated_at" gorm:"column:updated_at;type:integer;size(11)"`
	DeletedAt   int64  `json:"deleted_at" gorm:"column:deleted_at;type:integer;size(11)"`
	DeviceCount int    `json:"device_count" gorm:"column:device_count;size(64);default(0)"`
	//Device []*Device `orm:"reverse(many)"` // 设置一对多的反向关系
}

func getDBConn() *gorm.DB {
	db, err := support.GetMysqlConnInstance().GetDBConn("airad")
	if err != nil {
	}
	return db
}

// 检测用户是否存在
func CheckUserId(userId int) bool {
	db := getDBConn()
	var user User
	err := db.First(&user, userId).Error
	if nil != err && 0 != user.Id {
		return true
	}
	return false
}

// 检测用户是否存在
func CheckUserName(username string) bool {
	db := getDBConn()
	var user User
	err := db.Where("username", username).First(&user).Error
	if nil != err && 0 != user.Id {
		return true
	}
	return false
}

// 检测用户是否存在
func CheckUserIdAndToken(userId int, token string) bool {
	db := getDBConn()
	var user User
	err := db.Where(&User{Id: userId, Token: token}).First(&user).Error
	if nil != err && 0 != user.Id {
		return true
	}
	return false
}

// 检测用户是否存在
func CheckEmail(email string) bool {
	db := getDBConn()
	var user User
	err := db.Where("email", email).First(&user).Error
	if nil != err && 0 != user.Id {
		return true
	}
	return false
}

// CheckPass compare input password.
func (u *User) CheckPassword(password string) (ok bool, err error) {
	hash, err := util.GeneratePassHash(password, u.Salt)
	if err != nil {
		return false, err
	}
	return u.Password == hash, nil
}

// 根据用户ID获取用户
func GetUserById(id int) (v *User, err error) {
	db := getDBConn()
	err = db.First(v, id).Error
	return v, err
}

// 根据用户名字获取用户
func GetUserByUserName(username string) (v *User, err error) {
	db := getDBConn()
	var user = &User{}
	// 这里不能直接使用 v ,此时v还是 nil,会报nil错误
	err = db.Where("username = ?", username).First(user).Error
	return user, err
}

func GetUserAll() (ml []User, err error) {
	db := getDBConn()
	err = db.Find(&ml).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetUserByToken(token string) (bool, User) {
	db := getDBConn()
	var user User
	err := db.Where("token", token).First(&user).Error
	if nil != err && 0 != user.Id {
		return true, user
	}
	return false, user
}

func Login(username string, password string) (bool, *User) {
	user, err := GetUserByUserName(username)
	if nil != err && 0 != user.Id {
		return true, user
	}
	passwordHash, err := util.GeneratePassHash(password, user.Salt)
	if err != nil || passwordHash != user.Password {
		return false, nil
	}
	return true, user
}

// UpdateDevice updates User by DeviceCount and returns error if
// the record to be updated doesn't exist
func UpdateUserDeviceCount(m *User) (err error) {
	/*o := orm.NewOrm()
	  v := User{Id: m.Id}
	  m.DeviceCount += 1
	  // ascertain id exists in the database
	  if err = o.Read(&v); err == nil {
	      var num int64
	      if num, err = o.Update(m); err == nil {
	          fmt.Println("Number of records updated in database:", num)
	      }
	  }*/
	return
}

// updates User's Token and returns error if
// the record to be updated doesn't exist
func UpdateUserToken(m *User, token string) (err error) {
	//err = getDBConn().Model(m).Update(map[string]interface{}{"token": token, "updated_at": time.Now().UTC().Unix()}).Error
	err = getDBConn().Model(m).UpdateColumns(User{Token: token, UpdatedAt: time.Now().UTC().Unix()}).Error
	return
}

// updates User's LastLogin and returns error if
// the record to be updated doesn't exist
func UpdateUserLastLogin(m *User) (err error) {
	// todo 学习gorm tag
	// err = getDBConn().Model(m).Update(map[string]interface{}{"last_login": time.Now().UTC().Unix(), "updated_at": time.Now().UTC().Unix()}).Error
	err = getDBConn().Model(m).UpdateColumns(User{LastLogin: time.Now().UTC().Unix(), UpdatedAt: time.Now().UTC().Unix()}).Error
	return
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	db := getDBConn()
	err = db.Update(*m).Error
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	db := getDBConn()
	user, err := GetUserById(id)
	if err == nil {
		err = db.Delete(user).Error
	}
	return
}

//func HashPassword(password string) (string, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//	return string(bytes), err
//}
//
//func CheckPasswordHash(password, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}

//func generateToken() (tokenString string, err error) {
//	/* Create the token */
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	/* Create a map to store our claims
//	claims := token.Claims.(jwt.MapClaims)
//
//	/* Set token claims */
//	claims["admin"] = true
//	claims["name"] = "Ado Kukic"
//	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//
//	/* Sign the token with our secret */
//	tokenString, _ := token.SignedString(mySigningKey)
//}
