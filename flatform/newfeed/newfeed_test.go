package newfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})

	if len(feed.Items) != 1 {
		t.Errorf("Item was not added")
	}

}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{"Item", "Item1"})

	results := feed.GetAll()

	if len(results) == 0 {
		t.Errorf("Item is empty")
	}

}
