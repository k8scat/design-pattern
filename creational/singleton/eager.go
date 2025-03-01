package singleton

// 示例模式：单例模式 -> 饿汉方法.

// 初始化一个全局的单例变量.
var eager *Eager

// 定义单例模式类型
type Eager struct {
	count int
}

// Eager 类型 Inc 方法.
func (e *Eager) Inc() {
	e.count++
}

// 初始化单例实例，这里不建议用 init() 函数，因为使用 init 函数初始化时，
// 开发者无感知，不利于维护，容易出篓子.
func InitEager(count int) {
	eager = &Eager{count: count}
}

// 获取全局的单例实例，这里只读，是并发安全的.
func GetEager() *Eager {
	return eager
}
