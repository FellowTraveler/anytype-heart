package schema

// Code generated by go generate; DO NOT EDIT.
//go:generate go run embed/embed.go

var SchemaByURL = map[string]string{

	"https://anytype.io/schemas/page": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "id": "https://anytype.io/schemas/page",
  "title": "Anytype Page",
  "description": "This schema contains the base properties of Anytype Page and should be refereed if you want to extend it",
  "type": "array",
  "items": {
    "$ref": "https://anytype.io/schemas/relation"
  },
  "uniqueItems": true,
  "default": [
    {
      "id": "name",
      "name": "Name",
      "type": "https://anytype.io/schemas/types/title"
    },
    {
      "id": "iconEmoji",
      "name": "Emoji",
      "type": "https://anytype.io/schemas/types/emoji"
    },
    {
      "id": "iconImage",
      "name": "Image",
      "type": "https://anytype.io/schemas/types/image"
    },
    {
      "id": "isArchived",
      "name": "Archived",
      "type": "https://anytype.io/schemas/types/checkbox"
    }
  ]
}
`,
	"https://anytype.io/schemas/person": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "id": "https://anytype.io/schemas/person",
  "title": "Anytype Page",
  "description": "This schema contains the base properties of Anytype Person",

  "allOf": [
    { "$ref": "https://anytype.io/schemas/page" }
  ],

  "type": "object"
}
`,
	"https://anytype.io/schemas/relation-definitions": `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "id": "https://anytype.io/schemas/relation-definitions",
  "title": "Anytype Page",
  "description": "This schema contains all base definitions",
  "definitions": {
    "relation": {
      "$id": "https://anytype.io/schemas/relation",
      "type": "object",
      "$comment": "fills relation from specific details",
      "properties": {
        "id": {
          "type": "string",
          "$comment": "detail's ID"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string",
          "$comment": "json schema $id for the relation type, starting from https://anytype.io/schemas/types/"
        }
      }
    },
    "title": {
      "$id": "https://anytype.io/schemas/types/title",
      "type": "string",
      "description": "Title renders name plus first emoji/image relation for the same relation"
    },
    "description": {
      "$id": "https://anytype.io/schemas/types/description",
      "type": "string"
    },
    "select": {
      "$id": "https://anytype.io/schemas/types/select",
      "type": "string"
    },
    "multiselect": {
      "$id": "https://anytype.io/schemas/types/multiselect",
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "number": {
      "$id": "https://anytype.io/schemas/types/number",
      "type": "number"
    },
    "url": {
      "$id": "https://anytype.io/schemas/types/url",
      "type": "string",
      "description": "External URL",
      "format": "uri"
    },
    "email": {
      "$id": "https://anytype.io/schemas/types/email",
      "type": "string",
      "format": "email"
    },
    "phone": {
      "$id": "https://anytype.io/schemas/types/phone",
      "type": "string"
    },
    "date": {
      "$id": "https://anytype.io/schemas/types/date",
      "type": "string",
      "description": "UNIX timestamp as a string"
    },
    "checkbox": {
      "$id": "https://anytype.io/schemas/types/checkbox",
      "type": "boolean"
    },
    "page": {
      "$id": "https://anytype.io/schemas/types/page",
      "type": "string",
      "description": "ID of the page"
    },
    "person": {
      "$id": "https://anytype.io/schemas/types/person",
      "type": "string",
      "description": "ID of the profile"
    },
    "image": {
      "$id": "https://anytype.io/schemas/types/image",
      "type": "string",
      "description": "CID of image node in the IPFS"
    },
    "file": {
      "$id": "https://anytype.io/schemas/types/file",
      "type": "string",
      "description": "CID of file node in the IPFS"
    },
    "emoji": {
      "$id": "https://anytype.io/schemas/types/emoji",
      "type": "string",
      "description": "One emoji as unicode"
    }
  }
}
`}
