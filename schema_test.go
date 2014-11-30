package swagger

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var schema = Schema{
	Ref:              "Cat",
	Description:      "the description of this schema",
	Maximum:          100,
	ExclusiveMaximum: true,
	ExclusiveMinimum: true,
	Minimum:          5,
	MaxLength:        100,
	MinLength:        5,
	Pattern:          "\\w{1,5}\\w+",
	MaxItems:         100,
	MinItems:         5,
	UniqueItems:      true,
	MultipleOf:       5,
	Enum:             []interface{}{"hello", "world"},
	Type:             &StringOrArray{Single: "string"},
	Format:           "date",
	Title:            "the title",
	Default:          "blah",
	MaxProperties:    5,
	MinProperties:    1,
	Required:         []string{"id", "name"},
	Items:            &SchemaOrArray{Single: &Schema{Type: &StringOrArray{Single: "string"}}},
	AllOf:            []Schema{Schema{Type: &StringOrArray{Single: "string"}}},
	Properties: map[string]Schema{
		"id":   Schema{Type: &StringOrArray{Single: "integer"}, Format: "int64"},
		"name": Schema{Type: &StringOrArray{Single: "string"}},
	},
	Discriminator: "not this",
	ReadOnly:      true,
	XML:           &XMLObject{"sch", "swagger.io", "sw", true, true},
	ExternalDocs: &ExternalDocumentation{
		Description: "the documentation etc",
		URL:         "http://readthedocs.org/swagger",
	},
	Example: []interface{}{
		map[string]interface{}{
			"id":   1,
			"name": "a book",
		},
		map[string]interface{}{
			"id":   2,
			"name": "the thing",
		},
	},
	AdditionalProperties: &Schema{
		Type:   &StringOrArray{Single: "integer"},
		Format: "int32",
	},
	Extensions: map[string]interface{}{"x-framework": "go-swagger"},
}

var schemaJson = `{
	"x-framework": "go-swagger",
  "$ref": "Cat",
  "description": "the description of this schema",
  "maximum": 100,
  "minimum": 5,
  "exclusiveMaximum": true,
  "exclusiveMinimum": true,
  "maxLength": 100,
  "minLength": 5,
  "pattern": "\\w{1,5}\\w+",
  "maxItems": 100,
  "minItems": 5,
  "uniqueItems": true,
  "multipleOf": 5,
  "enum": ["hello", "world"],
  "type": "string",
  "format": "date",
  "title": "the title",
  "default": "blah",
  "maxProperties": 5,
  "minProperties": 1,
  "required": ["id", "name"],
  "items": { 
    "type": "string" 
  },
  "allOf": [
    { 
      "type": "string" 
    }
  ],
  "properties": {
    "id": { 
      "type": "integer",
      "format": "int64"
    },
    "name": {
      "type": "string"
    }
  },
  "discriminator": "not this",
  "readOnly": true,
  "xml": {
    "name": "sch",
    "namespace": "swagger.io",
    "prefix": "sw",
    "wrapped": true,
    "attribute": true
  },
  "externalDocs": {
    "description": "the documentation etc",
    "url": "http://readthedocs.org/swagger"
  },
  "example": [
    {
      "id": 1,
      "name": "a book"
    },
    { 
      "id": 2,
      "name": "the thing"
    }
  ],
  "additionalProperties": {
    "type": "integer",
    "format": "int32"
  }
}
`

func TestSchema(t *testing.T) {

	Convey("Schema should", t, func() {

		Convey("serialize", func() {
			expected := map[string]interface{}{}
			json.Unmarshal([]byte(schemaJson), &expected)
			b, err := json.Marshal(schema)
			So(err, ShouldBeNil)
			var actual map[string]interface{}
			json.Unmarshal(b, &actual)
			So(actual, ShouldBeEquivalentTo, expected)
		})
		Convey("deserialize", func() {
			actual := Schema{}
			err := json.Unmarshal([]byte(schemaJson), &actual)
			So(err, ShouldBeNil)
			So(actual.Ref, ShouldEqual, schema.Ref)
			So(actual.Description, ShouldEqual, schema.Description)
			So(actual.Maximum, ShouldEqual, schema.Maximum)
			So(actual.Minimum, ShouldEqual, schema.Minimum)
			So(actual.ExclusiveMinimum, ShouldEqual, schema.ExclusiveMinimum)
			So(actual.ExclusiveMaximum, ShouldEqual, schema.ExclusiveMaximum)
			So(actual.MaxLength, ShouldEqual, schema.MaxLength)
			So(actual.MinLength, ShouldEqual, schema.MinLength)
			So(actual.Pattern, ShouldEqual, schema.Pattern)
			So(actual.MaxItems, ShouldEqual, schema.MaxItems)
			So(actual.MinItems, ShouldEqual, schema.MinItems)
			So(actual.UniqueItems, ShouldBeTrue)
			So(actual.MultipleOf, ShouldEqual, schema.MultipleOf)
			So(actual.Enum, ShouldResemble, schema.Enum)
			So(actual.Type, ShouldResemble, schema.Type)
			So(actual.Format, ShouldEqual, schema.Format)
			So(actual.Title, ShouldEqual, schema.Title)
			So(actual.Default, ShouldResemble, schema.Default)
			So(actual.MaxProperties, ShouldResemble, schema.MaxProperties)
			So(actual.MinProperties, ShouldResemble, schema.MinProperties)
			So(actual.Required, ShouldResemble, schema.Required)
			So(actual.Items, ShouldResemble, schema.Items)
			So(actual.AllOf, ShouldResemble, schema.AllOf)
			So(actual.Properties, ShouldResemble, schema.Properties)
			So(actual.Discriminator, ShouldEqual, schema.Discriminator)
			So(actual.ReadOnly, ShouldEqual, schema.ReadOnly)
			So(actual.XML, ShouldResemble, schema.XML)
			So(actual.ExternalDocs, ShouldResemble, schema.ExternalDocs)
			examples := actual.Example.([]interface{})
			expEx := schema.Example.([]interface{})
			ex1 := examples[0].(map[string]interface{})
			ex2 := examples[1].(map[string]interface{})
			exp1 := expEx[0].(map[string]interface{})
			exp2 := expEx[1].(map[string]interface{})
			So(ex1["name"], ShouldEqual, exp1["name"])
			So(ex1["id"], ShouldEqual, exp1["id"])
			So(ex2["name"], ShouldEqual, exp2["name"])
			So(ex2["id"], ShouldEqual, exp2["id"])
			//So(actual.Example, ShouldBeEquivalentTo, schema.Example)
			So(actual.AdditionalProperties, ShouldResemble, schema.AdditionalProperties)
			So(actual.Extensions, ShouldBeEquivalentTo, schema.Extensions)
		})
	})

}
