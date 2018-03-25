package main

import (
	"fmt"
  "sync"
	"time"
  "math/rand"
)

// Genera un unsigned integer de 64 bits entre min y max
func randonuint(min, max uint64) uint64 {
	return uint64(rand.Int63n(int64(max-min+1)))
}

// Suma todos los elementos de un slice de unsigned integers
func sliceSum(a []uint64) uint64 {
	var s uint64 = 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

// Genera los primeros n kbonaccis dado un k
func first_n_kbonaccis(k uint64) []uint64 {
  a := []uint64{}
  var i uint64
  for i = 0; i < k; i++ {
    a = append(a, i)
  }
  return a
}

// Funcion auxiliar para calcular el kbonacci con una recursion de cola
func kbonacciaux(k uint64, j uint64, a[]uint64, i uint64) uint64  {
  if i == k {
    i = 0
  }
  a[i] = sliceSum(a)
  if j == k {
    return a[i]
  }
  return kbonacciaux(k, j-1, a, i+1)
}

// Calcula el kbonacci de un numero j
func kbonacci(k uint64, j uint64) uint64  {
  if j < k {
    return j
  }
  return kbonacciaux(k, j, first_n_kbonaccis(k), 0)
}

// Funcion para traducir un estructura tuple a kbonacci de su j y guardar el resultado en un apuntador r
func kbonacci_tuple(t tuple, r *uint64)  {
  *r = kbonacci(t.k, t.j)
}

// Estructura para almacenar el k y j para un kbonacci
type tuple struct {
  k uint64;
  j uint64;
}

// Funcion principal
func main() {

	rand.Seed(time.Now().UnixNano())

  var waiter sync.WaitGroup

  ts := []tuple{}
  results := []uint64{}

  for i := 0; i < 10; i++ {
    ts = append(ts, tuple{randonuint(100,10000), randonuint(1000,1000000)})
    results = append(results, 0)
    waiter.Add(1)
  }

	fmt.Println("Tuplas Generadas, Calculando kbonaccis")

	startTime := time.Now()

  for i := 0; i < len(ts); i++ {
    go func(i int) {
      var r uint64
      kbonacci_tuple(ts[i], &r)
      results[i] = r
      defer waiter.Done()
    }(i)
  }

  waiter.Wait()

	finishTime := time.Now()

  for i := 0; i < len(ts); i++ {
  	fmt.Printf("%d-bonacci de %d = %d\n", ts[i].k, ts[i].j, results[i])
  }

	fmt.Println("Calculado en ", finishTime.Sub(startTime))

}
