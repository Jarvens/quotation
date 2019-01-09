# 行情服务器

- codec 编解码
- common 公共
- config 配置资源
- domain 实体对象
- errors 异常
- main 启动
- server socket & ws 服务
- utils 工具


### 命名规范
- 接口只有一个方法时，需要以er结尾 例如 
```go
type Handler interface{
	Searcher()(message string)
}
```
- 接口中存在多个方法需要明确方法动作
- 启动类一般命名为command
- go 文件使用小写+下划线方式组合命名