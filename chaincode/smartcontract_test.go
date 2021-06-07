package chaincode_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-samples/asset-transfer-fabcar/chaincode-go/chaincode"
	"github.com/hyperledger/fabric-samples/asset-transfer-fabcar/chaincode-go/chaincode/mocks"
	"github.com/stretchr/testify/require"
)

type transactionContext interface {
	contractapi.TransactionContextInterface
}

type chaincodeStub interface {
	shim.ChaincodeStubInterface
}

func TestAddCar(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	assetTransfer := chaincode.SmartContract{}
	err := assetTransfer.AddCar(transactionContext, "CAR0", "KIA", "EV6", 1, "sanggi")
	require.NoError(t, err)

	chaincodeStub.PutStateReturns(fmt.Errorf("CAR0"))
	err = assetTransfer.AddCar(transactionContext, "CAR0", "KIA", "EV6", 3, "sanggi")
	require.EqualError(t, err, "Failed to put to world state. CAR0")
}

func TestQuerycar(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	expectedAsset := &chaincode.Car{ID: "Car1"}
	// expectedAsset1 := &chaincode.Car{ID: "Car2"}
	bytes, err := json.Marshal(expectedAsset)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	assetTransfer := chaincode.SmartContract{}
	asset, err := assetTransfer.QueryCar(transactionContext, "")
	require.NoError(t, err)
	fmt.Print(expectedAsset)
	fmt.Print(asset)
	require.Equal(t, expectedAsset, asset)
}
