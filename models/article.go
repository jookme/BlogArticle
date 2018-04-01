package models

import (
	db "BlogArticle/database"
	"log"
)

//定义一个文章结构体类型，成员数据包括：Id、标题、作者、内容和最后一次修改时间
type Article struct {
	Id       int    `json:"id" form:"id"`
	Title    string `json:"Title" form:"Title"`
	Author   string `json:"Author" form:"Author"`
	Content  string `json:"Content" form:"Content"`
	LastTime string `json:"LastTime" form:"LastTime"`
}

//GetArticleList 函数用于获取文章列表*****
//传入参数：无
//传出参数：文章列表信息，错误信息
func GetArticleList() ([]string, error) {
	//从表Article中查询文章列表
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

	return ArticleList, nil
}

//GetArtiContent 函数用于获取文章内容*****
//传入参数：文章id
//传出参数：对应文章内容，错误信息
func GetArtiContent(id int) (string, error) {
	//用content存放获取的文章内容
	var content string
	//从表Article中查询对应标题title的文章内容
	err := db.SqlDb.QueryRow("SELECT Content FROM Article where Id = ?", id).Scan(&content)
	if err != nil {
		log.Fatal(err)
	}

	return content, nil
}

//UpdateArticle 函数用于更新文章*****
//传入参数：Article 结构体类型变量
//传出参数：Article表中影响的数据行数，错误信息
func UpdateArticle(article Article) (ra int64, err error) {
	// PreparedStatement 可以防止SQL注入攻击
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
	return
}

//DeleteArticle 函数用于删除文章*****
//传入参数：一个文章id
//传出参数：Article表中影响的数据行数，错误信息
func DeleteArticle(id int) (ra int64, err error) {
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
	return
}
