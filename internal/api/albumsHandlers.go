package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/Services/albums"
	"github.com/iavorskyi/s3gallery/internal/model"
	"github.com/iavorskyi/s3gallery/s3Gallery"
	"net/http"
)

func (s *server) listAlbums(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	albumList, code, err := albums.ListAlbums(s.s3store)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to get list of albums: "+err.Error())
		return
	}
	ctx.JSON(code, albumList)
}

func (s *server) getAlbum(ctx *gin.Context) {
	albumId := ctx.Param("albumId")
	album, code, err := albums.GetAlbum(albumId, s.s3store)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to get album: "+albumId+". "+err.Error())
		return
	}
	ctx.JSON(code, album)
}

func (s *server) createAlbum(ctx *gin.Context) {
	var album model.Album
	err := ctx.BindJSON(&album)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, http.StatusBadRequest, "Failed to unmarshal payload: "+err.Error())
		return
	}
	code, err := albums.CreateAlbum(album.Name, s.s3store)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to create album: "+album.Name+". "+err.Error())
		return
	}
	ctx.JSON(code, album.Name+" is created")
}
