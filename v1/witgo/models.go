// Copyright 2016 Arne Roomann-Kurrik
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package witgo

import (
	"fmt"
)

type Value struct {
	Expressions []string `json:"expressions,omitempty"`
	Value       string   `json:"value,omitempty"`
}

type Entity struct {
	Lang       string   `json:"lang,omitempty"`
	Closed     bool     `json:"closed,omitempty"`
	Exotic     bool     `json:"exotic,omitempty"`
	Value      string   `json:"value,omitempty"`
	Values     []*Value `json:"values,omitempty"`
	Builtin    bool     `json:"builtin,omitempty"`
	Doc        string   `json:"doc,omitempty"`
	Name       string   `json:"name,omitempty"`
	ID         string   `json:"id,omitempty"`
	Confidence float32  `json:"confidence,omitempty"`
	Type       string   `json:"type,omitempty"`
}

type EntityMap map[string][]*Entity

func (m EntityMap) FirstEntityValue(key string) (out string, err error) {
	var (
		entities []*Entity
		found bool
	)
	if entities, found = m[key]; !found || len(entities) == 0 {
		err = fmt.Errorf("No entities associated with key %v", key)
		return
	}
	out = entities[0].Value
	return
}

type ConverseResponse struct {
	Type         string    `json:"type,omitempty"`
	Msg          string    `json:"msg,omitempty"`
	Action       string    `json:"action,omitempty"`
	Entities     EntityMap `json:"entities,omitempty"`
	Confidence   float64   `json:"confidence,omitempty"`
	QuickReplies []string`json:"quickreplies,omitempty"`
}
