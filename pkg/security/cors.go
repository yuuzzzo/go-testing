package security

import "net/http"

func EnableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Permite qualquer origem (em produção, troque pelo seu domínio)
        w.Header().Set("Access-Control-Allow-Origin", "*")
        
        // Permite os métodos que você já criou (GET, POST, DELETE, etc.)
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        
        // Permite os cabeçalhos comuns, como Content-Type
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Importante: Trata a requisição de "preflight" (OPTIONS)
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}