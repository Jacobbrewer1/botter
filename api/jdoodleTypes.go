package api

type (
	ExecuteInput struct {
		ClientId     *string `json:"clientId"`
		ClientSecret *string `json:"clientSecret"`
		Script       *string `json:"script,omitempty"`
		Language     *string `json:"language"`
		VersionIndex *string `json:"versionIndex,omitempty"`
	}

	ExecuteOutput struct {
		Output     *string `json:"output,omitempty"`
		StatusCode *int64  `json:"statusCode,omitempty"`
		Memory     *string `json:"memory,omitempty"`
		CpuTime    *string `json:"cpuTime,omitempty"`
		Error      *string `json:"error,omitempty"`
	}
)
