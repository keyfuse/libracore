// libracore
//
// Copyright 2019 by KeyFuse Labs
//
// GPLv3 License

package xcore

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/keyfuse/libracore/xproto"
)

// Client --
type Client struct {
	client           xproto.AdmissionControlClient
	network          Network
	grpcServerHost   string
	faucetServerHost string
}

// NewClient -- creates new Client with network.
func NewClient(network Network) (*Client, error) {
	client := &Client{}

	switch network {
	case TestNet:
		client.grpcServerHost = TestNetGRPCServerHost
		client.faucetServerHost = TestNetFaucetServerHost
	case MainNet:
		// (TODO)
	}
	client.network = network

	conn, err := grpc.Dial(client.grpcServerHost, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client.client = xproto.NewAdmissionControlClient(conn)
	return client, nil
}

// MintWithFaucetService -- mint coins from facuet server, this only works in testnet.
func (c *Client) MintWithFaucetService(address string, coins uint64) error {
	switch c.network {
	case TestNet:
		req := NewRequest()
		url := fmt.Sprintf("http://%s?amount=%d&address=%s", c.faucetServerHost, coins, address)
		rsp, err := req.Post(url, []byte{})
		if err != nil {
			return err
		}
		if rsp.StatusCode() != http.StatusOK {
			return fmt.Errorf("faucet.server[%v].rsp.status:%v", url, rsp.StatusCode())
		}
	case MainNet:
		return fmt.Errorf("mint.faucet.server.testnet.only")
	}
	return nil
}

// QueryAccountStates -- query the account states from the chain database.
func (c *Client) QueryAccountStates(addrs []string) ([]AccountState, error) {
	var err error
	var accountStates []AccountState
	var rsp *xproto.UpdateToLatestLedgerResponse

	// UpdateToLatestLedgerRequest.
	{
		req := &xproto.UpdateToLatestLedgerRequest{}
		for _, addr := range addrs {
			address, err := hex.DecodeString(addr)
			if err != nil {
				return nil, err
			}
			accountState := &xproto.GetAccountStateRequest{Address: address}
			accountStateReq := &xproto.RequestItem_GetAccountStateRequest{GetAccountStateRequest: accountState}
			requestItem := &xproto.RequestItem{RequestedItems: accountStateReq}
			req.RequestedItems = append(req.RequestedItems, requestItem)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		if rsp, err = c.client.UpdateToLatestLedger(ctx, req, grpc.WaitForReady(true)); err != nil {
			return nil, err
		}
	}

	// AccountStates.
	{
		for _, item := range rsp.ResponseItems {
			// Default seq is 1.
			accountState := &AccountState{SequenceNumber: 1}
			itm := item.ResponseItems.(*xproto.ResponseItem_GetAccountStateResponse)
			if itm.GetAccountStateResponse.AccountStateWithProof.Blob != nil {
				accountResource := &AccountResource{}
				if err := accountResource.Deserialize(itm.GetAccountStateResponse.AccountStateWithProof.Blob.Blob); err != nil {
					return nil, err
				}
				if err := accountState.Deserialize(accountResource.State[AccountStatePath]); err != nil {
					return nil, err
				}
			}
			accountStates = append(accountStates, *accountState)
		}
	}
	return accountStates, nil
}

// QueryAccountSequenceNumber -- query the account SequenceNumber of this address.
func (c *Client) QueryAccountSequenceNumber(addr string) (uint64, error) {
	states, err := c.QueryAccountStates([]string{addr})
	if err != nil {
		return 0, err
	}
	state := states[0]
	return state.SequenceNumber, nil
}

// WaitForTransaction -- block the request and wait the sequence until timeout.
func (c *Client) WaitForTransaction(addr string, sequenceNumber uint64) error {
	maxtrys := 50
	for {
		maxtrys--
		sequence, _ := c.QueryAccountSequenceNumber(addr)
		if sequence >= sequenceNumber {
			break
		}
		if maxtrys <= 0 {
			return fmt.Errorf("wait.for.transaction.timeout")
		}
		time.Sleep(time.Millisecond * 100)
	}
	return nil
}

// SubmitTransaction -- submit the txn req  to Libra chain.
func (c *Client) SubmitTransaction(req *xproto.SubmitTransactionRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	resp, err := c.client.SubmitTransaction(ctx, req, grpc.WaitForReady(true))
	if err != nil {
		return err
	}
	if acStatus := resp.GetAcStatus(); acStatus.Code != xproto.AdmissionControlStatusCode_Accepted {
		return fmt.Errorf("transfer.transaction.failed:%+v", acStatus)
	}
	return nil
}

// TransferCoins -- transfer coins with block.
func (c *Client) TransferCoins(sender *Account, to string, amount uint64, gasUnitPrice uint64, maxGasAmount uint64) error {
	return c.transferCoins(sender, to, amount, gasUnitPrice, maxGasAmount, true)
}

// TransferCoinsAsync -- transfer the coins with non-block.
func (c *Client) TransferCoinsAsync(sender *Account, to string, amount uint64, gasUnitPrice uint64, maxGasAmount uint64) error {
	return c.transferCoins(sender, to, amount, gasUnitPrice, maxGasAmount, false)
}

func (c *Client) transferCoins(sender *Account, to string, amount uint64, gasUnitPrice uint64, maxGasAmount uint64, block bool) error {
	sequence, err := c.QueryAccountSequenceNumber(sender.Address.ToString())
	if err != nil {
		return err
	}

	tx, err := NewTransaction().
		From(sender.Address).
		AddKey(sender.KeyPair.Private).
		To(to, amount).
		SetSequence(sequence).
		SetGasUintPrice(gasUnitPrice).
		SetMaxGasAmount(maxGasAmount).
		Sign().
		Build()
	if err != nil {
		return err
	}

	signedTxn := &xproto.SignedTransaction{
		SenderPublicKey: sender.KeyPair.Public.Value.Serialize(),
		RawTxnBytes:     tx.GetBytes(),
		SenderSignature: tx.GetSignature(),
	}
	req := &xproto.SubmitTransactionRequest{
		SignedTxn: signedTxn,
	}

	if err := c.SubmitTransaction(req); err != nil {
		return err
	}

	if block {
		return c.WaitForTransaction(sender.Address.ToString(), sequence+1)
	}
	return nil
}
