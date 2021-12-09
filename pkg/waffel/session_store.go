package waffel

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type Flash struct {
	Text  string
	Color string
}

const sessionName = "qaaa-session"

type Session struct {
	session *sessions.CookieStore
}

func newSession() (*Session, error) {
	gob.Register(Flash{})
	st := &Session{
		session: sessions.NewCookieStore([]byte("sup3r4sEcRETT!!!")),
	}

	return st, nil
}

func (st *Session) Flashes(w http.ResponseWriter, r *http.Request) ([]Flash, error) {
	session, err := st.session.Get(r, sessionName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rawFlashes := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		return nil, err
	}

	f := []Flash{}
	for _, rawFlash := range rawFlashes {
		f = append(f, rawFlash.(Flash))
	}
	return f, nil
}

func (st *Session) AddFlash(text string, color string, w http.ResponseWriter, r *http.Request) error {
	session, err := st.session.Get(r, sessionName)
	if err != nil {
		log.Println(err)
		return err
	}
	session.AddFlash(Flash{Text: text, Color: color})
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func (st *Session) Authenticate(userId int, w http.ResponseWriter, r *http.Request) error {
	session, err := st.session.Get(r, sessionName)
	if err != nil {
		log.Println(err)
		return err
	}

	session.Values["userId"] = userId
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (st *Session) IsAuthenticated(r *http.Request) bool {
	session, err := st.session.Get(r, sessionName)
	if err != nil {
		log.Println(err)
		return false
	}

	if session.Values["userId"] == nil {
		return false
	}
	return true
}
