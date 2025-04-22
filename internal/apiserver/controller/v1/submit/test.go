package submit

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/internal/apiserver/config"
	pb "github.com/ividernvi/algohub/internal/apiserver/proto/submit"
	"github.com/ividernvi/algohub/pkg/core"
	"google.golang.org/grpc"
)

func (c *SubmitController) Judge(ctx *gin.Context) {

	var requestBody struct {
		CodeText  string                   `json:"code_text"`
		Language  string                   `json:"language"`
		ProblemID string                   `json:"problem_id"`
		TestCase  []map[string]interface{} `json:"test_cases"`
	}

	// 解析 JSON 数据
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	problemId := ctx.Query("id")
	problem, err := c.Service.Problems().Get(ctx, problemId, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	conn, err := grpc.NewClient(config.ALGOHUB_JUDGE_RPC_ENDPOINT)
	if err != nil {
		panic(err)
	}

	var cases []*pb.Case
	for _, testCase := range requestBody.TestCase {
		cases = append(cases, &pb.Case{
			Input: func() string {
				if input, ok := testCase["test_case"].(string); ok {
					return input
				}
				return ""
			}(),
		})
	}

	client := pb.NewJudgeServiceClient(conn)
	req := &pb.Request{
		Code:        requestBody.CodeText,
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

	core.WriteResponse(ctx, nil, resp.CaseInfo.ActualOutput)
}
