package apis

import (
	"BlogArticle/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

//GetListApi 获取文章列表接口*****
func GetListAPI(c *gin.Context) {
	list, err := models.GetArticleList()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}

//GetContentApi 获取文章内容接口*****
func GetContentApi(c *gin.Context) {
	aid := c.Param("id")
	id, err := strconv.Atoi(aid)
	if err != nil {
		log.Fatalln(err)
	}
	content := models.GetArtiContent(id)

	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}

//UpdateArticApi 更新文章信息接口*****
func UpdateArticApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
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
}

//DelArticApi 删除文章接口*****
func DelArticApi(c *gin.Context) {
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
}
