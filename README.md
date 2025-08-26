# Workout Dev

## Workout Tracker App

This is an application which uses a REST API to manage workouts.

## Features

1. Authentication and Authorization: Sign-Up, Login and JWT for authentication

2. Workout management:
    * Create Workout with multiple exercises
    * Add comments
    * Delete Workout
    * Schedule Workouts for specific date and time

## API Usage

All authenticated routes require a `Authorization: Bearer <YOUR_JWT_TOKEN>` header.

### Authentication

* **Register a new user**
    * **Endpoint:** **`POST /register`**
    * **Body:**

        ```json
        {
            "first_name": "Test",
            "last_name": "User",
            "email": "test@example.com",
            "password": "password123"
        }
        ```

* **Log in**
    * **Endpoint:** **`POST /login`**
    * **Body:**

        ```json
        {
            "email": "test@example.com",
            "password": "password123"
        }
        ```

    *   **Response:**
        ```json
        {
            "message": "Welcome, Test!",
            "token": "YOUR_JWT_TOKEN"
        }
        ```

### Workouts

*   **Get all workouts for the logged-in user**
    *   **Endpoint:** `GET /workouts`

*   **Get a specific workout by ID**
    *   **Endpoint:** `GET /workouts/:id`

*   **Create a new workout**
    *   **Endpoint:** `POST /workouts`
    *   **Body:**
        ```json
        {
            "name": "Morning Workout",
            "due_date": "2025-08-27T10:00:00Z"
        }
        ```

*   **Add exercises to a workout**
    *   **Endpoint:** `POST /workouts/:id/exercises`
    *   **Body:**
        ```json
        {
            "exercise_names": ["Squat", "Deadlift", "Overhead Press"]
        }
        ```

*   **Add sets to a workout**
    *   **Endpoint:** `POST /workouts/sets`
    *   **Body:**
        ```json
        {
            "workout_id": 1,
            "exercise_name": "Squat",
            "reps": 10,
            "weight": 100
        }
        ```

*   **Add comments to a workout**
    *   **Endpoint:** `POST /workouts/:id/comments`
    *   **Body:**
        ```json
        {
            "comments": ["Felt strong today!", "Need to work on my form."]
        }
        ```

*   **Delete a workout**
    *   **Endpoint:** `DELETE /workouts/:id`
