package conf

import (
	"mywebmall/dao"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	// 在cache包里面进行读取设置
	// RedisDb     string
	// RedisAddr   string
	// RedisPw     string
	// RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	// 配置载入
	LoadServer(file)
	LoadMySql(file)
	// LoadRedis(file)
	LoadEmail(file)
	LoadPath(file)
	// mysql读路径
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	// mysql写路径
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	dao.Database(pathRead, pathWrite)
}

func LoadServer(f *ini.File) {
	AppMode = f.Section("service").Key("AppMode").String()
	HttpPort = f.Section("service").Key("HttpPort").String()
}

func LoadEmail(f *ini.File) {
	ValidEmail = f.Section("email").Key("ValidEmail").String()
	SmtpHost = f.Section("email").Key("SmtpHost").String()
	SmtpEmail = f.Section("email").Key("SmtpEmail").String()
	SmtpPass = f.Section("email").Key("SmtpPass").String()
}

func LoadMySql(f *ini.File) {
	Db = f.Section("mysql").Key("Db").String()
	DbHost = f.Section("mysql").Key("DbHost").String()
	DbPort = f.Section("mysql").Key("DbPort").String()
	DbUser = f.Section("mysql").Key("DbUser").String()
	DbPassword = f.Section("mysql").Key("DbPassword").String()
	DbName = f.Section("mysql").Key("DbName").String()
}

// func LoadRedis(f *ini.File) {
// 	RedisDb = f.Section("redis").Key("RedisDb").String()
// 	RedisAddr = f.Section("redis").Key("RedisAddr").String()
// 	RedisPw = f.Section("redis").Key("RedisPw").String()
// 	RedisDbName = f.Section("redis").Key("RedisDbName").String()
// }

func LoadPath(f *ini.File) {
	Host = f.Section("path").Key("Host").String()
	ProductPath = f.Section("path").Key("ProductPath").String()
	AvatarPath = f.Section("path").Key("AvatarPath").String()
}
