package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/student/ginblog/common"
	"github.com/student/ginblog/datamodels"
)

type IUser interface {
	//连接数据
	Conn() error
	Insert(*datamodels.User) int64
	Delete(int64) bool
	Update(*datamodels.User) error
	// SelectByKey(int64) (*datamodels.User, error)
	SelectAll() ([]*datamodels.User, error)
	// SelectAll()
	SelectByName(*datamodels.User, string) int64
	FindUser(*datamodels.LoginUser, string) *datamodels.LoginUser
	FindUserById(*datamodels.User, int64) *datamodels.User
}
type UserManager struct {
	table     string
	mysqlConn *gorm.DB
}

func NewUserManager(table string, db *gorm.DB) IUser {
	return &UserManager{table: table, mysqlConn: db}
}

//数据连接
func (p *UserManager) Conn() (err error) {
	if p.mysqlConn == nil {
		mysql := common.GormInit()
		if err != nil {
			return err
		}
		p.mysqlConn = mysql
	}
	if p.table == "" {
		p.table = "users"
	}
	return
}

//插入
func (p *UserManager) Insert(user *datamodels.User) int64 {
	//1.判断连接是否存在
	if err := p.Conn(); err != nil {
		return 0
	}
	// u := new(datamodels.User)
	result := p.mysqlConn.Create(user).RowsAffected
	return result
}

//根据用户凭查询
func (p *UserManager) SelectByName(user *datamodels.User, name string) int64 {
	if err := p.Conn(); err != nil {
		return 0
	}
	result := p.mysqlConn.Where("username = ?", name).Find(user).RowsAffected
	return result
}
func (p *UserManager) FindUser(user *datamodels.LoginUser, name string) *datamodels.LoginUser {
	if err := p.Conn(); err != nil {
		return &datamodels.LoginUser{}
	}
	p.mysqlConn.Table("users").Where("username = ?", name).Find(user)
	return user
}
func (p *UserManager) SelectAll() (userArray []*datamodels.User, err error) {
	var user []*datamodels.User
	if err := p.Conn(); err != nil {
		return nil, err
	}
	errors := p.mysqlConn.Table("users").Find(&user).Error
	if len(user) == 0 {
		return nil, nil
	}
	return user, errors
}

//通过id查询用户信息
func (p *UserManager) FindUserById(user *datamodels.User, id int64) *datamodels.User {
	if err := p.Conn(); err != nil {
		return &datamodels.User{}
	}
	p.mysqlConn.Table("users").Where("id = ?", id).Find(user)
	return user
}

//更新用户信息
func (p *UserManager) Update(user *datamodels.User) error {
	if err := p.Conn(); err != nil {
		return err
	}
	// id := user.ID
	err1 := p.mysqlConn.Model(&user).Debug().Update(map[string]interface{}{"username": user.Username, "password": user.Password}).Error
	return err1
}
func (p *UserManager) Delete(id int64) bool {
	if err := p.Conn(); err != nil {
		return false
	}
	err := p.mysqlConn.Exec("DELETE FROM users WHERE id = ? ", id).Error
	if err != nil {
		return false
	}
	return true
}
