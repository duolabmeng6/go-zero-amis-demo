// Code generated by goctl. DO NOT EDIT.
package types

type ArticleItems struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ArticleIndexRequest struct {
	Page     int64  `form:"page"`
	PerPage  int64  `form:"perPage"`
	Keywords string `form:"keywords,optional"`
}

type ArticleIndexResponse struct {
	Items []ArticleItems `json:"items"`
	Total int64          `json:"total"`
}

type ArticleDestroyRequest struct {
	Id int64 `path:"id"`
}

type ArticleDestroyResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
}

type ArticleEditRequest struct {
	Id int64 `path:"id"`
}

type ArticleEditResponse struct {
	Status int64        `json:"status"`
	Msg    string       `json:"msg"`
	Data   ArticleItems `json:"data"`
}

type ArticleUpdateRequest struct {
	Id        int64  `path:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ArticleUpdateResponse struct {
	Status int64        `json:"status"`
	Msg    string       `json:"msg"`
	Data   ArticleItems `json:"data"`
}

type ArticleStoreRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleStoreResponse struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ArticleBulkDeleteRequest struct {
	Ids string `path:"ids"`
}

type ArticleBulkDeleteResponse struct {
	Status int64  `json:"status"`
	Msg    string `json:"msg"`
}
