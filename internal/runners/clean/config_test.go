package clean

import (
	"os"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/testhelpers/outputhelper"
)

func (suite *CleanTestSuite) TestConfig_PromptNo() {
	runner := newConfig(&outputhelper.TestOutputer{}, &confirmMock{}, newConfigMock(suite.T(), suite.cachePath, suite.configPath), nil)
	err := runner.Run(&ConfigParams{})
	suite.Require().Error(err)
	suite.Require().True(locale.IsInputError(err), "non-confirmed clean config should return an error, not silently halt")

	suite.Require().DirExists(suite.configPath)
	suite.Require().DirExists(suite.cachePath)
	suite.Require().FileExists(suite.installPath)
}

func (suite *CleanTestSuite) TestConfig_Activated() {
	os.Setenv(constants.ActivatedStateEnvVarName, "true")
	defer func() {
		os.Unsetenv(constants.ActivatedStateEnvVarName)
	}()

	runner := newConfig(&outputhelper.TestOutputer{}, &confirmMock{}, newConfigMock(suite.T(), suite.cachePath, suite.configPath), nil)
	err := runner.Run(&ConfigParams{})
	suite.Require().Error(err)
}
