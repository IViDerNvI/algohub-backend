package submit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/internal/apiserver/config"
	pb "github.com/ividernvi/algohub/internal/apiserver/proto/submit"
	"github.com/ividernvi/algohub/pkg/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *SubmitController) Judge(ctx *gin.Context) {

	var requestBody struct {
		ProblemID string `json:"problem_id"`
		Code      string `json:"code"`
		Language  string `json:"language"`
		Cases     []struct {
			Input          string `json:"input"`
			ExpectedOutput string `json:"expected_output"`
		} `json:"cases"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	problemId := requestBody.ProblemID
	problem, err := c.Service.Problems().Get(ctx, problemId, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	conn, err := grpc.NewClient(config.ALGOHUB_JUDGE_RPC_ENDPOINT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	var cases []*pb.Case
	for _, testCase := range requestBody.Cases {
		cases = append(cases, &pb.Case{
			Input: testCase.Input,
		})
	}

	client := pb.NewJudgeServiceClient(conn)
	req := &pb.Request{
		Code:        requestBody.Code,
		Language:    requestBody.Language,
		Cases:       cases,
		TimeLimit:   int64(problem.TimeLimit),
		MemoryLimit: problem.MemoryLimit,
	}

	resp, err := client.Judge(ctx, req)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, resp)
}
