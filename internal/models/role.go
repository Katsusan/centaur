package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	RoleID       string
	RoleName     string
	RolePriority int
	RoleMemo     string
	Creator      string
	Menus        RoleMenus
}

type Roles []*Role

type RoleMenu struct {
	MenuID    string
	Actions   []string
	Resources []string
}

type RoleMenus []*RoleMenu

type RoleQueryParam struct {
	RoleIDs      []string
	RoleName     string
	RoleNameLike string
	UserID       string
}

type RoleQueryOptions struct {
	PageParam    *PaginationParam
	IncludeMenus bool
}

type RoleQueryResult struct {
	Res     Roles
	PageRes *PaginationResult
}

func (r RoleMenus) ToMenusIDs() []string {
	ids := make([]string, len(r))
	for i, rolemenu := range r {
		ids[i] = rolemenu.MenuID
	}
	return ids
}

func (roles Roles) ToMenuIDs() []string {
	var ids []string
	for _, role := range roles {
		ids = append(ids, role.Menus.ToMenusIDs()...)
	}
	return ids
}

//将Role数组转化为RoleID->Role键值对
func (roles Roles) ToMap() map[string]*Role {
	mRoles := make(map[string]*Role)
	for _, role := range roles {
		mRoles[role.RoleID] = role
	}
	return mRoles
}

//获取Role数组中的角色名称列表
func (roles Roles) ToNames() []string {
	names := make([]string, len(roles))
	for i, role := range roles {
		names[i] = role.RoleName
	}
	return names
}

func (rolemenus RoleMenus) ToMenuIDs() []string {
	menuids := make([]string, len(rolemenus))
	for i, rolemenu := range rolemenus {
		menuids[i] = rolemenu.MenuID
	}
	return menuids
}
