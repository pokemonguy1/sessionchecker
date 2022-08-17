package model

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SeancePropertiesStatistics struct {
	Reserved  int64 `bson:"reserved" json:"reserved,omitempty"`
	Payed     int64 `bson:"payed" json:"payed,omitempty"`
	Activated int64 `bson:"activated" json:"activated,omitempty"`
	Returned  int64 `bson:"returned" json:"returned,omitempty"`
	Cleared   int64 `bson:"cleared" json:"cleared,omitempty"`
}

type SeanceProperties struct {
	Format     []string                    `bson:"format" json:"format,omitempty"`           // Format
	Language   string                      `bson:"language" json:"language,omitempty"`       // Language
	IsPreSale  bool                        `bson:"is_pre_sale" json:"is_pre_sale,omitempty"` //IsPreSale
	Statistics *SeancePropertiesStatistics `bson:"statistics" json:"statistics,omitempty"`
}

type SeanceTimeframe struct {
	Start    *string `bson:"start" json:"start,omitempty"`       // Start time
	End      *string `bson:"end" json:"end,omitempty"`           // End time
	Interval uint32  `bson:"interval" json:"interval,omitempty"` // Internal
}

type SeanceDeadline struct {
	Name       string `bson:"name" json:"name,omitempty"`
	ContractId string `bson:"contract_id" json:"contract_id,omitempty"` // Contract id
	Type       string `bson:"type" json:"type,omitempty"`               // Deadline type
	Start      string `bson:"start" json:"start,omitempty"`             // Start time
	End        string `bson:"end" json:"end,omitempty"`                 // End time
	Automatic  bool   `bson:"automatic" json:"automatic,omitempty"`     // Auto change when changing timeframe
}

type SeanceDeadlines []*SeanceDeadline

func (m SeanceDeadlines) RemoveByType(t string) (SeanceDeadlines, error) {
	var index = -1
	for i := 0; i < len(m); i++ {
		if m[i].Type == t {
			index = i
		}
	}

	if index == -1 {
		return nil, fmt.Errorf("can't find deadline with type %s", t)
	}

	copy(m[index:], m[index+1:])
	m[len(m)-1] = nil
	m = m[:len(m)-1]

	return m, nil
}

type SeanceDiscount struct {
	Id          string  `bson:"id" json:"id,omitempty"`                   // local id
	Type        string  `bson:"type" json:"type,omitempty"`               // type
	Value       float64 `bson:"value" json:"value,omitempty"`             // value
	Name        string  `bson:"name" json:"name,omitempty"`               // Object name
	Code        string  `bson:"code" json:"code,omitempty"`               // Unique code or slug
	Description string  `bson:"description" json:"description,omitempty"` // Description
}

type SeanceReference struct {
	City   *primitive.ObjectID `bson:"city" json:"city,omitempty"`     // Object id
	Object *primitive.ObjectID `bson:"object" json:"object,omitempty"` // Object id
	Hall   *primitive.ObjectID `bson:"hall" json:"hall,omitempty"`     // Object id
	Plan   *primitive.ObjectID `bson:"plan" json:"plan,omitempty"`     // Object id
	Movie  *primitive.ObjectID `bson:"movie" json:"movie,omitempty"`   // Object id
}

type Seance struct {
	Id          primitive.ObjectID `bson:"_id" json:"id,omitempty"`                // Local id
	Properties  *SeanceProperties  `bson:"properties" json:"properties,omitempty"` // Properties
	Timeframe   *SeanceTimeframe   `bson:"timeframe" json:"timeframe,omitempty"`   // Timeframes
	Deadlines   SeanceDeadlines    `bson:"deadlines" json:"deadlines,omitempty"`   // time intervals
	Discounts   []*SeanceDiscount  `bson:"discounts" json:"discounts,omitempty"`
	Uuid        string             `bson:"uuid" json:"uuid,omitempty"`                      // Unique object uuid
	SortOrder   int64              `bson:"sort_order" json:"sort_order,omitempty"`          // Sort order
	IsActive    *bool              `bson:"is_active" json:"is_active,omitempty"`            // Is active?
	Name        string             `bson:"name" json:"name,omitempty"`                      // Object name
	Code        string             `bson:"code" json:"code,omitempty"`                      // Unique code or slug
	Description string             `bson:"description" json:"description,omitempty"`        // Description
	Reference   *SeanceReference   `bson:"reference" json:"reference,omitempty"`            // References
	Comment     string             `bson:"comment" json:"comment"`                          // Comment
	Hash        string             `bson:"hash" json:"hash,omitempty" hash:"-"`             // Hash fields
	CreatedAt   *string            `bson:"created_at" json:"created_at,omitempty" hash:"-"` // Indicate created at
	UpdatedAt   *string            `bson:"updated_at" json:"updated_at,omitempty" hash:"-"` // Indicate updated at
	DeletedAt   *string            `bson:"deleted_at" json:"deleted_at,omitempty" hash:"-"` // Indicate deleted at
}
