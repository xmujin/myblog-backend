package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xmujin/myblog-backend/internal/model"
	"github.com/xmujin/myblog-backend/internal/service"
)

type PostController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (p *PostController) GetPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := p.postService.GetPosts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"posts": posts})
	}
}

func (p *PostController) GetPostById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"非法请求参数：": err.Error()})
			return
		}
		post, err := p.postService.GetPost(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"获取文章失败：": err.Error()})
			return
		}
		c.JSON(http.StatusOK, post)
	}
}

func (p *PostController) DeletePostById() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"非法请求参数：": err.Error()})
			return
		}
		err = p.postService.DeletePostById(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"删除文章失败": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "删除文章成功",
		})

	}
}

func (p *PostController) CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postRequest model.PostRequest
		if err := c.BindJSON(&postRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err := p.postService.CreatePost(&postRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "文章创建成功",
		})
	}
}

func (p *PostController) UpdatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"非法请求参数：": err.Error()})
			return
		}
		var postRequest model.PostRequest
		if err := c.BindJSON(&postRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"错误的请求": err.Error()})
			return
		}
		err = p.postService.UpdatePostById(uint(id), &postRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "更新成功",
		})
	}
}

func (p *PostController) DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
