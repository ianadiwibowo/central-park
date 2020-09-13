// package pupuruhashmap_test
package pupuruhashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPupuruHashmap(t *testing.T) {
	hashmap := NewPupuruHashmap()

	assert.Equal(
		t,
		hashmap,
		&PupuruHashmap{},
	)
}

func TestSet_JustOneAssignment(t *testing.T) {
	h := NewPupuruHashmap()

	h.Set("cat", "Lupita")

	assert.Equal(t, "cat", h.Data[9].Key)
	assert.Equal(t, "Lupita", h.Data[9].Value)
}

func TestSet_WhenCollision_WithSameKeys(t *testing.T) {
	h := NewPupuruHashmap()
	h.Set("cat", "Lupita")
	h.Set("cat", "Vito")

	assert.Equal(t, "cat", h.Data[9].Key)
	assert.Equal(t, "Vito", h.Data[9].Value)
}

func TestSet_WhenCollision_WithDifferentKeys(t *testing.T) {
	h := NewPupuruHashmap()
	h.Set("cat1", "Lupita")
	h.Set("cat2", "Vito")
	h.Set("cat3", "Rijong")

	assert.Equal(t, "cat1", h.Data[9].Key)
	assert.Equal(t, "Lupita", h.Data[9].Value)
	assert.Equal(t, "cat2", h.Data[9].Next.Key)
	assert.Equal(t, "Vito", h.Data[9].Next.Value)
	assert.Equal(t, "cat3", h.Data[9].Next.Next.Key)
	assert.Equal(t, "Rijong", h.Data[9].Next.Next.Value)
}

func TestGet_JustOneRetrieval(t *testing.T) {
	h := NewPupuruHashmap()

	h.Set("cat", "Lupita")

	assert.Equal(t, "Lupita", h.Get("cat"))
}

func TestGet_WhenNotFound_ReturnEmptyString(t *testing.T) {
	h := NewPupuruHashmap()

	h.Set("cat", "Lupita")

	assert.Equal(t, "", h.Get("otter"))
}

func TestGet_WhenCollision_WithSameKeys(t *testing.T) {
	h := NewPupuruHashmap()
	h.Set("cat", "Lupita")
	h.Set("cat", "Vito")

	assert.Equal(t, "Vito", h.Get("cat"))
}

func TestGet_WhenCollision_WithDifferentKeys(t *testing.T) {
	h := NewPupuruHashmap()
	h.Set("cat1", "Lupita")
	h.Set("cat2", "Vito")
	h.Set("cat3", "Rijong")

	assert.Equal(t, "Lupita", h.Get("cat1"))
	assert.Equal(t, "Vito", h.Get("cat2"))
	assert.Equal(t, "Rijong", h.Get("cat3"))
}

func TestMapKeyToIndex(t *testing.T) {
	assert.Equal(t, 4, mapKeyToIndex("Test"))
	assert.Equal(t, 5, mapKeyToIndex("some_key"))
	assert.Equal(t, 9, mapKeyToIndex("cat"))
}
