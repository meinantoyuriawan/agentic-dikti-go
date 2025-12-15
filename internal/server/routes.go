package server

import "net/http"

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.helloWorldHandler)
	mux.HandleFunc("/health", s.healthHandler)
	mux.HandleFunc("/chatbot", s.chatHandler)

	return s.authMiddleware(mux)
}

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	// can do auth here
	return next
}

// func (s *Server) authMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		path := r.URL.Path
// 		if path == "/health" || path == "/v1/login" || path == "/v1/register" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		authorizationHeader := r.Header.Get("Authorization")
// 		if !strings.Contains(authorizationHeader, "Bearer") {
// 			sendErrorResponse(w, http.StatusUnauthorized, "unauthorized")
// 			return
// 		}

// 		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

// 		userID, err := utils.ParseUserIDFromToken(tokenString)
// 		if err != nil {
// 			sendErrorResponse(w, http.StatusUnauthorized, "unauthorized")
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), constants.UserIDCtxKey, userID)

// 		// Proceed with the next handler
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
