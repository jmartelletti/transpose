// input steps +/- to transpose
// read file
// parse file
//  dumb
//   - search for any numbers
//   - add/multiple steps from numbers
//  smart
//   what if note is lower than string allows?
//   ignore numbers that aren't positions - e.g. multiples (4x) numbers in intro text/lyrics/etc, 
// output updated file
package main

import (
  "os"
  "flag"
  "bufio"
  "log"
  "strconv"
  "fmt"
  "regexp"
)

func readFile(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func writeFile(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}

func scanFile(lines []string, steps *int) error {
  r, _ := regexp.Compile("[\\Whp](\\d+)")

  for _, line := range lines {
    derps := []byte(line)

    out := r.ReplaceAllFunc(derps, func(s []byte) []byte {
      position, _ := strconv.ParseInt(string(s), 0, 64)
      new_position := position + int64(*steps)
      return []byte(string(strconv.FormatInt(new_position, 10)))
    })

    fmt.Println(string(out))
  }
  return nil
}

func main() {
  steps := flag.Int("steps", 0, "Number of steps to transpose (default is down)")

  flag.Parse()

  fmt.Printf("%d steps\n", *steps)

  lines, err := readFile("tabs/Flatliners: Do or Die.txt")
  if err != nil {
    log.Fatalf("readFile: %s", err)
  }

  scanFile(lines, steps)
}
