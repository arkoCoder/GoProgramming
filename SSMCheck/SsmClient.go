package main

import (
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("TestSSM")
	var sess = session.Must(session.NewSession())
	var ssmClient = ssm.New(sess, aws.NewConfig().WithRegion("us-west-2"))
	serverPublicKey := "/daas/secrets/default/lh_security_lh-keys_privkey_pem"
	provisionKey := "/daas/secrets/default/lh_security_prov-keys_privkey_pem"
	publicKeySignature := "/daas/secrets/default/prov_keys_pubkey_der_sig_base64"
	keyArray := []*string{aws.String(serverPublicKey), aws.String(provisionKey), aws.String(publicKeySignature)}
	//// get the Key name from env and add public key
	withDecryption := true
	ssmResponse, err := ssmClient.GetParameters(&ssm.GetParametersInput{
		Names:          keyArray,
		WithDecryption: &withDecryption,
	})
	fmt.Printf("Response from SSM is %v \n", ssmResponse)
	var testParameters []*ssm.Parameter
	testParameters = ssmResponse.Parameters
	fmt.Printf("Fetching value of only parameters %v \n", testParameters[0])

	resp, err := ssmClient.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(serverPublicKey),
		WithDecryption: &withDecryption,
	})
	fmt.Printf("Response is %v %T \n", resp, resp)
	fmt.Printf("Error is %v \n", err)
	if err != nil {
		fmt.Println("do something")
	}
	var privateKey string

	privateKey = *resp.Parameter.Value
	fmt.Printf("private key %v \n", privateKey)
	mySigningKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		fmt.Println("improve error handling")
	}

	fmt.Printf("private key %v \n", mySigningKey.PublicKey)

	x509.MarshalPKCS1PublicKey(&mySigningKey.PublicKey)
	fmt.Printf("Encoded public key %v \n", base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&mySigningKey.PublicKey)))
	publicRsaKey, err := ssh.NewPublicKey(&mySigningKey.PublicKey)
	if err != nil {
		fmt.Println("improve error handling")
	}
	fmt.Printf("Public key %v %T \n", publicRsaKey, publicRsaKey)
	fmt.Printf("Private key %v %T \n", mySigningKey, mySigningKey)
	fmt.Printf("%v %T \n", *&mySigningKey.PublicKey, *&mySigningKey.PublicKey)

	//fmt.Printf("Public Key %v %T \n", mySigningKey.PublicKey, mySigningKey.PublicKey)
	//decoded, err := base64.StdEncoding.DecodeString()
	//if err != nil {
	//	fmt.Println("decode error:", err)
	//	return
	//}
	//fmt.Println(string(decoded))
}
