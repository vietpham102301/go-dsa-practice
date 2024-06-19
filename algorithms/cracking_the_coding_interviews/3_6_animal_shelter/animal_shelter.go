package main

import "fmt"

type Node struct {
	next *Node
	data Animal
}

type Queue struct {
	head *Node
	last *Node
}

func (q *Queue) enqueue(animal Animal) {
	newNode := &Node{data: animal}
	if q.last != nil {
		q.last.next = newNode
		q.last = newNode
		return
	}
	q.last = newNode
	q.head = newNode

	return
}

func (q *Queue) dequeue() (Animal, bool) {
	if q.head == nil {
		animal := Animal{}
		return animal, false
	}
	res := q.head.data
	q.head = q.head.next
	return res, true
}

func (q Queue) isEmpty() bool {
	return q.head == nil
}

type Animal struct {
	animalType string
	name       string
}

type AnimalShelter struct {
	allAnimalQueue Queue
	dogQueue       Queue
	catQueue       Queue
}

func (s *AnimalShelter) enqueueAnimal(animal Animal) {
	if animal.animalType == "dog" {
		s.dogQueue.enqueue(animal)
	} else if animal.animalType == "cat" {
		s.catQueue.enqueue(animal)
	} else {
		fmt.Println("just can enqueue cat or dog")
		return
	}

	s.allAnimalQueue.enqueue(animal)
}

func (s *AnimalShelter) dequeueAny() (Animal, bool) {
	res, success := s.allAnimalQueue.dequeue()

	if success && res.animalType == "cat" {
		resCat, successCat := s.catQueue.dequeue()
		tempQueue := Queue{}
		for successCat && !s.allAnimalQueue.isEmpty() &&
			s.allAnimalQueue.head.data.animalType != resCat.animalType &&
			s.allAnimalQueue.head.data.name != resCat.name {
			data, success3 := s.allAnimalQueue.dequeue()
			if success3 {
				tempQueue.enqueue(data)
			}
		}
		s.allAnimalQueue.dequeue()
		for !tempQueue.isEmpty() {
			data, suc := tempQueue.dequeue()
			if suc {
				s.allAnimalQueue.enqueue(data)
			}
		}
	} else if success && res.animalType == "dog" {
		resDog, successDog := s.dogQueue.dequeue()
		tempQueue := Queue{}
		for successDog && !s.allAnimalQueue.isEmpty() &&
			s.allAnimalQueue.head.data.animalType != resDog.animalType &&
			s.allAnimalQueue.head.data.name != resDog.name {
			data, success3 := s.allAnimalQueue.dequeue()
			if success3 {
				tempQueue.enqueue(data)
			}
		}
		s.allAnimalQueue.dequeue()
		for !tempQueue.isEmpty() {
			data, suc := tempQueue.dequeue()
			if suc {
				s.allAnimalQueue.enqueue(data)
			}
		}
	}
	return res, success
}

func (s *AnimalShelter) dequeueDog() (Animal, bool) {
	res, success := s.dogQueue.dequeue()
	tempQueue := Queue{}
	for success && !s.allAnimalQueue.isEmpty() && s.allAnimalQueue.head.data.name != res.name {
		res2, success2 := s.allAnimalQueue.dequeue()
		if success2 {
			tempQueue.enqueue(res2)
		}
	}
	s.allAnimalQueue.dequeue()

	for !tempQueue.isEmpty() {
		res3, success3 := tempQueue.dequeue()
		if success3 {
			s.allAnimalQueue.enqueue(res3)
		}
	}

	return res, success
}

func (s *AnimalShelter) dequeueCat() (Animal, bool) {
	res, success := s.catQueue.dequeue()
	tempQueue := Queue{}
	for success && !s.allAnimalQueue.isEmpty() && s.allAnimalQueue.head.data.name != res.name {
		res2, success2 := s.allAnimalQueue.dequeue()
		if success2 {
			tempQueue.enqueue(res2)
		}
	}
	s.allAnimalQueue.dequeue()

	for !tempQueue.isEmpty() {
		res3, success3 := tempQueue.dequeue()
		if success3 {
			s.allAnimalQueue.enqueue(res3)
		}
	}
	return res, success
}

func main() {
	animal1 := Animal{animalType: "dog", name: "pitbull"}
	animal2 := Animal{animalType: "cat", name: "englishcat"}
	animal3 := Animal{animalType: "dog", name: "longbody-dog"}
	animal4 := Animal{animalType: "cat", name: "euro-cat"}
	animal5 := Animal{animalType: "cat", name: "egypt-cat"}

	animal6 := Animal{animalType: "bear", name: "jungle-bear"}
	animal7 := Animal{animalType: "lion", name: "big-cat"}

	animalShelter := AnimalShelter{}

	animalShelter.enqueueAnimal(animal1)
	animalShelter.enqueueAnimal(animal2)
	animalShelter.enqueueAnimal(animal3)
	animalShelter.enqueueAnimal(animal4)
	animalShelter.enqueueAnimal(animal5)

	animalShelter.enqueueAnimal(animal6)
	animalShelter.enqueueAnimal(animal7)

	fmt.Println(animalShelter.dequeueAny())
	fmt.Println(animalShelter.dequeueCat())
	fmt.Println(animalShelter.dequeueDog())
	fmt.Println(animalShelter.dequeueCat())
	fmt.Println(animalShelter.dequeueCat())

}
