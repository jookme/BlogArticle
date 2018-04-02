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
#### 使用的数据库<br>
mysql<br>
#### 数据库建表 (主键为 Id）<br>
![404 找不到！](https://github.com/jookme/BlogArticle/blob/master/img/database/article%E8%A1%A8.png "artic表")<br>
#### 初始化数据库连接池<br>
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

### 功能
#### 查询文章列表
* HTTP方法<br>
get<br>
* 请求url地址<br>
/Article<br>
* 输入参数<br>
无<br>
* 输出参数<br>
文章列表信息 list 或 错误信息 err<br>
* 参数说明<br>

|参数名|说明|类型| 
|:--------------:|:---------------:|:-----------:|
|list| 所有文章标题列表 |  []string  |  
|err|错误信息|   error    |  

#### 查询文章内容
#### 文章记录更新
#### 文章记录删除



