package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username          string    `gorm:"column:username"`
	RealName          string    `gorm:"column:realname"`
	Orgid             string    `gorm:"column:orgid"`
	Password          string    `gorm:"column:password"`
	Status            string    `gorm:"column:status"`
	Roles             UserRoles `gorm:"column:roles"`
	Expireddate       time.Time `gorm:"column:expireddate"`
	Logintime         time.Time `gorm:"column:logintime"`
	Loginip           string    `gorm:"column:loginip"`
	Lasttime          time.Time `gorm:"column:lasttime"`
	Lastip            string    `gorm:"column:lastip"`
	Skin              string    `gorm:"column:skin"`
	Langcode          string    `gorm:"column:langcode"`
	Sex               string    `gorm:"column:sex"`
	Birthday          time.Time `gorm:"column:birthday"`
	Idcard            string    `gorm:"column:idcard"`
	School            string    `gorm:"column:school"`
	Graduation        string    `gorm:"column:graduation"`
	Degree            string    `gorm:"column:degree"`
	Major             string    `gorm:"column:major"`
	Country           string    `gorm:"column:country"`
	Province          string    `gorm:"column:province"`
	City              string    `gorm:"column:city"`
	Address           string    `gorm:"column:address"`
	Postcode          string    `gorm:"column:postcode"`
	Phone             string    `gorm:"column:phone"`
	Fax               string    `gorm:"column:fax"`
	Mobile            string    `gorm:"column:mobile"`
	Email             string    `gorm:"column:email"`
	Remark            string    `gorm:"column:remark"`
	Creator           string    `gorm:"column:creator"`
	Modifier          string    `gorm:"column:modifier"`
	Usertype          string    `gorm:"column:usertype"`
	Postid            string    `gorm:"column:postid"`
	Isleader          bool      `gorm:"column:isleader;null;default(false)"`
	Expired           string    `gorm:"column:expired;null;default(0)"`
	Ipconfig          string    `gorm:"column:ipconfig"`
	EnglishName       string    `gorm:"column:english_name"`
	Nationality       string    `gorm:"column:nationality"`
	Employeeid        string    `gorm:"column:employeeid"`
	Entrydate         time.Time `gorm:"column:entrydate"`
	ResidenceAddr     string    `gorm:"column:residence_addres)"`
	ResidenceType     string    `gorm:"column:residence_type"`
	MaritalStatus     string    `gorm:"column:marital_status"`
	NativePlace       string    `gorm:"column:native_place"`
	WorkDate          time.Time `gorm:"column:work_date"`
	ContactWay        string    `gorm:"column:contact_way"`
	ContactPerson     string    `gorm:"column:contact_person"`
	ProfessionalTitle string    `gorm:"column:professional_title"`
	ComputerLevel     string    `gorm:"column:computer_level"`
	ComputerCert      string    `gorm:"column:computer_cert"`
	EnglishLevel      string    `gorm:"column:english_level"`
	EnglishCert       string    `gorm:"column:english_cert"`
	JapaneseLevel     string    `gorm:"column:japanese_level"`
	JapaneseCert      string    `gorm:"column:japanese_cert"`
	Speciality        string    `gorm:"column:speciality"`
	SpecialityCert    string    `gorm:"column:speciality_cert"`
	HobbySport        string    `gorm:"column:hobby_sport"`
	HobbyArt          string    `gorm:"column:hobby_art"`
	HobbyOther        string    `gorm:"column:hobby_other"`
	KeyUser           string    `gorm:"column:key_user"`
	WorkCard          string    `gorm:"column:work_card"`
	GuardCard         string    `gorm:"column:guard_card"`
	Computer          string    `gorm:"column:computer"`
	Ext               string    `gorm:"column:ext"`
	Msn               string    `gorm:"column:msn"`
	Rank              string    `gorm:"column:rank"`
}

type Profile struct {
	UserID      string
	UserName    string
	CompanyCode string
	LoginIP     string
	LastLogin   time.Time
}

type UserRole struct {
	RoleID     string
	ExpireDate time.Time
}

type UserRoles []*UserRole

type Users []*User

//用户查询条件参数
type UserQueryParam struct {
	UserName     string
	RealName     string
	UserNameLike string
	RealNameLike string
	Status       int
	RoleIDs      []string
}

type UserQueryOptions struct {
	PageParam    *PaginationParam //分页参数
	IncludeRoles bool             //包含角色
}

type UserQueryResult struct {
	Res     Users
	PageRes *PaginationResult
}

type UserPageShow struct {
	UserName    string
	RealName    string
	PhoneNumber string
	Email       string
	Status      int
	CreatedAt   time.Time
	Roles       []*Role
}

func (User) TableName() string {
	return "user_tb"
}

//ToRoleIDs will traverse UserRoles and return aggregation of every userrole's roleid.
func (uRoles UserRoles) ToRoleIDs() []string {
	roleIDs := make([]string, len(uRoles))
	for i, urole := range uRoles {
		roleIDs[i] = urole.RoleID
	}
	return roleIDs
}

func (users Users) ToRoleIDs() []string {
	var roleIDs []string
	for _, u := range users {
		roleIDs = append(roleIDs, u.Roles.ToRoleIDs()...)
	}
	return roleIDs
}

func (users Users) ToPageShows() []*PaginationResult
