/*
 * Copyright (c) 2024. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log/slog"
	"os"
)

func SignRSA256(privateKeyPath, message string) (string, error) {
	// Load private key from file
	privateKeyFile, err := os.ReadFile(privateKeyPath)
	if err != nil {
		slog.Error("failed to read private key", "error", err)
		return "", err
	}

	// Decode PEM encoded private key
	privateKeyBlock, _ := pem.Decode(privateKeyFile)
	if privateKeyBlock == nil {
		slog.Error("failed to decode private key", "error", err)
		return "", err
	}

	// Parse RSA private key
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		slog.Error("failed to parse private key", "error", err)
		return "", err
	}

	// Message to sign
	messageByte := []byte(message)

	// Hash the messageByte
	hashed := sha256.Sum256(messageByte)

	// Sign the hash with private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hashed[:])
	// signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		slog.Error("failed to sign signature", "error", err)
		return "", err
	}

	toString := base64.StdEncoding.EncodeToString(signature)

	return toString, nil
}
