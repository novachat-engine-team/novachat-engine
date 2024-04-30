package crypto

import (
	"bytes"
	"math/big"
	"math/rand"
	"time"
)

type PasswordKdfAlgoModPow struct {
	Salt1 []byte
	Salt2 []byte
	G     int32
	P     []byte
}

type SRPUtil struct {
	*PasswordKdfAlgoModPow
	g      *big.Int
	gBytes []byte
	p      *big.Int
	k      *big.Int
}

func MakeSRPUtil(algo *PasswordKdfAlgoModPow) *SRPUtil {
	rand.Seed(time.Now().UnixNano())

	srp := &SRPUtil{
		PasswordKdfAlgoModPow: algo,
		g:                     big.NewInt(int64(algo.G)),
		p:                     new(big.Int).SetBytes(algo.P),
	}

	srp.gBytes = getBigIntegerBytes(srp.g)
	kBytes := calcSHA256(algo.P, srp.gBytes)
	srp.k = new(big.Int).SetBytes(kBytes)

	return srp
}

func (m *SRPUtil) CheckNewSalt1(newSalt1 []byte) bool {
	if len(newSalt1) < 8 {
		return false
	}

	return bytes.Equal(m.Salt1, newSalt1[:8])
}

func (m *SRPUtil) CalcSRPB(vBytes []byte) ([]byte, []byte) {
	v := new(big.Int).SetBytes(vBytes)

	bNonce := RandomBytes(256)
	b := new(big.Int).SetBytes(bNonce)
	B := new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p)), m.p)
	BBytes := getBigIntegerBytes(B)

	return bNonce, BBytes
}

func (m *SRPUtil) CalcSRPB2(bNonce, vBytes []byte) []byte {
	v := new(big.Int).SetBytes(vBytes)

	b := new(big.Int).SetBytes(bNonce)
	B := new(big.Int).Mod(new(big.Int).Add(new(big.Int).Mul(m.k, v), new(big.Int).Exp(m.g, b, m.p)), m.p)
	BBytes := getBigIntegerBytes(B)

	return BBytes
}

func (m *SRPUtil) CalcM(newSalt1, vBytes, srpA, srpb, srpB []byte) []byte {
	v := new(big.Int).SetBytes(vBytes)

	A := new(big.Int).SetBytes(srpA)
	if A.Cmp(bigIntZero) <= 0 || A.Cmp(m.p) >= 0 {
		return nil
	}
	ABytes := getBigIntegerBytes(A)

	b := new(big.Int).SetBytes(srpb)

	uBytes := calcSHA256(ABytes, srpB)
	u := new(big.Int).SetBytes(uBytes)
	if u.Cmp(bigIntZero) == 0 {
		return nil
	}
	S := new(big.Int).Exp(new(big.Int).Mod(new(big.Int).Mul(A, new(big.Int).Exp(v, u, m.p)), m.p), b, m.p)
	SBytes := getBigIntegerBytes(S)

	KBytes := calcSHA256(SBytes)

	pHash := calcSHA256(m.P)
	gHash := calcSHA256(m.gBytes)
	for i := 0; i < len(pHash); i++ {
		pHash[i] = gHash[i] ^ pHash[i]
	}

	return calcSHA256(pHash, calcSHA256(newSalt1), calcSHA256(m.Salt2), ABytes, srpB, KBytes)
}

func (m *SRPUtil) GetX(newSalt1, passwordBytes []byte) []byte {
	var xBytes []byte

	xBytes = calcSHA256(newSalt1, passwordBytes, newSalt1)
	xBytes = calcSHA256(m.Salt2, xBytes, m.Salt2)
	xBytes = calcPBKDF2(xBytes, newSalt1)

	return calcSHA256(m.Salt2, xBytes, m.Salt2)
}

func (m *SRPUtil) GetV(newSalt1, passwordBytes []byte) *big.Int {
	xBytes := m.GetX(newSalt1, passwordBytes)
	x := new(big.Int).SetBytes(xBytes)

	return new(big.Int).Exp(m.g, x, m.p)
}

func (m *SRPUtil) GetVBytes(newSalt1, passwordBytes []byte) []byte {
	return getBigIntegerBytes(m.GetV(newSalt1, passwordBytes))
}

func (m *SRPUtil) CalcClientM(newSalt1, xBytes, srpB []byte) ([]byte, []byte) {
	if len(xBytes) == 0 || len(srpB) == 0 || !isGoodPrime(m.P, int(m.G)) {
		// fmt.Println("check error")
		return nil, nil
	}

	x := new(big.Int).SetBytes(xBytes)

	aBytes := RandomBytes(256)
	//fmt.Println(hex.EncodeToString(aBytes))
	a := new(big.Int).SetBytes(aBytes)

	A := new(big.Int).Exp(m.g, a, m.p)
	ABytes := getBigIntegerBytes(A)

	B := new(big.Int).SetBytes(srpB)
	if B.Cmp(bigIntZero) <= 0 || B.Cmp(m.p) >= 0 {
		// fmt.Println("b error")
		return nil, nil
	}
	BBytes := getBigIntegerBytes(B)

	uBytes := calcSHA256(ABytes, BBytes)
	u := new(big.Int).SetBytes(uBytes)
	if u.Cmp(bigIntZero) == 0 {
		// fmt.Println("u error")
		return nil, nil
	}

	BKgx := new(big.Int).Sub(B, new(big.Int).Mod(new(big.Int).Mul(m.k, new(big.Int).Exp(m.g, x, m.p)), m.p))
	if BKgx.Cmp(bigIntZero) < 0 {
		// fmt.Println("<0")
		BKgx = new(big.Int).Add(BKgx, m.p)
	}

	if !isGoodGaAndGb(BKgx, m.p) {
		// fmt.Println("isGoodGaAndGb error")
		return nil, nil
	}

	S := new(big.Int).Exp(BKgx, new(big.Int).Add(a, new(big.Int).Mul(u, x)), m.p)
	SBytes := getBigIntegerBytes(S)

	KBytes := calcSHA256(SBytes)

	pHash := calcSHA256(m.P)
	gHash := calcSHA256(m.gBytes)
	for i := 0; i < len(pHash); i++ {
		pHash[i] = gHash[i] ^ pHash[i]
	}

	return ABytes, calcSHA256(pHash, calcSHA256(newSalt1), calcSHA256(m.Salt2), ABytes, BBytes, KBytes)
}

func (m *SRPUtil) CalcClientM2(newSalt1, aBytes, ABytes, xBytes, srpB []byte) []byte {
	if len(xBytes) == 0 || len(srpB) == 0 || !isGoodPrime(m.P, int(m.G)) {
		// fmt.Println("check error")
		return nil
	}

	x := new(big.Int).SetBytes(xBytes)

	a := new(big.Int).SetBytes(aBytes)

	B := new(big.Int).SetBytes(srpB)
	if B.Cmp(bigIntZero) <= 0 || B.Cmp(m.p) >= 0 {
		// fmt.Println("b error")
		return nil
	}
	BBytes := getBigIntegerBytes(B)

	uBytes := calcSHA256(ABytes, BBytes)
	u := new(big.Int).SetBytes(uBytes)
	if u.Cmp(bigIntZero) == 0 {
		// fmt.Println("u error")
		return nil
	}

	BKgx := new(big.Int).Sub(B, new(big.Int).Mod(new(big.Int).Mul(m.k, new(big.Int).Exp(m.g, x, m.p)), m.p))
	if BKgx.Cmp(bigIntZero) < 0 {
		BKgx = new(big.Int).Add(BKgx, m.p)
	}

	if !isGoodGaAndGb(BKgx, m.p) {
		// fmt.Println("isGoodGaAndGb error")
		return nil
	}

	S := new(big.Int).Exp(BKgx, new(big.Int).Add(a, new(big.Int).Mul(u, x)), m.p)
	SBytes := getBigIntegerBytes(S)

	KBytes := calcSHA256(SBytes)

	pHash := calcSHA256(m.P)
	gHash := calcSHA256(m.gBytes)

	for i := 0; i < len(pHash); i++ {
		pHash[i] = gHash[i] ^ pHash[i]
	}

	return calcSHA256(pHash, calcSHA256(newSalt1), calcSHA256(m.Salt2), ABytes, BBytes, KBytes)
}
