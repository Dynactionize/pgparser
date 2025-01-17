// Copyright 2015 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package security_test

import (
	"os"
	"testing"

	"github.com/ruiaylin/pgparser/security"
	"github.com/ruiaylin/pgparser/security/securitytest"
	"github.com/ruiaylin/pgparser/server"
	"github.com/ruiaylin/pgparser/testutils/serverutils"
	_ "github.com/ruiaylin/pgparser/utils/log" // for flags
)

// ResetTest sets up the test environment. In particular, it embeds the
// EmbeddedCertsDir folder and makes the tls package load from there.
func ResetTest() {
	security.SetAssetLoader(securitytest.EmbeddedAssets)
}

func TestMain(m *testing.M) {
	ResetTest()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	os.Exit(m.Run())
}

//go:generate ../util/leaktest/add-leaktest.sh *_test.go
