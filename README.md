# Weather App Level 1
A weather app using a third party api with the goal being a build upon project with additional features as time goes on. The goal of this project is to simply be able to connect to an api to pull weather from and display in a browser using very few technologies. Future versions of this project will include more advanced features, but for now this is just meant to be a smaller more achievable goal.

## Table of Contents
1. [Project Breakdown](#project-breakdown)
2. [Technologies Used](#technologies-used)
3. [Project Goals](#project-goals)
4. [Findings](#findings)
5. [Conclusion](#conclusion)

## Project Breakdown
The approach to building this app is to pull data from a third party weather api service whether it be [Free Weather API](https://www.weatherapi.com), [the National Weather Service](https://www.weather.gov), or [Open Weather Map](https://openweathermap.org/api) to get the forecasts to be displayed to the user with a web browser. The goal is to focus more on using backend programming with a minimal ui. This version mostly focus on using Golang to handle most of the load along with an sql lite database to store local favorites. The user would have the option to remove those favorites. Html will display the page and a minimal css framework will be used to add some design with pico css. Go will be used for some dynamic loading and possibly some javascript. That is the breakdown of technologies used and project breakdown.

## Technologies Used
1. [Go](https://go.dev) for backend
2. HTML
3. [Pico Css](https://picocss.com)
4. [SQLlite](https://www.sqlite.org)

## Project Goals
- [x] Get the Api key and practice a call to the service
- [x] Setup the HTML pages
- [x] Setup the design with pico css
- [x] Pull forecast data from api
- [x] Format the api request result into proper format for the user with icons and a weather report

## Findings

## Conclusion