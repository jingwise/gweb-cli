// Package gweb_cli
// @Author: itcyy@HuaWei
// @File: /main.go
// @Time: 2024/1/3 15:41
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "version" {
			fmt.Println("Windows v0.0.3")
			return
		}
		if arg == "help" {
			fmt.Println("version View version\n ")
			return
		}
	}
	projectName := getProjectName()
	directoryPath := getDirectoryPath(projectName)
	createDirectory(directoryPath)
	createMainDirectory(directoryPath)
	createGoFile(directoryPath, "main.go")
}

func getProjectName() string {
	fmt.Print("project name > ")
	var projectName string
	fmt.Scanln(&projectName)
	return strings.TrimSpace(projectName)
}

func getPackageName(projectName string) string {
	return strings.ToLower(projectName[:1]) + projectName[1:] + "package"
}

func getDirectoryPath(projectName string) string {
	homeDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	projectDir := filepath.Join(homeDir, projectName)
	return projectDir
}

func createDirectory(directoryPath string) {
	err := os.MkdirAll(directoryPath, 0755)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Project directory created: %s\n", directoryPath)
}

func createGoFile(directoryPath, packageName string) {
	goFilePath := filepath.Join(directoryPath, packageName)
	goCode := `package main  
		import "fmt"  
		func main() {  
		 	fmt.Println("Hello, %s!")  
		}`

	err := os.WriteFile(goFilePath, []byte(goCode), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Go file created: %s\n", goFilePath)
}
func createYamlFile(directoryPath string) {
	goFilePath := filepath.Join(directoryPath, "config.yaml")
	goCode := `#MySql配置
mysql:
  username: root
  password: 123456
  host: 127.0.0.1
  port: 3306
  database: database

#MongoDB配置
mongod:
  username: root
  password: 123456
  host: 127.0.0.1
  port: 27017

# 缓存配置 Redis配置
cache:
  driver: memory
  redis:
    addr: 127.0.0.1:6379
    username: default
    password: 123456
    db: 0

# 邮件配置
mail:
  host: smtp.qq.com
  port: 465
  username: <EMAIL>
  password: 123456
  sender: <EMAIL>
  subject: 邮件主题
  body: 邮件内容

# 微信小程序配置
wechat:
  appid: wxa6***
  secret: d418f3****
#开发环境  dev: 开发环境   localhost:  本地开发环境  main:  正式环境
main:
  environment: dev

#腾讯云配置 公司查询接口
tencentcould:
  secretid: AKIDom6vWakK***
  secretkey: 21p17Ta8R895***

service:
  port: 9090
  crt: ./www/test.crt
  key: ./www/test.key
`

	err := os.WriteFile(goFilePath, []byte(goCode), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Yaml file created: %s\n", goFilePath)
}
func createMainDirectory(directoryPath string) {
	createDirectory(directoryPath + "/application/config")
	createYamlFile(directoryPath)
	createYamlGoFile(directoryPath + "/application/config")
	createDirectory(directoryPath + "/interfaces/service")
	createDirectory(directoryPath + "/interfaces/handler")
	createDirectory(directoryPath + "/domain/mapper")
	createDirectory(directoryPath + "/domain/cache")
	createDirectory(directoryPath + "/infrastructure/pkg/logger/log")
	createLoggerGoFile(directoryPath + "/infrastructure/pkg/logger/log")
	createDirectory(directoryPath + "/www")
	createDirectory(directoryPath + "/application/response")
	createResponseCodeGoFile(directoryPath + "/application/response")
	createResponseModelGoFile(directoryPath + "/application/response")
	createResponseReturnGoFile(directoryPath + "/application/response")
	createResponseTypeGoFile(directoryPath + "/application/response")

}
func createYamlGoFile(directoryPath string) {
	goFilePath := filepath.Join(directoryPath, "config.go")
	goCode := "type Config struct {\n\tMySQL struct {\n\t\tUsername string `yaml:\"username\"`\n\t\tPassword string `yaml:\"password\"`\n\t\tHost     string `yaml:\"host\"`\n\t\tPort     int    `yaml:\"port\"`\n\t\tDatabase string `yaml:\"database\"`\n\t} `yaml:\"mysql\"`\n\n\tMongod struct {\n\t\tUsername string `yaml:\"username\"`\n\t\tPassword string `yaml:\"password\"`\n\t\tHost     string `yaml:\"host\"`\n\t\tPort     int    `yaml:\"port\"`\n\t} `yaml:\"mongod\"`\n\n\tCache struct {\n\t\tDriver string `yaml:\"driver\"`\n\t\tRedis  struct {\n\t\t\tAddr     string `yaml:\"addr\"`\n\t\t\tUsername string `yaml:\"username\"`\n\t\t\tPassword string `yaml:\"password\"`\n\t\t\tDB       int    `yaml:\"db\"`\n\t\t} `yaml:\"redis\"`\n\t} `yaml:\"cache\"`\n\n\tMail struct {\n\t\tHost     string `yaml:\"host\"`\n\t\tPort     int    `yaml:\"port\"`\n\t\tUsername string `yaml:\"username\"`\n\t\tPassword string `yaml:\"password\"`\n\t\tSender   string `yaml:\"sender\"`\n\t\tSubject  string `yaml:\"subject\"`\n\t\tBody     string `yaml:\"body\"`\n\t} `yaml:\"mail\"`\n\n\tWechat struct {\n\t\tAppID  string `yaml:\"appid\"`\n\t\tSecret string `yaml:\"secret\"`\n\t} `yaml:\"wechat\"`\n\n\tMain struct {\n\t\tEnvironment string `yaml:\"environment\"`\n\t} `yaml:\"main\"`\n\n\tTencentCould struct {\n\t\tSecretId  string `yaml:\"secretid\"`\n\t\tSecretKey string `yaml:\"secretkey\"`\n\t} `yaml:\"tencentcould\"`\n\n\tService struct {\n\t\tPort string `yaml:\"port\"`\n\t\tCrt  string `yaml:\"crt\"`\n\t\tKey  string `yaml:\"key\"`\n\t} `yaml:\"service\"`\n}\n\nfunc YamlConfig() (*Config, error) {\n\tvar config *Config\n\tdata, err := os.ReadFile(\"config.yaml\")\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\terr = yaml.Unmarshal(data, &config)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn config, nil\n}"
	err := os.WriteFile(goFilePath, []byte(goCode), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Go file created: %s\n", goFilePath)
}
func createResponseCodeGoFile(directoryPath string) {
	goFilePath := filepath.Join(directoryPath, "code.go")
	f, err := os.Open("pkg/response/code.go")
	if err != nil {
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	err = os.WriteFile(goFilePath, data, 0644)
	fmt.Printf("Go file created: %s\n", goFilePath)

}
func createResponseModelGoFile(directoryPath string) {
	goFilePath := filepath.Join(directoryPath, "model.go")
	f, err := os.Open("pkg/response/model.go")
	if err != nil {
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	err = os.WriteFile(goFilePath, data, 0644)
	fmt.Printf("Go file created: %s\n", goFilePath)

}
func createResponseReturnGoFile(directoryPath string) {
	goFilePath := filepath.Join(directoryPath, "return.go")
	f, err := os.Open("pkg/response/return.go")
	if err != nil {
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	err = os.WriteFile(goFilePath, data, 0644)
	fmt.Printf("Go file created: %s\n", goFilePath)

}
func createResponseTypeGoFile(directoryPath string) {
	goFilePath := filepath.Join(directoryPath, "type.go")
	f, err := os.Open("pkg/response/type.go")
	if err != nil {
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	err = os.WriteFile(goFilePath, data, 0644)
	fmt.Printf("Go file created: %s\n", goFilePath)

}
func createLoggerGoFile(directoryPath string) {
	goFilePath := filepath.Join(directoryPath, "logger.go")
	f, err := os.Open("pkg/logger/log.go")
	if err != nil {
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	err = os.WriteFile(goFilePath, data, 0644)
	fmt.Printf("Go file created: %s\n", goFilePath)

}
