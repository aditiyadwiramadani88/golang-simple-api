package response
import "net/http"

func JsonResponse(w http.ResponseWriter, resp string, status int) (int, error) { 
	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	return w.Write([]byte(resp))
} 