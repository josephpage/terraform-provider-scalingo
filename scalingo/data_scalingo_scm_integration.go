package scalingo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	scalingo "github.com/Scalingo/go-scalingo/v5"
)

func dataSourceScalingoScmIntegration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceScmIntegrationRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scm_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"avatar_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"profile_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func filterScmIntegrations(ss []scalingo.SCMIntegration, test func(scalingo.SCMIntegration) bool) []scalingo.SCMIntegration {
	var ret []scalingo.SCMIntegration
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func dataSourceScmIntegrationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, _ := meta.(*scalingo.Client)

	url, _ := d.Get("url").(string)

	integrations, err := client.SCMIntegrationsList(ctx)
	if err != nil {
		return diag.Errorf("fail to fetch integrations: %v", err)
	}

	selectedIntegrations := filterScmIntegrations(integrations, func(element scalingo.SCMIntegration) bool {
		return element.URL == url
	})

	if len(selectedIntegrations) != 1 {
		return diag.Errorf("fail to find the selected integration")
	}

	integration := selectedIntegrations[0]
	err = SetAll(d, map[string]interface{}{
		"scm_type":    integration.SCMType,
		"url":         integration.URL,
		"uid":         integration.Uid,
		"username":    integration.Username,
		"avatar_url":  integration.AvatarURL,
		"email":       integration.Email,
		"profile_url": integration.ProfileURL,
	})
	if err != nil {
		return diag.Errorf("fail to store scm integration information: %v", err)
	}
	d.SetId(integration.ID)

	return nil
}
