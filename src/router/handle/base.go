package handle

import (
	"github.com/fastgox/fastgox-api-starter/src/models/dto"
	"github.com/lonng/nano/session"
)

func getCharacter(s *session.Session) *dto.CharacterSession {
	character := s.Value("character")
	characterSession, _ := character.(*dto.CharacterSession)
	return characterSession
}

func setCharacter(s *session.Session, character *dto.CharacterSession) {
	s.Set("character", character)
}
