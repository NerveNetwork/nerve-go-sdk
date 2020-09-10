// @Title
// @Description
// @Author  Niels  2020/9/10
package acc

type Account interface {
	GetAddr() string
	GetAddrBytes() []byte
	GetPriKey() []byte
	GetPriKeyHex() string
	GetPubKey() []byte
	GetPubKeyHex() string
	GetChainId() uint16
	GetType() uint8
	GetPrefix() string
}

type NerveAccount struct {
	addr      string
	prefix    string
	chainId   uint16
	prikeyHex string
	pubkeyHex string
	accType   uint8
	pubkey    []byte
	prikey    []byte
	addrBytes []byte
}

func (a *NerveAccount) GetAddr() string {
	return a.addr
}

func (a *NerveAccount) GetAddrBytes() []byte {
	return a.addrBytes
}

func (a *NerveAccount) GetPriKey() []byte {
	return a.prikey
}

func (a *NerveAccount) GetPriKeyHex() string {
	return a.prikeyHex
}

func (a *NerveAccount) GetPubKey() []byte {
	return a.pubkey
}

func (a *NerveAccount) GetPubKeyHex() string {
	return a.pubkeyHex
}

func (a *NerveAccount) GetChainId() uint16 {
	return a.chainId
}

func (a *NerveAccount) GetType() uint8 {
	return a.accType
}

func (a *NerveAccount) GetPrefix() string {
	return a.prefix
}
