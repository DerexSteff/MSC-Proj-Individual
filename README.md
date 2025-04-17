# Individual Project MSC description
This project is dedicated to describing how 2 different applications can communicate using both oneM2M and FIWARE.
- For the oneM2M implementation it was chosen ACME
- For the FIWARER implementation it was chosen Orion GE

## Project architecture
This project consists of 2 very simple apps developed in Go
- **App A (temperature_sensor)**: This app "reads" a value and stores it in the communication medium
- **App B (display_temperature)**: This app is notified of the new stored value and uses it do some function

- For the example scenario it was defined that app A will read the temperature of a sensor in a controlled chamber (every 3 seconds)
- App B will read the stored temperature and display to the user if the temperature is OK, too low, or too high
- This scenario is purely examplatory and the values read are randomly generated
