package main

import (
	"context"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Encrypt(plaintext, key string) string {
	plaintext = strings.ToUpper(plaintext)
	key = strings.ToUpper(key)

	keyRunes := []rune(key)
	result := make([]rune, len(plaintext))
	keyIndex := 0

	for i, ch := range plaintext {
		if ch < 'A' || ch > 'Z' {
			result[i] = ch
			continue
		}

		currentKeyChar := keyRunes[keyIndex]
		keyVal := int(currentKeyChar - 'A')

		encrypted := (int(ch-'A') + keyVal) % 26
		result[i] = rune(encrypted) + 'A'

		keyRunes[keyIndex] = rune((int(currentKeyChar-'A')+1)%26 + 'A')

		keyIndex = (keyIndex + 1) % len(keyRunes)
	}

	return string(result)
}

func (a *App) Decrypt(ciphertext, key string) string {
	ciphertext = strings.ToUpper(ciphertext)
	key = strings.ToUpper(key)

	keyRunes := []rune(key)
	result := make([]rune, len(ciphertext))
	keyIndex := 0

	for i, ch := range ciphertext {
		if ch < 'A' || ch > 'Z' {
			result[i] = ch
			continue
		}

		currentKeyChar := keyRunes[keyIndex]
		keyVal := int(currentKeyChar - 'A')

		decrypted := (int(ch-'A') - keyVal) % 26
		if decrypted < 0 {
			decrypted += 26
		}
		result[i] = rune(decrypted) + 'A'

		keyRunes[keyIndex] = rune((int(currentKeyChar-'A')+1)%26 + 'A')

		keyIndex = (keyIndex + 1) % len(keyRunes)
	}

	return string(result)
}
