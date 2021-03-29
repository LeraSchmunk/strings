//Пакет для чтения данных из файла(строки)

package datafile

import (
	"bufio"
	"os"
)

//GetFloats читает значение string из каждой строки файла
func GetStrings(fileName string) ([]string, error) {
	var lines []string             //Объявление возвращаемого массива
	file, err := os.Open(fileName) //Открытие файла с переданным именем
	//Если при открытии файла произошла ошибка, сообщить о ней и завершить работу
	if err != nil {
		return nil, err //Возвращает nill вместо сегмента
	}
	//i := 0                            //Переменная для хранения индекса, по которому должно выполняться присваивание
	scanner := bufio.NewScanner(file) //Для файла создается новое значение
	for scanner.Scan() {              //Читает строку из файла
		line := scanner.Text() //Выводит строку
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	} //Если при закрытии файла произошла ошибка, сообщить о ней и завершить работу

	if scanner.Err() != nil {
		return nil, scanner.Err()
	} //Если при сканировании файла произошла ошибка, сообщить о ней и завершить работу
	return lines, nil //Если выполнение дошло до этой точки, значит, ошибок не было, поэтому программа возвращает массив чисел и значение ошибки nil
}
