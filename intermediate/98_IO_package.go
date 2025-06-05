*******LOTS OF LORE*******


package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}

	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // stack
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // heap
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}
	fmt.Println(buf.String())
}

// We're creating a multireader so we need a pointer, right? To read from two resources. So we need to create a pointer that is going to store the read values, the values that it reads into a memory address. And then, we're going to access that memory address and extract the value that is stored in that buffer. So this time, in case 2, we're creating a pointer instead of a value.   
// There's two declarations above: 
// 1) var buf bytes.Buffer  
// 2) buf := new(bytes.Buffer)

// Both do the same thing, create a byte buffer but the internal workings are different. 
// So, "buf" in 2 is a pointer to bytes.Buffer. "Buf" here is not of type bytebuffer but buf is of type pointer to bytes.Buffer.

// This is a common way to create a new instance of a type and get a pointer to it. Just use the 
// new() function; built-in func that allocates memory. 

// In 1) it is a declaration with var and it creates a variable buff of type byte.Buffer. This allocates a Bytebuffer instance directly and not a pointer, and initialises it to its zero value. More trad way to declare a var and directly create an instance of a type. 

// SO 1 is of value type and 2 is of pointer type. 

// Another diff is that the new(bytes.Buffer) allocates memory on the heap, which can be useful if you need to pass around the buffer or manage its lifetime explicitly.
// However, var buf bytes.Buffer creates memory on the stack, and unless its part of a struct or otherwise managed by the Go runtime, the mem alloc will be on the stack. If unless(struct/managed by runtime) case is true, then it *could* be on the heap.
 
func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	//Goroutines are essentially any func, any immediate func. Immediate func are the func that are executed immediately once they are defined; go func(){}() //immd. execution
	//If we add the go keyword before a func, it becomes a goroutine. 
	//Goroutines are basically anonymous functions that are executed immediately as soon as they're defined. 

	//Async await promises in Javascript -> takes that func out of the main thread, takes it to the background, and once it finishes off, then it joins the main thread. Same for goroutines. 
	
	//V8 engine. 

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creating file:", err)
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error opening/creating file:", err)
	}

	// // Type(value)
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln("Error opening/creating file:", err)
	// }
}

func main() {

	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello Reader!"))

	fmt.Println("=== Write to Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Writer!")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi Reader Example ===")
	multiReaderExample()

	fmt.Println("=== Pipe Example ===")
	pipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "Hello File!")

	resource := &MyResource{name: "TestResource"}
	closeResource(resource)

}

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Println("Closing resource:", m.name)
	return nil
}