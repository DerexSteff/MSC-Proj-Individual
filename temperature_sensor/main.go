package main

import (
    "fmt"
    "math/rand/v2"
    "time"
)

func main() {
    for (true) { // read temperature and send it infinitely
        // generate a value for temperature
        temperature_reading := 18 + (rand.NormFloat64() * 0.1)
        var is_valid_temperature string
        if (temperature_reading > 18.1){
            is_valid_temperature = "Temperature too high!"
        } else if (temperature_reading < 17.9) {
            is_valid_temperature = "Temperature too low!"
        } else {
            is_valid_temperature = "Temperature is OK"
        }
        fmt.Printf("Temperature: %.3f°C — %s\n", temperature_reading, is_valid_temperature)
        time.Sleep(3 * time.Second)
    }
}
