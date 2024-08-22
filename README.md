# Groupie Tracker

## Overview

Groupie Tracker is a web application designed to display information about artists, their concert dates, and locations using data from a RESTful API. The application is built with a backend written in Go and a frontend using HTML and CSS, offering a user-friendly interface for exploring and visualizing artist information, concert schedules, and locations.

## Objectives

- **Data Manipulation**: Retrieve and process data from a provided API that includes information about artists, concert dates, locations, and their relations.
- **Visualization**: Present the data through various visualizations such as cards, tables, and graphics to enhance user experience.
- **Client-Server Interaction**: Implement client-server communication to fetch data on user actions, ensuring a responsive and interactive web application.

## Components

The application is divided into four main parts based on the API provided:

1. **Artists**: Contains information about bands and artists, including names, images, the start year of their activity, their first album release date, and members.
2. **Locations**: Lists the last and upcoming concert locations, with locations normalized to ensure readability.
3. **Dates**: Provides details on the last and upcoming concert dates, displayed in a visually appealing format.
4. **Relation**: Connects artists, dates, and locations, forming the link between the other components, allowing for a comprehensive view of artist activities.

## Features

- **Artist Information**: View details about artists, including images, biographical information, and members. The index page lists all artists with limited information, while detailed views are available for each artist.
- **Concert Dates and Locations**: Display upcoming and past concert dates for each artist, along with the venues where the concerts are held. Dates are presented in a visually distinctive format.
- **Interactive Elements**: Engage with various UI components to fetch and view data dynamically, such as clicking on an artist card to view more details.
- **Styling**: Consistent and professional styling throughout the application, including a polished look for the detail pages and normalization of location names (e.g., converting "florida-usa" to "Florida USA").

## Setup and Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/hamzamaach/groupie-tracker.git
   ```

2. **Navigate to the Project Directory**
   ```bash
   cd groupie-tracker
   ```

3. **Configure the Application**
   - Update the `config` package to set the API URL and port number as required.

4. **Run the Backend**
   - Ensure you have Go installed.
   - Run the server:
     ```bash
     go run main.go
     ```

5. **Open the Application**
   - Open a browser and navigate to `http://localhost:<port_number>` to view the application.

## Packages and Structure

- **`handlers`**: Manages HTTP requests and routes, including fetching data from the API and rendering templates.
- **`utils`**: Contains utility functions for tasks such as normalizing location names and handling common logic.
- **`models`**: Defines the data structures used in the application, including separate structs for index and detail views.
- **`config`**: Holds configuration data, including API URLs and server port numbers.

## Contributing

Feel free to contribute to the project by opening issues or submitting pull requests. Ensure that any new code adheres to the project's coding standards and includes relevant tests.

## Contact

For any questions or support, please contact [hamzamaach56@gmail.com](mailto:hamzamaach56@gmail.com).

---

Thank you for using Groupie Tracker! We hope you enjoy exploring artist information and concert details.