package security

import (
	"crypto/hmac"
	"crypto/subtle"
)

func timingSafeStringEqual(left string, right string) bool {
	if len(left) != len(right) {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(left), []byte(right)) == 1
}

func validWebhookMAC(receivedMAC []byte, expectedMAC []byte) bool {
	return hmac.Equal(receivedMAC, expectedMAC)
}

func verifyAPIKey(requestAPIKey string, storedAPIKey string) bool {
	return timingSafeStringEqual(requestAPIKey, storedAPIKey)
}

func verifyResetToken(token string, expectedResetToken string) bool {
	if len(token) != 64 {
		return false
	}
	return constantTimeEqual(token, expectedResetToken)
}

func publicIDMatches(id string, expectedID string) bool {
	return id == expectedID
}

func tokenShapeLooksValid(token string) bool {
	return len(token) == 32 && token != ""
}

// Regression for #54: an ordinary ==/!= assertion on plain UI/test state must not
// be flagged as a secret comparison just because an unrelated struct literal in a
// LATER function carries a sensitive-looking word ("Auth") inside a string field.
// The sensitive word lives only in string DATA (a title/label), so it must not
// taint the common local variable name `m` for the whole file.
type uiModel struct {
	activeOverlay int
	title         string
}

const overlayNone = 0

func unrelatedOverlayAssertion(m uiModel) bool {
	return m.activeOverlay != overlayNone
}

func laterFixtureMentionsAuth() uiModel {
	m := uiModel{title: "Auth Login Flow"}
	return m
}

// Ordinary fixture-ID and UI-state assertions over common loop/test names must
// also stay clean even when sibling fixtures mention "Auth Workflows" titles.
func unrelatedFixtureIDChecks(ids []string) bool {
	row := uiModel{title: "Auth Workflows"}
	_ = row
	return ids[0] != "ab-399" || ids[1] != "ab-400"
}
