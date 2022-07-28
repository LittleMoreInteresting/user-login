// Code generated by goctl. DO NOT EDIT.
package types

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type UserInfoResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TagResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	State int64  `json:"state,default=1"`
}

type TagListRequest struct {
	Name string `json:"name"`
	Pager
}

type TagListResponse struct {
	List     []TagResponse `json:"list"`
	Matedata Pager         `json:"matedata"`
}

type Pager struct {
	Page      int `json:"page,default=1"`
	PageSize  int `json:"page_size,default=10"`
	TotalSize int `json:"total_size"`
}

func (p *Pager) Offsite() int {
	return (p.Page - 1) * p.PageSize
}
