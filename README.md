# golang-united-school-homework-8

In the task you have to write console application for managing users list. It should accept there types of operation:
`add
list
findById
remove`
Users list should be stored in the json file. When you start your application and tries to perform some operations, existing file should be used or new one should be created if it does not exist.
Example of the json file (users.json):
`[{id: "1", email: "test@test.com", age: 31}, {id: "2", email: "test2@test.com", age: 41}]`
In the `main.go` file you can find a function called Perform(args Arguments, writer io.Writer) error.
You have to call this function from the main function and pass arguments from the console and os.Stdout stream. Perform function body you have to write by yourself :).
Arguments - is a `map[string]string` with the following fields:
`id, item, operation and fileName`. Create a separate type for the arguments map: `type Arguments map[string]string` to prevent duplication of `map[string]string`. Unit tests use Arguments type also.
Arguments should be passed via console flags:
`./main.go -operation "add" -item ‘{"id": "1", "email": "email@test.com", "age": 23}’ -fileName "users.json"`
`-operation`, `-item`and `-fileName` are console flags. To parse them and build map you can take a Look at "flag" package: https://golang.org/pkg/flag/.
Pay attention that `-fileName` flag should be provided every time with the name of file where you store users!

#### Getting list of items:
Application has to retrieve list from the users.json file and print it to the `io.Writer` stream. Use writer from the argument of Perform function to print the result! It is important for passing unit tests. It can be smth like `writer.Write(bytes)`
File content: `[{"id": "1", "email": "email@test.com", "age": 23}]`
Command: `./main -operation "list" -fileName "users.json"` (main is bult go application. Binary file after go build command)
Output to the console: `[{"id": "1", "email": "email@test.com", "age": 23}]`
If file is empty then nothing should be printed to the console.
**Errors:** 
1. If `-operation` flag is missing, then error `-operation` flag has to be specified" has to be returned from Perform function. Package `errors` can be used for creating errors (https://golang.org/pkg/errors/).
2. If `—fileName` flag is missing, then error "-fileName flag has to be specified" has to be returned from Perform function.
3. If operation can not be handled, for example "abc" operation, then "Operation abcd not allowed!" error has to be return from the Perform function
All cases are covered by unit tests. If you want to be sure your solution works correct, just start `go test -v` command in the root of the repo
#### Adding new item:
For adding new item to the array inside users.json file, application should provide the following cmd command:
`./main -operation "add" -item '{"id": "1", "email": "email@test.com", "age": 23}' -fileName "users.json"`
`-item` - valid json object with the id, email and age fields
**Errors:**
1. All errors about operation and fileName flags mentioned above
2. If `-item` flag is not provided Perform function should return error "-item flag has to be specified"

#### Remove user
Application should allow removing user with the following command:
`./main -operation "remove" -id "2" -fileName "users.json"`
If user with id `"2"`, for example, does not exist, Perform functions should print message to the `io.Writer` "Item with id 2 not found".
If user with specified id exists, it should be removed from the users.json file.
**Errors:**
1. All errors about operation and fileName flags mentioned above
2. If `-id` flag is not provided error "-id flag has to be specified" should be returned from Perform function

#### Find by id
Application should allow finding user by id with the following command:
`./main -operation "findById" -id "1" -fileName "users.json"`
If user with specified id does not exist in the users.json file, then empty string has to be written to  the `io.Writer`
If user exists, then json object should be written in `io.Writer`
**Errors:**
1. All errors about operation and fileName flags mentioned above
2. If `-id` flag is not provided error "-id flag has to be specified" should be returned from Perform function
All cases of the task are covered by unit tests, So, you can check your solution during the implementation.

### Useful info:
1. For opening and creating file use `os` package and `OpenFile` function https://golang.org/pkg/os/
2. To simply read file use `ioutil` package and `ReadAll` function https://golang.org/pkg/io/ioutil/
3. To convert json string to the object use `encoding/json` package and `Unmarshal` function: https://golang.org/pkg/encoding/json/
4. To convert json array or object to string use json package and `Marshal` function.
5. Go does not have throw operator and try catch statement. Instead, it has multi return and allows returning error from a function: `func () (User, error) {}`
Take a look: https://medium.com/@hussachai/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76
6. If you receive error in Perform function, just call panic function for exiting the execution and printing error

**Note that flags and operations names should be the same as mentioned above or unit tests will never pass.**



В задаче необходимо написать консольное приложение для управления списком пользователей. Он должен принимать следующие типы операций:
`add
list
findById
remove`
Список пользователей должен храниться в json-файле. Когда вы запускаете приложение и пытаетесь выполнить некоторые операции, следует использовать существующий файл или создать новый, если он не существует.
Пример файла json (users.json):
`[{id: "1", email: "test@test.com", age: 31}, {id: "2", email: "test2@test.com", age: 41}]`
В файле `main.go` вы можете найти функцию с именем Perform(args Arguments, writer io. Writer) ошибка.
Вы должны вызвать эту функцию из основной функции и передать аргументы из консоли и потока os.Stdout. Тело выполнения функции вы должны написать сами :).

Arguments - это `map[string]string` со следующими полями:
`id, item, operation and fileName`. Создайте отдельный тип для карты аргументов: `type Arguments map[string]string`, чтобы предотвратить дублирование `map[string]string`. Модульные тесты также используют тип Arguments.
Аргументы должны передаваться через консольные флаги:
`./main.go -operation "add" -item '{"id": "1", "email": "email@test.com", "age": 23}' -fileName "users.json"`
`-operation`, `-item` и `-fileName` являются консольными флагами. Чтобы разобрать их и построить карту, вы можете взглянуть на пакет "flag": https://golang.org/pkg/flag/.
Обратите внимание, что флаг `-fileName` должен указываться каждый раз вместе с именем файла, в котором вы храните пользователей!

#### Получение списка предметов:
Приложение должно получить список из файла users.json и распечатать его в поток io.Writer. Используйте средство записи из аргумента функции Perform для печати результата! Это важно для прохождения модульных тестов. Это может быть что-то вроде `writer.Write(bytes)`.
Содержимое файла: `[{"id": "1", "email": "email@test.com", "age": 23}]`.
Команда: `./main -operation "list" -fileName "users.json"` (главным является приложение bult go. Двоичный файл после команды go build)
Вывод в консоль: `[{"id": "1", "email": "email@test.com", "age": 23}]`
Если файл пуст, то в консоль ничего не должно выводиться.

**Ошибки:** 
1. Если флаг `-operation` отсутствует, то из функции Perform должна быть возвращена ошибка `-operation flag has to be specified`. Пакет `errors` можно использовать для создания ошибок (https://golang.org/pkg/errors/).

2. Если флаг `—fileName` отсутствует, то из функции Perform должна быть возвращена ошибка `-fileName flag has to be specified`.

3. Если операция не может быть обработана, например операция `abc`, то должна быть возвращена из функции выполнения ошибка `Operation abcd not allowed!`.

Все случаи покрываются модульными тестами. Если вы хотите убедиться, что ваше решение работает правильно, просто запустите команду go test -v в корне репозитория.

#### Добавление нового элемента:
Для добавления нового элемента в массив внутри файла users.json приложение должно предоставить следующую команду cmd:
`./main -operation "add" -item "{"id": "1", "email": "email@test.com", "age": 23}" -fileName "users.json"`
`-item` - действительный объект json с полями идентификатора, электронной почты и возраста

**Ошибки:**
1. Все ошибки, связанные с операцией и флагами fileName, упомянутыми выше
2. Если флаг `-item` не указан, функция выполнения должна возвращать ошибку `-item flag has to be specified`.

#### Удаление пользователя
Приложение должно разрешать удаление пользователя с помощью следующей команды:
`./main -operation "remove" -id "2" -fileName "users.json"`
Если пользователя с id `"2"`, например, не существует, функции Perform должны вывести на `io.Writer` сообщение `Item with id 2 not found`.
Если пользователь с указанным идентификатором существует, его следует удалить из файла users.json.

**Ошибки:**
1. Все ошибки, связанные с операцией и флагами fileName, упомянутыми выше
2. Если флаг `-id` не указан, из функции Perform должна быть возвращена ошибка `-id flag has to be specified`.

#### Найти по идентификатору
Приложение должно позволять найти пользователя по id с помощью следующей команды:
`./main -operation "findById" -id "1" -fileName "users.json"`
Если пользователь с указанным идентификатором не существует в файле users.json, то в `io.Writer` необходимо записать пустую строку.
Если пользователь существует, то объект json должен быть записан в `io.Writer`.

**Ошибки:**
1. Все ошибки, связанные с операцией и флагами fileName, упомянутыми выше
2. Если флаг `-id` не указан, из функции Perform должна быть возвращена ошибка `-id flag has to be specified`.
Все случаи задачи покрыты модульными тестами, так что вы можете проверить свое решение в процессе реализации.

### Полезная информация:
1. Для открытия и создания файла используйте пакет `os` и функцию `OpenFile` https://golang.org/pkg/os/
2. Чтобы просто прочитать файл, используйте пакет `ioutil` и функцию `ReadAll` https://golang.org/pkg/io/ioutil/
3. Чтобы преобразовать строку json в объект, используйте пакет `encoding/json` и функцию `Unmarshal`: https://golang.org/pkg/encoding/json/
4. Чтобы преобразовать массив или объект json в строку, используйте пакет json и функцию `Marshal`.
5. В Go нет оператора throw и оператора try catch. Вместо этого он имеет множественный возврат и позволяет возвращать ошибку из функции: `func () (User, error){}`.
Взглянем: https://medium.com/@hussachai/error-handling-in-go-a-quick-opinionated-guide-9199dd7c7f76
6. Если вы получаете ошибку в функции выполнения, просто вызовите функцию паники для выхода из выполнения и печати ошибки

**Обратите внимание, что флаги и имена операций должны быть такими же, как указано выше, иначе модульные тесты никогда не пройдут.**