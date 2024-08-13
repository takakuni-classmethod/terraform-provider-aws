// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package rds

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	rds_sdkv2 "github.com/aws/aws-sdk-go-v2/service/rds"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	rds_sdkv1 "github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newClusterParameterGroupDataSource,
			Name:    "Cluster Parameter Group",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newIntegrationResource,
			Name:    "Integration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory: newResourceExportTask,
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceClusterSnapshot,
			TypeName: "aws_db_cluster_snapshot",
			Name:     "DB Cluster Snapshot",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceEventCategories,
			TypeName: "aws_db_event_categories",
			Name:     "Event Categories",
		},
		{
			Factory:  dataSourceInstance,
			TypeName: "aws_db_instance",
			Name:     "DB Instance",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceInstances,
			TypeName: "aws_db_instances",
			Name:     "DB Instances",
		},
		{
			Factory:  dataSourceParameterGroup,
			TypeName: "aws_db_parameter_group",
			Name:     "DB Parameter Group",
		},
		{
			Factory:  dataSourceProxy,
			TypeName: "aws_db_proxy",
			Name:     "DB Proxy",
		},
		{
			Factory:  dataSourceSnapshot,
			TypeName: "aws_db_snapshot",
			Name:     "DB Snapshot",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceSubnetGroup,
			TypeName: "aws_db_subnet_group",
			Name:     "DB Subnet Group",
		},
		{
			Factory:  dataSourceCertificate,
			TypeName: "aws_rds_certificate",
			Name:     "Certificate",
		},
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_rds_cluster",
			Name:     "Cluster",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceClusters,
			TypeName: "aws_rds_clusters",
			Name:     "Clusters",
		},
		{
			Factory:  dataSourceEngineVersion,
			TypeName: "aws_rds_engine_version",
			Name:     "Engine Version",
		},
		{
			Factory:  dataSourceOrderableInstance,
			TypeName: "aws_rds_orderable_db_instance",
			Name:     "Orderable DB Instance",
		},
		{
			Factory:  dataSourceReservedOffering,
			TypeName: "aws_rds_reserved_instance_offering",
			Name:     "Reserved Instance Offering",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceClusterSnapshot,
			TypeName: "aws_db_cluster_snapshot",
			Name:     "DB Cluster Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "db_cluster_snapshot_arn",
			},
		},
		{
			Factory:  resourceEventSubscription,
			TypeName: "aws_db_event_subscription",
			Name:     "Event Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceInstance,
			TypeName: "aws_db_instance",
			Name:     "DB Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceInstanceAutomatedBackupsReplication,
			TypeName: "aws_db_instance_automated_backups_replication",
			Name:     "Instance Automated Backups Replication",
		},
		{
			Factory:  resourceInstanceRoleAssociation,
			TypeName: "aws_db_instance_role_association",
			Name:     "DB Instance IAM Role Association",
		},
		{
			Factory:  resourceOptionGroup,
			TypeName: "aws_db_option_group",
			Name:     "DB Option Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceParameterGroup,
			TypeName: "aws_db_parameter_group",
			Name:     "DB Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceProxy,
			TypeName: "aws_db_proxy",
			Name:     "DB Proxy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceProxyDefaultTargetGroup,
			TypeName: "aws_db_proxy_default_target_group",
			Name:     "DB Proxy Default Target Group",
		},
		{
			Factory:  resourceProxyEndpoint,
			TypeName: "aws_db_proxy_endpoint",
			Name:     "DB Proxy Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceProxyTarget,
			TypeName: "aws_db_proxy_target",
			Name:     "DB Proxy Target",
		},
		{
			Factory:  resourceSnapshot,
			TypeName: "aws_db_snapshot",
			Name:     "DB Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "db_snapshot_arn",
			},
		},
		{
			Factory:  resourceSnapshotCopy,
			TypeName: "aws_db_snapshot_copy",
			Name:     "DB Snapshot Copy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "db_snapshot_arn",
			},
		},
		{
			Factory:  resourceSubnetGroup,
			TypeName: "aws_db_subnet_group",
			Name:     "DB Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCertificate,
			TypeName: "aws_rds_certificate",
			Name:     "Default Certificate",
		},
		{
			Factory:  resourceCluster,
			TypeName: "aws_rds_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceClusterActivityStream,
			TypeName: "aws_rds_cluster_activity_stream",
			Name:     "Cluster Activity Stream",
		},
		{
			Factory:  resourceClusterEndpoint,
			TypeName: "aws_rds_cluster_endpoint",
			Name:     "Cluster Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceClusterInstance,
			TypeName: "aws_rds_cluster_instance",
			Name:     "Cluster Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceClusterParameterGroup,
			TypeName: "aws_rds_cluster_parameter_group",
			Name:     "Cluster Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceClusterRoleAssociation,
			TypeName: "aws_rds_cluster_role_association",
			Name:     "Cluster IAM Role Association",
		},
		{
			Factory:  resourceCustomDBEngineVersion,
			TypeName: "aws_rds_custom_db_engine_version",
			Name:     "Custom DB Engine Version",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGlobalCluster,
			TypeName: "aws_rds_global_cluster",
			Name:     "Global Cluster",
		},
		{
			Factory:  resourceReservedInstance,
			TypeName: "aws_rds_reserved_instance",
			Name:     "Reserved Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RDS
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*rds_sdkv1.RDS, error) {
	sess := config[names.AttrSession].(*session_sdkv1.Session)

	cfg := aws_sdkv1.Config{}

	if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
		tflog.Debug(ctx, "setting endpoint", map[string]any{
			"tf_aws.endpoint": endpoint,
		})
		cfg.Endpoint = aws_sdkv1.String(endpoint)
	} else {
		cfg.EndpointResolver = newEndpointResolverSDKv1(ctx)
	}

	return rds_sdkv1.New(sess.Copy(&cfg)), nil
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*rds_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return rds_sdkv2.NewFromConfig(cfg,
		rds_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
