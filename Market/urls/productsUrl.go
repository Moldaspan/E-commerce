package urls

import "github.com/gin-gonic/gin"

func SetupRouter(pc *ProductController, cc *CategoryController) *gin.Engine {
	r := gin.Default()

	// Products endpoints
	r.POST("/products", pc.CreateProduct)
	r.GET("/products/:id", pc.GetProductByID)
	r.PUT("/products", pc.UpdateProduct)
	r.DELETE("/products/:id", pc.DeleteProduct)

	// Categories endpoints
	r.POST("/categories", cc.CreateCategory)
	r.GET("/categories/:id", cc.GetCategoryByID)
	r.PUT("/categories", cc.UpdateCategory)
	r.DELETE("/categories/:id", cc.DeleteCategory)

	return r
}
