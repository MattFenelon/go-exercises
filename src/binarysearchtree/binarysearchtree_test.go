package binarysearchtree

import (
	"reflect"
	"testing"
)

func Test_when_retrieving_an_unknown_key(t *testing.T) {
	bst := NewBinarySearchTree()

	actualValue := bst.Get("unknown")

	t.Log("It should return nil")
	if actualValue != nil {
		t.Errorf("\tExpected nil but was %#v", actualValue)
	}
}

func Test_when_retrieving_a_known_key(t *testing.T) {
	expectedValue := "value"

	bst := NewBinarySearchTree()
	bst.Set("key", expectedValue)

	actualValue := bst.Get("key")

	t.Log("It should return the stored instance for the key")
	if actualValue != expectedValue {
		t.Errorf("\tExpected %#v but was %#v", expectedValue, actualValue)
	}
}

func Test_when_2_keys_are_added_and_the_second_key_is_greater_than_the_first(t *testing.T) {
	expectedValue1 := "key1:value1"
	expectedValue2 := "key2:value2"

	bst := NewBinarySearchTree()
	bst.Set("key1", expectedValue1)
	bst.Set("key2", expectedValue2)

	actualValue1 := bst.Get("key1")
	actualValue2 := bst.Get("key2")

	t.Log("It should return the stored instances for all keys")
	if actualValue1 != expectedValue1 || actualValue2 != expectedValue2 {
		t.Errorf("\tExpected %#v, %#v but was %#v, %#v", expectedValue1, expectedValue2, actualValue1, actualValue2)
	}
}

func Test_when_3_keys_are_added_where_each_is_greater_than_the_last(t *testing.T) {
	expectedValue1 := "key1:value1"
	expectedValue2 := "key2:value2"
	expectedValue3 := "key3:value3"

	bst := NewBinarySearchTree()
	bst.Set("key1", expectedValue1)
	bst.Set("key2", expectedValue2)
	bst.Set("key3", expectedValue3)

	actualValue1 := bst.Get("key1")
	actualValue2 := bst.Get("key2")
	actualValue3 := bst.Get("key3")

	t.Log("It should return the stored instances for all keys")
	if actualValue1 != expectedValue1 || actualValue2 != expectedValue2 || actualValue3 != expectedValue3 {
		t.Errorf("\tExpected %#v, %#v, %#v but was %#v, %#v, %#v", expectedValue1, expectedValue2, expectedValue3, actualValue1, actualValue2, actualValue3)
	}
}

func Test_when_2_keys_are_added_and_the_second_key_is_less_than_the_first(t *testing.T) {
	expectedValue1 := "key1:value1"
	expectedValue2 := "key0:value0"

	bst := NewBinarySearchTree()
	bst.Set("key1", expectedValue1)
	bst.Set("key0", expectedValue2)

	actualValue1 := bst.Get("key1")
	actualValue2 := bst.Get("key0")

	t.Log("It should return the stored instances for all keys")
	if actualValue1 != expectedValue1 || actualValue2 != expectedValue2 {
		t.Errorf("\tExpected %#v, %#v but was %#v, %#v", expectedValue1, expectedValue2, actualValue1, actualValue2)
	}
}

func Test_when_3_keys_are_added_where_each_is_less_than_the_last(t *testing.T) {
	expectedValue1 := "key1:value1"
	expectedValue2 := "key2:value2"
	expectedValue3 := "key3:value3"

	bst := NewBinarySearchTree()
	bst.Set("key3", expectedValue3)
	bst.Set("key2", expectedValue2)
	bst.Set("key1", expectedValue1)

	actualValue3 := bst.Get("key3")
	actualValue2 := bst.Get("key2")
	actualValue1 := bst.Get("key1")

	t.Log("It should return the stored instances for all keys")
	if actualValue1 != expectedValue1 || actualValue2 != expectedValue2 || actualValue3 != expectedValue3 {
		t.Errorf("\tExpected %#v, %#v, %#v but was %#v, %#v, %#v", expectedValue1, expectedValue2, expectedValue3, actualValue1, actualValue2, actualValue3)
	}
}

func Test_when_multiple_unordered_keys_are_set_and_retrieved(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Set("key3", 3)
	bst.Set("key0", 0)
	bst.Set("key10", 10)
	bst.Set("key5", 5)
	bst.Set("key1", 1)
	bst.Set("key9", 9)
	bst.Set("key15", 15)
	bst.Set("key2", 2)
	bst.Set("key7", 7)
	bst.Set("key4", 4)

	actualValues := []interface{}{
		bst.Get("key3"),
		bst.Get("key0"),
		bst.Get("key10"),
		bst.Get("key5"),
		bst.Get("key1"),
		bst.Get("key9"),
		bst.Get("key15"),
		bst.Get("key2"),
		bst.Get("key7"),
		bst.Get("key4"),
	}

	expectedValues := []interface{}{3, 0, 10, 5, 1, 9, 15, 2, 7, 4}

	t.Log("It should return the stored instances for all keys")
	if reflect.DeepEqual(actualValues, expectedValues) == false {
		t.Errorf("\tExpected %#v but was %#v", expectedValues, actualValues)
	}
}

func Test_when_retrieving_an_overwritten_key(t *testing.T) {
	expectedValue := "expected"

	bst := NewBinarySearchTree()
	bst.Set("key1", "value to be overwritten")
	bst.Set("key1", expectedValue)

	actualValue := bst.Get("key1")

	t.Log("It should return the overwritten value")
	if actualValue != expectedValue {
		t.Errorf("\tExpected %#v but was %#v", expectedValue, actualValue)
	}
}
