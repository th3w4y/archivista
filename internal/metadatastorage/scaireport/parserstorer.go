// Copyright 2023 The Archivista Contributors
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

package scaireport

import (
	"context"
	"encoding/json"

	"github.com/in-toto/archivista/ent"
	"github.com/in-toto/archivista/internal/metadatastorage"
)

const (
	Predicate = "https://in-toto.io/attestation/scai/attribute-report/v0.2"
)

// attestation.Collection from go-witness will try to parse each of the attestations by calling their factory functions,
// which require the attestations to be registered in the go-witness library.  We don't really care about the actual attestation
// data for the purposes here, so just leave it as a raw message.
type ParsedAtributesReport struct {
	Attributes []struct {
		Attribute  string          `json:"attribute"`
		Evidence   json.RawMessage `json:"evidence"`
		Target     json.RawMessage `json:"target"`
		Conditions json.RawMessage `json:"conditions"`
	} `json:"attributes"`
}

func Parse(data []byte) (metadatastorage.Storer, error) {
	parsedCollection := ParsedAtributesReport{}
	if err := json.Unmarshal(data, &parsedCollection); err != nil {
		return parsedCollection, err
	}

	return parsedCollection, nil
}

func (parsedReport ParsedAtributesReport) Store(ctx context.Context, tx *ent.Tx, stmtID int) error {
	report, err := tx.AttributeReport.Create().
		SetStatementID(stmtID).
		Save(ctx)
	if err != nil {
		return err
	}

	for _, a := range parsedReport.Attributes {
		target := make(map[string]interface{})
		json.Unmarshal(a.Target, &target)
		conditions := make(map[string]interface{})
		json.Unmarshal(a.Conditions, &conditions)
		evidence := make(map[string]interface{})
		json.Unmarshal(a.Evidence, &evidence)
		if err := tx.AttributeAssertion.Create().
			SetAttributesReportID(report.ID).
			SetAttribute(a.Attribute).
			SetTarget(target).
			SetConditions(conditions).
			SetEvidence(evidence).
			Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}
