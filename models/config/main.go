package main

import (
  "github.com/gin-gonic/gin"
    "net/http"
    "your_project_path/config"
    "your_project_path/models" // 替换为你的项目路径
    "log"
    "gorm.io/gorm"
    "strconv"
)

func main() {
    config.Connect() // Connect to the database

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to the Sports Blog!")
    })
	 // 用户注册
	 router.POST("/register", func(c *gin.Context) {
        var newUser models.User
        if err := c.ShouldBindJSON(&newUser); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if result := config.DB.Create(&newUser); result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "User registered successfully"})
    })
	    // 用户登录
		router.POST("/login", func(c *gin.Context) {
			var loginUser models.User
			if err := c.ShouldBindJSON(&loginUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			var user models.User
			if result := config.DB.Where("username = ? AND password = ?", loginUser.Username, loginUser.Password).First(&user); result.Error != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "Logged in successfully"})
		})
	
		// 文章信息列表
		router.GET("/articles", func(c *gin.Context) {
			var articles []models.Article
			if result := config.DB.Find(&articles); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return// 文章详情
				router.GET("/articles/:id", func(c *gin.Context) {
					articleIDStr := c.Param("id")
					articleID, err := strconv.Atoi(articleIDStr)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
						return
					}
					var article models.Article
					if result := config.DB.Preload("Comments").First(&article, articleID); result.Error != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}
					c.JSON(http.StatusOK, article)
				})
		 
				// 发布文章
				router.POST("/articles", func(c *gin.Context) {
					var newArticle models.Article
					if err := c.ShouldBindJSON(&newArticle); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
						return
					}
					if result := config.DB.Create(&newArticle); result.Error != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}
					c.JSON(http.StatusOK, gin.H{"status": "Article created successfully"})
				})
		 
				// 修改文章
				router.PATCH("/articles/:id", func(c *gin.Context) {
					articleIDStr := c.Param("id")
					articleID, err := strconv.Atoi(articleIDStr)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
						return
					}
					var article models.Article
					if result := config.DB.First(&article, articleID); result.Error != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}
					if err := c.ShouldBindJSON(&article); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
						return
					}
					if result := config.DB.Save(&article).Error; result != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}
					c.JSON(http.StatusOK, gin.H{"status": "Article updated successfully"})
				})
		 
				// 删除文章
				router.DELETE("/articles/:id", func(c *gin.Context) {
					articleIDStr := c.Param("id")
					articleID, err := strconv.Atoi(articleIDStr)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
						return
					}
					if result := config.DB.Delete(&models.Article{}, articleID).Error; result != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}
					c.JSON(http.StatusOK, gin.H{"status": "Article deleted successfully"})
				})
		 
				// 搜索文章
				router.GET("/search", func(c *gin.Context) {
					query := c.Query("q")
					var articles []models.Article
					if result := config.DB.Where("title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%").Find(&articles).Error; result != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}
					c.JSON(http.StatusOK, articles)
				})
		 
				// 发布评论
				router.POST("/articles/:articleID/comments", func(c *gin.Context) {
					articleIDStr := c.Param("articleID")
					articleID, err := strconv.Atoi(articleIDStr)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
						return
					}
					var comment models.Comment
					if err := c.ShouldBindJSON(&comment); err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
						return
					}
					comment.ArticleID = uint(articleID)
					if result := config.DB.Create(&comment).Error; result != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}
					c.JSON(http.StatusOK, gin.H{"status": "Comment created successfully"})
				})
		 
				// 删除评论
				router.DELETE("/comments/:id", func(c *gin.Context) {
					commentIDStr := c.Param("id")
					commentID, err := strconv.Atoi(commentIDStr)
					if err != nil {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
						return
					}
					if result := config.DB.Delete(&models.Comment{}, commentID).Error; result != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
						return
					}
					c.JSON(http.StatusOK, gin.H{"status": "Comment deleted successfully"})
				})
		 
				router.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
			}