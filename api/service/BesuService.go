package service

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/Mpetrato/goledger-challenge-besu/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const contractABI = `[{"inputs":[],"name":"get","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"set","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

type BesuService struct{}

func NewBesuService() *BesuService {
	return &BesuService{}
}

func (s *BesuService) SetBesuContractValue(value uint64) error {
	abi, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	BESU_NODE_URL, err := helpers.GetOSEnv("BESU_NODE_URL")
	if err != nil {
		log.Fatal("BESU_NODE_URL not found o env")
	}

	client, err := ethclient.DialContext(ctx, *BESU_NODE_URL)
	if err != nil {
		log.Fatalf("error dialing node: %v", err)
	}

	slog.Info("querying chain id")

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("error querying chain id: %v", err)
	}
	defer client.Close()

	CONTRACT_ADDRESS, err := helpers.GetOSEnv("CONTRACT_ADDRESS")
	if err != nil {
		log.Fatal("CONTRACT_ADDRESS not found o env")
	}

	contractAddress := common.HexToAddress(*CONTRACT_ADDRESS)

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	PRIVATE_KEY, err := helpers.GetOSEnv("PRIVATE_KEY")
	if err != nil {
		log.Fatal("PRIVATE_KEY not found o env")
	}

	priv, err := crypto.HexToECDSA(*PRIVATE_KEY) // this can be found in the genesis.json file
	if err != nil {
		log.Fatalf("error loading private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(priv, chainId)
	if err != nil {
		log.Fatalf("error creating transactor: %v", err)
	}

	formattedValue := big.NewInt(int64(value))

	tx, err := boundContract.Transact(auth, "set", formattedValue)
	if err != nil {
		log.Fatalf("error transacting: %v", err)
	}

	fmt.Println("waiting until transaction is mined",
		"tx", tx.Hash().Hex(),
	)

	receipt, err := bind.WaitMined(
		context.Background(),
		client,
		tx,
	)
	if err != nil {
		log.Fatalf("error waiting for transaction to be mined: %v", err)
	}

	fmt.Printf("transaction mined: %v\n", receipt)

	return nil
}

func (s *BesuService) GetBesuContractValue() (*uint64, error) {
	abi, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatalf("error parsing abi: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	BESU_NODE_URL, err := helpers.GetOSEnv("BESU_NODE_URL")
	if err != nil {
		log.Fatal("BESU_NODE_URL not found o env")
	}

	client, err := ethclient.DialContext(ctx, *BESU_NODE_URL)
	if err != nil {
		log.Fatalf("error connecting to eth client: %v", err)
	}
	defer client.Close()

	CONTRACT_ADDRESS := os.Getenv("CONTRACT_ADDRESS")
	if CONTRACT_ADDRESS == "" {
		log.Fatal("CONTRACT_ADDRESS not found o env")
	}
	contractAddress := common.HexToAddress(CONTRACT_ADDRESS) // will be returned during startDev.sh execution
	caller := bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	var output []interface{}
	err = boundContract.Call(&caller, &output, "get")
	if err != nil {
		log.Fatalf("error calling contract: %v", err)
	}

	if len(output) == 0 {
		return nil, fmt.Errorf("empty result")
	}

	contractStoreValue := output[0].(*big.Int).Uint64()

	fmt.Println("Successfully called contract!", contractStoreValue)
	return &contractStoreValue, nil
}
