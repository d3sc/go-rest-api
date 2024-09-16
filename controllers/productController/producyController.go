package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/d3sc/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	// membuat variable product dengan tipe struktur product yang telah dibuat di model
	var products []models.Product

	// mencari semua product, lalu dimasukan ke dalam var struktur product
	models.DB.Find(&products)

	// mengembalikan status ok dan menampilkan product
	c.JSON(http.StatusOK, gin.H{"products": products})
}
func Show(c *gin.Context) {
	// membuat variable product dengan tipe struktur product yang telah dibuat di model
	var product models.Product

	// membuat var id yang isi nya diambil dari parameter id
	id := c.Param("id")

	// mencari data yang pertama kali ditemukan berdasarkan id
	if err := models.DB.First(&product, id).Error; err != nil {
		// kondisi error
		switch err {
		// jika data tidak ditemukan
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data tidak ditemukan",
			})
			return

		// jika server error
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	// jika tidak ada error maka tampilkan product
	c.JSON(http.StatusOK, gin.H{"product": product})

}
func Create(c *gin.Context) {
	// membuat variable product dengan tipe struktur product yang telah dibuat di model
	var product models.Product

	// mengambil body json dan dimasukan ke dalam var struktur product.
	if err := c.ShouldBindJSON(&product); err != nil {
		// dan jika error maka akan mengembalikan bad request
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messages": err.Error()})
		return
	}

	// jika tidak ada error maka create product
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}
func Update(c *gin.Context) {
	// membuat variable product dengan tipe struktur product yang telah dibuat di model
	var product models.Product

	// mengambil parameter id
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		// dan jika error maka akan mengembalikan bad request
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messages": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		// dan jika error maka akan mengembalikan bad request
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messages": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil diperbarui"})
}
func Delete(c *gin.Context) {
	// membuat variable product dengan tipe struktur product yang telah dibuat di model
	var product models.Product

	// membuat struct untuk value dari body
	var input struct {
		Id json.Number
	}

	// menangkap nilai yang ada di body yang sesuai dengan struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		// dan jika error maka akan mengembalikan bad request
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messages": err.Error()})
		return
	}

	// mengubah input Id menjadi Int64 untuk bisa menghapus data berdasarkan int yang ada di model
	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"messages": "tidak dapat menghapus product"})
		return
	}

	// jika tidak ada error maka berhasil
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil dihapus"})
}
