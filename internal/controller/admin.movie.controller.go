package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/response"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminMovieController struct {
	movieService *service.AdminMovieService
}

func AdminNewMovieController(movieService *service.AdminMovieService) *AdminMovieController {
	return &AdminMovieController{
		movieService: movieService,
	}
}

// @Summary      Get Admin List Movies
// @Description  Mengambil daftar semua film dengan support searching, filtering, dan pagination
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        page     query    int    false  "Halaman aktif (default: 1)"
// @Param        limit    query    int    false  "Jumlah data per halaman (default: 10)"
// @Param        month    query    int    false  "Filter Bulan"
// @Param        year     query    int    false  "Filter Tahun"
// @Success      200 {object} dto.AdminResponseSuccess{data=dto.AdminMovieListResponse} "Sukses mengambil list film"
// @Failure      400 {object} dto.AdminResponseError "Query parameter tidak valid"
// @Failure      500 {object} dto.AdminResponseError "Internal server error"
// @Router       /admin/movies [get]
func (c *AdminMovieController) AdminGetMovies(ctx *gin.Context) {
	var params dto.AdminMovieQueryParams

	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid query parameters")
		return
	}

	res, err := c.movieService.AdminGetMovieList(ctx.Request.Context(), params)
	if err != nil {
		fmt.Println("ERROR DATABASE:", err)
		response.Error(ctx, http.StatusInternalServerError, "Failed to fetch movie list")
		return
	}

	response.Success(ctx, http.StatusOK, "Get Movie List Success", res)
}

// @Summary      Add New Movie
// @Description  Menambahkan data film baru beserta file poster dan otomatisasi jadwal tayang
// @Tags         Admin
// @Accept       multipart/form-data
// @Produce      json
// @Param        title           formData string true  "Judul Film"
// @Param        poster          formData file   true  "File Gambar Poster Film"
// @Param        release_date    formData string true  "Tanggal Rilis (Format: YYYY-MM-DD)"
// @Param        duration_hour   formData int    true  "Durasi Jam"
// @Param        duration_minute formData int    true  "Durasi Menit"
// @Param        synopsis        formData string true  "Sinopsis"
// @Param        genre_ids       formData []int  true  "Array ID Genre" collectionFormat(multi)
// @Param        cast_ids        formData []int  true  "Array ID Cast" collectionFormat(multi)
// @Param        director_ids    formData []int  true  "Array ID Sutradara" collectionFormat(multi)
// @Param        location_ids    formData []int  true  "Array ID Kota/Lokasi Bioskop" collectionFormat(multi)
// @Param        dates           formData []string true "Array Tanggal (YYYY-MM-DD)" collectionFormat(multi)
// @Param        times           formData []string true "Array Jam Tayang (HH:MM)" collectionFormat(multi)
// @Success      201 {object}     dto.AdminResponseSuccess "Sukses menambahkan film baru"
// @Failure      400 {object}     dto.AdminResponseError "Payload input atau file tidak valid"
// @Failure      500 {object}     dto.AdminResponseError "Internal server error"
// @Router       /admin/movies [post]
func (c *AdminMovieController) AdminCreateMovie(ctx *gin.Context) {
	var req dto.AdminAddMovieRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if req.Poster == nil {
		response.Error(ctx, http.StatusBadRequest, "Poster file is required")
		return
	}

	const maxFileSize = 2 * 1024 * 1024
	if req.Poster.Size > maxFileSize {
		response.Error(ctx, http.StatusBadRequest, "Image size too large. Maximum size allowed is 2MB")
		return
	}

	extension := strings.ToLower(filepath.Ext(req.Poster.Filename))

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	if !allowedExtensions[extension] {
		response.Error(ctx, http.StatusBadRequest, "Invalid image format. Only .jpg, .jpeg, and .png are allowed")
		return
	}

	filename := fmt.Sprintf("%d_movie%s", time.Now().Unix(), extension)
	dst := filepath.Join("public/img", filename)

	if err := ctx.SaveUploadedFile(req.Poster, dst); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to save image: "+err.Error())
		return
	}

	movieID, err := c.movieService.AdminCreateMovie(ctx.Request.Context(), req, filename)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, http.StatusCreated, "Create New Movie Success", gin.H{
		"movie_id":   movieID,
		"poster_url": fmt.Sprintf("/img/%s", filename),
	})
}

// @Summary      Edit Movie
// @Description  Mengubah data film. Biarkan input kosong jika tidak ingin mengubah data spesifik.
// @Tags         Admin
// @Accept       multipart/form-data
// @Produce      json
// @Param        id              path     int    true  "ID Film yang ingin diedit"
// @Param        title           formData string false "Judul Film Baru"
// @Param        poster          formData file   false "File Gambar Poster Baru"
// @Param        release_date    formData string false "Tanggal Rilis Baru (YYYY-MM-DD)"
// @Param        duration_hour   formData int    false "Durasi Jam Baru"
// @Param        duration_minute formData int    false "Durasi Menit Baru"
// @Param        synopsis        formData string false "Sinopsis Baru"
// @Param        genre_ids       formData []int  false "Array ID Genre Baru" collectionFormat(multi)
// @Param        cast_ids        formData []int  false "Array ID Cast Baru" collectionFormat(multi)
// @Param        director_ids    formData []int  false "Array ID Sutradara Baru" collectionFormat(multi)
// @Param        location_ids    formData []int  false "Array ID Lokasi Baru" collectionFormat(multi)
// @Param        dates           formData []string false "Array Tanggal Baru (YYYY-MM-DD)" collectionFormat(multi)
// @Param        times           formData []string false "Array Waktu Baru (HH:MM)" collectionFormat(multi)
// @Success      200 {object}    dto.AdminResponseSuccess "Sukses mengubah data film"
// @Failure      400 {object}    dto.AdminResponseError "Payload input tidak valid"
// @Failure      404 {object}    dto.AdminResponseError "Film tidak ditemukan"
// @Failure      500 {object}    dto.AdminResponseError "Internal server error"
// @Router       /admin/movies/{id} [put]
func (c *AdminMovieController) AdminUpdateMovie(ctx *gin.Context) {
	idStr := ctx.Param("id")
	movieID, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid movie ID")
		return
	}

	var req dto.AdminEditMovieRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var filename string
	if req.Poster != nil {
		const maxFileSize = 2 * 1024 * 1024
		if req.Poster.Size > maxFileSize {
			response.Error(ctx, http.StatusBadRequest, "Image size too large. Max 2MB")
			return
		}

		extension := strings.ToLower(filepath.Ext(req.Poster.Filename))
		allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}

		if !allowedExtensions[extension] {
			response.Error(ctx, http.StatusBadRequest, "Invalid image format")
			return
		}

		filename = fmt.Sprintf("%d_movie_update%s", time.Now().Unix(), extension)
		dst := filepath.Join("public/img", filename)

		if err := ctx.SaveUploadedFile(req.Poster, dst); err != nil {
			response.Error(ctx, http.StatusInternalServerError, "Failed to save new image")
			return
		}
	}

	err = c.movieService.AdminUpdateMovie(ctx.Request.Context(), movieID, req, filename)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			response.Error(ctx, http.StatusNotFound, err.Error())
			return
		}
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Update Movie Success", nil)
}
