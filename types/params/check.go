package params

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys"
	tokentypes "github.com/okex/okexchain-go-sdk/module/token/types"
)

const (
	tokenDescLenLimit       = 256
	countDefault            = 10
	perPageDefault          = 50
	perPageMax              = 200
	reWholeName             = `[a-zA-Z0-9[:space:]]{1,30}`
	bech32AddrLen           = 41
	ethAddrWithPrefixLen    = 42
	ethAddrWithoutPrefixLen = 40
)

var (
	reWhole = regexp.MustCompile(fmt.Sprintf(`^%s$`, reWholeName))
)

// CheckCreatePoolParams gives a quick validity check for the input params of creating pool in farm
func CheckCreatePoolParams(fromInfo keys.Info, passWd, poolName, minLockAmountStr, yieldToken string) error {
	if err := CheckPoolNameParams(fromInfo, passWd, poolName); err != nil {
		return err
	}

	if len(minLockAmountStr) == 0 || len(yieldToken) == 0 {
		return errors.New("failed. empty min lock token or empty yield token")
	}

	return nil
}

// CheckPoolNameParams gives a quick validity check for the input params of the pool name in farm
func CheckPoolNameParams(fromInfo keys.Info, passWd, poolName string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	if len(poolName) == 0 {
		return errors.New("failed. empty pool name")
	}

	return nil
}

// CheckProposalOperation gives a quick validity check for the input params of the proposal operation with proposal ID
func CheckProposalOperation(fromInfo keys.Info, passWd string, proposalID uint64) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	if proposalID <= 0 {
		return errors.New("failed. proposal ID must be positive")
	}

	return nil
}

// CheckTokenEditParams gives a quick validity check for the input params of token info editing
func CheckTokenEditParams(fromInfo keys.Info, passWd, symbol, description, wholeName string, isDescEdit, isWholeNameEdit bool) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	if len(symbol) == 0 {
		return errors.New("failed. empty symbol")
	}

	if isWholeNameEdit && !isWholeNameValid(wholeName) {
		return fmt.Errorf("failed. invalid whole name of token: %s", wholeName)
	}

	if isDescEdit && len(description) > tokenDescLenLimit {
		return errors.New("failed. invalid token description")
	}

	return nil
}

// CheckProductParams gives a quick validity check for the input product params
func CheckProductParams(fromInfo keys.Info, passWd, product string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	if len(product) == 0 {
		return errors.New("failed. empty product")
	}

	return nil
}

// CheckDexAssetsParams gives a quick validity check for the input params of dex assets
func CheckDexAssetsParams(fromInfo keys.Info, passWd, baseAsset, quoteAsset string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	if len(baseAsset) == 0 {
		return errors.New("failed. empty base asset")
	}

	if len(quoteAsset) == 0 {
		return errors.New("failed. empty quote asset")
	}

	return nil
}

// CheckQueryTokenInfoParams gives a quick validity check for the input params of query token info
func CheckQueryTokenInfoParams(ownerAddr, symbol string) error {
	if len(ownerAddr) == 0 && len(symbol) == 0 {
		return errors.New("failed. empty input")
	}

	return nil
}

// CheckTokenIssueParams gives a quick validity check for the input params of token issuing
func CheckTokenIssueParams(fromInfo keys.Info, passWd, orgSymbol, wholeName, tokenDesc string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	if len(orgSymbol) == 0 {
		return errors.New("failed. empty original symbol")
	}

	tokenDescLen := len(tokenDesc)
	if tokenDescLen == 0 || tokenDescLen > tokenDescLenLimit {
		return errors.New("failed. invalid token description")
	}

	if len(wholeName) == 0 {
		return errors.New("failed. empty whole name")
	}

	return nil
}

// CheckTransferUnitsParams gives a quick validity check for the input params of multi-send
func CheckTransferUnitsParams(fromInfo keys.Info, passWd string, transfers []tokentypes.TransferUnit) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}
	transLen := len(transfers)
	if transLen == 0 {
		return errors.New("failed. no receiver input")
	}
	for i := 0; i < transLen; i++ {
		if transfers[i].Coins.IsAllPositive() {
			continue
		} else {
			return errors.New("failed. only positive amount of coins is available")
		}
	}

	return nil
}

// CheckAddSharesParams gives a quick validity check for the input params of multi-voting
func CheckAddSharesParams(fromInfo keys.Info, passWd string, valAddrs []string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}
	if len(valAddrs) == 0 {
		return errors.New("failed. no validator address input")
	}

	// check duplicated
	filter := make(map[string]struct{}, len(valAddrs))
	for _, valAddr := range valAddrs {
		if _, ok := filter[valAddr]; ok {
			return fmt.Errorf("failed. validator address: %s is duplicated", valAddr)
		}
		filter[valAddr] = struct{}{}
	}

	return nil
}

// CheckKeyParams gives a basic validity check for the input key params
func CheckKeyParams(fromInfo keys.Info, passWd string) error {
	if fromInfo == nil {
		return errors.New("failed. input invalid keys info")
	}
	if len(passWd) == 0 {
		return errors.New("failed. no password input")
	}

	return nil
}

// CheckSendParams gives a quick validity check for the input params of transferring
func CheckSendParams(fromInfo keys.Info, passWd, toAddr string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	addrLen := len(toAddr)
	if !(addrLen == bech32AddrLen || addrLen == ethAddrWithoutPrefixLen || addrLen == ethAddrWithPrefixLen) {
		return errors.New("failed. invalid receiver address with incorrect length")
	}

	return nil
}

// CheckNewOrderParams gives a quick validity check for the input params for placing orders
func CheckNewOrderParams(fromInfo keys.Info, passWd string, products, sides, prices, quantities []string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	productsLen := len(products)
	if productsLen == 0 {
		return errors.New("failed. no product input")
	}

	if len(sides) != productsLen {
		return errors.New("failed. invalid param side counts")
	}

	if len(prices) != productsLen {
		return errors.New("failed. invalid param price counts")
	}

	if len(quantities) != productsLen {
		return errors.New("failed. invalid param quantity counts")
	}

	for _, side := range sides {
		if side != "BUY" && side != "SELL" {
			return errors.New(`failed. side must only be "BUY" or "SELL"`)
		}
	}

	return nil
}

// CheckCancelOrderParams gives a quick validity check for the input params for cancelling orders
func CheckCancelOrderParams(fromInfo keys.Info, passWd string, orderIDs []string) error {
	if err := CheckKeyParams(fromInfo, passWd); err != nil {
		return err
	}

	// check duplicated
	filter := make(map[string]struct{})
	for _, id := range orderIDs {
		if _, ok := filter[id]; ok {
			return fmt.Errorf("failed. duplicated orderID: %s", id)
		}

		filter[id] = struct{}{}
	}

	return nil
}

// CheckQueryOrderDetailParams gives a quick validity check for the input params of query order detail
func CheckQueryOrderDetailParams(orderID string) error {
	if len(orderID) == 0 {
		return errors.New("failed. empty order ID")
	}

	return nil
}

// CheckQueryTickersParams gives a quick validity check for the input params of query tickers
func CheckQueryTickersParams(count []int) (countRet int, err error) {
	if len(count) > 1 {
		return countRet, errors.New("failed. invalid params input for tickers query")
	}

	if len(count) == 0 {
		countRet = countDefault
	} else {
		if count[0] < 0 {
			return countRet, errors.New(`failed. "count" is negative`)
		}
		countRet = count[0]
	}
	return
}

// CheckQueryRecentTxRecordParams gives a quick validity check for the input params of query recent tx record
func CheckQueryRecentTxRecordParams(product string, start, end, page, perPage int) (perPageRet int, err error) {
	if len(product) == 0 {
		return perPageRet, errors.New("failed. empty product")
	}

	return checkParamsPaging(start, end, page, perPage)
}

// CheckQueryOrdersParams gives a quick validity check for the input params of query orders
func CheckQueryOrdersParams(addrStr, product, side string, start, end, page, perPage int) (perPageRet int, err error) {
	if err = IsValidAccAddr(addrStr); err != nil {
		return
	}

	if len(product) == 0 {
		return perPageRet, errors.New("failed. empty product")
	}

	if !isValidSide(side) {
		return perPageRet, errors.New(`failed. "side" must only be "BUY" or "SELL"`)

	}

	return checkParamsPaging(start, end, page, perPage)
}

// CheckQueryTransactionsParams gives a quick validity check for the input params of query transactions
func CheckQueryTransactionsParams(addrStr string, typeCode, start, end, page, perPage int) (perPageRet int, err error) {
	if err = IsValidAccAddr(addrStr); err != nil {
		return
	}

	if typeCode < 0 {
		return perPageRet, errors.New("failed. type code isn't allowed to be negative")

	}

	return checkParamsPaging(start, end, page, perPage)
}

// CheckQueryHeightParams gives a quick validity check for the input params of query tendermint data with height
func CheckQueryHeightParams(height int64) error {
	if height < 0 {
		return errors.New("failed. negative height is not available")
	}

	return nil
}

// IsValidAccAddr gives a quick validity check for an address string
func IsValidAccAddr(addrStr string) error {
	if len(addrStr) != bech32AddrLen || !strings.HasPrefix(addrStr, sdk.GetConfig().GetBech32AccountAddrPrefix()) {
		return fmt.Errorf("failed. invalid account address: %s", addrStr)
	}
	return nil
}

// CheckQueryTxResultParams gives a quick validity check for txs query by searching string
func CheckQueryTxResultParams(tmEventStrs []string, page, perPage int) error {
	if len(tmEventStrs) == 0 {
		return errors.New("failed. empty event to search")
	}

	if page <= 0 {
		return errors.New("failed. page must be greater than 0")
	}

	if perPage <= 0 {
		return errors.New("failed. limit number in a page must be greater than 0")
	}

	return nil
}

func checkParamsPaging(start, end, page, perPage int) (perPageRet int, err error) {
	if start < 0 || end < 0 || page < 0 || perPage < 0 {
		return perPageRet, errors.New(`failed. "start","end","page","perPage" must be positive`)
	}

	if start > end {
		return perPageRet, errors.New(`failed. "start" isn't allowed to be larger than "end"`)
	}

	if perPage == 0 {
		perPageRet = perPageDefault
	} else if perPage > perPageMax {
		perPageRet = perPageMax
	} else {
		perPageRet = perPage
	}
	return

}

func isValidSide(side string) bool {
	return side == "BUY" || side == "SELL"
}

func isWholeNameValid(wholeName string) bool {
	return reWhole.MatchString(wholeName)
}
