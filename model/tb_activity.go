package model

//AcActivity of the model
type AcActivity struct {
	ID            int64  `json:"id"`                //主键自增长ID
	CreateAt      int    `json:"create_at"`         //创建时间
	CreatedBy     int64  `json:"created_by"`        //创建人员
	UpdateAt      int    `json:"update_at"`         //更新时间
	UpdatedBy     int64  `json:"updated_by"`        //更新人员
	OnlineTime    int    `json:"online_time" `      //上线时间
	OfflineTime   int    `json:"offline_time" `     //下线时间
	TaskBeginTime int    `json:"task_begin_time"  ` //任务开始时间
	TaskEndTime   int    `json:"task_end_time"  `   //任务结束时间
	NameZh        string `json:"name_zh" `          //中文名称
	ContentZh     string `json:"content_zh" `       //中文标题
	PicZhURL      string `json:"pic_zh_url" `       //活动海报 中文
	NameEn        string `json:"name_en" `          //英文标题
	ContentEn     string `json:"content_en" `       //英文活动内容
	PicEnURL      string `json:"pic_en_url" `       //活动海报 英文
	AcType        int    `json:"ac_type" `          //活动类型[0:其他，1:手数活动]
	Status        int    `json:"status"`            //活动状态[0:未开启，1:排期中，2:进行中，3:提前下线，4:已结束]
	PageURL       string `json:"page_url"`          //活动页面链接
	HandNum       int    `json:"hand_num" `         //活动门槛手数
	Remark        string `json:"remark"`            //备注
}

//TableName ...
func (AcActivity) TableName() string {
	return "ac_activity"
}
