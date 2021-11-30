package web

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const sessionName = "qaaa-session"

type sessionStore struct {
	session *sessions.CookieStore
}

func newSessionStore() (*sessionStore, error) {
	st := &sessionStore{
		session: sessions.NewCookieStore([]byte("sup3r4sEcRETT!!!")),
	}

	return st, nil
}

func (st *sessionStore) flashes(w http.ResponseWriter, r *http.Request) ([]interface{}, error) {
	session, err := st.session.Get(r, sessionName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	flashes := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		return nil, err
	}
	return flashes, nil
}

func (st *sessionStore) addFlash(flash string, w http.ResponseWriter, r *http.Request) error {
	session, err := st.session.Get(r, sessionName)
	if err != nil {
		log.Println(err)
		return err
	}
	session.AddFlash(flash)
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func (st *sessionStore) authenticate(userId int, w http.ResponseWriter, r *http.Request) error {
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

func (st *sessionStore) isAuthenticated(r *http.Request) bool {
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
