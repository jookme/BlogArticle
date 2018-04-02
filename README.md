# BlogArticle

## 简易博客系统文章部分接口

### 目录
* [数据库](#数据库)
* [功能实现](#功能)
  * [查询文章列表](#查询文章列表)
  * [查询文章内容](#查询文章内容)
  * [文章记录更新](#文章记录更新)
  * [文章记录删除](#文章记录删除)
* [单元测试](#单元测试)

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
/article<br>
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

* model函数实现  （源码地址：models/article.go:GetArticleList）<br>
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
* api接口实现  （源码地址：apis/handler.go:GetListAPI）<br>
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
* HTTP方法<br>
get<br>
* 请求url地址<br>
/article/:id<br>
* 输入参数<br>
前端页面传来的文章id<br>
* 输出参数<br>
文章内容 content 或 错误信息 err<br>
* 参数说明<br>

|参数名|说明|类型| 
|:--------------:|:---------------:|:-----------:|
|id| 文章id |  int  |  
|content| 文章内容 |  string  |  
|err|错误信息|   error    |  

* 流程图<br>
![404 找不到！](https://github.com/jookme/BlogArticle/blob/master/img/flowchart/%E6%9F%A5%E8%AF%A2%E6%96%87%E7%AB%A0%E5%86%85%E5%AE%B9.png "")<br>

* model函数实现  （源码地址：models/article.go:GetArtiContent）<br>
```
//用content存放获取的文章内容
	var content string
	//从表Article中查询对应标题title的文章内容
	err := db.SqlDb.QueryRow("SELECT Content FROM Article where Id = ?", id).Scan(&content)
	if err != nil {
		log.Fatal(err)
	}

```
* api接口实现  （源码地址：apis/handler.go:GetContentApi）<br>
```
aid := c.Param("id")
	id, err := strconv.Atoi(aid)
	if err != nil {
		log.Fatalln(err)
	}

	var content string
	content, err = models.GetArtiContent(id)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
```
#### 文章记录更新
* HTTP方法<br>
put<br>
* 请求url地址<br>
/article/:id<br>
* 输入参数<br>
文章id<br>
* 输出参数<br>
更新成功提示msg 或 错误信息 err<br>
* 参数说明<br>

|参数名|说明|类型| 
|:--------------:|:---------------:|:-----------:|
|msg|更新成功提示|string|  
|err|错误信息|error|  

* 流程图<br>
![404 找不到！](https://github.com/jookme/BlogArticle/blob/master/img/flowchart/%E6%96%87%E7%AB%A0%E8%AE%B0%E5%BD%95%E6%9B%B4%E6%96%B0.png "")<br>

* model函数实现  （源码地址：models/article.go:UpdateArticle）<br>
```
stmt, err := db.SqlDb.Prepare("UPDATE Article SET Title=?, Author=? , Content=? , LastTime=? WHERE Id=?")
	defer stmt.Close()
	if err != nil {
		log.Fatalln(err)
	}
	//执行更新
	res, err := stmt.Exec(article.Title, article.Author, article.Content, article.LastTime, article.Id)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err = res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
```
* api接口实现  （源码地址：apis/handler.go:UpdateArticApi）<br>
```
aid := c.Param("id")
	id, err := strconv.Atoi(aid)
	if err != nil {
		log.Fatalln(err)
	}

	//通过bind类型绑定的方式来获取更新的内容
	article := models.Article{Id: id}
	err = c.Bind(&article)
	if err != nil {
		log.Fatalln(err)
	}

	//更新
	_, err = models.UpdateArticle(article)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Update Article Success!",
	})
```
#### 文章记录删除
* HTTP方法<br>
delete<br>
* 请求url地址<br>
/article/:id<br>
* 输入参数<br>
文章id<br>
* 输出参数<br>
删除成功提示msg 或 错误信息 err<br>
* 参数说明<br>

|参数名|说明|类型| 
|:--------------:|:---------------:|:-----------:|
|msg|删除成功提示|string|  
|err|错误信息|error|  

* 流程图<br>
![404 找不到！](https://github.com/jookme/BlogArticle/blob/master/img/flowchart/%E6%96%87%E7%AB%A0%E8%AE%B0%E5%BD%95%E5%88%A0%E9%99%A4.png "")<br>

* model函数实现  （源码地址：models/article.go:DeleteArticle）<br>
```
stmt, err := db.SqlDb.Prepare("DELETE FROM Article WHERE Id=?")
	if err != nil {
		log.Fatalln(err)
	}
	//执行删除
	res, err := stmt.Exec(1)
	if err != nil {
		log.Fatalln(err)
	}

	ra, err = res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
```
* api接口实现  （源码地址：apis/handler.go:DelArticApi）<br>
```
aid := c.Param("id")
	id, err := strconv.Atoi(aid)
	if err != nil {
		log.Fatalln(err)
	}
	//删除
	_, err = models.DeleteArticle(id)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "Delete Article Success!",
	})
```
### 单元测试
#### 对TestGetList进行测试
* 测试前Article表<br>
![404 找不到！](https://github.com/jookme/BlogArticle/blob/master/img/database/article%E8%A1%A8.png "artic表")<br>
* 测试代码<br>
```
list, err := models.GetArticleList()
	fmt.Printf("The list is :%s", list)
	fmt.Println("")
	if err != nil {
		t.Error(`Error in GetArticleList!`)
	}
```
* 在vscode终端输入测试命令<br>
![404 找不到！](https://github.com/jookme/BlogArticle/blob/master/img/test/TestGetlist1.png "")<br>
* 测试结果输出<br>
![404 找不到！](https://github.com/jookme/BlogArticle/blob/master/img/test/TestGetlist2.png "")<br>

