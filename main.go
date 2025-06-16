package main

import (
	"fmt"
	"strconv"
)

type Coordinat struct {
	next *Coordinat
	x, y int
}

func IsBigger(first, second Coordinat) bool {
	if first.x > second.x {
		return true
	} else if first.x == second.x {
		if first.y > second.y {
			return true
		}
	}
	return false
}

func Initialization(x, y int) Coordinat {
	var head Coordinat
	head.x, head.y = x, y
	return head
}

func NewAdd(head Coordinat, x, y int) Coordinat {
	var to_add Coordinat
	to_add.x, to_add.y = x, y
	if IsBigger(to_add, head) {
		to_add.next = &head
		return to_add
	}
	if head.next == nil {
		if IsBigger(head, to_add) {
			head.next = &to_add
			return head
		} else {
			to_add.next = &head
			return to_add
		}
	} else {
		prev := &head
		elem := head.next
		for elem.next != nil {
			if IsBigger(to_add, *elem) {
				prev.next = &to_add
				to_add.next = elem
				return head
			}
			prev = elem
			elem = elem.next
		}
		if IsBigger(to_add, *elem) {
			prev.next = &to_add
			to_add.next = elem
		} else {
			elem.next = &to_add
		}
	}
	return head
}

func Add(head Coordinat, x, y int) Coordinat {
	var to_add Coordinat
	to_add.x, to_add.y = x, y
	if IsBigger(head, to_add) {
		to_add.next = &head
		return to_add
	}
	if head.next == nil {
		head.next = &to_add
		return head
	} else {
		prev := &head
		elem := head.next
		for elem.next != nil {
			if IsBigger(*elem, to_add) {
				prev.next = &to_add
				to_add.next = elem
				return head
			}
			prev = elem
			elem = elem.next
		}
		if IsBigger(*elem, to_add) {
			prev.next = &to_add
			to_add.next = elem
			return head
		} else {
			elem.next = &to_add
			return head
		}
	}
	return head
}

func Delete(head Coordinat, x, y, count int) Coordinat {
	if head.next == nil {
		fmt.Println("Can`t delete head.\nUse Clear")
	}
	if count <= 0 {
		fmt.Println("Wrong count")
		return head
	}
	elem := head
	prev := head
	now_count := 0
	for elem.next != nil {
		if elem.x == x && elem.y == y && now_count < count {
			prev.next = elem.next
			now_count++
		}
		elem = *elem.next
	}
	return head
}

func Find(head, elem Coordinat) (Coordinat, int) {
	count := 0
	now_elem := head
	for now_elem.next != nil {
		if now_elem.x == elem.x && now_elem.y == elem.y {
			count++
		}
		now_elem = *now_elem.next
	}
	if now_elem == elem {
		count++
	}
	return head, count
}

func Length(head Coordinat) (Coordinat, int) {
	elem := head
	_len := 1
	for elem.next != nil {
		_len++
		elem = *elem.next
	}
	return head, _len
}

func Clear(head Coordinat) *int {
	for head.next.next != nil {
		head.next = head.next.next
	}
	head.next = nil
	return nil
}

func ListPrint(head Coordinat) Coordinat {
	elem := head
	for elem.next != nil {
		fmt.Println(elem.x, elem.y)
		elem = *elem.next
	}
	fmt.Println(elem.x, elem.y)
	return head
}

func Intersection(first_head, second_head Coordinat) Coordinat {
	var third_head Coordinat
	elem := &first_head
	for elem.next != nil {
		_, count := Find(second_head, *elem)
		if count > 0 {
			third_head = Add(third_head, elem.x, elem.y)
		}
		elem = elem.next
	}
	_, count := Find(second_head, *elem)
	if count > 0 {
		third_head = Add(third_head, elem.x, elem.y)
	}
	elem = elem.next
	return third_head
}

func main() {
	var first_head, second_head Coordinat
	var str_count, str_x, str_y string
	fmt.Println("Enter number of elems of first array")
	fmt.Scan(&str_count)
	count, err := strconv.Atoi(str_count)
	if err != nil {
		panic(err)
	}
	for i := 0; i < count; i++ {
		fmt.Println("Enter coords on different lines ")
		fmt.Scan(&str_x)
		fmt.Scan(&str_y)
		x, err_x := strconv.Atoi(str_x)
		y, err_y := strconv.Atoi(str_y)
		if err_x != nil || err_y != nil {
			panic(err_x)
		}
		if i == 0 {
			first_head = Initialization(x, y)
		} else {
			first_head = NewAdd(first_head, x, y)
		}
	}
	first_head = ListPrint(first_head)
	fmt.Println("Enter number of elems of second array")
	fmt.Scan(&str_count)
	count, err = strconv.Atoi(str_count)
	if err != nil {
		panic(err)
	}
	for i := 0; i < count; i++ {
		fmt.Println("Enter coords on different lines ")
		fmt.Scan(&str_x)
		fmt.Scan(&str_y)
		x, err_x := strconv.Atoi(str_x)
		y, err_y := strconv.Atoi(str_y)
		if err_x != nil || err_y != nil {
			panic(err_x)
		}
		if i == 0 {
			second_head = Initialization(x, y)
		} else {
			second_head = NewAdd(second_head, x, y)
		}
	}
	second_head = ListPrint(second_head)
	fmt.Println("Result:")
	result_head := Intersection(first_head, second_head)
	_ = ListPrint(result_head)
}
