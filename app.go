package main

import (
	"context"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

const RussianAlphabet = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

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
	alphabet := []rune(RussianAlphabet)
	alphabetLen := len(alphabet)

	for i, ch := range plaintext {
		chIndex := -1
		for idx, letter := range alphabet {
			if letter == ch {
				chIndex = idx
				break
			}
		}

		if chIndex == -1 {
			result[i] = ch
			continue
		}

		currentKeyChar := keyRunes[keyIndex]
		keyIdx := -1
		for idx, letter := range alphabet {
			if letter == currentKeyChar {
				keyIdx = idx
				break
			}
		}
		if keyIdx == -1 {
			result[i] = ch
			continue
		}

		encryptedIdx := (chIndex + keyIdx) % alphabetLen
		result[i] = alphabet[encryptedIdx]

		newKeyIdx := (keyIdx + 1) % alphabetLen
		keyRunes[keyIndex] = alphabet[newKeyIdx]

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
	alphabet := []rune(RussianAlphabet)
	alphabetLen := len(alphabet)

	for i, ch := range ciphertext {
		chIndex := -1
		for idx, letter := range alphabet {
			if letter == ch {
				chIndex = idx
				break
			}
		}

		if chIndex == -1 {
			result[i] = ch
			continue
		}

		currentKeyChar := keyRunes[keyIndex]
		keyIdx := -1
		for idx, letter := range alphabet {
			if letter == currentKeyChar {
				keyIdx = idx
				break
			}
		}
		if keyIdx == -1 {
			result[i] = ch
			continue
		}

		decryptedIdx := (chIndex - keyIdx) % alphabetLen
		if decryptedIdx < 0 {
			decryptedIdx += alphabetLen
		}
		result[i] = alphabet[decryptedIdx]

		newKeyIdx := (keyIdx + 1) % alphabetLen
		keyRunes[keyIndex] = alphabet[newKeyIdx]

		keyIndex = (keyIndex + 1) % len(keyRunes)
	}

	return string(result)
}
