package keeper_test

import (
	"fmt"
	"github.com/osmosis-labs/osmosis/v11/x/gamm/pool-models/balancer"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v11/x/gamm/types"
)

func (suite *KeeperTestSuite) TestBalancerPoolSimpleMultihopSwapExactAmountIn() {
	type param struct {
		routes            []types.SwapAmountInRoute
		tokenIn           sdk.Coin
		tokenOutMinAmount sdk.Int
	}

	tests := []struct {
		name              string
		param             param
		expectPass        bool
		reducedFeeApplied bool
	}{
		{
			name: "Proper swap - foo -> bar(pool 1) - bar(pool 2) -> baz",
			param: param{
				routes: []types.SwapAmountInRoute{
					{
						PoolId:        1,
						TokenOutDenom: "bar",
					},
					{
						PoolId:        2,
						TokenOutDenom: "baz",
					},
				},
				tokenIn:           sdk.NewCoin("foo", sdk.NewInt(100000)),
				tokenOutMinAmount: sdk.NewInt(1),
			},
			expectPass: true,
		},
		{
			name: "Swap - foo -> uosmo(pool 1) - uosmo(pool 2) -> baz with a half fee applied",
			param: param{
				routes: []types.SwapAmountInRoute{
					{
						PoolId:        1,
						TokenOutDenom: "uosmo",
					},
					{
						PoolId:        2,
						TokenOutDenom: "baz",
					},
				},
				tokenIn:           sdk.NewCoin("foo", sdk.NewInt(100000)),
				tokenOutMinAmount: sdk.NewInt(1),
			},
			reducedFeeApplied: true,
			expectPass:        true,
		},
	}

	for _, test := range tests {
		// Init suite for each test.
		suite.SetupTest()

		suite.Run(test.name, func() {
			// Prepare 2 pools
			suite.PrepareBalancerPoolWithPoolParams(balancer.PoolParams{
				SwapFee: sdk.NewDecWithPrec(1, 2),
				ExitFee: sdk.NewDec(0),
			})
			suite.PrepareBalancerPoolWithPoolParams(balancer.PoolParams{
				SwapFee: sdk.NewDecWithPrec(1, 2),
				ExitFee: sdk.NewDec(0),
			})

			keeper := suite.App.GAMMKeeper

			if test.expectPass {
				// Calculate the chained spot price.
				spotPriceBefore := func() sdk.Dec {
					dec := sdk.NewDec(1)
					tokenInDenom := test.param.tokenIn.Denom
					for i, route := range test.param.routes {
						if i != 0 {
							tokenInDenom = test.param.routes[i-1].TokenOutDenom
						}
						pool, err := keeper.GetPoolAndPoke(suite.Ctx, route.PoolId)
						suite.NoError(err, "test: %v", test.name)

						sp, err := pool.SpotPrice(suite.Ctx, tokenInDenom, route.TokenOutDenom)
						suite.NoError(err, "test: %v", test.name)
						dec = dec.Mul(sp)
					}
					return dec
				}()

				calcOutAmountAsSeparateSwaps := func(feeMultiplier sdk.Dec) sdk.Coin {
					cacheCtx, _ := suite.Ctx.CacheContext()
					nextTokenIn := test.param.tokenIn
					for _, hop := range test.param.routes {
						tokenOut, err := keeper.SwapExactAmountIn(cacheCtx, suite.TestAccs[0], hop.PoolId, nextTokenIn, hop.TokenOutDenom, sdk.OneInt(), feeMultiplier)
						suite.Require().NoError(err)
						nextTokenIn = sdk.NewCoin(hop.TokenOutDenom, tokenOut)
					}
					return nextTokenIn
				}

				tokenOutCalculatedAsSeparateSwaps := calcOutAmountAsSeparateSwaps(sdk.OneDec())
				tokenOutCalculatedAsSeparateSwapsWithoutFee := calcOutAmountAsSeparateSwaps(sdk.ZeroDec())

				tokenOutAmount, err := keeper.MultihopSwapExactAmountIn(suite.Ctx, suite.TestAccs[0], test.param.routes, test.param.tokenIn, test.param.tokenOutMinAmount)
				suite.NoError(err, "test: %v", test.name)

				// Calculate the chained spot price.
				spotPriceAfter := func() sdk.Dec {
					dec := sdk.NewDec(1)
					tokenInDenom := test.param.tokenIn.Denom
					for i, route := range test.param.routes {
						if i != 0 {
							tokenInDenom = test.param.routes[i-1].TokenOutDenom
						}

						pool, err := keeper.GetPoolAndPoke(suite.Ctx, route.PoolId)
						suite.NoError(err, "test: %v", test.name)

						sp, err := pool.SpotPrice(suite.Ctx, tokenInDenom, route.TokenOutDenom)
						suite.NoError(err, "test: %v", test.name)
						dec = dec.Mul(sp)
					}
					return dec
				}()

				// Ratio of the token out should be between the before spot price and after spot price.
				sp := test.param.tokenIn.Amount.ToDec().Quo(tokenOutAmount.ToDec())
				suite.True(sp.GT(spotPriceBefore) && sp.LT(spotPriceAfter), "test: %v", test.name)

				if test.reducedFeeApplied {
					suite.Require().True(tokenOutAmount.GT(tokenOutCalculatedAsSeparateSwaps.Amount))
					suite.Require().True(tokenOutAmount.LTE(tokenOutCalculatedAsSeparateSwapsWithoutFee.Amount))
				} else {
					suite.Require().True(tokenOutAmount.Equal(tokenOutCalculatedAsSeparateSwaps.Amount))
				}

			} else {
				_, err := keeper.MultihopSwapExactAmountIn(suite.Ctx, suite.TestAccs[0], test.param.routes, test.param.tokenIn, test.param.tokenOutMinAmount)
				suite.Error(err, "test: %v", test.name)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestBalancerPoolSimpleMultihopSwapExactAmountOut() {
	type param struct {
		routes           []types.SwapAmountOutRoute
		tokenInMaxAmount sdk.Int
		tokenOut         sdk.Coin
	}

	tests := []struct {
		name              string
		param             param
		expectPass        bool
		reducedFeeApplied bool
	}{
		{
			name: "Proper swap: foo -> bar (pool 1), bar -> baz (pool 2)",
			param: param{
				routes: []types.SwapAmountOutRoute{
					{
						PoolId:       1,
						TokenInDenom: "foo",
					},
					{
						PoolId:       2,
						TokenInDenom: "bar",
					},
				},
				tokenInMaxAmount: sdk.NewInt(90000000),
				tokenOut:         sdk.NewCoin("baz", sdk.NewInt(100000)),
			},
			expectPass: true,
		},
		{
			name: "Swap - foo -> uosmo(pool 1) - uosmo(pool 2) -> baz with a half fee applied",
			param: param{
				routes: []types.SwapAmountOutRoute{
					{
						PoolId:       1,
						TokenInDenom: "foo",
					},
					{
						PoolId:       2,
						TokenInDenom: "uosmo",
					},
				},
				tokenInMaxAmount: sdk.NewInt(90000000),
				tokenOut:         sdk.NewCoin("baz", sdk.NewInt(100000)),
			},
			expectPass:        true,
			reducedFeeApplied: true,
		},
	}

	for _, test := range tests {
		// Init suite for each test.
		suite.SetupTest()

		suite.Run(test.name, func() {
			// Prepare 2 pools
			suite.PrepareBalancerPoolWithPoolParams(balancer.PoolParams{
				SwapFee: sdk.NewDecWithPrec(1, 2),
				ExitFee: sdk.NewDec(0),
			})
			suite.PrepareBalancerPoolWithPoolParams(balancer.PoolParams{
				SwapFee: sdk.NewDecWithPrec(1, 2),
				ExitFee: sdk.NewDec(0),
			})

			keeper := suite.App.GAMMKeeper

			if test.expectPass {
				// Calculate the chained spot price.
				spotPriceBefore := func() sdk.Dec {
					dec := sdk.NewDec(1)
					for i, route := range test.param.routes {
						tokenOutDenom := test.param.tokenOut.Denom
						if i != len(test.param.routes)-1 {
							tokenOutDenom = test.param.routes[i+1].TokenInDenom
						}

						pool, err := keeper.GetPoolAndPoke(suite.Ctx, route.PoolId)
						suite.NoError(err, "test: %v", test.name)

						sp, err := pool.SpotPrice(suite.Ctx, route.TokenInDenom, tokenOutDenom)
						suite.NoError(err, "test: %v", test.name)
						dec = dec.Mul(sp)
					}
					return dec
				}()

				calcInAmountAsSeparateSwaps := func(feeMultiplier sdk.Dec) sdk.Coin {
					cacheCtx, _ := suite.Ctx.CacheContext()
					nextTokenOut := test.param.tokenOut
					for i := len(test.param.routes) - 1; i >= 0; i-- {
						hop := test.param.routes[i]
						tokenOut, err := keeper.SwapExactAmountOut(cacheCtx, suite.TestAccs[0], hop.PoolId, hop.TokenInDenom, sdk.NewInt(100000000), nextTokenOut, feeMultiplier)
						suite.Require().NoError(err)
						nextTokenOut = sdk.NewCoin(hop.TokenInDenom, tokenOut)
					}
					return nextTokenOut
				}

				tokenInCalculatedAsSeparateSwaps := calcInAmountAsSeparateSwaps(sdk.OneDec())
				tokenInCalculatedAsSeparateSwapsWithoutFee := calcInAmountAsSeparateSwaps(sdk.ZeroDec())

				tokenInAmount, err := keeper.MultihopSwapExactAmountOut(suite.Ctx, suite.TestAccs[0], test.param.routes, test.param.tokenInMaxAmount, test.param.tokenOut)
				suite.Require().NoError(err, "test: %v", test.name)

				// Calculate the chained spot price.
				spotPriceAfter := func() sdk.Dec {
					dec := sdk.NewDec(1)
					for i, route := range test.param.routes {
						tokenOutDenom := test.param.tokenOut.Denom
						if i != len(test.param.routes)-1 {
							tokenOutDenom = test.param.routes[i+1].TokenInDenom
						}

						pool, err := keeper.GetPoolAndPoke(suite.Ctx, route.PoolId)
						suite.NoError(err, "test: %v", test.name)

						sp, err := pool.SpotPrice(suite.Ctx, route.TokenInDenom, tokenOutDenom)
						suite.NoError(err, "test: %v", test.name)
						dec = dec.Mul(sp)
					}
					return dec
				}()

				// Ratio of the token out should be between the before spot price and after spot price.
				// This is because the swap increases the spot price
				sp := tokenInAmount.ToDec().Quo(test.param.tokenOut.Amount.ToDec())
				fmt.Printf("spBefore %s, spAfter %s, sp actual %s\n", spotPriceBefore, spotPriceAfter, sp)
				suite.True(sp.GT(spotPriceBefore) && sp.LT(spotPriceAfter), "multi-hop spot price wrong, test: %v", test.name)

				if test.reducedFeeApplied {
					suite.Require().True(tokenInAmount.LT(tokenInCalculatedAsSeparateSwaps.Amount))
					suite.Require().True(tokenInAmount.GTE(tokenInCalculatedAsSeparateSwapsWithoutFee.Amount))
				} else {
					suite.Require().True(tokenInAmount.Equal(tokenInCalculatedAsSeparateSwaps.Amount))
				}
			} else {
				_, err := keeper.MultihopSwapExactAmountOut(suite.Ctx, suite.TestAccs[0], test.param.routes, test.param.tokenInMaxAmount, test.param.tokenOut)
				suite.Error(err, "test: %v", test.name)
			}
		})
	}
}
