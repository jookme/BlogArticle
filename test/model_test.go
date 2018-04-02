package testing

import (
	"BlogArticle/models"
	"fmt"
	"testing"
)

//测试函数GetArticleList() []string
func TestGetList(t *testing.T) {
	list, err := models.GetArticleList()
	fmt.Printf("The list is :%s", list)
	fmt.Println("")
	if err != nil {
		t.Error(`Error in GetArticleList!`)
	}
}

//测试函数GetArtiContent(id int) string
func TestGetArtiContent(t *testing.T) {
	content, err := models.GetArtiContent(2)
	fmt.Printf("The content is :%s", content)
	fmt.Println("")
	if err != nil {
		t.Error(`Error in GetArtiContent!`)
	}
}

//测试函数UpdateArticle(article Article) (ra int64, err error)
func TestUpdateArticle(t *testing.T) {
	var article models.Article
	article.Id = 3
	article.Title = "testTitle"
	article.Author = "testAuthor"
	article.Content = "testContent"
	article.LastTime = "2019"

	num, err := models.UpdateArticle(article)
	fmt.Println(num)
	if err != nil {
		t.Error(`Error in UpdateArticle!`)
	} else {
		fmt.Println("Update Success!")
	}
}

//测试函数DeleteArticle(id int) (ra int64, err error)
func TestDeleteArticle(t *testing.T) {
	testId := 2
	num, err := models.DeleteArticle(testId)
	fmt.Println(num)
	if err != nil {
		t.Error(`Error`)
	} else {
		fmt.Println("Delete Success!")
	}
}
