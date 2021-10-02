package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: Post
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] title                                          TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] body                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] published                                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
[ 4] user_id                                        INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 5] createdAt                                      TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
[ 6] updatedAt                                      TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 86,    "title": "yZsqNAZtuXaQMjRidPcFZWfLl",    "body": "aayujCgdMrydRqickhVyQjJrL",    "published": false,    "user_id": 59,    "created_at": "2053-05-25T16:18:33.711750301-04:00",    "updated_at": "2266-11-24T20:01:39.325315716-04:00"}



*/

// Post struct is a row record of the Post table in the myApp database
type Post struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:INT4;"`
	//[ 1] title                                          TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Title string `gorm:"column:title;type:TEXT;"`
	//[ 2] body                                           TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Body string `gorm:"column:body;type:TEXT;"`
	//[ 3] published                                      BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Published bool `gorm:"column:published;type:BOOL;"`
	//[ 4] user_id                                        INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	UserID int32 `gorm:"column:user_id;type:INT4;"`
	//[ 5] createdAt                                      TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:createdAt;type:TIMESTAMP;default:CURRENT_TIMESTAMP;"`
	//[ 6] updatedAt                                      TIMESTAMP            null: false  primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updatedAt;type:TIMESTAMP;"`
}

var PostTableInfo = &TableInfo{
	Name: "Post",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "title",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "body",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Body",
			GoFieldType:        "string",
			JSONFieldName:      "body",
			ProtobufFieldName:  "body",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "published",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "Published",
			GoFieldType:        "bool",
			JSONFieldName:      "published",
			ProtobufFieldName:  "published",
			ProtobufType:       "bool",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "user_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "createdAt",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "updatedAt",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Post) TableName() string {
	return "Post"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Post) BeforeSave(*gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Post) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Post) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Post) TableInfo() *TableInfo {
	return PostTableInfo
}
