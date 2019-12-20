package harbor

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Ouest-France/goharbor/client/products"
	"github.com/Ouest-France/goharbor/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHarborProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborProjectCreate,
		Read:   resourceHarborProjectRead,
		Update: resourceHarborProjectUpdate,
		Delete: resourceHarborProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"auto_scan": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"content_trust": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"prevent_vulnerability": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "low",
			},
		},
	}
}

func resourceHarborProjectCreate(d *schema.ResourceData, m interface{}) error {
	hc := m.(*HarborClient)

	params := products.NewPostProjectsParams()
	params.Project = &models.ProjectReq{
		ProjectName: d.Get("name").(string),
		Metadata: &models.ProjectMetadata{
			Public:             strconv.FormatBool(d.Get("public").(bool)),
			AutoScan:           strconv.FormatBool(d.Get("auto_scan").(bool)),
			PreventVul:         strconv.FormatBool(d.Get("prevent_vulnerability").(bool)),
			Severity:           d.Get("severity").(string),
			EnableContentTrust: strconv.FormatBool(d.Get("content_trust").(bool)),
		},
	}

	_, err := hc.client.Products.PostProjects(params, hc.auth)
	if err != nil {
		return err
	}

	id, err := getProjectID(d.Get("name").(string), d, m)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%d", id))

	return resourceHarborProjectRead(d, m)
}

func resourceHarborProjectRead(d *schema.ResourceData, m interface{}) error {
	hc := m.(*HarborClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	params := products.NewGetProjectsProjectIDParams()
	params.ProjectID = id

	project, err := hc.client.Products.GetProjectsProjectID(params, hc.auth)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			d.SetId("")
			return nil
		}
		return err
	}

	public, err := strconv.ParseBool(project.Payload.Metadata.Public)
	if err != nil {
		return err
	}

	autoscan, err := strconv.ParseBool(project.Payload.Metadata.AutoScan)
	if err != nil {
		return err
	}

	preventVul, err := strconv.ParseBool(project.Payload.Metadata.PreventVul)
	if err != nil {
		return err
	}

	contentTrust, err := strconv.ParseBool(project.Payload.Metadata.EnableContentTrust)
	if err != nil {
		return err
	}

	attributes := map[string]interface{}{
		"name":                  project.Payload.Name,
		"public":                public,
		"auto_scan":             autoscan,
		"prevent_vulnerability": preventVul,
		"severity":              project.Payload.Metadata.Severity,
		"content_trust":         contentTrust,
	}
	for key, val := range attributes {
		err = d.Set(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceHarborProjectUpdate(d *schema.ResourceData, m interface{}) error {
	hc := m.(*HarborClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	params := products.NewPutProjectsProjectIDParams()
	params.ProjectID = id
	params.Project = &models.ProjectReq{
		Metadata: &models.ProjectMetadata{
			Public:     strconv.FormatBool(d.Get("public").(bool)),
			AutoScan:   strconv.FormatBool(d.Get("auto_scan").(bool)),
			PreventVul: strconv.FormatBool(d.Get("prevent_vulnerability").(bool)),
			Severity:   d.Get("severity").(string),
		},
	}

	_, err = hc.client.Products.PutProjectsProjectID(params, hc.auth)
	if err != nil {
		return err
	}

	return resourceHarborProjectRead(d, m)
}

func resourceHarborProjectDelete(d *schema.ResourceData, m interface{}) error {
	hc := m.(*HarborClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	params := products.NewDeleteProjectsProjectIDParams()
	params.ProjectID = id

	_, err = hc.client.Products.DeleteProjectsProjectID(params, hc.auth)

	return err
}

func getProjectID(name string, d *schema.ResourceData, m interface{}) (int32, error) {
	hc := m.(*HarborClient)

	params := products.NewGetSearchParams()
	params.Q = name

	res, err := hc.client.Products.GetSearch(params, hc.auth)
	if err != nil {
		return 0, err
	}

	if len(res.GetPayload().Project) == 0 {
		return 0, errors.New("project not found")
	}

	for _, project := range res.GetPayload().Project {
		if project.Name == name {
			return project.ProjectID, nil
		}
	}

	return 0, fmt.Errorf("project not found in %d result(s)", len(res.GetPayload().Project))
}
