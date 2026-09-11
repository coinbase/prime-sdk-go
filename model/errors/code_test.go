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

package errors

import "testing"

func TestMetadataCompleteness(t *testing.T) {
	if len(errorCodeInfo) != 10 {
		t.Fatalf("errorCodeInfo size %d, want 10", len(errorCodeInfo))
	}
	if len(subcodeInfo) != 287 {
		t.Fatalf("subcodeInfo size %d, want 287", len(subcodeInfo))
	}

	codes := []ErrorCode{
		ErrorCodeResourceNotFound,
		ErrorCodeValidationError,
		ErrorCodeRequiredFieldMissing,
		ErrorCodeFailedPrecondition,
		ErrorCodeAuthenticationFailed,
		ErrorCodePermissionDenied,
		ErrorCodeRateLimitExceeded,
		ErrorCodeInternalError,
		ErrorCodeNotImplemented,
		ErrorCodeServiceUnavailable,
	}
	for _, c := range codes {
		info, ok := LookupErrorCode(c)
		if !ok || info.Description == "" || info.HTTPStatus == 0 {
			t.Errorf("missing metadata for %s", c)
		}
		if info.Code != c {
			t.Errorf("code mismatch %s vs %s", info.Code, c)
		}
	}

	info, ok := LookupSubcode(SubcodeOrderNotFound)
	if !ok {
		t.Fatal("ORDER_NOT_FOUND missing")
	}
	if info.ErrorCode != ErrorCodeResourceNotFound || info.HTTPStatus != 404 {
		t.Fatalf("ORDER_NOT_FOUND metadata %+v", info)
	}
}
