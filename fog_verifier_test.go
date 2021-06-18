package api

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	_ "embed"
	"encoding/hex"
	"encoding/pem"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed credentials/Dev_AttestationReportSigningCACert.pem
var certPEM []byte

func TestFogVerifier(t *testing.T) {
	assert := assert.New(t)

	block, _ := pem.Decode(certPEM)
	cert, err := x509.ParseCertificate(block.Bytes)
	assert.Nil(err)
	_ = cert
	//log.Printf("%#v", cert)

	derBytes := []byte{48, 130, 4, 161, 48, 130, 3, 9, 160, 3, 2, 1, 2, 2, 9, 0, 209, 7, 118, 93, 50, 163, 176, 150, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 48, 126, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 11, 48, 9, 6, 3, 85, 4, 8, 12, 2, 67, 65, 49, 20, 48, 18, 6, 3, 85, 4, 7, 12, 11, 83, 97, 110, 116, 97, 32, 67, 108, 97, 114, 97, 49, 26, 48, 24, 6, 3, 85, 4, 10, 12, 17, 73, 110, 116, 101, 108, 32, 67, 111, 114, 112, 111, 114, 97, 116, 105, 111, 110, 49, 48, 48, 46, 6, 3, 85, 4, 3, 12, 39, 73, 110, 116, 101, 108, 32, 83, 71, 88, 32, 65, 116, 116, 101, 115, 116, 97, 116, 105, 111, 110, 32, 82, 101, 112, 111, 114, 116, 32, 83, 105, 103, 110, 105, 110, 103, 32, 67, 65, 48, 30, 23, 13, 49, 54, 49, 49, 50, 50, 48, 57, 51, 54, 53, 56, 90, 23, 13, 50, 54, 49, 49, 50, 48, 48, 57, 51, 54, 53, 56, 90, 48, 123, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 11, 48, 9, 6, 3, 85, 4, 8, 12, 2, 67, 65, 49, 20, 48, 18, 6, 3, 85, 4, 7, 12, 11, 83, 97, 110, 116, 97, 32, 67, 108, 97, 114, 97, 49, 26, 48, 24, 6, 3, 85, 4, 10, 12, 17, 73, 110, 116, 101, 108, 32, 67, 111, 114, 112, 111, 114, 97, 116, 105, 111, 110, 49, 45, 48, 43, 6, 3, 85, 4, 3, 12, 36, 73, 110, 116, 101, 108, 32, 83, 71, 88, 32, 65, 116, 116, 101, 115, 116, 97, 116, 105, 111, 110, 32, 82, 101, 112, 111, 114, 116, 32, 83, 105, 103, 110, 105, 110, 103, 48, 130, 1, 34, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 1, 5, 0, 3, 130, 1, 15, 0, 48, 130, 1, 10, 2, 130, 1, 1, 0, 169, 122, 45, 224, 230, 110, 166, 20, 124, 158, 231, 69, 172, 1, 98, 104, 108, 113, 146, 9, 154, 252, 75, 63, 4, 15, 173, 109, 224, 147, 81, 29, 116, 232, 2, 245, 16, 215, 22, 3, 129, 87, 220, 175, 132, 244, 16, 75, 211, 254, 215, 230, 184, 249, 156, 136, 23, 253, 31, 245, 185, 184, 100, 41, 108, 61, 129, 250, 143, 27, 114, 158, 2, 210, 29, 114, 255, 238, 76, 237, 114, 94, 254, 116, 190, 166, 143, 188, 77, 66, 68, 40, 111, 205, 212, 191, 100, 64, 106, 67, 154, 21, 188, 180, 207, 103, 117, 68, 137, 196, 35, 151, 43, 74, 128, 223, 92, 46, 124, 91, 194, 219, 175, 45, 66, 187, 123, 36, 79, 124, 149, 191, 146, 199, 93, 59, 51, 252, 84, 16, 103, 138, 137, 88, 157, 16, 131, 218, 58, 204, 69, 159, 39, 4, 205, 153, 89, 140, 39, 94, 124, 24, 120, 224, 7, 87, 229, 189, 180, 232, 64, 34, 108, 17, 192, 161, 127, 247, 156, 128, 177, 92, 29, 219, 90, 242, 28, 194, 65, 112, 97, 251, 210, 162, 218, 129, 158, 211, 183, 43, 126, 250, 163, 191, 235, 226, 128, 92, 155, 138, 193, 154, 163, 70, 81, 45, 72, 76, 252, 129, 148, 30, 21, 245, 88, 129, 204, 18, 126, 143, 122, 161, 35, 0, 205, 90, 251, 87, 66, 250, 29, 32, 203, 70, 122, 91, 235, 28, 102, 108, 247, 106, 54, 137, 120, 181, 2, 3, 1, 0, 1, 163, 129, 164, 48, 129, 161, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 120, 67, 123, 118, 166, 126, 188, 208, 175, 126, 66, 55, 235, 53, 124, 59, 135, 1, 81, 60, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 6, 192, 48, 12, 6, 3, 85, 29, 19, 1, 1, 255, 4, 2, 48, 0, 48, 96, 6, 3, 85, 29, 31, 4, 89, 48, 87, 48, 85, 160, 83, 160, 81, 134, 79, 104, 116, 116, 112, 58, 47, 47, 116, 114, 117, 115, 116, 101, 100, 115, 101, 114, 118, 105, 99, 101, 115, 46, 105, 110, 116, 101, 108, 46, 99, 111, 109, 47, 99, 111, 110, 116, 101, 110, 116, 47, 67, 82, 76, 47, 83, 71, 88, 47, 65, 116, 116, 101, 115, 116, 97, 116, 105, 111, 110, 82, 101, 112, 111, 114, 116, 83, 105, 103, 110, 105, 110, 103, 67, 65, 46, 99, 114, 108, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 3, 130, 1, 129, 0, 103, 8, 182, 27, 92, 43, 210, 21, 71, 62, 43, 70, 175, 153, 40, 79, 187, 147, 157, 63, 59, 21, 44, 153, 111, 26, 106, 243, 179, 41, 189, 34, 11, 29, 59, 97, 15, 107, 206, 46, 103, 83, 189, 237, 48, 77, 178, 25, 18, 243, 133, 37, 98, 22, 207, 203, 164, 86, 189, 150, 148, 11, 232, 146, 245, 105, 12, 38, 13, 30, 248, 79, 22, 6, 4, 2, 34, 229, 254, 8, 229, 50, 104, 8, 33, 42, 68, 124, 253, 214, 74, 70, 233, 75, 242, 159, 107, 75, 154, 114, 29, 37, 179, 196, 226, 246, 47, 88, 186, 237, 93, 119, 197, 5, 36, 143, 15, 128, 31, 159, 191, 183, 253, 117, 32, 128, 9, 92, 238, 128, 147, 139, 51, 159, 109, 187, 78, 22, 86, 0, 226, 14, 74, 113, 136, 18, 212, 157, 153, 1, 227, 16, 169, 181, 29, 102, 199, 153, 9, 198, 153, 101, 153, 250, 230, 215, 106, 121, 239, 20, 93, 153, 67, 191, 29, 62, 53, 211, 180, 45, 31, 185, 164, 92, 190, 142, 227, 52, 193, 102, 238, 231, 211, 47, 205, 201, 147, 93, 184, 236, 139, 177, 216, 235, 55, 121, 221, 138, 185, 43, 110, 56, 127, 1, 71, 69, 15, 30, 56, 29, 8, 88, 31, 184, 61, 243, 59, 21, 224, 0, 165, 155, 229, 126, 169, 74, 58, 82, 220, 100, 189, 174, 201, 89, 179, 70, 76, 145, 231, 37, 187, 218, 234, 61, 153, 232, 87, 227, 128, 162, 60, 157, 159, 177, 239, 88, 233, 228, 45, 113, 241, 33, 48, 249, 38, 29, 114, 52, 214, 195, 126, 43, 3, 219, 164, 13, 253, 251, 19, 172, 74, 216, 225, 63, 211, 117, 99, 86, 182, 181, 0, 21, 163, 236, 149, 128, 184, 21, 216, 124, 44, 239, 113, 92, 210, 141, 240, 11, 191, 42, 60, 64, 62, 191, 102, 145, 179, 240, 94, 221, 145, 67, 128, 60, 160, 133, 207, 245, 126, 5, 62, 236, 47, 143, 234, 70, 234, 119, 138, 104, 201, 190, 136, 91, 194, 130, 37, 188, 95, 48, 155, 228, 162, 183, 77, 58, 3, 148, 83, 25, 221, 60, 113, 34, 254, 214, 255, 83, 187, 139, 140, 179, 160, 60}

	cert, err = x509.ParseCertificate(derBytes)
	assert.Nil(err)
	log.Printf("Extensions :: %#v", cert.Extensions)
	log.Printf("ExtraExtensions:: %#v", cert.ExtraExtensions)

	pub := cert.PublicKey.(*rsa.PublicKey)

	message := `{"nonce":"ca1bb26d4a756cabf422206fc1953e4b","id":"179687352362288239547319787000716174273","timestamp":"2020-09-14T23:07:16.215597","version":4,"epidPseudonym":"g4cL6vn6M9IDTPSqhX8Pf7Sr9+T7z4gDo9AS85sRtTzb/TwNlXWinJvc32CaMyYxBS47BasT0X28+sZcwivjU0sMLvw4m6+fzHNNn35aDNSpxb0Uex3jzgDuCRFnf8ALnusnQCta9T4+pdSa8q+jiH/rH8o5rhWhbMEWQOn6eL4=","advisoryURL":"https://security-center.intel.com","advisoryIDs":["INTEL-SA-00334"],"isvEnclaveQuoteStatus":"SW_HARDENING_NEEDED","isvEnclaveQuoteBody":"AgABAMYLAAALAAoAAAAAAJa61F5HK4XuN+hpUAosFDUAAAAAAAAAAAAAAAAAAAAADw8DBf+ABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABwAAAAAAAAAHAAAAAAAAAEX7JCJMNjPsjbUdCQvxHeTedsKGbAYBAjFQINmXhrgsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADRH0aZv+C3tUfOY+GILgHu0MZUeSireJoxWoeJjyxTTQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACrVp3CmSVw8JKk216nJxDjuvgQhd5061+C3IFKOR4zFbRGu2agQhwp2GNkGUHW8zZaRLp4BJ0UyeGr0mJbxhkU"}`

	signature := []byte{164, 105, 80, 134, 234, 173, 20, 233, 176, 192, 25, 170, 37, 122, 173, 94, 120, 55, 98, 212, 183, 187, 59, 31, 240, 29, 174, 87, 172, 54, 130, 3, 13, 59, 86, 196, 184, 158, 92, 217, 70, 198, 227, 246, 144, 228, 146, 81, 119, 241, 39, 69, 6, 15, 100, 53, 62, 28, 53, 194, 127, 121, 234, 167, 234, 97, 45, 195, 138, 118, 4, 207, 165, 114, 78, 22, 85, 167, 77, 74, 135, 25, 115, 81, 97, 222, 27, 227, 110, 0, 210, 66, 161, 3, 166, 188, 114, 73, 50, 201, 9, 138, 41, 27, 144, 163, 91, 255, 221, 42, 194, 86, 198, 103, 130, 155, 90, 64, 61, 249, 48, 106, 69, 205, 196, 118, 35, 153, 243, 197, 124, 204, 79, 205, 125, 181, 12, 190, 13, 25, 192, 30, 53, 190, 149, 11, 230, 63, 116, 15, 55, 231, 226, 169, 242, 126, 181, 8, 81, 98, 140, 166, 26, 138, 66, 4, 170, 178, 111, 158, 129, 140, 217, 171, 157, 212, 23, 225, 191, 137, 187, 254, 127, 111, 138, 209, 39, 250, 26, 250, 96, 217, 48, 113, 99, 175, 107, 179, 17, 213, 139, 116, 98, 193, 149, 89, 202, 239, 248, 42, 155, 39, 67, 173, 142, 59, 191, 54, 26, 196, 19, 67, 25, 159, 210, 199, 112, 156, 218, 117, 76, 1, 30, 251, 240, 15, 57, 141, 41, 242, 70, 42, 134, 68, 224, 117, 137, 47, 152, 246, 220, 192, 32, 201, 242, 58}

	log.Println(hex.EncodeToString(signature))
	hashed := sha256.Sum256([]byte(message))
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
	assert.Nil(err)
}
