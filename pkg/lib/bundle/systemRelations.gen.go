/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/systemRelations.json
*/
package bundle

const SystemRelationsChecksum = "bfb1c15c23c61f26cec64e4abbc4f8f95f6e53fc2fbe40e0ac2077bd41302733"

// SystemRelations contains relations that have some special biz logic depends on them in some objects
// in case EVERY object depend on the relation please add it to RequiredInternalRelations
var SystemRelations = append(RequiredInternalRelations, []RelationKey{
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
	RelationKeySpaceAccessibility,
	RelationKeyWidthInPixels,
	RelationKeyHeightInPixels,
	RelationKeyFileExt,
	RelationKeySizeInBytes,
	RelationKeyOriginalCreatedDate,
	RelationKeySourceFilePath,
}...)
