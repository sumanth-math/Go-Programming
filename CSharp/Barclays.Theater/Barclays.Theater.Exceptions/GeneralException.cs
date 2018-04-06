using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Barclays.Theater.Exceptions
{
    public class GeneralException : Exception
    {
        /// <summary>
        /// GeneralException Constructor
        /// </summary>
        /// <param name="message">Exception Message</param>
        public GeneralException(string message) : base(message)
        {

        }

        /// <summary>
        /// GeneralException Constructor
        /// </summary>
        /// <param name="message">message</param>
        /// <param name="innerException">innerException</param>
        public GeneralException(string message, Exception innerException) : base(message, innerException)
        {

        }
    }
}
