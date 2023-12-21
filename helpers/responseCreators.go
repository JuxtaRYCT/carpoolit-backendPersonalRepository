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
