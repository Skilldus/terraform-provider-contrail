package resources

import (
	"fmt"

	"github.com/Skilldus/contrail-go-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"log"
)

var mutexKV = mutexkv.NewMutexKV()

// Provider is a terraform provider implementation
func Provider() *schema.Provider {
	log.Printf("Resources map (for custom Contrail provider): %+v", ContrailResourcesMap)
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONTRAIL_API_SERVER", nil),
				Description: descriptions["server"],
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     8082,
				Description: descriptions["port"],
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USERNAME", ""),
				Description: descriptions["username"],
			},
			"tenant_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"OS_TENANT_NAME", "OS_PROJECT_NAME"}, ""),
				Description: descriptions["tenant_name"],
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PASSWORD", ""),
				Description: descriptions["password"],
			},
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"OS_AUTH_TOKEN", "OS_TOKEN"}, ""),
				Description: descriptions["token"],
			},
			"auth_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_URL", nil),
				Description: descriptions["auth_url"],
			},
			"domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_DOMAIN_NAME", nil),
				Description: descriptions["domain_name"],
			},
			"project_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_NAME", nil),
				Description: descriptions["project_name"],
			},
			"project_domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PROJECT_DOMAIN_NAME", nil),
				Description: descriptions["project_domain_name"],
			},
		},
		ResourcesMap:  ContrailResourcesMap,
		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"server":              "Contrail API server IP address.",
		"port":                "Contrail API server port number (default 8082).",
		"username":            "Username to login with.",
		"tenant_name":         "The name of the Tenant or Project to login with.",
		"password":            "Password to login with.",
		"token":               "Authentication token to use as an alternative to username/password.",
		"auth_url":            "The Identity authentication URL.",
		"domain_name":         "The name of the domain to login with.",
		"project_name":        "The name of the project.",
		"project_domain_name": "The name of the project domain to login with.",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := contrail.NewClient(d.Get("server").(string), d.Get("port").(int))
	log.Printf("Configuring Contrail client: %v:%v", client.GetServer(), d.Get("port").(int))
	if authURLRaw, isok := d.GetOk("auth_url"); isok {
		authURL := authURLRaw.(string)
		tenant := d.Get("tenant_name").(string)
		username := d.Get("username").(string)
		password := d.Get("password").(string)
		token := d.Get("token").(string)
		domain_name := d.Get("domain_name").(string)
		project_name := d.Get("project_name").(string)
		project_domain_name := d.Get("project_domain_name").(string)
		keyst := contrail.NewKeystoneClient(authURL, tenant, username, password, token, domain_name, project_name, project_domain_name)
		if err := keyst.AuthenticateV3(); err != nil {
			return nil, fmt.Errorf("Authentication error: %v", err)
		}
		client.SetAuthenticator(keyst)
		log.Printf("Using keystone authenticator: '%v' for '%v' as '%v' and token '%v'", authURL, tenant, username, token)
	}
	return client, nil
}
