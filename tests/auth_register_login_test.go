package tests

import (
	"testing"
	"time"

	ssov1 "github.com/curtrika/UMetrika_server/pkg/proto/auth/v1"
	"github.com/curtrika/UMetrika_server/tests/suite"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	emptyAppID = 0
	appID      = 1
	appSecret  = "test-secret"

	passDefaultLen = 10
)

// TODO: add token fail validation cases

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	// Initialize test context and suite
	ctx, st := suite.New(t)

	// Generate fake credentials
	email := gofakeit.Email()
	pass := randomFakePassword()

	// Register the user
	respReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: pass,
	})
	require.NoError(t, err, "registration should succeed")
	require.NotEmpty(t, respReg.GetUserId(), "user ID should not be empty upon successful registration")

	// Login with the same credentials
	respLogin, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: pass,
		AppId:    appID,
	})
	require.NoError(t, err, "login should succeed")

	token := respLogin.GetToken()
	require.NotEmpty(t, token, "login token should not be empty")

	// Capture login time to check token expiration
	loginTime := time.Now()

	// Parse and validate the JWT token
	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(appSecret), nil
	})
	require.NoError(t, err, "token parsing should succeed")

	// Type assertion for JWT claims
	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	require.True(t, ok, "claims should be of type jwt.MapClaims")

	// Check user ID
	userID, ok := claims["uid"].(string)
	require.True(t, ok, "user ID claim should be a string")
	assert.Equal(t, respReg.GetUserId(), userID, "user ID should match registered user")

	// Check email
	emailClaim, ok := claims["email"].(string)
	require.True(t, ok, "email claim should be a string")
	assert.Equal(t, email, emailClaim, "email should match registered email")

	// Check app ID
	appIDClaim, ok := claims["app_id"].(float64)
	require.True(t, ok, "app ID claim should be a float64")
	assert.Equal(t, appID, int(appIDClaim), "app ID should match expected app ID")

	// Check token expiration
	expClaim, ok := claims["exp"].(float64)
	require.True(t, ok, "expiration claim should be a float64")
	const deltaSeconds = 1

	assert.InDelta(t,
		loginTime.Add(st.Cfg.TokenTTL).Unix(),
		expClaim,
		deltaSeconds,
		"token expiration should be within acceptable range")
}

func TestRegisterLogin_DuplicatedRegistration(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	pass := randomFakePassword()

	respReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: pass,
	})
	require.NoError(t, err)
	require.NotEmpty(t, respReg.GetUserId())

	respReg, err = st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: pass,
	})
	require.Error(t, err)
	assert.Empty(t, respReg.GetUserId())
	// assert.ErrorContains(t, err, "exists") // client error does not contain these words
	assert.ErrorContains(t, err, "failed to register user")
}

func TestRegister_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		name        string
		email       string
		password    string
		expectedErr string
	}{
		{
			name:        "Register with Empty Password",
			email:       gofakeit.Email(),
			password:    "",
			expectedErr: "password is required",
		},
		{
			name:        "Register with Empty Email",
			email:       "",
			password:    randomFakePassword(),
			expectedErr: "email is required",
		},
		{
			name:        "Register with Both Empty",
			email:       "",
			password:    "",
			expectedErr: "email is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
				Email:    tt.email,
				Password: tt.password,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}

func TestLogin_FailCases(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		name        string
		email       string
		password    string
		appID       int32
		expectedErr string
	}{
		{
			name:        "Login with Empty Password",
			email:       gofakeit.Email(),
			password:    "",
			appID:       appID,
			expectedErr: "password is required",
		},
		{
			name:        "Login with Empty Email",
			email:       "",
			password:    randomFakePassword(),
			appID:       appID,
			expectedErr: "email is required",
		},
		{
			name:        "Login with Both Empty Email and Password",
			email:       "",
			password:    "",
			appID:       appID,
			expectedErr: "email is required",
		},
		{
			name:        "Login with Non-Matching Password",
			email:       gofakeit.Email(),
			password:    randomFakePassword(),
			appID:       appID,
			expectedErr: "invalid email or password",
		},
		{
			name:        "Login without AppID",
			email:       gofakeit.Email(),
			password:    randomFakePassword(),
			appID:       emptyAppID,
			expectedErr: "app_id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
				Email:    gofakeit.Email(),
				Password: randomFakePassword(),
			})
			require.NoError(t, err)

			_, err = st.AuthClient.Login(ctx, &ssov1.LoginRequest{
				Email:    tt.email,
				Password: tt.password,
				AppId:    tt.appID,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, passDefaultLen)
}
