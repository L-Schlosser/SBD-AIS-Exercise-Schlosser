package mapred

// todo implement mapreduce

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct{}

// Run implements the MapReduceInterface
func (mr MapReduce) Run(input []string) map[string]int {
	// Channel for mapped key-values
	mapCh := make(chan []KeyValue, len(input)) //buffered channel (can store len(input) items)
	var wg sync.WaitGroup

	// MAP PHASE
	for _, line := range input {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			mapCh <- mr.wordCountMapper(text)
		}(line)
	}

	wg.Wait()
	close(mapCh)

	// COLLECT ALL KEY-VALUES
	intermediate := make(map[string][]int)
	for kvs := range mapCh {
		for _, kv := range kvs {
			intermediate[kv.Key] = append(intermediate[kv.Key], kv.Value)
		}
	}

	// REDUCE PHASE
	result := make(map[string]int)
	var reduceWg sync.WaitGroup
	mu := sync.Mutex{}

	for key, values := range intermediate {
		reduceWg.Add(1)
		go func(k string, v []int) {
			defer reduceWg.Done()
			kv := mr.wordCountReducer(k, v)
			mu.Lock()
			result[kv.Key] = kv.Value
			mu.Unlock()
		}(key, values)
	}

	reduceWg.Wait()
	return result
}

// wordCountMapper splits text into words and emits KeyValue pairs
func (mr MapReduce) wordCountMapper(text string) []KeyValue {
	// Filter special chars and numbers
	re := regexp.MustCompile(`[^a-zA-Z]+`) // keep only ASCII letters
	text = re.ReplaceAllString(text, " ")
	text = strings.ToLower(text)

	words := strings.Fields(text)
	kvs := make([]KeyValue, 0, len(words))
	for _, w := range words {
		kvs = append(kvs, KeyValue{Key: w, Value: 1})
	}
	return kvs
}

// wordCountReducer sums the values for a given key
func (mr MapReduce) wordCountReducer(key string, values []int) KeyValue {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return KeyValue{Key: key, Value: sum}
}
