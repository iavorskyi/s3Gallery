package main

import (
	"github.com/gin-gonic/gin"
	s3Gallery "github.com/iavorskyi/s3gallery"
	"github.com/iavorskyi/s3gallery/Services/items"
	"net/http"
)

func listItems(ctx *gin.Context) {
	var user s3Gallery.User
	albumID := ctx.Param("albumId")
	itemList, code, err := items.ListItems(albumID)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to get list of items for user"+user.ID+err.Error())
		return
	}

	ctx.IndentedJSON(code, itemList)
}

func getItem(ctx *gin.Context) {
	isFullSize := ctx.Query("full-size")
	albumID := ctx.Param("albumId")
	itemID := ctx.Param("id")
	if isFullSize != "true" {
		itemID = "resized/" + itemID
	}

	item, code, err := items.GetItem(albumID, itemID)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, err.Error())
		return
	}

	ctx.JSON(code, item)
}

func uploadItem(ctx *gin.Context) {
	albumID := ctx.Param("albumId")
	code := http.StatusOK

	var opts items.UploadOpts
	err := ctx.BindJSON(&opts)

	err = items.UploadItem(opts, albumID)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to upload item: "+err.Error())
		return
	}
	ctx.String(code, "Done")
}

func updateItem(ctx *gin.Context) {

}

func deleteItem(ctx *gin.Context) {
	albumID := ctx.Param("albumId")
	code := http.StatusOK
	id := ctx.Param("id")

	err := items.DeleteItem(albumID, id)
	if err != nil {
		s3Gallery.NewErrorResponse(ctx, code, "Failed to delete item"+id+err.Error())
		return
	}

	ctx.JSON(code, map[string]interface{}{"msg": id + " was deleted"})
}
