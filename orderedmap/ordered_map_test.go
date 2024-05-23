package orderedmap

import (
	"testing"
)

func TestOrderedMap_Add(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")

	if val, ok := om.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected 'value1', got '%s'", val)
	}
}

func TestOrderedMap_Delete(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")
	om.Delete("key1")

	if _, ok := om.Get("key1"); ok {
		t.Errorf("Expected key1 to be deleted")
	}
}

func TestOrderedMap_Get(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")

	if val, ok := om.Get("key1"); !ok || val != "value1" {
		t.Errorf("Expected 'value1', got '%s'", val)
	}

	if _, ok := om.Get("nonexistent"); ok {
		t.Errorf("Expected 'nonexistent' key to not be found")
	}
}

func TestOrderedMap_GetAll(t *testing.T) {
	om := NewOrderedMap()
	om.Add("key1", "value1")
	om.Add("key2", "value2")

	allItems := om.GetAll()
	expected := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	if len(allItems) != len(expected) {
		t.Errorf("Expected map length %d, got %d", len(expected), len(allItems))
	}

	for k, v := range expected {
		if val, ok := allItems[k]; !ok || val != v {
			t.Errorf("Expected key '%s' to have value '%s', got '%s'", k, v, val)
		}
	}
}
