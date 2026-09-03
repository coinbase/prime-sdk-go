/**
 * Copyright 2026-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type errorCodeDef struct {
	Name        string
	HTTPStatus  int
	Description string
}

type subcodeDef struct {
	Name        string
	ErrorCode   string
	HTTPStatus  int
	Description string
}

func parseSpec(data []byte) ([]errorCodeDef, []subcodeDef, error) {
	lines := strings.Split(string(data), "\n")

	codesStart := -1
	subcodesStart := -1
	pathsStart := -1
	for i, line := range lines {
		switch line {
		case "  x-error-codes:":
			codesStart = i
		case "  x-subcodes:":
			subcodesStart = i
		case "paths:":
			pathsStart = i
		}
		if pathsStart >= 0 {
			break
		}
	}
	if codesStart < 0 || subcodesStart < 0 || pathsStart < 0 {
		return nil, nil, fmt.Errorf("spec missing x-error-codes, x-subcodes, or paths")
	}

	codeItems, err := parseListItems(lines[codesStart+1 : subcodesStart])
	if err != nil {
		return nil, nil, fmt.Errorf("x-error-codes: %w", err)
	}
	subItems, err := parseListItems(lines[subcodesStart+1 : pathsStart])
	if err != nil {
		return nil, nil, fmt.Errorf("x-subcodes: %w", err)
	}

	codes := make([]errorCodeDef, 0, len(codeItems))
	for _, item := range codeItems {
		c, err := toErrorCode(item)
		if err != nil {
			return nil, nil, err
		}
		codes = append(codes, c)
	}

	subcodes := make([]subcodeDef, 0, len(subItems))
	for _, item := range subItems {
		s, err := toSubcode(item)
		if err != nil {
			return nil, nil, err
		}
		subcodes = append(subcodes, s)
	}

	return codes, subcodes, nil
}

type listItem map[string]string

func parseListItems(lines []string) ([]listItem, error) {
	var items []listItem
	var cur listItem
	var curKey string
	var curVal strings.Builder
	inQuoted := false

	flushVal := func() {
		if curKey == "" {
			return
		}
		cur[curKey] = strings.TrimSpace(unescapeYAML(curVal.String()))
		curKey = ""
		curVal.Reset()
		inQuoted = false
	}

	flushItem := func() {
		flushVal()
		if cur != nil && cur["name"] != "" {
			items = append(items, cur)
		}
		cur = nil
	}

	for _, raw := range lines {
		if inQuoted {
			curVal.WriteString(raw)
			curVal.WriteByte('\n')
			if quotedScalarComplete(curVal.String()) {
				flushVal()
			}
			continue
		}

		if strings.HasPrefix(raw, "  - name:") {
			flushItem()
			cur = make(listItem)
			curKey = "name"
			curVal.WriteString(strings.TrimSpace(strings.TrimPrefix(raw, "  - name:")))
			flushVal()
			continue
		}

		if cur == nil {
			continue
		}

		trimmed := strings.TrimLeft(raw, " ")
		if strings.HasPrefix(raw, "    ") && strings.Contains(trimmed, ":") && !strings.HasPrefix(trimmed, "\"") {
			key, rest, ok := strings.Cut(trimmed, ":")
			if ok && isIdentKey(key) {
				flushVal()
				curKey = key
				rest = strings.TrimSpace(rest)
				curVal.WriteString(rest)
				if strings.HasPrefix(rest, "\"") && !quotedScalarComplete(rest) {
					curVal.WriteByte('\n')
					inQuoted = true
					continue
				}
				if strings.HasPrefix(rest, "\"") && quotedScalarComplete(rest) {
					flushVal()
				}
				// Unquoted (or empty) scalars may continue on following indented lines.
				continue
			}
		}

		if curKey != "" && strings.HasPrefix(raw, "      ") {
			if curVal.Len() > 0 {
				curVal.WriteByte(' ')
			}
			curVal.WriteString(strings.TrimSpace(raw))
			continue
		}
	}
	flushItem()
	return items, nil
}

func isIdentKey(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}
	return true
}

func quotedScalarComplete(s string) bool {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "\"") {
		return true
	}
	escaped := false
	for i := 1; i < len(s); i++ {
		c := s[i]
		if escaped {
			escaped = false
			continue
		}
		if c == '\\' {
			escaped = true
			continue
		}
		if c == '"' {
			return i == len(s)-1 || strings.TrimSpace(s[i+1:]) == ""
		}
	}
	return false
}

func unescapeYAML(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
		s = strings.ReplaceAll(s, "\\\n", "")
		s = strings.ReplaceAll(s, "\\ ", " ")
		s = strings.ReplaceAll(s, `\"`, `"`)
		s = strings.ReplaceAll(s, `\\`, `\`)
	}
	s = strings.Join(strings.Fields(s), " ")
	return s
}

func toErrorCode(item listItem) (errorCodeDef, error) {
	status, err := strconv.Atoi(item["httpStatus"])
	if err != nil {
		return errorCodeDef{}, fmt.Errorf("error code %s: invalid httpStatus %q", item["name"], item["httpStatus"])
	}
	if item["description"] == "" {
		return errorCodeDef{}, fmt.Errorf("error code %s: missing description", item["name"])
	}
	return errorCodeDef{
		Name:        item["name"],
		HTTPStatus:  status,
		Description: item["description"],
	}, nil
}

func toSubcode(item listItem) (subcodeDef, error) {
	status, err := strconv.Atoi(item["httpStatus"])
	if err != nil {
		return subcodeDef{}, fmt.Errorf("subcode %s: invalid httpStatus %q", item["name"], item["httpStatus"])
	}
	if item["errorCode"] == "" || item["description"] == "" {
		return subcodeDef{}, fmt.Errorf("subcode %s: missing errorCode or description", item["name"])
	}
	return subcodeDef{
		Name:        item["name"],
		ErrorCode:   item["errorCode"],
		HTTPStatus:  status,
		Description: item["description"],
	}, nil
}
