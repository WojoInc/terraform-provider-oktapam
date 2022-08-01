package descriptions

import (
	"fmt"

	"github.com/okta/terraform-provider-oktapam/oktapam/constants/typed_strings"
)

var (
	// Attribute Descriptions
	AccessAddress                    = "Access Address of the gateway."
	AccessAddressAttribute           = "AD Attribute mapped to IP address or DNS name used by the gateway to connect to a discovered server."
	ADConnectionID                   = "UUID of the AD Connection with which this AD Task Settings is associated."
	AdditionalAttributeMapping       = "Additional AD attributes mappings to Advanced Server Access labels."
	AdditionalAttributeMappingLabel  = "ASA label"
	AdditionalAttributeMappingValue  = "AD Attribute mapped to ASA label"
	AdditionalAttributeMappingIsGuid = "Identifies an AD attribute as a Globally Unique Identifier (GUID)"
	ADRuleAssignments                = "Assignment rules determine how servers are synced from Active Directory (AD) and assigned to projects."
	ADRuleAssignmentsBaseDN          = "Specifies where the rule searches for servers."
	ADRuleAssignmentsLDAPQueryFilter = "Specifies the specific criteria used to filter servers."
	ADRuleAssignmentsProjectID       = "Specifies a project to associate with matching servers"
	ADTaskFrequency                  = "Frequency of the AD Task"
	ADTaskIsActive                   = "If 'true', enables AD task"
	ADTaskRunTest                    = "If 'true', test is performed based on specified AD Task Settings"
	ADTaskStartHourUTC               = "If AD task is scheduled to run daily, then specify start hour in UTC"
	AltNamesAttributes               = "AD Attribute mapped to alternative hostnames or DNS entries used to resolve a discovered server."
	BastionAttribute                 = "AD Attribute mapped to bastion host that Advanced Server Access clients can use to tunnel traffic to a discovered server."
	CertificateCommonName            = "Common Name or FQDN to which certificate is issued to."
	CertificateContent               = "Certificate Signing Request (CSR)/ Self Signed Certificate content."
	CertificateID                    = "Certificate ID used for password less access method."
	CertificateRequestType           = "Specifies the type of certificate request - Certificate Signing Request (CSR)/ Self Signed Certificate."
	CertificateStatus                = "Certificate status - Valid/Request Created."
	CloudProvider                    = "Cloud provider name of the host where gateway is running."
	ClusterGroupClaims               = "A map of claims that will be included in credentials issued to users that are used to authenticate to Kubernetes clusters. Claims correspond to pre-configured role bindings on the cluster."
	ClusterSelector                  = "A label selector to used to match Kubernetes clusters."
	CreateServerGroup                = "If 'true', `sftd` (ASA Server Agent) creates a corresponding local (unix or windows) group in the ASA Project's servers."
	CreateServerUsers                = "If 'true', `sftd` (ASA Server Agent) creates corresponding local (unix or windows) user accounts in the ASA Project's servers."
	CreatedAt                        = "The UTC time of resource deletion. Format is `2022-01-01 00:00:00 +0000 UTC`."
	CreatedByUser                    = "The ASA User that created the resource."
	CSRDetails                       = "Certificate Signing Request (CSR) Details."
	DefaultAddress                   = "Default address of the gateway."
	DeletedAt                        = "The UTC time of resource creation. Format is `2022-01-01 00:00:00 +0000 UTC`."
	Description                      = "The human-readable description of the resource."
	DescriptionContains              = "If a value is provided, the results are filtered to only contain resources whose description contains that value."
	Domain                           = "The domain against which to query."
	DomainControllers                = "A comma-separated list of the specific domain controller(s) that should be used to query the domain. Can be specified as a hostname or IP."
	EnterpriseSigned                 = "If 'true', certificate is signed by AD Certificate Services."
	ForwardTraffic                   = "If 'true', all traffic in the ASA Project be forwarded through selected ASA Gateways."
	GatewayID                        = "The UUID of the Gateway with which this AD Connection is associated."
	GatewaySelector                  = "Assigns ASA Gateways with labels matching all selectors. At least one selector is required for traffic forwarding."
	GroupName                        = "The human-readable name of the ASA Group. Values are case-sensitive."
	HostnameAttribute                = "AD Attribute mapped to hostname used to identify a discovered server within Advanced Server Access."
	IssuedAt                         = "The UTC issuance time of the resource. Format is `2022-01-01 00:00:00 +0000 UTC`."
	KubernetesAuthMechanism          = "Mechanism to provide auth details to your Kubernetes cluster (eg. OIDC_RSA2048, NONE)"
	KubernetesAPIURL                 = "The URL to access the control plane of the Kubernetes cluster."
	KubernetesClusterID              = "The ASA ID of a Kubernetes cluster."
	KubernetesClusterKey             = "The human-friendly key to associate with the Kubernetes cluster. Must be simple alphanumeric without spaces."
	KubernetesClusterLabels          = "Map of labels to assign to the Kubernetes cluster."
	KubernetesPublicCertificate      = "The certificate expected when connecting to the Kubernetes cluster."
	Labels                           = "A map of key-value pairings that define access to the ASA Gateway."
	Name                             = "The human-readable name of the resource. Values are case-sensitive."
	NextUnixGID                      = "The GID to use when creating a new ASA Server User. Default value starts at 63001."
	NextUnixUID                      = "The UID to use when creating a new ASA Server User. Default value starts at 60001."
	OIDCIssuerURL                    = "The OIDC Issuer URL to use when configuring your Kubernetes cluster. "
	OSAttribute                      = "AD Attribute mapped to server operating system of a discovered server."
	ProjectName                      = "The human-readable name of the ASA Project. Values are case-sensitive."
	RDPSessionRecording              = "If 'true', enable remote desktop protocol (RDP) recording on all servers in the ASA Project."
	RefuseConnections                = "If 'true', gateway refuse connection."
	RemovedAt                        = "UTC time of resource removal from parent resource. Format is `2022-01-01 00:00:00 +0000 UTC`."
	RequirePreauthForCreds           = "If 'true', require preauthorization before an ASA User can retrieve credentials to sign in."
	Roles                            = "A list of roles for the ASA Group. Options are `access_user`, `access_admin`, and `reporting_user`."
	ServerAccess                     = "If 'true', members of this ASA Group have access to the ASA Project servers."
	ServerAdmin                      = "If 'true', members of ASA Group have sudo permissions on ASA Project servers."
	ServerUserName                   = "The name of the corresponding ASA Server User."
	ServersSelector                  = "Enables access to ASA Servers with labels matching all selectors. Only available to customers that have the Early Access Policy Sync feature enabled on their team."
	ServiceAccountUsername           = "The username of the service account that can be used to query the domain."
	ServiceAccountPassword           = "The password of the service account that can be used to query the domain."
	SSHCertificateType               = fmt.Sprintf("The SSH certificate type used by access requests. Options include: [%s]. `%s` is a deprecated key algorithm type. "+
		"This option should only be used to connect to legacy systems that cannot use newer SSH versions. If you do need to use `%s`, it is recommended to connect via a gateway with traffic forwarding. "+
		"Otherwise, please use a more current key algorithm. If left unspecified, `%s` is used by default.", typed_strings.SSHCertTypeListFormat(), typed_strings.CertTypeRsa, typed_strings.CertTypeRsa, typed_strings.CertTypeEd25519)
	SSHSessionRecording = "If 'true', enables ssh recording on server access requests."
	Status              = "The status of the ASA User. Valid statuses are `ACTIVE`, `DISABLED`, and `DELETED`."
	TeamName            = "The human-readable name of the ASA Team that owns the resource. Values are lower-case."
	Token               = "The secret used for resource enrollment."
	UsePasswordless     = "if 'true', Users will not need password to login."
	UserType            = "The user type. Valid types are `human` and `service`."
)
