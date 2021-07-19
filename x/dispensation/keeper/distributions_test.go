package keeper_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Sifchain/sifnode/x/dispensation/test"
	"github.com/Sifchain/sifnode/x/dispensation/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestKeeper_GetDistributions(t *testing.T) {
	app, ctx := test.CreateTestApp(false)
	distributionTypes := []types.DistributionType{types.DistributionType_DISTRIBUTION_TYPE_AIRDROP, types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING, types.DistributionType_DISTRIBUTION_TYPE_VALIDATOR_SUBSIDY}
	keeper := app.DispensationKeeper
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // local pseudorandom generator
	list := keeper.GetDistributions(ctx)
	assert.Len(t, list.Distributions, 0)
	for i := 0; i < 10; i++ {
		name := uuid.New().String()
		selectType := distributionTypes[r.Intn(len(distributionTypes))]
		distribution := types.NewDistribution(selectType, name)
		err := keeper.SetDistribution(ctx, distribution)
		assert.NoError(t, err)
		_, err = keeper.GetDistribution(ctx, name, selectType)
		assert.NoError(t, err)
	}
	list = keeper.GetDistributions(ctx)
	assert.Len(t, list.Distributions, 10)
}

func TestKeeper_FailGetDistributionsIfNotSet(t *testing.T) {
	app, ctx := test.CreateTestApp(false)
	keeper := app.DispensationKeeper
	list := keeper.GetDistributions(ctx)
	assert.Len(t, list.Distributions, 0)
	for i := 0; i < 5; i++ {
		name := uuid.New().String()
		selectType := types.DistributionType_DISTRIBUTION_TYPE_AIRDROP
		_, err := keeper.GetDistribution(ctx, name, selectType)
		assert.Error(t, err)
	}
	list = keeper.GetDistributions(ctx)
	assert.Len(t, list.Distributions, 0)
}

func TestKeeper_GetDistribution(t *testing.T) {
	app, ctx := test.CreateTestApp(false)
	keeper := app.DispensationKeeper
	name := uuid.New().String()
	selectType := types.DistributionType_DISTRIBUTION_TYPE_AIRDROP
	distribution := types.NewDistribution(selectType, name)
	assert.Equal(t, distribution.DistributionName, name)
	assert.Equal(t, distribution.DistributionType, selectType)
	err := keeper.SetDistribution(ctx, distribution)
	assert.NoError(t, err)
	distr, err := keeper.GetDistribution(ctx, name, selectType)
	assert.NoError(t, err)
	assert.Equal(t, distr.DistributionName, name)
	assert.Equal(t, distr.DistributionType, selectType)
}
