package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type (
	Arguments map[string]string

	User struct {
		Id    string `json:"id"`
		Email string `json:"email"`
		Age   int8   `json:"age"`
	}

	opRequirements []string
)

var (
	pOperation = "operation"
	pFileName  = "fileName"
	pItem      = "item"
	pId        = "id"

	opAdd      = "add"
	opList     = "list"
	opFindById = "findById"
	opRemove   = "remove"

	operationsIndex = []string{opAdd, opList, opFindById, opRemove}

	requiredFlag = &pOperation

	additionalRequiredFlags = map[string]opRequirements{
		opAdd:      []string{pFileName, pItem},
		opList:     []string{pFileName},
		opFindById: []string{pFileName, pId},
		opRemove:   []string{pFileName, pId},
	}
)

// Вспомогательные функции
// Функция проверки параметров
func validateParamEntered(args Arguments, params []string) error {
	for _, param := range params {
		value := args[param]
		fmt.Println("value =", value)
		if len(value) == 0 {
			return errors.New("-" + param + " flag has to be specified")
		}
	}

	return nil
}

func validateOpAllowed(operation string) error {
	b := false
	for _, v := range operationsIndex {
		if v == operation {
			b = true
		}
	}

	if !b {
		return errors.New("Operation " + operation + " not allowed!")
	}

	return nil
}

func validateConsequently(args Arguments, reqFlag string, params map[string]opRequirements) error {
	errOperationSpecified := validateParamEntered(args, []string{reqFlag})

	if errOperationSpecified != nil {
		return errOperationSpecified
	}

	errAllowed := validateOpAllowed(args[reqFlag])

	if errAllowed != nil {
		return errAllowed
	}

	op := args[reqFlag]
	errEntered := validateParamEntered(args, params[op])

	if errEntered != nil {
		return errEntered
	}

	return nil
}

func readFileToBuffer(fileName string) ([]byte, error) {
	r, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)

	defer r.Close()

	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func readFileToJSON(fileName string) ([]User, error) {
	buf, err := readFileToBuffer(fileName)

	if err != nil {
		return nil, err
	}

	u := []User{}

	if len(buf) != 0 {
		err := json.Unmarshal(buf, &u)
		if err != nil {
			return nil, err
		}
	}

	return u, nil
}

func parseArgs() Arguments {
	flagNames := [4]string{pOperation, pFileName, pItem, pId}
	flagValues := [4]string{"", "", "", ""}
	args := make(Arguments)

	flag.StringVar(&flagValues[0], flagNames[0], "", "Possible values are 'list', 'add', 'findById', 'remove'")
	flag.StringVar(&flagValues[1], flagNames[1], "", "Path to a JSON file")
	flag.StringVar(&flagValues[2], flagNames[2], "", "Item to add")
	flag.StringVar(&flagValues[3], flagNames[3], "", "User ID to search for")

	flag.Parse()

	for i, flagName := range flagNames {
		if len(flagValues[i]) > 0 {
			args[flagName] = flagValues[i]
		}
	}

	return args
}

// Основные функции
func Perform(args Arguments, writer io.Writer) error {
	err0 := validateConsequently(args, *requiredFlag, additionalRequiredFlags)

	if err0 != nil {
		return err0
	}

	var (
		buf []byte
		err error
	)

	switch args[pOperation] {
	case opList:
		buf, err = list(args)
	case opAdd:
		buf, err = add(args)
	case opFindById:
		buf, err = findById(args)
	case opRemove:
		buf, err = remove(args)
	}

	if err != nil {
		return err
	}

	writer.Write(buf)

	return nil
}

func add(args Arguments) ([]byte, error) {
	f, err := readFileToJSON(args[pFileName])

	if err != nil {
		return nil, err
	}

	u := User{}
	item := args[pItem]
	err3 := json.Unmarshal([]byte(item), &u)

	if err3 != nil {
		return nil, err3
	}

	for _, rec := range f {
		if rec.Id == u.Id {
			return []byte("Item with id " + rec.Id + " already exists"), nil
		}
	}

	f = append(f, u)
	marshalled, errLast := json.Marshal(f)

	if errLast == nil {
		if err := ioutil.WriteFile(args[pFileName], marshalled, 0666); err != nil {
			return nil, err
		}
	}

	return marshalled, errLast
}

func list(args Arguments) ([]byte, error) {
	return readFileToBuffer(args[pFileName])
}

func findById(args Arguments) ([]byte, error) {
	f, err := readFileToJSON(args[pFileName])

	if err != nil {
		return nil, err
	}

	for _, rec := range f {
		if rec.Id == args[pId] {
			marshalled, err := json.Marshal(rec)
			return marshalled, err
		}
	}

	return nil, nil
}

func remove(args Arguments) ([]byte, error) {
	f, err := readFileToJSON(args[pFileName])

	if err != nil {
		return nil, err
	}

	var foundId string

	for i, rec := range f {
		if rec.Id == args[pId] {
			foundId = rec.Id
			f = append(f[:i], f[i+1:]...)
		}
	}

	if foundId == "" {
		return []byte("Item with id " + args[pId] + " not found"), nil
	}

	marshalled, errLast := json.Marshal(f)

	if errLast == nil {
		if err := ioutil.WriteFile(args[pFileName], marshalled, 0666); err != nil {
			return nil, err
		}
	}

	return marshalled, errLast
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
