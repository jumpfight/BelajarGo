package handler

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Mahasiswa struct {
    ID       int     `json:"id"`
    Nama     string  `json:"nama"`
    Jurusan  string  `json:"jurusan"`
    Angkatan int     `json:"angkatan"`
    IPK      float64 `json:"ipk"`
}

// Fungsi handler untuk POST mahasiswa
func PostMahasiswa(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var mhs Mahasiswa
		if err := c.ShouldBindJSON(&mhs); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := "INSERT INTO mahasiswa (nama, jurusan, angkatan, ipk) VALUES (?, ?, ?, ?)"
		result, err := db.Exec(query, mhs.Nama, mhs.Jurusan, mhs.Angkatan, mhs.IPK)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		id, _ := result.LastInsertId()
		mhs.ID = int(id)
		c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa ditambahkan", "data": mhs})
	}
}

// Fungsi handler untuk GET mahasiswa
func GetMahasiswa(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        rows, err := db.Query("SELECT id, nama, jurusan, angkatan, ipk FROM mahasiswa")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()

        var list []Mahasiswa
        for rows.Next() {
            var mhs Mahasiswa
            if err := rows.Scan(&mhs.ID, &mhs.Nama, &mhs.Jurusan, &mhs.Angkatan, &mhs.IPK); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            list = append(list, mhs)
        }

        c.JSON(http.StatusOK, gin.H{"data": list})
    }
}

// Fungsi untuk GET mahasiswa by ID
func GetMahasiswaByID(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id") // Ambil parameter ID dari URL

        var mhs Mahasiswa
        err := db.QueryRow("SELECT id, nama, jurusan, angkatan, ipk FROM mahasiswa WHERE id = ?", id).
            Scan(&mhs.ID, &mhs.Nama, &mhs.Jurusan, &mhs.Angkatan, &mhs.IPK)

        if err != nil {
            if err == sql.ErrNoRows {
                c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            }
            return
        }

        c.JSON(http.StatusOK, gin.H{"data": mhs})
    }
}

// Fungsi handler untuk UPDATE mahasiswa
func UpdateMahasiswa(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var mhs Mahasiswa
        if err := c.ShouldBindJSON(&mhs); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        stmt, err := db.Prepare("UPDATE mahasiswa SET nama=?, jurusan=?, angkatan=?, ipk=? WHERE id=?")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer stmt.Close()

        _, err = stmt.Exec(mhs.Nama, mhs.Jurusan, mhs.Angkatan, mhs.IPK, id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa berhasil diupdate"})
    }
}

// FUngsi handler untuk DELETE mahasiswa
func DeleteMahasiswa(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")

        stmt, err := db.Prepare("DELETE FROM mahasiswa WHERE id=?")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer stmt.Close()

        _, err = stmt.Exec(id)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Mahasiswa berhasil dihapus"})
    }
}

