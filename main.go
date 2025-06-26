package main

import (
	"fmt"
)

type Coordinat struct {
	next *Coordinat
	x, y int
}

func Initialization(x, y int) *Coordinat {
	var head Coordinat
	head.x, head.y = x, y
	return &head
}

func IsBigger(first, second *Coordinat) bool {
	if first.x > second.x {
		return true
	} else if first.x == second.x {
		if first.y > second.y {
			return true
		}
	}
	return false
}

func Add(head *Coordinat, x, y int) *Coordinat {
	var to_add Coordinat
	to_add.x, to_add.y = x, y
	if IsBigger(&to_add, head) {
		to_add.next = head
		return &to_add
	}
	prev := head
	elem := head
	for elem != nil {
		if IsBigger(&to_add, elem) {
			prev.next = &to_add
			to_add.next = elem
			return head
		}
		prev = elem
		elem = elem.next
	}
	prev.next = &to_add
	return head
}

func Find(head, elem *Coordinat) (*Coordinat, int) {
	count := 0
	now_elem := head
	// _, _len := Length(head)
	// fmt.Println(_len)
	for now_elem != nil {
		// fmt.Println("Infinity")
		if now_elem.x == elem.x && now_elem.y == elem.y {
			count++
		}
		now_elem = now_elem.next
	}
	return head, count
}

func NewDelete(head *Coordinat, x, y, count int) *Coordinat {
	var prev *Coordinat
	var new_elem Coordinat
	now_count := 0
	if head.next == nil {
		fmt.Println("Can`t delete head.\nUse Clear")
		return head
	}
	if count <= 0 {
		fmt.Println("Wrong count")
		return head
	}
	new_elem.x, new_elem.y = x, y
	head, count_of_elem := Find(head, &new_elem)
	if count > count_of_elem {
		fmt.Println("Can`t delete that count of elems. Count so big")
		return head
	}
	prev = nil
	elem := head
	for elem != nil {
		if elem.x == x && elem.y == y && now_count < count {
			if prev == nil {
				head = elem.next
				prev = nil
				elem = head
				now_count++
			} else {
				prev.next = elem.next
				elem = elem.next
				now_count++
				continue
			}
		}
		prev = elem
		elem = elem.next
	}
	if head.x == x && head.y == y && now_count+1 == count {
		return head.next
	}
	return head
}

func Length(head *Coordinat) (*Coordinat, int) {
	elem := head
	_len := 0
	for elem != nil {
		_len++
		elem = elem.next
	}
	return head, _len
}

func Clear(head *Coordinat) *Coordinat {
	return nil
}

func ListPrint(head *Coordinat) *Coordinat {
	elem := head
	for elem != nil {
		fmt.Println(elem.x, elem.y)
		elem = elem.next
	}
	return head
}

func Intersection(first_head, second_head *Coordinat) *Coordinat {
	var third_head *Coordinat
	elem := first_head
	for elem != nil {
		_, second_count := Find(second_head, elem)
		_, first_count := Find(first_head, elem)
		_, third_count := Find(third_head, elem)
		if third_count == 0 {
			count := min(first_count, second_count)
			for i := 0; i < count; i++ {
				if third_head == nil {
					third_head = Initialization(elem.x, elem.y)
				} else {
					third_head = Add(third_head, elem.x, elem.y)
				}
			}
		}
		elem = elem.next

	}
	return third_head
}

//	func main() {
//		var first_head, second_head *Coordinat
//		var str_count, str_x, str_y string
//		fmt.Println("Enter number of elems of first array")
//		fmt.Scan(&str_count)
//		count, err := strconv.Atoi(str_count)
//		if err != nil {
//			panic(err)
//		}
//		for i := 0; i < count; i++ {
//			fmt.Println("Enter coords on different lines ")
//			fmt.Scan(&str_x)
//			fmt.Scan(&str_y)
//			x, err_x := strconv.Atoi(str_x)
//			y, err_y := strconv.Atoi(str_y)
//			if err_x != nil {
//				panic(err_x)
//			}
//			if err_y != nil {
//				panic(err_y)
//			}
//			if i == 0 {
//				first_head = Initialization(x, y)
//			} else {
//				first_head = Add(first_head, x, y)
//			}
//		}
//		first_head = ListPrint(first_head)
//		fmt.Println("Enter number of elems of second array")
//		fmt.Scan(&str_count)
//		count, err = strconv.Atoi(str_count)
//		if err != nil {
//			panic(err)
//		}
//		for i := 0; i < count; i++ {
//			fmt.Println("Enter coords on different lines ")
//			fmt.Scan(&str_x)
//			fmt.Scan(&str_y)
//			x, err_x := strconv.Atoi(str_x)
//			y, err_y := strconv.Atoi(str_y)
//			if err_x != nil || err_y != nil {
//				panic(err_x)
//			}
//			if i == 0 {
//				second_head = Initialization(x, y)
//			} else {
//				second_head = Add(second_head, x, y)
//			}
//		}
//		second_head = ListPrint(second_head)
//		fmt.Println("After delete 1")
//		first_head = NewDelete(first_head, 2, 2, 2)
//		ListPrint(first_head)
//		fmt.Println("After delete 2")
//		second_head = NewDelete(second_head, 2, 2, 2)
//		ListPrint(second_head)
//		fmt.Println("Result:")
//		result_head := Intersection(first_head, second_head)
//		_ = ListPrint(result_head)
//	}
func main() {
	fmt.Println("---------------- TEST 1------------------------:")
	fmt.Println("---------------- СФОРМИРОВАЛИ СПИСОК------------------------:")
	var second_head1 *Coordinat

	second_head1 = Initialization(2, 2)

	for i := 0; i < 6; i++ {
		second_head1 = Add(second_head1, 2, 2)
	}
	second_head1 = ListPrint(second_head1)
	var _count int
	el := Initialization(2, 2)
	second_head1, _count = Find(second_head1, el)
	fmt.Println("------------ ПОИСК В СПИСКЕ (2,2) ----------------------------:")
	fmt.Println("Find:", _count)

	fmt.Println("?????????????????? УДАЛЕНИЕ (2,2) в количестве =7-----------------------------:")
	second_head1 = NewDelete(second_head1, 2, 2, 7)
	_ = ListPrint(second_head1)

	fmt.Println("---------------- TEST 2------------------------:")
	fmt.Println("---------------- СФОРМИРОВАЛИ СПИСОК------------------------:")
	second_head1 = Initialization(2, 2)

	for i := 0; i < 6; i++ {
		second_head1 = Add(second_head1, i/2+1, i/2+1)
	}
	second_head1 = ListPrint(second_head1)

	el = Initialization(2, 2)
	second_head1, _count = Find(second_head1, el)
	fmt.Println("------------ ПОИСК В СПИСКЕ (2,2) ----------------------------:")
	fmt.Println("Find:", _count)

	fmt.Println("?????????????????? УДАЛЕНИЕ (2,2) в количестве =2-----------------------------:")
	second_head1 = NewDelete(second_head1, 2, 2, 2)
	_ = ListPrint(second_head1)

	fmt.Println("---------------- TEST 3------------------------:")
	fmt.Println("---------------- СФОРМИРОВАЛИ СПИСОК------------------------:")
	second_head1 = Initialization(2, 2)

	for i := 0; i < 6; i++ {
		if i > 3 {
			second_head1 = Add(second_head1, 2, 2)

		} else {
			second_head1 = Add(second_head1, 1, 1)

		}
	}
	second_head1 = ListPrint(second_head1)

	el = Initialization(2, 2)
	second_head1, _count = Find(second_head1, el)
	fmt.Println("------------ ПОИСК В СПИСКЕ (2,2) ----------------------------:")
	fmt.Println("Find:", _count)

	fmt.Println("?????????????????? УДАЛЕНИЕ (2,2) в количестве =3-----------------------------:")
	second_head1 = NewDelete(second_head1, 2, 2, 3)
	_ = ListPrint(second_head1)

	fmt.Println("---------------- TEST 4------------------------:")
	fmt.Println("---------------- СФОРМИРОВАЛИ СПИСОК------------------------:")
	second_head1 = Initialization(2, 2)

	for i := 0; i < 6; i++ {

		second_head1 = Add(second_head1, i/2*3+2, i/2*3+2)

	}
	second_head1 = ListPrint(second_head1)

	el = Initialization(2, 2)
	second_head1, _count = Find(second_head1, el)
	fmt.Println("------------ ПОИСК В СПИСКЕ (2,2) ----------------------------:")
	fmt.Println("Find:", _count)

	fmt.Println("?????????????????? УДАЛЕНИЕ (2,2) в количестве =3-----------------------------:")
	second_head1 = NewDelete(second_head1, 8, 8, 1)
	second_head1 = NewDelete(second_head1, 2, 2, 1)
	second_head1 = NewDelete(second_head1, 5, 5, 1)
	_ = ListPrint(second_head1)
	fmt.Println("---------------- TEST 5------------------------:")
	fmt.Println("---------------- СФОРМИРОВАЛИ СПИСОК------------------------:")

	for i := 0; i < 6; i++ {

		second_head1 = Add(second_head1, i/2*3+2, i/2*3+2)

	}
	_ = ListPrint(second_head1)
	fmt.Println("------------ ПОИСК В СПИСКЕ (2,2) ----------------------------:")
	fmt.Println("Find:", _count)

}
