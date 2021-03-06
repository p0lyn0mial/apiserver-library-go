/*
Copyright 2016 The Kubernetes Authors.

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

package app

import (
	"fmt"

	"net/http"

	"k8s.io/kubernetes/pkg/controller/bootstrap"
)

func startBootstrapSignerController(ctx ControllerContext) (http.Handler, bool, error) {
	bsc, err := bootstrap.NewBootstrapSigner(
		ctx.ClientBuilder.ClientOrDie("bootstrap-signer"),
		ctx.InformerFactory.Core().V1().Secrets(),
		ctx.InformerFactory.Core().V1().ConfigMaps(),
		bootstrap.DefaultBootstrapSignerOptions(),
	)
	if err != nil {
		return nil, true, fmt.Errorf("error creating BootstrapSigner controller: %v", err)
	}
	go bsc.Run(ctx.Stop)
	return nil, true, nil
}

func startTokenCleanerController(ctx ControllerContext) (http.Handler, bool, error) {
	tcc, err := bootstrap.NewTokenCleaner(
		ctx.ClientBuilder.ClientOrDie("token-cleaner"),
		ctx.InformerFactory.Core().V1().Secrets(),
		bootstrap.DefaultTokenCleanerOptions(),
	)
	if err != nil {
		return nil, true, fmt.Errorf("error creating TokenCleaner controller: %v", err)
	}
	go tcc.Run(ctx.Stop)
	return nil, true, nil
}
