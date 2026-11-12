package models

const (
	TitleField       = "title"
	DescriptionField = "description"
	PriceField       = "price"
	QuantityField    = "quantity"
	PhotosField      = "photos"
)

type CreateRequest struct {
	Fields     map[string]string
	PhotoPaths []string
}
