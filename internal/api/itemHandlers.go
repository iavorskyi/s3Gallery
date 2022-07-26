package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iavorskyi/s3gallery/Services/items"
	"github.com/iavorskyi/s3gallery/s3Gallery"
	"net/http"
	"os"
)

func (s *server) listItems(ctx *gin.Context) {
	albumID := ctx.Param("albumId")
	itemList, code, err := items.ListItems(albumID, s.s3store)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, err.Error())
		return
	}
	ctx.IndentedJSON(code, itemList)
}

func (s *server) getItem(ctx *gin.Context) {
	isFullSize := ctx.Query("full-size")
	albumID := ctx.Param("albumId")
	itemID := ctx.Param("id")
	if isFullSize != "true" {
		itemID = "resized/" + itemID
	}

	item, code, err := items.GetItem(albumID, itemID, s.s3store)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, err.Error())
		return
	}
	ctx.JSON(code, item)
}

func (s *server) uploadItem(ctx *gin.Context) {
	albumID := ctx.Param("albumId")
	code := http.StatusOK

	file, err := ctx.FormFile("image")
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to get file from request: "+err.Error())
		return
	}

	path := file.Filename
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, http.StatusInternalServerError, "Failed to save file: "+err.Error())
		return
	}
	defer os.Remove(path)

	err = items.UploadItem(file.Filename, path, albumID, s.s3store)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to upload item: "+err.Error())
		return
	}
	ctx.String(code, "Done")
}

func (s *server) updateItem(ctx *gin.Context) {

}

func (s *server) deleteItem(ctx *gin.Context) {
	albumID := ctx.Param("albumId")
	code := http.StatusOK
	id := ctx.Param("id")

	err := items.DeleteItem(albumID, id, s.s3store)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to delete item"+id+err.Error())
		return
	}

	ctx.JSON(code, map[string]interface{}{"msg": id + " was deleted"})
}
