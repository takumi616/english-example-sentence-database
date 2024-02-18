package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator"
	"github.com/takumi616/go-english-vocabulary-api/entity"
	"github.com/takumi616/go-english-vocabulary-api/store"
	"github.com/takumi616/go-english-vocabulary-api/testutil"
)

func TestAddTask(t *testing.T) {
	type want struct {
		status       int
		responseFile string
	}
	tests := map[string]struct {
		requestFile string
		want        want
	}{
		"ok": {
			requestFile: "testdata/add/ok_req.json.golden",
			want: want{
				status:       http.StatusCreated,
				responseFile: "testdata/add/ok_rsp.json.golden",
			},
		},
		"badrequest": {
			requestFile: "testdata/add/bad_req.json.golden",
			want: want{
				status:       http.StatusBadRequest,
				responseFile: "testdata/add/bad_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/vocabularies",
				bytes.NewReader(testutil.LoadFile(t, tt.requestFile)),
			)

			sut := AddVocabulary{Store: &store.VocabularyStore{
				Vocabularies: map[entity.VocabularyID]*entity.Vocabulary{},
			}, Validator: validator.New()}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.responseFile),
			)
		})
	}
}
