package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func isValidOrderStatus(status string) bool {
	return status == pendingOrderStatus || status == completedOrderStatus || status == cancelledOrderStatus
}

func addOrder(db *gorm.DB, order *Order) error {
	if order == nil {
		return errors.New("order is required")
	}
	if order.Quantity <= 0 {
		return errors.New("order quantity must be greater than zero")
	}
	if order.Status == "" {
		order.Status = pendingOrderStatus
	}
	if !isValidOrderStatus(order.Status) {
		return errors.New("status must be Pending, Completed or Cancelled")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&Customer{}, order.CustomerID).Error; err != nil {
			return fmt.Errorf("customer not found: %w", err)
		}

		var product Product
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, order.ProductID).Error; err != nil {
			return fmt.Errorf("product not found: %w", err)
		}
		if product.Quantity < order.Quantity {
			return fmt.Errorf("insufficient stock: available %d, requested %d", product.Quantity, order.Quantity)
		}

		order.TotalPrice = product.Price * float64(order.Quantity)
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		return tx.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", order.Quantity)).Error
	})
}

func fetchOrdersForCustomer(db *gorm.DB, customerID uint) ([]Order, error) {
	var orders []Order
	err := db.Preload("Customer").Preload("Product").Where("customer_id = ?", customerID).Order("created_at DESC").Find(&orders).Error
	return orders, err
}

func fetchOrdersByStatus(db *gorm.DB, status string) ([]Order, error) {
	status = strings.TrimSpace(status)
	if !isValidOrderStatus(status) {
		return nil, errors.New("status must be Pending, Completed or Cancelled")
	}
	var orders []Order
	err := db.Preload("Customer").Preload("Product").Where("status = ?", status).Order("created_at DESC").Find(&orders).Error
	return orders, err
}

func fetchOrdersInRange(db *gorm.DB, startDate, endDate time.Time) ([]Order, error) {
	if endDate.Before(startDate) {
		return nil, errors.New("end date must not be before start date")
	}
	var orders []Order
	err := db.Preload("Customer").Preload("Product").Where("created_at BETWEEN ? AND ?", startDate, endDate).Order("created_at DESC").Find(&orders).Error
	return orders, err
}

func fetchOrdersWithDetails(db *gorm.DB) ([]Order, error) {
	var orders []Order
	err := db.Preload("Customer").Preload("Product").Preload("Product.Category").Order("created_at DESC").Find(&orders).Error
	return orders, err
}

func changeOrderStatus(db *gorm.DB, id uint, status string) error {
	status = strings.TrimSpace(status)
	if !isValidOrderStatus(status) {
		return errors.New("status must be Pending, Completed or Cancelled")
	}
	result := db.Model(&Order{}).Where("id = ?", id).Update("status", status)
	return handleRowsAffected(result, "order", id)
}
