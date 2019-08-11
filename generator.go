package generator

import "crypto/rand"

func Generate(length int) string {
    output := make([]byte, length)
    randomness := make([]byte, length)
    rand.Read(randomness)
    for pos := range output {
        random := randomness[pos]
        randomPos := random % uint8(64)
        output[pos] = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"[randomPos]
    }
    return string(output)
}
