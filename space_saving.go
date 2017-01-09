package space_saving

import "fmt"

type SpaceSaving struct {
	streamSummary  *StreamSummary
	maxCountersLen int
}

func New(maxCountersLen int) *SpaceSaving {
	return &SpaceSaving{maxCountersLen: maxCountersLen, streamSummary: NewStreamSummary()}
}

func (ss *SpaceSaving) Process(key string) {
	if ss.streamSummary.HasKey(key) {
		ss.streamSummary.Increment(key)
		return
	}
	if ss.streamSummary.Len() < ss.maxCountersLen {
		ss.streamSummary.Add(key)
		return
	}
	ss.streamSummary.ReplaceWith(key)
}

func (ss *SpaceSaving) Report(k int) {
	i := 0
	for bucketElement := ss.streamSummary.bucketList.Front(); bucketElement != nil && i < k; bucketElement = bucketElement.Next() {
		bucket := bucketElement.Value.(*Bucket)
		for counterElement := bucket.counterList.Front(); counterElement != nil; counterElement = counterElement.Next() {
			fmt.Printf("%s = %d\n", counterElement.Value.(*Counter).key, bucket.value)
			i++
		}
	}
}
