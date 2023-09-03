package Api

const ProductsUrl = "api.php/v1/products"

type ProductsSet struct {
	Page     int       `json:"page"`
	Total    int       `json:"total"`
	Limit    int       `json:"limit"`
	Products []Product `json:"products"`
}

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetProductsSet(productsSet *ProductsSet) error {
	if err := Get(URL+ProductsUrl, productsSet); err != nil {
		return err
	}
	return nil
}

func GetAllProduct(allProduct *[]Product) error {
	var productsSet ProductsSet
	if err := GetProductsSet(&productsSet); err != nil {
		return err
	}
	*allProduct = productsSet.Products
	return nil
}
