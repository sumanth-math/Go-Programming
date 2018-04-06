using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Barclays.Theater.BusinessLogic;
using Barclays.Theater.Entities;

namespace Barclays.Theater.Application
{
    class Program
    {
        static void Main(string[] args)
        {
            List<string> theatreSeatingRowDetails = new List<string>();
            List<string> reservationRequests = new List<string>();
            List<ReservationInformation> reservations = null;

            try
            {
                SeatingReservationBc theatreSeatingReservation = new SeatingReservationBc();

                string option = string.Empty;

                SelectMenuOptions();
                while (!string.IsNullOrWhiteSpace(option = Console.ReadLine()))
                {
                    switch (option)
                    {
                        case "1":
                            {
                                Console.WriteLine("Provide the Theatre Seating Arrangement (Press 'Enter' to complete the seating arrangement):");
                                Console.WriteLine("[Specify the seats in each section of the row");
                                Console.WriteLine("Ex:- 3 6 3 meaing a row containing 3 sections with 3, 6 and 3 seats in each section respectively]");

                                string seatingRowDetail = string.Empty;
                                while (!string.IsNullOrWhiteSpace(seatingRowDetail = Console.ReadLine()))
                                {
                                    theatreSeatingRowDetails.Add(seatingRowDetail);
                                }

                                try
                                {
                                    theatreSeatingReservation.MapTheatreSeatingLayout(theatreSeatingRowDetails);
                                }
                                catch (Exception ex)
                                {
                                    Console.WriteLine(ex.Message);
                                }
                                break;
                            }
                        case "2":
                            {
                                Console.WriteLine("Provide Ticket Request (Press 'Enter' to complete the request):");
                                Console.WriteLine("[Name of the person buying the ticket followed by # of Tickets Ex:- Smith 5]");
                                string reservationRequest = string.Empty;
                                while (!string.IsNullOrWhiteSpace(reservationRequest = Console.ReadLine()))
                                {
                                    reservationRequests.Add(reservationRequest);
                                }

                                try
                                {
                                    reservations = theatreSeatingReservation.ProcessSeatingReservations(reservationRequests);
                                }
                                catch(Exception ex)
                                {
                                    Console.WriteLine(ex.Message);
                                }
                                break;
                            }
                        case "3":
                            {
                                if (reservations != null && reservations.Count > 0)
                                {
                                    Console.WriteLine("********* Ticketing Reservations ************");
                                    foreach (ReservationInformation reservationInfo in reservations)
                                    {
                                        if (reservationInfo != null && string.IsNullOrWhiteSpace(reservationInfo.ReservationStatusMessage))
                                        {
                                            Console.WriteLine("{0} : {1} seat(s) reserved - Row {2} Section {3}", reservationInfo.ReservationName,
                                                reservationInfo.NumberOfSeatsReserved, reservationInfo.ReservedRowNumber, reservationInfo.ReservedSectionNumber);
                                        }
                                        else
                                        {
                                            Console.WriteLine("{0} {1}", reservationInfo.ReservationName, reservationInfo.ReservationStatusMessage);
                                        }

                                    }
                                }
                                else
                                {
                                    Console.WriteLine("No Reservations found at this time");
                                }
                                break;
                            }
                        case "4":
                            {
                                goto ApplicationExit;
                            }
                    }

                    option = string.Empty;
                    SelectMenuOptions();
                }

                ApplicationExit:
                ;
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
                Console.ReadLine();
            }
        }

        static void SelectMenuOptions()
        {
            // Select Menu Options
            Console.WriteLine("Please choose from the below menu options:");
            Console.WriteLine("Select 1 for Theatre Layout");
            Console.WriteLine("Select 2 for Reservations");
            Console.WriteLine("Select 3 for Seating Display");
            Console.WriteLine("Select 4 for Exit");
        }
    }
}
