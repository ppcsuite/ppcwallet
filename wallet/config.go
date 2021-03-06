/*
 * Copyright (c) 2015 The btcsuite developers
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package wallet

import (
	"github.com/ppcsuite/ppcd/chaincfg"
	"github.com/ppcsuite/ppcwallet/waddrmgr"
	"github.com/ppcsuite/ppcwallet/walletdb"
	"github.com/ppcsuite/ppcwallet/wtxmgr"
)

// Config is a structure used to initialize a Wallet
// All values are required for successfully opening a Wallet
type Config struct {
	ChainParams *chaincfg.Params
	Db          *walletdb.DB
	TxStore     *wtxmgr.Store
	Waddrmgr    *waddrmgr.Manager
}
