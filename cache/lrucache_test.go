package cache

import (
	"log"
	"os"
	"reflect"
	"search-nearest-places/models"
	"testing"
)

var (
	blankLRUCache    *LRUCache
	lruCacheWithData *LRUCache
	placeWithData    *models.Places
	placeWithData2   *models.Places
)

func TestMain(m *testing.M) {
	blankLRUCache = New(2)

	lruCacheWithData = New(2)
	placeWithData = &models.Places{Restaurents: []models.PlaceInfo{{Title: "title"}}}
	placeWithData2 = &models.Places{Restaurents: []models.PlaceInfo{{Title: "title2"}}}
	lruCacheWithData.Put("London", *placeWithData)

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {

	tests := []struct {
		name          string
		cap           int
		wantCacheSize int
	}{
		{
			name:          "cache capacity var should be 10, and initial length of map and list is 0",
			cap:           10,
			wantCacheSize: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := New(tt.cap)
			//cache is initalized with 10
			if got.capacity != tt.cap {
				t.Errorf("got cache size = %d, want %d", got.capacity, tt.cap)
			}
			//initial size is 0
			mapLen := len(got.elementMap)
			listLen := got.list.Len()
			if mapLen != 0 && listLen != 0 {
				t.Errorf("1. want: map size %d got: %d, 2. want: list size %d got: %d", tt.wantCacheSize, mapLen, tt.wantCacheSize, listLen)
			}
		})
	}
}

func TestLRUCache_Get(t *testing.T) {

	tests := []struct {
		name          string
		location      string
		lruCache      *LRUCache
		newCacheValue *models.Places
		want          *models.Places
	}{
		{
			name:          "if key is not found, it should return nil",
			location:      "London",
			lruCache:      blankLRUCache,
			newCacheValue: placeWithData2,
			want:          nil,
		},
		{
			name:     "if key is found, it should return corresponding value",
			location: "London",
			lruCache: lruCacheWithData,
			want:     placeWithData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.lruCache.Get(tt.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRUCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Put(t *testing.T) {

	tests := []struct {
		name                     string
		location                 string
		lruCache                 *LRUCache
		newCacheValueForLocation *models.Places
		want                     *models.Places
	}{
		{
			name:                     "data should be updated for london, as cache have value of london",
			location:                 "London",
			lruCache:                 lruCacheWithData,
			newCacheValueForLocation: placeWithData2,
			want:                     placeWithData2,
		},
		{
			name:                     "data should be inserted, if key is not available",
			location:                 "Berlin",
			lruCache:                 lruCacheWithData,
			newCacheValueForLocation: placeWithData,
			want:                     placeWithData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.lruCache.Put(tt.location, *tt.newCacheValueForLocation)

			if got := tt.lruCache.Get(tt.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LRUCache.Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLRUCache_Put_Capacity(t *testing.T) {

	lruCacheWithData.Put("Berlin", *placeWithData2)

	tests := []struct {
		name                     string
		location                 string
		lruCache                 *LRUCache
		newCacheValueForLocation *models.Places
		want                     *models.Places
	}{
		{
			name:                     "if capacity is reached, it should delete the least recently accessed and add recent one in front of list",
			location:                 "Stolkholm",
			lruCache:                 lruCacheWithData,
			newCacheValueForLocation: placeWithData2,
			want:                     placeWithData2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.lruCache.Put(tt.location, *tt.newCacheValueForLocation)

			if got := tt.lruCache.Get(tt.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match for Stolkholm :: LRUCache.Get() = %v, want %v", got, tt.want)
			}

			if got := tt.lruCache.Get("London"); got != nil {
				log.Println(got)
				t.Errorf("London data should be nil :: LRUCache.Get() = %v, want %v", got, nil)
			}
		})
	}
}
