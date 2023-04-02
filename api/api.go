package api

import (
	_ "app/api/docs"

	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, store, logger)
	// category api
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	// brand api
	r.POST("/brand", handler.CreateBrand)
	r.GET("/brand/:id", handler.GetByIdBrand)
	r.GET("/brand", handler.GetListBrand)
	r.PUT("/brand/:id", handler.UpdateBrand)
	r.DELETE("/brand/:id", handler.DeleteBrand)

	// product api
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	// stock api  -- not ready for using
	r.POST("/stock", handler.CreateStock)
	r.GET("/stock/:id", handler.GetByIdStock)
	r.GET("/stock", handler.GetListStock)
	r.PUT("/stock/:id", handler.UpdateStock)
	r.PUT("/stock/send_product", handler.UpdateStock)
	r.DELETE("/stock/:id", handler.DeleteStock)

	// store api
	r.POST("/store", handler.CreateStore)
	r.GET("/store/:id", handler.GetByIdStore)
	r.GET("/store", handler.GetListStore)
	r.PUT("/store/:id", handler.UpdateStore)
	r.PATCH("/store/:id", handler.UpdatePatchStore)
	r.DELETE("/store/:id", handler.DeleteStore)

	// customer api
	r.POST("/customer", handler.CreateCustomer)
	r.GET("/customer/:id", handler.GetByIdCustomer)
	r.GET("/customer", handler.GetListCustomer)
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.PATCH("/customer/:id", handler.UpdatePatchCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)

	// staff api
	r.POST("/staff", handler.CreateStaff)
	r.GET("/staff/:id", handler.GetByIdStaff)
	r.GET("/staff", handler.GetListStaff)
	r.GET("/staffreport", handler.GetListReportStaff)
	r.PUT("/staff/:id", handler.UpdateStaff)
	r.PATCH("/staff/:id", handler.UpdatePatchStaff)
	r.DELETE("/staff/:id", handler.DeleteStaff)

	// order api
	r.POST("/order", handler.CreateOrder)
	r.GET("/order/:id", handler.GetByIdOrder)
	r.GET("/order", handler.GetListOrder)
	r.GET("/order/total_sum", handler.OrderTotalSum)
	r.PUT("/order/:id", handler.UpdateOrder)
	r.PATCH("/order/:id", handler.UpdatePatchOrder)
	r.DELETE("/order/:id", handler.DeleteOrder)
	r.POST("/order_item/", handler.CreateOrderItem)
	r.DELETE("/order_item/:id", handler.DeleteOrderItem)

	// code api
	r.POST("/code", handler.CreateCode)
	r.GET("/code/:id", handler.GetByIdCode)
	r.GET("/code", handler.GetListCode)
	r.PUT("/code/:id", handler.UpdateCode)
	r.DELETE("/code/:id", handler.DeleteCode)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
