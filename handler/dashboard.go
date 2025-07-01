package handler

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
)

// Fungsi handler untuk GET dashboard
func ShowDashboard(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, nama, jurusan, angkatan, ipk FROM mahasiswa")
		if err != nil {
			c.String(http.StatusInternalServerError, "Query error: %s", err.Error())
			return
		}
		defer rows.Close()

		var list []Mahasiswa
		for rows.Next() {
			var mhs Mahasiswa
			if err := rows.Scan(&mhs.ID, &mhs.Nama, &mhs.Jurusan, &mhs.Angkatan, &mhs.IPK); err != nil {
				c.String(http.StatusInternalServerError, "Scan error: %s", err.Error())
				return
			}
			list = append(list, mhs)
		}

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"title":     "Dashboard Mahasiswa",
			"mahasiswa": list,
		})
	}
}
