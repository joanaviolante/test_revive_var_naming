package cache

import (
	"database/sql"
	"mp-shipping-charge-service/model"
	"mp-shipping-charge-service/util"

	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetRegionServiceAreaAndShipChargeCache(t *testing.T) {
	logr = zap.NewNop()
	var cartId = "CART-001"
	var serviceAreaId int64 = 1
	var shippingHeaderId int64 = 0
	var shippingPercent decimal.Decimal = decimal.NewFromInt(6)
	var serviceAreaName = ""
	var shipGroupId int64 = 4
	var partnerLocationId int64 = 1s
	var serviceAreaType = "REGION"
	var shipOptionCode = "STANDARD"
	var isResidentialAddress = false
	var shipChargeType = "PERCENT"
	var serviceAreaFound = true
	var endTs time.Time = time.Now()
	var err error = nil
	var currentTimestamp time.Time = time.Now()
	var found bool = true
	var serviceAreaFoundResult bool

	var dbResults []model.RegionServiceAreaAndShipCharge
	var dbResult model.RegionServiceAreaAndShipCharge
	var cacheResult []model.ServiceAreaAndShipChargeCache

	dbResults = append(dbResults, dbResult)
	dbResults[0].ServiceAreaName = serviceAreaName
	dbResults[0].ShipGroupId = shipGroupId
	dbResults[0].PartnerLocationId = partnerLocationId
	dbResults[0].ServiceAreaType = serviceAreaType
	dbResults[0].ShipOptionCode = shipOptionCode
	dbResults[0].ServiceAreaFound = serviceAreaFound
	dbResults[0].ShippingRegionHeaderId = shippingHeaderId
	dbResults[0].ServiceAreadId = serviceAreaId
	dbResults[0].ShipPercentCommercial = shippingPercent
	dbResults[0].EndTs = endTs

	var cacheDumpResult model.RegionServiceAreaAndShipChargeCacheDump

	getRegionServiceAreasAndShipChargeDBFunc = func(
		dbConn *sql.DB,
		stmtTimeoutSecs int,
		cartId string,
		currentTs time.Time,
		serviceAreaName string,
		shipGroupId int64,
		partnerLocationId int64,
		serviceAreaType string,
		shipOptionCode string,
		isResidentialAddress bool,
		shipChargeType string) ([]model.RegionServiceAreaAndShipCharge, bool, error) {

		return dbResults, found, err
	}

	var props util.PropsType
	props.Cache.LocationCleanupSecs = 10
	props.Cache.LocationExpireSecs = 300
	props.Cache.DumpCacheMaxItems = 10

	cacheDumpResult = DumpRegionServiceAreaAndShipChargeCache(props)
	assert.False(t, cacheDumpResult.Enabled)

	InitRegionServiceAreaAndShipChargeCache(props)

	// The first call fetches from DB
	cacheResult, serviceAreaFoundResult, err = GetRegionServiceAreaAndShipChargeCache(cartId, currentTimestamp, serviceAreaName, shipGroupId, partnerLocationId, serviceAreaType, shipOptionCode, isResidentialAddress, shipChargeType, props)

	assert.Nil(t, err)
	assert.Equal(t, serviceAreaType, cacheResult[0].ServiceAreaType)
	assert.Equal(t, serviceAreaName, cacheResult[0].ServiceAreaName)
	assert.Equal(t, shippingHeaderId, cacheResult[0].ShippingRegionHeaderId)
	assert.Equal(t, serviceAreaId, cacheResult[0].ServiceAreadId)
	assert.Equal(t, shippingPercent, cacheResult[0].ShipPercentCommercial)
	assert.Equal(t, serviceAreaFoundResult, cacheResult[0].ServiceAreaFound)
	assert.True(t, cacheResult[0].ServiceAreaFound)

	// The second call gets same result as first call, but fetches from cache
	cacheResult, serviceAreaFoundResult, err = GetRegionServiceAreaAndShipChargeCache(cartId, currentTimestamp, serviceAreaName, shipGroupId, partnerLocationId, serviceAreaType, shipOptionCode, isResidentialAddress, shipChargeType, props)

	assert.Nil(t, err)
	assert.Equal(t, serviceAreaType, cacheResult[0].ServiceAreaType)
	assert.Equal(t, serviceAreaName, cacheResult[0].ServiceAreaName)
	assert.Equal(t, shippingHeaderId, cacheResult[0].ShippingRegionHeaderId)
	assert.Equal(t, serviceAreaId, cacheResult[0].ServiceAreadId)
	assert.Equal(t, shippingPercent, cacheResult[0].ShipPercentCommercial)
	assert.Equal(t, serviceAreaFoundResult, cacheResult[0].ServiceAreaFound)
	assert.True(t, cacheResult[0].ServiceAreaFound)

	cacheDumpResult = DumpRegionServiceAreaAndShipChargeCache(props)
	assert.True(t, cacheDumpResult.Enabled)
}