package test

import (
	"BlogArticle/models"

	"fmt"
	"testing"
)

//测试函数GetArticleList() []string
func TestGetList(t *testing.T) {
	list, err := models.GetArticleList()
	fmt.Println(list)
	if err != nil {
		t.Error(`Error in GetArticleList!`)
	}
}

//测试函数GetArtiContent(id int) string
func TestGetArtiContent(t *testing.T) {
	list, err := models.GetArtiContent(6)
	fmt.Println(list)
	if err != nil {
		t.Error(`Error in GetArtiContent!`)
	}
}

//测试函数UpdateArticle(article Article) (ra int64, err error)
func TestUpdateArticle(t *testing.T) {
	var article models.Article
	article.Id = 1
	article.Title = "testTitle"
	article.Author = "testAuthor"
	article.Content = "testContent"
	article.LastTime = "testTime"

	num, err := models.UpdateArticle(article)
	fmt.Println(num)
	if err != nil {
		t.Error(`Error in UpdateArticle!`)
	}
}

//测试函数DeleteArticle(id int) (ra int64, err error)
func TestDeleteArticle(t *testing.T) {
	testId := 1
	num, err := models.DeleteArticle(testId)
	fmt.Println(num)
	if err != nil {
		t.Error(`Error`)
	}
}
