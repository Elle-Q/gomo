package response

type Responses interface {
	SetCode(int)
	SetMsg(string)
	SetData(interface{})
	SetSuccess(bool)
	Clone() Responses
}

type Response struct {
	Code   int
	Msg    string
	Status string
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

type Page struct {
	Count     int `json:"count"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"page_size"`
}

type page struct {
	Page
	List interface{} `json:"list"`
}

func (e *response) SetData(data interface{})  {
	e.Data = data
}

func (e *response) SetMsg(s string) {
	e.Msg = s
}

func (e *response) SetCode(code int) {
	e.Code = code
}

func (e *response) SetSuccess(success bool) {
	if !success {
		e.Status = "error"
	} else {
		e.Status = "ok"
	}
}

func (e response) Clone() Responses {
	return &e
}
