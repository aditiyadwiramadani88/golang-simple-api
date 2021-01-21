package response
import "net/http"
func HttpResponse(w http.ResponseWriter, resp string)(int, error)   {
		w.Header().Add("content-type", "text/html")
		return w.Write([]byte(resp))
}
func JsonResponse(w http.ResponseWriter, resp string) (int, error) { 
	w.Header().Add("content-type", "application/json")
	return w.Write([]byte(resp))
} 