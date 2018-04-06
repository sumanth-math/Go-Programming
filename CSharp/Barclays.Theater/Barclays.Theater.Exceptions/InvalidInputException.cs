using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Barclays.Theater.Exceptions
{
    public class InvalidInputException : Exception
    {
        /// <summary>
        /// InvalidInputException Constructor
        /// </summary>
        /// <param name="message">Exception Message</param>
        public InvalidInputException(string message)
            : base(message)
        {

        }

        /// <summary>
        /// InvalidInputException Constructor
        /// </summary>
        /// <param name="message">message</param>
        /// <param name="innerException">innerException</param>
        public InvalidInputException(string message, Exception innerException)
            : base(message, innerException)
        {

        }
    }
}
