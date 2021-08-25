package bloom

import (
	"github.com/spaolacci/murmur3"
	"hash"
	"log"
	"math"
	"sync"
)

type Filter struct {
	lock   *sync.RWMutex
	bytes  []byte
	hashes []hash.Hash64
	n      uint32
	m      uint32
	k      uint32
	p      float64
}

func (f *Filter) Add(bs []byte) {
	f.lock.Lock()
	defer f.lock.Unlock()
	indexs := hashes(&f.hashes, bs, f.m)
	for _, index := range indexs {
		var bits = f.bytes[index/8]
		var t byte = (1 << (7 - (index % 8)))
		f.bytes[index/8] = bits | t
	}
}

func (f *Filter) Contain(bs []byte) bool {
	f.lock.RLock()
	defer f.lock.RUnlock()
	indexs := hashes(&f.hashes, bs, f.m)
	for _, index := range indexs {
		var bits = f.bytes[index/8]
		var t byte = (1 << (7 - (index % 8)))
		if (bits & t) == 0 {
			return false
		}
	}
	return true
}

func (f *Filter) BitCount() uint32 {
	return BitCount(f.bytes)
}

func New(n uint32, p float64) Filter {
	filter := Filter{}
	m := OptimalNumOfBits(n, p)
	k := OptimalNumOfHashFunctions(n, m)
	filter.n = n
	filter.p = p
	filter.m = m
	filter.k = k
	filter.lock = &sync.RWMutex{}
	filter.hashes = createHashes(k)
	filter.bytes = make([]byte, uint32(math.Ceil(float64(m)/8)))
	return filter
}

func createHashes(k uint32) []hash.Hash64 {
	var seed uint32 = 2
	var hashes = make([]hash.Hash64, k)
	for i := 0; i < int(k); i++ {
		hashes[i] = murmur3.New64WithSeed(uint32((seed << i) - 1))
	}
	return hashes
}

func hashes(hashes *[]hash.Hash64, bs []byte, length uint32) []uint32 {
	var indexs = make([]uint32, len(*hashes))
	for i := 0; i < len(*hashes); i++ {
		_, err := (*hashes)[i].Write(bs)
		if err != nil {
			log.Panic("murmur hash error", err)
		}
		indexs[i] = uint32((*hashes)[i].Sum64() % uint64(length))
		(*hashes)[i].Reset()
	}
	return indexs
}
