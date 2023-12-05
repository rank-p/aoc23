package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
) 

type mapValue struct {
	dst int
	mapRange int
}


func getMappedValue(srcValue int, valueMap map[int]mapValue) int {
	for mapSrc, value := range valueMap {
		if mapSrc > srcValue {
			continue
		} else if mapSrc + value.mapRange < srcValue {
			continue
		} else {
			return value.dst + (srcValue - mapSrc)
		}
	}
	return srcValue

}

func min(slice []int) int {
	var m int
	for i,e := range slice {
		if i==0 || e < m {
			m = e
		}
	}
	return m
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
			currentMap[src] = mapValue{dst, mapRange}
		}
	
	}
	
	var locations []int
	for _, seed := range seeds {
		soil := getMappedValue(seed, seedToSoil)
		fertilizer := getMappedValue(soil, soilToFertilizer)
		water := getMappedValue(fertilizer, fertilizerToWater)
		light := getMappedValue(water, waterToLight)
		temperature := getMappedValue(light, lightToTemperature)
		humidity := getMappedValue(temperature, temperatureToHumidity)
		location := getMappedValue(humidity, humidityToLocation)
		locations = append(locations, location)
	}
	fmt.Println(min(locations))
}
