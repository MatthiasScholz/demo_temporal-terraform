package tfactivity

import (
	"context"
	"testing"

	"github.com/dynajoe/temporal-terraform-demo/config/awsconfig"
	"github.com/dynajoe/temporal-terraform-demo/terraform"

	//"github.com/dynajoe/temporal-terraform-demo/tfactivity"
	"github.com/dynajoe/temporal-terraform-demo/tfworkspace"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type UnitTestSuite struct {
	suite.Suite

	// WorkflowTestSuite usable for workflows and activities
	testsuite.WorkflowTestSuite

	env *testsuite.TestActivityEnvironment
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func (s *UnitTestSuite) SetupTest() {
	s.env = s.NewTestActivityEnvironment()
}

func (s *UnitTestSuite) Test_Plan() {
	// Prepare
	vars := map[string]interface{}{
		"cidr_block": "10.0.0.0/16",
		"name":       "Test_Apply",
	}
	creds, err := awsconfig.LoadConfig("tw-beach-push").Credentials.Retrieve(context.Background())
	if err != nil {
		s.FailNowf("credentials", "failed to retrieve aws credentials: %e", err)
	}
	input := tfworkspace.ApplyInput{
		Vars:           vars,
		AwsCredentials: creds,
	}
	ws := tfworkspace.Config{
		TerraformPath: "aws/vpc",
		TerraformFS:   terraform.FS,
	}
	a := New(ws)
	s.env.RegisterActivity(a.Plan)

	// Execute
	result, err := s.env.ExecuteActivity(a.Plan, input)

	// Verify
	s.True(result.HasValue())
	s.NoError(err)
}
