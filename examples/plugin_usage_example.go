// Package examples demonstrates how to use the DBv2 GORM-like API in plugins
//
// This file is for documentation purposes only and shows common usage patterns.
package examples

import (
	"database/sql"
	"time"

	shared "github.com/pikselisbusiness/common-plugin"
)

// ========================================
// MODEL DEFINITIONS
// ========================================

// Product - example of reading from existing table
type Product struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"column:name;size:255" json:"name"`
	SKU        string    `gorm:"column:sku;uniqueIndex" json:"sku"`
	Price      float64   `gorm:"column:price" json:"price"`
	CategoryID uint      `gorm:"column:category_id" json:"categoryId"`
	IsActive   bool      `gorm:"column:is_active;default:true" json:"isActive"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Product) TableName() string {
	return "pb_products"
}

// PluginOrder - example of a plugin-specific model that needs migration
type PluginOrder struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	OrderID         uint           `gorm:"column:order_id;index" json:"orderId"`
	ExternalOrderID string         `gorm:"column:external_order_id;size:100" json:"externalOrderId"`
	Status          string         `gorm:"column:status;size:50;default:'pending'" json:"status"`
	SyncedAt        sql.NullTime   `gorm:"column:synced_at" json:"syncedAt"`
	ErrorMessage    sql.NullString `gorm:"column:error_message;type:text" json:"errorMessage"`
	RetryCount      int            `gorm:"column:retry_count;default:0" json:"retryCount"`
	CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (PluginOrder) TableName() string {
	return "pb_plugin_orders"
}

// PluginSetting - example model for plugin settings
type PluginSetting struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Key       string    `gorm:"column:key;size:100;uniqueIndex" json:"key"`
	Value     string    `gorm:"column:value;type:text" json:"value"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (PluginSetting) TableName() string {
	return "pb_plugin_settings"
}

// ExamplePlugin demonstrates DBv2 usage
type ExamplePlugin struct {
	shared.PBPlugin
}

// ========================================
// MIGRATION EXAMPLES
// ========================================

// OnActivate - typical plugin activation with migrations
func (p *ExamplePlugin) OnActivate() error {
	// Run migrations on plugin activation
	if err := shared.AutoMigrate(p.DB, &PluginOrder{}, &PluginSetting{}); err != nil {
		return err
	}

	// Initialize default settings if needed
	var count int64
	p.DBv2.Table("pb_plugin_settings").Count(&count)
	if count == 0 {
		p.DBv2.Table("pb_plugin_settings").Create(map[string]interface{}{
			"key":   "sync_enabled",
			"value": "true",
		})
	}

	return nil
}

// ExampleAutoMigrate - using the helper function (recommended)
func (p *ExamplePlugin) ExampleAutoMigrate() error {
	// Migrate single model
	err := shared.AutoMigrate(p.DB, &PluginOrder{})
	if err != nil {
		return err
	}

	// Migrate multiple models at once
	err = shared.AutoMigrate(p.DB,
		&PluginOrder{},
		&PluginSetting{},
	)
	return err
}

// ExampleManualMigrate - manual migration with GetMigrateQueries
func (p *ExamplePlugin) ExampleManualMigrate() error {
	// Get fields from model
	fields, err := shared.GetMigrateQueries(&PluginOrder{})
	if err != nil {
		return err
	}

	// Get table name
	tableName := shared.GetTableNameFromModel(&PluginOrder{})

	// Run migration
	return p.DB.MigrateModel(tableName, fields)
}

// ExampleCustomMigration - migration with custom fields (when you need more control)
func (p *ExamplePlugin) ExampleCustomMigration() error {
	fields := []shared.MigrateField{
		{FieldName: "ID", FieldType: "uint", GormTag: "primaryKey"},
		{FieldName: "Name", FieldType: "string", GormTag: "column:name;size:255;not null"},
		{FieldName: "Code", FieldType: "string", GormTag: "column:code;size:50;uniqueIndex"},
		{FieldName: "Value", FieldType: "float64", GormTag: "column:value;default:0"},
		{FieldName: "IsEnabled", FieldType: "bool", GormTag: "column:is_enabled;default:true"},
		{FieldName: "CreatedAt", FieldType: "Time", GormTag: "column:created_at;autoCreateTime"},
	}

	return p.DB.MigrateModel("pb_plugin_custom_table", fields)
}

// ========================================
// CRUD EXAMPLES
// ========================================

// ExampleCRUD - basic CRUD operations
func (p *ExamplePlugin) ExampleCRUD() {
	// CREATE - Insert new record
	product := Product{
		Name:       "New Product",
		SKU:        "SKU-001",
		Price:      99.99,
		CategoryID: 1,
		IsActive:   true,
	}
	result := p.DBv2.Table("pb_products").Create(&product)
	if result.Error != nil {
		// handle error
	}

	// READ - Find multiple records
	var products []Product
	p.DBv2.Table("pb_products").
		Where("is_active = ?", true).
		Order("created_at", shared.DESC).
		Limit(10).
		Find(&products)

	// READ - Find single record
	var singleProduct Product
	p.DBv2.Table("pb_products").
		Where("id = ?", 1).
		First(&singleProduct)

	// UPDATE - Update specific fields
	p.DBv2.Table("pb_products").
		Where("id = ?", 1).
		Update("price", 149.99)

	// UPDATE - Update multiple fields
	p.DBv2.Table("pb_products").
		Where("id = ?", 1).
		Updates(map[string]interface{}{
			"name":  "Updated Name",
			"price": 199.99,
		})

	// DELETE
	p.DBv2.Table("pb_products").
		Where("id = ?", 1).
		Delete(nil)
}

// ========================================
// QUERY EXAMPLES
// ========================================

// ExampleComplexQueries - joins, grouping, aggregations
func (p *ExamplePlugin) ExampleComplexQueries() {
	type ProductWithCategory struct {
		ProductID    uint    `json:"product_id"`
		ProductName  string  `json:"product_name"`
		CategoryName string  `json:"category_name"`
		Price        float64 `json:"price"`
	}

	var results []ProductWithCategory

	// Join query
	p.DBv2.Table("pb_products p").
		Select("p.id as product_id", "p.name as product_name", "c.name as category_name", "p.price").
		LeftJoin("pb_categories c", "c.id = p.category_id").
		Where("p.is_active = ?", true).
		Where("p.price > ?", 50.0).
		Order("p.price", shared.DESC).
		Find(&results)

	// Group by with having
	type CategoryStats struct {
		CategoryID   uint    `json:"category_id"`
		ProductCount int64   `json:"product_count"`
		AvgPrice     float64 `json:"avg_price"`
	}

	var stats []CategoryStats
	p.DBv2.Table("pb_products").
		Select("category_id", "COUNT(*) as product_count", "AVG(price) as avg_price").
		Where("is_active = ?", true).
		Group("category_id").
		Having("COUNT(*) > ?", 5).
		Find(&stats)
}

// ExamplePagination - paginated queries
func (p *ExamplePlugin) ExamplePagination(page, pageSize int) ([]Product, int64) {
	var products []Product
	var total int64

	// Get total count
	p.DBv2.Table("pb_products").
		Where("is_active = ?", true).
		Count(&total)

	// Get paginated results
	offset := (page - 1) * pageSize
	p.DBv2.Table("pb_products").
		Where("is_active = ?", true).
		Order("id", shared.ASC).
		Limit(pageSize).
		Offset(offset).
		Find(&products)

	return products, total
}

// ExampleScopes - reusable query logic
func (p *ExamplePlugin) ExampleScopes() {
	// Define reusable scopes
	activeScope := func(q shared.QueryBuilder) shared.QueryBuilder {
		return q.Where("is_active = ?", true)
	}

	priceRangeScope := func(min, max float64) func(shared.QueryBuilder) shared.QueryBuilder {
		return func(q shared.QueryBuilder) shared.QueryBuilder {
			return q.Where("price >= ?", min).Where("price <= ?", max)
		}
	}

	var products []Product
	p.DBv2.Table("pb_products").
		Scopes(activeScope, priceRangeScope(10.0, 100.0)).
		Find(&products)
}

// ExampleRawSQL - raw SQL when needed
func (p *ExamplePlugin) ExampleRawSQL() {
	// Use the basic DB interface for raw queries
	rows, err := p.DB.Raw(`
		SELECT 
			p.id,
			p.name,
			(SELECT COUNT(*) FROM pb_order_lines ol WHERE ol.product_id = p.id) as order_count
		FROM pb_products p
		WHERE p.is_active = 1
		ORDER BY order_count DESC
		LIMIT 10
	`)
	if err != nil {
		// handle error
	}
	_ = rows
}

// ExampleModelUsage - using Model() instead of Table()
func (p *ExamplePlugin) ExampleModelUsage() {
	// When your struct implements TableName(), you can use Model()
	var products []Product
	p.DBv2.Model(&Product{}).
		Where("is_active = ?", true).
		Find(&products)

	// Save uses Model internally
	product := Product{Name: "Test", SKU: "TEST-001", Price: 10.0}
	p.DBv2.Model(&product).Save(&product)
}

// ExampleDebug - debug mode for logging SQL
func (p *ExamplePlugin) ExampleDebug() {
	var products []Product
	// Debug() will log the SQL query
	p.DBv2.Table("pb_products").
		Debug().
		Where("is_active = ?", true).
		Find(&products)
}
