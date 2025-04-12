package config

// Mysql 表示 MySQL 数据库的配置结构体，继承了 GeneralDB 结构体的字段。
type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn 返回用于连接 MySQL 数据库的数据源名称（DSN）字符串。
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
