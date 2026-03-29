package constants

import "errors"

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInvalidRefreshToken = errors.New("invalid refresh token")
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidToken = errors.New("invalid token")
var ErrNilUser = errors.New("user is nil")
var ErrUserIdRequired = errors.New("user id is required")
var ErrNoPassword = errors.New("password is required")
var ErrNilTokenManager = errors.New("token manager is nil")
