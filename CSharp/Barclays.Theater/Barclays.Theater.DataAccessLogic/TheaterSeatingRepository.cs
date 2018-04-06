using System;
using System.IO;
using System.Collections.Generic;
using Newtonsoft.Json;
using Barclays.Theater.DataAccessLogic.Interface;
using Barclays.Theater.Entities;

namespace Barclays.Theater.DataAccessLogic
{
    public class TheaterSeatingRepository : IRepository<SectionSeating>
    {

        private const string THEATER_SEATING_DATA_STORE_NAME = @"TheaterSeating.json";
        /// <summary>
        /// Create theater seating repository
        /// </summary>
        /// <param name="seating"></param>
        public void Create(IEnumerable<SectionSeating> seating)
        {
            if (File.Exists(THEATER_SEATING_DATA_STORE_NAME))
                File.Delete(THEATER_SEATING_DATA_STORE_NAME);
            File.WriteAllText(THEATER_SEATING_DATA_STORE_NAME, JsonConvert.SerializeObject((List<SectionSeating>)seating, Formatting.Indented));
        }

        // Update theater seating
        public void UpdateAll(IEnumerable<SectionSeating> seating)
        {
            if (File.Exists(THEATER_SEATING_DATA_STORE_NAME))
                File.WriteAllText(THEATER_SEATING_DATA_STORE_NAME, String.Empty); 
            File.WriteAllText(THEATER_SEATING_DATA_STORE_NAME, JsonConvert.SerializeObject((List<SectionSeating>)seating, Formatting.Indented));
        }

        // Get all theater seating information
        public IEnumerable<SectionSeating> GetAll()
        {
            List<SectionSeating> theaterSeating = 
                JsonConvert.DeserializeObject<List<SectionSeating>>(File.ReadAllText(THEATER_SEATING_DATA_STORE_NAME));
            return theaterSeating;
        }

        /// <summary>
        /// Delete
        /// </summary>
        /// <param name="seating"></param>
        public void Delete(SectionSeating seating)
        {
            throw new NotImplementedException();
        }
    }
}
