package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/Services/albums"
	"github.com/iavorskyi/s3gallery/s3Gallery"
	"net/http"
)

func (s *APIServer) listAlbums(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	albumList, code, err := albums.ListAlbums()
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to get list of albums: "+err.Error())
		return
	}
	ctx.JSON(code, albumList)
}

func (s *APIServer) getAlbum(ctx *gin.Context) {
	albumId := ctx.Param("albumId")
	album, code, err := albums.GetAlbum(albumId)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to get album: "+albumId+". "+err.Error())
		return
	}
	ctx.JSON(code, album)
}

func (s *APIServer) createAlbum(ctx *gin.Context) {
	var album s3Gallery.Album
	err := ctx.BindJSON(&album)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, http.StatusBadRequest, "Failed to unmarshal payload: "+err.Error())
		return
	}
	code, err := albums.CreateAlbum(album.Name)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to create album: "+album.Name+". "+err.Error())
		return
	}
	ctx.JSON(code, album.Name+" is created")
}
