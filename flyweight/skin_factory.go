package flyweight

import "fmt"

const (
	T  = "coolguy"
	CT = "agent"
)

var skinFactoryInstance = &SkinFactory{
	skins: make(map[string]Skin),
}

type SkinFactory struct {
	skins map[string]Skin
}

func (s *SkinFactory) getSkinByType(playerType string) (Skin, error) {
	if s.skins[playerType] != nil {
		return s.skins[playerType], nil
	}

	if playerType == T {
		s.skins[playerType] = TSkin{}
		return s.skins[playerType], nil
	}

	if playerType == CT {
		s.skins[playerType] = CTSkin{}
		return s.skins[playerType], nil
	}

	return nil, fmt.Errorf("unrecognized skin type")
}
