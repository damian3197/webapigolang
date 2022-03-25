package controller

import (
	"net/http"
	"tugaswebapi/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type MahasiswaInput struct {
	ID       int    `json:"id" binding: "required,number"`
	Nama     string `json:"nama" binding: "required,gt=5"`
	Prodi    string `json:"prodi" binding: "required"`
	Fakultas string `json:"fakultas" binding: "required"`
	NIM      int    `json:"nim" binding: "required,number,gt=5"`
	Angkatan int    `json:"angkatan" binding: "required,number"`
}

// GET DATA /  READ DATA
func GetData(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{
		"data" : mhs,
	})

}

// POST DATA  CREATE DATA
func CreateData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validasi Inputan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput) 
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	// proses input data
	mhs := models.Mahasiswa{
		ID: dataInput.ID,
		Nama: dataInput.Nama,
		Prodi: dataInput.Prodi,
		Fakultas: dataInput.Fakultas,
		NIM: dataInput.NIM,
		Angkatan: dataInput.Angkatan,
	}
	db.Create(&mhs)
	// Menampilkan Hasil
	c.JSON(http.StatusOK, gin.H{
		"Message" : "Berhasil Input Data",
		"Data" : mhs,
	})
}

// UPDATE DATA
func UpdateData(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)

	// cek data
	var mhs models.Mahasiswa
	if err := db.Where("ID = ?", c.Param("ID")).First(&mhs).Error;
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Data Mahasiswa Tidak Ditemukan",
		})
		return
	}
}
