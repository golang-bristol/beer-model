package model_test

import (
	"testing"
	"time"

	model "github.com/golang-bristol/beer-model"
)

func TestReview(t *testing.T) {
	review := model.Review{1, 100, "Tim", "Lawrence", 7, "This is a great beer", time.Now()}

	if review.Beer_id != 100 {
		t.Error("Expected 100 got", review.Beer_id)
	}

}
