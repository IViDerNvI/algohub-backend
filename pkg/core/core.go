package core

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		if customErr, ok := err.(interface{ ErrCode() int }); ok {
			ctx.JSON(customErr.ErrCode(), gin.H{
				"code":    customErr.ErrCode(),
				"status":  "error",
				"message": err.Error(),
				"data":    data,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
			"data":    data,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    200,
		"status":  "success",
		"message": "success",
		"data":    data,
	})
}

func WriteResponseWithFile(ctx *gin.Context, err error, file []byte) {
	if err != nil {
		if customErr, ok := err.(interface{ ErrCode() int }); ok {
			ctx.JSON(customErr.ErrCode(), gin.H{
				"code":    customErr.ErrCode(),
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	len := strconv.Itoa(len(file))

	ctx.Header("Content-Disposition", "inline; filename=avatar.png")
	ctx.Header("Content-Type", "image/png")
	ctx.Header("Content-Length", len)
	ctx.Header("Accept-Ranges", "bytes")
	ctx.Header("Cache-Control", "public, max-age=31536000")
	ctx.Data(200, "image/png", file)
}
