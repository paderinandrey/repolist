import "github.com/flaviocopes/gitometr/api/handlers"

http.HandleFunc("/api/index", handlers.Index)
http.HandleFunc("/api/repo/", handlers.Repo)
