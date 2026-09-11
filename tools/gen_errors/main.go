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
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	specPath := flag.String("spec", "apiSpec/prime-public-api-spec.yaml", "path to the Prime OpenAPI spec")
	outDir := flag.String("out", "model/errors", "output directory for generated files")
	flag.Parse()

	spec, err := os.ReadFile(*specPath)
	if err != nil {
		fatalf("read spec: %v", err)
	}

	codes, subcodes, err := parseSpec(spec)
	if err != nil {
		fatalf("parse spec: %v", err)
	}

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		fatalf("mkdir: %v", err)
	}

	files, err := generate(*outDir, codes, subcodes)
	if err != nil {
		fatalf("generate: %v", err)
	}

	for _, f := range files {
		abs, _ := filepath.Abs(f)
		fmt.Println(abs)
	}
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
