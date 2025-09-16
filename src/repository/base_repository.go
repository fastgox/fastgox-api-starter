package repository

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/database"
	"github.com/fastgox/utils/logger"
	"gorm.io/gorm"
)

// Entity 定义实体接口，所有实体都应该有ID
type Entity interface {
	comparable
}

// RepoFactory 仓储工厂函数类型
type RepoFactory func(*gorm.DB) interface{}

// 全局仓储注册表
var repoFactories = make(map[string]RepoFactory)

// 全局数据库实例
var GlobalDB *gorm.DB

// InitGlobalDB 初始化全局数据库实例
func InitGlobalDB(db *gorm.DB) {
	GlobalDB = db
}

// initDatabase 内部初始化数据库函数
func initDatabase() {
	logger.Info("初始化数据库...")

	// 初始化数据库连接
	db, err := database.Initialize()
	if err != nil {
		logger.Error("数据库初始化失败: %v", err)
		panic(err)
	}

	// 设置全局数据库实例
	GlobalDB = db

	logger.Info("数据库初始化完成")
}

// BaseRepository 基础仓储实现，提供通用的CRUD操作
type BaseRepository[T Entity] struct {
	DB *gorm.DB
}

// NewBaseRepository 创建基础仓储实例，使用全局DB
func NewBaseRepository[T Entity]() *BaseRepository[T] {
	return &BaseRepository[T]{DB: GlobalDB}
}

// GetRepository 获取仓储实例的泛型函数
func GetRepository[T Entity]() *BaseRepository[T] {
	if GlobalDB == nil {
		initDatabase()
	}
	return &BaseRepository[T]{DB: GlobalDB}
}

// Create 创建实体
func (r *BaseRepository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

// GetByID 根据ID获取实体
func (r *BaseRepository[T]) GetByID(id int64) (*T, error) {
	var entity T
	err := r.DB.Where("id = ?", id).First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // 没找到，返回nil而不是错误
		}
		return nil, err
	}
	return &entity, nil
}

// Update 更新实体
func (r *BaseRepository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

// Delete 删除实体
func (r *BaseRepository[T]) Delete(id string) error {
	var entity T
	return r.DB.Where("id = ?", id).Delete(&entity).Error
}

// DeleteByID 根据ID删除实体（别名方法）
func (r *BaseRepository[T]) DeleteByID(id string) error {
	return r.Delete(id)
}

// Find 根据条件查询多个实体
func (r *BaseRepository[T]) Find(condition string, args ...interface{}) ([]T, error) {
	var entities []T
	query := r.DB
	if condition != "" {
		query = query.Where(condition, args...)
	}
	err := query.Find(&entities).Error
	return entities, err
}

// First 查询第一个实体
func (r *BaseRepository[T]) First(condition string, args ...interface{}) (*T, error) {
	var entity T
	query := r.DB
	if condition != "" {
		query = query.Where(condition, args...)
	}
	err := query.First(&entity).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // 没找到，返回nil而不是错误
		}
		return nil, err
	}
	return &entity, nil
}

// Count 获取总数
func (r *BaseRepository[T]) Count(condition string, args ...interface{}) (int64, error) {
	var count int64
	var entity T
	query := r.DB.Model(&entity)
	if condition != "" {
		query = query.Where(condition, args...)
	}
	err := query.Count(&count).Error
	return count, err
}

// GetDB 获取数据库实例（用于复杂查询）
func (r *BaseRepository[T]) GetDB() *gorm.DB {
	return r.DB
}

// UpdateFields 更新指定字段
func (r *BaseRepository[T]) UpdateFields(id string, fields map[string]interface{}) error {
	// 自动设置更新时间
	fields["updated_at"] = time.Now().Unix()

	var entity T
	return r.DB.Model(&entity).Where("id = ?", id).Updates(fields).Error
}

// Model 返回 GORM 的 Model 方法结果
func (r *BaseRepository[T]) Model() *gorm.DB {
	var entity T
	return r.DB.Model(&entity)
}

// Transaction 执行事务
func (r *BaseRepository[T]) Transaction(fn func(tx *gorm.DB) error) error {
	return r.DB.Transaction(fn)
}

// Limit 设置查询限制
func (r *BaseRepository[T]) Limit(limit int) *BaseRepository[T] {
	return &BaseRepository[T]{
		DB: r.DB.Limit(limit),
	}
}

// Offset 设置查询偏移
func (r *BaseRepository[T]) Offset(offset int) *BaseRepository[T] {
	return &BaseRepository[T]{
		DB: r.DB.Offset(offset),
	}
}

// Order 设置排序
func (r *BaseRepository[T]) Order(value interface{}) *BaseRepository[T] {
	return &BaseRepository[T]{
		DB: r.DB.Order(value),
	}
}

// Select 指定查询字段
func (r *BaseRepository[T]) Select(value interface{}) *BaseRepository[T] {
	return &BaseRepository[T]{
		DB: r.DB.Select(value),
	}
}

// Page 分页查询
func (r *BaseRepository[T]) Page(page, size int, condition string, args ...interface{}) ([]T, int64, error) {
	var entities []T
	var total int64

	// 添加SQL调试日志
	fmt.Printf("🔍 [DEBUG] BaseRepository.Page 查询参数: page=%d, size=%d, condition='%s', args=%v\n",
		page, size, condition, args)

	// 计算总数
	countQuery := r.DB.Model(new(T))
	if condition != "" {
		countQuery = countQuery.Where(condition, args...)
	}

	// 获取生成的SQL（调试用）
	sql := countQuery.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Count(&total)
	})
	fmt.Printf("🔍 [DEBUG] 生成的COUNT SQL: %s\n", sql)

	if err := countQuery.Count(&total).Error; err != nil {
		fmt.Printf("❌ COUNT查询失败: %v\n", err)
		return nil, 0, err
	}
	fmt.Printf("🔍 [DEBUG] COUNT查询结果: total=%d\n", total)

	// 分页查询
	offset := (page - 1) * size
	query := r.DB.Offset(offset).Limit(size)
	if condition != "" {
		query = query.Where(condition, args...)
	}

	// 获取生成的SQL（调试用）
	findSQL := query.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Find(&entities)
	})
	fmt.Printf("🔍 [DEBUG] 生成的FIND SQL: %s\n", findSQL)

	err := query.Find(&entities).Error
	if err != nil {
		fmt.Printf("❌ FIND查询失败: %v\n", err)
	} else {
		fmt.Printf("🔍 [DEBUG] FIND查询结果: 找到%d条记录\n", len(entities))
	}

	return entities, total, err
}
func (r *BaseRepository[T]) ApplyFiltersReflect(f interface{}) *gorm.DB {
	if f == nil {
		return r.DB
	}
	db := r.DB

	quoteCol := func(col string) string {
		// 如果已经包含引号，直接返回
		if strings.ContainsAny(col, "`\"") {
			return col
		}
		parts := strings.Split(col, ".")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
			if parts[i] == "" {
				return col
			}
			// 每个标识符只允许字母数字和下划线，避免注入
			for _, ch := range parts[i] {
				if !(ch == '_' || (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')) {
					return col
				}
			}
			// Postgres 风格：用双引号包裹每个标识符
			parts[i] = `"` + parts[i] + `"`
		}
		return strings.Join(parts, ".")
	}

	var apply func(reflect.Value)
	apply = func(v reflect.Value) {
		if !v.IsValid() {
			return
		}
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return
			}
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return
		}

		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			sf := t.Field(i)
			fv := v.Field(i)

			if sf.Anonymous {
				apply(fv)
				continue
			}
			if !fv.IsValid() {
				continue
			}
			if fv.Kind() == reflect.Ptr && fv.IsNil() {
				continue
			}

			tag := sf.Tag.Get("filter")
			if tag == "" {
				if fv.Kind() == reflect.Struct {
					apply(fv)
				}
				continue
			}

			parts := strings.Split(tag, ",")
			col := strings.TrimSpace(parts[0])
			if col == "" {
				continue
			}

			// 简单白名单：只允许字母、数字、下划线和点
			valid := true
			for _, ch := range col {
				if !(ch == '.' || ch == '_' || (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')) {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			col = quoteCol(col)
			op := "eq"
			if len(parts) > 1 {
				op = strings.TrimSpace(parts[1])
			}

			var val interface{}
			if fv.Kind() == reflect.Ptr {
				val = fv.Elem().Interface()
			} else {
				val = fv.Interface()
			}
			strVal := fmt.Sprintf("%v", val)
			if strVal == "" {
				continue
			}

			switch op {
			case "eq":
				db = db.Where(col+" = ?", val)
			case "like":
				// 检查是否是 JSONB 字段（meta 字段）
				if strings.Contains(col, "meta") {
					// 对于 JSONB 字段，使用 PostgreSQL 的 JSONB 文本搜索
					db = db.Where(col+"::text LIKE ?", "%"+strVal+"%")
				} else {
					db = db.Where(col+" LIKE ?", "%"+strVal+"%")
				}
			case "jsonb_contains":
				// 可以添加更多 JSONB 特定的操作符
				db = db.Where(col+" @> ?", val)
			case "jsonb_key":
				// 检查 JSONB 是否包含某个键
				db = db.Where(col+" ? ?", val)
			default:
				db = db.Where(col+" = ?", val)
			}
		}
	}

	apply(reflect.ValueOf(f))
	return db
}
