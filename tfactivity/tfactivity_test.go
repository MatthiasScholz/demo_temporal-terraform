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

// func TestUnitTestSuite(t *testing.T) {
// 	suite.Run(t, new(UnitTestSuite))
// }
//
// func (s *UnitTestSuite) SetupTest() {
// 	s.env = s.NewTestActivityEnvironment()
// }
//
//func (s *UnitTestSuite) Test_Apply(t *testing.T) {
func Test_Apply(t *testing.T) {
	vars := map[string]interface{}{
		"cidr_block": "10.0.0.0/16",
		"name":       "Test_Apply",
	}
	creds, err := awsconfig.LoadConfig("tw-beach-push").Credentials.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("failed to retrieve aws credentials: %e", err)
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

	suite := UnitTestSuite{}
	suite.env = suite.NewTestActivityEnvironment()
	suite.env.RegisterActivity(a.Apply)
	result, err := suite.env.ExecuteActivity(a.Apply, input)

	if err != nil {
		t.Fatalf("An error occured during execution: %e", err)
	}
	if !result.HasValue() {
		t.Fatalf("Result as no value: %v, %e", result, err)
	}
	//s.env.RegisterActivity(a.Apply)
	//result, err := s.env.ExecuteActivity(a.Apply, input)

	//s.True(result.HasValue())
	//s.NoError(err)
}
