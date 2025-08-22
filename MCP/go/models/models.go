package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// UpdateServerRequest represents the UpdateServerRequest schema from the OpenAPI specification
type UpdateServerRequest struct {
	Servername interface{} `json:"ServerName"`
	Backupretentioncount interface{} `json:"BackupRetentionCount,omitempty"`
	Disableautomatedbackup interface{} `json:"DisableAutomatedBackup,omitempty"`
	Preferredbackupwindow string `json:"PreferredBackupWindow,omitempty"` // <p> <code>DDD:HH:MM</code> (weekly start time) or <code>HH:MM</code> (daily start time). </p> <p> Time windows always use coordinated universal time (UTC). Valid strings for day of week (<code>DDD</code>) are: <code>Mon</code>, <code>Tue</code>, <code>Wed</code>, <code>Thr</code>, <code>Fri</code>, <code>Sat</code>, or <code>Sun</code>.</p>
	Preferredmaintenancewindow string `json:"PreferredMaintenanceWindow,omitempty"` // <p> <code>DDD:HH:MM</code> (weekly start time) or <code>HH:MM</code> (daily start time). </p> <p> Time windows always use coordinated universal time (UTC). Valid strings for day of week (<code>DDD</code>) are: <code>Mon</code>, <code>Tue</code>, <code>Wed</code>, <code>Thr</code>, <code>Fri</code>, <code>Sat</code>, or <code>Sun</code>.</p>
}

// AssociateNodeRequest represents the AssociateNodeRequest schema from the OpenAPI specification
type AssociateNodeRequest struct {
	Engineattributes interface{} `json:"EngineAttributes"`
	Nodename interface{} `json:"NodeName"`
	Servername interface{} `json:"ServerName"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Tagkeys interface{} `json:"TagKeys"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// DescribeNodeAssociationStatusRequest represents the DescribeNodeAssociationStatusRequest schema from the OpenAPI specification
type DescribeNodeAssociationStatusRequest struct {
	Nodeassociationstatustoken interface{} `json:"NodeAssociationStatusToken"`
	Servername interface{} `json:"ServerName"`
}

// DescribeEventsResponse represents the DescribeEventsResponse schema from the OpenAPI specification
type DescribeEventsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Serverevents interface{} `json:"ServerEvents,omitempty"`
}

// RestoreServerRequest represents the RestoreServerRequest schema from the OpenAPI specification
type RestoreServerRequest struct {
	Keypair interface{} `json:"KeyPair,omitempty"`
	Servername interface{} `json:"ServerName"`
	Backupid interface{} `json:"BackupId"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
}

// ExportServerEngineAttributeRequest represents the ExportServerEngineAttributeRequest schema from the OpenAPI specification
type ExportServerEngineAttributeRequest struct {
	Exportattributename interface{} `json:"ExportAttributeName"`
	Inputattributes interface{} `json:"InputAttributes,omitempty"`
	Servername interface{} `json:"ServerName"`
}

// ExportServerEngineAttributeResponse represents the ExportServerEngineAttributeResponse schema from the OpenAPI specification
type ExportServerEngineAttributeResponse struct {
	Engineattribute interface{} `json:"EngineAttribute,omitempty"`
	Servername interface{} `json:"ServerName,omitempty"`
}

// CreateBackupRequest represents the CreateBackupRequest schema from the OpenAPI specification
type CreateBackupRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Servername interface{} `json:"ServerName"`
}

// DeleteServerResponse represents the DeleteServerResponse schema from the OpenAPI specification
type DeleteServerResponse struct {
}

// DescribeServersRequest represents the DescribeServersRequest schema from the OpenAPI specification
type DescribeServersRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Servername interface{} `json:"ServerName,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// Backup represents the Backup schema from the OpenAPI specification
type Backup struct {
	Backuptype interface{} `json:"BackupType,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
	Preferredbackupwindow interface{} `json:"PreferredBackupWindow,omitempty"`
	S3dataurl interface{} `json:"S3DataUrl,omitempty"`
	Userarn interface{} `json:"UserArn,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Backuparn interface{} `json:"BackupArn,omitempty"`
	Engine interface{} `json:"Engine,omitempty"`
	S3logurl interface{} `json:"S3LogUrl,omitempty"`
	Keypair interface{} `json:"KeyPair,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Preferredmaintenancewindow interface{} `json:"PreferredMaintenanceWindow,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Servername interface{} `json:"ServerName,omitempty"`
	Enginemodel interface{} `json:"EngineModel,omitempty"`
	Instanceprofilearn interface{} `json:"InstanceProfileArn,omitempty"`
	S3datasize interface{} `json:"S3DataSize,omitempty"`
	Statusdescription interface{} `json:"StatusDescription,omitempty"`
	Backupid interface{} `json:"BackupId,omitempty"`
	Toolsversion interface{} `json:"ToolsVersion,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// AccountAttribute represents the AccountAttribute schema from the OpenAPI specification
type AccountAttribute struct {
	Used interface{} `json:"Used,omitempty"`
	Maximum interface{} `json:"Maximum,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// DeleteBackupResponse represents the DeleteBackupResponse schema from the OpenAPI specification
type DeleteBackupResponse struct {
}

// DescribeBackupsResponse represents the DescribeBackupsResponse schema from the OpenAPI specification
type DescribeBackupsResponse struct {
	Backups interface{} `json:"Backups,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StartMaintenanceRequest represents the StartMaintenanceRequest schema from the OpenAPI specification
type StartMaintenanceRequest struct {
	Engineattributes interface{} `json:"EngineAttributes,omitempty"`
	Servername interface{} `json:"ServerName"`
}

// DeleteServerRequest represents the DeleteServerRequest schema from the OpenAPI specification
type DeleteServerRequest struct {
	Servername interface{} `json:"ServerName"`
}

// CreateBackupResponse represents the CreateBackupResponse schema from the OpenAPI specification
type CreateBackupResponse struct {
	Backup interface{} `json:"Backup,omitempty"`
}

// DescribeEventsRequest represents the DescribeEventsRequest schema from the OpenAPI specification
type DescribeEventsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Servername interface{} `json:"ServerName"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DescribeServersResponse represents the DescribeServersResponse schema from the OpenAPI specification
type DescribeServersResponse struct {
	Servers interface{} `json:"Servers,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateServerRequest represents the CreateServerRequest schema from the OpenAPI specification
type CreateServerRequest struct {
	Backupretentioncount interface{} `json:"BackupRetentionCount,omitempty"`
	Customdomain interface{} `json:"CustomDomain,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Instanceprofilearn interface{} `json:"InstanceProfileArn"`
	Engine interface{} `json:"Engine"`
	Enginemodel interface{} `json:"EngineModel,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn"`
	Tags interface{} `json:"Tags,omitempty"`
	Servername interface{} `json:"ServerName"`
	Keypair interface{} `json:"KeyPair,omitempty"`
	Instancetype interface{} `json:"InstanceType"`
	Customprivatekey interface{} `json:"CustomPrivateKey,omitempty"`
	Disableautomatedbackup interface{} `json:"DisableAutomatedBackup,omitempty"`
	Preferredbackupwindow interface{} `json:"PreferredBackupWindow,omitempty"`
	Customcertificate interface{} `json:"CustomCertificate,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Backupid interface{} `json:"BackupId,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Engineattributes interface{} `json:"EngineAttributes,omitempty"`
	Preferredmaintenancewindow interface{} `json:"PreferredMaintenanceWindow,omitempty"`
	Associatepublicipaddress interface{} `json:"AssociatePublicIpAddress,omitempty"`
}

// DescribeNodeAssociationStatusResponse represents the DescribeNodeAssociationStatusResponse schema from the OpenAPI specification
type DescribeNodeAssociationStatusResponse struct {
	Engineattributes interface{} `json:"EngineAttributes,omitempty"`
	Nodeassociationstatus interface{} `json:"NodeAssociationStatus,omitempty"`
}

// StartMaintenanceResponse represents the StartMaintenanceResponse schema from the OpenAPI specification
type StartMaintenanceResponse struct {
	Server interface{} `json:"Server,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DisassociateNodeResponse represents the DisassociateNodeResponse schema from the OpenAPI specification
type DisassociateNodeResponse struct {
	Nodeassociationstatustoken interface{} `json:"NodeAssociationStatusToken,omitempty"`
}

// CreateServerResponse represents the CreateServerResponse schema from the OpenAPI specification
type CreateServerResponse struct {
	Server interface{} `json:"Server,omitempty"`
}

// Server represents the Server schema from the OpenAPI specification
type Server struct {
	Disableautomatedbackup interface{} `json:"DisableAutomatedBackup,omitempty"`
	Associatepublicipaddress interface{} `json:"AssociatePublicIpAddress,omitempty"`
	Engineversion interface{} `json:"EngineVersion,omitempty"`
	Instanceprofilearn interface{} `json:"InstanceProfileArn,omitempty"`
	Preferredmaintenancewindow interface{} `json:"PreferredMaintenanceWindow,omitempty"`
	Servername interface{} `json:"ServerName,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Subnetids interface{} `json:"SubnetIds,omitempty"`
	Enginemodel interface{} `json:"EngineModel,omitempty"`
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Preferredbackupwindow interface{} `json:"PreferredBackupWindow,omitempty"`
	Customdomain interface{} `json:"CustomDomain,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Backupretentioncount interface{} `json:"BackupRetentionCount,omitempty"`
	Maintenancestatus interface{} `json:"MaintenanceStatus,omitempty"`
	Endpoint interface{} `json:"Endpoint,omitempty"`
	Engine interface{} `json:"Engine,omitempty"`
	Serverarn interface{} `json:"ServerArn,omitempty"`
	Statusreason interface{} `json:"StatusReason,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Keypair interface{} `json:"KeyPair,omitempty"`
	Cloudformationstackarn interface{} `json:"CloudFormationStackArn,omitempty"`
	Engineattributes interface{} `json:"EngineAttributes,omitempty"`
	Instancetype interface{} `json:"InstanceType,omitempty"`
}

// UpdateServerEngineAttributesResponse represents the UpdateServerEngineAttributesResponse schema from the OpenAPI specification
type UpdateServerEngineAttributesResponse struct {
	Server interface{} `json:"Server,omitempty"`
}

// DescribeAccountAttributesResponse represents the DescribeAccountAttributesResponse schema from the OpenAPI specification
type DescribeAccountAttributesResponse struct {
	Attributes interface{} `json:"Attributes,omitempty"`
}

// DisassociateNodeRequest represents the DisassociateNodeRequest schema from the OpenAPI specification
type DisassociateNodeRequest struct {
	Engineattributes interface{} `json:"EngineAttributes,omitempty"`
	Nodename interface{} `json:"NodeName"`
	Servername interface{} `json:"ServerName"`
}

// UpdateServerResponse represents the UpdateServerResponse schema from the OpenAPI specification
type UpdateServerResponse struct {
	Server interface{} `json:"Server,omitempty"`
}

// ServerEvent represents the ServerEvent schema from the OpenAPI specification
type ServerEvent struct {
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Logurl interface{} `json:"LogUrl,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Servername interface{} `json:"ServerName,omitempty"`
}

// UpdateServerEngineAttributesRequest represents the UpdateServerEngineAttributesRequest schema from the OpenAPI specification
type UpdateServerEngineAttributesRequest struct {
	Attributename interface{} `json:"AttributeName"`
	Attributevalue interface{} `json:"AttributeValue,omitempty"`
	Servername interface{} `json:"ServerName"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// DeleteBackupRequest represents the DeleteBackupRequest schema from the OpenAPI specification
type DeleteBackupRequest struct {
	Backupid interface{} `json:"BackupId"`
}

// RestoreServerResponse represents the RestoreServerResponse schema from the OpenAPI specification
type RestoreServerResponse struct {
	Server Server `json:"Server,omitempty"` // Describes a configuration management server.
}

// AssociateNodeResponse represents the AssociateNodeResponse schema from the OpenAPI specification
type AssociateNodeResponse struct {
	Nodeassociationstatustoken interface{} `json:"NodeAssociationStatusToken,omitempty"`
}

// DescribeBackupsRequest represents the DescribeBackupsRequest schema from the OpenAPI specification
type DescribeBackupsRequest struct {
	Backupid interface{} `json:"BackupId,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Servername interface{} `json:"ServerName,omitempty"`
}

// DescribeAccountAttributesRequest represents the DescribeAccountAttributesRequest schema from the OpenAPI specification
type DescribeAccountAttributesRequest struct {
}

// EngineAttribute represents the EngineAttribute schema from the OpenAPI specification
type EngineAttribute struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}
