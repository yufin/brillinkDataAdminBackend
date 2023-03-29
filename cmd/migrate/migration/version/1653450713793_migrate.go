package version

import (
	"go-admin/cmd/migrate/migration/models"
	"gorm.io/gorm"
	"runtime"

	"go-admin/cmd/migrate/migration"
	common "go-admin/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1653450713793Test)
}

func _1653450713793Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		list := make([]models.SysConfig, 0)

		list = append(list, models.SysConfig{ConfigName: "皮肤样式", ConfigKey: "sys_index_skinName", ConfigValue: "skin-green", ConfigType: "Y", Remark: "主框架页-默认皮肤样式名称:蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow", IsFrontend: false, ConfigModule: "base"})
		list = append(list, models.SysConfig{ConfigName: "初始密码", ConfigKey: "sys_user_initPassword", ConfigValue: "123456", ConfigType: "Y", Remark: "用户管理-账号初始密码:123456", IsFrontend: false, ConfigModule: "security"})
		list = append(list, models.SysConfig{ConfigName: "侧栏主题", ConfigKey: "sys_index_sideTheme", ConfigValue: "theme-dark", ConfigType: "Y", Remark: "主框架页-侧边栏主题:深色主题theme-dark，浅色主题theme-light", IsFrontend: false, ConfigModule: "base"})
		list = append(list, models.SysConfig{ConfigName: "系统名称", ConfigKey: "sys_app_name", ConfigValue: common.System.SystemName, ConfigType: "Y", Remark: "", IsFrontend: true, ConfigModule: "base"})
		list = append(list, models.SysConfig{ConfigName: "系统logo", ConfigKey: "sys_app_logo", ConfigValue: "https://gitee.com/mydearzwj/image/raw/master/img/go-admin.png", ConfigType: "Y", Remark: "", IsFrontend: true, ConfigModule: "base"})

		list = append(list, models.SysConfig{ConfigName: "站点标志", ConfigKey: "sys_site_logo", ConfigValue: "https://doc-image.zhangwj.com/img/go-admin.png", ConfigType: "Y", Remark: "logo", IsFrontend: true, ConfigModule: "base"})
		list = append(list, models.SysConfig{ConfigName: "站点徽标", ConfigKey: "sys_site_favicon", ConfigValue: "", ConfigType: "Y", Remark: "", IsFrontend: true, ConfigModule: "base"})
		list = append(list, models.SysConfig{ConfigName: "用户头像", ConfigKey: "sys_user_avatar", ConfigValue: "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png", ConfigType: "Y", Remark: "", IsFrontend: true, ConfigModule: "base"})
		list = append(list, models.SysConfig{ConfigName: "站点描述", ConfigKey: "sys_site_desc", ConfigValue: "站点描述", ConfigType: "Y", Remark: "", IsFrontend: true, ConfigModule: "base"})

		list = append(list, models.SysConfig{ConfigName: "是否启⽤", ConfigKey: "oxs_enable", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "云存储", ConfigKey: "oxs_type", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "临时授权访问", ConfigKey: "oxs_provisional_auth", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "访问域名", ConfigKey: "oxs_access_domain", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "AccessKeyId", ConfigKey: "oxs_access_key", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "Secret", ConfigKey: "oxs_secret_key", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs", IsSecret: true})
		list = append(list, models.SysConfig{ConfigName: "Bucket", ConfigKey: "oxs_bucket", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "地区", ConfigKey: "oxs_region", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "⻆⾊ARN", ConfigKey: "obx_oss_role_arn", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "RAM⻆⾊名称", ConfigKey: "obx_oss_role_session_name", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "持续时间", ConfigKey: "oxs_duration_seconds", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "IAM主账号", ConfigKey: "oxs_obs_main_username", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "IAM子账号", ConfigKey: "oxs_obs_iam_username", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs"})
		list = append(list, models.SysConfig{ConfigName: "IAM子密码", ConfigKey: "oxs_obs_iam_password", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs", IsSecret: true})
		list = append(list, models.SysConfig{ConfigName: "端点", ConfigKey: "oxs_obs_endpoint", ConfigType: "Y", IsFrontend: false, ConfigModule: "oxs", IsSecret: true})

		list = append(list, models.SysConfig{ConfigName: "通知类型", ConfigKey: "sys_notification_type", ConfigType: "Y", IsFrontend: false, ConfigModule: "notification"})
		list = append(list, models.SysConfig{ConfigName: "Webhook", ConfigKey: "sys_wechat_webhook", ConfigType: "Y", IsFrontend: false, ConfigModule: "notification"})
		err := tx.Model(models.SysConfig{}).CreateInBatches(list, len(list)).Error
		if err != nil {
			return err
		}

		err = tx.Model(models.SysConfig{}).Where("config_key in ?", []string{"sys_index_skinName", "sys_index_sideTheme", "sys_app_name", "sys_app_logo"}).Update("config_module", "base").Error
		if err != nil {
			return err
		}
		err = tx.Model(models.SysConfig{}).Where("config_key in ?", []string{"sys_user_initPassword"}).Update("config_module", "security").Error
		if err != nil {
			return err
		}

		err = tx.Model(models.SysUser{}).Create(&models.SysUser{Username: common.System.Username, Password: common.System.Password, Avatar: "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png", NickName: "zhangwj", Phone: "13818888888", RoleId: 1, Sex: "1", Email: "1@qq.com", DeptId: 1, PostId: 1, Status: "2"}).Error
		if err != nil {
			return err
		}
		err = tx.Model(models.SysRole{}).Create(&models.SysRole{RoleKey: "admin", RoleName: "系统管理员", Status: "2", Admin: true}).Error
		if err != nil {
			return err
		}
		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
