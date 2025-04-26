package post

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
	"github.com/ividernvi/algohub/pkg/util/idutil"
)

func (p *PostController) PutImage(ctx *gin.Context) {
	avatarBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if ctx.ContentType() != "image/png" {
		core.WriteResponse(ctx, core.ErrInvalidFileType, nil)
		return
	}

	imageUrl, err := p.Service.Subjects().Put(ctx, &v1.PutOptions{
		Realm:      "algohub-post",
		File_UUID:  fmt.Sprintf("%s-%s", ctx.Param("id"), idutil.UUID()),
		File_Bytes: avatarBytes,
		File_type:  "image/png",
	})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, imageUrl)
}
