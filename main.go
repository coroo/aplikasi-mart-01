package main

import (
	"database/sql"
	"fmt"
	"github/coroo/aplikasi-mart-01/config"
	"github/coroo/aplikasi-mart-01/deliveries"
	"github/coroo/aplikasi-mart-01/models"
	"github/coroo/aplikasi-mart-01/repositories"
	"github/coroo/aplikasi-mart-01/usecases"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type app struct {
	db *sql.DB
}
type menuChoosed string

func newApp() app {
	c := config.NewConfig()
	err := c.InitDb()
	if err != nil {
		panic(err)
	}
	myapp := app{
		db: c.Db,
	}
	return myapp
}

func MainMenuForm() {
	var appMenu = map[string]string{
		"01": "Produk barang",
		"02": "Transaksi penjualan",
		"03": "Laporan penjualan",
		"q":  "Exit",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Main Menu Application")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range MenuChoiceOrdered(appMenu) {
		fmt.Printf("%s. %s\n", menuCode, appMenu[menuCode])
	}
}
func SubMainMenuForm(main menuChoosed) {
	var menuName string
	var subAppMenu map[string]string
	subAppMenu = map[string]string{}

	if main == "01" {
		menuName = "Product barang"
		subAppMenu["A"] = "Tambah produk"
		subAppMenu["B"] = "Hapus produk"
		subAppMenu["C"] = "Detail produk"
		subAppMenu["q"] = "Back to Main Menu"
	} else if main == "02" {
		menuName = "Transaksi Penjualan"
		subAppMenu["A"] = "Tambah transaksi penjualan"
		subAppMenu["B"] = "Hapus transaksi penjualan"
		subAppMenu["C"] = "Detail transaksi penjualan"
		subAppMenu["q"] = "Back to Main Menu"
	} else if main == "03" {
		menuName = "Laporan Penjualan"
		subAppMenu["A"] = "Laporan harian"
		subAppMenu["B"] = "Laporan bulanan"
		subAppMenu["C"] = "Semua laporan"
		subAppMenu["q"] = "Back to Main Menu"
	} else {
		menuName = "Unknown Menu"
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", menuName)
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range MenuChoiceOrdered(subAppMenu) {
		fmt.Printf("%s. %s\n", menuCode, subAppMenu[menuCode])
	}
}
func MenuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
func (a app) Run() {
	var isExist = false
	var userChoice string

	MainMenuForm()

	for isExist == false {
		fmt.Printf("\n%s", "Your Choice: ")
		fmt.Scan(&userChoice)
		switch {
		case userChoice == "01":
			var isBack = false
			var userChoice string

			SubMainMenuForm("01")
			repo := repositories.NewProductRepository(a.db)
			usecase := usecases.NewProductUseCase(repo)
			productDelivery := deliveries.NewProductDelivery(usecase)

			for isBack == false {
				fmt.Printf("\n%s", "Your Next Choice: ")
				fmt.Scan(&userChoice)
				switch {
				case userChoice == "A":
					var productCode string
					var productName string
					var productPrice int
					fmt.Printf("\n%s", "Product Code: ")
					fmt.Scan(&productCode)
					fmt.Printf("%s", "Product Name: ")
					fmt.Scan(&productName)
					fmt.Printf("%s", "Product Price: ")
					fmt.Scan(&productPrice)
					newProduct, err := productDelivery.RegisterNewProduct(models.Product{
						ProductCode:  productCode,
						ProductName:  productName,
						ProductPrice: productPrice,
					})
					if err != nil {
						panic(err)
					}
					productDelivery.PrintOneProduct(newProduct)
					isBack = true
					isExist = true
				case userChoice == "B":
					fmt.Printf("\n%s", "Delete Product")
					var productId string
					fmt.Printf("\n%s", "Product id: ")
					fmt.Scan(&productId)
					result, _ := productDelivery.DeleteProductById(productId)
					productDelivery.PrintOneProduct(result)
					isBack = true
					isExist = true
				case userChoice == "C":
					fmt.Printf("\n%s", "Detail Product")
					var productId string
					fmt.Printf("\n%s", "Product id: ")
					fmt.Scan(&productId)
					result, _ := productDelivery.GetProductById(productId)
					productDelivery.PrintOneProduct(result)
					isBack = true
					isExist = true
				case userChoice == "q":
					isBack = true
					MainMenuForm()
				default:
					fmt.Println("Unknown Menu Code")
				}
			}
		case userChoice == "02":
			var isBack = false
			var userChoice string

			SubMainMenuForm("02")
			repo := repositories.NewOrderRepository(a.db)
			usecase := usecases.NewOrderUseCase(repo)
			orderDelivery := deliveries.NewOrderDelivery(usecase)

			for isBack == false {
				fmt.Printf("\n%s", "Your Next Choice: ")
				fmt.Scan(&userChoice)
				switch {
				case userChoice == "A":
					var productCode string
					var adminName string
					var orderQuantity int
					fmt.Printf("\n%s", "Product Code: ")
					fmt.Scan(&productCode)
					fmt.Printf("%s", "Admin Name: ")
					fmt.Scan(&adminName)
					fmt.Printf("%s", "Quantity: ")
					fmt.Scan(&orderQuantity)
					_, err := orderDelivery.RegisterNewOrder(models.OrderWithProduct{
						Order: models.Order{
							ProductCode:   productCode,
							AdminName:     adminName,
							OrderQuantity: orderQuantity,
						},
					})
					if err != nil {
						panic(err)
					}
					isBack = true
					isExist = true
				case userChoice == "B":
					fmt.Printf("\n%s", "Delete Order")
					var orderId string
					fmt.Printf("\n%s", "Order id: ")
					fmt.Scan(&orderId)
					result, _ := orderDelivery.DeleteOrderById(orderId)
					orderDelivery.PrintOneOrder(result)
					isBack = true
					isExist = true
				case userChoice == "C":
					fmt.Printf("\n%s", "Detail Product")
					var orderId string
					fmt.Printf("\n%s", "Product id: ")
					fmt.Scan(&orderId)
					result, _ := orderDelivery.GetOrderById(orderId)
					orderDelivery.PrintOneOrder(result)
					isBack = true
					isExist = true
				case userChoice == "q":
					isBack = true
					MainMenuForm()
				default:
					fmt.Println("Unknown Menu Code")
				}
			}
		case userChoice == "03":
			var isBack = false
			var userChoice string

			SubMainMenuForm("03")
			repo := repositories.NewOrderRepository(a.db)
			usecase := usecases.NewOrderSummaryUseCase(repo)
			orderSummaryDelivery := deliveries.NewOrderSummaryDelivery(usecase)

			for isBack == false {
				fmt.Printf("\n%s", "Your Next Choice: ")
				fmt.Scan(&userChoice)
				switch {
				case userChoice == "A":
					result, _ := orderSummaryDelivery.GetDailyReport()
					orderSummaryDelivery.PrintOneOrder(result)
					isBack = true
					isExist = true
				case userChoice == "B":
					result, _ := orderSummaryDelivery.GetMonthlyReport()
					orderSummaryDelivery.PrintOneOrder(result)
					isBack = true
					isExist = true
				case userChoice == "C":
					result, _ := orderSummaryDelivery.GetAllReport()
					orderSummaryDelivery.PrintOneOrder(result)
					isBack = true
					isExist = true
				case userChoice == "q":
					isBack = true
					MainMenuForm()
				default:
					fmt.Println("Unknown Menu Code")
				}
			}
		case userChoice == "q":
			isExist = true
		default:
			fmt.Println("Unknown Menu Code")

		}
	}
}

func main() {
	newApp().Run()
}
