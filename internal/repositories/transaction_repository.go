package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/internal/models"
	"strconv"
	"strings"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		// Cari informasi produk
		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id =$1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		// Hitung subtotal dan totalAmount
		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		// Kurangi jumlah stock
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		// Item dimasukkin ke transaction details
		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING ID", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID
		var detailID int
		err := tx.QueryRow("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING ID", details[i].TransactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal).Scan(&detailID)
		if err != nil {
			return nil, err
		}
		details[i].ID = detailID
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func (repo *TransactionRepository) Summary(startDate string, endDate string) (*models.Summary, error) {
	today := time.Now()
	datePlaceholder := "2006-01-02"

	if startDate != "" {
		_, err := time.Parse(datePlaceholder, startDate)
		if err != nil {
			return nil, err
		}
	} else {
		startDate = today.Format(datePlaceholder)
	}

	if endDate != "" {
		_, err := time.Parse(datePlaceholder, endDate)
		if err != nil {
			return nil, err
		}
	} else {
		endDate = today.AddDate(0, 0, 1).Format(datePlaceholder)
	}

	// Ambil semua transaksi
	rows, err := repo.db.Query("SELECT * FROM transactions WHERE created_at::date >= $1::date AND created_at::date <= $2", startDate, endDate)
	if err == sql.ErrNoRows {
		fmt.Printf("Tidak ada transaksi pada %s", today)
		return nil, err
	}
	defer rows.Close()

	// Init
	totalTransaksi := 0
	totalRevenue := 0
	transactionIds := make([]string, 0)
	for rows.Next() {
		// Scan baris
		var tr models.Transaction
		err := rows.Scan(&tr.ID, &tr.TotalAmount, &tr.CreatedAt)
		if err != nil {
			fmt.Printf(`Scan error untuk id = %d, total_amount = %d`, tr.ID, tr.TotalAmount)
			continue
		}

		// Increment transaksi dan hitung revenue
		totalTransaksi++
		totalRevenue += tr.TotalAmount

		// Simpan id transaksi untuk query detail transaksi nanti
		transactionIds = append(transactionIds, strconv.Itoa(tr.ID))
	}

	// Buat query untuk mendapatkan produk terlaris
	// Misal produk terlalis adalah yang kuantitas penjualannya tertinggi
	var query string
	query = `
		SELECT 
			p.name as product_name, 
			SUM(quantity) as total_quantity
		FROM
			transaction_details td JOIN products p ON td.product_id = p.id
		WHERE
			transaction_id IN (` + strings.Join(transactionIds, ",") + `)
		GROUP BY
			p.name
		ORDER BY
			total_quantity DESC
		LIMIT
			1;
		`

	// Jalankan query
	var produkTerlaris models.ProdukTerlaris
	err = repo.db.QueryRow(query).Scan(&produkTerlaris.Nama, &produkTerlaris.QtyTerjual)
	if err != nil {
		return nil, err
	}

	summary := models.Summary{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: produkTerlaris,
	}

	return &summary, nil
}

func (repo *TransactionRepository) SummaryToday() (*models.Summary, error) {
	today := time.Now()

	// Ambil semua transaksi
	rows, err := repo.db.Query("SELECT * FROM transactions WHERE created_at::date >= $1::date", today)
	if err == sql.ErrNoRows {
		fmt.Printf("Tidak ada transaksi pada %s", today)
		return nil, err
	}
	defer rows.Close()

	// Init
	totalTransaksi := 0
	totalRevenue := 0
	transactionIds := make([]string, 0)
	for rows.Next() {
		// Scan baris
		var tr models.Transaction
		err := rows.Scan(&tr.ID, &tr.TotalAmount, &tr.CreatedAt)
		if err != nil {
			fmt.Printf(`Scan error untuk id = %d, total_amount = %d`, tr.ID, tr.TotalAmount)
			continue
		}

		// Increment transaksi dan hitung revenue
		totalTransaksi++
		totalRevenue += tr.TotalAmount

		// Simpan id transaksi untuk query detail transaksi nanti
		transactionIds = append(transactionIds, strconv.Itoa(tr.ID))
	}

	// Buat query untuk mendapatkan produk terlaris
	// Misal produk terlalis adalah yang kuantitas penjualannya tertinggi
	var query string
	query = `
		SELECT 
			p.name as product_name, 
			SUM(quantity) as total_quantity
		FROM
			transaction_details td JOIN products p ON td.product_id = p.id
		WHERE
			transaction_id IN (` + strings.Join(transactionIds, ",") + `)
		GROUP BY
			p.name
		ORDER BY
			total_quantity DESC
		LIMIT
			1;
		`

	// Jalankan query
	var produkTerlaris models.ProdukTerlaris
	err = repo.db.QueryRow(query).Scan(&produkTerlaris.Nama, &produkTerlaris.QtyTerjual)
	if err != nil {
		return nil, err
	}

	summary := models.Summary{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: produkTerlaris,
	}

	return &summary, nil
}
