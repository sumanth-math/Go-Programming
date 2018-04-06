using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;
using Barclays.Theater.BusinessLogic.Interface;
using Barclays.Theater.DataAccessLogic.Interface;
using Barclays.Theater.DataAccessLogic;
using Barclays.Theater.Entities;
using Barclays.Theater.Exceptions;

namespace Barclays.Theater.BusinessLogic
{
    public class SeatingReservationBc : ISeatingReservationBc
    {
        IRepository<SectionSeating> _theaterSeatingRepository;
        IRepository<ReservationInformation> _reservationRepository;
        private Regex VALID_SEATING_ROW_PATTERN = new Regex(@"^[0-9\s+]*$");
        private Regex VALID_TICKETING_REQUEST_PATTERN = new Regex(@"^[a-zA-Z0-9_]*\s+[0-9]*$");
        private const string INVALID_THEATER_SEATING_ERROR_MESSAGE =
            "Invalid theater seaating layout. Please provide a valid theater seating! Ex:- A row with values 3 6 3 means 3 sections with 3, 6 and 3 seats in each section respectively.";
        private const string INVALID_RESERVATION_REQUEST_ERROR_MESSAGE =
            "Reservation request was invalid.Please provide valid request<Requestor Name> <Number of Tickets>!";
        private const string THEATER_LAYOUT_CREATION_ERROR_MESSAGE = "Unable to create theater layout at this time.Please try back again later!";
        private const string RESERVATION_REQUEST_INFLATION_ERROR_MESSAGE = "Reservation request validation error.Please try back again later!";
        private const string THEATER_SEATING_LAYOUT_UNAVAILABLE_ERROR_MESSAGE = "Theater seating layout unavailable.Please ensure that the seating layout has been mapped!";
        private const string RESERVATION_ERROR_MESSAGE = "Unable to reserve seats at this time.Please try back again later!";
        private const string SPLIT_RESERVATION_PARTY_MESSAGE = "Call to split party!";
        private const string UNABLE_TO_HANDLE_RESERVATION_MESSAGE = "Sorry, we can't handle your party!";

        /// <summary>
        /// Constructor
        /// </summary>
        public SeatingReservationBc()
        {
            _theaterSeatingRepository = new TheaterSeatingRepository();
            _reservationRepository = new ReservationRepository();
        }

        /// <summary>
        /// Creates and Maps the theater seating layout
        /// </summary>
        /// <param name="seatingDetails"></param>
        public void MapTheatreSeatingLayout(List<string> seatingDetails)
        {
            List<SectionSeating> sectionSeatings = new List<SectionSeating>();

            // Validate if the theater seating layout is valid, if not throw exception
            if (seatingDetails == null || seatingDetails.Count == 0 ||
                seatingDetails.Any(r => string.IsNullOrWhiteSpace(r) || !VALID_SEATING_ROW_PATTERN.IsMatch(r)))
            {
                throw new InvalidInputException(INVALID_THEATER_SEATING_ERROR_MESSAGE);
            }

            try
            {
                // Run through the seating layout and create the layout map
                int rowNumber = 1;
                foreach (string seatingRowDetail in seatingDetails)
                {
                    string cleanedSeatingRowDetail = Regex.Replace(seatingRowDetail, @"\s+", " ");
                    List<int> sections = cleanedSeatingRowDetail.Split(' ').Select(Int32.Parse).ToList();

                    if (sections.Any(s => s <= 0))
                    {
                        break;
                    }
                    int sectionNumber = 1;
                    foreach (int section in sections)
                    {
                        SectionSeating sectionSeating = new SectionSeating();
                        sectionSeating.RowNumber = rowNumber;
                        sectionSeating.SectionNumber = sectionNumber;
                        sectionSeating.NumderOfSeatsInSection = section;
                        sectionSeatings.Add(sectionSeating);
                        sectionNumber++;
                    }

                    rowNumber++;
                }

                // create the layout map
                _theaterSeatingRepository.Create(sectionSeatings);
            }

            // Throw exception in case of any erro during layout creation
            catch (Exception)
            {
                throw new GeneralException(THEATER_LAYOUT_CREATION_ERROR_MESSAGE);
            }
        }

        private List<TicketReservationRequest> ValidateAndInflateTicketingRequest(List<string> ticketRequests)
        {
            List<TicketReservationRequest> ticketingRequests = null;

            try
            {
                // Validate reservation request and if invalid throw exception
                if (ticketRequests == null || ticketRequests.Count == 0 ||
                    ticketRequests.Any(r => string.IsNullOrWhiteSpace(r) || !VALID_TICKETING_REQUEST_PATTERN.IsMatch(r)))
                {
                    return ticketingRequests;
                }

                // Inflate reservation request
                ticketingRequests = new List<TicketReservationRequest>();
                foreach (string ticketRequest in ticketRequests)
                {
                    string cleanedTicketRequest = Regex.Replace(ticketRequest, @"\s+", " ");
                    string[] ticketRequestDetail = cleanedTicketRequest.Split(' ');
                    TicketReservationRequest ticketingRequest = new TicketReservationRequest();
                    ticketingRequest.RequestorName = ticketRequestDetail[0];
                    ticketingRequest.NumberOfTicketsRequested = int.Parse(ticketRequestDetail[1]);
                    ticketingRequests.Add(ticketingRequest);
                }
            }
            catch (Exception ex)
            {
                throw new GeneralException(RESERVATION_REQUEST_INFLATION_ERROR_MESSAGE);
            }

            return ticketingRequests;
        }

        /// <summary>
        /// Process the reservations request
        /// </summary>
        /// <param name="reservationRequests"></param>
        /// <returns></returns>
        public List<ReservationInformation> ProcessSeatingReservations(List<string> reservationRequests)
        {
            List<TicketReservationRequest> ticketReservationRequests = null;
            List<ReservationInformation> reservations = null;
            List<SectionSeating> theaterSeatings = null;

            // Retrieve the theater seating layout
            theaterSeatings = (List<SectionSeating>)_theaterSeatingRepository.GetAll();

            // Throw exception if theater seating is not available
            if (theaterSeatings == null || theaterSeatings.Count == 0)
            {
                throw new GeneralException(THEATER_SEATING_LAYOUT_UNAVAILABLE_ERROR_MESSAGE);
            }

            // Validate & inflate reservation requests
            ticketReservationRequests = ValidateAndInflateTicketingRequest(reservationRequests);

            // Throw exception in case reservations request(s) are invalid
            if (ticketReservationRequests == null || ticketReservationRequests.Count == 0)
            {
                throw new InvalidInputException(INVALID_RESERVATION_REQUEST_ERROR_MESSAGE);
            }

            try
            {
                // Evaluate the largest avaibale seating in a section
                int largestAvailableSectionForReservation = theaterSeatings.Where(s => !s.IsSectionFullyOccupied).Max(s => s.NumberOfAvailableSeats);
                
                // Evaluate the number of seats available for reservation 
                int totalAvailableSeatsForReservation = theaterSeatings.Where(s => !s.IsSectionFullyOccupied).Sum(s => s.NumberOfAvailableSeats);

                // Loop through the reservation requests
                reservations = new List<ReservationInformation>();
                foreach (TicketReservationRequest ticketingRequest in ticketReservationRequests)
                {
                    ReservationInformation reservationInfo = new ReservationInformation();
                    
                    // Check seats are available for the reservation request, if not set the appropriate status message
                    if (ticketingRequest.NumberOfTicketsRequested > totalAvailableSeatsForReservation)
                    {
                        reservationInfo.ReservationStatusMessage = UNABLE_TO_HANDLE_RESERVATION_MESSAGE;
                    }
                    // Check if the reservation request can be accomodated in a single section, if not set the appropriate status message
                    else if (ticketingRequest.NumberOfTicketsRequested > largestAvailableSectionForReservation)
                    {
                        reservationInfo.ReservationStatusMessage = SPLIT_RESERVATION_PARTY_MESSAGE;
                    }
                    else
                    {
                        // Find the section first availabe section which can accomodate the reservation request which is closest to the front
                        SectionSeating sectionReservable = theaterSeatings.FindAll(s => 
                        s.NumberOfAvailableSeats >= ticketingRequest.NumberOfTicketsRequested).OrderBy(s => s.RowNumber).ThenBy(s => s.SectionNumber).FirstOrDefault();

                        // Updated the section with the appropriate resrvation information
                        if (sectionReservable != null)
                        {
                            sectionReservable.NumberOfSeatsReserved += ticketingRequest.NumberOfTicketsRequested;
                            reservationInfo.ReservationName = ticketingRequest.RequestorName;
                            reservationInfo.NumberOfSeatsReserved = ticketingRequest.NumberOfTicketsRequested;
                            reservationInfo.ReservedRowNumber = sectionReservable.RowNumber;
                            reservationInfo.ReservedSectionNumber = sectionReservable.SectionNumber;
                        }
                        // If no sections is available for reservation, set the status message accordingly
                        else
                        {
                            reservationInfo.ReservationStatusMessage = UNABLE_TO_HANDLE_RESERVATION_MESSAGE;
                        }
                    }

                    reservations.Add(reservationInfo);
                    // Re-evaluate the largest avaibale seating in a section, due to reservation updates
                    largestAvailableSectionForReservation = theaterSeatings.Where(s => !s.IsSectionFullyOccupied).Max(s => s.NumberOfAvailableSeats);
                    // Re-valuate the number of seats available for reservation 
                    totalAvailableSeatsForReservation = theaterSeatings.Where(s => !s.IsSectionFullyOccupied).Sum(s => s.NumberOfAvailableSeats);
                }
            }
            // Throw exception in case of any error during processing
            catch (Exception ex)
            {
                throw new GeneralException(UNABLE_TO_HANDLE_RESERVATION_MESSAGE);
            }

            // Confirm the reservations by updating the data repository
            _theaterSeatingRepository.UpdateAll(theaterSeatings);
            _reservationRepository.UpdateAll(reservations);

            // return the confirmed reservations
            return reservations;
        }
    }
}
