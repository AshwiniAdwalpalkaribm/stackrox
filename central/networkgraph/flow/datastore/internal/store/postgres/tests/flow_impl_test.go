//go:build sql_integration

package tests

import (
	"testing"

	postgresFlowStore "github.com/stackrox/rox/central/networkgraph/flow/datastore/internal/store/postgres"
	"github.com/stackrox/rox/central/networkgraph/flow/datastore/internal/store/testcommon"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stretchr/testify/suite"
)

func TestFlowStore(t *testing.T) {
	testDB := pgtest.ForT(t)
	defer testDB.DB.Close()

	store := postgresFlowStore.NewClusterStore(testDB.DB)
	flowSuite := testcommon.NewFlowStoreTest(store)
	suite.Run(t, flowSuite)
}
