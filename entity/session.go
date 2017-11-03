package entity

import (
	"encoding/json"
	"io"

	"../logger"
)

type Session struct {
	CurrentUser string
}

func (session *Session) Serialize(w io.Writer) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(session)
	logger.FatalIf(err)
}

func DeserializeSession(r io.Reader) *Session {
	decoder := json.NewDecoder(r)
	session := new(Session)

	err := decoder.Decode(session)
	if err == io.EOF {
		return session
	}
	logger.FatalIf(err)
	return session
}
