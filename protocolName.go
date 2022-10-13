package ethaddr

import "fmt"

type ProtocolName string

func (name ProtocolName) Str() string {
	return fmt.Sprint(name)
}

// Protocol lend is aave like.
func (name ProtocolName) IsAaveLike() bool {
	for _, pro := range AaveLikeProtocols {
		if pro == name.Str() {
			return true
		}
	}
	return false
}

// Protocol lend is compound like.
func (name ProtocolName) IsCompoundLike() bool {
	for _, pro := range CompoundLikeProtocols {
		if pro == name.Str() {
			return true
		}
	}
	return false
}
