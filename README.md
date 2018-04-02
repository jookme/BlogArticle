# BlogArticle

## 简易博客系统文章部分接口

### 目录
* [数据库](#数据库)
* [功能实现](#功能)
  * [查询文章列表](#查询文章列表)
  * [查询文章内容](#查询文章内容)
  * [文章记录更新](#文章记录更新)
  * [文章记录删除](#文章记录删除)
* [封装](#封装)
* [入口](#入口)
* [文件目录说明](#文件目录)

### 数据库
使用的数据库：mysql<br>
数据库建表<br>
![404 找不到！](https://github.com/jookme/BlogArticle/blob/master/img/database/article%E8%A1%A8.png "artic表")<br>
初始化数据库连接池<br>
```
//初始化一个sql.DB对象,接口为mysql中的blog数据库，账号为root，密码为1996
	SqlDb, err := sql.Open("mysql", "root:1996@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer SqlDb.Close()

	//测试与数据库的连接是否可用
	err = SqlDb.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
```    
<br>
### 功能
#### 查询文章列表
#### 查询文章内容
#### 文章记录更新
#### 文章记录删除



