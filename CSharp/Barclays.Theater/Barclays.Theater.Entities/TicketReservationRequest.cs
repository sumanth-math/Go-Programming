using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Barclays.Theater.Entities
{
    /// <summary>
    /// This entity holds the information regarding the reservation request
    /// </summary>
    public class TicketReservationRequest
    {
        /// <summary>
        /// Name of the person requesting the reservation
        /// </summary>
        public string RequestorName
        {
            get;
            set;
        }

        /// <summary>
        /// Number of tickets being requested for
        /// </summary>
        public int NumberOfTicketsRequested
        {
            get;
            set;
        }
    }
}
