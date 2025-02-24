# Music Library API

## Overview
The **Music Library API** is a RESTful service for managing songs, artists, and lyrics. It provides endpoints to create, retrieve, update, and delete songs, along with filtering and pagination capabilities. The application is built in **Go (Golang)** and uses **PostgreSQL** for data storage and **Redis** for caching.

## Technologies Used
- **Golang** (v1.24)
- **PostgreSQL** (v16)
- **Redis** (v9)
- **Docker & Docker Compose**
- **Gin Framework**

## Getting Started

### Prerequisites
Ensure you have the following installed:
- **Docker** & **Docker Compose**
- **Go (1.24)**

### Installation & Running

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/music-lib.git
   cd music-lib
   ```

2. Create a **.env** file in the root directory:
   ```sh
   PORT=8080
   DB_PORT=5432
   DB_NAME=music_db
   DB_USER=root_user
   DB_PASSWORD=Dost0n1k
   DB_HOST=postgres_db
   DB_SSLMODE=disable
   REDIS_HOST=redis_cache
   REDIS_PORT=6379
   REDIS_TTL=3600
   ```

3. Start the services:
   ```sh
   docker-compose up --build
   ```

4. The API will be available at:
   ```
   http://localhost:8080
   ```

---
## API Endpoints

### **1. Create a Song**
- **Endpoint:** `POST /songs`
- **Description:** Adds a new song to the database.
- **Request Body (JSON):**
  ```json
  {
    "name": "Song Title",
    "group": "Group Name",
    "artists": ["Artist 1", "Artist 2"],
    "lyrics": "Song lyrics...",
    "release_date": "2025-02-25T00:00:00Z"
  }
  ```
- **Response:**
  ```json
  {
    "id": "uuid",
    "name": "Song Title",
    "group": "Group Name",
    "artists": ["Artist 1", "Artist 2"],
    "lyrics": "Song lyrics...",
    "release_date": "2025-02-25T00:00:00Z"
  }
  ```

### **2. Get All Songs**
- **Endpoint:** `GET /songs`
- **Description:** Retrieves a list of all songs.
- **Response:**
  ```json
  [
    {
      "id": "uuid",
      "name": "Song Title",
      "group": "Group Name",
      "artists": ["Artist 1", "Artist 2"],
      "release_date": "2025-02-25T00:00:00Z"
    }
  ]
  ```

### **3. Get Songs with Filters**
- **Endpoint:** `GET /songs/filtered`
- **Description:** Retrieves songs based on filters (e.g., artist, group, release date range).
- **Query Parameters:**
  - `artist` - Filter by artist name
  - `group` - Filter by group name
  - `start_date` & `end_date` - Filter by release date range

### **4. Get a Song by ID**
- **Endpoint:** `GET /songs/:id`
- **Description:** Retrieves a song by its unique ID.

### **5. Get Lyrics (Paginated)**
- **Endpoint:** `GET /songs/:id/lyrics?page={page}&limit={limit}`
- **Description:** Retrieves paginated lyrics of a song.

### **6. Get Songs by Artist**
- **Endpoint:** `GET /songs/artists?name={artist_name}`
- **Description:** Retrieves all songs by a specific artist.

### **7. Update a Song**
- **Endpoint:** `PUT /songs/:id`
- **Description:** Updates an existing song.
- **Request Body:** (Same as Create)

### **8. Delete a Song**
- **Endpoint:** `DELETE /songs/:id`
- **Description:** Soft deletes a song from the database.

---
## Contributing
Feel free to fork this repository and submit pull requests.
