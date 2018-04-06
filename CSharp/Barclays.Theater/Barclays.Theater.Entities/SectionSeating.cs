using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Barclays.Theater.Entities
{
    /// <summary>
    /// This entity holds the information regarding the theater seating layout
    /// </summary>
    public class SectionSeating
    {
        /// <summary>
        /// Seating Row number
        /// </summary>
        public int RowNumber
        {
            get;
            set;
        }

        /// <summary>
        /// Seating Section number
        /// </summary>
        public int SectionNumber
        {
            get;
            set;
        }

        /// <summary>
        /// Number of seats available in the Section
        /// </summary>
        public int NumderOfSeatsInSection
        {
            get;
            set;
        }

        /// <summary>
        /// Number of seats reserved in the section (if any)
        /// </summary>
        public int NumberOfSeatsReserved
        {
            get;
            set;
        }

        /// <summary>
        /// Number of seats still available in the section for reservation
        /// </summary>
        public int NumberOfAvailableSeats
        {
            get
            {
                return (NumderOfSeatsInSection - NumberOfSeatsReserved);
            }
        }

        /// <summary>
        /// Whether the section seating is full or not
        /// </summary>
        public bool IsSectionFullyOccupied
        {
            get
            {
                if (NumderOfSeatsInSection == NumberOfSeatsReserved)
                    return true;
                else
                    return false;
            }
        }
    }
}
