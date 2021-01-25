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
	TypePrefix = "https://anytype.io/schemas/object/bundled/"
)
const (
	TypeKeyNote       TypeKey = "note"
	TypeKeyDashboard  TypeKey = "dashboard"
	TypeKeyContact    TypeKey = "contact"
	TypeKeyIdea       TypeKey = "idea"
	TypeKeyTask       TypeKey = "task"
	TypeKeyRelation   TypeKey = "relation"
	TypeKeyVideo      TypeKey = "video"
	TypeKeyObjectType TypeKey = "objectType"
	TypeKeySet        TypeKey = "set"
	TypeKeyPage       TypeKey = "page"
	TypeKeyImage      TypeKey = "image"
	TypeKeyProfile    TypeKey = "profile"
	TypeKeyAudio      TypeKey = "audio"
	TypeKeyDocument   TypeKey = "document"
	TypeKeyFile       TypeKey = "file"
	TypeKeyProject    TypeKey = "project"
	TypeKeyCollection TypeKey = "collection"
)

var (
	types = map[TypeKey]*relation.ObjectType{
		TypeKeyAudio: {

			Description: "",
			Layout:      relation.ObjectType_basic,
			Name:        "Audio",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription], relations[RelationKeyArtist], relations[RelationKeyAudioAlbum], relations[RelationKeyAudioAlbumTrackNumber], relations[RelationKeyAudioGenre], relations[RelationKeyReleasedYear], relations[RelationKeyThumbnailImage], relations[RelationKeyComposer], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType]},
			Url:         TypePrefix + "audio",
		},
		TypeKeyCollection: {

			Description: "",
			Layout:      relation.ObjectType_database,
			Name:        "Collection",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription], relations[RelationKeyCollectionOf]},
			Url:         TypePrefix + "collection",
		},
		TypeKeyContact: {

			Description: "",
			Layout:      relation.ObjectType_profile,
			Name:        "Contact",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription]},
			Url:         TypePrefix + "contact",
		},
		TypeKeyDashboard: {

			Description: "",
			Layout:      relation.ObjectType_dashboard,
			Name:        "Dashboard",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription]},
			Url:         TypePrefix + "dashboard",
		},
		TypeKeyDocument: {

			Description: "",
			Layout:      relation.ObjectType_basic,
			Name:        "Document",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription]},
			Url:         TypePrefix + "document",
		},
		TypeKeyFile: {

			Description: "",
			Layout:      relation.ObjectType_basic,
			Name:        "File",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription], relations[RelationKeyFileMimeType], relations[RelationKeySizeInBytes]},
			Url:         TypePrefix + "file",
		},
		TypeKeyIdea: {

			Description: "",
			Layout:      relation.ObjectType_basic,
			Name:        "Idea",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription], relations[RelationKeyTag]},
			Url:         TypePrefix + "idea",
		},
		TypeKeyImage: {

			Description: "",
			Layout:      relation.ObjectType_basic,
			Name:        "Image",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription], relations[RelationKeyFileMimeType], relations[RelationKeyWidthInPixels], relations[RelationKeyCamera], relations[RelationKeyHeightInPixels], relations[RelationKeySizeInBytes], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure]},
			Url:         TypePrefix + "image",
		},
		TypeKeyNote: {

			Description: "",
			Layout:      relation.ObjectType_basic,
			Name:        "Note",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription]},
			Url:         TypePrefix + "note",
		},
		TypeKeyObjectType: {

			Description: "Object that contains a definition of some object type",
			Layout:      relation.ObjectType_objectType,
			Name:        "Type",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription], relations[RelationKeyRecommendedRelations], relations[RelationKeyRecommendedLayout]},
			Url:         TypePrefix + "objectType",
		},
		TypeKeyPage: {

			Description: "Base type to start with",
			Layout:      relation.ObjectType_basic,
			Name:        "Undefined",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations]},
			Url:         TypePrefix + "page",
		},
		TypeKeyProfile: {

			Description: "",
			Layout:      relation.ObjectType_profile,
			Name:        "Human",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription]},
			Url:         TypePrefix + "profile",
		},
		TypeKeyProject: {

			Description: "",
			Layout:      relation.ObjectType_basic,
			Name:        "Project",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription]},
			Url:         TypePrefix + "project",
		},
		TypeKeyRelation: {

			Description: "",
			Layout:      relation.ObjectType_relation,
			Name:        "Relation",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyLayout], relations[RelationKeyDescription]},
			Url:         TypePrefix + "relation",
		},
		TypeKeySet: {

			Description: "",
			Layout:      relation.ObjectType_set,
			Name:        "Set of objects",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription], relations[RelationKeySetOf]},
			Url:         TypePrefix + "set",
		},
		TypeKeyTask: {

			Description: "",
			Layout:      relation.ObjectType_action,
			Name:        "Task",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription], relations[RelationKeyAssignee], relations[RelationKeyDueDate], relations[RelationKeyDescription], relations[RelationKeyAttachments], relations[RelationKeyStatus], relations[RelationKeyDone], relations[RelationKeyPriority], relations[RelationKeyLinkedTasks], relations[RelationKeyLinkedProjects], relations[RelationKeyTag]},
			Url:         TypePrefix + "task",
		},
		TypeKeyVideo: {

			Description: "",
			Layout:      relation.ObjectType_basic,
			Name:        "Video",
			Relations:   []*relation.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyDescription], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType], relations[RelationKeyCamera], relations[RelationKeyThumbnailImage], relations[RelationKeyHeightInPixels], relations[RelationKeyWidthInPixels], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure]},
			Url:         TypePrefix + "video",
		},
	}
)
