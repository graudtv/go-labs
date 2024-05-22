package main

import (
  "fmt"
  "flag"
  "log"
  "io"
  "io/ioutil"
  "bytes"
  "strings"
  "strconv"
  "encoding/xml"
  "encoding/json"
  "sort"
  "golang.org/x/text/encoding/charmap"
  "gopkg.in/yaml.v2"
)

type Config struct {
  InputFilename string `yaml:"input"`
  OutputFilename string `yaml:"output"`
}

type FPValue float64

/* Custom parser for currency values: for some reason comma is used as
 * separator instead of point */
func (v *FPValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
  var str string
  err := d.DecodeElement(&str, &start)
  if err != nil {
    return err
  }
  str = strings.Replace(str, ",", ".", 1)
  fp, err := strconv.ParseFloat(str, 64)
  if err != nil {
    return err
  }
  *v = FPValue(fp)
  return nil
}

type XML struct {
  XMLName xml.Name `xml:"ValCurs"`
  Quotes [] struct {
    NumCode int 
    CharCode string
    Value FPValue
  } `xml:"Valute"`
}

type Quote struct {
  NumCode int `json:"num_code"`
  CharCode string `json:"char_code"`
  Value float64 `json:"value"`
}

type QuotesByValue []Quote
func (quotes QuotesByValue) Len() int { return len(quotes) }
func (quotes QuotesByValue) Swap(i, j int) { quotes[i], quotes[j] = quotes[j], quotes[i] }
func (quotes QuotesByValue) Less(i, j int) bool { return quotes[i].Value < quotes[j].Value }

func readFile(path string) []byte {
  bytes, err := ioutil.ReadFile(path)
  if err != nil {
    log.Fatal(err)
  }
  return bytes
}

func parseConfigFile(cfgPath string) Config {
  var cfg Config
  err := yaml.Unmarshal(readFile(cfgPath), &cfg)
  if err != nil {
    log.Fatal(err)
  }
  return cfg
}

func parseQuotesXML(xmlFile []byte) XML {
  /* Windows 1251 xml parsing: https://ru.stackoverflow.com/a/713848 */
  d := xml.NewDecoder(bytes.NewReader(xmlFile))
  d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
      switch charset {
      case "windows-1251":
          return charmap.Windows1251.NewDecoder().Reader(input), nil
      default:
          return nil, fmt.Errorf("unknown charset: %s", charset)
      }
  }
  var x XML
  err := d.Decode(&x)
  if err != nil {
    log.Fatal(err)
  }
  return x
}

func extractQuotes(x XML) [] Quote {
  var quotes []Quote
  for _, q := range x.Quotes {
    quotes = append(quotes, Quote{NumCode: q.NumCode, CharCode: q.CharCode, Value: float64(q.Value)})
  }
  return quotes
}

func main() {
  log.SetFlags(0)
  cfgPath := flag.String("config", "", "path to config.yaml")
  verbose := flag.Bool("verbose", false, "enable verbose mode")
  flag.Parse()
  if len(*cfgPath) == 0 {
    log.Fatal("error: -config option required")
  }
  cfg := parseConfigFile(*cfgPath)
  if *verbose { fmt.Println("config:", cfg) }
  quotes := extractQuotes(parseQuotesXML(readFile(cfg.InputFilename)))
  if *verbose { fmt.Println("input:", quotes) }
  sort.Sort(QuotesByValue(quotes))
  if *verbose { fmt.Println("sorted by value:", quotes) }

  output, _ := json.MarshalIndent(quotes, "", "  ")
  if *verbose { fmt.Println("output:", string(output)) }
  err := ioutil.WriteFile(cfg.OutputFilename, output, 0644)
  if err != nil {
    log.Fatal(err)
  }
}
