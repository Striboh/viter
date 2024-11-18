package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/Striboh/viter/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTIssuingPeriod - period until the JWT is considered
const JWTIssuingPeriod = time.Minute * 10

var (
	// ErrAuthBearerHeader - user provided invalid "Authentication" header
	ErrAuthBearerHeader = errors.New("invalid token provided in auth header")
	// ErrAuthExpiredJWT - provided by user JWT is expired
	ErrAuthExpiredJWT = errors.New("JWT token expired")
)

// GetAPIToken - (POST /auth) handler
func (s Server) GetAPIToken(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Debug("failed to decode auth request", slog.Any("err", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedUUID, err := uuid.Parse(req.ClientId)
	if err != nil {
		slog.Debug("user sent invalid UUID", slog.Any("err", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hash := sha256.Sum256([]byte(req.ClientSecret))
	hashHex := hex.EncodeToString(hash[:])

	apiClient, err := models.GetApiClient(s.db, parsedUUID)
	if err != nil {
		slog.Error("failed getting api client entry from DB", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if apiClient.Hash != hashHex {
		slog.Debug("user provided not matching to DB secret hash",
			slog.String("hash", hashHex),
			slog.String("dbHash", apiClient.Hash))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := s.GetAPIClientJWT(apiClient)
	if err != nil {
		slog.Error("failed signing creating JWT", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	success := true
	resp := AuthResponse{
		Success: &success,
		Token:   &token,
	}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("failed to encode the json response", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetAPIClientJWT - create and sign JWT token for api client from DB
func (s Server) GetAPIClientJWT(apiClient models.ApiClient) (string, error) { //
	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         "data-hub",
		"iat":         jwt.NewNumericDate(now),
		"clientId":    apiClient.ID,
		"clientRoles": apiClient.Roles,
	})

	return token.SignedString(s.config.JwtSecret)
}

// ValidateJWT - validate JWT parsed from request.
func (s Server) ValidateJWT(r *http.Request) error {
	header := strings.Split(r.Header.Get("Authentication"), " ")
	if len(header) < 2 {
		return ErrAuthBearerHeader
	}
	if header[0] != "Bearer" {
		return ErrAuthBearerHeader
	}
	token, err := jwt.Parse(header[1], func(t *jwt.Token) (interface{}, error) {
		return s.config.JwtSecret, nil
	})
	if err != nil {
		return err
	}

	issuedAt, err := token.Claims.GetIssuedAt()
	if err != nil {
		return err
	}

	now := time.Now()

	if now.Add(JWTIssuingPeriod).Compare(issuedAt.Time) <= 0 {
		return ErrAuthExpiredJWT
	}

	return nil
}
