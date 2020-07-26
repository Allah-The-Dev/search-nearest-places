package cache

import (
	"container/list"

	"search-nearest-places/models"
)

//Cache ... generic cache type
type Cache interface {
	Get(string) *models.Places
	Put(string, models.Places)
}

//LRUCache ...  using golang container/list - a doubly linked list package
type LRUCache struct {
	capacity   int
	list       *list.List
	elementMap map[string]*list.Element
}

//PlaceInfoCache ... store places data mapped to location name
type PlaceInfoCache struct {
	LocationName string
	POIData      models.Places
}

//New ... method to create a LRU new cache
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
func (cache *LRUCache) Put(locationName string, newPOIPlaces models.Places) {

	if node, ok := cache.elementMap[locationName]; ok {

		node.Value.(*list.Element).Value = PlaceInfoCache{locationName, newPOIPlaces}

		cache.list.MoveToFront(node)
	} else {

		if cache.list.Len() == cache.capacity {

			lastLocationInCache := cache.list.Back().Value.(*list.Element).Value.(PlaceInfoCache).LocationName

			delete(cache.elementMap, lastLocationInCache)

			cache.list.Remove(cache.list.Back())
		}

		newNode := &list.Element{
			Value: PlaceInfoCache{
				LocationName: locationName,
				POIData:      newPOIPlaces,
			},
		}

		nodePtr := cache.list.PushFront(newNode)

		cache.elementMap[locationName] = nodePtr
	}
}
