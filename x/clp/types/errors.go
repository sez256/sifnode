package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalid                         = sdkerrors.Register(ModuleName, 1, "invalid")
	ErrPoolDoesNotExist                = sdkerrors.Register(ModuleName, 2, "pool does not exist")
	ErrLiquidityProviderDoesNotExist   = sdkerrors.Register(ModuleName, 3, "liquidity Provider does not exist")
	ErrInValidAsset                    = sdkerrors.Register(ModuleName, 4, "asset is invalid")
	ErrInValidAmount                   = sdkerrors.Register(ModuleName, 5, "amount is invalid")
	ErrTotalAmountTooLow               = sdkerrors.Register(ModuleName, 7, "total amount is less than minimum threshold")
	ErrNotEnoughAssetTokens            = sdkerrors.Register(ModuleName, 8, "not enough received asset tokens to swap")
	ErrInvalidAsymmetry                = sdkerrors.Register(ModuleName, 9, "Asymmetry has to be between -10000 and 10000")
	ErrInvalidWBasis                   = sdkerrors.Register(ModuleName, 10, "WBasisPoints has to be positive")
	ErrBalanceTooHigh                  = sdkerrors.Register(ModuleName, 11, "Pool Balance too high to be decommissioned")
	ErrUnableToSetPool                 = sdkerrors.Register(ModuleName, 12, "Unable to set pool")
	ErrUnableToDestroyPool             = sdkerrors.Register(ModuleName, 13, "Unable to destroy pool")
	ErrUnableToCreatePool              = sdkerrors.Register(ModuleName, 14, "Unable to create pool")
	ErrBalanceNotAvailable             = sdkerrors.Register(ModuleName, 18, "user does not have enough balance of the required coin")
	ErrTokenNotSupported               = sdkerrors.Register(ModuleName, 19, "Token not supported by sifchain")
	ErrUnableToAddBalance              = sdkerrors.Register(ModuleName, 20, "unable to add balance")
	ErrNotEnoughLiquidity              = sdkerrors.Register(ModuleName, 21, "pool does not have sufficient balance")
	ErrPoolTooShallow                  = sdkerrors.Register(ModuleName, 23, "Cannot withdraw pool is too shallow")
	ErrOverFlow                        = sdkerrors.Register(ModuleName, 24, "IntegerOverflow")
	ErrUnableToAddLiquidity            = sdkerrors.Register(ModuleName, 25, "Unable to add liquidity")
	ErrUnableToRemoveLiquidity         = sdkerrors.Register(ModuleName, 26, "Unable to remove liquidity")
	ErrUnableToSwap                    = sdkerrors.Register(ModuleName, 27, "Unable to swap")
	ErrUnableToRemoveLiquidityProvider = sdkerrors.Register(ModuleName, 28, "Unable to add liquidity provider")
	ErrUnableToDecommissionPool        = sdkerrors.Register(ModuleName, 29, "Unable to decommission pool")
	ErrUnableToParseInt                = sdkerrors.Register(ModuleName, 30, "Unable to parse to Int")
	ErrReceivedAmountBelowExpected     = sdkerrors.Register(ModuleName, 31, "Unable to swap, received amount is below expected")
	ErrAmountTooLow                    = sdkerrors.Register(ModuleName, 32, "Tx amount is too low")
	ErrNotEnoughPermissions            = sdkerrors.Register(ModuleName, 33, "Signer does not have permissions to execute this action")
	ErrRateCannotBeNegative            = sdkerrors.Register(ModuleName, 34, "Rate cannot be negative")
	ErrCannotStartPolicy               = sdkerrors.Register(ModuleName, 35, "A new policy can be started only after the current policy has ended")

