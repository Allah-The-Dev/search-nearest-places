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
	//if key exist in map
	if node, ok := cache.elementMap[location]; ok {
		// fetch vale
		val := node.Value.(*list.Element).Value.(PlaceInfoCache).POIData
		//move itemto front in list
		cache.list.MoveToFront(node)
		return &val
	}
	return nil
}

//Put ... put item in list and map to a hashmap if not exist
//if exist move it to front
func (cache *LRUCache) Put(location string, poiPlaces models.Places) {
	//if value exist 1. update the value 2. move it to front
	if node, ok := cache.elementMap[location]; ok {
		//update the value in list node
		node.Value.(*list.Element).Value = poiPlaces
		//move it to the first node
		cache.list.MoveToFront(node)
	} else {
		//check size and delete the last one
		if cache.list.Len() == cache.capacity {
			//get the last key
			lastLocationInCache := cache.list.Back().Value.(*list.Element).Value.(PlaceInfoCache).Location
			//delete this key from element map
			delete(cache.elementMap, lastLocationInCache)
			//remove it from cache list
			cache.list.Remove(cache.list.Back())
		}
		//create a new node
		newNode := &list.Element{
			Value: PlaceInfoCache{
				Location: location,
				POIData:  poiPlaces,
			},
		}
		//insert this node in front of list
		nodePtr := cache.list.PushFront(newNode)
		//update map index for fast lookup
		cache.elementMap[location] = nodePtr
	}
}
