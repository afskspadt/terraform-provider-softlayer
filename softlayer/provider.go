package softlayer

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/riethm/gopherlayer.git/session"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOFTLAYER_USERNAME", nil),
				Description: "The user name for SoftLayer API operations.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOFTLAYER_API_KEY", nil),
				Description: "The API key for SoftLayer API operations.",
			},
			"endpoint_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOFTLAYER_ENDPOINT_URL", session.DefaultEndpoint),
				Description: "The endpoint url for the SoftLayer API.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"softlayer_virtual_guest":          resourceSoftLayerVirtualGuest(),
			"softlayer_ssh_key":                resourceSoftLayerSSHKey(),
			"softlayer_dns_domain_record":      resourceSoftLayerDnsDomainRecord(),
			"softlayer_dns_domain":             resourceSoftLayerDnsDomain(),
			"softlayer_lb_vpx":                 resourceSoftLayerLbVpx(),
			"softlayer_lb_vpx_vip":             resourceSoftLayerLbVpxVip(),
			"softlayer_lb_vpx_service":         resourceSoftLayerLbVpxService(),
			"softlayer_lb_local":               resourceSoftLayerLbLocal(),
			"softlayer_lb_local_service_group": resourceSoftLayerLbLocalServiceGroup(),
			"softlayer_lb_local_service":       resourceSoftLayerLbLocalService(),
			"softlayer_security_certificate":   resourceSoftLayerSecurityCertificate(),
			"softlayer_user":                   resourceSoftLayerUserCustomer(),
			"softlayer_objectstorage_account":  resourceSoftLayerObjectStorageAccount(),
			"softlayer_provisioning_hook":      resourceSoftLayerProvisioningHook(),
			"softlayer_scale_policy":           resourceSoftLayerScalePolicy(),
			"softlayer_scale_group":            resourceSoftLayerScaleGroup(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	sess := session.Session{
		UserName: d.Get("username").(string),
		APIKey:   d.Get("api_key").(string),
		Endpoint: d.Get("endpoint_url").(string),
	}

	if os.Getenv("TF_LOG") != "" {
		sess.Debug = true
	}

	return &sess, nil
}
