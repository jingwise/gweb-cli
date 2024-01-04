package response

type Response struct {
	Code int    `json:"errcode"`
	Msg  string `json:"msg,omitempty"`
}
type response struct {
	Response
	Data any `json:"data"`
}

func (e *response) SetMsg(msg string) {
	e.Msg = msg
}

func (e *response) SetCode(code int) {
	e.Code = code
}

type Page struct {
	Count     int `json:"count"`
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

type page struct {
	Page
	List any `json:"list"`
}

func (e *response) SetData(data any) {
	e.Data = data
}

func (e response) Clone() *response {
	return &e
}
