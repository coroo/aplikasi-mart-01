package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github/coroo/aplikasi-mart-01/models"
	"time"

	guuid "github.com/google/uuid"
)

type IOrderRepository interface {
	Insert(orderWithProduct models.OrderWithProduct) (*models.OrderWithProduct, error)
	FindOneById(id string) (*models.Order, error)
	DeleteOneById(id string) (*models.Order, error)

	FindDailyReportOder() (*models.OrderSummary, error)
	FindMonthlyReportOder() (*models.OrderSummary, error)
	FindReportOrder() (*models.OrderSummary, error)
}

var (
	orderQueries = map[string]string{
		"deleteOneById":             "DELETE FROM orders WHERE id=?",
		"orderFindOneById":          "select id,product_code,admin_name,order_quantity,total_price from orders where id=?",
		"orderReportAll":            "select sum(order_quantity),sum(total_price) from orders",
		"orderReportAllByMonthYear": "select sum(order_quantity),sum(total_price) from orders WHERE month(created_at)=? AND year(created_at)=?",
		"orderReportAllByToday":     "select sum(order_quantity),sum(total_price) from orders WHERE created_at>=?",
		"insertOrder":               "insert into orders(id,product_code,admin_name,order_quantity,total_price) values(?,?,?,?,?)",
	}
)

type OrderRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func NewOrderRepository(db *sql.DB) IOrderRepository {
	ps := make(map[string]*sql.Stmt, len(orderQueries))
	for n, v := range orderQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &OrderRepository{
		db, ps,
	}
}

func (r *OrderRepository) Insert(orderWithProduct models.OrderWithProduct) (*models.OrderWithProduct, error) {
	productRow := r.db.QueryRow("select count(id), sum(product_price) from products where product_code=?", orderWithProduct.Order.ProductCode)
	checkProduct := new(models.TotalProduct)
	err := productRow.Scan(&checkProduct.Count, &checkProduct.ProductPrice)
	if err != nil {
		return nil, err
	}
	if checkProduct.Count < 1 {
		fmt.Println("\nErr: Product code not found")
		return nil, nil
	}

	orderId := guuid.New().String()
	calculateTotalPrice := checkProduct.ProductPrice * orderWithProduct.OrderQuantity
	res, err := r.ps["insertOrder"].Exec(orderId, orderWithProduct.Order.ProductCode, orderWithProduct.AdminName, orderWithProduct.OrderQuantity, calculateTotalPrice)
	if err != nil {
		return nil, err
	}

	affectedNo, err := res.RowsAffected()
	if err != nil || affectedNo == 0 {
		return nil, errors.New(fmt.Sprintf("%s:%v", "Insert failed", err))
	}
	return &orderWithProduct, nil
}

func (r *OrderRepository) FindOneById(id string) (*models.Order, error) {
	row := r.ps["orderFindOneById"].QueryRow(id)
	res := new(models.Order)
	err := row.Scan(&res.Id, &res.ProductCode, &res.AdminName, &res.OrderQuantity, &res.TotalPrice)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *OrderRepository) DeleteOneById(id string) (*models.Order, error) {
	row := r.ps["deleteOneById"].QueryRow(id)
	res := new(models.Order)
	err := row.Scan(&res.Id, &res.ProductCode, &res.AdminName, &res.OrderQuantity, &res.TotalPrice)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *OrderRepository) FindDailyReportOder() (*models.OrderSummary, error) {
	currentTime := time.Now()
	row := r.ps["orderReportAllByToday"].QueryRow(currentTime.Format("2006-01-02"))
	res := new(models.OrderSummary)
	err := row.Scan(&res.OrderCount, &res.OrderTotal)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *OrderRepository) FindMonthlyReportOder() (*models.OrderSummary, error) {
	currentTime := time.Now()
	row := r.ps["orderReportAllByMonthYear"].QueryRow(currentTime.Month(), currentTime.Year())
	res := new(models.OrderSummary)
	err := row.Scan(&res.OrderCount, &res.OrderTotal)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *OrderRepository) FindReportOrder() (*models.OrderSummary, error) {
	row := r.ps["orderReportAllByMonthYear"].QueryRow()
	res := new(models.OrderSummary)
	err := row.Scan(&res.OrderCount, &res.OrderTotal)
	if err != nil {
		return res, err
	}
	return res, nil
}
