package models

import "go-admin/common/models"

type SysMenu struct {
	MenuId     int       `json:"menuId" gorm:"primaryKey;autoIncrement"`
	MenuName   string    `json:"menuName" gorm:"size:128;"`
	Title      string    `json:"title" gorm:"size:128;"`
	Icon       string    `json:"icon" gorm:"size:128;"`
	Path       string    `json:"path" gorm:"size:128;"`
	Paths      string    `json:"paths" gorm:"size:128;"`
	MenuType   string    `json:"menuType" gorm:"size:1;"`
	Permission string    `json:"permission" gorm:"size:255;"`
	ParentId   int       `json:"parentId" gorm:"size:11;"`
	NoCache    bool      `json:"noCache" gorm:"size:8;"`
	Breadcrumb string    `json:"breadcrumb" gorm:"size:255;"`
	Component  string    `json:"component" gorm:"size:255;"`
	Sort       int       `json:"sort" gorm:"size:11;"`
	Visible    bool      `json:"visible" gorm:"size:1;"`
	IsFrame    string    `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	SysApi     []SysApi  `json:"sysApi" gorm:"many2many:sys_menu_api_rule"`
	Apis       []int     `json:"apis" gorm:"-"`
	DataScope  string    `json:"dataScope" gorm:"-"`
	Params     string    `json:"params" gorm:"-"`
	RoleId     int       `gorm:"-"`
	Children   []SysMenu `json:"children,omitempty" gorm:"-"`
	IsSelect   bool      `json:"is_select" gorm:"-"`
	models.ControlBy
	models.ModelTime
}

//children?: MenuDataItem[];
///** @name 在菜单中隐藏子节点 */
//hideChildrenInMenu?: boolean;
///** @name 在菜单中隐藏自己和子节点 */
//hideInMenu?: boolean;
///** @name 菜单的icon */
//icon?: React.ReactNode;
///** @name 自定义菜单的国际化 key */
//locale?: string | false;
///** @name 菜单的名字 */
//name?: string;
///** @name 用于标定选中的值，默认是 path */
//key?: string;
///** @name disable 菜单选项 */
//disabled?: boolean;
///** @name 路径,可以设定为网页链接 */
//path?: string;
///**
// * 当此节点被选中的时候也会选中 parentKeys 的节点
// *
// * @name 自定义父节点
// */
//parentKeys?: string[];
///** @name 隐藏自己，并且将子节点提升到与自己平级 */
//flatMenu?: boolean;
///** @name 指定外链打开形式，同a标签 */
//target?: string;
//[key: string]: any;

type Menu struct {
	MenuId     int    `json:"-" gorm:"primaryKey;autoIncrement"`
	MenuName   string `json:"menuName" gorm:"size:128;"`
	Title      string `json:"name,omitempty" gorm:"size:128;"`
	Icon       string `json:"icon,omitempty" gorm:"size:128;"`
	Path       string `json:"path,omitempty" gorm:"size:128;"`
	MenuType   string `json:"menuType" gorm:"size:1;"`
	Permission string `json:"-" gorm:"size:255;"`
	ParentId   int    `json:"-" gorm:"size:11;"`
	NoCache    bool   `json:"-" gorm:"size:8;"`
	Breadcrumb string `json:"-" gorm:"size:255;"`
	Component  string `json:"component" gorm:"size:255;"`
	Sort       int    `json:"-" gorm:"size:11;"`
	Visible    bool   `json:"hideInMenu" gorm:"size:1;"`
	IsFrame    string `json:"-" gorm:"size:1;DEFAULT:0;"`
	Children   []Menu `json:"children" gorm:"-"`
	Redirect   string `json:"redirect,omitempty" gorm:"-"`
}

type SysMenuSlice []Menu

func (x SysMenuSlice) Len() int           { return len(x) }
func (x SysMenuSlice) Less(i, j int) bool { return x[i].Sort < x[j].Sort }
func (x SysMenuSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (Menu) TableName() string {
	return "sys_menu"
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

func (e *SysMenu) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysMenu) GetId() interface{} {
	return e.MenuId
}
