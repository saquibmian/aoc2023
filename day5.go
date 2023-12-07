package main

import (
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var day5part1 = RunFunc(func(input []string) string {
	identity := func(i int) int { return i }
	wrap := func(inner func(int) int, wrapper func(i int, f func(int) int) int) func(int) int {
		return func(i int) int {
			return wrapper(i, inner)
		}
	}
	type almanac struct {
		seeds                 []int
		seedToSoil            func(int) int
		soilToFertilizer      func(int) int
		fertilizerToWater     func(int) int
		waterToLight          func(int) int
		lightToTemperature    func(int) int
		temperatureToHumidity func(int) int
		humidityToLocation    func(int) int
	}
	a := almanac{
		seeds: Map(
			strings.Fields(
				strings.TrimPrefix(input[0], "seeds: "),
			),
			func(s string) int { return MustInt(s) },
		),
		seedToSoil:            identity,
		soilToFertilizer:      identity,
		fertilizerToWater:     identity,
		waterToLight:          identity,
		lightToTemperature:    identity,
		temperatureToHumidity: identity,
		humidityToLocation:    identity,
	}
	input = input[1:]
	for len(input) > 0 {
		line := input[0]
		input = input[1:]
		if Trim(line) == "" {
			continue
		}
		if Trim(line) == "seed-to-soil map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.seedToSoil = wrap(a.seedToSoil, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "soil-to-fertilizer map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.soilToFertilizer = wrap(a.soilToFertilizer, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "fertilizer-to-water map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.fertilizerToWater = wrap(a.fertilizerToWater, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "water-to-light map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.waterToLight = wrap(a.waterToLight, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "light-to-temperature map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.lightToTemperature = wrap(a.lightToTemperature, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "temperature-to-humidity map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.temperatureToHumidity = wrap(a.temperatureToHumidity, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "humidity-to-location map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.humidityToLocation = wrap(a.humidityToLocation, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		}
	}
	lowestLocationNumber := math.MaxInt
	for _, seed := range a.seeds {
		lowestLocationNumber = Min(
			lowestLocationNumber,
			a.humidityToLocation(
				a.temperatureToHumidity(
					a.lightToTemperature(
						a.waterToLight(
							a.fertilizerToWater(
								a.soilToFertilizer(
									a.seedToSoil(
										seed,
									),
								),
							),
						),
					),
				),
			),
		)
	}
	return strconv.Itoa(lowestLocationNumber)
})

var day5part2 = RunFunc(func(input []string) string {
	identity := func(i int) int { return i }
	wrap := func(inner func(int) int, wrapper func(i int, f func(int) int) int) func(int) int {
		return func(i int) int {
			return wrapper(i, inner)
		}
	}
	type almanac struct {
		seeds                 []int
		seedToSoil            func(int) int
		soilToFertilizer      func(int) int
		fertilizerToWater     func(int) int
		waterToLight          func(int) int
		lightToTemperature    func(int) int
		temperatureToHumidity func(int) int
		humidityToLocation    func(int) int
	}
	a := almanac{
		seedToSoil:            identity,
		soilToFertilizer:      identity,
		fertilizerToWater:     identity,
		waterToLight:          identity,
		lightToTemperature:    identity,
		temperatureToHumidity: identity,
		humidityToLocation:    identity,
	}
	seedParts := Map(
		strings.Fields(
			strings.TrimPrefix(input[0], "seeds: "),
		),
		func(s string) int { return MustInt(s) },
	)
	input = input[1:]
	for len(input) > 0 {
		line := input[0]
		input = input[1:]
		if Trim(line) == "" {
			continue
		}
		if Trim(line) == "seed-to-soil map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.seedToSoil = wrap(a.seedToSoil, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "soil-to-fertilizer map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.soilToFertilizer = wrap(a.soilToFertilizer, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "fertilizer-to-water map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.fertilizerToWater = wrap(a.fertilizerToWater, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "water-to-light map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.waterToLight = wrap(a.waterToLight, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "light-to-temperature map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.lightToTemperature = wrap(a.lightToTemperature, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "temperature-to-humidity map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.temperatureToHumidity = wrap(a.temperatureToHumidity, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		} else if Trim(line) == "humidity-to-location map:" {
			for len(input) > 0 && Trim(input[0]) != "" {
				line := input[0]
				input = input[1:]
				rangee := strings.Fields(line)
				destStart, srcStart, size := MustInt(rangee[0]), MustInt(rangee[1]), MustInt(rangee[2])
				a.humidityToLocation = wrap(a.humidityToLocation, func(i int, f func(int) int) int {
					if i >= srcStart && i < srcStart+size {
						return destStart + (i - srcStart)
					}
					return f(i)
				})
			}
		}
	}
	requests := make(chan int)
	results := make(chan int, runtime.GOMAXPROCS(0))
	var wg sync.WaitGroup
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		wg.Add(1)
		go func() {
			lowestLocationNumber := math.MaxInt
			for seed := range requests {
				lowestLocationNumber = Min(
					lowestLocationNumber,
					a.humidityToLocation(
						a.temperatureToHumidity(
							a.lightToTemperature(
								a.waterToLight(
									a.fertilizerToWater(
										a.soilToFertilizer(
											a.seedToSoil(
												seed,
											),
										),
									),
								),
							),
						),
					),
				)
			}
			results <- lowestLocationNumber
			wg.Done()
		}()
	}
	var (
		total int
		done  int
	)
	for i := 0; i < len(seedParts); i = i + 2 {
		total += seedParts[i+1]
	}
	for i := 0; i < len(seedParts); i = i + 2 {
		start, size := seedParts[i], seedParts[i+1]
		for j := 0; j < size; j++ {
			seed := start + j
			done++
			requests <- seed
		}
	}
	close(requests)
	wg.Wait()
	close(results)
	lowestLocationNumber := math.MaxInt
	for result := range results {
		lowestLocationNumber = Min(lowestLocationNumber, result)
	}
	return strconv.Itoa(lowestLocationNumber)
})
