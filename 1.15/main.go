package main

// Небольшой фрагмент кода — проблемы и решение

// Рассмотреть следующий код и ответить на вопросы:
// к каким негативным последствиям он может привести и как это исправить?

// Приведите корректный пример реализации.

var justString string

func createHugeString(size int) string {
	bytes := make([]byte, size)
	for i := 0; i < size; i++ {
		bytes[i] = 'A'
	}
	return string(bytes)
}

func someFunc() {
	v := createHugeString(1 << 10)
	tmp := make([]byte, 100)
	copy(tmp, v[:100])
	justString = string(tmp)
}

func main() {
	someFunc()
}

//Вопрос: что происходит с переменной justString?
//Ответ:
//1. justString - это глобальная строковая переменная
//2. В someFunc() она получает значение - строку из 100 символов 'A'
//3. Благодаря использованию временного буфера и copy() избегаем утечки памяти
//4. Исходная большая строка может быть освобождена сборщиком мусора
