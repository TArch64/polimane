package xor

func Encrypt(data, key []byte) []byte {
	result := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ key[i%len(key)]
	}

	return result
}
