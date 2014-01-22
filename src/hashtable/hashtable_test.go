package hashtable

import (
	"testing"
)

func Test_when_the_hashtable_does_not_contain_the_key(t *testing.T) {
	ht := NewHashTable()
	actualValue := ht.Get("key")

	t.Log("It should return nil")
	if actualValue != nil {
		t.Errorf("\tExpected nil but was %v", actualValue)
	}
}

func Test_when_retrieving_a_known_key(t *testing.T) {
	expectedValue := "1"

	ht := NewHashTable()
	ht.Set("key", expectedValue)

	actualValue := ht.Get("key")

	t.Log("It should return the instance that was stored at that key")
	if actualValue != expectedValue {
		t.Errorf("Expected %#v but was %#v", expectedValue, actualValue)
	}
}

func Test_when_retrieving_multiple_known_keys(t *testing.T) {
	expectedValue1 := "value1"
	expectedValue2 := "value2"

	ht := NewHashTable()
	ht.Set("key1", expectedValue1)
	ht.Set("key2", expectedValue2)

	actualValue1 := ht.Get("key1")
	actualValue2 := ht.Get("key2")

	t.Log("It should return the instances that were stored at the keys")
	if actualValue1 != expectedValue1 || actualValue2 != expectedValue2 {
		t.Errorf("Expected key1 = %#v but was %#v and key2 = %#v but was %#v", expectedValue1, actualValue1, expectedValue2, actualValue2)
	}
}

func Test_when_retrieving_a_key_that_has_been_overwritten(t *testing.T) {
	firstValue := "value1"
	expectedValue := "value2"

	ht := NewHashTable()
	ht.Set("key1", firstValue)
	ht.Set("key1", expectedValue)

	actualValue := ht.Get("key1")

	t.Log("It should return the last value stored at the key")
	if actualValue != expectedValue {
		t.Errorf("\tExpected %#v but was %#v", expectedValue, actualValue)
	}
}
