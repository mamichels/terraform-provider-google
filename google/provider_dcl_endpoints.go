// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	framework_schema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// empty string is passed for dcl default since dcl
// [hardcodes the values](https://github.com/GoogleCloudPlatform/declarative-resource-client-library/blob/main/services/google/eventarc/beta/trigger_internal.go#L96-L103)

var ApikeysEndpointEntryKey = "apikeys_custom_endpoint"
var ApikeysEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var AssuredWorkloadsEndpointEntryKey = "assured_workloads_custom_endpoint"
var AssuredWorkloadsEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var CloudBuildWorkerPoolEndpointEntryKey = "cloud_build_worker_pool_custom_endpoint"
var CloudBuildWorkerPoolEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var ClouddeployEndpointEntryKey = "clouddeploy_custom_endpoint"
var ClouddeployEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var CloudResourceManagerEndpointEntryKey = "cloud_resource_manager_custom_endpoint"
var CloudResourceManagerEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var DataplexEndpointEntryKey = "dataplex_custom_endpoint"
var DataplexEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var EventarcEndpointEntryKey = "eventarc_custom_endpoint"
var EventarcEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var FirebaserulesEndpointEntryKey = "firebaserules_custom_endpoint"
var FirebaserulesEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var NetworkConnectivityEndpointEntryKey = "network_connectivity_custom_endpoint"
var NetworkConnectivityEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var OrgPolicyEndpointEntryKey = "org_policy_custom_endpoint"
var OrgPolicyEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

var RecaptchaEnterpriseEndpointEntryKey = "recaptcha_enterprise_custom_endpoint"
var RecaptchaEnterpriseEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
}

type DCLConfig struct {
	ApikeysBasePath              string
	AssuredWorkloadsBasePath     string
	CloudBuildWorkerPoolBasePath string
	ClouddeployBasePath          string
	CloudResourceManagerBasePath string
	DataplexBasePath             string
	EventarcBasePath             string
	FirebaserulesBasePath        string
	NetworkConnectivityBasePath  string
	OrgPolicyBasePath            string
	RecaptchaEnterpriseBasePath  string
}

func configureDCLProvider(provider *schema.Provider) {
	provider.Schema[ApikeysEndpointEntryKey] = ApikeysEndpointEntry
	provider.Schema[AssuredWorkloadsEndpointEntryKey] = AssuredWorkloadsEndpointEntry
	provider.Schema[CloudBuildWorkerPoolEndpointEntryKey] = CloudBuildWorkerPoolEndpointEntry
	provider.Schema[ClouddeployEndpointEntryKey] = ClouddeployEndpointEntry
	provider.Schema[CloudResourceManagerEndpointEntryKey] = CloudResourceManagerEndpointEntry
	provider.Schema[DataplexEndpointEntryKey] = DataplexEndpointEntry
	provider.Schema[EventarcEndpointEntryKey] = EventarcEndpointEntry
	provider.Schema[FirebaserulesEndpointEntryKey] = FirebaserulesEndpointEntry
	provider.Schema[NetworkConnectivityEndpointEntryKey] = NetworkConnectivityEndpointEntry
	provider.Schema[OrgPolicyEndpointEntryKey] = OrgPolicyEndpointEntry
	provider.Schema[RecaptchaEnterpriseEndpointEntryKey] = RecaptchaEnterpriseEndpointEntry
}

func HandleDCLProviderDefaults(d *schema.ResourceData) {
	if d.Get(ApikeysEndpointEntryKey) == "" {
		d.Set(ApikeysEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_APIKEYS_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(AssuredWorkloadsEndpointEntryKey) == "" {
		d.Set(AssuredWorkloadsEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_ASSURED_WORKLOADS_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(CloudBuildWorkerPoolEndpointEntryKey) == "" {
		d.Set(CloudBuildWorkerPoolEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_CLOUD_BUILD_WORKER_POOL_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(ClouddeployEndpointEntryKey) == "" {
		d.Set(ClouddeployEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_CLOUDDEPLOY_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(CloudResourceManagerEndpointEntryKey) == "" {
		d.Set(CloudResourceManagerEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_CLOUD_RESOURCE_MANAGER_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(DataplexEndpointEntryKey) == "" {
		d.Set(DataplexEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_DATAPLEX_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(EventarcEndpointEntryKey) == "" {
		d.Set(EventarcEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_EVENTARC_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(FirebaserulesEndpointEntryKey) == "" {
		d.Set(FirebaserulesEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_FIREBASERULES_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(NetworkConnectivityEndpointEntryKey) == "" {
		d.Set(NetworkConnectivityEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_NETWORK_CONNECTIVITY_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(OrgPolicyEndpointEntryKey) == "" {
		d.Set(OrgPolicyEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_ORG_POLICY_CUSTOM_ENDPOINT",
		}))
	}
	if d.Get(RecaptchaEnterpriseEndpointEntryKey) == "" {
		d.Set(RecaptchaEnterpriseEndpointEntryKey, multiEnvSearch([]string{
			"GOOGLE_RECAPTCHA_ENTERPRISE_CUSTOM_ENDPOINT",
		}))
	}
}

// plugin-framework provider set-up
func configureDCLFrameworkProvider(frameworkSchema *framework_schema.Schema) {
	frameworkSchema.Attributes["apikeys_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["assured_workloads_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["cloud_build_worker_pool_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["clouddeploy_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["cloud_resource_manager_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["dataplex_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["eventarc_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["firebaserules_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["network_connectivity_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["org_policy_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
	frameworkSchema.Attributes["recaptcha_enterprise_custom_endpoint"] = framework_schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			CustomEndpointValidator(),
		},
	}
}

func providerDCLConfigure(d *schema.ResourceData, config *Config) interface{} {
	config.ApikeysBasePath = d.Get(ApikeysEndpointEntryKey).(string)
	config.AssuredWorkloadsBasePath = d.Get(AssuredWorkloadsEndpointEntryKey).(string)
	config.CloudBuildWorkerPoolBasePath = d.Get(CloudBuildWorkerPoolEndpointEntryKey).(string)
	config.ClouddeployBasePath = d.Get(ClouddeployEndpointEntryKey).(string)
	config.CloudResourceManagerBasePath = d.Get(CloudResourceManagerEndpointEntryKey).(string)
	config.DataplexBasePath = d.Get(DataplexEndpointEntryKey).(string)
	config.EventarcBasePath = d.Get(EventarcEndpointEntryKey).(string)
	config.FirebaserulesBasePath = d.Get(FirebaserulesEndpointEntryKey).(string)
	config.NetworkConnectivityBasePath = d.Get(NetworkConnectivityEndpointEntryKey).(string)
	config.OrgPolicyBasePath = d.Get(OrgPolicyEndpointEntryKey).(string)
	config.RecaptchaEnterpriseBasePath = d.Get(RecaptchaEnterpriseEndpointEntryKey).(string)
	config.CloudBuildWorkerPoolBasePath = d.Get(CloudBuildWorkerPoolEndpointEntryKey).(string)
	return config
}
