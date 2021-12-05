package dto

type UserApiReq struct {
	Id int `uri:"id"`
}

func (s *UserApiReq) GetId() int {
	return s.Id
}