package cache

import (
	"container/list"

	"search-nearest-places/models"
)

//LRUCache ...  using golang container/list - a doubly linked list package
type LRUCache struct {
	capacity   int
	list       *list.List
	elementMap map[string]*list.Element
}

//PlaceInfoCache ... store places data mapped to location name
type PlaceInfoCache struct {
	Location string
	POIData  models.Places
}

//New ... method to create a new cache
func New(cap int) *LRUCache {
	return &LRUCache{
		capacity:   cap,
		list:       new(list.List),
		elementMap: make(map[string]*list.Element, cap),
	}
}

//Get ... get the element value via key if it exists
func (cache *LRUCache) Get(location string) *models.Places {

	if node, ok := cache.elementMap[location]; ok {

		val := node.Value.(*list.Element).Value.(PlaceInfoCache).POIData

		cache.list.MoveToFront(node)
		return &val
	}
	return nil
}

//Put ... put item in list and map to a hashmap if not exist
//if exist move it to front
func (cache *LRUCache) Put(location string, newPOIPlaces models.Places) {

	if node, ok := cache.elementMap[location]; ok {

		node.Value.(*list.Element).Value = PlaceInfoCache{location, newPOIPlaces}

		cache.list.MoveToFront(node)
	} else {

		if cache.list.Len() == cache.capacity {

			lastLocationInCache := cache.list.Back().Value.(*list.Element).Value.(PlaceInfoCache).Location

			delete(cache.elementMap, lastLocationInCache)

			cache.list.Remove(cache.list.Back())
		}

		newNode := &list.Element{
			Value: PlaceInfoCache{
				Location: location,
				POIData:  newPOIPlaces,
			},
		}

		nodePtr := cache.list.PushFront(newNode)

		cache.elementMap[location] = nodePtr
	}
}
