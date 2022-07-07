package scalingo

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	scalingo "github.com/Scalingo/go-scalingo/v4"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALINGO_API_TOKEN", nil),
			},
			"api_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALINGO_API_URL", nil),
			},
			"db_api_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALINGO_DB_API_URL", nil),
			},
			"auth_api_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALINGO_AUTH_URL", nil),
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALINGO_REGION", "osc-fr1"),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"scalingo_notification_platform": dataSourceScNotificationPlatform(),
			"scalingo_stack":                 dataSourceScStack(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"scalingo_addon":          resourceScalingoAddon(),
			"scalingo_app":            resourceScalingoApp(),
			"scalingo_collaborator":   resourceScalingoCollaborator(),
			"scalingo_container_type": resourceScalingoContainerType(),
			"scalingo_autoscaler":     resourceScalingoAutoscaler(),
			"scalingo_domain":         resourceScalingoDomain(),
			"scalingo_github_link":    resourceScalingoGithubLink(),
			"scalingo_notifier":       resourceScalingoNotifier(),
			"scalingo_run":            resourceScalingoRun(),
			"scalingo_log_drain":      resourceScalingoLogDrain(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	apiURL, _ := data.Get("api_url").(string)
	authAPIURL, _ := data.Get("auth_api_url").(string)
	dbAPIURL, _ := data.Get("db_api_url").(string)
	apiToken, _ := data.Get("api_token").(string)
	region, _ := data.Get("region").(string)

	client, err := scalingo.New(scalingo.ClientConfig{
		Region:              region,
		APIToken:            apiToken,
		APIEndpoint:         apiURL,
		DatabaseAPIEndpoint: dbAPIURL,
		AuthEndpoint:        authAPIURL,
	})
	if err != nil {
		return nil, fmt.Errorf("fail to initialize Scalingo client: %v", err)
	}

	return client, nil
}
