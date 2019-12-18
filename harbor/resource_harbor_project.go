package harbor

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Ouest-France/goharbor/client/products"
	"github.com/Ouest-France/goharbor/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHarborProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceHarborProjectCreate,
		Read:   resourceHarborProjectRead,
		Delete: resourceHarborProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceHarborProjectCreate(d *schema.ResourceData, m interface{}) error {
	hc := m.(*HarborClient)

	params := products.NewPostProjectsParams()
	params.Project = &models.ProjectReq{
		ProjectName: d.Get("name").(string),
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
		return err
	}

	d.Set("name", project.Payload.Name)

	return nil
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

	params := products.NewGetProjectsParams()
	params.Name = &name

	res, err := hc.client.Products.GetProjects(params, hc.auth)
	if err != nil {
		return 0, err
	}
	projects := res.GetPayload()

	if len(projects) != 1 {
		return 0, errors.New("project not found")
	}

	return projects[0].ProjectID, nil
}
