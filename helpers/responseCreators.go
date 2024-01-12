package helpers

import "carpool-backend/models"

func CreateResponseRide(ride models.Ride) models.RideResponse {
	return models.RideResponse{
		HostUserID:    ride.HostUserID,
		StartLocation: ride.StartLocation,
		EndLocation:   ride.EndLocation,
		StartTime:     ride.StartTime,
		TotalSeats:    ride.TotalSeats,
		BookedSeats:   ride.BookedSeats,
		TotalPrice:    ride.TotalPrice,
	}
}

func CreateResponseBooking(booking models.Booking) models.BookingResponse {
	return models.BookingResponse{
		RideID:      booking.RideID,
		PassengerID: booking.PassengerID,
	}
}

func CreateResponseUser(user models.User) models.UserResponse {
	return models.UserResponse{
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
		ContactNumber:     user.ContactNumber,
		Gender:            user.Gender,
		YOB:               user.YOB,
	}
}
