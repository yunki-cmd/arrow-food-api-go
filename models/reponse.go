package models

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Mensage string      `json:"mensage,omitempty"`
	Status  int         `json:"status"`
}

func (r *Response) Contructor(data interface{}, mensaje string, status int) {
	r.Data = data
	r.Mensage = mensaje
	r.Status = status
}
