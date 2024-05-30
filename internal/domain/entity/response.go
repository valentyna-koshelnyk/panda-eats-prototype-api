package entity

type Response struct {
	APIVersion string `json:"apiVersion"`
	Data       Data   `json:"data"`
}

type Data struct {
	StartIndex   int    `json:"start_index"`
	ItemsCount   int    `json:"items_count"`
	ItemsPerPage int    `json:"items_per_page"`
	Items        []Item `json:"items"`
}

type Item interface{}

const (
	DefaultAPIVersion = "1.0"
	DefaultStartIndex = 1
)

func NewResponse(items []Item) *Response {
	itemsCount := len(items)
	return &Response{
		APIVersion: DefaultAPIVersion,
		Data: Data{
			StartIndex:   DefaultStartIndex,
			ItemsCount:   itemsCount,
			ItemsPerPage: itemsCount,
			Items:        items,
		},
	}
}
