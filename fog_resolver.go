package api

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"errors"

	"github.com/ChainSafe/go-schnorrkel"
	account "github.com/jadeydi/mobilecoin-account"
	"github.com/jadeydi/mobilecoin-account/block"
)

const (
	SUPER_CONTEXT = "Fog authority signature"
)

type FullyValidatedFogPubkey struct {
}

// https://github.com/mobilecoinfoundation/mobilecoin/blob/2f90154a445c769594dfad881463a2d4a003d7d6/fog/report/validation/src/ingest_report.rs#L23
type IngestReportVerifier struct {
	Verifier *Verifier
}

type FogResolver struct {
	Responses map[string]*block.ReportResponse
	Verifier  *IngestReportVerifier
}

// https://github.com/mobilecoinfoundation/mobilecoin/blob/2f90154a445c769594dfad881463a2d4a003d7d6/account-keys/src/account_keys.rs#L180
// https://github.com/mobilecoinfoundation/mobilecoin/blob/2f90154a445c769594dfad881463a2d4a003d7d6/fog/sig/src/public_address.rs#L44
func verifyAuthority(recipient *account.PublicAddress, certs []*x509.Certificate, sig string) (bool, error) {
	cert, err := VerifiedRoot(certs)
	if err != nil {
		return false, err
	}
	/*
		pub, err := x509.ParsePKIXPublicKey(cert.RawSubjectPublicKeyInfo)
		if err != nil {
			return err
		}
	*/

	signingCtx := []byte(SUPER_CONTEXT)
	verifyTranscript := schnorrkel.NewSigningContext(signingCtx, cert.RawSubjectPublicKeyInfo)

	view, err := hex.DecodeString(recipient.ViewPublicKey)
	if err != nil {
		return false, err
	}
	var view32 [32]byte
	copy(view32[:], view)
	public := schnorrkel.NewPublicKey(view32)
	sigbuf, err := hex.DecodeString(sig)
	if err != nil {
		return false, err
	}
	var sig64 [64]byte
	copy(sig64[:], sigbuf)
	signature := schnorrkel.Signature{}
	err = signature.Decode(sig64)
	if err != nil {
		return false, err
	}
	return public.Verify(&signature, verifyTranscript), nil
}

func mcPublicKey(cert *x509.Certificate) (ed25519.PublicKey, error) {
	pub, err := x509.ParsePKIXPublicKey(cert.RawSubjectPublicKeyInfo)
	if err != nil {
		return nil, err
	}
	switch pub := pub.(type) {
	case ed25519.PublicKey:
		return pub, nil
	default:
		return nil, errors.New("unknown type of public key")
	}
}

// https://github.com/mobilecoinfoundation/mobilecoin/blob/2f90154a445c769594dfad881463a2d4a003d7d6/fog/sig/src/public_address.rs#L22
func verifyFogSig(recipient *account.PublicAddress, responses *block.ReportResponse) error {
	var certs []*x509.Certificate
	for _, buf := range responses.GetChain() {
		cert, err := x509.ParseCertificate(buf)
		if err != nil {
			return err
		}
		certs = append(certs, cert)
	}

	if len(certs) == 0 {
		return errors.New("Empty Chain Error")
	}

	valid, err := verifyAuthority(recipient, certs, recipient.FogAuthoritySig)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("Verify Authority Error")
	}

	// leaf
	leaf := certs[0]
	public, err := mcPublicKey(leaf)
	if err != nil {
		return err
	}
	return VerifyReports(public, responses.GetReports(), responses.GetSignature())
}

// https://github.com/mobilecoinfoundation/mobilecoin/blob/2f90154a445c769594dfad881463a2d4a003d7d6/fog/report/validation/src/lib.rs#L108
// get_fog_pubkey
func (resolver *FogResolver) GetFogPubkey(recipient *account.PublicAddress) error {
	response := resolver.Responses[recipient.FogReportUrl]
	if response == nil {
		return errors.New("No Matching Report Response")
	}

	err := verifyFogSig(recipient, response)
	if err != nil {
		return err
	}
	for _, report := range response.Reports {
		if recipient.FogReportId == report.GetFogReportId() {
		}
	}
	return errors.New("No Matching Report Id")
}

// https://github.com/mobilecoinfoundation/mobilecoin/blob/2f90154a445c769594dfad881463a2d4a003d7d6/fog/report/validation/src/ingest_report.rs#L23
// validate_ingest_ias_report
func (resolver *FogResolver) ValidateIngestIasReport(report *block.VerificationReport) {
}
