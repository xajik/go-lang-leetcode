package main

import (
	"fmt"
	"math"
	"container/heap"
)

//POI for ex. restaurants 
type POI struct {
	name string
	coordinates Coordinates
}

//Coordinates pair of lat/lng
type Coordinates struct {
	longtitute float64
	latitude float64
}

//DistancePOI heap data structure
type DistancePOI struct {
	poi POI
	distance float64
	index int // The index of the item in the heap.
}

//User current
var myLocation = Coordinates{16.6, 19.6}

//Server API 
func getPoiPlaces() []POI {
	places := []POI{ 
		POI{"Prata", Coordinates{6.3, 8.99} }, 
		POI{"Raffeles", Coordinates{8.3, 9.99}},
		POI{"KFC", Coordinates{10.3, 15.9}}, 
		POI{"MD", Coordinates{16.3, 18.99}}, 
		POI{"IndianExpress", Coordinates{29.3, 98.99}},
		POI{"ChikenMassala", Coordinates{129.3, 498.99}},
		POI{"Waffles", Coordinates{1293.3, 4198.99}},
	}
	return places
}

//Util -> return distance from 2 pair of lat/lng
func getDistance(from Coordinates, to Coordinates) float64 {
	longtitute := from.longtitute + to.longtitute
	latitude := from.latitude + to.latitude
	return math.Sqrt(math.Pow(longtitute, 2) + math.Pow(latitude, 2))
}

//															Map IMPL all places
func findAllPoi(places []POI) []DistancePOI {
	closestPoi := make([]DistancePOI, len(places))
	for i, p := range places {
		distance := getDistance(myLocation, p.coordinates)
		closestPoi = append(closestPoi, DistancePOI{p, distance, i})
	}

	return closestPoi
}

//Heap Sort is O(nLogn)
func heapSort(items []DistancePOI) {
	lenght := len(items)
	for i := lenght / 2 - 1; i >= 0 ; i-- {
		heapify(items, lenght, i)
	}

	for i:= lenght - 1; i >= 0; i-- {
		items[0], items[i] = items[i], items[0]
		heapify(items, i, 0)
	}
}

//heapify is O(Logn)
func heapify(items []DistancePOI, lenght int, i int) {
	largest := i
	left := 2 * i + 1
	right := 2 * i + 2
	if(left < lenght && items[left].distance < items[largest].distance) {
		largest = left
	}
	if(right < lenght && items[right].distance < items[largest].distance) {
		largest = right
	}
	if(largest != i) {
		items[largest], items[i] = items[i], items[largest]
		heapify(items, lenght, largest)
	}
}

func testDistanceToPlacesSimple() {
	fmt.Println()
	fmt.Println("Closest POI")
	places := findAllPoi(getPoiPlaces())
	heapSort(places)
	for i , p := range places {
		fmt.Printf("Distance[%d] %f to name %s [ %f, %f ] \n", i, p.distance, p.poi.name, p.poi.coordinates.latitude, p.poi.coordinates.longtitute)
	}
}

//PriorityQueue default implementation - Heap data struct https://golang.org/pkg/container/heap/
type PriorityQueue []*DistancePOI

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

//Push item to priority Q
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*DistancePOI)
	item.index = n
	*pq = append(*pq, item)
}

//Pop item from priority Q
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *DistancePOI, value POI, priority float64) {
	item.poi = value
	item.distance = priority
	heap.Fix(pq, item.index)
}

func findAllPoiHeap(places []POI) PriorityQueue {
	pq := make(PriorityQueue, len(places))
	for i, p := range places {
		distance := getDistance(myLocation, p.coordinates)
		pq[i] = &DistancePOI{p, distance, i}
	}
	heap.Init(&pq)
	return pq
}

func testDistanceToPlacesHeap() {
	fmt.Println()
	fmt.Println("Closest POI, priority Q")
	places := findAllPoiHeap(getPoiPlaces())
	for i , p := range places {
		fmt.Printf("Distance[%d] %f  to name %s [ %f, %f ] \n", i, p.distance, p.poi.name, p.poi.coordinates.latitude, p.poi.coordinates.longtitute)
	}
}
