package common

type CatalogPath struct {
	IsMain bool `path:"/catalog"`                // show categories
	IsCat  bool `path:"/catalog/:CatId"`         // show items of category
	IsItem bool `path:"/catalog/:CatId/:ItemId"` // show item
	CatId  string
	ItemId int
	Page   *int `query:"page"` // query param for pagination (used pointer to avoid showing 0 default value)
}

// prev one, keep it
type HomePath struct {
	Main bool `path:"/"`
}
