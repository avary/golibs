package token

import (
	"errors"
	"fmt"
	"strings"

	"github.com/minio/sha256-simd"

	"github.com/skerkour/golibs/base32"
	"github.com/skerkour/golibs/crypto"
	"github.com/skerkour/golibs/uuid"
)

const (
	SecretSize = crypto.KeySize256
	HashSize   = crypto.KeySize256
)

var (
	ErrTokenIsNotValid = errors.New("token is not valid")
	ErrDataIsTooLong   = errors.New("data is too long")
)

type Token struct {
	id     uuid.UUID
	secret []byte
	hash   []byte
	str    string
	prefix string
}

func New(prefix string) (token Token, err error) {
	secret, err := newSecret()
	if err != nil {
		return
	}

	return newToken(prefix, secret), nil
}

// func NewWithSecret(secret []byte) (token Token, err error) {
// 	return new("", secret)
// }

// func NewWithPrefix(prefix string) (token Token, err error) {
// 	secret, err := newSecret()
// 	if err != nil {
// 		return
// 	}
// 	return newToken(prefix, secret)
// }

func newSecret() (secret []byte, err error) {
	secret, err = crypto.RandBytes(SecretSize)
	if err != nil {
		err = fmt.Errorf("token: Generating secret: %w", err)
		return
	}
	return
}

func newToken(prefix string, secret []byte) (token Token) {
	id := uuid.New()
	idBytes, _ := id.MarshalBinary()

	hash := generateHash(idBytes, secret)

	data := append(idBytes, secret...)
	str := base32.EncodeToString(data)
	str = prefix + str

	token = Token{
		id,
		secret,
		hash,
		str,
		prefix,
	}
	return
}

func (token *Token) String() string {
	return token.str
}

func (token *Token) ID() uuid.UUID {
	return token.id
}

func (token *Token) Secret() []byte {
	return token.secret
}

func (token *Token) Hash() []byte {
	return token.hash
}

func Parse(prefix, input string) (token Token, err error) {
	var tokenBytes []byte

	token.str = input

	if prefix != "" {
		if !strings.HasPrefix(input, prefix) {
			err = ErrTokenIsNotValid
			return
		}
		input = strings.TrimPrefix(input, prefix)
		token.prefix = prefix
	}

	tokenBytes, err = base32.DecodeString(input)
	if err != nil {
		err = ErrTokenIsNotValid
		return
	}

	if len(tokenBytes) != uuid.Size+SecretSize {
		err = ErrTokenIsNotValid
		return
	}

	tokenIDBytes := tokenBytes[:uuid.Size]
	token.secret = tokenBytes[uuid.Size:]

	token.id, err = uuid.FromBytes(tokenIDBytes)
	if err != nil {
		err = ErrTokenIsNotValid
		return
	}

	token.hash = generateHash(tokenIDBytes, token.secret)

	return
}

// FromIdAndHash creates a new token from an ID and a Hash
// it means that the token needs to be refreshed with `Refresh` before being able to use it
// as we din't have the secret, and thus cannot convert it to a valid string
func FromIdAndHash(prefix string, id uuid.UUID, hash []byte) (token Token, err error) {
	if len(hash) != HashSize {
		err = ErrTokenIsNotValid
		return
	}

	token = Token{
		id:     id,
		secret: nil,
		hash:   hash,
		str:    "",
		prefix: prefix,
	}

	return
}

func (token *Token) Verify(hash []byte) (err error) {
	// in case we need to update hash size later
	// if len(hash) == OldHashSize {
	// token.hash = crypto.DeriveKeyFromKey(secret, idBytes, OldHashSize)
	// ..
	// }

	if !crypto.ConstantTimeCompare(hash, token.hash) {
		err = ErrTokenIsNotValid
	}
	return
}

func (token *Token) Refresh() (err error) {
	idBytes, _ := token.id.MarshalBinary()
	token.secret, err = newSecret()
	if err != nil {
		return
	}

	token.hash = generateHash(idBytes, token.secret)

	data := append(idBytes, token.secret...)
	str := base32.EncodeToString(data)
	token.str = token.prefix + str

	return
}

func generateHash(tokenID, secret []byte) (hash []byte) {
	hasher := sha256.New()
	hasher.Write(tokenID)
	hasher.Write(secret)
	hash = hasher.Sum(nil)
	return
}
