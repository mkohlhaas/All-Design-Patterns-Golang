package main

import (
	"fmt"
	"log"
)

////////////////// Error Type ////////////////////

type facadeError struct {
	err error
}

func (f *facadeError) Unwrap() {
	if f.err != nil {
		log.Fatalf("Error: %s.\n", f.err.Error())
	}
}

//////////////////// Facade //////////////////////

type walletFacade struct {
	wallet       *wallet       // a wallet (the only thing the client cares about)
	account      *account      // needed to make life easier for client
	securityCode *securityCode // needed to make life easier for client
	notification *notification // needed to make life easier for client
	ledger       *ledger       // needed to make life easier for client
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Start creating account.")
	walletFacacde := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Account created.")
	fmt.Println("--------------------------------------------------------")
	return walletFacacde
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) *facadeError {
	fmt.Println("Start adding money to wallet.")
	err := w.account.checkAccount(accountID)
	if err.err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err.err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return &facadeError{nil}
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) *facadeError {
	fmt.Println("Start debiting money from wallet.")
	err := w.account.checkAccount(accountID)
	if err.err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err.err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err.err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "debit", amount)
	return &facadeError{nil}
}

///////////////////////// Account ///////////////

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

func (a *account) checkAccount(accountName string) *facadeError {
	if a.name != accountName {
		return &facadeError{fmt.Errorf("Account name is incorrect.")}
	}
	fmt.Println("Account verified.")
	return &facadeError{nil}
}

//////////////////// Security Code ///////////////

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkCode(incomingCode int) *facadeError {
	if s.code != incomingCode {
		return &facadeError{fmt.Errorf("Security code is incorrect.")}
	}
	fmt.Println("Security code verified.")
	return &facadeError{nil}
}

///////////////////// Wallet /////////////////////

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully.")
	return
}

func (w *wallet) debitBalance(amount int) *facadeError {
	if w.balance < amount {
		return &facadeError{fmt.Errorf("Balance is not sufficient.")}
	}
	fmt.Println("Wallet balance is sufficient.")
	w.balance = w.balance - amount
	return &facadeError{nil}
}

/////////////////// Ledger ///////////////////////

type ledger struct{}

func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for account id '%s' with txn type '%s' for amount %d.\n", accountID, txnType, amount)
	fmt.Println("--------------------------------------------------------")
	return
}

///////////////// Notification ///////////////////

type notification struct{}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification.")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification.")
}

//////////////// Main ////////////////////////////

func main() {
	walletFacade := newWalletFacade("abc", 1234)
	walletFacade.addMoneyToWallet("abc", 1234, 10).Unwrap()
	walletFacade.deductMoneyFromWallet("abc", 1234, 5).Unwrap()
}
