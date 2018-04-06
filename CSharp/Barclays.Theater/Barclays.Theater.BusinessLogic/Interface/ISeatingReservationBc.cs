using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Barclays.Theater.Entities;

namespace Barclays.Theater.BusinessLogic.Interface
{
    interface ISeatingReservationBc
    {
        /// <summary>
        /// Maps theater seating layout
        /// </summary>
        /// <param name="seatingDetails"></param>
        /// <returns>Seating layout mapping was successful or not</returns>
        void MapTheatreSeatingLayout(List<string> seatingDetails);

        /// <summary>
        /// Processes the reservations for the given ticketing requests
        /// </summary>
        /// <param name="ticketingRequests"></param>
        /// <returns></returns>
        List<ReservationInformation> ProcessSeatingReservations(List<string> reservationRequests);
    }
}
