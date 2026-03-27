package ports

import "github.com/vinicius-benevides/go-rest-api/src/domain/models"

// EventRepository defines persistence operations for events.
type EventRepository interface {
    GetAllEvents() ([]models.Event, error)
    GetEventByID(id int64) (*models.Event, error)
    CreateEvent(event *models.Event) error
    UpdateEvent(event *models.Event) error
    DeleteEvent(id int64) error
}

// RegistrationRepository exposes data access for event registrations.
type RegistrationRepository interface {
    GetRegistrationsByEvent(eventID int64) ([]int64, error)
    CreateRegistration(eventID, userID int64) error
    CancelRegistration(eventID, userID int64) error
}

// UserRepository provides persistence methods for users.
type UserRepository interface {
    GetUserByEmail(email string) (*models.User, error)
    CreateUser(user *models.User) error
}
