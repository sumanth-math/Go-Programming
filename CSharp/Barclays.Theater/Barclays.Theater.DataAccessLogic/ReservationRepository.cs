using System;
using System.IO;
using System.Collections.Generic;
using Newtonsoft.Json;
using Barclays.Theater.DataAccessLogic.Interface;
using Barclays.Theater.Entities;

namespace Barclays.Theater.DataAccessLogic
{
    public class ReservationRepository : IRepository<ReservationInformation>
    {
        private const string RESERVATION_DATA_STORE_NAME = @"Reservation.json";

        /// <summary>
        /// Create reservation repository
        /// </summary>
        /// <param name="reservations"></param>
        public void Create(IEnumerable<ReservationInformation> reservations)
        {
            // Store resrvation information to a file data store in json format
            File.WriteAllText(RESERVATION_DATA_STORE_NAME, JsonConvert.SerializeObject((List<SectionSeating>)reservations, Formatting.Indented));
        }

        /// <summary>
        /// Update reservation information
        /// </summary>
        /// <param name="reservations"></param>
        public void UpdateAll(IEnumerable<ReservationInformation> reservations)
        {
            // Update resrvation information to file data store in json format
            if (File.Exists(RESERVATION_DATA_STORE_NAME))
                File.WriteAllText(RESERVATION_DATA_STORE_NAME, String.Empty);
            File.WriteAllText(RESERVATION_DATA_STORE_NAME, JsonConvert.SerializeObject((List<ReservationInformation>)reservations, Formatting.Indented));
        }

        /// <summary>
        /// Get all reservation information
        /// </summary>
        /// <returns></returns>
        public IEnumerable<ReservationInformation> GetAll()
        {
            // Get all the reservation information from the file data store
            List<ReservationInformation> reservations =
                JsonConvert.DeserializeObject<List<ReservationInformation>>(File.ReadAllText(RESERVATION_DATA_STORE_NAME));
            return reservations;
        }

        public void Delete(ReservationInformation seating)
        {
            throw new NotImplementedException();
        }
    }
}
