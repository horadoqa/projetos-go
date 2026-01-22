package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// loggingResponseWriter captura status code e tamanho da resposta
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	if lrw.statusCode == 0 {
		lrw.statusCode = http.StatusOK
	}
	n, err := lrw.ResponseWriter.Write(b)
	lrw.size += n
	return n, err
}

// Middleware de logging
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ignora requisições de "ruído" do navegador
		if r.URL.Path == "/cdn-cgi/rum" {
			next.ServeHTTP(w, r)
			return
		}

		lrw := &loggingResponseWriter{ResponseWriter: w}

		log.Println("Request:", r.Method, r.URL.Path)
		log.Println("Headers:", r.Header)

		next.ServeHTTP(lrw, r)

		log.Printf("Response status: %d, size: %d bytes\n", lrw.statusCode, lrw.size)
	})
}

// Handler da API
func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	response := `{"message":"Dados recebidos","body":"` + string(bodyBytes) + `"}`

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	// Configura log para console + arquivo
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Não foi possível abrir arquivo de log:", err)
	}
	defer file.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, file))

	// Serve arquivos estáticos
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", logMiddleware(fs))

	// Se quiser healthcheck separado
	http.Handle("/healthcheck/", logMiddleware(http.StripPrefix("/healthcheck/", http.FileServer(http.Dir("./public/healthcheck")))))

	// API
	http.Handle("/api/data", logMiddleware(http.HandlerFunc(apiDataHandler)))

	log.Println("HTTP Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
