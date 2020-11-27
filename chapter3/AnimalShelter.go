package chapter3

type AnimalShelter struct {
	dogQueue    *queue
	catQueue    *queue
	isDogOldest bool
	currAge     int
}

type Animal int

const (
	Dog Animal = iota
	Cat
)

func (animalShelter *AnimalShelter) enqueue(name string, species Animal) {
	animalShelter.currAge++

	if species == Dog {
		animalShelter.dogQueue.Enqueue(name, animalShelter.currAge)
	} else {
		animalShelter.catQueue.Enqueue(name, animalShelter.currAge)
	}
}

func (animalShelter *AnimalShelter) dequeueAll() {
	if animalShelter.isDogOldest {
		animalShelter.dogQueue.Dequeue()
	} else {
		animalShelter.catQueue.Dequeue()
	}

	animalShelter.findOldest()
}

func (animalShelter *AnimalShelter) dequeueDog() {
	animalShelter.dogQueue.Dequeue()
	animalShelter.findOldest()
}

func (animalShelter *AnimalShelter) dequeueCat() {
	animalShelter.catQueue.Dequeue()
	animalShelter.findOldest()
}

func (animalShelter *AnimalShelter) findOldest() {
	if animalShelter.dogQueue.isEmpty() {
		animalShelter.isDogOldest = false
		return
	}

	if animalShelter.catQueue.isEmpty() {
		animalShelter.isDogOldest = true
		return
	}

	dog := animalShelter.dogQueue.Peek
	cat := animalShelter.catQueue.Peek

	if dog().age < cat().age {
		animalShelter.isDogOldest = true
	} else {
		animalShelter.isDogOldest = false
	}
}
