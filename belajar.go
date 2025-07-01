package main

import (
    "database/sql"
    "log"

	"BelajarGo/handler"

    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "IT:--@tcp(127.0.0.1:3306)/dbgo")
    if err != nil {
        log.Fatal("Koneksi DB gagal:", err)
    }
    defer db.Close()

    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

	router.POST("/login", handler.Login)

	// Group route yang butuh login
	auth := router.Group("/")

    // Routing
    auth.Use(handler.AuthRequired())
    {
        auth.GET("/mahasiswa", handler.GetMahasiswa(db))
        auth.GET("/mahasiswa/:id", handler.GetMahasiswaByID(db))
        auth.POST("/mahasiswa", handler.PostMahasiswa(db))
        auth.POST("/mahasiswa/:id", handler.UpdateMahasiswa(db))
        auth.DELETE("/mahasiswa/:id", handler.DeleteMahasiswa(db))
    }

    // Dashboard
    router.GET("/dashboard", handler.ShowDashboard(db))

    router.Run(":8080")
}
