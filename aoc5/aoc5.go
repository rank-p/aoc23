package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
) 

type mapValue struct {
	src int
	mapRange int
}


func getMappedValue(dstValue int, valueMap map[int]mapValue) int {
	for mapDst, value := range valueMap {
		if dstValue < mapDst {
			continue
		} else if mapDst + value.mapRange < dstValue {
			continue
		} else {
			return (dstValue - mapDst) + value.src
		}
	}
	return dstValue

}

func min(valueMap map[int]mapValue, k int) int {
	keys := make([]int, len(valueMap))

	i := 0
	for k:= range valueMap {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys[k]
	

}

func isValidSeed(seed int, seeds []int) bool {
	for i:= 0; i < len(seeds); i++ {
		if i % 2 == 1 {
			if seed >= seeds[i-1] && seed <= seeds[i-1]+seeds[i]-1 {
				return true
			}

		}
	}
	return false 
	
}

func main() {
	file, _ := os.Open("aoc5.txt")
	defer file.Close()
	
	seedToSoil := make(map[int]mapValue)
	soilToFertilizer := make(map[int]mapValue)
	fertilizerToWater := make(map[int]mapValue)
	waterToLight := make(map[int]mapValue)
	lightToTemperature := make(map[int]mapValue)
	temperatureToHumidity := make(map[int]mapValue)
	humidityToLocation := make(map[int]mapValue)
	var seeds []int
	var currentMap map[int]mapValue

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch scanner.Text() {
		case "":
			continue
		case "seed-to-soil map:":
			currentMap = seedToSoil
		case "soil-to-fertilizer map:":
			currentMap = soilToFertilizer
		case "fertilizer-to-water map:": 
			currentMap = fertilizerToWater
		case "water-to-light map:":
			currentMap = waterToLight
		case "light-to-temperature map:":
			currentMap = lightToTemperature
		case "temperature-to-humidity map:":
			currentMap = temperatureToHumidity
		case "humidity-to-location map:":
			currentMap = humidityToLocation
		default:
			if strings.Contains(scanner.Text(), "seeds") {
				inputs := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")
				for i := range inputs {
					input, _ := strconv.Atoi(inputs[i])
					seeds = append(seeds, input)
				}
				continue
			}
			inputs := strings.Split(scanner.Text(), " ")
			dst, _ := strconv.Atoi(inputs[0])
			src, _ := strconv.Atoi(inputs[1])
			mapRange, _ := strconv.Atoi(inputs[2])
			currentMap[dst] = mapValue{src, mapRange}
		}
	
	}
	
	var soil int
	var fertilizer int
	var water int
	var light int
	var temperature int
	var humidity int
	var seed int
	location := 1
	for true {
		humidity = getMappedValue(location, humidityToLocation)
		temperature = getMappedValue(humidity, temperatureToHumidity)
		light = getMappedValue(temperature, lightToTemperature)
		water = getMappedValue(light, waterToLight)
		fertilizer = getMappedValue(water, fertilizerToWater)
		soil = getMappedValue(fertilizer, soilToFertilizer)
		seed = getMappedValue(soil, seedToSoil)
		if location % 1000000 == 0 {
			fmt.Println(location)
		}
		if isValidSeed(seed, seeds) {
			fmt.Println(location)
			break
			}
		
		location++
	}

	
}
