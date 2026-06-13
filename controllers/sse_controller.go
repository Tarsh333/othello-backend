package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"othello-backend/events"
	"strconv"
)

func HandleSSE(w http.ResponseWriter, r *http.Request) {
	otpStr := r.PathValue("id")
	otp, _ := strconv.Atoi(otpStr)
	// if err != nil {
	// 	http.Error(
	// 		w,
	// 		"invalid otp",
	// 		http.StatusBadRequest,
	// 	)
	// 	return
	// }
	flusher, _ := w.(http.Flusher)
	// if !ok {
	// 	http.Error(
	// 		w,
	// 		"stream unsupported",
	// 		http.StatusInternalServerError,
	// 	)
	// 	return
	// }
	w.Header().Set(
		"Content-Type",
		"text/event-stream",
	)

	w.Header().Set(
		"Cache-Control",
		"no-cache",
	)

	w.Header().Set(
		"Connection",
		"keep-alive",
	)
	ch := events.Subscribe(otp)
	defer events.Unsubscribe(
		otp,
		ch,
	)
	for {

		select {

		case game := <-ch:

			payload, err := json.Marshal(game)

			if err != nil {
				continue
			}

			fmt.Fprintf(
				w,
				"data: %s\n\n",
				payload,
			)

			flusher.Flush()

		case <-r.Context().Done():

			return
		}
	}
}
