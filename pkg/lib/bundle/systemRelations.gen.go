/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/systemRelations.json
*/
package bundle

import domain "github.com/anyproto/anytype-heart/core/domain"

const SystemRelationsChecksum = "4e6e9e08d02922886ef0146cf5ed44560b96cac0bcdfdeaf5393105473f15f8d"

// SystemRelations contains relations that have some special biz logic depends on them in some objects
// in case EVERY object depend on the relation please add it to RequiredInternalRelations
var SystemRelations = append(RequiredInternalRelations, []domain.RelationKey{
	RelationKeyAddedDate,
	RelationKeySource,
	RelationKeySourceObject,
	RelationKeySetOf,
	RelationKeyRelationFormat,
	RelationKeyRelationKey,
	RelationKeyRelationReadonlyValue,
	RelationKeyRelationDefaultValue,
	RelationKeyRelationMaxCount,
	RelationKeyRelationOptionColor,
	RelationKeyRelationFormatObjectTypes,
	RelationKeyIsReadonly,
	RelationKeyIsDeleted,
	RelationKeyIsHidden,
	RelationKeyIsHiddenDiscovery,
	RelationKeyDone,
	RelationKeyIsArchived,
	RelationKeyTemplateIsBundled,
	RelationKeyTag,
	RelationKeySmartblockTypes,
	RelationKeyTargetObjectType,
	RelationKeyRecommendedLayout,
	RelationKeyFileExt,
	RelationKeyFileMimeType,
	RelationKeySizeInBytes,
	RelationKeyOldAnytypeID,
	RelationKeySpaceDashboardId,
	RelationKeyRecommendedRelations,
	RelationKeyIconOption,
	RelationKeyWidthInPixels,
	RelationKeyHeightInPixels,
	RelationKeyFileExt,
	RelationKeySizeInBytes,
	RelationKeySourceFilePath,
	RelationKeyFileSyncStatus,
	RelationKeyDefaultTemplateId,
	RelationKeyUniqueKey,
	RelationKeyBacklinks,
	RelationKeyProfileOwnerIdentity,
	RelationKeyFileBackupStatus,
	RelationKeyFileId,
	RelationKeyFileIndexingStatus,
	RelationKeyOrigin,
	RelationKeyRevision,
	RelationKeyImageKind,
	RelationKeyImportType,
	RelationKeySpaceAccessType,
	RelationKeySpaceInviteFileCid,
	RelationKeySpaceInviteFileKey,
	RelationKeyParticipantPermissions,
	RelationKeyParticipantStatus,
	RelationKeyLatestAclHeadId,
	RelationKeyIdentity,
}...)
