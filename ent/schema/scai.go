// Copyright 2024 The Archivista Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AttributeAssertion holds the schema definition for the AttributeAssertion entity.
type AttributeAssertion struct {
	ent.Schema
}

// Fields of the AttributeAssertion.
func (AttributeAssertion) Fields() []ent.Field {
	return []ent.Field{
		field.String("attribute"),
		field.JSON("target", map[string]interface{}{}).
			Optional(),
		field.JSON("conditions", map[string]interface{}{}).
			Optional(),
		field.JSON("evidence", map[string]interface{}{}).
			Optional(),
	}
}

func (AttributeAssertion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attributes_report", AttributeReport.Type).Ref("attributes").Unique().Required(),
	}
}

func (AttributeAssertion) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}

// AttributeReport holds the schema definition for the AttributeReport entity.
type AttributeReport struct {
	ent.Schema
}

func (AttributeReport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attributes", AttributeAssertion.Type),
		edge.From("statement", Statement.Type).Ref("attributes_report").Unique().Required(),
	}
}
