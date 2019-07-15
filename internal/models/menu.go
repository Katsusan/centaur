package models

import (
	"strings"

	"github.com/jinzhu/gorm"
)

//Menu菜单属性
type Menu struct {
	gorm.Model
	MenuID         string
	MenuName       string
	MenuPriority   int
	MenuIcon       string
	MenuRouter     string
	Hidden         int
	MenuParentID   string
	MenuParentPath string
	MenuCreator    string
}

type Menus []*Menu

//MenuAction 定义菜单动作
type MenuAction struct {
	ActionCode string
	ActionName string
}

type MenuActions []*MenuAction

//MenuResource 定义菜单资源属性
type MenuResource struct {
	ResourceCode   string
	ResourceName   string
	ResourceMethod string
	ResourcePath   string
}

//MenuQueryParam 描述了菜单查询条件
type MenuQueryParam struct {
	QueryParamIDs    []string
	MenuNameLike     string //模糊查询
	ParentID         string
	PrefixParentPath string
	Hidden           int
}

type MenuQueryOptions struct {
	PageParam        *PaginationParam
	IncludeActions   bool
	IncludeResources bool
}

//MenuTree 定义菜单树
type MenuTree struct {
	MenuTreeID       string
	MenuTreeName     string
	MenuTreePriority int
	MenuTreeIcon     string
	Router           string
}

//ToMap 将Menus([]*Menu)转化为MenuID->*Menu的映射
func (menus Menus) ToMap() map[string]*Menu {
	mapMenus := make(map[string]*Menu)
	for _, menu := range menus {
		mapMenus[menu.MenuID] = menu
	}
	return mapMenus
}

//SplitAndGetAllRecordIDs 拆分父级路径并获取所有菜单ID
func (menus Menus) SplitAndGetAllRecordIDs() []string {
	var menuIDs []string
	mapIDs := make(map[string]bool)
	for _, menu := range menus {
		mapIDs[menu.MenuID] = true
		if menu.MenuParentPath == "" {
			continue
		}

		parpath := strings.Split(menu.MenuParentPath, "/")
		for _, path := range parpath {
			if _, ok := mapIDs[path]; !ok {
				mapIDs[path] = true
			}
		}
	}

	for k := range mapIDs {
		menuIDs = append(menuIDs, k)
	}
	return menuIDs
}

func (menus Menus) ToTrees() {

}
