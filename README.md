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
#### 初始化数据库连接池    （源码地址：database/mysql.go）<br>
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

* 流程图<br>
![404 找不到！](https://github.com/jookme/BlogArticle/blob/master/img/flowchart/%E6%9F%A5%E8%AF%A2%E6%96%87%E7%AB%A0%E5%88%97%E8%A1%A8.png "")<br>

* model函数实现  （源码地址：models/article.go）<br>
```
rows, err := db.SqlDb.Query("SELECT Title FROM Article")

	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	//ArticleList用于存放文章标题列表信息
	ArticleList := make([]string, 0)

	//遍历获取到的数据
	for rows.Next() {
		var article Article
		rows.Scan(&article.Title)
		//标题依次添加到ArticleList中
		ArticleList = append(ArticleList, article.Title)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
```
* api接口实现  （源码地址：apis/handler.go）<br>
```
list, err := models.GetArticleList()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
```

#### 查询文章内容
#### 文章记录更新
#### 文章记录删除



