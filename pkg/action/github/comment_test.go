package github_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/m-mizutani/alertchain/pkg/action/github"
	"github.com/m-mizutani/alertchain/pkg/domain/model"
	"github.com/m-mizutani/gt"
)

func TestComment(t *testing.T) {
	if _, ok := os.LookupEnv("TEST_GITHUB_COMMENT"); !ok {
		t.Skip("Skipping test because TEST_GITHUB_ISSUER is not set")
	}

	args := model.ActionArgs{
		"app_id":             float64(gt.R1(strconv.Atoi(os.Getenv("TEST_GITHUB_APP_ID"))).NoError(t)),
		"install_id":         float64(gt.R1(strconv.Atoi(os.Getenv("TEST_GITHUB_INSTALL_ID"))).NoError(t)),
		"secret_private_key": os.Getenv("TEST_GITHUB_PRIVATE_KEY"),
		"owner":              os.Getenv("TEST_GITHUB_OWNER"),
		"repo":               os.Getenv("TEST_GITHUB_REPO"),
		"issue_number":       float64(gt.R1(strconv.Atoi(os.Getenv("TEST_GITHUB_ISSUE_NUMBER"))).NoError(t)),
		"body":               "this is test comment\n",
	}
	alert := model.NewAlert(model.AlertMetaData{}, "test_schema", "test_raw")

	ctx := model.NewContext()
	resp := gt.R1(github.CreateComment(ctx, alert, args)).NoError(t)
	gt.V(t, resp).NotNil()
}
