/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/types.json
*/
package bundle

import "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/relation"

type TypeKey string

func (tk TypeKey) String() string {
	return string(tk)
}
func (tk TypeKey) URL() string {
	return string(TypePrefix + tk)
}

const (
	TypePrefix = "_ot"
)
const (
	TypeKeyNote       TypeKey = "note"
	TypeKeyContact    TypeKey = "contact"
	TypeKeyIdea       TypeKey = "idea"
	TypeKeyTask       TypeKey = "task"
	TypeKeyRelation   TypeKey = "relation"
	TypeKeyVideo      TypeKey = "video"
	TypeKeyDashboard  TypeKey = "dashboard"
	TypeKeyObjectType TypeKey = "objectType"
	TypeKeyTemplate   TypeKey = "template"
	TypeKeySet        TypeKey = "set"
	TypeKeyPage       TypeKey = "page"
	TypeKeyImage      TypeKey = "image"
	TypeKeyProfile    TypeKey = "profile"
	TypeKeyAudio      TypeKey = "audio"
	TypeKeyDocument   TypeKey = "document"
	TypeKeyFile       TypeKey = "file"
	TypeKeyProject    TypeKey = "project"
)

var (
	types = map[TypeKey]*relation.ObjectType{
		TypeKeyAudio: {

			Description: "",
			Hidden:      true,
			IconEmoji:   "🎵",
			Layout:      relation.ObjectType_basic,
			Name:        "Audio",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyArtist], relations[RelationKeyAudioAlbum], relations[RelationKeyAudioAlbumTrackNumber], relations[RelationKeyAudioGenre], relations[RelationKeyReleasedYear], relations[RelationKeyThumbnailImage], relations[RelationKeyComposer], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType], relations[RelationKeyAddedDate], relations[RelationKeyFileExt]},
			Url:         TypePrefix + "audio",
		},
		TypeKeyContact: {

			Description: "",
			IconEmoji:   "📇",
			Layout:      relation.ObjectType_profile,
			Name:        "Contact",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived]},
			Url:         TypePrefix + "contact",
		},
		TypeKeyDashboard: {

			Description: "Internal home dashboard",
			Hidden:      true,
			Layout:      relation.ObjectType_dashboard,
			Name:        "Dashboard",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden]},
			Url:         TypePrefix + "dashboard",
		},
		TypeKeyDocument: {

			Description: "",
			IconEmoji:   "📋",
			Layout:      relation.ObjectType_basic,
			Name:        "Document",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived]},
			Url:         TypePrefix + "document",
		},
		TypeKeyFile: {

			Description: "",
			IconEmoji:   "🗂️",
			Layout:      relation.ObjectType_basic,
			Name:        "File",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyFileMimeType], relations[RelationKeySizeInBytes], relations[RelationKeyAddedDate], relations[RelationKeyFileExt]},
			Url:         TypePrefix + "file",
		},
		TypeKeyIdea: {

			Description: "",
			IconEmoji:   "💡",
			Layout:      relation.ObjectType_basic,
			Name:        "Idea",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Url:         TypePrefix + "idea",
		},
		TypeKeyImage: {

			Description: "",
			IconEmoji:   "🌅",
			Layout:      relation.ObjectType_image,
			Name:        "Image",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyFileMimeType], relations[RelationKeyWidthInPixels], relations[RelationKeyCamera], relations[RelationKeyHeightInPixels], relations[RelationKeySizeInBytes], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure], relations[RelationKeyAddedDate], relations[RelationKeyFocalRatio], relations[RelationKeyFileExt]},
			Url:         TypePrefix + "image",
		},
		TypeKeyNote: {

			Description: "",
			IconEmoji:   "🗒️",
			Layout:      relation.ObjectType_basic,
			Name:        "Note",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived]},
			Url:         TypePrefix + "note",
		},
		TypeKeyObjectType: {

			Description: "Object that contains a definition of some object type",
			Hidden:      true,
			IconEmoji:   "🔮",
			Layout:      relation.ObjectType_objectType,
			Name:        "Type",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyRecommendedRelations], relations[RelationKeyRecommendedLayout], relations[RelationKeyMpAddedToLibrary], relations[RelationKeyIsHidden]},
			Url:         TypePrefix + "objectType",
		},
		TypeKeyPage: {

			Description: "Base type to start with",
			IconEmoji:   "📄",
			Layout:      relation.ObjectType_basic,
			Name:        "Draft",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived]},
			Url:         TypePrefix + "page",
		},
		TypeKeyProfile: {

			Description: "",
			IconEmoji:   "🧍",
			Layout:      relation.ObjectType_profile,
			Name:        "Human",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived]},
			Url:         TypePrefix + "profile",
		},
		TypeKeyProject: {

			Description: "",
			IconEmoji:   "🔨",
			Layout:      relation.ObjectType_basic,
			Name:        "Project",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived]},
			Url:         TypePrefix + "project",
		},
		TypeKeyRelation: {

			Description: "",
			Hidden:      true,
			IconEmoji:   "🔗",
			Layout:      relation.ObjectType_relation,
			Name:        "Relation",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyLayout], relations[RelationKeyDescription], relations[RelationKeyCreator], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyMpAddedToLibrary], relations[RelationKeyRelationFormat], relations[RelationKeyIsHidden]},
			Url:         TypePrefix + "relation",
		},
		TypeKeySet: {

			Description: "",
			Hidden:      true,
			IconEmoji:   "🗂️",
			Layout:      relation.ObjectType_set,
			Name:        "Set",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeySetOf]},
			Url:         TypePrefix + "set",
		},
		TypeKeyTask: {

			Description: "",
			IconEmoji:   "✔️",
			Layout:      relation.ObjectType_todo,
			Name:        "Task",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyAssignee], relations[RelationKeyDueDate], relations[RelationKeyAttachments], relations[RelationKeyStatus], relations[RelationKeyDone], relations[RelationKeyPriority], relations[RelationKeyLinkedTasks], relations[RelationKeyLinkedProjects], relations[RelationKeyTag]},
			Url:         TypePrefix + "task",
		},
		TypeKeyTemplate: {

			Description: "Special type to create objects from",
			Hidden:      true,
			IconEmoji:   "✨",
			Layout:      relation.ObjectType_basic,
			Name:        "Template",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTemplateType]},
			Url:         TypePrefix + "template",
		},
		TypeKeyVideo: {

			Description: "",
			IconEmoji:   "📹",
			Layout:      relation.ObjectType_basic,
			Name:        "Video",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType], relations[RelationKeyCamera], relations[RelationKeyThumbnailImage], relations[RelationKeyHeightInPixels], relations[RelationKeyWidthInPixels], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure], relations[RelationKeyAddedDate], relations[RelationKeyFileExt]},
			Url:         TypePrefix + "video",
		},
	}
)
