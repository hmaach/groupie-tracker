# Groupie Tracker

## Overview

Groupie Tracker is a web application designed to display information about artists, their concert dates, and locations using data from a RESTful API. The application leverages a backend written in Go and a frontend built with HTML to offer a user-friendly interface for exploring and visualizing artist information, concert schedules, and locations.

## Objectives

- **Data Manipulation**: Retrieve and process data from a provided API that includes information about artists, concert dates, locations, and their relations.
- **Visualization**: Present the data through various visualizations such as cards, tables, and graphics to enhance user experience.
- **Client-Server Interaction**: Implement client-server communication to fetch data on user actions, ensuring a responsive and interactive web application.

## Components

The application is divided into four main parts based on the API provided:

1. **Artists**: Contains information about bands and artists, including names, images, the start year of their activity, their first album release date, and members.
2. **Locations**: Lists the last and upcoming concert locations.
3. **Dates**: Provides details on the last and upcoming concert dates.
4. **Relation**: Connects artists, dates, and locations, forming the link between the other components.

## Instructions

### Backend

- **Language**: Go
- **Packages**: Only standard Go packages are used.
- **Functionality**: The backend handles API requests, processes data, and serves responses to the frontend.
- **Error Handling**: Ensure the server does not crash, and all pages work correctly. Implement error handling and validation.
- **Testing**: Unit tests are recommended to ensure code reliability.

### Frontend

- **Language**: HTML
- **Visualization**: Display data using a variety of UI components such as blocks, cards, tables, lists, and graphics.
- **Event Handling**: Implement actions triggered by user interactions that communicate with the server to fetch and display data dynamically.

### API Interaction

- **RESTful API**: The frontend communicates with the backend through RESTful API endpoints.
- **Request-Response**: Implement features that involve client-server requests and handle responses to update the user interface accordingly.

## Setup and Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/hamzamaach/groupie-tracker.git
   ```

2. **Navigate to the Project Directory**
   ```bash
   cd groupie-tracker
   ```

3. **Run the Backend**
   - Ensure you have Go installed.
   - Navigate to the backend directory and run the server:
     ```bash
     cd backend
     go run main.go
     ```

4. **Open the Frontend**
   - Open the HTML files in your browser to view the application.

## Features

- **Artist Information**: View details about artists including images and biographical information.
- **Concert Dates**: Display upcoming and past concert dates for each artist.
- **Location Details**: Show the venues for concerts and their details.
- **Interactive Elements**: Engage with various UI components to fetch and view data dynamically.

## Contributing

Feel free to contribute to the project by opening issues or submitting pull requests. Ensure that any new code adheres to the project's coding standards and includes relevant tests.

## Contact

For any questions or support, please contact [hamzamaach56@gmail.com](mailto:hamzamaach56@gmail.com).

---

Thank you for using Groupie Tracker! We hope you enjoy exploring artist information and concert details.