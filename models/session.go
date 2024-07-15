package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session.
	// We will use this to store the session token in the database.
	// This will be left empty, as i only gonna store the hash of a session token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// Todo : Create the session token
	// Todo : Implement SessionService.Create
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	// Todo : Implement SessionService.User
	return nil, nil
}
