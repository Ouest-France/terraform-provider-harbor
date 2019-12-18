package harbor

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"

	apiclient "github.com/Ouest-France/goharbor/client"
	httptransport "github.com/go-openapi/runtime/client"
)

type HarborClient struct {
	client *apiclient.Harbor
	auth   runtime.ClientAuthInfoWriter
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Harbor address",
			},
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Harbor username",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Harbor password",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"harbor_project": resourceHarborProject(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	client := apiclient.New(httptransport.New(d.Get("address").(string), "api", nil), strfmt.Default)
	basicAuth := httptransport.BasicAuth(d.Get("user").(string), d.Get("password").(string))

	return &HarborClient{client, basicAuth}, nil
}
