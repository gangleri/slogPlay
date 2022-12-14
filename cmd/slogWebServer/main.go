package main

import (
	"context"
	"io"
	"net/http"
	"os"

	"golang.org/x/exp/slog"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default().With("url", r.URL) // url will be converted to an Attrs
	ctx := slog.NewContext(r.Context(), logger)

	m := process(ctx, r)
	io.WriteString(w, m)
}

func process(ctx context.Context, r *http.Request) string {
	l := slog.FromContext(ctx)
	//l.Info("processing request", "result", "request processed")
	l.LogAttrs(slog.LevelInfo, "processing request",
		slog.String("result", "request processed"))
	return "request processed"
}

func main() {

	// would typically initialise at program start and pass this to an API layer
	logger := slog.New(slog.NewTextHandler(os.Stdout))
	slog.SetDefault(logger)

	http.HandleFunc("/", getRoot)
	http.ListenAndServe(":8888", nil)
}
