package fasthttp

// Book type
type Book struct {
	ID 		string  `json:"id"`
	Title 	string  `json:"title"`
	Price	float64 `json:"price"` 
}

// Response type
type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Book   	*Book  `json:"book"`
}