package accumulate

const testVersion = 1

func Accumulate(collection []string, converter func(string) string) []string {
  var newCollection = make([]string, 0)

  for _, element := range collection {
    newCollection = append(newCollection, converter(element))
  }

  return newCollection
}
