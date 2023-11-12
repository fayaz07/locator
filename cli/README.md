# locator command line interface

## How it works

- Show info about the program
- Request the user for dataset source
  - If the location of the file is invalid, the program will ask the user to enter a valid location
  - If the location is valid, application validates the source file and returns error if invalid
- Ask user if the user wants to use all countries or exclude specific countries or use specific countries
- Strategy selection
  - Search location in all countries
  - Search location in specific country/countries
  - Both of the above
- Request the user for target location
  - When target location is invalid, the program will ask the user to enter a valid location
