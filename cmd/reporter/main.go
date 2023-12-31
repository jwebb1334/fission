/*
Copyright 2021 The Fission Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"go.uber.org/zap"

	"github.com/jwebb1334/fission/cmd/reporter/app"
	"github.com/jwebb1334/fission/pkg/utils/loggerfactory"
)

func main() {
	logger := loggerfactory.GetLogger()
	defer logger.Sync()

	err := app.App().Execute()
	if err != nil {
		logger.Error("error occurred during analytics reporting", zap.Error(err))
	}
}
