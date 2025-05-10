package comment

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
	"github.com/ividernvi/algohub/pkg/util/idutil"
	"github.com/sirupsen/logrus"
)

func (p *CommentController) PutImage(ctx *gin.Context) {
	dataBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if ctx.ContentType() != "image/png" && ctx.ContentType() != "image/jpeg" && ctx.ContentType() != "image/jpg" && ctx.ContentType() != "image/gif" {
		logrus.Warnf("invalid file type: %s", ctx.ContentType())
		core.WriteResponse(ctx, core.ErrInvalidFileType, nil)
		return
	}

	imageUrl, err := p.Service.Subjects().Put(ctx, &v1.PutOptions{
		Realm:      "algohub-comment",
		File_UUID:  fmt.Sprintf("%s-%s-%s", ctx.Param("id"), ctx.Param("commentid"), idutil.UUID()),
		File_Bytes: dataBytes,
		File_type:  ctx.ContentType(),
	})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, imageUrl)
}
