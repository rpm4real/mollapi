package actions

func (as *ActionSuite) Test_Health_Check() {
	result := as.JSON("/health/check").Get()

    as.Equal(200, result.Code)
    as.Contains(result.Body.String(), "Up and running")
}

