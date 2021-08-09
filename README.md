# heavenutil

## service
### mysql
该包主要集合封装常用MySQL操作
## logran
以logrus为基础，提供仿logback的日志配置功能  
配置文件格式：[logran.xml](logran/testdata/logran.xml)  
配置类：[Configuration](logran/structs.go)  
读取配置文件方式：

```
func LoadConfiguration(path string) (*Configuration, error)
```
自定义[Hook](logran/hook.go)及[简单示例](logran/hook_test.go)
## comutil
该包提供工具类操作
加载YAML配置
```
func LoadProperties(path string, out interface{}) error
```
驼峰转下划线
```
func HumpToUnderline(title string) string
```
转float64
```
func ValueToFloat64(value interface{}) (float64, error)
```
转int64
```
func ValueToInt64(value interface{}) (int64, error)
```
## errutil
简单的对error的处理  
