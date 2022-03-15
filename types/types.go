package types

type SuccessResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type GetCategoryResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type GetCategoriesParams struct {
	Limit  int
	Offset int
	Title  string
}

type GetDealsParams struct {
	Limit  int
	Offset int
}

type GetProductResponse struct {
	Id         int      `json:"id"`
	Title      string   `json:"title"`
	Price      float32  `json:"price"`
	CategoryId int      `json:"categoryId"`
	Pictures   []string `json:"pictures"`
}

type GetProductsParams struct {
	Limit      int
	Offset     int
	CategoryId int
	Title      string
}

type AdminExistResponse struct {
	Sucess bool `json:"success"`
}

type AdminLoginResponse struct {
	Sucess bool `json:"success"`
}

type CreateAdminResponse struct {
	Qrcode string `json:"qrcode"`
}

type PostDealResponse struct {
	Qrcode string `json:"qrcode"`
}

type GetCartResponse struct {
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}

type GetDealResponse struct {
	Id        int     `json:"id"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
	Datestamp string  `json:"datestamp"`
	CartId    int     `json:"cartId"`
}

type PostItemRequest struct {
	Id        int `json:"id"`
	ProductId int `json:"productId"`
	CartId    int `json:"cartId"`
	Quantity  int `json:"quantity"`
}

type GetItemResponse struct {
	Id        int      `json:"id"`
	ProductId int      `json:"productId"`
	Title     string   `json:"title"`
	Price     float32  `json:"price"`
	Pictures  []string `json:"pictures"`
	Quantity  int      `json:"quantity"`
}

type GetItemsParams struct {
	Title      string
	Offset     int
	Limit      int
	CategoryId int
	CartId     int
}
