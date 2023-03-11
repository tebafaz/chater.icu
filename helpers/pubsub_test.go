package helpers_test

import (
	"fmt"
	"testing"

	"github.com/tebafaz/chater.icu/helpers"
)

// this is frankenstein of a test, but works)
func TestPubSub(t *testing.T) {
	input := []struct {
		topic string
		data  []interface{}
	}{
		{
			topic: "test-topic-1",
			data: []interface{}{
				"test-string",
				1234,
				true,
				false,
				9876,
				"some additional data",
			},
		},
		{
			topic: "test-topic-2",
			data: []interface{}{
				"test-string: topic 2",
				457,
			},
		},
	}

	//create pubsub
	ps := helpers.NewPubsub()

	//make new subscribers
	subs := make([]chan interface{}, len(input))
	for k := range input {
		subs[k] = ps.Subscribe(input[k].topic)
	}

	//check if topic has subscribers
	for _, v := range input {
		if !ps.HaveSubscriber(v.topic) {
			t.Error("no subscriber")
		}
	}

	if ps.HaveSubscriber("non-existent-topic") {
		t.Error("subscriber exist")
	}

	//check if all subscribers are present in topic subscription
	allSubs := ps.ShowAllSubscribers()

	for _, v := range input {
		if _, ok := allSubs[v.topic]; !ok {
			t.Error("no subscriber")
		}
	}

	go func() {
		//publish all input data
		for k := range input {
			for _, v := range input[k].data {
				ps.Publish(input[k].topic, v)
			}
		}
	}()

	//get subscribed data
	for _, v1 := range subs {
		go func(v1 chan interface{}) {
			for v2 := range v1 {
				fmt.Printf("%+v\n", v2)
			}
		}(v1)
	}

	//unsubscribe
	for k, v := range subs {
		ps.Unsubscribe(input[k].topic, v)
		close(v)
	}
}
