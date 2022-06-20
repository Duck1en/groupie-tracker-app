package delivery

import (
	"net/http"
	"text/template"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) Router() *http.ServeMux {
	var err error
	MainTmpl, err = template.ParseFiles("template/index.html")
	ErrorCheck(err)
	ErrTmpl, err = template.ParseFiles("template/error.html")
	ErrorCheck(err)
	ArtistTmpl, err = template.ParseFiles("template/artist.html")
	ErrorCheck(err)
	s.mux.HandleFunc("/", HomePage)
	s.mux.HandleFunc("/artist", ArtistPage)
	s.mux.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template/"))))
	s.mux.Handle("/template/assets/", http.StripPrefix("/template/assets/", http.FileServer(http.Dir("./template/assets/"))))
	// s.mux.Handle("/template/map/", http.StripPrefix("/template/map/", http.FileServer(http.Dir("./template/map/"))))
	return s.mux
}

func PathPage(page http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(page, http.StatusNotFound, ErrNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(page, http.StatusBadRequest, ErrBadRequest)
		return
	}
	MainTmpl.Execute(page, allArtists)
}

// func ArtistPage(page http.ResponseWriter, r *http.Request) {
// 	id := strings.TrimPrefix(r.URL.Path, "/artist/")
// 	n, _ := strconv.Atoi(id)
// 	fmt.Println(allArtists[n-1])
// 	// MainTmpl.Execute(page, allArtists[n-1:n])
// }
