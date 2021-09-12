package main

import (
	"testing"
)

func TestGenerateBasePoint0(t *testing.T) {
	basePoint := generateBasePoint(0)
	if basePoint.X.String() != "10457101036533406547632367118273992217979173478358440826365724437999023779287" ||
		basePoint.Y.String() != "19824078218392094440610104313265183977899662750282163392862422243483260492317" {
		t.Error("Basepoint incorrect")
	}
}

func TestGenerateBasePoint100(t *testing.T) {
	basePoint := generateBasePoint(100)
	if basePoint.X.String() != "19833279235750316492042296719910504725662169777030633907799544770859664074425" ||
		basePoint.Y.String() != "19728181840733450018645709815558234105947496435492456827296453808697087111057" {
		t.Error("Basepoint incorrect")
	}
}

func TestPedersenHashString(t *testing.T) {
	hash := pedersenHash([]byte("pedersen hashes yay!"))
	if hash.X.String() != "17128280986227683498588103109559827401834954064908010612959680331684132386223" ||
		hash.Y.String() != "740472655752945115283384209702187584259772600939885912301093536259195197955" {
		t.Error("Hashes did not match")
	}
}

func TestPedersenHashSingleByte(t *testing.T) {
	hash := pedersenHash([]byte{0})
	if hash.X.String() != "2713984616998054873485125083403724179682140658671583177610038376665425019990" ||
		hash.Y.String() != "6281144028007049357012765257133378775433463448755543459194783914343308083779" {
		t.Error("Hashes did not match")
	}
}

func TestPedersenHashEmpty(t *testing.T) {
	hash := pedersenHash([]byte{})
	if hash.X.String() != "0" ||
		hash.Y.String() != "1" {
		t.Error("Hashes did not match")
	}
}
