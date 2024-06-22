package middleware

import (
	"context"
	"net/http"

	"github.com/didikz/godisb/pkg/httpapi"
)

type ctxRequestSignature string
type ctxIdempotencyKey string

const (
	CtxSignature      ctxRequestSignature = "ctxRequestSignature"
	CtxIdempotencyKey ctxIdempotencyKey   = "ctxIdempotencyKey"
)

func HeaderValidator(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		signature := r.Header.Get("X-Signature")
		if signature == "" {
			httpapi.WriteJson(w, http.StatusUnauthorized, httpapi.GeneralResponseError{Error: "invalid signature"})
			return
		}

		ik := r.Header.Get("X-Idempotency-Key")
		if ik == "" {
			httpapi.WriteJson(w, http.StatusBadRequest, httpapi.GeneralResponseError{Error: "idempotency key is empty"})
			return
		}

		ctx = context.WithValue(ctx, CtxSignature, signature)
		ctx = context.WithValue(ctx, CtxIdempotencyKey, ik)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
