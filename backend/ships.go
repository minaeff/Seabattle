package main

// Флотилия игрока
type Fleet []Ship

// Корабль
type Ship struct {
	ID    int
	Name  string
	Class ClassShip
}

// Класс корабля
type ClassShip struct {
	ID          int
	ClassName   string // линкор и т.д.
	Length      int
	HealthPoint int // для расчета статистики урона
}

/* Описание методов флота и отдельных кораблей,
в планах реализации - создание флота, размещение на игровом поле,
аттака, получение урона, разрушение корабля*/

// Создание флотилии, функция принимает на вход массив кораблей и размещает их в одну флотилию
func (f Fleet) NewFleet(Ships []Ship) Fleet {
	// тут должно быть тело функции
	return f
}

// Атака клетки игрового поля
func (s Ship) Attack(Target Cell) {
	//если клетка не пуста, то наносим урон врагу
	if !Target.IsEmpty {
		//
	} else {
		//иначе отмечаем как неактивную
		Target.IsEmpty = true
	}
}
