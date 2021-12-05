package dto


type CatApiReq struct {
	Id int `uri:"id"`
}

func (s *CatApiReq) GetId() int {
	return s.Id
}
