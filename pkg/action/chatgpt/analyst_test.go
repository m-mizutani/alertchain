package chatgpt_test

import (
	_ "embed"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/m-mizutani/alertchain/pkg/action/chatgpt"
	"github.com/m-mizutani/alertchain/pkg/domain/model"
	"github.com/m-mizutani/gt"
	"github.com/sashabaranov/go-openai"
)

//go:embed testdata/aws_guardduty_example.json
var alertData []byte

func TestAnalystInquiry(t *testing.T) {
	if _, ok := os.LookupEnv("TEST_CHATGPT_ANALYST"); !ok {
		t.Skip("Skipping test because TEST_CHATGPT_ANALYST is not set")
	}

	var body any
	gt.NoError(t, json.Unmarshal(alertData, &body))

	ctx := model.NewContext()
	alert := model.Alert{
		AlertMetaData: model.AlertMetaData{
			Title:       "test",
			Description: "test",
		},
		Data:      body,
		CreatedAt: time.Now(),
	}

	resp := gt.R1(chatgpt.Query(ctx, alert, model.ActionArgs{
		"secret_api_key": os.Getenv("TEST_CHATGPT_API_KEY"),
	})).NoError(t)
	data := gt.Cast[openai.ChatCompletionResponse](t, resp)
	gt.A(t, data.Choices).Length(1)
	t.Log(data.Choices[0].Message.Content)
}
