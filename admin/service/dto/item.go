package dto

type ItemUpdateReq struct {
	ID       int    `json:"ID" comment:"id"`        // id
	Name    string `json:"Title" comment:"名称"`     //标题
	BLink string `json:"SubTitle" comment:"B站链接"` //副标题
	Preview  string `json:"Preview" comment:"预览图"`   //主图
	Desc     string `json:"Desc" comment:"描述"`      //描述
	Author   string `json:"Author" comment:"作者"`    //状态
	Price   string `json:"Price" comment:"价格"`    //状态
	Status   string `json:"Status" comment:"状态"`    //状态
}
