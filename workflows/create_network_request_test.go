package workflows

import (
	// "context"
	// "errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	// "go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"
)

type UnitTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func (s *UnitTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *UnitTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func (s *UnitTestSuite) Test_CreateNetworkRequestWorkflow() {
	input := CreateNetworkRequestWorkflowInput{NetworkName: "test_success"}
	s.env.ExecuteWorkflow(CreateNetworkRequestWorkflow, &input)

	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())
}

func (s *UnitTestSuite) Test_SimpleWorkflow_ActivityParamCorrect() {
	s.env.OnActivity(CreateNetworkRequestWorkflow, mock.Anything, mock.Anything).Return(
		func(ctx workflow.Context, value *CreateNetworkRequestWorkflowInput) (*CreateNetworkRequestWorkflowResult, error) {
			s.Equal("test_success", value.NetworkName)
			result := CreateNetworkRequestWorkflowResult{NetworkName: value.NetworkName}
			return &result, nil
		})
	input := CreateNetworkRequestWorkflowInput{NetworkName: "test_success"}
	s.env.ExecuteWorkflow(CreateNetworkRequestWorkflow, &input)

	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())
}
