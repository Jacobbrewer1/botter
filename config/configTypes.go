package config

type (
	JsonConfigStruct struct {
		ConfigIpAddress *string          `json:"ConfigIpAddress"`
		BotPrefix       *string          `json:"BotPrefix"`
		Endpoints       *EndpointsStruct `json:"Endpoints"`
	}

	ApiSecretsStruct struct {
		BotToken               *string `json:"BotToken"`
		GithubApiToken         *string `json:"GithubApiToken"`
		GiphyApiToken          *string `json:"GiphyApiToken"`
		JdoodleApiClientId     *string `json:"JdoodleApiClientId"`
		JdoodleApiClientSecret *string `json:"JdoodleApiClientSecret"`
	}

	overrideStruct struct {
		Secrets            *ApiSecretsStruct `json:"Secrets"`
		BotPrefix          *string           `json:"BotPrefix,omitempty"`
		IgnoreVerification *bool             `json:"IgnoreVerification,omitempty"`
		IgnoreBadWords     *bool             `json:"IgnoreBadWords,omitempty"`
		IgnoreInvites      *bool             `json:"IgnoreInvites,omitempty"`
	}

	EndpointsStruct struct {
		GithubApiEndpoint     *string `json:"GithubApiEndpoint,omitempty"`
		GiphyApiEndpoint      *string `json:"GiphyApiEndpoint,omitempty"`
		FormulaOneApiEndpoint *string `json:"FormulaOneApiEndpoint,omitempty"`
		JdoodleApiEndpoint    *string `json:"JdoodleApiEndpoint,omitempty"`
	}
)
