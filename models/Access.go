package models

type Access struct {
	Id          int      `gorm:"primary_key" json:"id"`
	ModuleName  string   `json:"module_name"` //模块名称
	ActionName  string   `json:"action_name"` //操作名称
	Type        int      `json:"type"`        //节点类型 :  1、表示模块    2、表示菜单     3、操作
	Url         string   `json:"url"`         //路由跳转地址
	ModuleId    int      `json:"module_id"`   //此module_id和当前模型的id关联       module_id= 0 表示模块
	Sort        int      `json:"sort"`
	Description string   `json:"description"`
	Status      int      `json:"status"`
	AddTime     int      `json:"add_time"`
	AccessItem  []Access `gorm:"foreignKey:ModuleId;references:Id"`
	Checked     string   `gorm:"-" json:"checked"` // 忽略本字段
}

func (Access) TableName() string {
	return "access"
}
