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
	Actions        MenuActions
	Resources      MenuResources
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

type MenuResources []*MenuResource

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
	Hidden           int
	ParentID         string
	ParentPath       string
	Actions          MenuActions
	Resources        MenuResources
	Children         []*MenuTree
}

type MenuTrees []*MenuTree

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

//ToTrees 返回Menus的MenuTree列表
func (menus Menus) ToTrees() MenuTrees {
	menuTreesList := make(MenuTrees, len(menus))
	for i, menu := range menus {
		menuTreesList[i] = &MenuTree{
			MenuTreeID:       menu.MenuID,
			MenuTreeName:     menu.MenuName,
			MenuTreePriority: menu.MenuPriority,
			MenuTreeIcon:     menu.MenuIcon,
			Router:           menu.MenuRouter,
			Hidden:           menu.Hidden,
			ParentID:         menu.MenuParentID,
			ParentPath:       menu.MenuParentPath,
			Actions:          menu.Actions,
			Resources:        menu.Resources,
		}
	}
	return menuTreesList
}

func (menus Menus) fillLeafNodeID(menuTree []*MenuTree, leafNodeIDs *[]string) {
	for _, node := range menuTree {
		if node.Children == nil || len(node.Children) == 0 {
			*leafNodeIDs = append(*leafNodeIDs, node.MenuTreeID)
		}
		menus.fillLeafNodeID(node.Children, leafNodeIDs)
	}
}

//ToLeafNodeIDs 返回叶子节点列表
func (menus Menus) ToLeafNodeIDs() []string {
	var leafNodeIDs []string
	tree := menus.ToTrees().ToTree()
	menus.fillLeafNodeID(tree, &leafNodeIDs)
	return leafNodeIDs
}

func (menuTrees MenuTrees) ToTree() []*MenuTree {
	treemap := make(map[string]*MenuTree)
	for _, menutree := range menuTrees {
		treemap[menutree.MenuTreeID] = menutree
	}

	var mtreelist []*MenuTree
	for _, menutree := range menuTrees {
		if menutree.ParentID == "" {
			mtreelist = append(mtreelist, menutree)
			continue
		}

		if par, ok := treemap[menutree.ParentID]; ok {
			if par.Children == nil {
				var child []*MenuTree
				child = append(child, par)
				copy(par.Children, child)
				continue
			}
			par.Children = append(par.Children, menutree)
		}
	}
	return mtreelist
}

//ToMap 返回MenuResources的键值对形式(ResourceCode->*MenuResource)
func (menuRes MenuResources) ToMap() map[string]*MenuResource {
	res := make(map[string]*MenuResource)
	for _, mr := range menuRes {
		res[mr.ResourceCode] = mr
	}
	return res
}
