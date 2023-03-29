package version

import (
	"gorm.io/gorm"
	"runtime"
    "time"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"go-admin/cmd/migrate/migration"
	common "go-admin/common/models"
)

type Menu struct {
	MenuId     int    `json:"menuId" gorm:"primaryKey;autoIncrement"`
	MenuName   string `json:"menuName" gorm:"size:128;"`
	Title      string `json:"title" gorm:"size:128;"`
	Icon       string `json:"icon" gorm:"size:128;"`
	Path       string `json:"path" gorm:"size:128;"`
	Paths      string `json:"paths" gorm:"size:128;"`
	MenuType   string `json:"menuType" gorm:"size:1;"`
	Action     string `json:"action" gorm:"size:16;"`
	Permission string `json:"permission" gorm:"size:255;"`
	ParentId   int    `json:"parentId" gorm:"size:11;"`
	NoCache    bool   `json:"noCache" gorm:"size:8;"`
	Breadcrumb string `json:"breadcrumb" gorm:"size:255;"`
	Component  string `json:"component" gorm:"size:255;"`
	Sort       int    `json:"sort" gorm:"size:11;"`
	Visible    string `json:"visible" gorm:"size:1;"`
	CreateBy   string `json:"createBy" gorm:"size:128;"`
	UpdateBy   string `json:"updateBy" gorm:"size:128;"`
	IsFrame    string `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	CreatedAt time.Time  `json:"createdAt"`
    UpdatedAt time.Time  `json:"updatedAt"`
    DeletedAt *time.Time `json:"deletedAt"`
}

func (Menu) TableName() string {
	return "sys_menu"
}

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _{{.GenerateTime}}Test)
}

func _{{.GenerateTime}}Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
        timeNow := pkg.GetCurrentTime()
        m := Menu{}
        m.MenuName = "{{.TBName}}Manage"
        m.Title = "{{.TableComment}}"
        m.Icon = "pass"
        m.Path = "/{{.TBName}}"
        m.MenuType = "M"
        m.Action = "无"
        m.ParentId = 0
        m.NoCache = false
        m.Component = "Layout"
        m.Sort = 0
        m.Visible = "0"
        m.IsFrame = "0"
        m.CreateBy = "1"
        m.UpdateBy = "1"
        m.CreatedAt = timeNow
        m.UpdatedAt = timeNow
        err := tx.Create(&m).Error
        if err != nil {
            return err
        }
        c := Menu{}
        c.MenuName = "{{.TBName}}"
        c.Title = "{{.TableComment}}"
        c.Icon = "pass"
        c.Path = "{{.TBName}}"
        c.MenuType = "C"
        c.Action = "无"
        c.Permission = "{{.PackageName}}:{{.BusinessName}}:list"
        c.ParentId = Mmenu.MenuId
        c.NoCache = false
        c.Component = "/{{.BusinessName}}/index"
        c.Sort = 0
        c.Visible = "0"
        c.IsFrame = "0"
        c.CreateBy = "1"
        c.UpdateBy = "1"
        c.CreatedAt = timeNow
        c.UpdatedAt = timeNow
        err = tx.Create(&c).Error
        if err != nil {
            return err
        }

        list := Menu{}
        list.MenuName = ""
        list.Title = "分页获取{{.TableComment}}"
        list.Icon = ""
        list.Path = "{{.TBName}}"
        list.MenuType = "F"
        list.Action = "无"
        list.Permission = "{{.PackageName}}:{{.BusinessName}}:query"
        list.ParentId = c.MenuId
        list.NoCache = false
        list.Sort = 0
        list.Visible = "0"
        list.IsFrame = "0"
        list.CreateBy = "1"
        list.UpdateBy = "1"
        list.CreatedAt = timeNow
        list.UpdatedAt = timeNow
        err = tx.Create(&list).Error
        if err != nil {
            return err
        }

        create := Menu{}
        create.MenuName = ""
        create.Title = "创建{{.TableComment}}"
        create.Icon = ""
        create.Path = "{{.TBName}}"
        create.MenuType = "F"
        create.Action = "无"
        create.Permission = "{{.PackageName}}:{{.BusinessName}}:add"
        create.ParentId = c.MenuId
        create.NoCache = false
        create.Sort = 0
        create.Visible = "0"
        create.IsFrame = "0"
        create.CreateBy = "1"
        create.UpdateBy = "1"
        create.CreatedAt = timeNow
        create.UpdatedAt = timeNow
        err = tx.Create(&create).Error
        if err != nil {
            return err
        }

        update := Menu{}
        update.MenuName = ""
        update.Title = "修改{{.TableComment}}"
        update.Icon = ""
        update.Path = "{{.TBName}}"
        update.MenuType = "F"
        update.Action = "无"
        update.Permission ="{{.PackageName}}:{{.BusinessName}}:edit"
        update.ParentId = c.MenuId
        update.NoCache = false
        update.Sort = 0
        update.Visible = "0"
        update.IsFrame = "0"
        update.CreateBy = "1"
        update.UpdateBy = "1"
        update.CreatedAt = timeNow
        update.UpdatedAt = timeNow
        err = tx.Create(&update).Error
        if err != nil {
            return err
        }

        delete := Menu{}
        delete.MenuName = ""
        delete.Title = "删除{{.TableComment}}"
        delete.Icon = ""
        delete.Path = "{{.TBName}}"
        delete.MenuType = "F"
        delete.Action = "无"
        delete.Permission = "{{.PackageName}}:{{.BusinessName}}:remove"
        delete.ParentId = c.MenuId
        delete.NoCache = false
        delete.Sort = 0
        delete.Visible = "0"
        delete.IsFrame = "0"
        delete.CreateBy = "1"
        delete.UpdateBy = "1"
        delete.CreatedAt = timeNow
        delete.UpdatedAt = timeNow
        err = tx.Create(&delete).Error
        if err != nil {
            return err
        }

        var InterfaceId = 63
        a := Menu{}
        a.MenuName = "{{.TBName}}"
        a.Title = "{{.TableComment}}"
        a.Icon = "bug"
        a.Path = "{{.TBName}}"
        a.MenuType = "M"
        a.Action = "无"
        a.ParentId = InterfaceId
        a.NoCache = false
        a.Sort = 0
        a.Visible = "1"
        a.IsFrame = "0"
        a.CreateBy = "1"
        a.UpdateBy = "1"
        a.CreatedAt = timeNow
        a.UpdatedAt = timeNow
        err = tx.Create(&a).Error
        if err != nil {
            return err
        }

        AList := Menu{}
        AList.MenuName = ""
        AList.Title = "分页获取{{.TableComment}}"
        AList.Icon = "bug"
        AList.Path = "/api/v1/{{.ModuleName}}"
        AList.MenuType = "A"
        AList.Action = "GET"
        AList.ParentId = a.MenuId
        AList.NoCache = false
        AList.Sort = 0
        AList.Visible = "1"
        AList.IsFrame = "0"
        AList.CreateBy = "1"
        AList.UpdateBy = "1"
        AList.CreatedAt = timeNow
        AList.UpdatedAt = timeNow
        err = tx.Create(&AList).Error
        if err != nil {
            return err
        }

        AGet := Menu{}
        AGet.MenuName = ""
        AGet.Title = "根据id获取{{.TableComment}}"
        AGet.Icon = "bug"
        AGet.Path = "/api/v1/{{.ModuleName}}/:id"
        AGet.MenuType = "A"
        AGet.Action = "GET"
        AGet.ParentId = a.MenuId
        AGet.NoCache = false
        AGet.Sort = 0
        AGet.Visible = "1"
        AGet.IsFrame = "0"
        AGet.CreateBy = "1"
        AGet.UpdateBy = "1"
        AGet.CreatedAt = timeNow
        AGet.UpdatedAt = timeNow
        err = tx.Create(&AGet).Error
        if err != nil {
            return err
        }

        ACreate := Menu{}
        ACreate.MenuName = ""
        ACreate.Title = "创建{{.TableComment}}"
        ACreate.Icon = "bug"
        ACreate.Path = "/api/v1/{{.ModuleName}}"
        ACreate.MenuType = "A"
        ACreate.Action = "POST"
        ACreate.ParentId = a.MenuId
        ACreate.NoCache = false
        ACreate.Sort = 0
        ACreate.Visible = "1"
        ACreate.IsFrame = "0"
        ACreate.CreateBy = "1"
        ACreate.UpdateBy = "1"
        ACreate.CreatedAt = timeNow
        ACreate.UpdatedAt = timeNow
        err = tx.Create(&ACreate).Error
        if err != nil {
            return err
        }

        AUpdate := Menu{}
        AUpdate.MenuName = ""
        AUpdate.Title = "修改{{.TableComment}}"
        AUpdate.Icon = "bug"
        AUpdate.Path = "/api/v1/{{.ModuleName}}/:id"
        AUpdate.MenuType = "A"
        AUpdate.Action = "PUT"
        AUpdate.ParentId = a.MenuId
        AUpdate.NoCache = false
        AUpdate.Sort = 0
        AUpdate.Visible = "1"
        AUpdate.IsFrame = "0"
        AUpdate.CreateBy = "1"
        AUpdate.UpdateBy = "1"
        AUpdate.CreatedAt = timeNow
        AUpdate.UpdatedAt = timeNow
        err = tx.Create(&AUpdate).Error
        if err != nil {
            return err
        }

        ADelete := Menu{}
        ADelete.MenuName = ""
        ADelete.Title = "删除{{.TableComment}}"
        ADelete.Icon = "bug"
        ADelete.Path = "/api/v1/{{.ModuleName}}"
        ADelete.MenuType = "A"
        ADelete.Action = "DELETE"
        ADelete.ParentId = a.MenuId
        ADelete.NoCache = false
        ADelete.Sort = 0
        ADelete.Visible = "1"
        ADelete.IsFrame = "0"
        ADelete.CreateBy = "1"
        ADelete.UpdateBy = "1"
        ADelete.CreatedAt = timeNow
        ADelete.UpdatedAt = timeNow
        err = tx.Create(&ADelete).Error
        if err != nil {
            return err
        }

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}