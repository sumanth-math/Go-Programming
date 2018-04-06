using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Barclays.Theater.Entities
{
    /// <summary>
    /// This entity holds the reservation information
    /// </summary>
    public class ReservationInformation
    {
        /// <summary>
        /// Name of the person making the reservation
        /// </summary>
        public string ReservationName
        {
            get;
            set;
        }

        /// <summary>
        /// Number of Seats reserved
        /// </summary>
        public int NumberOfSeatsReserved
        {
            get;
            set;
        }

        // Reserved seat row number
        public int ReservedRowNumber
        {
            get;
            set;
        }

        // Reserved seat section number
        public int ReservedSectionNumber
        {
            get;
            set;
        }

        // Reservation status message if unable to proceed with reservation
        public string ReservationStatusMessage
        {
            get;
            set;
        }
    }
}
