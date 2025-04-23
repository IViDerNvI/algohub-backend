package submit

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/internal/apiserver/config"
	pb "github.com/ividernvi/algohub/internal/apiserver/proto/submit"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *SubmitController) Create(ctx *gin.Context) {

	var requestBody struct {
		ProblemID string `json:"problem_id"`
		Code      string `json:"code"`
		Language  string `json:"language"`
		Cases     []struct {
			Input          string `json:"input"`
			ExpectedOutput string `json:"expected_output"`
		} `json:"cases"`
		TimeLimit   int `json:"time_limit"`
		MemoryLimit int `json:"memory_limit"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	var submit v1.Submit

	operatorName, exists := ctx.Get("X-Operation-User-Name")
	if !exists {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	var valid bool
	submit.Author, valid = operatorName.(string)
	if !valid {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	submit.Status = v1.SubmitStatusPending
	submit.CodeText = requestBody.Code
	submit.Language = requestBody.Language
	submit.ProblemID = requestBody.ProblemID

	if err := submit.Validate(); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	logrus.Infof("Judge endpoint: %s", config.ALGOHUB_JUDGE_RPC_ENDPOINT)
	conn, err := grpc.NewClient(config.ALGOHUB_JUDGE_RPC_ENDPOINT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	problem, err := c.Service.Problems().Get(ctx, submit.ProblemID, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	mapper := map[string]string{
		"problem_id": submit.ProblemID,
	}
	selector := v1.Selector(mapper)

	testcases, err := c.Service.Solutions().List(ctx, &v1.ListOptions{
		Offset:   0,
		Limit:    10000,
		Selector: selector,
	})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	var cases []*pb.Case
	for _, testcase := range testcases.Items {
		cases = append(cases, &pb.Case{
			Input:          testcase.TestData,
			ExpectedOutput: testcase.TestResult,
		})
	}

	client := pb.NewJudgeServiceClient(conn)
	req := &pb.Request{
		Code:        submit.CodeText,
		Language:    submit.Language,
		Cases:       cases,
		TimeLimit:   int64(problem.MemoryLimit),
		MemoryLimit: problem.MemoryLimit,
	}

	resp, err := client.Judge(ctx, req)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	submit.Status = resp.Status
	submit.Details = resp.Message

	if err := c.Service.Submits().Create(ctx, &submit, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, submit)
}
