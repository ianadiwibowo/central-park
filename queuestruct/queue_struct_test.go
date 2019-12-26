package queuestruct_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ianadiwibowo/central-park/queuestruct"
)

// cat demonstrates that a custom struct can be used as the queue internal data structure
type cat struct {
	ID   int
	Name string
}

func TestNewQueue(t *testing.T) {
	q := queuestruct.NewQueueStruct()

	if q.Print() != "[]" {
		t.Errorf("Expected: []. Got: %v", q.Print())
	}
}

func TestEnqueue(t *testing.T) {
	q := queuestruct.NewQueueStruct()
	q.Enqueue(&cat{1, "Pupuru"})
	q.Enqueue(&cat{3, "Hosico"})

	if q.Print() != "[&{1 Pupuru} &{3 Hosico}]" {
		t.Errorf("Expected: [&{1 Pupuru} &{3 Hosico}], Got: %v", q.Print())
	}
}

func TestDequeue(t *testing.T) {
	q := queuestruct.NewQueueStruct()
	q.Enqueue(&cat{5, "Yuen Jumbo"})
	q.Enqueue(&cat{3, "Fomo"})
	q.Enqueue(&cat{1, "Ah Fei"})
	v := q.Dequeue().(*cat)

	if cmp.Equal(v, &cat{5, "Yuen Jumbo"}) == false {
		t.Errorf("Expected: &{5 Yuen Jumbo}, Got: %v", v)
	}

	if q.Print() != "[&{3 Fomo} &{1 Ah Fei}]" {
		t.Errorf("Expected: [&{3 Fomo} &{1 Ah Fei}], Got: %v", q.Print())
	}
}

func TestDequeueOnEmptyQueue(t *testing.T) {
	q := queuestruct.NewQueueStruct()
	v := q.Dequeue()

	if v != nil {
		t.Errorf("Expected: nil, Got: %v", q.Print())
	}

	if q.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", q.Print())
	}
}

func TestPeek(t *testing.T) {
	q := queuestruct.NewQueueStruct()
	q.Enqueue(&cat{1, "Lupita"})
	q.Enqueue(&cat{2, "Bombi"})
	v := q.Peek().(*cat)

	if cmp.Equal(v, &cat{1, "Lupita"}) == false {
		t.Errorf("Expected: &cat{1 Lupita}, Got: %v", q.Print())
	}

	if q.Print() != "[&{1 Lupita} &{2 Bombi}]" {
		t.Errorf("Expected: [&{1 Lupita} &{2 Bombi}], Got: %v", q.Print())
	}
}

func TestPeekOnEmptyQueue(t *testing.T) {
	q := queuestruct.NewQueueStruct()
	v := q.Peek()

	if v != nil {
		t.Errorf("Expected: nil, Got: %v", q.Print())
	}

	if q.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", q.Print())
	}
}

func TestClear(t *testing.T) {
	q := queuestruct.NewQueueStruct()
	q.Enqueue(&cat{10, "Simabossneko"})
	q.Enqueue(&cat{20, "Big Billy"})
	q.Clear()

	if q.Print() != "[]" {
		t.Errorf("Expected: [], Got: %v", q.Print())
	}
}

func TestIsEmptyOnEmptyQueue(t *testing.T) {
	q := queuestruct.NewQueueStruct()

	if q.IsEmpty() == false {
		t.Errorf("Expected: true, Got: %v", q.Print())
	}
}

func TestIsEmptyOnNotEmptyQueue(t *testing.T) {
	q := queuestruct.NewQueueStruct()
	q.Enqueue(&cat{7, "Bellamy"})

	if q.IsEmpty() == true {
		t.Errorf("Expected: false, Got: %v", q.Print())
	}
}
