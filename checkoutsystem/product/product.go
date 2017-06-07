//Package product helps to add and delete products available in the market
package product

//PSchema provides Layout for a product entry to be added in the inventry
type PSchema struct {
	ProductCode string  //code for a product
	Name        string  //name of the product
	Price       float64 //price of the product
}

//Products Global handle for list of currently available products in the market
var Products map[string]*PSchema

//addProduct adds the new product in the markets inventry
func addProduct(code, name string, price float64) {
	item := new(PSchema)
	item.ProductCode = code
	item.Name = name
	item.Price = price
	Products[code] = item
}

//deleteProduct deletes the product from the inventory.
func deleteProduct(code string) {
	delete(Products, code)
}

func init() {
	Products = make(map[string]*PSchema)
	addProduct("CH1", "Chai", 3.11)
	addProduct("AP1", "Apples", 6.00)
	addProduct("CF1", "Cofee", 11.23)
	addProduct("MK1", "Milk", 4.75)
	addProduct("OM1", "Oatmeal", 3.69)
}
