// Copyright 2016 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package eval_test

import (
	"os"
	"testing"

	"github.com/ruiaylin/pgparser/security"
	"github.com/ruiaylin/pgparser/security/securitytest"
	"github.com/ruiaylin/pgparser/pkg/server"
	"github.com/ruiaylin/pgparser/pkg/testutils/serverutils"
	"github.com/ruiaylin/pgparser/pkg/testutils/testcluster"
	"github.com/ruiaylin/pgparser/utils/randutil"
)

//go:generate ../../../../util/leaktest/add-leaktest.sh *_test.go

func TestMain(m *testing.M) {
	security.SetAssetLoader(securitytest.EmbeddedAssets)
	randutil.SeedForTests()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	serverutils.InitTestClusterFactory(testcluster.TestClusterFactory)
	os.Exit(m.Run())
}
