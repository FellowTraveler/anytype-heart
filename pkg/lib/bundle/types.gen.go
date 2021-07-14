/*
Code generated by pkg/lib/bundle/generator. DO NOT EDIT.
source: pkg/lib/bundle/types.json
*/
package bundle

import "github.com/anytypeio/go-anytype-middleware/pkg/lib/pb/model"

const TypeChecksum = "2f79013bf6bbd30748e3840b548edf8f1abc9ed98f2e7d436467f987bff0127f"

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
	TypeKeyDailyReflection TypeKey = "dailyReflection"
	TypeKeyRecipe          TypeKey = "recipe"
	TypeKeyNote            TypeKey = "note"
	TypeKeyContact         TypeKey = "contact"
	TypeKeyIdea            TypeKey = "idea"
	TypeKeyTask            TypeKey = "task"
	TypeKeyRelation        TypeKey = "relation"
	TypeKeyVideo           TypeKey = "video"
	TypeKeyDashboard       TypeKey = "dashboard"
	TypeKeyDailyPlan       TypeKey = "dailyPlan"
	TypeKeyMeetingNote     TypeKey = "meetingNote"
	TypeKeyArticle         TypeKey = "article"
	TypeKeyObjectType      TypeKey = "objectType"
	TypeKeyTemplate        TypeKey = "template"
	TypeKeySet             TypeKey = "set"
	TypeKeyDiaryEntry      TypeKey = "diaryEntry"
	TypeKeyPage            TypeKey = "page"
	TypeKeyImage           TypeKey = "image"
	TypeKeyProfile         TypeKey = "profile"
	TypeKeyAudio           TypeKey = "audio"
	TypeKeyActionPlan      TypeKey = "actionPlan"
	TypeKeyDocument        TypeKey = "document"
	TypeKeyFile            TypeKey = "file"
	TypeKeyProject         TypeKey = "project"
)

var (
	types = map[TypeKey]*model.ObjectType{
		TypeKeyActionPlan: {

			Description: "Is a detailed plan outlining actions needed to reach one or more goals",
			IconEmoji:   "🤝",
			Layout:      model.ObjectType_basic,
			Name:        "Action Plan",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyTasks], relations[RelationKeyResources], relations[RelationKeyResult], relations[RelationKeyDueDate], relations[RelationKeyResponsible]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "actionPlan",
		},
		TypeKeyArticle: {

			Description: "A piece of writing included with others in a newspaper, magazine, or other publication.",
			IconEmoji:   "📄",
			Layout:      model.ObjectType_basic,
			Name:        "Article",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "article",
		},
		TypeKeyAudio: {

			Description: "Sound, especially when recorded, transmitted, or reproduced.",
			Hidden:      true,
			IconEmoji:   "🎵",
			Layout:      model.ObjectType_basic,
			Name:        "Audio",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyArtist], relations[RelationKeyAudioAlbum], relations[RelationKeyAudioAlbumTrackNumber], relations[RelationKeyAudioGenre], relations[RelationKeyReleasedYear], relations[RelationKeyThumbnailImage], relations[RelationKeyComposer], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType], relations[RelationKeyAddedDate], relations[RelationKeyFileExt]},
			Types:       []model.SmartBlockType{model.SmartBlockType_File},
			Url:         TypePrefix + "audio",
		},
		TypeKeyContact: {

			Description: "Information to make action of communicating or meeting with Human or Company",
			IconEmoji:   "📇",
			Layout:      model.ObjectType_profile,
			Name:        "Contact",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyPhone], relations[RelationKeyEmail], relations[RelationKeyDateOfBirth], relations[RelationKeyPlaceOfBirth], relations[RelationKeyCompany], relations[RelationKeySocialProfile], relations[RelationKeyJob], relations[RelationKeyLinkedContacts], relations[RelationKeyOccupation], relations[RelationKeyInstagram], relations[RelationKeyGender], relations[RelationKeyFacebook]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "contact",
		},
		TypeKeyDailyPlan: {

			Description: "A detailed proposal for doing or achieving something",
			IconEmoji:   "📋",
			Layout:      model.ObjectType_basic,
			Name:        "Daily Plan",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyNotes], relations[RelationKeyTasks], relations[RelationKeyEvents]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "dailyPlan",
		},
		TypeKeyDailyReflection: {

			Description: "Serious thought or consideration",
			IconEmoji:   "💭",
			Layout:      model.ObjectType_basic,
			Name:        "Daily Reflection",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyIntentions], relations[RelationKeyHappenings], relations[RelationKeyGratefulFor], relations[RelationKeyActionItems], relations[RelationKeyMood], relations[RelationKeyWorriedAbout]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "dailyReflection",
		},
		TypeKeyDashboard: {

			Description: "Internal home dashboard with favourite objects",
			Hidden:      true,
			Layout:      model.ObjectType_dashboard,
			Name:        "Dashboard",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Home},
			Url:         TypePrefix + "dashboard",
		},
		TypeKeyDiaryEntry: {

			Description: "Record of events and experiences",
			IconEmoji:   "✨",
			Layout:      model.ObjectType_basic,
			Name:        "Diary Entry",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyMood]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "diaryEntry",
		},
		TypeKeyDocument: {

			Description: "A piece of matter that provides information or evidence or that serves as an official record.",
			IconEmoji:   "📃",
			Layout:      model.ObjectType_basic,
			Name:        "Document",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "document",
		},
		TypeKeyFile: {

			Description: "Computer resource for recording data in a computer storage device",
			IconEmoji:   "🗂️",
			Layout:      model.ObjectType_basic,
			Name:        "File",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyFileMimeType], relations[RelationKeySizeInBytes], relations[RelationKeyAddedDate], relations[RelationKeyFileExt]},
			Types:       []model.SmartBlockType{model.SmartBlockType_File},
			Url:         TypePrefix + "file",
		},
		TypeKeyIdea: {

			Description: "A thought or suggestion as to a possible course of action",
			IconEmoji:   "💡",
			Layout:      model.ObjectType_basic,
			Name:        "Idea",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyProblem], relations[RelationKeySolution], relations[RelationKeyAlternative]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "idea",
		},
		TypeKeyImage: {

			Description: "A representation of the external form of a person or thing in art.",
			IconEmoji:   "🌅",
			Layout:      model.ObjectType_image,
			Name:        "Image",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyFileMimeType], relations[RelationKeyWidthInPixels], relations[RelationKeyCamera], relations[RelationKeyHeightInPixels], relations[RelationKeySizeInBytes], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure], relations[RelationKeyAddedDate], relations[RelationKeyFocalRatio], relations[RelationKeyFileExt]},
			Types:       []model.SmartBlockType{model.SmartBlockType_File},
			Url:         TypePrefix + "image",
		},
		TypeKeyMeetingNote: {

			Description: "Quick references to ideas, goals, deadlines, data, and anything else important that's covered in your meeting.",
			IconEmoji:   "✏️",
			Layout:      model.ObjectType_basic,
			Name:        "Meeting Note",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyParticipants], relations[RelationKeyActionItems], relations[RelationKeyAgenda]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "meetingNote",
		},
		TypeKeyNote: {

			Description: "A brief record of points written down as an aid to memory.",
			IconEmoji:   "🗒️",
			Layout:      model.ObjectType_basic,
			Name:        "Note",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "note",
		},
		TypeKeyObjectType: {

			Description: "Object that contains a definition of some object type",
			Hidden:      true,
			IconEmoji:   "🔮",
			Layout:      model.ObjectType_objectType,
			Name:        "Type",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyRecommendedRelations], relations[RelationKeyRecommendedLayout], relations[RelationKeyMpAddedToLibrary]},
			Types:       []model.SmartBlockType{model.SmartBlockType_STObjectType, model.SmartBlockType_BundledObjectType},
			Url:         TypePrefix + "objectType",
		},
		TypeKeyPage: {

			Description: "Proto type to start with",
			IconEmoji:   "⚪",
			Layout:      model.ObjectType_basic,
			Name:        "Draft",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "page",
		},
		TypeKeyProfile: {

			Description: "Description of individuals' characteristics that identify you",
			IconEmoji:   "🧍",
			Layout:      model.ObjectType_profile,
			Name:        "Human",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page, model.SmartBlockType_ProfilePage},
			Url:         TypePrefix + "profile",
		},
		TypeKeyProject: {

			Description: "An individual or collaborative enterprise that is carefully planned to achieve a particular aim",
			IconEmoji:   "🔨",
			Layout:      model.ObjectType_basic,
			Name:        "Project",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "project",
		},
		TypeKeyRecipe: {

			Description: "A recipe is a set of instructions that describes how to prepare or make something, especially a dish of prepared food",
			IconEmoji:   "🍲",
			Layout:      model.ObjectType_basic,
			Name:        "Recipe",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyTime], relations[RelationKeyServings], relations[RelationKeyIngredients], relations[RelationKeyInstructions], relations[RelationKeyDifficulty]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "recipe",
		},
		TypeKeyRelation: {

			Description: "Meaningful connection between objects",
			Hidden:      true,
			IconEmoji:   "🔗",
			Layout:      model.ObjectType_relation,
			Name:        "Relation",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyLayout], relations[RelationKeyDescription], relations[RelationKeyCreator], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyMpAddedToLibrary], relations[RelationKeyRelationFormat], relations[RelationKeyRelationMaxCount], relations[RelationKeyRelationDict], relations[RelationKeyRelationDefaultValue], relations[RelationKeyRelationFormatObjectTypes]},
			Types:       []model.SmartBlockType{model.SmartBlockType_IndexedRelation, model.SmartBlockType_BundledRelation},
			Url:         TypePrefix + "relation",
		},
		TypeKeySet: {

			Description: "Collection of objects sharing one characteristic",
			IconEmoji:   "🗂️",
			Layout:      model.ObjectType_set,
			Name:        "Set",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeySetOf]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Set},
			Url:         TypePrefix + "set",
		},
		TypeKeyTask: {

			Description: "A piece of work to be done or undertaken",
			IconEmoji:   "✔️",
			Layout:      model.ObjectType_todo,
			Name:        "Task",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTag], relations[RelationKeyAssignee], relations[RelationKeyDueDate], relations[RelationKeyAttachments], relations[RelationKeyStatus], relations[RelationKeyDone], relations[RelationKeyPriority], relations[RelationKeyTasks], relations[RelationKeyLinkedProjects]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Page},
			Url:         TypePrefix + "task",
		},
		TypeKeyTemplate: {

			Description: "Sample object that has already some details in place and used to create objects from",
			Hidden:      true,
			Layout:      model.ObjectType_basic,
			Name:        "Template",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyTargetObjectType], relations[RelationKeyTemplateIsBundled]},
			Types:       []model.SmartBlockType{model.SmartBlockType_Template},
			Url:         TypePrefix + "template",
		},
		TypeKeyVideo: {

			Description: "The recording, reproducing, or broadcasting of moving visual images.",
			IconEmoji:   "📹",
			Layout:      model.ObjectType_basic,
			Name:        "Video",
			Readonly:    true,
			Relations:   []*model.Relation{relations[RelationKeyId], relations[RelationKeyName], relations[RelationKeyDescription], relations[RelationKeyType], relations[RelationKeyCreator], relations[RelationKeyCreatedDate], relations[RelationKeyLayout], relations[RelationKeyLastModifiedBy], relations[RelationKeyIconImage], relations[RelationKeyIconEmoji], relations[RelationKeyCoverId], relations[RelationKeyLastModifiedDate], relations[RelationKeyLastOpenedDate], relations[RelationKeyCoverX], relations[RelationKeyCoverY], relations[RelationKeyCoverScale], relations[RelationKeyToBeDeletedDate], relations[RelationKeyFeaturedRelations], relations[RelationKeyCoverType], relations[RelationKeyLayoutAlign], relations[RelationKeyIsHidden], relations[RelationKeyIsArchived], relations[RelationKeyDurationInSeconds], relations[RelationKeySizeInBytes], relations[RelationKeyFileMimeType], relations[RelationKeyCamera], relations[RelationKeyThumbnailImage], relations[RelationKeyHeightInPixels], relations[RelationKeyWidthInPixels], relations[RelationKeyCameraIso], relations[RelationKeyAperture], relations[RelationKeyExposure], relations[RelationKeyAddedDate], relations[RelationKeyFileExt]},
			Types:       []model.SmartBlockType{model.SmartBlockType_File},
			Url:         TypePrefix + "video",
		},
	}
)
