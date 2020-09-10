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
	chainId   string
	prikeyHex string
	pubkeyHex string
}

func (a *NerveAccount) GetAddr() string {
	return a.addr
}
